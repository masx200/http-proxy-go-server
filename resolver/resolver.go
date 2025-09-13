package resolver

import (
	"context"
	"net"
)

type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
}
