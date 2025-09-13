package tests

import (
	"context"
	"log"
	"os/exec"
	"runtime"
	"testing"
	"time"
)

// TestMainDOH 主测试函数 - DOH配置测试
func TestMainDOH(t *testing.T) {
	// 创建带有30秒超时的上下文（增加超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 创建通道来接收测试结果
	resultChan := make(chan int, 1)

	// 在goroutine中运行测试
	go func() {
		// 运行测试
		runProxyServer(t)
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

		// 在Windows上强制终止所有go进程和可能的子进程
		if runtime.GOOS == "windows" {
			// 使用taskkill终止所有go进程
			killCmd := exec.Command("taskkill", "/F", "/IM", "go.exe")
			killCmd.Run() // 忽略错误

			// 终止可能的代理服务器进程（在8080端口上）
			findCmd := exec.Command("netstat", "-ano", "|", "findstr", ":8080")
			findCmd.Run() // 忽略错误
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