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

// resolveTargetAddressForHttpWithRoundRobin 从解析的IP数组中轮询选择一个地址（HTTP代理使用）
func resolveTargetAddressForHttpWithRoundRobin(addrs []string, target string) string {
	if len(addrs) == 0 {
		return target
	}

	if len(addrs) == 1 {
		return addrs[0]
	}

	// 简单轮询：基于目标字符串哈希来选择一个相对稳定的IP
	// 这样相同的域名总是会选择相同的IP，但如果有多个IP可以实现负载分散
	hash := 0
	for _, c := range target {
		hash = (hash*31 + int(c)) % len(addrs)
	}

	selectedAddr := addrs[hash]
	log.Printf("Http RoundRobin selected address %s from %v for target %s", selectedAddr, addrs, target)

	return selectedAddr
}
