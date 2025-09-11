package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// TestProxyServer 测试HTTP代理服务器的基本功能
func TestProxyServer(t *testing.T) {
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
	// 创建缓冲区来捕获代理服务器的输出
	var proxyOutput bytes.Buffer
	cmd.Stdout = &proxyOutput
	cmd.Stderr = &proxyOutput
	err := cmd.Start()
	if err != nil {
		t.Fatalf("启动代理服务器失败: %v", err)
	}
	// 移除defer语句，改为手动管理进程生命周期

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

	output1, err1 := exec.Command("curl", "-v", "-I", "http://www.baidu.com", "-x", "http://localhost:8080").CombinedOutput()
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

	output2, err2 := exec.Command("curl", "-v", "-I", "http://www.so.com", "-x", "http://localhost:8080").CombinedOutput()
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

	output3, err3 := exec.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "http://localhost:8080").CombinedOutput()
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

		// 立即关闭代理服务器进程
		err = cmd.Process.Kill()
		if err != nil {
			testResults = append(testResults, fmt.Sprintf("❌ 关闭代理服务器失败: %v", err))
		} else {
			testResults = append(testResults, "✅ 代理服务器进程已成功关闭")
		}

		// 等待进程完全关闭并释放资源
		time.Sleep(2 * time.Second)

		// 等待进程完全退出并获取输出
		cmd.Wait()

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

		// 关闭代理服务器进程
		err = cmd.Process.Kill()
		if err != nil {
			testResults = append(testResults, fmt.Sprintf("❌ 关闭代理服务器失败: %v", err))
		} else {
			testResults = append(testResults, "✅ 代理服务器进程已成功关闭")
		}

		// 等待进程完全退出并获取输出
		cmd.Wait()

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
	file, err := os.OpenFile("测试记录.md", os.O_RDWR|os.O_APPEND, 0644)
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
	// 运行测试
	code := m.Run()

	// 退出
	os.Exit(code)
}
