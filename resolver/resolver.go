package resolver

import (
	"context"
	"net"
	"net/http"

	"github.com/masx200/http-proxy-go-server/options"
)

type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
}
type HostsAndDohResolver struct{}

// LookupIP implements NameResolver.
func (h *HostsAndDohResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	panic("unimplemented")
}

// Resolve implements NameResolver.
func (h *HostsAndDohResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	panic("unimplemented")
}

func CreateHostsAndDohResolver(proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	return &HostsAndDohResolver{}
}
func CreateDOHResolver(proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) NameResolver {
	return &DOHResolver{}
}

type DOHResolver struct{}

// LookupIP implements NameResolver.
func (d *DOHResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	panic("unimplemented")
}

// Resolve implements NameResolver.
func (d *DOHResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	panic("unimplemented")
}

func CreateDOH3Resolver(proxyoptions options.ProxyOptions) NameResolver {
	return &DOH3Resolver{}
}

type DOH3Resolver struct{}

// LookupIP implements NameResolver.
func (d *DOH3Resolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	panic("unimplemented")
}

// Resolve implements NameResolver.
func (d *DOH3Resolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	panic("unimplemented")
}

func CreateHostsResolver() NameResolver {
	return &HostsResolver{}
}

type HostsResolver struct {
}

// LookupIP implements NameResolver.
func (h *HostsResolver) LookupIP(ctx context.Context, network string, host string) ([]net.IP, error) {
	panic("unimplemented")
}

// Resolve implements NameResolver.
func (h *HostsResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	panic("unimplemented")
}
