package main

import (
	"fmt"
	"net"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
)

func main() {
	fmt.Println("=== 简单AOF测试 ===")

	// 创建测试配置
	config := dnscache.DefaultConfig()
	config.FilePath = "./simple_dns_cache.json"
	config.AOFPath = "./simple_dns_cache.aof"
	config.SaveInterval = 60 * time.Second // 延长全量保存间隔
	config.AOFInterval = 1 * time.Second   // 1秒增量保存

	fmt.Printf("配置: 快照=%s, AOF=%s\n", config.FilePath, config.AOFPath)

	// 创建DNS缓存实例
	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建DNS缓存失败: %v\n", err)
		return
	}
	defer cache.Close()

	fmt.Println("DNS缓存创建成功！")

	// 测试1: 添加一个DNS记录
	fmt.Println("\n=== 测试1: 添加DNS记录 ===")
	domain := "test.example.com"
	ip := net.ParseIP("1.2.3.4")
	cache.SetIP("A", domain, ip, 5*time.Minute)
	fmt.Printf("设置 %s -> %s\n", domain, ip)

	// 等待AOF写入
	time.Sleep(2 * time.Second)

	// 测试2: 读取记录
	fmt.Println("\n=== 测试2: 读取记录 ===")
	retrievedIP, found := cache.GetIP("A", domain)
	if found {
		fmt.Printf("读取 %s -> %s ✓\n", domain, retrievedIP)
	} else {
		fmt.Printf("未找到 %s ✗\n", domain)
	}

	// 等待AOF写入
	time.Sleep(2 * time.Second)

	// 测试3: 删除记录
	fmt.Println("\n=== 测试3: 删除记录 ===")
	cache.Delete("A", domain)
	fmt.Printf("删除 %s\n", domain)

	// 等待AOF写入
	time.Sleep(2 * time.Second)

	// 测试4: 重新创建缓存，测试AOF重放
	fmt.Println("\n=== 测试4: 测试AOF重放 ===")
	cache.Close()

	time.Sleep(1 * time.Second)

	cache2, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建新缓存实例失败: %v\n", err)
		return
	}
	defer cache2.Close()

	// 检查记录是否被删除
	retrievedIP, found = cache2.GetIP("A", domain)
	if found {
		fmt.Printf("错误: %s 仍然存在 -> %s ✗\n", domain, retrievedIP)
	} else {
		fmt.Printf("正确: %s 已被删除 ✓\n", domain)
	}

	stats := cache2.Stats()
	fmt.Printf("缓存统计: %v\n", stats)

	fmt.Println("\n=== 测试完成 ===")
}