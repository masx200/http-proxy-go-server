package options

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/masx200/http-proxy-go-server/doh"
	"github.com/masx200/http-proxy-go-server/hosts"
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

type ProxyOption struct {
	Dohurl string
	Dohip  string
}
type ProxyOptions = []ProxyOption

func Proxy_net_Dial(network string, addr string, proxyoptions ProxyOptions) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	var ctx = context.Background()
	if IsIP(hostname) {
		dialer := &net.Dialer{}
		//				// 发起连接
		return dialer.DialContext(ctx, network, addr)
	}
	var ips []net.IP
	// var errors []error
	// hostname, _, err := net.SplitHostPort(address)
	// if err != nil {
	// 	return nil, err
	// }
	ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsaray = make([]error, 0)
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

				log.Println("success connect to addr=" + addr + " by network=" + network + " by serverIP=" + serverIP)
				return connection, err1
			}
		}
		return nil, ErrorArray(errorsaray)
	}
	if len(ips) == 0 && err != nil {
		log.Println(err)
	}

	//调用ResolveDomainToIPsWithHosts函数解析域名
	if len(proxyoptions) > 0 {
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
		return Proxy_net_DialContext(ctx, network, addr, proxyoptions)
	} else {
		connection, err1 := net.Dial(network, addr)

		if err1 != nil {
			log.Println("failure connect to " + addr + " by " + network + "" + err1.Error())
			return nil, err1
		}
		log.Println("success connect to " + addr + " by " + network + "")
		return connection, err1
	}
}
func Proxy_net_DialContext(ctx context.Context, network string, address string, proxyoptions ProxyOptions) (net.Conn, error) {
	hostname, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if IsIP(hostname) {
		dialer := &net.Dialer{}
		//				// 发起连接
		return dialer.DialContext(ctx, network, address)
	}
	var ips []net.IP
	// var errors []error
	// hostname, _, err := net.SplitHostPort(address)
	// if err != nil {
	// 	return nil, err
	// }
	ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsaray = make([]error, 0)
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

				log.Println("success connect to addr=" + address + " by network=" + network + " by serverIP=" + serverIP)
				return connection, err1
			}
		}
		return nil, ErrorArray(errorsaray)
	}
	if len(ips) == 0 && err != nil {
		log.Println(err)
	}

	//调用ResolveDomainToIPsWithHosts函数解析域名
	if len(proxyoptions) > 0 {
		var errorsaray = make([]error, 0)
		Shuffle(proxyoptions)
		for _, dohurlopt := range proxyoptions {

			var dohip = dohurlopt.Dohip
			var ips []net.IP
			var errors []error
			hostname, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}
			if dohip == "" {
				ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl)
			} else {
				ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, dohip)
			}

			if len(ips) == 0 && len(errors) > 0 {
				errorsaray = append(errorsaray, errors...)
				continue
			} else {
				lengthip := len(ips)
				Shuffle(ips)
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

						log.Println("success connect to " + address + " by " + network + " by " + dohurlopt.Dohurl + " by " + dohip + " by " + serverIP)
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
		if err1 != nil {
			log.Println("failure connect to " + address + " by " + network + "" + err1.Error())
			return nil, err1
		}
		log.Println("success connect to " + address + " by " + network + "")
		return connection, err1
	}
}

// Shuffle 对切片进行随机排序
func Shuffle[T any](slice []T) {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// IsIP 判断给定的字符串是否是有效的 IPv4 或 IPv6 地址。
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}
