package config

import "time"

// DohConfig DOH配置结构体
type DohConfig struct {
	IP       string `json:"ip"`
	Alpn     string `json:"alpn"`
	URL      string `json:"url"`
	Protocol string `json:"protocol"` // "doh", "doh3"
}

// DotConfig DoT配置结构体
type DotConfig struct {
	IP  string `json:"ip"`
	URL string `json:"url"`
}

// DoqConfig DoQ配置结构体
type DoqConfig struct {
	IP  string `json:"ip"`
	URL string `json:"url"`
}

// DNSCacheConfig DNS缓存配置
type DNSCacheConfig struct {
	Enabled       bool   `json:"enabled"`
	EnabledSet    bool   `json:"-"` // Internal flag to track if value was explicitly set
	File          string `json:"file"`
	TTL           string `json:"ttl"`
	SaveInterval  string `json:"save_interval"`
	AOFEnabled    bool   `json:"aof_enabled"`
	AOFEnabledSet bool   `json:"-"` // Internal flag to track if value was explicitly set
	AOFFile       string `json:"aof_file"`
	AOFInterval   string `json:"aof_interval"`
}

// UpStream 上游代理配置
type UpStream struct {
	TYPE        string   `json:"type"`
	HTTP_PROXY  string   `json:"http_proxy"`
	HTTPS_PROXY string   `json:"https_proxy"`
	BypassList  []string `json:"bypass_list"`
	// 新增WebSocket支持
	WS_PROXY    string `json:"ws_proxy"`    // WebSocket代理地址
	WS_USERNAME string `json:"ws_username"` // WebSocket代理用户名
	WS_PASSWORD string `json:"ws_password"` // WebSocket代理密码

	HTTP_USERNAME string `json:"http_username"` // http代理用户名
	HTTP_PASSWORD string `json:"http_password"` // http代理密码
	// 新增SOCKS5支持
	SOCKS5_PROXY    string `json:"socks5_proxy"`    // SOCKS5代理地址
	SOCKS5_USERNAME string `json:"socks5_username"` // SOCKS5代理用户名
	SOCKS5_PASSWORD string `json:"socks5_password"` // SOCKS5代理密码
}

// RoutingRule 路由规则
type RoutingRule struct {
	Filter   string `json:"filter"`
	Upstream string `json:"upstream"`
}

// Filter 过滤器配置
type Filter struct {
	Patterns []string `json:"patterns"`
}

// Config 主配置结构体
type Config struct {
	Hostname   string      `json:"hostname"`
	Port       int         `json:"port"`
	ServerCert string      `json:"server_cert"`
	ServerKey  string      `json:"server_key"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Doh        []DohConfig `json:"doh"`
	Dot        []DotConfig `json:"dot"`
	Doq        []DoqConfig `json:"doq"`

	// DNS缓存配置
	DNSCache DNSCacheConfig `json:"dns_cache"`

	UpStreams map[string]UpStream `json:"upstreams"`
	Rules     []RoutingRule       `json:"rules"`
	Filters   map[string]Filter   `json:"filters"`
}

// CacheConfig DNS缓存配置 (兼容现有代码)
type CacheConfig struct {
	Enabled      bool          `json:"enabled"`
	FilePath     string        `json:"file_path"`
	AOFPath      string        `json:"aof_path"`
	DefaultTTL   time.Duration `json:"default_ttl"`
	SaveInterval time.Duration `json:"save_interval"`
	AOFInterval  time.Duration `json:"aof_interval"`
	AOFEnabled   bool          `json:"aof_enabled"`
}

// DefaultCacheConfig 返回默认缓存配置 (兼容现有代码)
func DefaultCacheConfig() *CacheConfig {
	return &CacheConfig{
		Enabled:      true,
		FilePath:     "./dns_cache.json",
		AOFPath:      "./dns_cache.aof",
		DefaultTTL:   10 * time.Minute,
		SaveInterval: 30 * time.Second,
		AOFInterval:  1 * time.Second,
		AOFEnabled:   true,
	}
}

// ToCacheConfig 转换为兼容的CacheConfig结构
func (c *DNSCacheConfig) ToCacheConfig() (*CacheConfig, error) {
	config := &CacheConfig{
		Enabled:    c.Enabled,
		FilePath:   c.File,
		AOFPath:    c.AOFFile,
		AOFEnabled: c.AOFEnabled,
	}

	// Parse durations
	if ttl, err := time.ParseDuration(c.TTL); err == nil {
		config.DefaultTTL = ttl
	}
	if saveInterval, err := time.ParseDuration(c.SaveInterval); err == nil {
		config.SaveInterval = saveInterval
	}
	if aofInterval, err := time.ParseDuration(c.AOFInterval); err == nil {
		config.AOFInterval = aofInterval
	}

	return config, nil
}
