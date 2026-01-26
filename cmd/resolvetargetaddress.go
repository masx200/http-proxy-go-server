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
func resolveTargetAddress(addr string, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, ipPriority options.IPPriority) ([]string, error) {
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

	// 收集所有解析出的IP地址
	var resolvedAddrs []string
	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		resolvedAddrs = append(resolvedAddrs, resolvedAddr)
	}

	// 根据 IP 优先级策略排序
	resolvedAddrs = options.SortAddressesByPriority(resolvedAddrs, ipPriority)

	// 统计 IPv4 和 IPv6 地址数量
	var ipv4Count, ipv6Count int
	for _, addr := range resolvedAddrs {
		host, _, _ := net.SplitHostPort(addr)
		if ip := net.ParseIP(host); ip != nil {
			if ip.To4() != nil {
				ipv4Count++
			} else {
				ipv6Count++
			}
		}
	}

	log.Printf("Resolved target address %s to %d IP addresses (IPv4: %d, IPv6: %d, priority: %s): %v",
		addr, len(resolvedAddrs), ipv4Count, ipv6Count, ipPriority, resolvedAddrs)

	return resolvedAddrs, nil
}

// resolveTargetAddressWithRoundRobin 从解析的IP数组中根据优先级策略选择一个地址
func resolveTargetAddressWithRoundRobin(addrs []string, target string, ipPriority options.IPPriority) string {
	if len(addrs) == 0 {
		return target
	}

	if len(addrs) == 1 {
		return addrs[0]
	}

	var ipv4Addrs []string
	var ipv6Addrs []string

	// 分离 IPv4 和 IPv6 地址
	for _, addr := range addrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			// 如果无法解析，作为 IPv4 处理
			ipv4Addrs = append(ipv4Addrs, addr)
			continue
		}
		if ip := net.ParseIP(host); ip != nil {
			if ip.To4() != nil {
				ipv4Addrs = append(ipv4Addrs, addr)
			} else {
				ipv6Addrs = append(ipv6Addrs, addr)
			}
		} else {
			// 如果不是 IP 地址，作为 IPv4 处理
			ipv4Addrs = append(ipv4Addrs, addr)
		}
	}

	var candidateAddrs []string

	// 根据优先级策略选择候选地址列表
	switch ipPriority {
	case options.IPv6Priority:
		if len(ipv6Addrs) > 0 {
			candidateAddrs = ipv6Addrs
		} else {
			candidateAddrs = ipv4Addrs
		}
		log.Printf("IPv6 priority: selecting from %d IPv6 addresses", len(ipv6Addrs))
	case options.IPv4Priority:
		if len(ipv4Addrs) > 0 {
			candidateAddrs = ipv4Addrs
		} else {
			candidateAddrs = ipv6Addrs
		}
		log.Printf("IPv4 priority: selecting from %d IPv4 addresses", len(ipv4Addrs))
	case options.IPRandomPriority:
		candidateAddrs = addrs
		log.Printf("Random priority: selecting from all %d addresses", len(addrs))
	default:
		// 默认 IPv4 优先
		if len(ipv4Addrs) > 0 {
			candidateAddrs = ipv4Addrs
		} else {
			candidateAddrs = ipv6Addrs
		}
	}

	// 如果候选地址为空，使用所有地址
	if len(candidateAddrs) == 0 {
		candidateAddrs = addrs
	}

	// 从候选地址中随机选择一个
	candidateAddrs = Shuffle(candidateAddrs)
	selectedAddr := candidateAddrs[rand.Intn(len(candidateAddrs))]

	log.Printf("RoundRobin (priority=%s) selected address %s from %v for target %s", ipPriority, selectedAddr, addrs, target)

	return selectedAddr
}
func Shuffle[T any](slice []T) []T {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
