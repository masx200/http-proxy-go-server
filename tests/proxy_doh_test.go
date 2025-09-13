package tests

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

// logCommand 记录命令执行到文件
func logCommand(cmd *exec.Cmd, cmdType string) error {
	cmdStr := strings.Join(cmd.Args, " ")

	entry := fmt.Sprintf("[%s] [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		cmdType,
		cmdStr)

	return appendToFile("command_execution_log.md", entry)
}

// logCommandResult 记录命令执行结果
func logCommandResult(cmd *exec.Cmd, err error, output string) error {
	cmdStr := strings.Join(cmd.Args, " ")
	result := "成功"
	if err != nil {
		result = "失败"
	}

	entry := fmt.Sprintf("执行结果: %s\n进程PID: %d\n执行时间: %s\n输出: %s\n错误: %s\n---\n",
		result,
		cmd.Process.Pid,
		time.Now().Format("2006-01-02 15:04:05"),
		output,
		errToString(err))

	return appendToFile("command_execution_log.md", entry)
}

// errToString 将错误转换为字符串
func errToString(err error) string {
	if err == nil {
		return "无"
	}
	return err.Error()
}

// appendToFile 追加内容到文件
func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
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
	_, err = writer.WriteString("\n\n###\n\n")
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

// runProxyServerDOH 测试带DoH的HTTP代理服务器的基本功能
func runProxyServerDOH(t *testing.T) {
	// 创建进程管理器
	processManager := NewProcessManager()
	defer processManager.CleanupAll()

	// 先编译代理服务器
	var testResults []string
	testResults = append(testResults, "# DoH HTTP代理服务器测试")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `go build -o main.exe ../cmd/main.go`")
	testResults = append(testResults, "")

	// 先编译代理服务器
	testResults = append(testResults, "编译代理服务器...")
	buildCmd := exec.Command("go", "build", "-o", "main.exe", "../cmd/main.go")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	// 记录命令执行
	if err := logCommand(buildCmd, "BUILD"); err != nil {
		log.Printf("记录命令失败: %v", err)
	}

	if err := buildCmd.Run(); err != nil {
		logCommandResult(buildCmd, err, "")
		t.Fatalf("编译代理服务器失败: %v", err)
	}
	logCommandResult(buildCmd, nil, "")

	testResults = append(testResults, "✅ 代理服务器编译成功")
	testResults = append(testResults, "")

	// 启动代理服务器进程（使用编译后的可执行文件，带DoH参数）
	testResults = append(testResults, "启动代理服务器（DoH模式）...")
	testResults = append(testResults, "执行命令: `./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohalpn h2 -dohalpn h3`")
	testResults = append(testResults, "")

	cmd := exec.Command("./main.exe", "-dohurl", "https://dns.alidns.com/dns-query", "-dohip", "223.5.5.5", "-dohip", "223.6.6.6", "-dohurl", "https://dns.alidns.com/dns-query", "-dohalpn", "h2", "-dohalpn", "h3")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 记录命令执行
	if err := logCommand(cmd, "SERVER"); err != nil {
		log.Printf("记录命令失败: %v", err)
	}

	// 将进程添加到进程管理器
	processManager.AddProcess(cmd)

	// 启动代理服务器
	if err := cmd.Start(); err != nil {
		logCommandResult(cmd, err, "")
		t.Fatalf("启动代理服务器失败: %v", err)
	}

	// 记录启动结果
	logCommandResult(cmd, nil, "")

	// 等待代理服务器启动
	time.Sleep(2 * time.Second)

	// 检查代理服务器是否启动成功
	if !isProxyServerRunning() {
		t.Fatal("代理服务器未正常启动")
	}

	testResults = append(testResults, "✅ 代理服务器启动成功")
	testResults = append(testResults, "")
	testResults = append(testResults, "开始测试代理功能...")
	testResults = append(testResults, "")

	// 检查8080端口是否被占用
	if !isPortOccupied(8080) {
		testResults = append(testResults, "❌ 8080端口未被占用，代理服务器可能未正确启动")
	} else {
		testResults = append(testResults, "✅ 8080端口被占用，代理服务器运行正常")
	}

	// 测试1：通过代理访问HTTP网站
	testResults = append(testResults, "测试1: 通过代理访问HTTP网站 (baidu.com)")
	curlCmd1 := exec.Command("curl", "-v", "-I", "http://www.baidu.com", "-x", "http://localhost:8080")

	var curlOutput1 bytes.Buffer
	curlCmd1.Stdout = &curlOutput1
	curlCmd1.Stderr = &curlOutput1

	// 记录命令执行
	if err := logCommand(curlCmd1, "TEST"); err != nil {
		log.Printf("记录命令失败: %v", err)
	}

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd1)

	// 启动curl进程
	err1 := curlCmd1.Run()
	output1 := curlOutput1.Bytes()

	if err1 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试1失败: %v", err1))
	} else {
		testResults = append(testResults, "✅ 测试1成功")
	}

	// 记录命令执行结果
	if err := logCommandResult(curlCmd1, err1, string(output1)); err != nil {
		log.Printf("记录命令结果失败: %v", err)
	}
	testResults = append(testResults, "")

	// 测试2：通过代理访问HTTP网站（跟随重定向）
	testResults = append(testResults, "测试2: 通过代理访问HTTP网站 (so.com，跟随重定向)")
	curlCmd2 := exec.Command("curl", "-v", "-I", "-L", "http://www.so.com", "-x", "http://localhost:8080")

	var curlOutput2 bytes.Buffer
	curlCmd2.Stdout = &curlOutput2
	curlCmd2.Stderr = &curlOutput2

	// 记录命令执行
	if err := logCommand(curlCmd2, "TEST"); err != nil {
		log.Printf("记录命令失败: %v", err)
	}

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd2)

	// 启动curl进程
	err2 := curlCmd2.Run()
	output2 := curlOutput2.Bytes()

	if err2 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试2失败: %v", err2))
	} else {
		testResults = append(testResults, "✅ 测试2成功")
	}

	// 记录命令执行结果
	if err := logCommandResult(curlCmd2, err2, string(output2)); err != nil {
		log.Printf("记录命令结果失败: %v", err)
	}
	testResults = append(testResults, "")

	// 测试3：通过代理访问HTTPS网站
	testResults = append(testResults, "测试3: 通过代理访问HTTPS网站 (baidu.com)")
	curlCmd3 := exec.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "http://localhost:8080")

	var curlOutput3 bytes.Buffer
	curlCmd3.Stdout = &curlOutput3
	curlCmd3.Stderr = &curlOutput3

	// 记录命令执行
	if err := logCommand(curlCmd3, "TEST"); err != nil {
		log.Printf("记录命令失败: %v", err)
	}

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd3)

	// 启动curl进程
	err3 := curlCmd3.Run()
	output3 := curlOutput3.Bytes()

	if err3 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试3失败: %v", err3))
	} else {
		testResults = append(testResults, "✅ 测试3成功")
	}

	// 记录命令执行结果
	if err := logCommandResult(curlCmd3, err3, string(output3)); err != nil {
		log.Printf("记录命令结果失败: %v", err)
	}
	testResults = append(testResults, "")

	// 将测试结果写入文件
	if err := writeTestResults(testResults); err != nil {
		log.Printf("写入测试结果失败: %v\n", err)
	}

	// 等待一秒，确保所有输出都被捕获
	time.Sleep(1 * time.Second)
}

// TestMainDOH 主测试函数
func TestMainDOH(t *testing.T) {
	// 创建带有30秒超时的上下文（增加超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan int, 1)

	// 设置全局变量，让测试函数能够访问进程管理器
	var globalProcessManager *ProcessManager

	// 在goroutine中运行测试
	go func() {
		// 运行测试
		runProxyServerDOH(t)
		resultChan <- 0
	}()

	// 等待测试完成或超时
	select {
	case <-resultChan:
		// 测试正常完成
		return //os.Exit(code)
	case <-ctx.Done():
		// 超时或取消
		log.Println("\n⏰ 测试超时（30秒），强制退出...")

		// 强制终止所有记录的进程
		log.Println("正在终止所有运行中的进程...")

		// 在Windows上强制终止所有go进程和可能的子进程
		if runtime.GOOS == "windows" {
			// 使用taskkill终止所有go进程
			killCmd := exec.Command("taskkill", "/F", "/IM", "go.exe")
			if err := logCommand(killCmd, "SYSTEM"); err != nil {
				log.Printf("记录命令失败: %v", err)
			}
			killCmd.Run() // 忽略错误
			if err := logCommandResult(killCmd, nil, ""); err != nil {
				log.Printf("记录命令结果失败: %v", err)
			}

			// 终止可能的代理服务器进程（在8080端口上）
			findCmd := exec.Command("netstat", "-ano", "|", "findstr", ":8080")
			if err := logCommand(findCmd, "SYSTEM"); err != nil {
				log.Printf("记录命令失败: %v", err)
			}
			findCmd.Run() // 忽略错误
			if err := logCommandResult(findCmd, nil, ""); err != nil {
				log.Printf("记录命令结果失败: %v", err)
			}
		}

		// 清理全局进程管理器中的进程
		if globalProcessManager != nil {
			globalProcessManager.CleanupAll()
		}

		// 记录超时信息到测试记录
		timeoutMessage := []string{
			"# 测试超时记录",
			"",
			"## 超时时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"❌ 测试执行超过30秒超时限制，强制退出",
			"",
			"可能的原因:",
			"- 代理服务器进程未正常退出",
			"- curl命令卡住",
			"- 网络连接问题",
			"- 日志输出缓冲问题",
			"",
			"已尝试终止所有相关进程",
			"",
		}

		// 写入超时记录
		if err := writeTestResults(timeoutMessage); err != nil {
			log.Printf("写入超时记录失败: %v\n", err)
		}

		// 强制退出
		t.Fatal("测试超时") //os.Exit(1)
	}
}
