package options

import (
	"context"
	"log"
	"math/rand"
	"net"
	"net/http"
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
	Dohurl  string
	Dohip   string
	Dohalpn string
}
type ProxyOptions = []ProxyOption
// Proxy_net_Dial 通过指定的网络和地址建立连接，支持代理配置和传输层自定义配置。
// 如果目标地址是IP，则直接连接；否则尝试解析域名并使用解析出的IP进行连接。
// 如果本地 hosts 文件中没有解析到IP，且提供了代理选项，则使用代理连接。
//
// 参数:
//   - network: 网络类型，如 "tcp"、"udp" 等
//   - addr: 目标地址，格式为 "host:port"
//   - proxyoptions: 代理配置选项，用于指定代理服务器等信息
//   - tranportConfigurations: 可选的 http.Transport 配置函数，用于自定义传输层行为
//
// 返回值:
//   - net.Conn: 成功建立的网络连接
//   - error: 连接过程中发生的错误
func Proxy_net_Dial(network string, addr string, proxyoptions ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
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
	// hosts没有找到域名解析ip,可以忽略这个错误
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
		return Proxy_net_DialContext(ctx, network, addr, proxyoptions, tranportConfigurations...)
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

// Proxy_net_DialContext 是一个支持代理和 DoH 解析的网络连接拨号函数。
// 它会尝试通过本地 hosts 文件解析域名，如果失败则使用提供的 DoH 配置进行解析，
// 并尝试连接到解析出的 IP 地址。
//
// 参数:
//   - ctx: 上下文，用于控制连接的生命周期
//   - network: 网络协议，例如 "tcp"、"udp"
//   - address: 目标地址，格式为 "host:port"
//   - proxyoptions: 代理选项列表，包含 DoH 配置
//   - tranportConfigurations: 可选的 http.Transport 配置函数
//
// 返回值:
//   - net.Conn: 成功建立的网络连接
//   - error: 连接过程中发生的错误
func Proxy_net_DialContext(ctx context.Context, network string, address string, proxyoptions ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
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
			var dohalpn = dohurlopt.Dohalpn
			var ips []net.IP
			var errors []error
			hostname, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}

			if dohalpn == "h3" {
				if dohip == "" {

					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh3(hostname, dohurlopt.Dohurl, dohip)
				}
			} else {
				if dohip == "" {

					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, "", tranportConfigurations...)
				} else {
					ips, errors = doh.ResolveDomainToIPsWithDoh(hostname, dohurlopt.Dohurl, dohip, tranportConfigurations...)
				}
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

						log.Println("success connect to address=" + address + " by network=" + network + " by Dohurl=" + dohurlopt.Dohurl + " by dohip=" + dohip + " by serverIP=" + serverIP)
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
