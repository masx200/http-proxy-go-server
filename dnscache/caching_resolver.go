package dnscache

import (
	"context"
	"fmt"
	"net"
	"time"

	go_socks5 "gitee.com/masx200/go-socks5"
	"github.com/masx200/http-proxy-go-server/resolver"
)

// CachingResolver 包装现有的NameResolver以添加缓存功能
type CachingResolver struct {
	underlying resolver.NameResolver
	cache      *DNSCache
	enabled    bool
}

// NewCachingResolver 创建带有缓存功能的解析器包装器
func NewCachingResolver(underlying resolver.NameResolver, cache *DNSCache) *CachingResolver {
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
func (c *CachingResolver) GetUnderlyingResolver() resolver.NameResolver {
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