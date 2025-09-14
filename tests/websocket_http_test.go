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
	// "os/exec"
	"runtime"
	"strings"
	//	"syscall"
	"testing"
	"time"
)

// runWebSockethttpProxy 测试WebSocket和http级联代理服务器
func runWebSockethttpProxy(t *testing.T) {
	// 使用传入的进程管理器
	var processManager *ProcessManager = NewProcessManager()
	defer func() {

		// 清理所有进程
		processManager.CleanupAll()
		processManager.Close()
	}()
	// 创建缓冲区来捕获服务器输出
	var websocketOutput bytes.Buffer
	var httpOutput bytes.Buffer

	// 创建多写入器
	websocketWriter := io.MultiWriter(os.Stdout, &websocketOutput)
	httpWriter := io.MultiWriter(os.Stdout, &httpOutput)

	// 清理可能存在的旧的可执行文件
	if _, err := os.Stat("main.exe"); err == nil {
		os.Remove("main.exe")
	}

	// 添加测试超时检查
	timeoutTimer := time.AfterFunc(30*time.Second, func() {
		log.Println("\n⚠️ 测试即将超时，正在清理进程...")
		processManager.CleanupAll()
		t.Fatal("测试超时")
	})
	defer timeoutTimer.Stop()

	// 测试结果记录
	var testResults []string
	testResults = append(testResults, "# WebSocket和http级联代理测试记录")
	testResults = append(testResults, "")
	testResults = append(testResults, "## 测试时间")
	testResults = append(testResults, time.Now().Format("2006-01-02 15:04:05"))
	testResults = append(testResults, "")

	// 检查端口是否被占用
	if IsPortOccupied2(18080) {
		t.Fatal("端口18080已被占用，请先停止占用该端口的进程")
	}
	if IsPortOccupied2(18080) {
		t.Fatal("端口18080已被占用，请先停止占用该端口的进程")
	}

	// 编译代理服务器
	testResults = append(testResults, "## 1. 编译代理服务器")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd`")
	testResults = append(testResults, "")

	buildCmd1 := processManager.Command("go", "build", "-o", "socks5-websocket-proxy-golang.exe", "github.com/masx200/socks5-websocket-proxy-golang/cmd")
	buildCmd1.Stdout = websocketWriter
	buildCmd1.Stderr = websocketWriter

	// 记录命令执行
	processManager.LogCommand(buildCmd1, "BUILD")
	if err := buildCmd1.Run(); err != nil {
		processManager.LogCommandResult(buildCmd1, err, "")
		t.Fatalf("编译代理服务器失败: %v", err)
	}
	processManager.LogCommandResult(buildCmd1, nil, "")

	buildCmd := processManager.Command("go", "build", "-o", "main.exe", "../cmd/main.go")
	buildCmd.Stdout = websocketWriter
	buildCmd.Stderr = websocketWriter

	// 记录命令执行
	processManager.LogCommand(buildCmd, "BUILD")
	if err := buildCmd.Run(); err != nil {
		processManager.LogCommandResult(buildCmd, err, "")
		t.Fatalf("编译代理服务器失败: %v", err)
	}
	processManager.LogCommandResult(buildCmd, nil, "")
	testResults = append(testResults, "✅ 代理服务器编译成功")
	testResults = append(testResults, "")

	// 启动WebSocket服务器（作为上游）
	testResults = append(testResults, "## 2. 启动WebSocket服务器（上游）")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :28081`")
	testResults = append(testResults, "")

	websocketCmd := processManager.Command("./socks5-websocket-proxy-golang.exe", "-mode", "server", "-protocol", "websocket", "-addr", ":28081")
	websocketCmd.Stdout = websocketWriter
	websocketCmd.Stderr = websocketWriter

	// 记录命令执行
	processManager.LogCommand(websocketCmd, "WEBSOCKET")

	// 设置进程属性
	if runtime.GOOS == "windows" {
		//websocketCmd.SysProcAttr =

		websocketCmd.SysProcAttr = NewSysProcAttr()

		// &syscall.SysProcAttr{
		//			CreationFlags:
		//syscall.CREATE_NEW_PROCESS_GROUP,
		//	}
	}

	err := websocketCmd.Start()
	if err != nil {
		processManager.LogCommandResult(websocketCmd, err, "")
		t.Fatalf("启动WebSocket服务器失败: %v", err)
	}
	processManager.LogCommandResult(websocketCmd, nil, "")

	processManager.AddProcess(websocketCmd)
	log.Printf("WebSocket服务器已启动，PID: %d\n", websocketCmd.Process.Pid)
	testResults = append(testResults, fmt.Sprintf("📋 WebSocket服务器进程PID: %d", websocketCmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待WebSocket服务器启动
	testResults = append(testResults, "等待WebSocket服务器启动...")
	websocketStarted := false
	for i := 0; i < 10; i++ {
		if IsPortOccupied2(28081) {
			websocketStarted = true
			break
		}
		time.Sleep(1 * time.Second)
		log.Printf("等待WebSocket服务器启动... %d/10\n", i+1)
	}

	if !websocketStarted {
		t.Fatal("WebSocket服务器启动失败")
	}

	testResults = append(testResults, "✅ WebSocket服务器启动成功")
	testResults = append(testResults, "")

	// 启动http服务器（设置upstream为WebSocket服务器）
	testResults = append(testResults, "## 3. 启动http服务器（下游）")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `./main.exe  -port 18080 -upstream-type websocket -upstream-address ws://localhost:28081`")
	testResults = append(testResults, "")

	httpCmd := processManager.Command("./main.exe", "-port", "18080",
		"-upstream-type", "websocket", "-upstream-address", "ws://localhost:28081")
	httpCmd.Stdout = httpWriter
	httpCmd.Stderr = httpWriter

	// 记录命令执行
	processManager.LogCommand(httpCmd, "HTTP")

	// 设置进程属性
	if runtime.GOOS == "windows" {
		httpCmd.SysProcAttr = NewSysProcAttr()

		//&syscall.SysProcAttr{
		//	CreationFlags:
		//syscall.CREATE_NEW_PROCESS_GROUP,
		//		}
	}

	err = httpCmd.Start()
	if err != nil {
		processManager.LogCommandResult(httpCmd, err, "")
		t.Fatalf("启动http服务器失败: %v", err)
	}
	processManager.LogCommandResult(httpCmd, nil, "")

	processManager.AddProcess(httpCmd)
	log.Printf("http服务器已启动，PID: %d\n", httpCmd.Process.Pid)
	testResults = append(testResults, fmt.Sprintf("📋 http服务器进程PID: %d", httpCmd.Process.Pid))
	testResults = append(testResults, "")

	// 等待http服务器启动
	testResults = append(testResults, "等待http服务器启动...")
	httpStarted := false
	for i := 0; i < 10; i++ {
		if ishttpProxyRunning() {
			httpStarted = true
			break
		}
		time.Sleep(1 * time.Second)
		log.Printf("等待http服务器启动... %d/10\n", i+1)
	}

	if !httpStarted {
		t.Fatal("http服务器启动失败")
	}

	testResults = append(testResults, "✅ http服务器启动成功")
	testResults = append(testResults, "")

	// 等待额外的时间确保服务器完全启动
	time.Sleep(2 * time.Second)

	// 测试级联代理功能
	testResults = append(testResults, "## 4. 测试级联代理功能")
	testResults = append(testResults, "")

	// 测试HTTP代理
	testResults = append(testResults, "### 测试1: HTTP代理通过级联")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I http://www.baidu.com -x http://localhost:18080`")
	testResults = append(testResults, "")

	curlCmd1 := processManager.Command("curl", "-v", "-I", "http://www.baidu.com", "-x", "http://localhost:18080")
	var curlOutput1 bytes.Buffer
	curlCmd1.Stdout = &curlOutput1
	curlCmd1.Stderr = &curlOutput1

	// 记录命令执行
	processManager.LogCommand(curlCmd1, "CURL")
	err1 := curlCmd1.Run()
	output1 := curlOutput1.Bytes()

	// 检查进程退出状态码
	exitCode1 := 0
	if curlCmd1.ProcessState != nil {
		exitCode1 = curlCmd1.ProcessState.ExitCode()
	}

	// 记录命令执行结果
	processManager.LogCommandResult(curlCmd1, err1, string(output1))

	processManager.AddProcess(curlCmd1)
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试1进程PID: %d, 退出状态码: %d", curlCmd1.Process.Pid, exitCode1))
	testResults = append(testResults, "")

	if err1 != nil || exitCode1 != 0 {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err1))
		testResults = append(testResults, fmt.Sprintf("退出状态码: %d", exitCode1))
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

	// 测试HTTPS代理
	testResults = append(testResults, "### 测试2: HTTPS代理通过级联")
	testResults = append(testResults, "")
	testResults = append(testResults, "执行命令: `curl -v -I https://www.baidu.com -x http://localhost:18080`")
	testResults = append(testResults, "")

	curlCmd2 := processManager.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "http://localhost:18080")
	var curlOutput2 bytes.Buffer
	curlCmd2.Stdout = &curlOutput2
	curlCmd2.Stderr = &curlOutput2

	// 记录命令执行
	processManager.LogCommand(curlCmd2, "CURL")
	err2 := curlCmd2.Run()
	output2 := curlOutput2.Bytes()

	// 检查进程退出状态码
	exitCode2 := 0
	if curlCmd2.ProcessState != nil {
		exitCode2 = curlCmd2.ProcessState.ExitCode()
	}

	// 记录命令执行结果
	processManager.LogCommandResult(curlCmd2, err2, string(output2))

	processManager.AddProcess(curlCmd2)
	testResults = append(testResults, fmt.Sprintf("📋 Curl测试2进程PID: %d, 退出状态码: %d", curlCmd2.Process.Pid, exitCode2))
	testResults = append(testResults, "")

	if err2 != nil || exitCode2 != 0 {
		testResults = append(testResults, fmt.Sprintf("❌ 测试失败: %v", err2))
		testResults = append(testResults, fmt.Sprintf("退出状态码: %d", exitCode2))
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

	// 记录所有进程PID信息
	testResults = append(testResults, "### 📋 所有进程PID记录")
	testResults = append(testResults, "")
	allPIDs := processManager.GetPIDs()
	testResults = append(testResults, fmt.Sprintf("所有进程PID: %s", strings.Join(allPIDs, ", ")))
	testResults = append(testResults, "")

	// 写入测试记录到文件
	err = WriteTestResults2(testResults)
	if err != nil {
		t.Errorf("写入测试记录失败: %v", err)
	}

	// 验证测试结果
	if err1 != nil {
		t.Errorf("HTTP curl测试失败: %v", err1)
	}
	if err2 != nil {
		t.Errorf("HTTPS curl测试失败: %v", err2)
	}

	// 如果测试成功，关闭服务器进程
	if err1 == nil && err2 == nil {
		testResults = append(testResults, "## 5. 关闭服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "✅ 所有测试成功，正在关闭服务器进程...")
		testResults = append(testResults, "")

		// 停止超时计时器
		timeoutTimer.Stop()

		// 终止WebSocket服务器
		testResults = append(testResults, "🛑 正在终止WebSocket服务器进程...")
		if websocketCmd.Process != nil {
			if err := websocketCmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止WebSocket服务器进程失败: %v", err))
			} else {
				websocketCmd.Wait()
				testResults = append(testResults, "✅ WebSocket服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 终止http服务器
		testResults = append(testResults, "🛑 正在终止http服务器进程...")
		if httpCmd.Process != nil {
			if err := httpCmd.Process.Kill(); err != nil {
				testResults = append(testResults, fmt.Sprintf("❌ 终止http服务器进程失败: %v", err))
			} else {
				httpCmd.Wait()
				testResults = append(testResults, "✅ http服务器进程已终止")
			}
		}
		testResults = append(testResults, "")

		// 清理所有进程
		testResults = append(testResults, "🧹 正在清理所有子进程...")
		processManager.CleanupAll()
		testResults = append(testResults, "✅ 所有子进程已清理完成")
		testResults = append(testResults, "")

		// 等待进程完全关闭
		time.Sleep(2 * time.Second)

		// 清理编译的可执行文件
		if _, err := os.Stat("main.exe"); err == nil {
			os.Remove("main.exe")
			testResults = append(testResults, "🧹 已清理编译的可执行文件")
		}

		// 添加服务器日志输出
		testResults = append(testResults, "### WebSocket服务器日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "```")
		websocketLines := strings.Split(websocketOutput.String(), "\n")
		for _, line := range websocketLines {
			if strings.TrimSpace(line) != "" {
				testResults = append(testResults, line)
			}
		}
		testResults = append(testResults, "```")
		testResults = append(testResults, "")

		testResults = append(testResults, "### http服务器日志输出")
		testResults = append(testResults, "")
		testResults = append(testResults, "```")
		httpLines := strings.Split(httpOutput.String(), "\n")
		for _, line := range httpLines {
			if strings.TrimSpace(line) != "" {
				testResults = append(testResults, line)
			}
		}
		testResults = append(testResults, "```")
		testResults = append(testResults, "")

		// 验证端口是否已释放
		if !IsPortOccupied2(18080) {
			testResults = append(testResults, "✅ 端口18080已成功释放")
		} else {
			testResults = append(testResults, "❌ 端口18080仍被占用")
		}
		if !IsPortOccupied2(18080) {
			testResults = append(testResults, "✅ 端口18080已成功释放")
		} else {
			testResults = append(testResults, "❌ 端口18080仍被占用")
		}

		// 重新写入测试记录
		err = WriteTestResults2(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}
	} else {
		// 如果有测试失败，也记录关闭进程的信息
		testResults = append(testResults, "## 5. 关闭服务器")
		testResults = append(testResults, "")
		testResults = append(testResults, "⚠️ 部分测试失败，但仍需关闭服务器进程...")
		testResults = append(testResults, "")

		// 终止WebSocket服务器
		if websocketCmd.Process != nil {
			websocketCmd.Process.Kill()
			websocketCmd.Wait()
		}

		// 终止http服务器
		if httpCmd.Process != nil {
			httpCmd.Process.Kill()
			httpCmd.Wait()
		}

		// 清理所有进程
		processManager.CleanupAll()

		// 等待进程完全关闭
		time.Sleep(2 * time.Second)

		// 清理编译的可执行文件
		if _, err := os.Stat("main.exe"); err == nil {
			os.Remove("main.exe")
		}

		// 重新写入测试记录
		err = WriteTestResults2(testResults)
		if err != nil {
			t.Errorf("更新测试记录失败: %v", err)
		}
	}
}

// ishttpProxyRunning 检查http代理服务器是否正在运行
func ishttpProxyRunning() bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 创建一个测试请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return false
	}

	// 设置代理
	proxyURL, err := url.Parse("http://localhost:18080")
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

// IsPortOccupied2 检查端口是否被占用
func IsPortOccupied2(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return true
	}
	listener.Close()
	return false
}

// WriteTestResults2 写入测试结果到文件
func WriteTestResults2(results []string) error {
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

// RunMainWebSocket 主测试函数
func RunMainWebSocket(t *testing.T) {

	var processManager *ProcessManager = NewProcessManager()
	defer func() {

		// 清理所有进程
		processManager.CleanupAll()
		processManager.Close()
	}()
	// 创建带有35秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan bool, 1)

	// 在goroutine中运行测试
	go func() {
		// 运行测试，并传递进程管理器
		runWebSockethttpProxy(t)
		resultChan <- true
	}()

	// 等待测试完成或超时
	select {
	case <-resultChan:

		// 测试正常完成
		return
	case <-ctx.Done():
		// 超时或取消
		log.Println("\n⏰ 测试超时（35秒），强制退出...")

		// 在Windows上强制终止所有go进程
		if runtime.GOOS == "windows" {
			killCmd := processManager.Command("taskkill", "/F", "/IM", "go.exe")
			processManager.LogCommand(killCmd, "CLEANUP")
			killCmd.Run()
			processManager.LogCommandResult(killCmd, nil, "")
		}

		// 记录超时信息
		timeoutMessage := []string{
			"# WebSocket和http级联测试超时记录",
			"",
			"## 超时时间",
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"❌ 测试执行超过35秒超时限制，强制退出",
			"",
			"可能的原因:",
			"- 服务器进程未正常退出",
			"- curl命令卡住",
			"- 网络连接问题",
			"- 级联代理配置问题",
			"",
			"已尝试终止所有相关进程",
			"",
		}

		// 写入超时记录
		if err := WriteTestResults2(timeoutMessage); err != nil {
			log.Printf("写入超时记录失败: %v\n", err)
		}

		// 强制退出
		t.Fatal("测试超时")
	}
}
