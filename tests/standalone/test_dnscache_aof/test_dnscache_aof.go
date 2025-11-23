package main

import (
	"fmt"
	"net"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
)

func main() {
	fmt.Println("=== DNS缓存AOF持久化功能测试 ===")

	// 创建测试配置
	config := dnscache.DefaultConfig()
	config.FilePath = "./test_dns_cache.json"
	config.AOFPath = "./test_dns_cache.aof"
	config.SaveInterval = 5 * time.Second // 5秒全量保存，测试用
	config.AOFInterval = 1 * time.Second  // 1秒增量保存

	// 创建DNS缓存实例
	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建DNS缓存失败: %v\n", err)
		return
	}
	defer cache.Close()

	fmt.Println("DNS缓存创建成功！")
	fmt.Printf("配置信息:\n")
	fmt.Printf("- 全量保存间隔: %v\n", config.SaveInterval)
	fmt.Printf("- 增量保存间隔: %v\n", config.AOFInterval)
	fmt.Printf("- 快照文件: %s\n", config.FilePath)
	fmt.Printf("- AOF文件: %s\n", config.AOFPath)

	// 测试1: 添加一些DNS记录
	fmt.Println("\n=== 测试1: 添加DNS记录 ===")
	testRecords := []struct {
		domain string
		ips    []string
		ttl    time.Duration
	}{
		{"example.com", []string{"93.184.216.34"}, 5 * time.Minute},
		{"google.com", []string{"142.250.191.14", "142.250.191.78"}, 10 * time.Minute},
		{"localhost", []string{"127.0.0.1", "::1"}, time.Hour},
	}

	for _, record := range testRecords {
		var ips []net.IP
		for _, ipStr := range record.ips {
			if ip := net.ParseIP(ipStr); ip != nil {
				ips = append(ips, ip)
			}
		}

		cache.SetIPs("A", record.domain, ips, record.ttl)
		fmt.Printf("设置 %s -> %v (TTL: %v)\n", record.domain, record.ips, record.ttl)

		// 短暂延迟，确保AOF写入
		time.Sleep(100 * time.Millisecond)
	}

	// 测试2: 读取缓存
	fmt.Println("\n=== 测试2: 读取缓存 ===")
	for _, record := range testRecords {
		ips, found := cache.GetIPs("A", record.domain)
		if found {
			var ipStrs []string
			for _, ip := range ips {
				ipStrs = append(ipStrs, ip.String())
			}
			fmt.Printf("读取 %s -> %v\n", record.domain, ipStrs)
		} else {
			fmt.Printf("未找到 %s 的缓存记录\n", record.domain)
		}
	}

	// 测试3: 删除记录
	fmt.Println("\n=== 测试3: 删除记录 ===")
	cache.Delete("A", "localhost")
	fmt.Println("删除 localhost 记录")

	// 测试4: 等待全量保存
	fmt.Println("\n=== 测试4: 等待全量保存 ===")
	fmt.Println("等待6秒进行全量保存...")
	time.Sleep(6 * time.Second)

	// 测试5: 显示统计信息
	fmt.Println("\n=== 测试5: 缓存统计 ===")
	stats := cache.Stats()
	for k, v := range stats {
		fmt.Printf("%s: %v\n", k, v)
	}

	// 测试6: 创建新缓存实例，测试加载
	fmt.Println("\n=== 测试6: 测试缓存加载 ===")
	fmt.Println("关闭当前缓存实例...")
	cache.Close()

	fmt.Println("创建新的缓存实例进行加载测试...")
	cache2, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建新缓存实例失败: %v\n", err)
		return
	}
	defer cache2.Close()

	fmt.Println("检查加载的记录:")
	for _, record := range testRecords[:2] { // 跳过已删除的localhost
		ips, found := cache2.GetIPs("A", record.domain)
		if found {
			var ipStrs []string
			for _, ip := range ips {
				ipStrs = append(ipStrs, ip.String())
			}
			fmt.Printf("✓ %s -> %v\n", record.domain, ipStrs)
		} else {
			fmt.Printf("✗ 未找到 %s 的缓存记录\n", record.domain)
		}
	}

	stats2 := cache2.Stats()
	fmt.Printf("新缓存实例统计: 项数量 = %v\n", stats2["item_count"])

	fmt.Println("\n=== 测试完成 ===")
}
