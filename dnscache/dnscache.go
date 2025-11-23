package dnscache

import (
	"bufio"
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
	// DefaultSaveInterval 默认全量保存间隔 30 秒
	DefaultSaveInterval = 30 * time.Second
	// DefaultAOFInterval 默认AOF增量保存间隔 1 秒
	DefaultAOFInterval = 1 * time.Second
)

// DNSCache DNS缓存管理器
type DNSCache struct {
	cache      *cache.Cache
	filePath   string
	aofPath    string
	mu         sync.RWMutex
	saveTicker *time.Ticker
	aofTicker  *time.Ticker
	done       chan bool
	wg         sync.WaitGroup
	aofFile    *os.File
	aofEncoder *json.Encoder
	closed     bool
}

// Record DNS记录结构 (用于可能的统计和调试)
type Record struct {
	Type   string        `json:"type"`
	Domain string        `json:"domain"`
	Value  interface{}   `json:"value"`
	TTL    time.Duration `json:"ttl"`
}

// cacheItem 用于序列化的缓存项
type cacheItem struct {
	Value      interface{} `json:"value"`
	Expiration time.Time   `json:"expiration"`
}

// AOFEntry AOF日志条目
type AOFEntry struct {
	Timestamp time.Time   `json:"timestamp"`
	Operation string      `json:"operation"` // "SET" or "DELETE"
	Key       string      `json:"key"`
	Value     interface{} `json:"value,omitempty"`
	TTL       int64       `json:"ttl,omitempty"` // TTL秒数
}

// MarshalJSON 自定义JSON序列化，处理无效时间
func (ci cacheItem) MarshalJSON() ([]byte, error) {
	type Alias cacheItem
	aux := struct {
		*Alias
		Expiration int64 `json:"expiration"`
	}{
		Alias: (*Alias)(&ci),
	}

	// 如果时间是零值，使用特殊标记
	if ci.Expiration.IsZero() {
		aux.Expiration = 0
	} else {
		// 验证年份是否有效
		year := ci.Expiration.Year()
		if year < 0 || year > 9999 {
			aux.Expiration = 0 // 使用0作为安全默认值
		} else {
			aux.Expiration = ci.Expiration.Unix()
		}
	}

	return json.Marshal(&aux)
}

// UnmarshalJSON 自定义JSON反序列化，处理时间戳
func (ci *cacheItem) UnmarshalJSON(data []byte) error {
	type Alias cacheItem
	aux := struct {
		*Alias
		Expiration int64 `json:"expiration"`
	}{
		Alias: (*Alias)(ci),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 处理时间戳
	if aux.Expiration == 0 {
		ci.Expiration = time.Time{} // 零值表示永不过期
	} else {
		// 验证时间戳范围
		if aux.Expiration < -62135596800 || aux.Expiration > 253402300799 { // 约束在有效年份范围内
			ci.Expiration = time.Time{} // 无效时间戳使用零值
		} else {
			ci.Expiration = time.Unix(aux.Expiration, 0)
			// 再次验证年份
			year := ci.Expiration.Year()
			if year < 0 || year > 9999 {
				ci.Expiration = time.Time{} // 无效年份使用零值
			}
		}
	}

	return nil
}

// Config DNS缓存配置
type Config struct {
	FilePath        string        `json:"file_path"`
	AOFPath         string        `json:"aof_path"`
	DefaultTTL      time.Duration `json:"default_ttl"`
	CleanupInterval time.Duration `json:"cleanup_interval"`
	SaveInterval    time.Duration `json:"save_interval"`
	AOFInterval     time.Duration `json:"aof_interval"`
	Enabled         bool          `json:"enabled"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		FilePath:        "./dns_cache.json",
		AOFPath:         "./dns_cache.aof",
		DefaultTTL:      DefaultTTL,
		CleanupInterval: DefaultCleanupInterval,
		SaveInterval:    DefaultSaveInterval,
		AOFInterval:     DefaultAOFInterval,
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
		aofPath:    config.AOFPath,
		done:       make(chan bool),
		saveTicker: time.NewTicker(config.SaveInterval),
		aofTicker:  time.NewTicker(config.AOFInterval),
	}

	// 初始化AOF文件
	if err := dc.initAOF(); err != nil {
		return nil, fmt.Errorf("初始化AOF文件失败: %w", err)
	}

	// 加载已有缓存
	if err := dc.Load(); err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("警告: 加载缓存失败: %v，将创建新的缓存\n", err)
		}
	}

	// 启动定期保存任务
	dc.wg.Add(2)
	go dc.periodicSave()
	go dc.periodicAOFCheckpoint()

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

// convertToStringSlice 将任意类型转换为字符串切片
// 处理JSON反序列化后可能出现的[]interface{}情况
func convertToStringSlice(value interface{}) []string {
	switch v := value.(type) {
	case []string:
		return v
	case []interface{}:
		result := make([]string, 0, len(v))
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	default:
		return nil
	}
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

	// 追加到AOF日志
	go dc.appendAOF("SET", key, ips, int64(ttl.Seconds()))
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

	// 追加到AOF日志
	go dc.appendAOF("SET", key, value, int64(ttl.Seconds()))
}

// Delete 删除DNS记录
func (dc *DNSCache) Delete(dnsType, domain string) {
	if dc.cache == nil {
		return
	}

	key := dc.makeKey(dnsType, domain)
	dc.cache.Delete(key)

	// 追加到AOF日志
	go dc.appendAOF("DELETE", key, nil, 0)
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
		// go-cache 的 Expiration 是 int64 类型的 Unix 时间戳（纳秒精度），或者 0 表示永不过期
		// 某些情况下可能是纳秒时间戳而不是秒时间戳
		if item.Expiration == 0 {
			// 对于永不过期的项，使用零时间
			expirationTime := time.Time{}
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
		} else {
			// 尝试将时间戳转换为有效的时间
			var expirationTime time.Time
			var unixSeconds int64

			// 判断是秒时间戳还是纳秒时间戳
			if item.Expiration > 1e18 { // 纳秒时间戳
				unixSeconds = item.Expiration / 1e9
			} else { // 秒时间戳
				unixSeconds = item.Expiration
			}

			// 验证时间戳范围
			if unixSeconds < 0 || unixSeconds > 253402300799 { // 9999-12-31 23:59:59 UTC
				fmt.Printf("警告: 跳过无效的过期时间戳 %d for key %s (转换为秒: %d)\n", item.Expiration, k, unixSeconds)
				continue
			}

			expirationTime = time.Unix(unixSeconds, 0)

			// 验证是否未过期
			if expirationTime.After(now) {
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

	// 尝试原子重命名
	if err := os.Rename(tempFile, dc.filePath); err != nil {
		// 如果重命名失败（常见于Docker bind mount），使用复制和删除的备选方案
		fmt.Printf("重命名失败，尝试备选方案: %v\n", err)

		// 先读取临时文件内容
		tempData, readErr := os.ReadFile(tempFile)
		if readErr != nil {
			return fmt.Errorf("读取临时文件失败: %w", readErr)
		}

		// 直接写入目标文件
		if writeErr := os.WriteFile(dc.filePath, tempData, 0644); writeErr != nil {
			return fmt.Errorf("直接写入目标文件失败: %w", writeErr)
		}

		// 删除临时文件
		if removeErr := os.Remove(tempFile); removeErr != nil {
			// 删除失败不作为致命错误，只记录警告
			fmt.Printf("警告: 删除临时文件失败: %v\n", removeErr)
		}
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

	// 1. 先加载全量缓存文件
	if err := dc.loadSnapshot(); err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("警告: 加载快照文件失败: %v\n", err)
		}
	}

	// 2. 重放AOF日志，恢复增量数据
	if err := dc.replayAOF(); err != nil {
		fmt.Printf("警告: 重放AOF日志失败: %v\n", err)
	}

	return nil
}

// loadSnapshot 加载快照文件
func (dc *DNSCache) loadSnapshot() error {
	data, err := os.ReadFile(dc.filePath)
	if err != nil {
		return err
	}

	// 检查文件是否为空
	if len(data) == 0 {
		fmt.Printf("缓存快照文件为空\n")
		return nil
	}

	// 检查是否只包含空白字符
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 {
		fmt.Printf("缓存快照文件只包含空白字符\n")
		return nil
	}

	var fileItems map[string]cacheItem
	if err := json.Unmarshal(data, &fileItems); err != nil {
		return fmt.Errorf("快照文件反序列化失败: %w", err)
	}

	// 检查是否解析出了有效的数据结构
	if fileItems == nil {
		fmt.Printf("缓存快照文件格式无效\n")
		return nil
	}

	now := time.Now()
	loadedCount := 0

	for k, item := range fileItems {
		// 如果过期时间是零值，表示永不过期，使用默认TTL加载
		if item.Expiration.IsZero() {
			// 尝试将字符串IP转换为net.IP
			if ipsStr := convertToStringSlice(item.Value); len(ipsStr) > 0 {
				var ips []net.IP
				for _, ipStr := range ipsStr {
					if ip := net.ParseIP(ipStr); ip != nil {
						ips = append(ips, ip)
					}
				}
				if len(ips) > 0 {
					dc.cache.Set(k, ips, DefaultTTL) // 使用默认TTL
					loadedCount++
				}
			} else {
				dc.cache.Set(k, item.Value, DefaultTTL) // 使用默认TTL
				loadedCount++
			}
			continue
		}

		// 验证过期时间是否有效
		if item.Expiration.Year() < 0 || item.Expiration.Year() > 9999 {
			fmt.Printf("警告: 跳过无效的过期时间 for key %s: %v\n", k, item.Expiration)
			continue
		}

		if item.Expiration.After(now) {
			ttl := time.Until(item.Expiration)
			// 尝试将字符串IP转换为net.IP
			if ipsStr := convertToStringSlice(item.Value); len(ipsStr) > 0 {
				var ips []net.IP
				for _, ipStr := range ipsStr {
					if ip := net.ParseIP(ipStr); ip != nil {
						ips = append(ips, ip)
					}
				}
				if len(ips) > 0 {
					dc.cache.Set(k, ips, ttl)
					loadedCount++
				}
			} else {
				dc.cache.Set(k, item.Value, ttl)
				loadedCount++
			}
		}
	}

	fmt.Printf("快照文件加载完成，共加载 %d 条记录\n", loadedCount)
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
	if dc.cache == nil || dc.closed {
		return
	}

	dc.closed = true
	dc.saveTicker.Stop()
	dc.aofTicker.Stop()
	close(dc.done)
	dc.wg.Wait()

	// 关闭AOF文件
	if dc.aofFile != nil {
		dc.aofFile.Close()
		dc.aofFile = nil
	}

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
			"enabled":    false,
			"item_count": 0,
			"file_path":  "",
			"aof_path":   "",
		}
	}

	return map[string]interface{}{
		"enabled":    true,
		"item_count": dc.cache.ItemCount(),
		"file_path":  dc.filePath,
		"aof_path":   dc.aofPath,
	}
}

// AOF相关方法

// initAOF 初始化AOF文件
func (dc *DNSCache) initAOF() error {
	file, err := os.OpenFile(dc.aofPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	dc.aofFile = file
	dc.aofEncoder = json.NewEncoder(file)

	return nil
}

// appendAOF 追加AOF日志条目
func (dc *DNSCache) appendAOF(operation, key string, value interface{}, ttl int64) error {
	if dc.aofFile == nil || dc.aofEncoder == nil {
		return nil // AOF未初始化，忽略
	}

	entry := AOFEntry{
		Timestamp: time.Now(),
		Operation: operation,
		Key:       key,
		Value:     value,
		TTL:       ttl,
	}

	dc.mu.Lock()
	defer dc.mu.Unlock()

	return dc.aofEncoder.Encode(entry)
}

// replayAOF 重放AOF日志
func (dc *DNSCache) replayAOF() error {
	file, err := os.Open(dc.aofPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // AOF文件不存在，正常情况
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	replayedCount := 0

	for scanner.Scan() {
		var entry AOFEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			fmt.Printf("警告: AOF日志条目解析失败: %v\n", err)
			continue
		}

		// 根据操作类型重放
		switch entry.Operation {
		case "SET":
			ttl := time.Duration(entry.TTL) * time.Second
			if ttl <= 0 {
				ttl = DefaultTTL
			}

			// 特殊处理IP类型数据
			if ipsStr := convertToStringSlice(entry.Value); len(ipsStr) > 0 {
				var ips []net.IP
				for _, ipStr := range ipsStr {
					if ip := net.ParseIP(ipStr); ip != nil {
						ips = append(ips, ip)
					}
				}
				if len(ips) > 0 {
					dc.cache.Set(entry.Key, ips, ttl)
				}
			} else {
				dc.cache.Set(entry.Key, entry.Value, ttl)
			}
			replayedCount++

		case "DELETE":
			dc.cache.Delete(entry.Key)
			replayedCount++

		default:
			fmt.Printf("警告: 未知的AOF操作类型: %s\n", entry.Operation)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取AOF文件错误: %w", err)
	}

	fmt.Printf("AOF重放完成，共处理 %d 条记录\n", replayedCount)
	return nil
}

// periodicAOFCheckpoint 定期AOF检查点
func (dc *DNSCache) periodicAOFCheckpoint() {
	defer dc.wg.Done()

	for {
		select {
		case <-dc.aofTicker.C:
			// 检查是否需要压缩AOF
			if err := dc.checkAOFSize(); err != nil {
				fmt.Printf("AOF检查点错误: %v\n", err)
			}
		case <-dc.done:
			return
		}
	}
}

// checkAOFSize 检查AOF文件大小并在需要时压缩
func (dc *DNSCache) checkAOFSize() error {
	if dc.aofPath == "" {
		return nil
	}

	info, err := os.Stat(dc.aofPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	// 如果AOF文件超过10MB，进行压缩
	if info.Size() > 10*1024*1024 {
		fmt.Printf("AOF文件过大(%d bytes)，开始压缩...\n", info.Size())
		return dc.compactAOF()
	}

	return nil
}

// compactAOF 压缩AOF文件
func (dc *DNSCache) compactAOF() error {
	tempPath := dc.aofPath + ".compact"

	// 创建临时文件
	tempFile, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	encoder := json.NewEncoder(tempFile)

	// 获取当前所有有效缓存项
	items := dc.cache.Items()
	now := time.Now()

	for key, item := range items {
		var expirationTime time.Time
		var hasExpiration bool

		// 只包含未过期的项
		if item.Expiration > 0 {
			hasExpiration = true
			if item.Expiration > 1e18 { // 纳秒时间戳
				expirationTime = time.Unix(item.Expiration/1e9, 0)
			} else { // 秒时间戳
				expirationTime = time.Unix(item.Expiration, 0)
			}

			if expirationTime.Before(now) {
				continue // 已过期，跳过
			}
		}

		// 计算TTL
		var ttl int64
		if item.Expiration == 0 {
			ttl = int64(DefaultTTL.Seconds())
		} else if hasExpiration {
			ttl = int64(time.Until(expirationTime).Seconds())
			if ttl <= 0 {
				continue // 已过期，跳过
			}
		} else {
			continue // 无效的过期时间，跳过
		}

		// 写入压缩后的AOF条目
		entry := AOFEntry{
			Timestamp: time.Now(),
			Operation: "SET",
			Key:       key,
			Value:     item.Object,
			TTL:       ttl,
		}

		if err := encoder.Encode(entry); err != nil {
			return err
		}
	}

	// 原子替换文件
	if err := tempFile.Close(); err != nil {
		return err
	}

	// 关闭当前AOF文件
	if dc.aofFile != nil {
		dc.aofFile.Close()
	}

	// 重命名
	if err := os.Rename(tempPath, dc.aofPath); err != nil {
		// 如果重命名失败，尝试复制方案
		data, err := os.ReadFile(tempPath)
		if err != nil {
			return err
		}

		if err := os.WriteFile(dc.aofPath, data, 0644); err != nil {
			return err
		}

		os.Remove(tempPath)
	}

	// 重新打开AOF文件
	return dc.initAOF()
}
