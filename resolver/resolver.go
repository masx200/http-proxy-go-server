package resolver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	go_socks5 "gitee.com/masx200/go-socks5"
	"github.com/masx200/http-proxy-go-server/doh"
	"github.com/masx200/http-proxy-go-server/hosts"
	"github.com/masx200/http-proxy-go-server/options"
)

type NameResolver = go_socks5.NameResolver

/*
	  interface {
		Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
		LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
	}
*/
type HostsAndDohResolver struct {
	proxyoptions           options.ProxyOptionsDNSSLICE
	Proxy                  func(*http.Request) (*url.URL, error)
	transportConfigurations []func(*http.Transport) *http.Transport
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
				ips, errors = doh.ResolveDomainToIPsWithDoh(host, opt.Dohurl, opt.Dohip, h.Proxy, h.transportConfigurations...)
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

func CreateHostsAndDohResolver(Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, transportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	return &HostsAndDohResolver{
		proxyoptions:           proxyoptions,
		Proxy:                  Proxy,
		transportConfigurations: transportConfigurations,
	}
}
func CreateDOHResolver(Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, transportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	return &DOHResolver{
		proxyoptions:           proxyoptions,
		Proxy:                  Proxy,
		transportConfigurations: transportConfigurations,
	}
}

type DOHResolver struct {
	proxyoptions           options.ProxyOptionsDNSSLICE
	Proxy                  func(*http.Request) (*url.URL, error)
	transportConfigurations []func(*http.Transport) *http.Transport
}

// LookupIP implements NameResolver.
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

		ips, errors := doh.ResolveDomainToIPsWithDoh(host, opt.Dohurl, opt.Dohip, d.Proxy, d.transportConfigurations...)
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

func CreateDOH3Resolver(proxyoptions options.ProxyOptionsDNSSLICE) NameResolver {
	return &DOH3Resolver{
		proxyoptions: proxyoptions,
	}
}

type DOH3Resolver struct {
	proxyoptions options.ProxyOptionsDNSSLICE
}

// LookupIP implements NameResolver.
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

func CreateHostsResolver() NameResolver {
	return &HostsResolver{}
}

type HostsResolver struct {
}

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
