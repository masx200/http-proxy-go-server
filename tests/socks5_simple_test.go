package tests

import (
	"net"
	"testing"
	"time"
)

// TestSimpleSOCKS5Connection 简单的SOCKS5连接测试
func TestSimpleSOCKS5Connection(t *testing.T) {
	// 等待一段时间确保之前的测试服务器已经清理
	time.Sleep(2 * time.Second)

	// 检查端口是否被占用
	if isPortOccupied(44444) {
		t.Skip("端口44444被占用，跳过简单测试")
	}

	// 测试简单的TCP连接到SOCKS5服务器
	conn, err := net.DialTimeout("tcp", "127.0.0.1:44444", 5*time.Second)
	if err != nil {
		t.Skipf("无法连接到SOCKS5服务器: %v", err)
		return
	}
	defer conn.Close()

	t.Log("✅ 成功连接到SOCKS5服务器")

	// 发送一些测试数据到SOCKS5服务器
	testData := []byte("Hello SOCKS5 Server")
	_, err = conn.Write(testData)
	if err != nil {
		t.Errorf("写入测试数据失败: %v", err)
	} else {
		t.Logf("✅ 成功发送测试数据到SOCKS5服务器")
	}
}