package options

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"

	"github.com/masx200/http-proxy-go-server/hosts"
)

type ErrorArray []error

// Error implements error interface.
func (e ErrorArray) Error() string {
	var errorMessages []string
	for _, err := range e {
		errorMessages = append(errorMessages, err.Error())
	}
	return strings.Join(errorMessages, "; ")
}

type ProxyOption struct {
	Dohurl   string
	Dohip    string
	Dohalpn  string
	Doturl   string
	Dotip    string
	Doqurl   string
	Doqip    string
	// 新增DNS协议类型字段，用于标识使用哪种DNS协议
	Protocol string // "doh", "dot", "doq", "doh3"
}

type ProxyOptions = []ProxyOption
func IsIP(domain string) bool {
	return net.ParseIP(domain) != nil
}

// formatIPs 将IP地址列表格式化为字符串
func formatIPs(ips []net.IP) string {
	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}
	return strings.Join(ipStrings, ", ")
}

// Shuffle 对切片进行随机排序
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// ResolveUpstreamDomainToIPs 解析上游代理域名到IP地址列表
// 参数:
//   - upstreamAddress: 上游代理地址，可能是域名或IP
//   - proxyoptions: DNS/DoH配置选项
//   - dnsCache: DNS缓存实例
//
// 返回值:
//   - []net.IP: 解析出的IP地址列表
//   - error: 解析过程中发生的错误
func ResolveUpstreamDomainToIPs(upstreamAddress string, proxyoptions ProxyOptions, dnsCache interface{}) ([]net.IP, error) {
	hostname, _, err := net.SplitHostPort(upstreamAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid upstream address: %s", upstreamAddress)
	}

	// 如果已经是IP地址，直接返回
	if ip := net.ParseIP(hostname); ip != nil {
		return []net.IP{ip}, nil
	}

	// 使用现有DNS基础设施解析域名
	ips, err := hosts.ResolveDomainToIPsWithCache(hostname, dnsCache)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve upstream domain %s: %v", hostname, err)
	}

	if len(ips) == 0 {
		return nil, fmt.Errorf("no IP addresses resolved for upstream domain %s", hostname)
	}

	log.Printf("Resolved upstream domain %s to %d IP addresses: %v", hostname, len(ips), ips)
	return ips, nil
}

// Proxy_net_Dial 通过指定的网络和地址建立连接，支持代理配置和传输层自定义配置。
// 如果目标地址是IP，则直接连接；否则尝试解析域名并使用解析出的IP进行连接。
// 如果本地 hosts 文件中没有解析到IP，且提供了代理选项，则使用代理连接。
//
// 参数:
//   - network: 网络类型，如 "tcp"、"udp" 等
//   - addr: 目标地址，格式为 "host:port"
//   - proxyoptions: 代理配置选项，用于指定代理服务器等信息
//   - upstreamResolveIPs: 是否启用上游IP解析功能
//   - tranportConfigurations: 可选的 http.Transport 配置函数，用于自定义传输层行为
//
// 返回值:
//   - net.Conn: 成功建立的网络连接
//   - error: 连接过程中发生的错误
func Proxy_net_Dial(network string, addr string, proxyoptions ProxyOptions, upstreamResolveIPs bool, dnsCache interface{}, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
	var ctx = context.Background()

	// DNS缓存功能现在通过interface{}调用，避免循环导入
	// 实际的DNS解析逻辑在hosts包中处理
	hostname, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	if IsIP(hostname) {
		dialer := &net.Dialer{}
		return dialer.DialContext(ctx, network, addr)
	}
	// 使用基本hosts解析
	var ips []net.IP
	ips, err = hosts.ResolveDomainToIPsWithHosts(hostname)
	if err != nil {
		connection, err1 := net.Dial(network, addr)
		if err1 != nil {
			log.Println("failure connect to " + addr + " by " + network + ": " + err1.Error())
			return nil, err1
		}
		log.Println("success connect to " + addr + " by " + network + "")
		return connection, err1
	}

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsArray = make([]error, 0)
		for i := 0; i < lengthip; i++ {

			var serverIP = ips[i].String()
			newAddr := net.JoinHostPort(serverIP, port)
			// 创建 net.Dialer 实例
			//				dialer := &net.Dialer{}
			dialer := &net.Dialer{}
			connection, err1 := dialer.DialContext(ctx, network, newAddr)

			if err1 != nil {
				errorsArray = append(errorsArray, err1)
				continue
			} else {

				log.Println("success connect to addr=" + addr + " by network=" + network + " by serverIP=" + serverIP)
				return connection, err1
			}
		}
		return nil, ErrorArray(errorsArray)
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
//   - ctx: 上下文对象
//   - network: 网络类型 (如 "tcp", "udp")
//   - address: 目标地址，格式为 "host:port"
//   - proxyoptions: 代理配置选项
//   - dnsCache: DNS缓存实例
//   - upstreamResolveIPs: 是否启用上游IP解析
//   - tranportConfigurations: 可选的传输配置函数
//
// 返回值:
//   - net.Conn: 成功建立的网络连接
//   - error: 连接过程中发生的错误
func Proxy_net_DialContext(ctx context.Context, network string, address string, proxyoptions ProxyOptions, dnsCache interface{}, upstreamResolveIPs bool, tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
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
	// 使用正确的DNS缓存解析器，支持DoH和上游IP解析
	ips, err = hosts.ResolveDomainToIPsWithCache(hostname, dnsCache)

	if len(ips) > 0 {
		Shuffle(ips)
		lengthip := len(ips)
		var errorsArray = make([]error, 0)
		for i := 0; i < lengthip; i++ {

			var serverIP = ips[i].String()
			newAddr := net.JoinHostPort(serverIP, port)
			// 创建 net.Dialer 实例
			//				dialer := &net.Dialer{}
			dialer := &net.Dialer{}
			connection, err1 := dialer.DialContext(ctx, network, newAddr)

			if err1 != nil {
				errorsArray = append(errorsArray, err1)
				continue
			} else {

				log.Println("success connect to addr=" + address + " by network=" + network + " by serverIP=" + serverIP)
				return connection, err1
			}
		}
		return nil, ErrorArray(errorsArray)
	}
	if len(ips) == 0 && err != nil {
		log.Println(err)
	}

	// 调用正确的DNS缓存函数解析域名
	if len(proxyoptions) > 0 {
		// DNS缓存功能现在通过interface{}调用，避免循环导入
		// 回退到基础连接
		connection, err1 := net.Dial(network, address)
		if err1 != nil {
			log.Println("failure connect to " + address + " by " + network + "" + err1.Error())
			return nil, err1
		}
		log.Println("success connect to " + address + " by " + network + "")
		return connection, err1
	} else {
		connection, err1 := net.Dial(network, address)

		if err1 != nil {
			log.Println("failure connect to " + address + " by " + network + "" + err1.Error())
			return nil, err1
		}
		log.Println("success connect to " + address + " by " + network + "")
		return connection, err1
	}
}