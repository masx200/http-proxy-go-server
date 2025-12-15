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
	resolver := dnscache.CreateHostsAndDohResolverCached(proxyoptions, dnsCache)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for target %s: %v", host, err)
		return []string{addr}, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for target %s", host)
		return []string{addr}, fmt.Errorf("no IP addresses resolved for target %s", host)
	}

	// 返回所有解析出的IP地址
	var resolvedAddrs []string
	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		resolvedAddrs = append(resolvedAddrs, resolvedAddr)
	}

	log.Printf("Resolved target address %s to %d IP addresses: %v", addr, len(resolvedAddrs), resolvedAddrs)

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
	addrs = Shuffle(addrs)
	// 简单轮询：基于目标字符串哈希来选择一个相对稳定的IP
	// 这样相同的域名总是会选择相同的IP，但如果有多个IP可以实现负载分散
	hash := 0
	for _, c := range target {
		hash = (hash*31 + int(c)) % len(addrs)
	}

	selectedAddr := addrs[hash]
	log.Printf("RoundRobin selected address %s from %v for target %s", selectedAddr, addrs, target)

	return selectedAddr
}
func Shuffle[T any](slice []T) []T {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
