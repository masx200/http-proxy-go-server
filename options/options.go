package options

import (
	"context"
	"net"
)

type ProxyOptions struct {
	Dohurls []string
	Dohips  []string
}

func Proxy_net_Dial(network string, address string, proxyoptions ProxyOptions) (net.Conn, error) {

	if len(proxyoptions.Dohurls) > 0 {
		//		var addr=address
		//		_, port, err := net.SplitHostPort(addr)
		//		if err != nil {
		//			return nil, err
		//		}
		//		// 用指定的 IP 地址和原端口创建新地址
		//		newAddr := net.JoinHostPort(serverIP, port)
		//		// 创建 net.Dialer 实例
		//		dialer := &net.Dialer{}
		//		// 发起连接
		//		return dialer.DialContext(ctx, network, newAddr)
		var ctx = context.Background()
		return Proxy_net_DialContext(ctx, network, address, proxyoptions)
	} else {
		newVar, err1 := net.Dial(network, address)
		return newVar, err1
	}
}
func Proxy_net_DialContext(ctx context.Context, network string, address string, proxyoptions ProxyOptions) (net.Conn, error) {

	if len(proxyoptions.Dohurls) > 0 {
		return nil, nil
	} else {
		dialer := &net.Dialer{}
		newVar, err1 := dialer.DialContext(ctx, network, address)
		return newVar, err1
	}
}
