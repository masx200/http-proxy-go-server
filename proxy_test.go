package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// ProcessManager 进程管理器
type ProcessManager struct {
	processes []*exec.Cmd
	mutex     sync.Mutex
}

// NewProcessManager 创建新的进程管理器
func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make([]*exec.Cmd, 0),
	}
}

// AddProcess 添加进程到管理器
func (pm *ProcessManager) AddProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.processes = append(pm.processes, cmd)
}

// CleanupAll 清理所有进程
func (pm *ProcessManager) CleanupAll() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			cmd.Process.Kill()
			cmd.Wait()
		}
	}
	pm.processes = make([]*exec.Cmd, 0)
}

// GetPIDs 获取所有进程的PID
func (pm *ProcessManager) GetPIDs() []string {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	var pids []string
	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			pids = append(pids, strconv.Itoa(cmd.Process.Pid))
		}
	}
	return pids
}

// TestProxyServer 测试HTTP代理服务器的基本功能
func TestProxyServer(t *testing.T) {
	// 创建进程管理器
	processManager := NewProcessManager()
	defer processManager.CleanupAll()

	// 创建缓冲区来捕获代理服务器的输出
	var proxyOutput bytes.Buffer

	// 创建一个多写入器，同时写入到标准输出和缓冲区
	multiWriter := io.MultiWriter(os.Stdout, &proxyOutput)

	// 添加测试超时检查
	timeoutTimer := time.AfterFunc(18*time.Second, func() {
		fmt.Println("\n⚠️ 测试即将超时，正在清理进程...")
		// 在超时前记录代理服务器日志
		var timeoutTestResults []string
		if proxyOutput.Len() > 0 {
			timeoutTestResults = []string{
				"# HTTP代理服务器测试记录（超时）",
				"",
				"## 测试时间",
				time.Now().Format("2006-01-02 15:04:05"),
				"",
				"## 代理服务器日志输出（超时前捕获）",
				"",
				"```",
			}
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(proxyOutput.String(), "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					timeoutTestResults = append(timeoutTestResults, line)
				}
			}
			timeoutTestResults = append(timeoutTestResults, "```")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "❌ 测试超时，但已捕获代理服务器日志")
		} else {
			// 即使没有输出，也要记录超时信息
			timeoutTestResults = []string{
				"# HTTP代理服务器测试记录（超时）",
				"",
				"## 测试时间",
				time.Now().Format("2006-01-02 15:04:05"),
				"",
				"## 代理服务器状态",
				"",
				"⚠️ 代理服务器没有产生任何输出",
				"",
				"❌ 测试超时",
			}
		}

		// 调试信息：将proxyOutput状态添加到测试记录
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, "## 调试信息")
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, fmt.Sprintf("[DEBUG] proxyOutput长度: %d", proxyOutput.Len()))
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, "[DEBUG] proxyOutput内容:")
		timeoutTestResults = append(timeoutTestResults, "```")
		timeoutTestResults = append(timeoutTestResults, proxyOutput.String())
		timeoutTestResults = append(timeoutTestResults, "```")

		// 写入超时测试记录
		if err := writeTestResults(timeoutTestResults); err != nil {
			fmt.Printf("写入超时测试记录失败: %v\n", err)
		}
		processManager.CleanupAll()
	})
	defer timeoutTimer.Stop()

	// 测试结果记录
	var testResults []string
	testResults = append(testResults, "# HTTP代理服务器测试记录")
	testResults = append(testResults, "")
	testResults = append(testResults, "## 测试时间")
	testResults = append(testResults, time.Now().Format("2006-01-02 15:04:05"))
	testResults = append(testResults, "")

	// 检查端口是否被占用
	if isPortOccupied(8080) {
		t.Fatal("端口8080已被占用，请先停止占用该端口的进程")
	}

	// 启动代理服务器
	testResults = append(testResults, "## 1. 启动代理服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `go run -v ./main.go`")
	testResults = append(testResults, "")

	// 启动代理服务器进程
	cmd := exec.Command("go", "run", "-v", "./main.go")
	// 设置代理服务器的输出到多写入器，同时输出到控制台和缓冲区
	cmd.Stdout = multiWriter
	cmd.Stderr = multiWriter
	err := cmd.Start()
	if err != nil {
		t.Fatalf("启动代理服务器失败: %v", err)
	}

	// 将代理服务器进程添加到管理器
	processManager.AddProcess(cmd)

	// 记录代理服务器PID
	testResults = append(testResults, fmt.Sprintf("📋 代理服务器进程PID: %d", cmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待服务器启动
	testResults = append(testResults, "等待服务器启动...")
	time.Sleep(3 * time.Second)

	// 检查服务器是否正常启动
	if !isProxyServerRunning() {
		t.Fatal("代理服务器启动失败")
	}

	testResults = append(testResults, "✅ 代理服务器启动成功")
	testResults = append(testResults, "")

	// 测试HTTP代理功能
	testResults = append(testResults, "## 2. 测试HTTP代理功能")
	testResults = append(testResults, "")

	// 第一个curl测试
	testResults = append(testResults, "### 测试1: 基本HTTP代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I http://www.baidu.com -x http://localhost:8080`")
	testResults = append(testResults, "")

	// 创建curl进程
	curlCmd1 := exec.Command("curl", "-v", "-I", "http://www.baidu.com", "-x", "http://localhost:8080")
	// 创建缓冲区来捕获curl输出
	var curlOutput1 bytes.Buffer
	curlCmd1.Stdout = &curlOutput1
	curlCmd1.Stderr = &curlOutput1

	// 启动curl进程
	err1 := curlCmd1.Run()
	output1 := curlOutput1.Bytes()

	// 将curl进程添加到管理器
	processManager.AddProcess(curlCmd1)

	// 记录curl进程PID
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试1进程PID: %d", curlCmd1.Process.Pid))
	testResults = append(testResults, "")
	if err1 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err1))
		testResults = append(testResults, fmt.Sprintf("错误输出: %s", string(output1)))
	} else {
		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, "输出结果:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(output1))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 第二个curl测试（重复测试）
	testResults = append(testResults, "### 测试2: HTTP代理www.so.com")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I http://www.so.com -x http://localhost:8080`")
	testResults = append(testResults, "")

	// 创建curl进程
	curlCmd2 := exec.Command("curl", "-v", "-I", "http://www.so.com", "-x", "http://localhost:8080")
	// 创建缓冲区来捕获curl输出
	var curlOutput2 bytes.Buffer
	curlCmd2.Stdout = &curlOutput2
	curlCmd2.Stderr = &curlOutput2

	// 启动curl进程
	err2 := curlCmd2.Run()
	output2 := curlOutput2.Bytes()

	// 将curl进程添加到管理器
	processManager.AddProcess(curlCmd2)

	// 记录curl进程PID
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试2进程PID: %d", curlCmd2.Process.Pid))
	testResults = append(testResults, "")
	if err2 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err2))
		testResults = append(testResults, fmt.Sprintf("错误输出: %s", string(output2)))
	} else {
		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, "输出结果:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(output2))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 测试HTTPS代理功能
	testResults = append(testResults, "### 测试3: HTTPS代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I https://www.baidu.com -x http://localhost:8080`")
	testResults = append(testResults, "")

	// 创建curl进程
	curlCmd3 := exec.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "http://localhost:8080")
	// 创建缓冲区来捕获curl输出
	var curlOutput3 bytes.Buffer
	curlCmd3.Stdout = &curlOutput3
	curlCmd3.Stderr = &curlOutput3

	// 启动curl进程
	err3 := curlCmd3.Run()
	output3 := curlOutput3.Bytes()

	// 将curl进程添加到管理器
	processManager.AddProcess(curlCmd3)

	// 记录curl进程PID
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试3进程PID: %d", curlCmd3.Process.Pid))
	testResults = append(testResults, "")
	if err3 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err3))
		testResults = append(testResults, fmt.Sprintf("错误输出: %s", string(output3)))
	} else {
		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, "输出结果:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(output3))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 记录所有进程PID信息
	testResults = append(testResults, "### 📋 所有进程PID记录")
	testResults = append(testResults, "")
	allPIDs := processManager.GetPIDs()
	testResults = append(testResults, fmt.Sprintf("所有进程PID: %s", strings.Join(allPIDs, ", ")))
	testResults = append(testResults, "")

	// 写入测试记录到文件
	err = writeTestResults(testResults)
	if err != nil {
		t.Errorf("写入测试记录失败: %v", err)
	}

	// 验证测试结果
	if err1 != nil {
		t.Errorf("第一个curl测试失败: %v", err1)
	}
	if err2 != nil {
		t.Errorf("第二个curl测试失败: %v", err2)
	}
	if err3 != nil {
		t.Errorf("HTTPS curl测试失败: %v", err3)
	}

	// 如果curl命令运行成功，关闭代理服务器进程
	if err1 == nil && err2 == nil && err3 == nil {
		testResults = append(testResults, "## 3. 关闭代理服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "✅ 所有curl测试成功，正在关闭代理服务器进程...")
		testResults = append(testResults, "")

		// 停止超时计时器
		timeoutTimer.Stop()

		// 明确终止代理服务器进程
		testResults = append(testResults, "🛑 正在终止代理服务器进程...")
		if cmd.Process != nil {
			if err := cmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止代理服务器进程失败: %v", err))
			} else {
				cmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ 代理服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 清理所有进程
		testResults = append(testResults, "🧹 正在清理所有子进程...")
		testResults = append(testResults, "")
		processManager.CleanupAll()
		testResults = append(testResults, "✅ 所有子进程已清理完成")
		testResults = append(testResults, "")

		// 等待进程完全关闭并释放资源
		time.Sleep(2 * time.Second)

		// 将代理服务器输出添加到测试记录
		if proxyOutput.Len() > 0 {
			testResults = append(testResults, "### 代理服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(proxyOutput.String(), "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		}

		// 将curl进程输出添加到测试记录
		testResults = append(testResults, "### 所有子进程日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "```")

		// 添加curl1输出
		if curlOutput1.Len() > 0 {
			testResults = append(testResults, "--- Curl测试1输出 ---")
			curl1Lines := strings.Split(curlOutput1.String(), "\n")
			for _, line := range curl1Lines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
		}

		// 添加curl2输出
		if curlOutput2.Len() > 0 {
			testResults = append(testResults, "--- Curl测试2输出 ---")
			curl2Lines := strings.Split(curlOutput2.String(), "\n")
			for _, line := range curl2Lines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
		}

		// 添加curl3输出
		if curlOutput3.Len() > 0 {
			testResults = append(testResults, "--- Curl测试3输出 ---")
			curl3Lines := strings.Split(curlOutput3.String(), "\n")
			for _, line := range curl3Lines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
		}

		testResults = append(testResults, "```")
		testResults = append(testResults, "")

		// 验证端口是否已释放
		if !isPortOccupied(8080) {
			testResults = append(testResults, "✅ 端口8080已成功释放")
		} else {
			testResults = append(testResults, "❌ 端口8080仍被占用")
		}

		// 重新写入测试记录
		err = writeTestResults(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}

	} else {
		// 如果有测试失败，也记录关闭进程的信息
		testResults = append(testResults, "## 3. 关闭代理服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "⚠️ 部分测试失败，但仍需关闭代理服务器进程...")
		testResults = append(testResults, "")

		// 明确终止代理服务器进程
		testResults = append(testResults, "🛑 正在终止代理服务器进程...")
		if cmd.Process != nil {
			if err := cmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止代理服务器进程失败: %v", err))
			} else {
				cmd.Wait() // 等待进程完全退出
				testResults = append(testResults, "✅ 代理服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 清理所有进程
		testResults = append(testResults, "🧹 正在清理所有子进程...")
		testResults = append(testResults, "")
		processManager.CleanupAll()
		testResults = append(testResults, "✅ 所有子进程已清理完成")
		testResults = append(testResults, "")

		// 等待进程完全关闭并释放资源
		time.Sleep(2 * time.Second)

		// 将代理服务器输出添加到测试记录
		if proxyOutput.Len() > 0 {
			testResults = append(testResults, "### 代理服务器日志输出")
			testResults = append(testResults, "")
			testResults = append(testResults, "```")
			// 按行分割输出并添加到测试结果
			outputLines := strings.Split(proxyOutput.String(), "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					testResults = append(testResults, line)
				}
			}
			testResults = append(testResults, "```")
			testResults = append(testResults, "")
		}

		// 重新写入测试记录
		err = writeTestResults(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}
	}
}

// isPortOccupied 检查端口是否被占用
func isPortOccupied(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return true
	}
	listener.Close()
	return false
}

// isProxyServerRunning 检查代理服务器是否正在运行
func isProxyServerRunning() bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 创建一个测试请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return false
	}

	// 设置代理
	proxyURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		return false
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client.Transport = transport

	// 发送测试请求
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

// writeTestResults 写入测试结果到文件
func writeTestResults(results []string) error {
	// 写入到测试记录.md
	file, err := os.OpenFile("测试记录.md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 移动到文件末尾
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// 写入分隔符
	_, err = writer.WriteString("\n\n---\n\n")
	if err != nil {
		return err
	}

	// 写入测试结果
	for _, line := range results {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// TestMain 主测试函数
func TestMain(m *testing.M) {
	// 创建带有20秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan int, 1)
	
	// 保存所有运行中的进程，以便在超时时强制终止
	var runningProcesses []*os.Process
	var processMutex sync.Mutex
	
	// 在goroutine中运行测试
	go func() {
		code := m.Run()
		resultChan <- code
	}()

	// 等待测试完成或超时
	select {
	case code := <-resultChan:
		// 测试正常完成
		os.Exit(code)
	case <-ctx.Done():
		// 超时或取消
		fmt.Println("\n⏰ 测试超时（20秒），强制退出...")

		// 强制终止所有记录的进程
		fmt.Println("正在终止所有运行中的进程...")
		processMutex.Lock()
		for _, proc := range runningProcesses {
			if proc != nil {
				proc.Kill()
			}
		}
		processMutex.Unlock()

		// 记录超时信息到测试记录
		timeoutMessage := []string{
			"# 测试超时记录",
			"",
			"## 超时时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"❌ 测试执行超过20秒超时限制，强制退出",
			"",
			"可能的原因:",
			"- 代理服务器进程未正常退出",
			"- curl命令卡住",
			"- 网络连接问题",
			"",
			fmt.Sprintf("已尝试终止 %d 个运行中的进程", len(runningProcesses)),
			"",
		}

		// 写入超时记录
		if err := writeTestResults(timeoutMessage); err != nil {
			fmt.Printf("写入超时记录失败: %v\n", err)
		}

		// 强制退出
		os.Exit(1)
	}
}
