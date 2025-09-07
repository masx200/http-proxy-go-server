package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/masx200/http-proxy-go-server/auth"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/tls"
	tls_auth "github.com/masx200/http-proxy-go-server/tls+auth"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

type multiString []string

func (m *multiString) String() string {
	return "[" + strings.Join(*m, ", ") + "]"
}

func (m *multiString) Set(value string) error {
	*m = append(*m, value)
	return nil
}

type UpStream struct {
	TYPE        string   `json:"type"`
	HTTP_PROXY  string   `json:"http_proxy"`
	HTTPS_PROXY string   `json:"https_proxy"`
	BypassList  []string `json:"bypass_list"`
	// 新增WebSocket支持
	WS_PROXY    string `json:"ws_proxy"`    // WebSocket代理地址
	WS_USERNAME string `json:"ws_username"` // WebSocket代理用户名
	WS_PASSWORD string `json:"ws_password"` // WebSocket代理密码
}

// func init() {
// 	http.ProxyFromEnvironment()
// }

// DohConfig DOH配置结构体
type DohConfig struct {
	IP   string `json:"ip"`
	Alpn string `json:"alpn"`
	URL  string `json:"url"`
}

// Config 结构体用于JSON配置文件
type Config struct {
	Hostname   string      `json:"hostname"`
	Port       int         `json:"port"`
	ServerCert string      `json:"server_cert"`
	ServerKey  string      `json:"server_key"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Doh        []DohConfig `json:"doh"`

	UpStreams map[string]UpStream `json:"upstreams"`
	Rules     []struct {
		Filter   string `json:"filter"`
		Upstream string `json:"upstream"`
	} `json:"rules"`
	Filters map[string]struct {
		Patterns []string `json:"patterns"`
	} `json:"filters"`
}

// loadConfig 从JSON文件加载配置
func loadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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

// SelectProxyURLWithCIDR 根据输入的域名或IP地址选择代理服务器的URL，支持CIDR匹配和WebSocket代理
func SelectProxyURLWithCIDR(upstreams map[string]UpStream, rules []struct {
	Filter   string `json:"filter"`
	Upstream string `json:"upstream"`
}, filters map[string]struct {
	Patterns []string `json:"patterns"`
}, domain string, scheme string) (string, error) {
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
						if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
							return upstream.WS_PROXY, nil
						}
						return "ws://" + upstream.WS_PROXY, nil
					}
					if upstream.HTTPS_PROXY != "" && scheme == "https" {
						return upstream.HTTPS_PROXY, nil
					}
					if upstream.HTTP_PROXY != "" && scheme == "http" {
						return upstream.HTTP_PROXY, nil
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
								if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
									return upstream.WS_PROXY, nil
								}
								return "ws://" + upstream.WS_PROXY, nil
							}
							if upstream.HTTPS_PROXY != "" && scheme == "https" {
								return upstream.HTTPS_PROXY, nil
							}
							if upstream.HTTP_PROXY != "" && scheme == "http" {
								return upstream.HTTP_PROXY, nil
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
						if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
							return upstream.WS_PROXY, nil
						}
						return "ws://" + upstream.WS_PROXY, nil
					}
					if scheme == "https" && upstream.HTTPS_PROXY != "" {
						return upstream.HTTPS_PROXY, nil
					}
					if scheme == "http" && upstream.HTTP_PROXY != "" {
						return upstream.HTTP_PROXY, nil
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
								if strings.HasPrefix(upstream.WS_PROXY, "ws://") || strings.HasPrefix(upstream.WS_PROXY, "wss://") {
									return upstream.WS_PROXY, nil
								}
								return "ws://" + upstream.WS_PROXY, nil
							}
							if scheme == "http" && upstream.HTTP_PROXY != "" {
								return upstream.HTTP_PROXY, nil
							}
							if scheme == "https" && upstream.HTTPS_PROXY != "" {
								return upstream.HTTPS_PROXY, nil
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
func IsBypassedWithCIDR(upstreams map[string]UpStream, rules []struct {
	Filter   string `json:"filter"`
	Upstream string `json:"upstream"`
}, filters map[string]struct {
	Patterns []string `json:"patterns"`
}, target string) bool {
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
func ProxySelector(r *http.Request, UpStreams map[string]UpStream, Rules []struct {
	Filter   string `json:"filter"`
	Upstream string `json:"upstream"`
}, Filters map[string]struct {
	Patterns []string `json:"patterns"`
}) (*url.URL, error) {
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
		upstreamType     = flag.String("upstream-type", "", "upstream proxy type (websocket)")
		upstreamAddress  = flag.String("upstream-address", "", "upstream proxy address (e.g., ws://127.0.0.1:1081)")
		upstreamUsername = flag.String("upstream-username", "", "upstream proxy username")
		upstreamPassword = flag.String("upstream-password", "", "upstream proxy password")
	)
	flag.Parse()

	// 如果指定了配置文件，则从配置文件读取参数
	var config *Config
	var err error
	if *configFile != "" {
		config, err = loadConfig(*configFile)
		if err != nil {
			fmt.Printf("读取配置文件失败: %v\n", err)
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
	fmt.Println("dohalpn:", dohalpns.String())
	//parse cmd flags
	fmt.Println(
		"hostname:", *hostname)
	fmt.Println(
		"port:", *port)
	fmt.Println(
		"server_cert:", *server_cert)
	fmt.Println(
		"server_key:", *server_key)
	fmt.Println(
		"username:", *username)
	fmt.Println(
		"password:", *password)
	fmt.Println(
		"dohurl:", dohurls.String())
	fmt.Println("dohip:", dohips.String())
	fmt.Println("upstream-type:", *upstreamType)
	fmt.Println("upstream-address:", *upstreamAddress)
	fmt.Println("upstream-username:", *upstreamUsername)
	fmt.Println("upstream-password:", *upstreamPassword)

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
			config.Rules = []struct {
				Filter   string `json:"filter"`
				Upstream string `json:"upstream"`
			}{}
		}
		// 如果Filters为空，则初始化
		if config.Filters == nil {
			config.Filters = make(map[string]struct {
				Patterns []string `json:"patterns"`
			})
		}

		// 创建WebSocket代理配置
		wsUpstream := UpStream{
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
		config.Rules = append(config.Rules, struct {
			Filter   string `json:"filter"`
			Upstream string `json:"upstream"`
		}{
			Filter:   "websocket_filter",
			Upstream: "websocket_upstream",
		})

		config.Filters["websocket_filter"] = struct {
			Patterns []string `json:"patterns"`
		}{
			Patterns: []string{"*"},
		}

		fmt.Println("WebSocket代理配置已添加")
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
			tranportConfigurations = append(tranportConfigurations, func(t *http.Transport) *http.Transport {
				t.Proxy = func(r *http.Request) (*url.URL, error) {

					fmt.Println("ProxySelector", r.URL.Host)
					proxyURL, err := ProxySelector(r, config.UpStreams, config.Rules, config.Filters)
					if err != nil {
						fmt.Printf("ProxySelector 出错: %v\n", err)
					} else {
						if proxyURL != nil {
							fmt.Printf("选择的代理 URL: %s\n", proxyURL.String())
						} else {
							fmt.Println("未选择代理")
						}
					}
					return proxyURL, err
				}

				// 检查是否有WebSocket代理配置
				for _, upstream := range config.UpStreams {
					if upstream.WS_PROXY != "" && upstream.TYPE == "websocket" {
						// 创建自定义的DialContext函数来处理WebSocket代理
						t.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
							// 实现WebSocket代理连接逻辑
							return websocketDialContext(ctx, network, addr, upstream)
						}
					}
				}
				return t
			})
		}
	}
	by, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(by))
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password, proxyoptions, tranportConfigurations...)
		return
	}
	// if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
	// 	tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password)
	// 	return
	// }
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		auth.Auth(*hostname, *port, *username, *password, proxyoptions, tranportConfigurations...)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls.Tls(*server_cert, *server_key, *hostname, *port, proxyoptions, tranportConfigurations...)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		simple.Simple(*hostname, *port, proxyoptions, tranportConfigurations...)
		return
	}
}

// websocketDialContext 实现WebSocket代理连接
func websocketDialContext(ctx context.Context, network, addr string, upstream UpStream) (net.Conn, error) {
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
			fmt.Printf("WebSocket ForwardData error: %v\n", err)
		}
	}()

	// 返回客户端连接
	return clientConn, nil
}

func init() {
	// var config interfaces.ClientConfig = interfaces.ClientConfig{}
	// socks5_websocket_proxy_golang_websocket.NewWebSocketClient(config)
}
