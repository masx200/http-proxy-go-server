package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/masx200/http-proxy-go-server/auth"
	"github.com/masx200/http-proxy-go-server/config"
	"github.com/masx200/http-proxy-go-server/dnscache"
	_ "github.com/masx200/http-proxy-go-server/doh"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/tls"
	tls_auth "github.com/masx200/http-proxy-go-server/tls+auth"
	"github.com/masx200/http-proxy-go-server/utils"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/socks5"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

// Type aliases for backward compatibility and easier migration
type Config = config.Config
type UpStream = config.UpStream
type RoutingRule = config.RoutingRule
type Filter = config.Filter
type DohConfig = config.DohConfig
type DNSCacheConfig = config.DNSCacheConfig

// 全局DNS缓存实例
var (
	globalDNSCache *dnscache.DNSCache
	cacheOnce      sync.Once
)

// CacheConfig DNS缓存配置
type CacheConfig struct {
	Enabled      bool          `json:"enabled"`
	FilePath     string        `json:"file_path"`
	AOFPath      string        `json:"aof_path"`
	DefaultTTL   time.Duration `json:"default_ttl"`
	SaveInterval time.Duration `json:"save_interval"`
	AOFInterval  time.Duration `json:"aof_interval"`
	AOFEnabled   bool          `json:"aof_enabled"`
}

// DefaultCacheConfig 返回默认缓存配置
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

// InitDNSCache 初始化DNS缓存
func InitDNSCache(config *CacheConfig) error {
	var err error
	cacheOnce.Do(func() {
		if !config.Enabled {
			globalDNSCache = &dnscache.DNSCache{}
			return
		}

		dnscacheConfig := &dnscache.Config{
			FilePath:     config.FilePath,
			AOFPath:      config.AOFPath,
			DefaultTTL:   config.DefaultTTL,
			SaveInterval: config.SaveInterval,
			AOFInterval:  config.AOFInterval,
			Enabled:      config.Enabled,
		}

		globalDNSCache, err = dnscache.NewWithConfig(dnscacheConfig)
		if err != nil {
			log.Printf("初始化DNS缓存失败: %v", err)
		} else {
			log.Printf("DNS缓存初始化成功，AOF: %v", config.AOFEnabled)
		}
	})
	return err
}

// GetDNSCache 获取全局DNS缓存实例
func GetDNSCache() *dnscache.DNSCache {
	return globalDNSCache
}

// CloseDNSCache 关闭DNS缓存
func CloseDNSCache() {
	if globalDNSCache != nil {
		globalDNSCache.Close()
	}
}

type multiString []string

func (m *multiString) String() string {
	return "[" + strings.Join(*m, ", ") + "]"
}

func (m *multiString) Set(value string) error {
	*m = append(*m, value)
	return nil
}

// loadConfig 从配置文件加载并验证配置
func loadConfig(configFile string) (*config.Config, error) {
	return config.LoadAndValidateConfig(configFile)
}

// isValidDomain 检查字符串是否为有效的域名格式
func isValidDomain(domain string) bool {
	// 检查是否包含协议前缀
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return false
	}
	// 检查是否包含非法字符
	if strings.Contains(domain, "/") || strings.Contains(domain, ":") {
		return false
	}
	// 检查是否为空
	if domain == "" {
		return false
	}
	// 检查是否为IP地址格式
	if net.ParseIP(domain) != nil {
		return false
	}
	return true
}

// matchWildcard 检查域名是否匹配通配符模式
func matchWildcard(pattern, domain string) bool {
	if pattern == "*" {
		return true
	}
	if strings.HasPrefix(pattern, "*.") {
		suffix := pattern[2:]
		return strings.HasSuffix(domain, suffix) || domain == suffix
	}
	return pattern == domain
}

// overrideProxyURLCredentials 覆盖代理URL中的用户名密码
func overrideProxyURLCredentials(proxyURL string, username, password string) (string, error) {
	if proxyURL == "" {
		return proxyURL, nil
	}

	// 解析URL
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return proxyURL, err
	}

	// 如果配置中提供了用户名，则覆盖URL中的用户名
	if username != "" {
		// 如果配置中提供了密码，则使用配置中的密码
		if password != "" {
			parsedURL.User = url.UserPassword(username, password)
		} else {
			// 如果配置中没有提供密码，但URL中有密码，则保留URL中的密码
			if parsedURL.User != nil {
				if _, hasPassword := parsedURL.User.Password(); hasPassword {
					parsedURL.User = url.UserPassword(username, "")
					if existingPassword, ok := parsedURL.User.Password(); ok {
						parsedURL.User = url.UserPassword(username, existingPassword)
					}
				} else {
					parsedURL.User = url.User(username)
				}
			} else {
				parsedURL.User = url.User(username)
			}
		}
	} else if password != "" {
		// 如果只提供了密码但没有提供用户名，则保留URL中的用户名，只覆盖密码
		if parsedURL.User != nil {
			existingUsername := parsedURL.User.Username()
			parsedURL.User = url.UserPassword(existingUsername, password)
		}
	}

	return parsedURL.String(), nil
}

// SelectProxyURLWithCIDR 根据输入的域名或IP地址选择代理服务器的URL，支持CIDR匹配和WebSocket代理
func SelectProxyURLWithCIDR(upstreams map[string]config.UpStream, rules []config.RoutingRule, filters map[string]config.Filter, domain string, scheme string) (string, error) {
	// 首先尝试解析为IP地址
	ip := net.ParseIP(domain)
	if ip != nil {
		// 如果是IP地址，检查CIDR匹配
		for _, rule := range rules {
			// 获取filter对应的patterns
			filter, exists := filters[rule.Filter]
			if !exists {
				continue
			}

			// 检查filter中的所有patterns
			for _, pattern := range filter.Patterns {
				if pattern == "*" {
					var upstream = upstreams[rule.Upstream]
					// 优先检查WebSocket代理
					if upstream.TYPE == "websocket" && upstream.WS_PROXY != "" {
						// 检查是否已经包含协议前缀
						var wsProxyURL string
						if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
							wsProxyURL = upstream.WS_PROXY
						} else {
							wsProxyURL = "ws://" + upstream.WS_PROXY
						}
						// 使用配置中的ws_username和ws_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(wsProxyURL, upstream.WS_USERNAME, upstream.WS_PASSWORD)
					}
					// 检查SOCKS5代理
					if upstream.TYPE == "socks5" && upstream.SOCKS5_PROXY != "" {
						// 使用配置中的socks5_username和socks5_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.SOCKS5_PROXY, upstream.SOCKS5_USERNAME, upstream.SOCKS5_PASSWORD)
					}
					if upstream.HTTPS_PROXY != "" && scheme == "https" {
						// 使用配置中的http_username和http_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.HTTPS_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
					}
					if upstream.HTTP_PROXY != "" && scheme == "http" {
						// 使用配置中的http_username和http_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.HTTP_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
					}
				}
				// 检查是否是CIDR格式
				if strings.Contains(pattern, "/") {
					_, ipNet, err := net.ParseCIDR(pattern)
					if err == nil && ipNet.Contains(ip) {
						// 找到匹配的CIDR，返回对应的代理
						if upstream, exists := upstreams[rule.Upstream]; exists {
							// 优先检查WebSocket代理
							if upstream.WS_PROXY != "" && upstream.TYPE == "websocket" {
								// 检查是否已经包含协议前缀
								var wsProxyURL string
								if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
									wsProxyURL = upstream.WS_PROXY
								} else {
									wsProxyURL = "ws://" + upstream.WS_PROXY
								}
								// 使用配置中的ws_username和ws_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(wsProxyURL, upstream.WS_USERNAME, upstream.WS_PASSWORD)
							}
							// 检查SOCKS5代理
							if upstream.SOCKS5_PROXY != "" && upstream.TYPE == "socks5" {
								// 使用配置中的socks5_username和socks5_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.SOCKS5_PROXY, upstream.SOCKS5_USERNAME, upstream.SOCKS5_PASSWORD)
							}
							if upstream.HTTPS_PROXY != "" && scheme == "https" {
								// 使用配置中的http_username和http_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.HTTPS_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
							}
							if upstream.HTTP_PROXY != "" && scheme == "http" {
								// 使用配置中的http_username和http_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.HTTP_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
							}
						}
					}
				} else if pattern == domain || strings.HasPrefix(domain, pattern) {
					// 精确IP匹配或前缀匹配
					if upstream, exists := upstreams[rule.Upstream]; exists {
						// 优先检查WebSocket代理
						if upstream.TYPE == "websocket" && upstream.WS_PROXY != "" {
							// 检查是否已经包含协议前缀
							if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
								return upstream.WS_PROXY, nil
							}
							return "ws://" + upstream.WS_PROXY, nil
						}
						// 检查SOCKS5代理
						if upstream.TYPE == "socks5" && upstream.SOCKS5_PROXY != "" {
							return upstream.SOCKS5_PROXY, nil
						}
						if upstream.HTTPS_PROXY != "" && scheme == "https" {
							return upstream.HTTPS_PROXY, nil
						}
						if upstream.HTTP_PROXY != "" && scheme == "http" {
							return upstream.HTTP_PROXY, nil
						}
					}
				}
			}
		}
	} else {

		// 检查是否为有效的域名格式
		if !isValidDomain(domain) {
			return "", fmt.Errorf("invalid domain format: %s", domain)
		}

		// 如果是域名，进行域名匹配
		for _, rule := range rules {
			// 获取filter对应的patterns
			filter, exists := filters[rule.Filter]
			if !exists {
				continue
			}

			// 检查filter中的所有patterns
			for _, pattern := range filter.Patterns {
				if pattern == "*" {
					var upstream = upstreams[rule.Upstream]
					// 优先检查WebSocket代理
					if upstream.WS_PROXY != "" && upstream.TYPE == "websocket" {
						// 检查是否已经包含协议前缀
						var wsProxyURL string
						if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
							wsProxyURL = upstream.WS_PROXY
						} else {
							wsProxyURL = "ws://" + upstream.WS_PROXY
						}
						// 使用配置中的ws_username和ws_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(wsProxyURL, upstream.WS_USERNAME, upstream.WS_PASSWORD)
					}
					// 检查SOCKS5代理
					if upstream.SOCKS5_PROXY != "" && upstream.TYPE == "socks5" {
						// 使用配置中的socks5_username和socks5_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.SOCKS5_PROXY, upstream.SOCKS5_USERNAME, upstream.SOCKS5_PASSWORD)
					}
					if scheme == "https" && upstream.HTTPS_PROXY != "" {
						// 使用配置中的http_username和http_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.HTTPS_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
					}
					if scheme == "http" && upstream.HTTP_PROXY != "" {
						// 使用配置中的http_username和http_password覆盖URL中的用户名密码
						return overrideProxyURLCredentials(upstream.HTTP_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
					}
				}
				// 检查是否是CIDR格式（域名不应该匹配CIDR）
				if !strings.Contains(pattern, "/") {
					if matchWildcard(pattern, domain) || strings.Contains(domain, pattern) {
						// 找到匹配的域名模式，返回对应的代理
						if upstream, exists := upstreams[rule.Upstream]; exists {
							// 优先检查WebSocket代理
							if upstream.WS_PROXY != "" && upstream.TYPE == "websocket" {
								// 检查是否已经包含协议前缀
								var wsProxyURL string
								if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
									wsProxyURL = upstream.WS_PROXY
								} else {
									wsProxyURL = "ws://" + upstream.WS_PROXY
								}
								// 使用配置中的ws_username和ws_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(wsProxyURL, upstream.WS_USERNAME, upstream.WS_PASSWORD)
							}
							// 检查SOCKS5代理
							if upstream.SOCKS5_PROXY != "" && upstream.TYPE == "socks5" {
								// 使用配置中的socks5_username和socks5_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.SOCKS5_PROXY, upstream.SOCKS5_USERNAME, upstream.SOCKS5_PASSWORD)
							}
							if scheme == "http" && upstream.HTTP_PROXY != "" {
								// 使用配置中的http_username和http_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.HTTP_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
							}
							if scheme == "https" && upstream.HTTPS_PROXY != "" {
								// 使用配置中的http_username和http_password覆盖URL中的用户名密码
								return overrideProxyURLCredentials(upstream.HTTPS_PROXY, upstream.HTTP_USERNAME, upstream.HTTP_PASSWORD)
							}
						}
					}
				}
			}
		}
	}

	// 如果没有匹配的规则，返回空字符串和错误
	return "", nil
}

// IsBypassedWithCIDR 检查目标是否在bypass列表中，支持CIDR匹配
func IsBypassedWithCIDR(upstreams map[string]config.UpStream, rules []config.RoutingRule, filters map[string]config.Filter, target string) bool {
	// 首先尝试解析为IP地址
	ip := net.ParseIP(target)
	if ip != nil {
		// 如果是IP地址，检查CIDR匹配
		for _, rule := range rules {
			if upstream, exists := upstreams[rule.Upstream]; exists {
				// 检查bypass列表中的CIDR
				for _, bypass := range upstream.BypassList {
					if strings.Contains(bypass, "/") {
						_, ipNet, err := net.ParseCIDR(bypass)
						if err == nil && ipNet.Contains(ip) {
							return true
						}
					} else if bypass == target || strings.HasPrefix(target, bypass) {
						return true
					}
				}
			}
		}
	} else {
		// 如果是域名，进行域名匹配
		for _, rule := range rules {
			if upstream, exists := upstreams[rule.Upstream]; exists {
				for _, bypass := range upstream.BypassList {
					if !strings.Contains(bypass, "/") {
						if strings.Contains(target, bypass) ||
							strings.HasPrefix(target, bypass) ||
							strings.HasSuffix(target, bypass) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// ProxySelector 使用SelectProxyURLWithCIDR和IsBypassedWithCIDR实现代理选择逻辑，支持WebSocket代理
func ProxySelector(r *http.Request, UpStreams map[string]config.UpStream, Rules []config.RoutingRule, Filters map[string]config.Filter) (*url.URL, error) {
	scheme := r.URL.Scheme
	// 提取请求的主机名
	host := r.URL.Host
	if host == "" {
		host = r.Host
	}

	// 移除端口号
	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}

	// 检查是否应该被绕过
	if IsBypassedWithCIDR(UpStreams, Rules, Filters, host) {
		return nil, nil
	}

	// 选择代理URL
	proxyURL, err := SelectProxyURLWithCIDR(UpStreams, Rules, Filters, host, scheme)
	if err != nil {
		return nil, err
	}

	// 解析代理URL
	if proxyURL != "" {
		// 检查是否是WebSocket代理
		if strings.HasPrefix(proxyURL, "ws://") || strings.HasPrefix(proxyURL, "wss://") {
			// 对于WebSocket代理，返回一个特殊的URL，后续在transport配置中处理
			return url.Parse(proxyURL)
		}
		return url.Parse(proxyURL)
	}

	return nil, nil
}

func main() {
	// 设置日志输出到标准错误，确保日志能够被正确捕获
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 添加配置文件参数
	configFile := flag.String("config", "", "JSON配置文件路径")

	// 自定义字符串切片类型，实现 flag.Value 接口
	var (
		dohurls  multiString
		dohips   multiString
		dohalpns multiString
	)
	// 注册可重复参数
	flag.Var(&dohurls, "dohurl", "DOH URL (可重复),支持http协议和https协议")
	flag.Var(&dohips, "dohip", "DOH IP (可重复),支持ipv4地址和ipv6地址")
	flag.Var(&dohalpns, "dohalpn", "DOH alpn (可重复),支持h2协议和h3协议")

	var (
		hostname    = flag.String("hostname", "0.0.0.0", "an String value for hostname")
		port        = flag.Int("port", 8080, "TCP port to listen on")
		server_cert = flag.String("server_cert", "", "tls server cert")
		server_key  = flag.String("server_key", "", "tls server key")
		username    = flag.String("username", "", "username")
		password    = flag.String("password", "", "password")
		// 新增WebSocket代理相关参数
		upstreamType     = flag.String("upstream-type", "", "upstream proxy type (websocket, socks5, http)")
		upstreamAddress  = flag.String("upstream-address", "", "upstream proxy address (e.g., ws://127.0.0.1:1081, socks5://127.0.0.1:1080 or http://127.0.0.1:8080)")
		upstreamUsername = flag.String("upstream-username", "", "upstream proxy username")
		upstreamPassword = flag.String("upstream-password", "", "upstream proxy password")
		// DNS缓存相关参数
		cacheEnabled      = flag.Bool("cache-enabled", true, "enable DNS caching")
		cacheFile         = flag.String("cache-file", "./dns_cache.json", "DNS cache file path")
		cacheTTL          = flag.String("cache-ttl", "10m", "DNS cache TTL (duration string, e.g., 5m, 10m, 1h)")
		cacheSaveInterval = flag.String("cache-save-interval", "30s", "DNS cache full save interval (duration string, e.g., 30s, 1m)")
		// DNS缓存AOF相关参数
		cacheAOFEnabled  = flag.Bool("cache-aof-enabled", true, "enable DNS cache AOF (append-only file) persistence")
		cacheAOFFile     = flag.String("cache-aof-file", "./dns_cache.aof", "DNS cache AOF file path")
		cacheAOFInterval = flag.String("cache-aof-interval", "1s", "DNS cache AOF save interval (duration string, e.g., 1s, 5s)")
	)
	flag.Parse()

	log.Println("代理服务器启动中...")

	// 如果指定了配置文件，则从配置文件读取参数
	var config *Config
	var err error
	if *configFile != "" {
		config, err = loadConfig(*configFile)
		if err != nil {
			log.Printf("读取配置文件失败: %v\n", err)
			os.Exit(1)
		}
		// 使用配置文件的值覆盖命令行参数的默认值
		if config.Hostname != "" {
			*hostname = config.Hostname
		}
		if config.Port != 0 {
			*port = config.Port
		}
		if config.ServerCert != "" {
			*server_cert = config.ServerCert
		}
		if config.ServerKey != "" {
			*server_key = config.ServerKey
		}
		if config.Username != "" {
			*username = config.Username
		}
		if config.Password != "" {
			*password = config.Password
		}
		if len(config.Doh) > 0 {
			for _, dohConfig := range config.Doh {
				if dohConfig.URL != "" {
					dohurls = append(dohurls, dohConfig.URL)
				}
				if dohConfig.IP != "" {
					dohips = append(dohips, dohConfig.IP)
				}
				if dohConfig.Alpn != "" {
					dohalpns = append(dohalpns, dohConfig.Alpn)
				}
			}
		}
	}
	log.Println("dohalpn:", dohalpns.String())
	//parse cmd flags
	log.Println(
		"hostname:", *hostname)
	log.Println(
		"port:", *port)
	log.Println(
		"server_cert:", *server_cert)
	log.Println(
		"server_key:", *server_key)
	log.Println(
		"username:", *username)
	log.Println(
		"password:", *password)
	log.Println(
		"dohurl:", dohurls.String())
	log.Println("dohip:", dohips.String())
	log.Println("upstream-type:", *upstreamType)
	log.Println("upstream-address:", *upstreamAddress)
	log.Println("upstream-username:", *upstreamUsername)
	log.Println("upstream-password:", *upstreamPassword)
	log.Println("cache-enabled:", *cacheEnabled)
	log.Println("cache-file:", *cacheFile)
	log.Println("cache-ttl:", *cacheTTL)
	log.Println("cache-save-interval:", *cacheSaveInterval)
	log.Println("cache-aof-enabled:", *cacheAOFEnabled)
	log.Println("cache-aof-file:", *cacheAOFFile)
	log.Println("cache-aof-interval:", *cacheAOFInterval)

	// 如果指定了配置文件，则从配置文件读取DNS缓存配置
	if config != nil {
		if config.DNSCache.Enabled {
			*cacheEnabled = config.DNSCache.Enabled
		}
		if config.DNSCache.File != "" {
			*cacheFile = config.DNSCache.File
		}
		if config.DNSCache.TTL != "" {
			*cacheTTL = config.DNSCache.TTL
		}
		if config.DNSCache.SaveInterval != "" {
			*cacheSaveInterval = config.DNSCache.SaveInterval
		}
		if config.DNSCache.AOFEnabled {
			*cacheAOFEnabled = config.DNSCache.AOFEnabled
		}
		if config.DNSCache.AOFFile != "" {
			*cacheAOFFile = config.DNSCache.AOFFile
		}
		if config.DNSCache.AOFInterval != "" {
			*cacheAOFInterval = config.DNSCache.AOFInterval
		}
	}

	// 解析DNS缓存配置并初始化
	if *cacheEnabled {
		// 解析TTL
		cacheTTLDuration, err := time.ParseDuration(*cacheTTL)
		if err != nil {
			log.Printf("解析cache-ttl失败，使用默认值: %v", err)
			cacheTTLDuration = 10 * time.Minute
		}

		// 解析保存间隔
		cacheSaveIntervalDuration, err := time.ParseDuration(*cacheSaveInterval)
		if err != nil {
			log.Printf("解析cache-save-interval失败，使用默认值: %v", err)
			cacheSaveIntervalDuration = 30 * time.Second
		}

		// 解析AOF保存间隔
		cacheAOFIntervalDuration, err := time.ParseDuration(*cacheAOFInterval)
		if err != nil {
			log.Printf("解析cache-aof-interval失败，使用默认值: %v", err)
			cacheAOFIntervalDuration = 1 * time.Second
		}

		// 创建缓存配置并初始化DNS缓存系统
		dnsCacheConfig := &CacheConfig{
			Enabled:      *cacheEnabled,
			FilePath:     *cacheFile,
			AOFPath:      *cacheAOFFile,
			DefaultTTL:   cacheTTLDuration,
			SaveInterval: cacheSaveIntervalDuration,
			AOFInterval:  cacheAOFIntervalDuration,
			AOFEnabled:   *cacheAOFEnabled,
		}

		// 初始化DNS缓存
		err = InitDNSCache(dnsCacheConfig)
		if err != nil {
			log.Printf("初始化DNS缓存失败，将禁用缓存: %v", err)
		} else {
			log.Printf("DNS缓存已启用，文件: %s, AOF: %v, TTL: %v", *cacheFile, *cacheAOFEnabled, cacheTTLDuration)
		}
	} else {
		// 即使禁用缓存也要初始化
		dnsCacheConfig := &CacheConfig{
			Enabled: false,
		}
		InitDNSCache(dnsCacheConfig)
		log.Println("DNS缓存已禁用")
	}

	// 添加信号处理
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("收到退出信号，正在关闭服务器...")
		// 关闭DNS缓存
		CloseDNSCache()
		log.Println("DNS缓存已关闭")
		os.Exit(0)
	}()

	var proxyoptions = options.ProxyOptions{}

	// 处理WebSocket代理参数
	if *upstreamType == "websocket" && *upstreamAddress != "" {
		// 如果配置为空，则创建一个默认配置
		if config == nil {
			config = &Config{}
		}
		// 如果UpStreams为空，则初始化
		if config.UpStreams == nil {
			config.UpStreams = make(map[string]UpStream)
		}
		// 如果Rules为空，则初始化
		if config.Rules == nil {
			config.Rules = []RoutingRule{}
		}
		// 如果Filters为空，则初始化
		if config.Filters == nil {
			config.Filters = make(map[string]Filter)
		}

		// 创建WebSocket代理配置
		wsUpstream := UpStream{
			TYPE:        "websocket",
			HTTP_PROXY:  "",
			HTTPS_PROXY: "",
			BypassList:  []string{},
			WS_PROXY:    *upstreamAddress,
			WS_USERNAME: *upstreamUsername,
			WS_PASSWORD: *upstreamPassword,
		}

		// 添加到UpStreams
		config.UpStreams["websocket_upstream"] = wsUpstream

		// 添加规则和过滤器
		config.Rules = append(config.Rules, RoutingRule{
			Filter:   "websocket_filter",
			Upstream: "websocket_upstream",
		})

		config.Filters["websocket_filter"] = Filter{
			Patterns: []string{"*"},
		}

		log.Println("WebSocket代理配置已添加")
	}

	// 处理SOCKS5代理参数
	if *upstreamType == "socks5" && *upstreamAddress != "" {
		// 如果配置为空，则创建一个默认配置
		if config == nil {
			config = &Config{}
		}
		// 如果UpStreams为空，则初始化
		if config.UpStreams == nil {
			config.UpStreams = make(map[string]UpStream)
		}
		// 如果Rules为空，则初始化
		if config.Rules == nil {
			config.Rules = []RoutingRule{}
		}
		// 如果Filters为空，则初始化
		if config.Filters == nil {
			config.Filters = make(map[string]Filter)
		}

		// 创建SOCKS5代理配置
		socks5Upstream := UpStream{
			TYPE:            "socks5",
			HTTP_PROXY:      "",
			HTTPS_PROXY:     "",
			BypassList:      []string{},
			SOCKS5_PROXY:    *upstreamAddress,
			SOCKS5_USERNAME: *upstreamUsername,
			SOCKS5_PASSWORD: *upstreamPassword,
		}

		// 添加到UpStreams
		config.UpStreams["socks5_upstream"] = socks5Upstream

		// 添加规则和过滤器
		config.Rules = append(config.Rules, RoutingRule{
			Filter:   "socks5_filter",
			Upstream: "socks5_upstream",
		})

		config.Filters["socks5_filter"] = Filter{
			Patterns: []string{"*"},
		}

		log.Println("SOCKS5代理配置已添加")
	}

	// 处理HTTP代理参数
	if *upstreamType == "http" && *upstreamAddress != "" {
		// 如果配置为空，则创建一个默认配置
		if config == nil {
			config = &Config{}
		}
		// 如果UpStreams为空，则初始化
		if config.UpStreams == nil {
			config.UpStreams = make(map[string]UpStream)
		}
		// 如果Rules为空，则初始化
		if config.Rules == nil {
			config.Rules = []RoutingRule{}
		}
		// 如果Filters为空，则初始化
		if config.Filters == nil {
			config.Filters = make(map[string]Filter)
		}

		// 创建HTTP代理配置
		httpUpstream := UpStream{
			TYPE:        "http",
			HTTP_PROXY:  *upstreamAddress,
			HTTPS_PROXY: *upstreamAddress,
			BypassList:  []string{},
			WS_PROXY:    "",
			WS_USERNAME: "",
			WS_PASSWORD: "",
		}

		// 如果提供了用户名和密码，则添加到代理地址中
		if *upstreamUsername != "" || *upstreamPassword != "" {
			// 解析代理地址
			parsedURL, err := url.Parse(*upstreamAddress)
			if err == nil {
				// 设置用户名和密码
				parsedURL.User = url.UserPassword(*upstreamUsername, *upstreamPassword)
				// 重新设置代理地址
				httpUpstream.HTTP_PROXY = parsedURL.String()
				httpUpstream.HTTPS_PROXY = parsedURL.String()
			}
		}

		// 添加到UpStreams
		config.UpStreams["http_upstream"] = httpUpstream

		// 添加规则和过滤器
		config.Rules = append(config.Rules, RoutingRule{
			Filter:   "http_filter",
			Upstream: "http_upstream",
		})

		config.Filters["http_filter"] = Filter{
			Patterns: []string{"*"},
		}

		log.Println("HTTP代理配置已添加")
	}

	for i, dohurl := range dohurls {

		var dohip string
		if len(dohips) > i {
			dohip = dohips[i]
		} else {
			dohip = ""
		}
		var dohalpn string
		if len(dohalpns) > i {
			dohalpn = dohalpns[i]
		} else {
			dohalpn = ""
		}

		proxyoptions = append(proxyoptions, options.ProxyOption{Dohurl: dohurl, Dohip: dohip, Dohalpn: dohalpn})
	}

	var tranportConfigurations = []func(*http.Transport) *http.Transport{}
	if config != nil {
		if config.UpStreams != nil && config.Rules != nil && len(config.Rules) > 0 && len(config.UpStreams) > 0 {

			for name, upstream := range config.UpStreams {
				var proxyURL string
				if upstream.TYPE == "http" {
					proxyURL = upstream.HTTP_PROXY
				} else if upstream.TYPE == "socks5" {
					proxyURL = upstream.SOCKS5_PROXY
				} else if upstream.TYPE == "websocket" {
					proxyURL = upstream.WS_PROXY
				}

				var username, password string
				if upstream.TYPE == "http" {
					username = upstream.HTTP_USERNAME
					password = upstream.HTTP_PASSWORD
				} else if upstream.TYPE == "socks5" {
					username = upstream.SOCKS5_USERNAME
					password = upstream.SOCKS5_PASSWORD
				} else if upstream.TYPE == "websocket" {
					username = upstream.WS_USERNAME
					password = upstream.WS_PASSWORD
				}

				modifedUpstreamurl, err := overrideProxyURLCredentials(proxyURL, username, password)
				if err != nil {
					log.Fatalf("overrideProxyURLCredentials 出错: %v\n", err)
					return
				}
				modifedUpstream := UpStream{
					TYPE:         upstream.TYPE,
					HTTP_PROXY:   modifedUpstreamurl,
					HTTPS_PROXY:  modifedUpstreamurl,
					SOCKS5_PROXY: modifedUpstreamurl,
					WS_PROXY:     modifedUpstreamurl,
				}
				config.UpStreams[name] = modifedUpstream
			}
			tranportConfigurations = append(tranportConfigurations, func(t *http.Transport) *http.Transport {
				t.Proxy = func(r *http.Request) (*url.URL, error) {

					log.Println("ProxySelector", r.URL.Host)
					var addr = r.URL.Host

					var host, _, err = net.SplitHostPort(addr)
					if err != nil {

						if addrErr, ok := err.(*net.AddrError); ok && addrErr.Err == "missing port in address" {
							host = addr // 整个字符串就是 host
						} else {
							return nil, err
						}

					}
					if utils.IsLoopbackIP(host) {

						return nil, nil
					}
					proxyURL, err := ProxySelector(r, config.UpStreams, config.Rules, config.Filters)
					if err != nil {
						log.Printf("ProxySelector 出错: %v\n", err)
					} else {
						if proxyURL != nil {
							log.Printf("选择的代理 URL: %s\n", proxyURL.String())
						} else {
							log.Println("未选择代理")
						}
					}
					return proxyURL, err
				}
				t.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {

					var host, _, err = net.SplitHostPort(addr)
					if err != nil {
						return nil, err
					}
					if utils.IsLoopbackIP(host) {
						var dialer = &net.Dialer{}
						return dialer.DialContext(ctx, network, addr)
					}

					r, err := http.NewRequest("GET", "https://"+addr, nil)
					if err != nil {
						return nil, err
					}
					proxyURL, err := ProxySelector(r, config.UpStreams, config.Rules, config.Filters)
					if err != nil {
						return nil, err
					}

					if proxyURL != nil {
						log.Printf("选择的代理 URL: %s\n", proxyURL.String())
						if proxyURL.Scheme == "ws" || proxyURL.Scheme == "wss" {
							var modifiedUpstream = UpStream{
								TYPE:     "websocket",
								WS_PROXY: proxyURL.String(),
							}
							return websocketDialContext(ctx, network, addr, modifiedUpstream)
						}
						if proxyURL.Scheme == "socks5" || proxyURL.Scheme == "socks5s" {
							var modifiedUpstream = UpStream{
								TYPE:         "socks5",
								SOCKS5_PROXY: proxyURL.String(),
							}
							return socks5DialContext(ctx, network, addr, modifiedUpstream)
						} else {
							log.Println("未选择代理")
							var dialer = &net.Dialer{}
							return dialer.DialContext(ctx, network, addr)

						}
					}
					var dialer = &net.Dialer{}
					return dialer.DialContext(ctx, network, addr)
				}

				return t

			})
		}
	}
	by, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(by))
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password, proxyoptions, GetDNSCache(), tranportConfigurations...)
		return
	}
	// if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
	// 	tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password)
	// 	return
	// }
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		auth.Auth(*hostname, *port, *username, *password, proxyoptions, GetDNSCache(), tranportConfigurations...)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls.Tls(*server_cert, *server_key, *hostname, *port, proxyoptions, GetDNSCache(), tranportConfigurations...)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		simple.Simple(*hostname, *port, proxyoptions, GetDNSCache(), tranportConfigurations...)
		return
	}
}

// websocketDialContext 实现WebSocket代理连接
func websocketDialContext(ctx context.Context, network, addr string, upstream config.UpStream) (net.Conn, error) {
	// 解析目标地址
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		// 如果没有端口，尝试添加默认端口
		if network == "tcp" {
			if strings.Contains(addr, ":") {
				// IPv6地址
				addr = "[" + addr + "]:80"
			} else {
				// 域名或IPv4地址
				addr = addr + ":80"
			}
			host, port, err = net.SplitHostPort(addr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse address %s: %v", addr, err)
			}
		} else {
			return nil, fmt.Errorf("failed to parse address %s: %v", addr, err)
		}
	}

	// 转换端口号为整数
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to parse port %s: %v", port, err)
	}

	// 创建WebSocket客户端配置
	wsConfig := interfaces.ClientConfig{
		Username:   upstream.WS_USERNAME,
		Password:   upstream.WS_PASSWORD,
		ServerAddr: upstream.WS_PROXY,
		Protocol:   "websocket",
		Timeout:    30 * time.Second,
	}
	// 详细打印wsConfig的每个字段
	log.Println("WebSocket Config Details:")
	log.Println("host, portNum", host, portNum)
	log.Printf("  Username: %s", wsConfig.Username)
	log.Printf("  Password: %s", wsConfig.Password)
	log.Printf("  ServerAddr: %s", wsConfig.ServerAddr)
	log.Printf("  Protocol: %s", wsConfig.Protocol)
	log.Printf("  Timeout: %v", wsConfig.Timeout)
	// 创建WebSocket客户端
	websocketClient := socks5_websocket_proxy_golang_websocket.NewWebSocketClient(wsConfig)

	// 连接到目标主机
	err = websocketClient.Connect(host, portNum)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s:%d via WebSocket proxy: %v", host, portNum, err)
	}

	// 创建一个管道连接来处理WebSocket数据转发
	clientConn, serverConn := net.Pipe()

	// 在goroutine中处理WebSocket数据转发
	go func() {
		defer clientConn.Close()
		defer serverConn.Close()
		// 使用ForwardData方法处理WebSocket连接
		err := websocketClient.ForwardData(serverConn)
		if err != nil {
			log.Printf("WebSocket ForwardData error: %v\n", err)
		}
	}()

	// 返回客户端连接
	return clientConn, nil
}

// socks5DialContext 实现SOCKS5代理连接
func socks5DialContext(ctx context.Context, network, addr string, upstream config.UpStream) (net.Conn, error) {
	// 解析SOCKS5代理地址
	proxyURL, err := url.Parse(upstream.SOCKS5_PROXY)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SOCKS5 proxy URL %s: %v", upstream.SOCKS5_PROXY, err)
	}

	// 提取代理主机和端口
	proxyHost := proxyURL.Hostname()
	proxyPort := proxyURL.Port()
	if proxyPort == "" {
		proxyPort = "1080" // SOCKS5默认端口
	}

	// 详细打印SOCKS5配置信息
	log.Println("SOCKS5 Config Details:")
	log.Printf("  Target Address: %s", addr)
	log.Printf("  Proxy Host: %s", proxyHost)
	log.Printf("  Proxy Port: %s", proxyPort)
	log.Printf("  Username: %s", upstream.SOCKS5_USERNAME)
	log.Printf("  Password: %s", upstream.SOCKS5_PASSWORD)

	// 创建SOCKS5客户端配置
	// 确保ServerAddr必须以socks5://、tcp://、tls://或socks5s://开头
	serverAddr := proxyURL.String() //proxyHost + ":" + proxyPort
	if !strings.HasPrefix(serverAddr, "socks5://") && !strings.HasPrefix(serverAddr, "tcp://") && !strings.HasPrefix(serverAddr, "tls://") && !strings.HasPrefix(serverAddr, "socks5s://") {
		// 默认使用tcp://协议
		serverAddr = "tcp://" + serverAddr
	}

	socks5Config := interfaces.ClientConfig{
		Username:   upstream.SOCKS5_USERNAME,
		Password:   upstream.SOCKS5_PASSWORD,
		ServerAddr: serverAddr,
		Protocol:   "socks5",
		Timeout:    30 * time.Second,
	}

	// 创建SOCKS5客户端
	socks5Client := socks5.NewSOCKS5Client(socks5Config)

	// 使用DialContext连接到目标主机
	conn, err := socks5Client.DialContext(ctx, network, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s via SOCKS5 proxy: %v", addr, err)
	}

	// 返回连接
	return conn, nil
}
