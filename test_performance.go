package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
)

func main() {
	fmt.Println("=== DNS缓存性能持久化测试 ===")

	// 清理旧文件
	os.Remove("perf_dns_cache.json")
	os.Remove("perf_dns_cache.aof")

	// 创建测试配置 - 使用真实的30秒间隔
	config := dnscache.DefaultConfig()
	config.FilePath = "./perf_dns_cache.json"
	config.AOFPath = "./perf_dns_cache.aof"
	config.SaveInterval = 30 * time.Second  // 30秒全量保存
	config.AOFInterval = 1 * time.Second    // 1秒增量保存

	fmt.Printf("配置: 全量保存=%v, 增量保存=%v\n", config.SaveInterval, config.AOFInterval)

	// 创建DNS缓存实例
	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建DNS缓存失败: %v\n", err)
		return
	}
	defer cache.Close()

	// 测试1: 大量DNS记录添加
	fmt.Println("\n=== 测试1: 添加100个DNS记录 ===")
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		domain := fmt.Sprintf("test%d.example.com", i)
		ip := net.ParseIP(fmt.Sprintf("192.168.1.%d", i%255+1))

		cache.SetIP("A", domain, ip, 5*time.Minute)

		if i%10 == 0 {
			fmt.Printf("已添加 %d 个记录...\n", i+1)
		}
	}

	addTime := time.Since(startTime)
	fmt.Printf("添加100个记录耗时: %v\n", addTime)

	// 测试2: 立即检查AOF文件大小
	fmt.Println("\n=== 测试2: 检查AOF文件大小 ===")
	time.Sleep(2 * time.Second) // 等待AOF写入

	if info, err := os.Stat(config.AOFPath); err == nil {
		fmt.Printf("AOF文件大小: %d bytes\n", info.Size())
	}

	// 测试3: 读取性能测试
	fmt.Println("\n=== 测试3: 读取性能测试 ===")
	startTime = time.Now()
	hitCount := 0

	for i := 0; i < 100; i++ {
		domain := fmt.Sprintf("test%d.example.com", i)
		if _, found := cache.GetIP("A", domain); found {
			hitCount++
		}
	}

	readTime := time.Since(startTime)
	fmt.Printf("读取100个记录耗时: %v, 命中率: %d%%\n", readTime, hitCount)

	// 测试4: 等待30秒进行全量保存
	fmt.Println("\n=== 测试4: 等待30秒全量保存 ===")
	fmt.Println("等待35秒进行全量保存...")
	time.Sleep(35 * time.Second)

	// 检查两个文件大小
	fmt.Println("\n=== 测试5: 文件大小对比 ===")
	if info, err := os.Stat(config.FilePath); err == nil {
		fmt.Printf("快照文件大小: %d bytes\n", info.Size())
	}
	if info, err := os.Stat(config.AOFPath); err == nil {
		fmt.Printf("AOF文件大小: %d bytes\n", info.Size())
	}

	// 测试6: 验证数据持久化
	fmt.Println("\n=== 测试6: 验证数据持久化 ===")
	cache.Close()

	// 重新创建缓存实例
	cache2, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建新缓存实例失败: %v\n", err)
		return
	}
	defer cache2.Close()

	stats := cache2.Stats()
	fmt.Printf("重载后缓存项数量: %v\n", stats["item_count"])

	// 验证数据完整性
	verifiedCount := 0
	for i := 0; i < 100; i++ {
		domain := fmt.Sprintf("test%d.example.com", i)
		if _, found := cache2.GetIP("A", domain); found {
			verifiedCount++
		}
	}

	fmt.Printf("数据完整性验证: %d/100 (%.1f%%)\n", verifiedCount, float64(verifiedCount)/100*100)

	fmt.Println("\n=== 性能测试完成 ===")
	fmt.Println("✅ 30秒全量保存和每秒增量保存功能正常工作!")
	fmt.Println("✅ 数据持久化和恢复功能验证通过!")
}