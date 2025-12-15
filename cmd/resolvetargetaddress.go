package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
	"github.com/masx200/http-proxy-go-server/options"
)

// resolveTargetAddress 解析目标地址的域名为IP地址（如果启用了upstreamResolveIPs）
// 返回所有解析的IP地址数组，供调用者实现轮询
func resolveTargetAddress(addr string, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool) ([]string, error) {
	if !upstreamResolveIPs || len(proxyoptions) == 0 || dnsCache == nil {
		return []string{addr}, nil
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return []string{addr}, err
	}

	// 如果已经是IP地址，直接返回
	if net.ParseIP(host) != nil {
		return []string{addr}, nil
	}

	log.Printf("Resolving target address %s using DoH infrastructure", host)

	// 使用DoH解析
	resolver := dnscache.CreateHostsAndDohResolverCachedSimple(proxyoptions, dnsCache, Proxy)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for target %s: %v", host, err)
		return []string{addr}, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for target %s", host)
		return []string{addr}, fmt.Errorf("no IP addresses resolved for target %s", host)
	}

	// 返回所有解析出的IP地址，优先使用IPv4
	var resolvedAddrs []string
	var ipv4Addrs []string
	var ipv6Addrs []string

	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		if ip.To4() != nil {
			// IPv4地址优先添加
			ipv4Addrs = append(ipv4Addrs, resolvedAddr)
		} else {
			// IPv6地址后添加
			ipv6Addrs = append(ipv6Addrs, resolvedAddr)
		}
	}

	// IPv4地址在前，IPv6地址在后
	resolvedAddrs = append(ipv4Addrs, ipv6Addrs...)

	log.Printf("Resolved target address %s to %d IP addresses (IPv4: %d, IPv6: %d): %v",
		addr, len(resolvedAddrs), len(ipv4Addrs), len(ipv6Addrs), resolvedAddrs)

	return resolvedAddrs, nil
}

// resolveTargetAddressWithRoundRobin 从解析的IP数组中轮询选择一个地址
func resolveTargetAddressWithRoundRobin(addrs []string, target string) string {
	if len(addrs) == 0 {
		return target
	}

	if len(addrs) == 1 {
		return addrs[0]
	}

	// 分离IPv4和IPv6地址
	var ipv4Addrs []string
	var ipv6Addrs []string

	for _, addr := range addrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			continue
		}
		if ip := net.ParseIP(host); ip != nil && ip.To4() != nil {
			ipv4Addrs = append(ipv4Addrs, addr)
		} else {
			ipv6Addrs = append(ipv6Addrs, addr)
		}
	}

	// 优先选择IPv4地址
	var candidateAddrs []string
	if len(ipv4Addrs) > 0 {
		candidateAddrs = ipv4Addrs
		log.Printf("Preferring IPv4 addresses for SOCKS5 compatibility: %v", ipv4Addrs)
	} else {
		candidateAddrs = ipv6Addrs
		log.Printf("No IPv4 addresses available, using IPv6: %v", ipv6Addrs)
	}

	candidateAddrs = Shuffle(candidateAddrs)

	// 简单轮询：基于目标字符串哈希来选择一个相对稳定的IP
	hash := 0
	for _, c := range target {
		hash = (hash*31 + int(c)) % len(candidateAddrs)
	}

	selectedAddr := candidateAddrs[hash]
	log.Printf("RoundRobin selected address %s from %v for target %s", selectedAddr, candidateAddrs, target)

	return selectedAddr
}
func Shuffle[T any](slice []T) []T {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
