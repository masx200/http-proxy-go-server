package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/masx200/http-proxy-go-server/auth"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/tls"
	tls_auth "github.com/masx200/http-proxy-go-server/tls+auth"
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
	HTTP_PROXY  string   `json:"http_proxy"`
	HTTPS_PROXY string   `json:"https_proxy"`
	BypassList  []string `json:"bypass_list"`
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
		Pattern  string `json:"pattern"`
		Upstream string `json:"upstream"`
	} `json:"rules"`
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

// SelectProxyURLWithCIDR 根据输入的域名或IP地址选择代理服务器的URL，支持CIDR匹配
func SelectProxyURLWithCIDR(upstreams map[string]UpStream, rules []struct {
	Pattern  string `json:"pattern"`
	Upstream string `json:"upstream"`
}, domain string, scheme string) (string, error) {
	// 首先尝试解析为IP地址
	ip := net.ParseIP(domain)
	if ip != nil {
		// 如果是IP地址，检查CIDR匹配
		for _, rule := range rules {

			if rule.Pattern == "*" {
				var upstream = upstreams[rule.Upstream]
				if upstream.HTTPS_PROXY != "" && scheme == "https" {
					return upstream.HTTPS_PROXY, nil
				}
				if upstream.HTTP_PROXY != "" && scheme == "http" {
					return upstream.HTTP_PROXY, nil
				}
			}
			// 检查是否是CIDR格式
			if strings.Contains(rule.Pattern, "/") {
				_, ipNet, err := net.ParseCIDR(rule.Pattern)
				if err == nil && ipNet.Contains(ip) {
					// 找到匹配的CIDR，返回对应的代理
					if upstream, exists := upstreams[rule.Upstream]; exists {
						if upstream.HTTPS_PROXY != "" && scheme == "https" {
							return upstream.HTTPS_PROXY, nil
						}
						if upstream.HTTP_PROXY != "" && scheme == "http" {
							return upstream.HTTP_PROXY, nil
						}
					}
				}
			} else if rule.Pattern == domain || strings.HasPrefix(domain, rule.Pattern) {
				// 精确IP匹配或前缀匹配
				if upstream, exists := upstreams[rule.Upstream]; exists {
					if upstream.HTTPS_PROXY != "" && scheme == "https" {
						return upstream.HTTPS_PROXY, nil
					}
					if upstream.HTTP_PROXY != "" && scheme == "http" {
						return upstream.HTTP_PROXY, nil
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
			if rule.Pattern == "*" {
				var upstream = upstreams[rule.Upstream]
				if scheme == "https" && upstream.HTTPS_PROXY != "" {
					return upstream.HTTPS_PROXY, nil
				}
				if scheme == "http" && upstream.HTTP_PROXY != "" {
					return upstream.HTTP_PROXY, nil
				}
			}
			// 检查是否是CIDR格式（域名不应该匹配CIDR）
			if !strings.Contains(rule.Pattern, "/") {
				if matchWildcard(rule.Pattern, domain) || strings.Contains(domain, rule.Pattern) {
					// 找到匹配的域名模式，返回对应的代理
					if upstream, exists := upstreams[rule.Upstream]; exists {
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

	// 如果没有匹配的规则，返回空字符串和错误
	return "", nil
}

// IsBypassedWithCIDR 检查目标是否在bypass列表中，支持CIDR匹配
func IsBypassedWithCIDR(upstreams map[string]UpStream, rules []struct {
	Pattern  string `json:"pattern"`
	Upstream string `json:"upstream"`
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

// ProxySelector 使用SelectProxyURLWithCIDR和IsBypassedWithCIDR实现代理选择逻辑
func ProxySelector(r *http.Request, UpStreams map[string]UpStream, Rules []struct {
	Pattern  string `json:"pattern"`
	Upstream string `json:"upstream"`
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
	if IsBypassedWithCIDR(UpStreams, Rules, host) {
		return nil, nil
	}

	// 选择代理URL
	proxyURL, err := SelectProxyURLWithCIDR(UpStreams, Rules, host, scheme)
	if err != nil {
		return nil, err
	}

	// 解析代理URL
	if proxyURL != "" {
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
	var proxyoptions = options.ProxyOptions{}

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
					proxyURL, err := ProxySelector(r, config.UpStreams, config.Rules)
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
