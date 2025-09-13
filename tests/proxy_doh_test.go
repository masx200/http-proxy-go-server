package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"
	"time"
)

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
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("编译代理服务器失败: %v", err)
	}
	testResults = append(testResults, "✅ 代理服务器编译成功")
	testResults = append(testResults, "")

	// 启动代理服务器进程（使用编译后的可执行文件，带DoH参数）
	testResults = append(testResults, "启动代理服务器（DoH模式）...")
	testResults = append(testResults, "执行命令: `./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohalpn h2 -dohalpn h3`")
	testResults = append(testResults, "")

	cmd := exec.Command("./main.exe", "-dohurl", "https://dns.alidns.com/dns-query", "-dohip", "223.5.5.5", "-dohip", "223.6.6.6", "-dohurl", "https://dns.alidns.com/dns-query", "-dohalpn", "h2", "-dohalpn", "h3")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 将进程添加到进程管理器
	processManager.AddProcess(cmd)

	// 启动代理服务器
	if err := cmd.Start(); err != nil {
		t.Fatalf("启动代理服务器失败: %v", err)
	}

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
	curlCmd1.Stdout = os.Stdout
	curlCmd1.Stderr = os.Stderr

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd1)

	// 启动curl进程
	err1 := curlCmd1.Run()
	if err1 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试1失败: %v", err1))
	} else {
		testResults = append(testResults, "✅ 测试1成功")
	}
	testResults = append(testResults, "")

	// 测试2：通过代理访问HTTP网站（跟随重定向）
	testResults = append(testResults, "测试2: 通过代理访问HTTP网站 (so.com，跟随重定向)")
	curlCmd2 := exec.Command("curl", "-v", "-I", "-L", "http://www.so.com", "-x", "http://localhost:8080")
	curlCmd2.Stdout = os.Stdout
	curlCmd2.Stderr = os.Stderr

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd2)

	// 启动curl进程
	err2 := curlCmd2.Run()
	if err2 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试2失败: %v", err2))
	} else {
		testResults = append(testResults, "✅ 测试2成功")
	}
	testResults = append(testResults, "")

	// 测试3：通过代理访问HTTPS网站
	testResults = append(testResults, "测试3: 通过代理访问HTTPS网站 (baidu.com)")
	curlCmd3 := exec.Command("curl", "-v", "-I", "https://www.baidu.com", "-x", "http://localhost:8080")
	curlCmd3.Stdout = os.Stdout
	curlCmd3.Stderr = os.Stderr

	// 将curl进程添加到进程管理器
	processManager.AddProcess(curlCmd3)

	// 启动curl进程
	err3 := curlCmd3.Run()
	if err3 != nil {
		testResults = append(testResults, fmt.Sprintf("❌ 测试3失败: %v", err3))
	} else {
		testResults = append(testResults, "✅ 测试3成功")
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
			killCmd.Run() // 忽略错误

			// 终止可能的代理服务器进程（在8080端口上）
			findCmd := exec.Command("netstat", "-ano", "|", "findstr", ":8080")
			findCmd.Run() // 忽略错误
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
