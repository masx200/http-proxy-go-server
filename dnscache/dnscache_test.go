package dnscache

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheItemMarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		item       cacheItem
		wantJSON   string
		shouldFail bool
	}{
		{
			name: "正常时间",
			item: cacheItem{
				Value:      "test-value",
				Expiration: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			wantJSON: `{"value":"test-value","expiration":1704110400}`,
		},
		{
			name: "零值时间（永不过期）",
			item: cacheItem{
				Value:      "test-value",
				Expiration: time.Time{},
			},
			wantJSON: `{"value":"test-value","expiration":0}`,
		},
		{
			name: "无效年份 - 小于0",
			item: cacheItem{
				Value:      "test-value",
				Expiration: time.Date(-100, 1, 1, 0, 0, 0, 0, time.UTC), // 无效年份
			},
			wantJSON: `{"value":"test-value","expiration":0}`,
		},
		{
			name: "无效年份 - 大于9999",
			item: cacheItem{
				Value:      "test-value",
				Expiration: time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC), // 无效年份
			},
			wantJSON: `{"value":"test-value","expiration":0}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.item)
			if (err != nil) != tt.shouldFail {
				t.Errorf("cacheItem.MarshalJSON() error = %v, shouldFail %v", err, tt.shouldFail)
				return
			}
			if !tt.shouldFail {
				gotJSON := string(data)
				if gotJSON != tt.wantJSON {
					t.Errorf("cacheItem.MarshalJSON() = %v, want %v", gotJSON, tt.wantJSON)
				}
			}
		})
	}
}

func TestCacheItemUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonData    string
		wantItem    cacheItem
		shouldFail  bool
	}{
		{
			name:     "正常时间戳",
			jsonData: `{"value":"test-value","expiration":1704110400}`,
			wantItem: cacheItem{
				Value:      "test-value",
				Expiration: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			},
		},
		{
			name:     "零值时间戳（永不过期）",
			jsonData: `{"value":"test-value","expiration":0}`,
			wantItem: cacheItem{
				Value:      "test-value",
				Expiration: time.Time{},
			},
		},
		{
			name:     "负时间戳（无效，应设为零值）",
			jsonData: `{"value":"test-value","expiration":-99999999999}`,
			wantItem: cacheItem{
				Value:      "test-value",
				Expiration: time.Time{},
			},
		},
		{
			name:     "超大时间戳（无效，应设为零值）",
			jsonData: `{"value":"test-value","expiration":999999999999}`,
			wantItem: cacheItem{
				Value:      "test-value",
				Expiration: time.Time{},
			},
		},
		{
			name:        "无效JSON",
			jsonData:    `{"value":"test-value","expiration":invalid}`,
			shouldFail:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item cacheItem
			err := json.Unmarshal([]byte(tt.jsonData), &item)
			if (err != nil) != tt.shouldFail {
				t.Errorf("cacheItem.UnmarshalJSON() error = %v, shouldFail %v", err, tt.shouldFail)
				return
			}
			if !tt.shouldFail {
				// 比较时间时考虑精度差异
				if item.Value != tt.wantItem.Value {
					t.Errorf("cacheItem.UnmarshalJSON() Value = %v, want %v", item.Value, tt.wantItem.Value)
				}
				if !item.Expiration.Equal(tt.wantItem.Expiration) {
					t.Errorf("cacheItem.UnmarshalJSON() Expiration = %v, want %v", item.Expiration, tt.wantItem.Expiration)
				}
			}
		})
	}
}

func TestDNSCacheSaveLoad(t *testing.T) {
	tempDir := t.TempDir()
	cacheFile := filepath.Join(tempDir, "test_cache.json")

	// 创建新的缓存实例
	cache, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	defer cache.Close()

	// 添加一些测试数据
	testIP1 := net.ParseIP("192.168.1.1")
	testIP2 := net.ParseIP("10.0.0.1")
	testIPs := []net.IP{testIP1, testIP2}

	cache.SetIPs("A", "example.com", testIPs, 5*time.Minute)
	cache.Set("TXT", "example.com", "v=spf1 include:_spf.example.com ~all", 10*time.Minute)

	// 保存到文件
	err = cache.Save()
	if err != nil {
		t.Fatalf("Failed to save cache: %v", err)
	}

	// 验证文件存在
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		t.Fatalf("Cache file was not created")
	}

	// 创建新的缓存实例并加载数据
	cache2, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create second cache: %v", err)
	}
	defer cache2.Close()

	// 验证加载的数据
	loadedIPs, found := cache2.GetIPs("A", "example.com")
	if !found {
		t.Fatalf("Expected to find IPs for example.com")
	}
	if len(loadedIPs) != len(testIPs) {
		t.Fatalf("Expected %d IPs, got %d", len(testIPs), len(loadedIPs))
	}

	loadedTXT, found := cache2.Get("TXT", "example.com")
	if !found {
		t.Fatalf("Expected to find TXT record for example.com")
	}
	if loadedTXT != "v=spf1 include:_spf.example.com ~all" {
		t.Fatalf("Expected TXT record content, got %v", loadedTXT)
	}
}

func TestDNSCacheWithInvalidTimestamps(t *testing.T) {
	tempDir := t.TempDir()
	cacheFile := filepath.Join(tempDir, "test_invalid_cache.json")

	// 创建包含无效时间戳的缓存文件
	invalidItems := map[string]cacheItem{
		"valid_item": {
			Value:      "valid-value",
			Expiration: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		"invalid_year_item": {
			Value:      "invalid-value",
			Expiration: time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC), // 无效年份
		},
		"zero_time_item": {
			Value:      "zero-value",
			Expiration: time.Time{}, // 零值
		},
	}

	// 手动序列化以创建包含无效时间的文件
	data, err := json.MarshalIndent(invalidItems, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	err = os.WriteFile(cacheFile, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write test cache file: %v", err)
	}

	// 尝试加载这个包含无效时间的缓存文件
	cache, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create cache with invalid timestamps: %v", err)
	}
	defer cache.Close()

	// 验证只有有效的项目被加载
	_, found := cache.Get("A", "valid_item")
	if found {
		t.Errorf("Expected valid_item to be loaded")
	}

	// 验证缓存统计
	stats := cache.Stats()
	if stats["item_count"] == 0 {
		t.Errorf("Expected some items to be loaded, got %d", stats["item_count"])
	}
}

func TestDNSCacheExpirationHandling(t *testing.T) {
	tempDir := t.TempDir()
	cacheFile := filepath.Join(tempDir, "test_expiration_cache.json")

	cache, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// 添加一个已过期的项目
	expiredIP := net.ParseIP("1.2.3.4")
	cache.SetIPs("A", "expired.com", []net.IP{expiredIP}, 1*time.Millisecond)

	// 等待过期
	time.Sleep(10 * time.Millisecond)

	// 添加一个未过期的项目
	validIP := net.ParseIP("5.6.7.8")
	cache.SetIPs("A", "valid.com", []net.IP{validIP}, 1*time.Hour)

	// 保存缓存
	err = cache.Save()
	if err != nil {
		t.Fatalf("Failed to save cache: %v", err)
	}

	cache.Close()

	// 创建新实例并加载
	cache2, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create second cache: %v", err)
	}
	defer cache2.Close()

	// 验证只有未过期的项目被加载
	_, found := cache2.GetIPs("A", "expired.com")
	if found {
		t.Errorf("Expected expired.com to not be loaded")
	}

	_, found = cache2.GetIPs("A", "valid.com")
	if !found {
		t.Errorf("Expected valid.com to be loaded")
	}
}

func TestDNSCacheKeyNormalization(t *testing.T) {
	cache, err := New("") // 使用内存缓存
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	testIP := net.ParseIP("192.168.1.1")

	// 使用不同的格式设置相同的域名
	cache.SetIPs("A", "example.com.", []net.IP{testIP}, 1*time.Hour)
	cache.SetIPs("A", "EXAMPLE.COM", []net.IP{testIP}, 1*time.Hour)
	cache.SetIPs("A", " example.com ", []net.IP{testIP}, 1*time.Hour)

	// 验证所有格式都指向相同的键
	_, found1 := cache.GetIPs("A", "example.com")
	_, found2 := cache.GetIPs("A", "EXAMPLE.COM")
	_, found3 := cache.GetIPs("A", " example.com ")

	if !found1 || !found2 || !found3 {
		t.Errorf("Expected all domain formats to be found")
	}

	// 验证缓存项数量只有1个
	if cache.ItemCount() != 1 {
		t.Errorf("Expected 1 cache item, got %d", cache.ItemCount())
	}
}

func TestDNSCacheConfiguration(t *testing.T) {
	config := DefaultConfig()

	if config.FilePath != "./dns_cache.json" {
		t.Errorf("Expected default file path './dns_cache.json', got %s", config.FilePath)
	}

	if config.DefaultTTL != DefaultTTL {
		t.Errorf("Expected default TTL %v, got %v", DefaultTTL, config.DefaultTTL)
	}

	if config.CleanupInterval != DefaultCleanupInterval {
		t.Errorf("Expected cleanup interval %v, got %v", DefaultCleanupInterval, config.CleanupInterval)
	}

	if config.SaveInterval != DefaultSaveInterval {
		t.Errorf("Expected save interval %v, got %v", DefaultSaveInterval, config.SaveInterval)
	}

	if !config.Enabled {
		t.Errorf("Expected cache to be enabled by default")
	}
}

func TestDNSCacheDisabled(t *testing.T) {
	config := &Config{
		Enabled: false,
	}

	cache, err := NewWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create disabled cache: %v", err)
	}

	// 禁用的缓存不应该保存任何数据
	testIP := net.ParseIP("192.168.1.1")
	cache.SetIPs("A", "example.com", []net.IP{testIP}, 1*time.Hour)

	_, found := cache.GetIPs("A", "example.com")
	if found {
		t.Errorf("Expected disabled cache to not store data")
	}

	if cache.ItemCount() != 0 {
		t.Errorf("Expected disabled cache to have 0 items, got %d", cache.ItemCount())
	}
}

func TestDNSCacheEmptyFileHandling(t *testing.T) {
	tempDir := t.TempDir()
	emptyCacheFile := filepath.Join(tempDir, "empty_cache.json")

	// 创建空文件
	err := os.WriteFile(emptyCacheFile, []byte("{}"), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty cache file: %v", err)
	}

	cache, err := New(emptyCacheFile)
	if err != nil {
		t.Fatalf("Failed to create cache from empty file: %v", err)
	}
	defer cache.Close()

	if cache.ItemCount() != 0 {
		t.Errorf("Expected empty cache to have 0 items, got %d", cache.ItemCount())
	}
}

func BenchmarkCacheSave(b *testing.B) {
	tempDir := b.TempDir()
	cacheFile := filepath.Join(tempDir, "bench_cache.json")

	cache, err := New(cacheFile)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}

	// 添加大量测试数据
	testIP := net.ParseIP("192.168.1.1")
	for i := 0; i < 1000; i++ {
		domain := fmt.Sprintf("test%d.com", i)
		cache.SetIPs("A", domain, []net.IP{testIP}, 1*time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := cache.Save()
		if err != nil {
			b.Fatalf("Failed to save cache: %v", err)
		}
	}
}