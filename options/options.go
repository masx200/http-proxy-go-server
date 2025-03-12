package options

import (
	"context"
	"net"
	"strings"

	"github.com/masx200/http-proxy-go-server/doh"
)

type ErrorArray []error

// Error implements error.
func (e ErrorArray) Error() string {
	// 将 ErrorArray 中的每个 error 转换为字符串
	errStrings := make([]string, len(e))
	for i, err := range e {
		errStrings[i] = err.Error()
	}
	// 使用逗号分隔符连接所有错误字符串
	return "ErrorArray:[" + strings.Join(errStrings, ", ") + "]"
}

func init() {
	var _ error = ErrorArray{}
}

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
		connection, err1 := net.Dial(network, address)
		return connection, err1
	}
}
func Proxy_net_DialContext(ctx context.Context, network string, address string, proxyoptions ProxyOptions) (net.Conn, error) {

	if len(proxyoptions.Dohurls) > 0 {
		var errorsaray = make([]error, 0)
		for index, dohurl := range proxyoptions.Dohurls {
			var dohip = proxyoptions.Dohips[index]
			var ips []net.IP

			hostname, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}
			ips, errors := doh.ResolveDomainToIPsWithDoh(hostname, dohurl, dohip)
			if len(ips) == 0 && len(errors) > 0 {
				errorsaray = append(errorsaray, errors...)
				continue
			} else {
				lengthip := len(ips)
				for i := 0; i < lengthip; i++ {

					var serverIP = ips[i].String()
					newAddr := net.JoinHostPort(serverIP, port)
					// 创建 net.Dialer 实例
					//				dialer := &net.Dialer{}
					dialer := &net.Dialer{}
					connection, err1 := dialer.DialContext(ctx, network, newAddr)

					if err1 != nil {
						errorsaray = append(errorsaray, err1)
						continue
					} else {
						return connection, err1
					}
				}
				// var serverIP = ips[0].String()
				// 用指定的 IP 地址和原端口创建新地址

			}
		}
		return nil, ErrorArray(errorsaray)
	} else {
		dialer := &net.Dialer{}
		connection, err1 := dialer.DialContext(ctx, network, address)
		return connection, err1
	}
}
