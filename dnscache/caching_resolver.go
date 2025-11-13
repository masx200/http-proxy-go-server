package dnscache

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/masx200/http-proxy-go-server/doh"
	"github.com/masx200/http-proxy-go-server/hosts"
	"github.com/masx200/http-proxy-go-server/options"
)

// NameResolver 接口定义
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
}

// CachingResolver 包装器，为DNS解析添加缓存功能
type CachingResolver struct {
	original NameResolver
	cache    *DNSCache
}

// NewCachingResolver 创建一个新的带缓存的解析器
func NewCachingResolver(original NameResolver, cache *DNSCache) *CachingResolver {
	return &CachingResolver{
		original: original,
		cache:    cache,
	}
}

// Resolve 使用缓存解析域名到IP
func (c *CachingResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("resolve:%s", name)
	if cached, found := c.cache.Get(cacheKey); found {
		log.Printf("DNS cache hit for resolve: %s", name)
		if ip, ok := cached.(net.IP); ok {
			return ctx, ip, nil
		}
	}

	// 缓存未命中，使用原始解析器
	resolvedCtx, ip, err := c.original.Resolve(ctx, name)
	if err != nil {
		return resolvedCtx, ip, err
	}

	// 存储到缓存，使用默认TTL
	c.cache.Set(cacheKey, ip)
	log.Printf("DNS cache set for resolve: %s -> %s", name, ip)

	return resolvedCtx, ip, nil
}

// LookupIP 使用缓存查找IP地址
func (c *CachingResolver) LookupIP(ctx context.Context, network, host string) ([]net.IP, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("lookupip:%s:%s", network, host)
	if cached, found := c.cache.Get(cacheKey); found {
		log.Printf("DNS cache hit for lookupip: %s (%s)", host, network)
		if ips, ok := cached.([]net.IP); ok {
			return ips, nil
		}
	}

	// 缓存未命中，使用原始解析器
	ips, err := c.original.LookupIP(ctx, network, host)
	if err != nil {
		return nil, err
	}

	// 存储到缓存，使用默认TTL
	c.cache.Set(cacheKey, ips)
	log.Printf("DNS cache set for lookupip: %s (%s) -> %v", host, network, ips)

	return ips, nil
}

// CreateHostsResolverCached 创建带缓存的Hosts解析器
func CreateHostsResolverCached(dnsCache *DNSCache) NameResolver {
	original := &HostsResolver{}
	return NewCachingResolver(original, dnsCache)
}

// CreateDOHResolverCached 创建带缓存的DOH解析器
func CreateDOHResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	original := &DOHResolver{
		proxyoptions:           proxyoptions,
		tranportConfigurations: tranportConfigurations,
	}
	return NewCachingResolver(original, dnsCache)
}

// CreateDOH3ResolverCached 创建带缓存的DOH3解析器
func CreateDOH3ResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache) NameResolver {
	original := &DOH3Resolver{
		proxyoptions: proxyoptions,
	}
	return NewCachingResolver(original, dnsCache)
}

// CreateHostsAndDohResolverCached 创建带缓存的Hosts+DOH解析器
func CreateHostsAndDohResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	original := &HostsAndDohResolver{
		proxyoptions:           proxyoptions,
		tranportConfigurations: tranportConfigurations,
	}
	return NewCachingResolver(original, dnsCache)
}

// 以下是从resolver包移过来的解析器实现

type HostsResolver struct{}

// LookupIP implements NameResolver.
func (h *HostsResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	// 使用 hosts 包解析域名
	ips, err := hosts.ResolveDomainToIPsWithHosts(host)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

// Resolve implements NameResolver.
func (h *HostsResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	// 使用 hosts 包解析域名
	ips, err := hosts.ResolveDomainToIPsWithHosts(name)
	if err != nil {
		return ctx, nil, err
	}
	if len(ips) == 0 {
		return ctx, nil, fmt.Errorf("no IP addresses found for domain %s", name)
	}
	// 返回第一个 IP 地址
	return ctx, ips[0], nil
}

type DOHResolver struct {
	proxyoptions           options.ProxyOptions
	tranportConfigurations []func(*http.Transport) *http.Transport
}

// LookupIP implements NameResolver.
func (d *DOHResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	if len(d.proxyoptions) == 0 {
		return nil, fmt.Errorf("no proxy options provided for DOH resolver")
	}

	// 随机打乱 proxyoptions 顺序
	Shuffle(d.proxyoptions)

	var allErrors []error
	for _, opt := range d.proxyoptions {
		// 跳过 h3 配置，因为这是 DOHResolver
		if opt.Dohalpn == "h3" {
			continue
		}

		ips, errors := doh.ResolveDomainToIPsWithDoh(host, opt.Dohurl, opt.Dohip, d.tranportConfigurations...)
		if len(ips) > 0 {
			return ips, nil
		}
		if len(errors) > 0 {
			allErrors = append(allErrors, errors...)
		}
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("DOH resolution failed for %s: %v", host, allErrors)
	}
	return nil, fmt.Errorf("no IP addresses found for domain %s", host)
}

// Resolve implements NameResolver.
func (d *DOHResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ips, err := d.LookupIP(ctx, "", name)
	if err != nil {
		return ctx, nil, err
	}
	if len(ips) == 0 {
		return ctx, nil, fmt.Errorf("no IP addresses found for domain %s", name)
	}
	// 返回第一个 IP 地址
	return ctx, ips[0], nil
}

type DOH3Resolver struct {
	proxyoptions options.ProxyOptions
}

// LookupIP implements NameResolver.
func (d *DOH3Resolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	if len(d.proxyoptions) == 0 {
		return nil, fmt.Errorf("no proxy options provided for DOH3 resolver")
	}

	// 随机打乱 proxyoptions 顺序
	Shuffle(d.proxyoptions)

	var allErrors []error
	for _, opt := range d.proxyoptions {
		// 只处理 h3 配置
		if opt.Dohalpn != "h3" {
			continue
		}

		var ips []net.IP
		var errors []error
		if opt.Dohip == "" {
			ips, errors = doh.ResolveDomainToIPsWithDoh3(host, opt.Dohurl)
		} else {
			ips, errors = doh.ResolveDomainToIPsWithDoh3(host, opt.Dohurl, opt.Dohip)
		}

		if len(ips) > 0 {
			return ips, nil
		}
		if len(errors) > 0 {
			allErrors = append(allErrors, errors...)
		}
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("DOH3 resolution failed for %s: %v", host, allErrors)
	}
	return nil, fmt.Errorf("no IP addresses found for domain %s", host)
}

// Resolve implements NameResolver.
func (d *DOH3Resolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ips, err := d.LookupIP(ctx, "", name)
	if err != nil {
		return ctx, nil, err
	}
	if len(ips) == 0 {
		return ctx, nil, fmt.Errorf("no IP addresses found for domain %s", name)
	}
	// 返回第一个 IP 地址
	return ctx, ips[0], nil
}

type HostsAndDohResolver struct {
	proxyoptions           options.ProxyOptions
	tranportConfigurations []func(*http.Transport) *http.Transport
}

// LookupIP implements NameResolver.
func (h *HostsAndDohResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	// 首先尝试使用 hosts 解析
	ips, err := hosts.ResolveDomainToIPsWithHosts(host)
	if err == nil && len(ips) > 0 {
		return ips, nil
	}

	// 如果 hosts 解析失败，尝试使用 DoH 解析
	if len(h.proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		Shuffle(h.proxyoptions)

		var allErrors []error
		for _, opt := range h.proxyoptions {
			var ips []net.IP
			var errors []error

			if opt.Dohalpn == "h3" {
				// 使用 DOH3
				if opt.Dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(host, opt.Dohurl)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(host, opt.Dohurl, opt.Dohip)
				}
			} else {
				// 使用 DOH
				ips, errors = doh.ResolveDomainToIPsWithDoh(host, opt.Dohurl, opt.Dohip, h.tranportConfigurations...)
			}

			if len(ips) > 0 {
				return ips, nil
			}
			if len(errors) > 0 {
				allErrors = append(allErrors, errors...)
			}
		}

		if len(allErrors) > 0 {
			return nil, fmt.Errorf("DOH resolution failed for %s: %v", host, allErrors)
		}
	}

	// 如果都失败了，返回原始的 hosts 错误
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("no IP addresses found for domain %s", host)
}

// Resolve implements NameResolver.
func (h *HostsAndDohResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ips, err := h.LookupIP(ctx, "", name)
	if err != nil {
		return ctx, nil, err
	}
	if len(ips) == 0 {
		return ctx, nil, fmt.Errorf("no IP addresses found for domain %s", name)
	}
	// 返回第一个 IP 地址
	return ctx, ips[0], nil
}

// ErrorArray 错误数组类型
type ErrorArray []error

// Error implements error.
func (e ErrorArray) Error() string {
	// 将 ErrorArray 中的每个 error 转换为字符串
	errStrings := make([]string, len(e))
	for i, err := range e {
		errStrings[i] = err.Error()
	}
	// 使用逗号分隔符连接所有错误字符串
	return "ErrorArray:[" + strings.Join(errStrings, ", ") + "]"
}

// Shuffle 对切片进行随机排序
func Shuffle[T any](slice []T) {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// IsIP 判断给定的字符串是否是有效的 IPv4 或 IPv6 地址。
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

// Proxy_net_DialCached 带DNS缓存的网络连接拨号函数
func Proxy_net_DialCached(network string, addr string, proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	if dnsCache != nil {
		return proxy_net_DialWithResolver(nil, network, addr, proxyoptions, CreateHostsAndDohResolverCached(proxyoptions, dnsCache, tranportConfigurations...), tranportConfigurations...)
	}
	return proxy_net_DialOriginal(network, addr, proxyoptions, tranportConfigurations...)
}

// Proxy_net_DialContextCached 带DNS缓存的上下文网络连接拨号函数
func Proxy_net_DialContextCached(ctx context.Context, network string, addr string, proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	if dnsCache != nil {
		return proxy_net_DialWithResolver(ctx, network, addr, proxyoptions, CreateHostsAndDohResolverCached(proxyoptions, dnsCache, tranportConfigurations...), tranportConfigurations...)
	}
	return proxy_net_DialContextOriginal(ctx, network, addr, proxyoptions, tranportConfigurations...)
}

// proxy_net_DialWithResolver 使用指定解析器的网络拨号函数
func proxy_net_DialWithResolver(ctx context.Context, network string, addr string, proxyoptions options.ProxyOptions, resolver NameResolver, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	if IsIP(hostname) {
		dialer := &net.Dialer{}
		if ctx != nil {
			return dialer.DialContext(ctx, network, addr)
		}
		return dialer.Dial(network, addr)
	}

	var ips []net.IP
	if resolver != nil {
		ips, err = resolver.LookupIP(ctx, network, hostname)
		if err != nil {
			log.Printf("Resolver failed for %s: %v", hostname, err)
		}
	}

	// 如果resolver解析失败，尝试使用本地hosts文件解析域名
	if len(ips) == 0 {
		ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)
		if err != nil {
			log.Printf("Hosts resolution failed for %s: %v", hostname, err)
		}
	}

	if len(ips) > 0 {
		Shuffle(ips)
		var errorsaray = make([]error, 0)
		for _, ip := range ips {
			serverIP := ip.String()
			newAddr := net.JoinHostPort(serverIP, port)
			dialer := &net.Dialer{}
			var connection net.Conn
			if ctx != nil {
				connection, err = dialer.DialContext(ctx, network, newAddr)
			} else {
				connection, err = dialer.Dial(network, newAddr)
			}

			if err != nil {
				errorsaray = append(errorsaray, err)
				continue
			} else {
				log.Printf("Successfully connected to %s via IP %s", addr, serverIP)
				return connection, nil
			}
		}
		return nil, ErrorArray(errorsaray)
	}

	// 如果提供了代理选项，尝试使用DOH解析
	if len(proxyoptions) > 0 {
		Shuffle(proxyoptions)
		var allErrors []error
		for _, opt := range proxyoptions {
			var ips []net.IP
			var errors []error

			if opt.Dohalpn == "h3" {
				if opt.Dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl, opt.Dohip)
				}
			} else {
				if opt.Dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, "", tranportConfigurations...)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
				}
			}

			if len(ips) == 0 && len(errors) > 0 {
				allErrors = append(allErrors, errors...)
				continue
			} else {
				lengthip := len(ips)
				Shuffle(ips)
				for i := 0; i < lengthip; i++ {
					var serverIP = ips[i].String()
					newAddr := net.JoinHostPort(serverIP, port)
					dialer := &net.Dialer{}
					var connection net.Conn
					var err1 error
					if ctx != nil {
						connection, err1 = dialer.DialContext(ctx, network, newAddr)
					} else {
						connection, err1 = dialer.Dial(network, newAddr)
					}

					if err1 != nil {
						allErrors = append(allErrors, err1)
						continue
					} else {
						log.Printf("Successfully connected to %s via DOH %s using IP %s", addr, opt.Dohurl, serverIP)
						return connection, nil
					}
				}
			}
		}
		return nil, ErrorArray(allErrors)
	}

	// 如果所有方法都失败了，使用原始地址
	dialer := &net.Dialer{}
	if ctx != nil {
		connection, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			log.Printf("Failed to connect to %s: %v", addr, err)
			return nil, err
		}
		log.Printf("Successfully connected to %s using original address", addr)
		return connection, nil
	}
	connection, err := dialer.Dial(network, addr)
	if err != nil {
		log.Printf("Failed to connect to %s: %v", addr, err)
		return nil, err
	}
	log.Printf("Successfully connected to %s using original address", addr)
	return connection, nil
}

// proxy_net_DialOriginal 原始的网络拨号函数（不使用缓存）
func proxy_net_DialOriginal(network string, addr string, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	if IsIP(hostname) {
		dialer := &net.Dialer{}
		return dialer.Dial(network, addr)
	}

	var ips []net.IP
	ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsaray = make([]error, 0)
		for i := 0; i < lengthip; i++ {
			var serverIP = ips[i].String()
			newAddr := net.JoinHostPort(serverIP, port)
			dialer := &net.Dialer{}
			connection, err1 := dialer.Dial(network, newAddr)

			if err1 != nil {
				errorsaray = append(errorsaray, err1)
				continue
			} else {
				log.Printf("success connect to addr=%s by network=%s by serverIP=%s", addr, network, serverIP)
				return connection, nil
			}
		}
		return nil, ErrorArray(errorsaray)
	}

	// hosts没有找到域名解析ip,可以忽略这个错误
	if len(ips) == 0 && err != nil {
		log.Println(err)
	}

	//调用ResolveDomainToIPsWithHosts函数解析域名
	if len(proxyoptions) > 0 {
		var errorsaray = make([]error, 0)
		Shuffle(proxyoptions)
		for _, dohurlopt := range proxyoptions {
			var dohip = dohurlopt.Dohip
			var dohalpn = dohurlopt.Dohalpn
			var ips []net.IP
			var errors []error
			hostname, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}

			if dohalpn == "h3" {
				if dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl, dohip)
				}
			} else {
				if dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, "", tranportConfigurations...)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, dohip, tranportConfigurations...)
				}
			}

			if len(ips) == 0 && len(errors) > 0 {
				errorsaray = append(errorsaray, errors...)
				continue
			} else {
				lengthip := len(ips)
				Shuffle(ips)
				for i := 0; i < lengthip; i++ {
					var serverIP = ips[i].String()
					newAddr := net.JoinHostPort(serverIP, port)
					dialer := &net.Dialer{}
					connection, err1 := dialer.Dial(network, newAddr)

					if err1 != nil {
						errorsaray = append(errorsaray, err1)
						continue
					} else {
						log.Printf("success connect to address=%s by network=%s by Dohurl=%s by dohip=%s by serverIP=%s", addr, network, dohurlopt.Dohurl, dohip, serverIP)
						return connection, nil
					}
				}
			}
		}
		return nil, ErrorArray(errorsaray)
	} else {
		dialer := &net.Dialer{}
		connection, err1 := dialer.Dial(network, addr)
		if err1 != nil {
			log.Printf("failure connect to %s by %s%s", addr, network, err1.Error())
			return nil, err1
		}
		log.Printf("success connect to %s by %s", addr, network)
		return connection, nil
	}
}

// proxy_net_DialContextOriginal 原始的上下文网络拨号函数（不使用缓存）
func proxy_net_DialContextOriginal(ctx context.Context, network string, address string, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if IsIP(hostname) {
		dialer := &net.Dialer{}
		return dialer.DialContext(ctx, network, address)
	}

	var ips []net.IP
	ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsaray = make([]error, 0)
		for i := 0; i < lengthip; i++ {
			var serverIP = ips[i].String()
			newAddr := net.JoinHostPort(serverIP, port)
			dialer := &net.Dialer{}
			connection, err1 := dialer.DialContext(ctx, network, newAddr)

			if err1 != nil {
				errorsaray = append(errorsaray, err1)
				continue
			} else {
				log.Printf("success connect to addr=%s by network=%s by serverIP=%s", address, network, serverIP)
				return connection, nil
			}
		}
		return nil, ErrorArray(errorsaray)
	}

	if len(ips) == 0 && err != nil {
		log.Println(err)
	}

	if len(proxyoptions) > 0 {
		var errorsaray = make([]error, 0)
		Shuffle(proxyoptions)
		for _, dohurlopt := range proxyoptions {
			var dohip = dohurlopt.Dohip
			var dohalpn = dohurlopt.Dohalpn
			var ips []net.IP
			var errors []error
			hostname, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}

			if dohalpn == "h3" {
				if dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl, dohip)
				}
			} else {
				if dohip == "" {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, "", tranportConfigurations...)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, dohip, tranportConfigurations...)
				}
			}

			if len(ips) == 0 && len(errors) > 0 {
				errorsaray = append(errorsaray, errors...)
				continue
			} else {
				lengthip := len(ips)
				Shuffle(ips)
				for i := 0; i < lengthip; i++ {
					var serverIP = ips[i].String()
					newAddr := net.JoinHostPort(serverIP, port)
					dialer := &net.Dialer{}
					connection, err1 := dialer.DialContext(ctx, network, newAddr)

					if err1 != nil {
						errorsaray = append(errorsaray, err1)
						continue
					} else {
						log.Printf("success connect to address=%s by network=%s by Dohurl=%s by dohip=%s by serverIP=%s", address, network, dohurlopt.Dohurl, dohip, serverIP)
						return connection, nil
					}
				}
			}
		}
		return nil, ErrorArray(errorsaray)
	} else {
		dialer := &net.Dialer{}
		connection, err1 := dialer.DialContext(ctx, network, address)
		if err1 != nil {
			log.Printf("failure connect to %s by %s%s", address, network, err1.Error())
			return nil, err1
		}
		log.Printf("success connect to %s by %s", address, network)
		return connection, nil
	}
}