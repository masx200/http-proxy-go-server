package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/net/proxy"
)

// runSOCKS5ProxyServer 测试SOCKS5代理服务器的基本功能，使用 golang.org/x/net/proxy + go-socks5
func runSOCKS5ProxyServer(t *testing.T, logfilename string) {
	var processManager *ProcessManager = NewProcessManager(logfilename)
	defer func() {
		// 清理所有进程
		processManager.CleanupAll()
		processManager.Close()
	}()

	// 创建缓冲区来捕获代理服务器输出
	var proxyOutput bytes.Buffer
	var proxyOutputMutex sync.Mutex

	// 创建一个多写入器，同时写入到标准输出和缓冲区
	multiWriter := io.MultiWriter(os.Stdout, &proxyOutput)

	// 清理可能存在的旧的可执行文件
	if _, err := os.Stat("main.exe"); err == nil {
		os.Remove("main.exe")
	}
	if _, err := os.Stat("socks5-test-server.exe"); err == nil {
		os.Remove("socks5-test-server.exe")
	}

	// 添加测试超时检查
	timeoutTimer := time.AfterFunc(35*time.Second, func() {
		log.Println("\n⚠️ SOCKS5代理测试即将超时，正在清理进程...")
		var timeoutTestResults []string

		// 使用互斥锁保护对proxyOutput的访问
		proxyOutputMutex.Lock()
		outputLen := proxyOutput.Len()
		outputContent := proxyOutput.String()
		proxyOutputMutex.Unlock()

		if outputLen > 0 {
			timeoutTestResults = []string{
				"# SOCKS5代理服务器测试记录（超时）",
				"",
				"## 测试时间",
				time.Now().Format("2006-01-02 15:04:05"),
				"",
				"## 代理服务器日志输出（超时前捕获）",
				"",
				"```",
			}
			outputLines := strings.Split(outputContent, "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					timeoutTestResults = append(timeoutTestResults, line)
				}
			}
			timeoutTestResults = append(timeoutTestResults, "```")
			timeoutTestResults = append(timeoutTestResults, "")
			timeoutTestResults = append(timeoutTestResults, "❌ 测试超时，但已捕获代理服务器日志")
		} else {
			timeoutTestResults = []string{
				"# SOCKS5代理服务器测试记录（超时）",
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

		// 调试信息
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, "## 调试信息")
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, fmt.Sprintf("[DEBUG] proxyOutput长度: %d", outputLen))
		timeoutTestResults = append(timeoutTestResults, "")
		timeoutTestResults = append(timeoutTestResults, "[DEBUG] proxyOutput内容:")
		timeoutTestResults = append(timeoutTestResults, "```")
		timeoutTestResults = append(timeoutTestResults, outputContent)
		timeoutTestResults = append(timeoutTestResults, "```")

		// 写入超时测试记录
		if err := WriteTestResultsToFile(timeoutTestResults, processManager.GetFile()); err != nil {
			log.Printf("写入超时测试记录失败: %v\n", err)
		}
		processManager.CleanupAll()
		t.Fatal("SOCKS5代理测试超时")
	})
	defer timeoutTimer.Stop()

	// 测试结果记录
	var testResults []string
	testResults = append(testResults, "# SOCKS5代理服务器测试记录 (使用 golang.org/x/net/proxy + go-socks5)")
	testResults = append(testResults, "")
	testResults = append(testResults, "## 测试时间")
	testResults = append(testResults, time.Now().Format("2006-01-02 15:04:05"))
	testResults = append(testResults, "")

	// 检查端口是否被占用
	if isPortOccupied(44444) {
		t.Fatal("端口44444已被占用，请先停止占用该端口的进程")
	}

	// 创建测试用的SOCKS5服务器代码
	socks5ServerCode := `package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gitee.com/masx200/go-socks5"
)

func main() {
	// 创建SOCKS5配置
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			&socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					"g7envpwz14b0u55": "juvytdsdzc225pq",
				},
			},
		},
		Rules: socks5.PermitAll(),
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	// 创建SOCKS5服务器
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// 监听端口
	listener, err := net.Listen("tcp", ":44444")
	if err != nil {
		log.Fatalf("Failed to listen on port 44444: %v", err)
	}
	defer listener.Close()

	fmt.Println("SOCKS5 server started on :44444")
	log.Println("SOCKS5 server started on :44444")

	// 接受连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go func() {
			defer conn.Close()
			err := server.ServeConn(conn)
			if err != nil {
				log.Printf("SOCKS5 connection error: %v", err)
			}
		}()
	}
}
`

	// 写入SOCKS5服务器代码
	serverFile := "socks5_test_server.go"
	if err := os.WriteFile(serverFile, []byte(socks5ServerCode), 0644); err != nil {
		t.Fatalf("创建SOCKS5服务器代码失败: %v", err)
	}
	defer os.Remove(serverFile)
	defer os.Remove("socks5-test-server.exe")

	// 启动SOCKS5代理服务器
	testResults = append(testResults, "## 1. 启动SOCKS5代理服务器 (go-socks5)")
	testResults = append(testResults, "")
	testResults = append(testResults, "编译并启动SOCKS5服务器...")
	testResults = append(testResults, "")

	// 编译SOCKS5服务器
	buildCmd := processManager.Command("go", "build", "-o", "socks5-test-server.exe", serverFile)
	buildCmd.Stdout = multiWriter
	buildCmd.Stderr = multiWriter

	// 记录命令执行
	processManager.LogCommand(buildCmd, "BUILD")
	if err := buildCmd.Run(); err != nil {
		processManager.LogCommandResult(buildCmd, err, "")
		t.Fatalf("编译SOCKS5服务器失败: %v", err)
	}
	processManager.LogCommandResult(buildCmd, nil, "")
	testResults = append(testResults, "✅ SOCKS5服务器编译成功")
	testResults = append(testResults, "")

	// 启动SOCKS5服务器进程
	cmd := processManager.Command("./socks5-test-server.exe")
	cmd.Stdout = multiWriter
	cmd.Stderr = multiWriter

	// 设置进程属性，确保能终止所有子进程（跨平台兼容）
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = NewSysProcAttr()
	}

	err := cmd.Start()
	if err != nil {
		t.Fatalf("启动SOCKS5服务器失败: %v", err)
	}

	// 将SOCKS5服务器进程添加到管理器
	processManager.AddProcess(cmd)
	log.Printf("SOCKS5服务器已启动，PID: %d\n", cmd.Process.Pid)

	// 确保进程能正确退出
	go func() {
		cmd.Wait()
		log.Println("SOCKS5服务器进程已退出")
	}()

	// 记录服务器PID
	testResults = append(testResults, fmt.Sprintf("📋 SOCKS5服务器进程PID: %d", cmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待服务器启动
	testResults = append(testResults, "等待SOCKS5服务器启动...")

	// 等待服务器启动，增加重试机制
	serverStarted := false
	for i := 0; i < 15; i++ {
		if isSOCKS5ProxyServerRunningWithGolangNetProxy() {
			serverStarted = true
			break
		}
		time.Sleep(1 * time.Second)
		log.Printf("等待SOCKS5服务器启动... %d/15\n", i+1)
	}

	if !serverStarted {
		t.Fatal("SOCKS5服务器启动失败")
	}

	testResults = append(testResults, "✅ SOCKS5服务器启动成功")
	testResults = append(testResults, "")

	// 添加启动成功的日志输出提示
	log.Println("SOCKS5服务器启动成功，开始执行测试...")

	// 等待额外的时间确保服务器完全启动
	time.Sleep(3 * time.Second)

	// 测试SOCKS5代理功能
	testResults = append(testResults, "## 2. 测试SOCKS5代理功能 (使用 golang.org/x/net/proxy)")
	testResults = append(testResults, "")

	// ===== 使用 golang.org/x/net/proxy 进行SOCKS5代理测试 =====

	// 创建SOCKS5代理拨号器，使用 golang.org/x/net/proxy
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:44444", &proxy.Auth{
		User:     "g7envpwz14b0u55",
		Password: "juvytdsdzc225pq",
	}, proxy.Direct)
	if err != nil {
		t.Fatalf("创建SOCKS5代理拨号器失败: %v", err)
	}

	// 测试1: 基本HTTP请求通过SOCKS5代理 (等效curl命令)
	testResults = append(testResults, "### 测试1: HTTP请求通过SOCKS5代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "等效命令: `curl -v -X GET http://httpbin.org/ip -x socks5://g7envpwz14b0u55:juvytdsdzc225pq@127.0.0.1:44444`")
	testResults = append(testResults, "")

	// 创建HTTP客户端，使用自定义的SOCKS5拨号器
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 30 * time.Second,
	}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	if err != nil {
		t.Fatalf("创建HTTP请求失败: %v", err)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err))
		testResults = append(testResults, "")
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, fmt.Sprintf("状态码: %d", resp.StatusCode))
		testResults = append(testResults, "")
		testResults = append(testResults, "响应内容:")
		testResults = append(testResults, "```")
		testResults = append(testResults, string(body))
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 测试2: HTTPS请求通过SOCKS5代理 (等同于用户提供的curl命令)
	testResults = append(testResults, "### 测试2: HTTPS请求通过SOCKS5代理")
	testResults = append(testResults, "")
	testResults = append(testResults, "等效命令: `curl -v -I -X GET https://dns.google -x socks5://g7envpwz14b0u55:juvytdsdzc225pq@127.0.0.1:44444`")
	testResults = append(testResults, "")

	// 创建HTTPS测试请求
	req2, err := http.NewRequest("HEAD", "https://dns.google", nil)
	if err != nil {
		t.Fatalf("创建HTTPS请求失败: %v", err)
	}

	// 使用自定义拨号器发送HTTPS请求
	client2 := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 30 * time.Second,
	}

	// 发送HTTPS请求
	resp2, err := client2.Do(req2)
	if err != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err))
		testResults = append(testResults, "")
	} else {
		defer resp2.Body.Close()

		testResults = append(testResults, "✅ 测试成功")
		testResults = append(testResults, "")
		testResults = append(testResults, fmt.Sprintf("状态码: %d", resp2.StatusCode))
		testResults = append(testResults, "")

		// 记录响应头
		testResults = append(testResults, "响应头:")
		testResults = append(testResults, "```")
		for key, values := range resp2.Header {
			for _, value := range values {
				testResults = append(testResults, fmt.Sprintf("%s: %s", key, value))
			}
		}
		testResults = append(testResults, "```")
	}
	testResults = append(testResults, "")

	// 测试3: 使用 golang.org/x/net/proxy 进行多种协议测试
	testResults = append(testResults, "### 测试3: 多目标URL测试 (使用 golang.org/x/net/proxy)")
	testResults = append(testResults, "")
	testResults = append(testResults, "测试多个网站通过SOCKS5代理的连接...")
	testResults = append(testResults, "")

	// 测试多个目标URL
	testURLs := []struct {
		url      string
		method   string
		expected int
		desc     string
	}{
		{"https://httpbin.org/get", "GET", 200, "获取IP信息"},
		{"https://ifconfig.me/ip", "GET", 200, "获取外部IP"},
		{"https://api.ipify.org?format=text", "GET", 200, "IP查询服务"},
		{"https://httpbin.org/status/200", "GET", 200, "状态码测试"},
		{"https://www.google.com", "HEAD", 200, "Google首页"},
	}

	for i, testCase := range testURLs {
		testResults = append(testResults, fmt.Sprintf("#### 子测试 3.%d: %s - %s", i+1, testCase.method, testCase.desc))
		testResults = append(testResults, "")
		testResults = append(testResults, fmt.Sprintf("URL: %s", testCase.url))
		testResults = append(testResults, "")

		// 创建新的HTTP客户端，使用SOCKS5拨号器
		client3 := &http.Client{
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
			Timeout: 20 * time.Second,
		}

		// 创建请求
		req3, err := http.NewRequest(testCase.method, testCase.url, nil)
		if err != nil {
			testResults = append(testResults, fmt.Sprintf("❌ 创建请求失败: %v", err))
			testResults = append(testResults, "")
			continue
		}

		// 发送请求
		startTime := time.Now()
		resp3, err := client3.Do(req3)
		responseTime := time.Since(startTime)

		if err != nil {
			testResults = append(testResults, fmt.Sprintf("❌ 请求失败: %v", err))
			testResults = append(testResults, "")
		} else {
			defer resp3.Body.Close()

			if resp3.StatusCode == testCase.expected {
				testResults = append(testResults, "✅ 请求成功")
				testResults = append(testResults, "")
				testResults = append(testResults, fmt.Sprintf("状态码: %d", resp3.StatusCode))
				testResults = append(testResults, fmt.Sprintf("响应时间: %v", responseTime))
				testResults = append(testResults, "")

				// 如果是GET请求且内容不长，显示响应内容
				if testCase.method == "GET" && resp3.ContentLength < 1000 {
					body3, _ := io.ReadAll(resp3.Body)
					testResults = append(testResults, "响应内容:")
					testResults = append(testResults, "```")
					testResults = append(testResults, string(body3))
					testResults = append(testResults, "```")
				}
			} else {
				testResults = append(testResults, fmt.Sprintf("❌ 请求失败，状态码: %d (期望: %d)", resp3.StatusCode, testCase.expected))
				testResults = append(testResults, "")
			}
		}
		testResults = append(testResults, "")
	}

	// 记录所有进程PID信息
	testResults = append(testResults, "### 📋 所有进程PID记录")
	testResults = append(testResults, "")
	allPIDs := processManager.GetPIDs()
	testResults = append(testResults, fmt.Sprintf("所有进程PID: %s", strings.Join(allPIDs, ", ")))
	testResults = append(testResults, "")

	// 写入测试记录到文件
	err = WriteTestResultsToFile(testResults, processManager.GetFile())
	if err != nil {
		t.Errorf("写入测试记录失败: %v", err)
	}

	// 停止超时计时器
	timeoutTimer.Stop()

	// 如果所有测试成功，关闭代理服务器进程
	testResults = append(testResults, "## 3. 关闭SOCKS5代理服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "✅ SOCKS5代理测试完成，正在关闭代理服务器进程...")
	testResults = append(testResults, "")

	// 明确终止代理服务器进程
	testResults = append(testResults, "🛑 正在终止SOCKS5代理服务器进程...")
	if cmd.Process != nil {
		log.Printf("正在终止SOCKS5代理服务器进程 PID: %d\n", cmd.Process.Pid)
		if err := cmd.Process.Kill(); err != nil {
			testResults = append(testResults, fmt.Sprintf("❌ 终止SOCKS5代理服务器进程失败: %v", err))
			log.Printf("终止SOCKS5代理服务器进程失败: %v\n", err)
		} else {
			cmd.Wait() // 等待进程完全退出
			testResults = append(testResults, "✅ SOCKS5代理服务器进程已终止")
			log.Println("SOCKS5代理服务器进程已终止")
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
	log.Println("正在记录SOCKS5代理服务器日志...")

	// 使用互斥锁保护对proxyOutput的访问
	proxyOutputMutex.Lock()
	outputLen := proxyOutput.Len()
	outputContent := proxyOutput.String()
	proxyOutputMutex.Unlock()

	if outputLen > 0 {
		testResults = append(testResults, "### SOCKS5代理服务器日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "```")
		// 按行分割输出并添加到测试结果
		outputLines := strings.Split(outputContent, "\n")
		for _, line := range outputLines {
			if strings.TrimSpace(line) != "" {
				testResults = append(testResults, line)
				log.Println("[代理日志]", line) // 同时打印到控制台
			}
		}
		testResults = append(testResults, "```")
		testResults = append(testResults, "")
	} else {
		testResults = append(testResults, "### SOCKS5代理服务器日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "⚠️ 没有捕获到SOCKS5代理服务器日志")
		testResults = append(testResults, "")
		log.Println("⚠️ 没有捕获到SOCKS5代理服务器日志")

		// 添加调试信息
		testResults = append(testResults, "### 调试信息")
		testResults = append(testResults, "")
		testResults = append(testResults, fmt.Sprintf("SOCKS5代理服务器输出缓冲区长度: %d", outputLen))
		testResults = append(testResults, "")
		testResults = append(testResults, "可能的原因:")
		testResults = append(testResults, "- SOCKS5代理服务器程序没有输出日志")
		testResults = append(testResults, "- 日志输出被重定向到其他地方")
		testResults = append(testResults, "- 缓冲区没有正确捕获输出")
		testResults = append(testResults, "")
	}

	// 验证端口是否已释放
	if !isPortOccupied(44444) {
		testResults = append(testResults, "✅ 端口44444已成功释放")
	} else {
		testResults = append(testResults, "❌ 端口44444仍被占用")
	}

	// 重新写入测试记录
	err = WriteTestResultsToFile(testResults, processManager.GetFile())
	if err != nil {
		t.Errorf("更新测试记录失败: %v", err)
	}
}

// isSOCKS5ProxyServerRunningWithGolangNetProxy 使用 golang.org/x/net/proxy 检查SOCKS5代理服务器是否正在运行
func isSOCKS5ProxyServerRunningWithGolangNetProxy() bool {
	// 简单检查端口是否开放
	conn, err := net.DialTimeout("tcp", "127.0.0.1:44444", 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()

	// 创建SOCKS5代理拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:44444", &proxy.Auth{
		User:     "g7envpwz14b0u55",
		Password: "juvytdsdzc225pq",
	}, proxy.Direct)
	if err != nil {
		return false
	}

	// 尝试通过SOCKS5代理建立TCP连接
	conn, err = dialer.Dial("tcp", "httpbin.org:80")
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

// TestSOCKS5Proxy 主测试函数
func TestSOCKS5Proxy(t *testing.T) {
	var processManager *ProcessManager = NewProcessManager("socks5_proxy_test.log")
	defer func() {
		// 清理所有进程
		processManager.CleanupAll()
		processManager.Close()
	}()

	// 创建带有45秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan int, 1)

	// 在goroutine中运行测试
	go func() {
		// 运行SOCKS5代理测试
		runSOCKS5ProxyServer(t, "socks5_proxy_test.log")
		resultChan <- 0
	}()

	// 等待测试完成或超时
	select {
	case <-resultChan:
		// 测试正常完成
		return
	case <-ctx.Done():
		// 超时或取消
		log.Println("\n⏰ SOCKS5代理测试超时（45秒），强制退出...")

		// 强制终止所有记录的进程
		log.Println("正在终止所有运行中的进程...")

		// 在Windows上强制终止所有go进程
		if runtime.GOOS == "windows" {
			// 使用taskkill终止所有go进程
			killCmd := processManager.Command("taskkill", "/F", "/IM", "go.exe")
			processManager.LogCommand(killCmd, "CLEANUP")
			killCmd.Run() // 忽略错误
			processManager.LogCommandResult(killCmd, nil, "")

			// 终止可能的代理服务器进程（在44444端口上）
			findCmd := processManager.Command("netstat", "-ano", "|", "findstr", ":44444")
			processManager.LogCommand(findCmd, "CLEANUP")
			findCmd.Run() // 忽略错误
			processManager.LogCommandResult(findCmd, nil, "")
		}

		// 清理全局进程管理器中的进程
		if processManager != nil {
			processManager.CleanupAll()
		}

		// 记录超时信息到测试记录
		timeoutMessage := []string{
			"# SOCKS5代理测试超时记录",
			"",
			"## 超时时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"❌ SOCKS5代理测试执行超过45秒超时限制，强制退出",
			"",
			"可能的原因:",
			"- SOCKS5代理服务器进程未正常退出",
			"- golang.org/x/net/proxy 连接卡住",
			"- 网络连接问题",
			"- SOCKS5代理配置问题",
			"",
			"已尝试终止所有相关进程",
			"",
		}

		// 写入超时记录
		if err := WriteTestResultsToFile(timeoutMessage, processManager.GetFile()); err != nil {
			log.Printf("写入超时记录失败: %v\n", err)
		}

		// 强制退出
		t.Fatal("SOCKS5代理测试超时")
	}
}