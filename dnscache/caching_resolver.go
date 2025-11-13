package dnscache

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	go_socks5 "gitee.com/masx200/go-socks5"
	"github.com/masx200/http-proxy-go-server/doh"
	"github.com/masx200/http-proxy-go-server/hosts"
	"github.com/masx200/http-proxy-go-server/options"
)

// NameResolver 接口定义
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
}

// CachingResolver 包装现有的NameResolver以添加缓存功能
type CachingResolver struct {
	underlying NameResolver
	cache      *DNSCache
	enabled    bool
}

// NewCachingResolver 创建带有缓存功能的解析器包装器
func NewCachingResolver(underlying NameResolver, cache *DNSCache) *CachingResolver {
	return &CachingResolver{
		underlying: underlying,
		cache:      cache,
		enabled:    cache != nil && cache.ItemCount() >= 0, // 检查cache是否有效
	}
}

// Resolve 实现NameResolver接口，返回单个IP地址
func (c *CachingResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	if !c.enabled {
		// 如果缓存未启用，直接使用底层解析器
		return c.underlying.Resolve(ctx, name)
	}

	// 尝试从缓存获取
	if ip, found := c.cache.GetIP("A", name); found {
		return ctx, ip, nil
	}

	// 缓存未命中，使用底层解析器
	resolvedCtx, ip, err := c.underlying.Resolve(ctx, name)
	if err != nil {
		return resolvedCtx, nil, err
	}

	// 将结果存入缓存
	if ip != nil {
		c.cache.SetIP("A", name, ip, 0) // 使用默认TTL
	}

	return resolvedCtx, ip, nil
}

// LookupIP 实现NameResolver接口，返回IP地址列表
func (c *CachingResolver) LookupIP(ctx context.Context, network, host string) ([]net.IP, error) {
	if !c.enabled {
		// 如果缓存未启用，直接使用底层解析器
		return c.underlying.LookupIP(ctx, network, host)
	}

	// 确定DNS记录类型
	dnsType := "A"
	if network == "tcp6" || network == "udp6" || network == "ip6" {
		dnsType = "AAAA"
	} else if network == "tcp4" || network == "udp4" || network == "ip4" {
		dnsType = "A"
	} else if network == "" {
		// 如果网络类型为空，同时尝试A和AAAA记录
		return c.lookupIPBoth(ctx, host)
	}

	// 尝试从缓存获取
	if ips, found := c.cache.GetIPs(dnsType, host); found {
		return ips, nil
	}

	// 缓存未命中，使用底层解析器
	ips, err := c.underlying.LookupIP(ctx, network, host)
	if err != nil {
		return nil, err
	}

	// 将结果存入缓存
	if len(ips) > 0 {
		c.cache.SetIPs(dnsType, host, ips, 0) // 使用默认TTL
	}

	return ips, nil
}

// lookupIPBoth 同时查找A和AAAA记录
func (c *CachingResolver) lookupIPBoth(ctx context.Context, host string) ([]net.IP, error) {
	var allIPs []net.IP

	// 尝试获取A记录
	if aIPs, found := c.cache.GetIPs("A", host); found {
		allIPs = append(allIPs, aIPs...)
	} else {
		// 缓存未命中，使用底层解析器
		if ips, err := c.underlying.LookupIP(ctx, "tcp4", host); err == nil && len(ips) > 0 {
			allIPs = append(allIPs, ips...)
			c.cache.SetIPs("A", host, ips, 0)
		}
	}

	// 尝试获取AAAA记录
	if aaaaIPs, found := c.cache.GetIPs("AAAA", host); found {
		allIPs = append(allIPs, aaaaIPs...)
	} else {
		// 缓存未命中，使用底层解析器
		if ips, err := c.underlying.LookupIP(ctx, "tcp6", host); err == nil && len(ips) > 0 {
			allIPs = append(allIPs, ips...)
			c.cache.SetIPs("AAAA", host, ips, 0)
		}
	}

	if len(allIPs) == 0 {
		return nil, fmt.Errorf("no IP addresses found for %s", host)
	}

	return allIPs, nil
}

// GetUnderlyingResolver 获取底层解析器（用于调试）
func (c *CachingResolver) GetUnderlyingResolver() NameResolver {
	return c.underlying
}

// GetCache 获取缓存实例（用于统计）
func (c *CachingResolver) GetCache() *DNSCache {
	return c.cache
}

// IsEnabled 检查缓存是否启用
func (c *CachingResolver) IsEnabled() bool {
	return c.enabled
}

// Invalidate 使特定域名的缓存失效
func (c *CachingResolver) Invalidate(domain string) {
	if !c.enabled {
		return
	}

	c.cache.Delete("A", domain)
	c.cache.Delete("AAAA", domain)
	c.cache.Delete("CNAME", domain)
	c.cache.Delete("MX", domain)
	c.cache.Delete("TXT", domain)
	c.cache.Delete("NS", domain)
}

// InvalidateAll 使所有缓存失效
func (c *CachingResolver) InvalidateAll() {
	if !c.enabled {
		return
	}

	c.cache.Flush()
}

// GetStats 获取缓存统计信息
func (c *CachingResolver) GetStats() map[string]interface{} {
	if !c.enabled {
		return map[string]interface{}{
			"enabled": false,
		}
	}

	stats := c.cache.Stats()
	stats["underlying_resolver_type"] = fmt.Sprintf("%T", c.underlying)
	return stats
}

// SetCustomTTL 为特定域名和类型设置自定义TTL
func (c *CachingResolver) SetCustomTTL(dnsType, domain string, ttl time.Duration) {
	if !c.enabled {
		return
	}

	if ips, found := c.cache.GetIPs(dnsType, domain); found {
		c.cache.SetIPs(dnsType, domain, ips, ttl)
	}
}

// ===== 缓存解析器创建函数 =====

// HostsResolver 仅使用hosts文件解析
type HostsResolver struct {
}

// LookupIP implements NameResolver
func (h *HostsResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	// 使用 hosts 包解析域名
	ips, err := hosts.ResolveDomainToIPsWithHosts(host)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

// Resolve implements NameResolver
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

// DOHResolver 使用DNS over HTTPS解析
type DOHResolver struct {
	proxyoptions           options.ProxyOptions
	tranportConfigurations []func(*http.Transport) *http.Transport
}

// LookupIP implements NameResolver
func (d *DOHResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	if len(d.proxyoptions) == 0 {
		return nil, fmt.Errorf("no proxy options provided for DOH resolver")
	}

	// 随机打乱 proxyoptions 顺序
	options.Shuffle(d.proxyoptions)

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

// Resolve implements NameResolver
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

// DOH3Resolver 使用DNS over HTTPS HTTP/3解析
type DOH3Resolver struct {
	proxyoptions options.ProxyOptions
}

// LookupIP implements NameResolver
func (d *DOH3Resolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	if len(d.proxyoptions) == 0 {
		return nil, fmt.Errorf("no proxy options provided for DOH3 resolver")
	}

	// 随机打乱 proxyoptions 顺序
	options.Shuffle(d.proxyoptions)

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

// Resolve implements NameResolver
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

// HostsAndDohResolver 组合hosts和DoH解析
type HostsAndDohResolver struct {
	proxyoptions           options.ProxyOptions
	tranportConfigurations []func(*http.Transport) *http.Transport
}

// LookupIP implements NameResolver
func (h *HostsAndDohResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	// 首先尝试使用 hosts 解析
	ips, err := hosts.ResolveDomainToIPsWithHosts(host)
	if err == nil && len(ips) > 0 {
		return ips, nil
	}

	// 如果 hosts 解析失败，尝试使用 DoH 解析
	if len(h.proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		options.Shuffle(h.proxyoptions)

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

// Resolve implements NameResolver
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

// ===== 带缓存的解析器创建函数 =====

// CreateHostsResolverCached 创建带缓存的hosts解析器
func CreateHostsResolverCached(dnsCache *DNSCache) NameResolver {
	if dnsCache == nil {
		return &HostsResolver{}
	}
	underlying := &HostsResolver{}
	return NewCachingResolver(underlying, dnsCache)
}

// CreateDOHResolverCached 创建带缓存的DOH解析器
func CreateDOHResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	if dnsCache == nil {
		return &DOHResolver{
			proxyoptions:           proxyoptions,
			tranportConfigurations: tranportConfigurations,
		}
	}
	underlying := &DOHResolver{
		proxyoptions:           proxyoptions,
		tranportConfigurations: tranportConfigurations,
	}
	return NewCachingResolver(underlying, dnsCache)
}

// CreateDOH3ResolverCached 创建带缓存的DOH3解析器
func CreateDOH3ResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache) NameResolver {
	if dnsCache == nil {
		return &DOH3Resolver{
			proxyoptions: proxyoptions,
		}
	}
	underlying := &DOH3Resolver{
		proxyoptions: proxyoptions,
	}
	return NewCachingResolver(underlying, dnsCache)
}

// CreateHostsAndDohResolverCached 创建带缓存的hosts和DoH组合解析器
func CreateHostsAndDohResolverCached(proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	if dnsCache == nil {
		return &HostsAndDohResolver{
			proxyoptions:           proxyoptions,
			tranportConfigurations: tranportConfigurations,
		}
	}
	underlying := &HostsAndDohResolver{
		proxyoptions:           proxyoptions,
		tranportConfigurations: tranportConfigurations,
	}
	return NewCachingResolver(underlying, dnsCache)
}

// ===== 缓存拨号函数 =====

// Proxy_net_DialCached 带DNS缓存的网络连接拨号函数
// 如果提供了dnsCache，则使用缓存的解析器；否则使用原始的直接解析方式
func Proxy_net_DialCached(network string, addr string, proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	if dnsCache != nil {
		return proxy_net_DialWithResolver(network, addr, proxyoptions, CreateHostsAndDohResolverCached(proxyoptions, dnsCache, tranportConfigurations...))
	}
	return proxy_net_DialOriginal(network, addr, proxyoptions, tranportConfigurations...)
}

// Proxy_net_DialContextCached 带DNS缓存的上下文网络连接拨号函数
// 如果提供了dnsCache，则使用缓存的解析器；否则使用原始的直接解析方式
func Proxy_net_DialContextCached(ctx context.Context, network string, address string, proxyoptions options.ProxyOptions, dnsCache *DNSCache, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	if dnsCache != nil {
		return proxy_net_DialContextWithResolver(ctx, network, address, proxyoptions, CreateHostsAndDohResolverCached(proxyoptions, dnsCache, tranportConfigurations...))
	}
	return proxy_net_DialContextOriginal(ctx, network, address, proxyoptions, tranportConfigurations...)
}

// ===== 内部拨号辅助函数 =====
// 这些函数是从options.go复制过来的，避免循环依赖

// proxy_net_DialOriginal 原始的网络拨号函数（不使用缓存）
func proxy_net_DialOriginal(network string, addr string, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	if len(proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		options.Shuffle(proxyoptions)

		var allErrors []error
		for _, opt := range proxyoptions {
			if opt.Dohalpn == "h3" {
				if opt.Dohip == "" {
					log.Printf("使用HTTP/3 DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用HTTP/3 DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl, opt.Dohip)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			} else {
				if opt.Dohip == "" {
					log.Printf("使用DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			}
		}

		if len(allErrors) > 0 {
			return nil, fmt.Errorf("所有DOH解析都失败了: %v", allErrors)
		}
	} else {
		// 尝试使用本地hosts文件解析域名
		ips, err := hosts.ResolveDomainToIPsWithHosts(hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIp(network, ips, port, options.ProxyOption{})
		}
	}

	// 如果都失败了，使用原始地址
	log.Printf("所有解析方法都失败了，使用原始地址连接: %s", addr)
	return net.Dial(network, addr)
}

// proxy_net_DialContextOriginal 原始的上下文网络拨号函数（不使用缓存）
func proxy_net_DialContextOriginal(ctx context.Context, network string, address string, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if len(proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		options.Shuffle(proxyoptions)

		var allErrors []error
		for _, opt := range proxyoptions {
			if opt.Dohalpn == "h3" {
				if opt.Dohip == "" {
					log.Printf("使用HTTP/3 DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用HTTP/3 DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl, opt.Dohip)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			} else {
				if opt.Dohip == "" {
					log.Printf("使用DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			}
		}

		if len(allErrors) > 0 {
			return nil, fmt.Errorf("所有DOH解析都失败了: %v", allErrors)
		}
	} else {
		// 尝试使用本地hosts文件解析域名
		ips, err := hosts.ResolveDomainToIPsWithHosts(hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIpContext(ctx, network, ips, port, options.ProxyOption{})
		}
	}

	// 如果都失败了，使用原始地址
	log.Printf("所有解析方法都失败了，使用原始地址连接: %s", address)
	var dialer net.Dialer
	connection, err1 := dialer.DialContext(ctx, network, address)
	return connection, err1
}

// proxy_net_DialWithResolver 使用指定的解析器进行网络连接
func proxy_net_DialWithResolver(network string, addr string, proxyoptions options.ProxyOptions, nameResolver NameResolver, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	if len(proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		options.Shuffle(proxyoptions)

		var allErrors []error
		for _, opt := range proxyoptions {
			if opt.Dohalpn == "h3" {
				if opt.Dohip == "" {
					log.Printf("使用HTTP/3 DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用HTTP/3 DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl, opt.Dohip)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			} else {
				if opt.Dohip == "" {
					log.Printf("使用DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIp(network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			}
		}

		if len(allErrors) > 0 {
			return nil, fmt.Errorf("所有DOH解析都失败了: %v", allErrors)
		}
	} else {
		// 使用提供的解析器
		ips, err := nameResolver.LookupIP(context.Background(), "tcp", hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIp(network, ips, port, options.ProxyOption{})
		}

		// 如果使用解析器失败，尝试使用本地hosts文件解析域名
		ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIp(network, ips, port, options.ProxyOption{})
		}
	}

	// 如果都失败了，使用原始地址
	log.Printf("所有解析方法都失败了，使用原始地址连接: %s", addr)
	return net.Dial(network, addr)
}

// proxy_net_DialContextWithResolver 使用指定的解析器和上下文进行网络连接
func proxy_net_DialContextWithResolver(ctx context.Context, network string, address string, proxyoptions options.ProxyOptions, nameResolver NameResolver, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if len(proxyoptions) > 0 {
		// 随机打乱 proxyoptions 顺序
		options.Shuffle(proxyoptions)

		var allErrors []error
		for _, opt := range proxyoptions {
			if opt.Dohalpn == "h3" {
				if opt.Dohip == "" {
					log.Printf("使用HTTP/3 DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用HTTP/3 DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh3(hostname, opt.Dohurl, opt.Dohip)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			} else {
				if opt.Dohip == "" {
					log.Printf("使用DOH %s 解析域名: %s", opt.Dohurl, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				} else {
					log.Printf("使用DOH %s 和IP %s 解析域名: %s", opt.Dohurl, opt.Dohip, hostname)
					ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, opt.Dohurl, opt.Dohip, tranportConfigurations...)
					if len(ips) > 0 {
						return doh.ConnectWithIpContext(ctx, network, ips, port, opt)
					}
					if len(errors) > 0 {
						allErrors = append(allErrors, errors...)
					}
				}
			}
		}

		if len(allErrors) > 0 {
			return nil, fmt.Errorf("所有DOH解析都失败了: %v", allErrors)
		}
	} else {
		// 使用提供的解析器
		ips, err := nameResolver.LookupIP(ctx, "tcp", hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIpContext(ctx, network, ips, port, options.ProxyOption{})
		}

		// 如果使用解析器失败，尝试使用本地hosts文件解析域名
		ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)
		if err == nil && len(ips) > 0 {
			// 使用解析到的IP地址直接连接
			return doh.ConnectWithIpContext(ctx, network, ips, port, options.ProxyOption{})
		}
	}

	// 如果都失败了，使用原始地址
	log.Printf("所有解析方法都失败了，使用原始地址连接: %s", address)
	var dialer net.Dialer
	connection, err1 := dialer.DialContext(ctx, network, address)
	return connection, err1
}