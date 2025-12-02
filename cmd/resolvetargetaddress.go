package main

import (
	"context"
	"fmt"
	"github.com/masx200/http-proxy-go-server/dnscache"
	"github.com/masx200/http-proxy-go-server/options"
	"log"
	"net"
)

// resolveTargetAddress 解析目标地址的域名为IP地址（如果启用了upstreamResolveIPs）
func resolveTargetAddress(addr string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool) (string, error) {
	if !upstreamResolveIPs || len(proxyoptions) == 0 || dnsCache == nil {
		return addr, nil
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return addr, err
	}

	// 如果已经是IP地址，直接返回
	if net.ParseIP(host) != nil {
		return addr, nil
	}

	log.Printf("Resolving target address %s using DoH infrastructure", host)

	// 使用DoH解析
	resolver := dnscache.CreateHostsAndDohResolverCached(proxyoptions, dnsCache)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for target %s: %v", host, err)
		return addr, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for target %s", host)
		return addr, fmt.Errorf("no IP addresses resolved for target %s", host)
	}

	// 使用第一个解析出的IP
	resolvedIP := ips[0]
	resolvedAddr := net.JoinHostPort(resolvedIP.String(), port)
	log.Printf("Resolved target address %s -> %s via IP %s", addr, resolvedAddr, resolvedIP)

	return resolvedAddr, nil
}
