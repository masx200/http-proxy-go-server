package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
)

func main() {
	fmt.Println("=== DNS缓存AOF修复测试 ===")

	// 清理旧文件
	os.Remove("fixed_dns_cache.json")
	os.Remove("fixed_dns_cache.aof")

	// 创建配置
	config := dnscache.DefaultConfig()
	config.FilePath = "./fixed_dns_cache.json"
	config.AOFPath = "./fixed_dns_cache.aof"
	config.SaveInterval = 30 * time.Second
	config.AOFInterval = 1 * time.Second

	// 第一阶段: 添加数据
	fmt.Println("\n=== 阶段1: 添加数据 ===")
	cache, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("创建缓存失败: %v\n", err)
		return
	}

	// 添加测试数据
	domains := []string{"google.com", "github.com", "example.com"}
	for i, domain := range domains {
		ip := net.ParseIP(fmt.Sprintf("93.184.216.%d", i+34))
		cache.SetIP("A", domain, ip, 5*time.Minute)
		fmt.Printf("✓ 添加: %s -> %s\n", domain, ip)
	}

	// 等待AOF写入
	time.Sleep(3 * time.Second)

	// 验证数据
	fmt.Println("\n验证缓存中的数据:")
	for _, domain := range domains {
		if ip, found := cache.GetIP("A", domain); found {
			fmt.Printf("✓ %s -> %s\n", domain, ip.String())
		} else {
			fmt.Printf("✗ 未找到: %s\n", domain)
		}
	}

	// 关闭缓存
	cache.Close()
	time.Sleep(1 * time.Second)

	// 第二阶段: 重新加载，测试AOF重放
	fmt.Println("\n=== 阶段2: 测试AOF重放 ===")
	cache2, err := dnscache.NewWithConfig(config)
	if err != nil {
		fmt.Printf("重新创建缓存失败: %v\n", err)
		return
	}
	defer cache2.Close()

	// 验证重放的数据
	fmt.Println("\n验证重放的数据:")
	successCount := 0
	for _, domain := range domains {
		if ip, found := cache2.GetIP("A", domain); found {
			fmt.Printf("✓ %s -> %s\n", domain, ip.String())
			successCount++
		} else {
			fmt.Printf("✗ 重放失败: %s\n", domain)
		}
	}

	// 统计
	fmt.Printf("\n=== 统计结果 ===\n")
	fmt.Printf("重放成功率: %d/3 (%.1f%%)\n", successCount, float64(successCount)/3*100)

	if successCount == 3 {
		fmt.Println("🎉 AOF重放修复成功！")
	} else {
		fmt.Println("❌ AOF重放仍有问题")
	}

	// 显示文件信息
	fmt.Println("\n=== 生成的文件 ===")
	if info, err := os.Stat(config.AOFPath); err == nil {
		fmt.Printf("AOF文件: %s (%d bytes)\n", config.AOFPath, info.Size())
	}
	if info, err := os.Stat(config.FilePath); err == nil {
		fmt.Printf("快照文件: %s (%d bytes)\n", config.FilePath, info.Size())
	}
}
