package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"runtime"
	"testing"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
	"go.uber.org/goleak"
)

// TestDNSCacheMemoryLeaks 测试DNS缓存的内存泄漏
func TestDNSCacheMemoryLeaks(t *testing.T) {
	defer goleak.VerifyNone(t)

	// 配置测试缓存
	config := &dnscache.Config{
		FilePath:     "/tmp/test_cache.json",
		AOFPath:      "/tmp/test_cache.aof",
		DefaultTTL:   5 * time.Minute,
		SaveInterval: 1 * time.Second,
		AOFInterval:  500 * time.Millisecond,
		Enabled:      true,
	}

	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	defer cache.Close()

	// 记录初始内存
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 测试大量域名缓存
	numDomains := 10000
	for i := 0; i < numDomains; i++ {
		domain := fmt.Sprintf("memleak-test-%d.com", i)
		ips := []string{fmt.Sprintf("192.168.1.%d", i%256)}

		err := cache.Set(domain, ips)
		if err != nil {
			t.Errorf("Failed to set %s: %v", domain, err)
		}

		// 每千个域名检查一次
		if (i+1)%1000 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// 强制GC
	runtime.GC()
	time.Sleep(1 * time.Second)

	// 检查最终内存
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	allocatedMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
	t.Logf("DNS Cache allocated %.2f MB for %d domains", allocatedMB, numDomains)

	// 内存增长不应超过100MB（考虑到缓存开销）
	if allocatedMB > 100 {
		t.Errorf("Excessive memory usage: %.2f MB allocated", allocatedMB)
	}

	// 验证缓存可以被正常读取
	result := cache.Get("memleak-test-0.com")
	if result == nil {
		t.Error("Expected cached result")
	}
}

// TestProxyServerNoGoroutineLeak 测试代理服务器没有goroutine泄漏
func TestProxyServerNoGoroutineLeak(t *testing.T) {
	defer goleak.VerifyNone(t, goleak.IgnoreCurrent())

	// 记录初始goroutine数量
	initialGoroutines := runtime.NumGoroutine()
	t.Logf("Initial goroutines: %d", initialGoroutines)

	// 创建测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}))
	defer testServer.Close()

	// 模拟多个客户端连接
	numClients := 10
	connectionsPerClient := 100

	for client := 0; client < numClients; client++ {
		go func(clientID int) {
			proxyURL, _ := url.Parse(testServer.URL)

			client := &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(proxyURL),
				},
				Timeout: 10 * time.Second,
			}

			for i := 0; i < connectionsPerClient; i++ {
				resp, err := client.Get(testServer.URL)
				if err == nil {
					resp.Body.Close()
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(client)
	}

	// 等待所有操作完成
	time.Sleep(10 * time.Second)

	// 强制GC
	runtime.GC()
	time.Sleep(2 * time.Second)

	// 检查goroutine数量
	finalGoroutines := runtime.NumGoroutine()
	t.Logf("Final goroutines: %d", finalGoroutines)

	goroutineDiff := finalGoroutines - initialGoroutines
	if goroutineDiff > 5 {
		t.Errorf("Possible goroutine leak: %d goroutines created", goroutineDiff)
	}
}

// TestHTTPConnectionMemoryLeaks 测试HTTP连接的内存泄漏
func TestHTTPConnectionMemoryLeaks(t *testing.T) {
	defer goleak.VerifyNone(t)

	// 创建测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Test response")
	}))
	defer testServer.Close()

	// 记录初始内存
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 发送大量请求
	numRequests := 1000
	for i := 0; i < numRequests; i++ {
		resp, err := client.Get(testServer.URL)
		if err != nil {
			t.Logf("Request %d failed: %v", i, err)
			continue
		}

		// 确保响应体被完全读取和关闭
		_, _ = resp.Body.Read(make([]byte, 1024))
		resp.Body.Close()

		// 每100个请求检查一次
		if (i+1)%100 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// 强制GC
	runtime.GC()
	time.Sleep(1 * time.Second)

	// 检查最终内存
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	allocatedMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
	t.Logf("HTTP connections allocated %.2f MB for %d requests", allocatedMB, numRequests)

	// 内存增长不应过多
	if allocatedMB > 50 {
		t.Errorf("Excessive memory usage: %.2f MB allocated", allocatedMB)
	}
}

// TestDNSCacheAOFMemory 测试DNS缓存AOF的内存使用
func TestDNSCacheAOFMemory(t *testing.T) {
	defer goleak.VerifyNone(t)

	config := &dnscache.Config{
		AOFPath:      "/tmp/test_aof.aof",
		DefaultTTL:   1 * time.Minute,
		AOFInterval:  100 * time.Millisecond,
		Enabled:      true,
	}

	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	defer cache.Close()

	// 记录初始内存
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 测试AOF文件内存使用
	for i := 0; i < 1000; i++ {
		domain := fmt.Sprintf("aof-test-%d.com", i)
		cache.Set(domain, []string{"1.2.3.4"})
		time.Sleep(10 * time.Millisecond)
	}

	// 等待AOF写入完成
	time.Sleep(2 * time.Second)

	// 强制GC
	runtime.GC()

	// 检查最终内存
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	allocatedMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
	t.Logf("DNS Cache AOF allocated %.2f MB for 1000 operations", allocatedMB)

	if allocatedMB > 20 {
		t.Errorf("Excessive memory usage in AOF: %.2f MB allocated", allocatedMB)
	}
}

// BenchmarkDNSCacheMemory DNS缓存内存基准测试
func BenchmarkDNSCacheMemory(b *testing.B) {
	config := &dnscache.Config{
		DefaultTTL: 10 * time.Minute,
		Enabled:    true,
	}

	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}
	defer cache.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		domain := fmt.Sprintf("bench-test-%d.com", i%1000)
		cache.Set(domain, []string{"1.2.3.4", "5.6.7.8"})
		cache.Get(domain)
	}
}

// BenchmarkProxyRequests 代理请求基准测试
func BenchmarkProxyRequests(b *testing.B) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}))
	defer testServer.Close()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get(testServer.URL)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}

// TestMemoryStressTest 内存压力测试
func TestMemoryStressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	defer goleak.VerifyNone(t, goleak.IgnoreCurrent())

	// 创建测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Stress test response")
	}))
	defer testServer.Close()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 记录初始内存
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 执行压力测试
	numRequests := 5000
	completedRequests := 0

	for i := 0; i < numRequests; i++ {
		resp, err := client.Get(testServer.URL)
		if err == nil {
			resp.Body.Close()
			completedRequests++
		}

		// 每1000个请求检查内存
		if (i+1)%1000 == 0 {
			var mTemp runtime.MemStats
			runtime.ReadMemStats(&mTemp)
			tempMB := float64(mTemp.Alloc-m1.Alloc) / 1024 / 1024
			t.Logf("After %d requests: %.2f MB allocated, completed: %d",
				i+1, tempMB, completedRequests)

			// 内存使用过高时强制GC
			if tempMB > 100 {
				runtime.GC()
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	// 最终检查
	runtime.GC()
	time.Sleep(1 * time.Second)

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	finalMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
	t.Logf("Stress test completed: %d/%d requests, %.2f MB allocated",
		completedRequests, numRequests, finalMB)

	if finalMB > 200 {
		t.Errorf("Excessive memory usage after stress test: %.2f MB", finalMB)
	}
}