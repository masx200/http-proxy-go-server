package dns_experiment

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/masx200/dnsproxy/upstream"
	"github.com/miekg/dns"
)

// DoqDNSOptions DoQ DNS 选项结构体
type DoqDNSOptions struct {
	ServerURL string
	ServerIP  string // 可选的服务器IP地址
	Timeout   time.Duration
}

// ResolveDomainToIPsWithDoQ 使用 DoQ (DNS over QUIC) 协议解析域名到 IP 地址
// 参数:
//   - domain: 要解析的域名
//   - options: DoQ 配置选项
//
// 返回值:
//   - []net.IP: 解析得到的 IP 地址列表
//   - []error: 解析过程中出现的错误列表
func ResolveDomainToIPsWithDoQ(domain string, options *DoqDNSOptions) ([]net.IP, []error) {
	log.Println("DoQ resolving domain:", domain, "server:", options.ServerURL)

	if options.Timeout == 0 {
		options.Timeout = 30 * time.Second
	}

	var errors = make([]error, 0)
	var results = make([]*dns.Msg, 0)

	// 查询 A 和 AAAA 记录
	recordTypes := []uint16{dns.TypeA, dns.TypeAAAA}

	for _, recordType := range recordTypes {
		msg := &dns.Msg{}
		msg.SetQuestion(domain+".", recordType)
		msg.Id = 0 // 设置为0以支持缓存

		resp, err := DoQClientWithOptions(msg, options)
		if err != nil {
			log.Printf("DoQ query failed for %s: %v", domain, err)
			errors = append(errors, err)
			continue
		}

		if resp != nil && len(resp.Answer) > 0 {
			results = append(results, resp)
		}
	}

	if len(results) == 0 && len(errors) > 0 {
		return nil, errors
	}

	var ips []net.IP
	for _, response := range results {
		for _, record := range response.Answer {
			switch r := record.(type) {
			case *dns.A:
				ips = append(ips, r.A)
			case *dns.AAAA:
				ips = append(ips, r.AAAA)
			}
		}
	}

	if len(ips) == 0 {
		return nil, []error{fmt.Errorf("no IP addresses found for domain %s", domain)}
	}

	// 打印日志
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}
	log.Println("DoQ resolved " + domain + " ips:[" + strings.Join(ipStrings, ",") + "]")

	return ips, nil
}

// DoQClientWithOptions 使用配置选项创建 DoQ 客户端并执行查询
func DoQClientWithOptions(msg *dns.Msg, options *DoqDNSOptions) (*dns.Msg, error) {
	// 构建DoQ地址
	var doqAddr string
	if strings.HasPrefix(options.ServerURL, "quic://") {
		doqAddr = options.ServerURL
	} else {
		doqAddr = fmt.Sprintf("quic://%s", options.ServerURL)
	}

	// 创建上游选项
	opts := &upstream.Options{
		Timeout: options.Timeout,
	}

	// 如果指定了服务器IP，需要使用自定义的UpstreamOptions
	if options.ServerIP != "" {
		log.Printf("DoQ using custom server IP: %s for server: %s", options.ServerIP, options.ServerURL)

		// 创建自定义的UpstreamOptions
		customOpts := &CustomUpstreamOptions{
			Options:   opts,
			serverIP:  options.ServerIP,
			serverURL: options.ServerURL,
		}

		up, err := upstream.AddressToUpstream(doqAddr, customOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to create DoQ upstream: %v", err)
		}
		defer up.Close()

		// 执行DNS查询
		resp, err := up.Exchange(msg)
		if err != nil {
			return nil, fmt.Errorf("DoQ exchange failed: %v", err)
		}

		return resp, nil
	}

	// 创建上游（使用默认选项）
	up, err := upstream.AddressToUpstream(doqAddr, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create DoQ upstream: %v", err)
	}
	defer up.Close()

	// 执行DNS查询
	resp, err := up.Exchange(msg)
	if err != nil {
		return nil, fmt.Errorf("DoQ exchange failed: %v", err)
	}

	return resp, nil
}

// TestDoQConnection 测试 DoQ 连接
func TestDoQConnection(options *DoqDNSOptions, testDomain string) (bool, int64, string) {
	if testDomain == "" {
		testDomain = "example.com"
	}

	startTime := time.Now()

	msg := &dns.Msg{}
	msg.SetQuestion(testDomain+".", dns.TypeA)
	msg.Id = 0

	resp, err := DoQClientWithOptions(msg, options)
	if err != nil {
		responseTime := time.Since(startTime).Milliseconds()
		return false, responseTime, fmt.Sprintf("DoQ connection test failed: %v", err)
	}

	if resp == nil || len(resp.Answer) == 0 {
		responseTime := time.Since(startTime).Milliseconds()
		return false, responseTime, "No answers received"
	}

	responseTime := time.Since(startTime).Milliseconds()
	return true, responseTime, "DoQ connection successful"
}
