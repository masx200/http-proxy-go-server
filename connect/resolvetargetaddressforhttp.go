package connect

import (
	"context"
	"fmt"
	"github.com/masx200/http-proxy-go-server/dnscache"
	"github.com/masx200/http-proxy-go-server/options"
	"log"
	"net"
)

// resolveTargetAddressForHttp 解析目标地址的域名为IP地址（用于HTTP代理）
// 返回所有解析的IP地址数组，供调用者实现轮询
func resolveTargetAddressForHttp(addr string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache) ([]string, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return []string{addr}, err
	}

	// 如果已经是IP地址，直接返回
	if net.ParseIP(host) != nil {
		return []string{addr}, nil
	}

	log.Printf("Resolving HTTP target address %s using DoH infrastructure", host)

	// 使用DoH解析
	resolver := dnscache.CreateHostsAndDohResolverCached(proxyoptions, dnsCache)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for HTTP target %s: %v", host, err)
		return []string{addr}, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for HTTP target %s", host)
		return []string{addr}, fmt.Errorf("no IP addresses resolved for HTTP target %s", host)
	}

	// 返回所有解析出的IP地址
	var resolvedAddrs []string
	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		resolvedAddrs = append(resolvedAddrs, resolvedAddr)
	}

	log.Printf("Resolved HTTP target address %s to %d IP addresses: %v", addr, len(resolvedAddrs), resolvedAddrs)

	return resolvedAddrs, nil
}
