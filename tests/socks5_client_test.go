package tests

import (
	"io"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/proxy"
)

// TestSOCKS5ClientWithStandaloneServer 测试golang.org/x/net/proxy客户端连接到独立SOCKS5服务器
func TestSOCKS5ClientWithStandaloneServer(t *testing.T) {
	// 等待SOCKS5服务器完全启动
	time.Sleep(2 * time.Second)

	// 创建SOCKS5代理拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:44444", &proxy.Auth{
		User:     "g7envpwz14b0u55",
		Password: "juvytdsdzc225pq",
	}, proxy.Direct)
	if err != nil {
		t.Fatalf("创建SOCKS5代理拨号器失败: %v", err)
	}

	t.Log("✅ 成功创建SOCKS5代理拨号器")

	// 创建HTTP客户端，使用自定义的SOCKS5拨号器
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 30 * time.Second,
	}

	// 测试1: HTTP请求通过SOCKS5代理
	t.Log("测试1: HTTP请求通过SOCKS5代理")
	req1, err := http.NewRequest("GET", "https://api.ip.sb/ip", nil)
	if err != nil {
		t.Fatalf("创建HTTP请求失败: %v", err)
	}

	resp1, err := client.Do(req1)
	if err != nil {
		t.Errorf("HTTP请求失败: %v", err)
	} else {
		defer resp1.Body.Close()
		body1, _ := io.ReadAll(resp1.Body)
		t.Logf("✅ HTTP请求成功，状态码: %d", resp1.StatusCode)
		t.Logf("响应内容: %s", string(body1))
	}

	// 测试2: HTTPS请求通过SOCKS5代理 (等同于用户提供的curl命令)
	t.Log("测试2: HTTPS请求通过SOCKS5代理")
	req2, err := http.NewRequest("HEAD", "https://www.baidu.com", nil)
	if err != nil {
		t.Fatalf("创建HTTPS请求失败: %v", err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Errorf("HTTPS请求失败: %v", err)
	} else {
		defer resp2.Body.Close()
		t.Logf("✅ HTTPS请求成功，状态码: %d", resp2.StatusCode)
		t.Logf("响应头数量: %d", len(resp2.Header))
	}
}

// TestSOCKS5ClientWithStandaloneServer 测试golang.org/x/net/proxy客户端连接到独立SOCKS5服务器
func TestSOCKS5ClientWithStandaloneServer8086(t *testing.T) {
	// 等待SOCKS5服务器完全启动
	time.Sleep(2 * time.Second)

	// 创建SOCKS5代理拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:8086", &proxy.Auth{
		User:     "g7envpwz14b0u55",
		Password: "juvytdsdzc225pq",
	}, proxy.Direct)
	if err != nil {
		t.Fatalf("创建SOCKS5代理拨号器失败: %v", err)
	}

	t.Log("✅ 成功创建SOCKS5代理拨号器")

	// 创建HTTP客户端，使用自定义的SOCKS5拨号器
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 30 * time.Second,
	}

	// 测试1: HTTP请求通过SOCKS5代理
	t.Log("测试1: HTTP请求通过SOCKS5代理")
	req1, err := http.NewRequest("GET", "https://api.ip.sb/ip", nil)
	if err != nil {
		t.Fatalf("创建HTTP请求失败: %v", err)
	}

	resp1, err := client.Do(req1)
	if err != nil {
		t.Errorf("HTTP请求失败: %v", err)
	} else {
		defer resp1.Body.Close()
		body1, _ := io.ReadAll(resp1.Body)
		t.Logf("✅ HTTP请求成功，状态码: %d", resp1.StatusCode)
		t.Logf("响应内容: %s", string(body1))
	}

	// 测试2: HTTPS请求通过SOCKS5代理 (等同于用户提供的curl命令)
	t.Log("测试2: HTTPS请求通过SOCKS5代理")
	req2, err := http.NewRequest("HEAD", "https://www.baidu.com", nil)
	if err != nil {
		t.Fatalf("创建HTTPS请求失败: %v", err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Errorf("HTTPS请求失败: %v", err)
	} else {
		defer resp2.Body.Close()
		t.Logf("✅ HTTPS请求成功，状态码: %d", resp2.StatusCode)
		t.Logf("响应头数量: %d", len(resp2.Header))
	}
}
