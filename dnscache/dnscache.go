package dnscache

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	// DefaultTTL 默认缓存时间 10 分钟
	DefaultTTL = 10 * time.Minute
	// DefaultCleanupInterval 默认清理间隔 5 分钟
	DefaultCleanupInterval = 5 * time.Minute
	// DefaultSaveInterval 默认自动保存间隔 30 秒
	DefaultSaveInterval = 30 * time.Second
)

// DNSCache DNS缓存管理器
type DNSCache struct {
	cache      *cache.Cache
	filePath   string
	mu         sync.RWMutex
	saveTicker *time.Ticker
	done       chan bool
	wg         sync.WaitGroup
}

// Record DNS记录结构 (用于可能的统计和调试)
type Record struct {
	Type   string      `json:"type"`
	Domain string      `json:"domain"`
	Value  interface{} `json:"value"`
	TTL    time.Duration `json:"ttl"`
}

// cacheItem 用于序列化的缓存项
type cacheItem struct {
	Value      interface{} `json:"value"`
	Expiration time.Time   `json:"expiration"`
}

// Config DNS缓存配置
type Config struct {
	FilePath         string        `json:"file_path"`
	DefaultTTL       time.Duration `json:"default_ttl"`
	CleanupInterval  time.Duration `json:"cleanup_interval"`
	SaveInterval     time.Duration `json:"save_interval"`
	Enabled          bool          `json:"enabled"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		FilePath:        "./dns_cache.json",
		DefaultTTL:      DefaultTTL,
		CleanupInterval: DefaultCleanupInterval,
		SaveInterval:    DefaultSaveInterval,
		Enabled:         true,
	}
}

// New 创建新的DNS缓存实例
func New(filePath string) (*DNSCache, error) {
	config := DefaultConfig()
	config.FilePath = filePath
	return NewWithConfig(config)
}

// NewWithConfig 使用配置创建DNS缓存实例
func NewWithConfig(config *Config) (*DNSCache, error) {
	if !config.Enabled {
		return &DNSCache{}, nil
	}

	// 确保缓存目录存在
	dir := filepath.Dir(config.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("创建缓存目录失败: %w", err)
	}

	dc := &DNSCache{
		cache:      cache.New(config.DefaultTTL, config.CleanupInterval),
		filePath:   config.FilePath,
		done:       make(chan bool),
		saveTicker: time.NewTicker(config.SaveInterval),
	}

	// 加载已有缓存
	if err := dc.Load(); err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("警告: 加载缓存失败: %v，将创建新的缓存\n", err)
		}
	}

	// 启动定期保存任务
	dc.wg.Add(1)
	go dc.periodicSave()

	return dc, nil
}

// makeKey 生成标准化缓存键
func (dc *DNSCache) makeKey(dnsType, domain string) string {
	return fmt.Sprintf("%s:%s", strings.ToUpper(dnsType), normalizeDomain(domain))
}

// normalizeDomain 标准化域名格式
func normalizeDomain(domain string) string {
	domain = strings.TrimSpace(domain)
	domain = strings.TrimSuffix(domain, ".")
	domain = strings.ToLower(domain)
	return domain
}

// GetIPs 获取DNS记录（IP地址列表）
func (dc *DNSCache) GetIPs(dnsType, domain string) ([]net.IP, bool) {
	if dc.cache == nil {
		return nil, false
	}

	key := dc.makeKey(dnsType, domain)
	if value, found := dc.cache.Get(key); found {
		if ips, ok := value.([]net.IP); ok {
			return ips, true
		}
		// 尝试转换字符串格式的IP
		if ipsStr, ok := value.([]string); ok {
			var ips []net.IP
			for _, ipStr := range ipsStr {
				if ip := net.ParseIP(ipStr); ip != nil {
					ips = append(ips, ip)
				}
			}
			if len(ips) > 0 {
				// 更新缓存为IP格式
				dc.SetIPs(dnsType, domain, ips, 0)
				return ips, true
			}
		}
	}
	return nil, false
}

// GetIP 获取单个IP地址
func (dc *DNSCache) GetIP(dnsType, domain string) (net.IP, bool) {
	ips, found := dc.GetIPs(dnsType, domain)
	if !found || len(ips) == 0 {
		return nil, false
	}
	return ips[0], true
}

// SetIPs 设置DNS记录（IP地址列表）
func (dc *DNSCache) SetIPs(dnsType, domain string, ips []net.IP, ttl time.Duration) {
	if dc.cache == nil {
		return
	}

	if ttl <= 0 {
		ttl = DefaultTTL
	}
	key := dc.makeKey(dnsType, domain)
	dc.cache.Set(key, ips, ttl)
}

// SetIP 设置单个IP地址
func (dc *DNSCache) SetIP(dnsType, domain string, ip net.IP, ttl time.Duration) {
	dc.SetIPs(dnsType, domain, []net.IP{ip}, ttl)
}

// Get 获取通用DNS记录
func (dc *DNSCache) Get(dnsType, domain string) (interface{}, bool) {
	if dc.cache == nil {
		return nil, false
	}

	key := dc.makeKey(dnsType, domain)
	return dc.cache.Get(key)
}

// Set 设置通用DNS记录
func (dc *DNSCache) Set(dnsType, domain string, value interface{}, ttl time.Duration) {
	if dc.cache == nil {
		return
	}

	if ttl <= 0 {
		ttl = DefaultTTL
	}
	key := dc.makeKey(dnsType, domain)
	dc.cache.Set(key, value, ttl)
}

// Delete 删除DNS记录
func (dc *DNSCache) Delete(dnsType, domain string) {
	if dc.cache == nil {
		return
	}

	key := dc.makeKey(dnsType, domain)
	dc.cache.Delete(key)
}

// Save 保存缓存到文件（原子操作）
func (dc *DNSCache) Save() error {
	if dc.cache == nil {
		return nil
	}

	dc.mu.Lock()
	defer dc.mu.Unlock()

	items := dc.cache.Items()
	now := time.Now()
	validItems := make(map[string]cacheItem)

	// 只保存未过期的项
	for k, item := range items {
		// go-cache 的 Expiration 是 int64 类型的 Unix 时间戳，或者 0 表示永不过期
		if item.Expiration == 0 || item.Expiration > now.Unix() {
			expirationTime := time.Unix(item.Expiration, 0)
			// 特殊处理IP列表，确保序列化格式一致
			if ips, ok := item.Object.([]net.IP); ok {
				var ipsStr []string
				for _, ip := range ips {
					ipsStr = append(ipsStr, ip.String())
				}
				validItems[k] = cacheItem{
					Value:      ipsStr,
					Expiration: expirationTime,
				}
			} else {
				validItems[k] = cacheItem{
					Value:      item.Object,
					Expiration: expirationTime,
				}
			}
		}
	}

	// 如果没有有效数据，删除文件
	if len(validItems) == 0 {
		if err := os.Remove(dc.filePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("删除空缓存文件失败: %w", err)
		}
		return nil
	}

	// 写入临时文件后重命名，保证原子性
	data, err := json.MarshalIndent(validItems, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化失败: %w", err)
	}

	tempFile := dc.filePath + ".tmp"
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("写入临时文件失败: %w", err)
	}

	if err := os.Rename(tempFile, dc.filePath); err != nil {
		return fmt.Errorf("重命名文件失败: %w", err)
	}

	return nil
}

// Load 从文件加载缓存
func (dc *DNSCache) Load() error {
	if dc.cache == nil {
		return nil
	}

	dc.mu.Lock()
	defer dc.mu.Unlock()

	data, err := os.ReadFile(dc.filePath)
	if err != nil {
		return err
	}

	var fileItems map[string]cacheItem
	if err := json.Unmarshal(data, &fileItems); err != nil {
		return fmt.Errorf("反序列化失败: %w", err)
	}

	now := time.Now()
	for k, item := range fileItems {
		if item.Expiration.After(now) {
			ttl := time.Until(item.Expiration)
			// 尝试将字符串IP转换为net.IP
			if ipsStr, ok := item.Value.([]string); ok {
				var ips []net.IP
				for _, ipStr := range ipsStr {
					if ip := net.ParseIP(ipStr); ip != nil {
						ips = append(ips, ip)
					}
				}
				if len(ips) > 0 {
					dc.cache.Set(k, ips, ttl)
				}
			} else {
				dc.cache.Set(k, item.Value, ttl)
			}
		}
	}

	return nil
}

// periodicSave 定期保存任务
func (dc *DNSCache) periodicSave() {
	defer dc.wg.Done()

	for {
		select {
		case <-dc.saveTicker.C:
			if err := dc.Save(); err != nil {
				fmt.Printf("定期保存缓存失败: %v\n", err)
			}
		case <-dc.done:
			return
		}
	}
}

// Close 关闭缓存（会触发最后一次保存）
func (dc *DNSCache) Close() {
	if dc.cache == nil {
		return
	}

	dc.saveTicker.Stop()
	close(dc.done)
	dc.wg.Wait()

	// 关闭时保存
	if err := dc.Save(); err != nil {
		fmt.Printf("关闭时保存缓存失败: %v\n", err)
	}
}

// Flush 清空所有缓存
func (dc *DNSCache) Flush() {
	if dc.cache == nil {
		return
	}

	dc.cache.Flush()
}

// ItemCount 返回缓存项数量
func (dc *DNSCache) ItemCount() int {
	if dc.cache == nil {
		return 0
	}

	return dc.cache.ItemCount()
}

// Stats 返回统计信息
func (dc *DNSCache) Stats() map[string]interface{} {
	if dc.cache == nil {
		return map[string]interface{}{
			"enabled":   false,
			"item_count": 0,
			"file_path":  "",
		}
	}

	return map[string]interface{}{
		"enabled":    true,
		"item_count": dc.cache.ItemCount(),
		"file_path":  dc.filePath,
	}
}