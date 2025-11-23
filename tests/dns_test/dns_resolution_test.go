package dns_test

import (
	"net"
	"testing"
	"time"

	"github.com/masx200/http-proxy-go-server/dns_experiment"
	"github.com/masx200/http-proxy-go-server/doh"
)

const (
	TestDomain     = "www.baidu.com"
	AliDNSDoHURL   = "https://dns.alidns.com/dns-query"
	AliDNSDoH3URL  = "https://dns.alidns.com/dns-query"
	AliDNSDoTURL   = "tls://dns.alidns.com:853"
	AliDNSDoQURL   = "quic://dns.alidns.com:853"
	AliDNSDoTIP    = "223.5.5.5"
	AliDNSDoQIP    = "223.5.5.5"
	AliDNSDoHIP    = "223.5.5.5"
	Timeout        = 10 * time.Second
)

func TestDoHResolution(t *testing.T) {
	t.Log("Testing DoH resolution with AliDNS")

	// 测试DoH解析
	ips, errors := doh.ResolveDomainToIPsWithDoh(TestDomain, AliDNSDoHURL, AliDNSDoHIP)

	// 检查错误
	if len(errors) > 0 {
		t.Logf("DoH resolution encountered errors: %v", errors)
	}

	// 检查结果
	if len(ips) == 0 {
		t.Fatal("DoH resolution failed: no IP addresses returned")
	}

	// 验证返回的IP地址
	for _, ip := range ips {
		t.Logf("DoH resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoH3Resolution(t *testing.T) {
	t.Log("Testing DoH3 resolution with AliDNS")

	// 测试DoH3解析
	ips, errors := doh.ResolveDomainToIPsWithDoh3(TestDomain, AliDNSDoH3URL)

	// 检查错误
	if len(errors) > 0 {
		t.Logf("DoH3 resolution encountered errors: %v", errors)
	}

	// 检查结果
	if len(ips) == 0 {
		t.Fatal("DoH3 resolution failed: no IP addresses returned")
	}

	// 验证返回的IP地址
	for _, ip := range ips {
		t.Logf("DoH3 resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoTResolution(t *testing.T) {
	t.Log("Testing DoT resolution with AliDNS")

	// 创建DoT选项
	options := &dns_experiment.DotDNSOptions{
		ServerURL: AliDNSDoTURL,
		ServerIP:  AliDNSDoTIP,
		Timeout:   Timeout,
	}

	// 测试DoT解析
	ips, errors := dns_experiment.ResolveDomainToIPsWithDoT(TestDomain, options)

	// 检查错误
	if len(errors) > 0 {
		t.Logf("DoT resolution encountered errors: %v", errors)
	}

	// 检查结果
	if len(ips) == 0 {
		t.Fatal("DoT resolution failed: no IP addresses returned")
	}

	// 验证返回的IP地址
	for _, ip := range ips {
		t.Logf("DoT resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoQResolution(t *testing.T) {
	t.Log("Testing DoQ resolution with AliDNS")

	// 创建DoQ选项
	options := &dns_experiment.DoqDNSOptions{
		ServerURL: AliDNSDoQURL,
		ServerIP:  AliDNSDoQIP,
		Timeout:   Timeout,
	}

	// 测试DoQ解析
	ips, errors := dns_experiment.ResolveDomainToIPsWithDoQ(TestDomain, options)

	// 检查错误
	if len(errors) > 0 {
		t.Logf("DoQ resolution encountered errors: %v", errors)
	}

	// 检查结果
	if len(ips) == 0 {
		t.Fatal("DoQ resolution failed: no IP addresses returned")
	}

	// 验证返回的IP地址
	for _, ip := range ips {
		t.Logf("DoQ resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoTResolutionWithoutIP(t *testing.T) {
	t.Log("Testing DoT resolution without custom IP")

	// 创建DoT选项（不指定IP）
	options := &dns_experiment.DotDNSOptions{
		ServerURL: AliDNSDoTURL,
		Timeout:   Timeout,
	}

	// 测试DoT解析
	ips, errors := dns_experiment.ResolveDomainToIPsWithDoT(TestDomain, options)

	// 检查错误
	if len(errors) > 0 {
		t.Logf("DoT resolution without IP encountered errors: %v", errors)
	}

	// 检查结果
	if len(ips) == 0 {
		t.Fatal("DoT resolution without IP failed: no IP addresses returned")
	}

	// 验证返回的IP地址
	for _, ip := range ips {
		t.Logf("DoT (no IP) resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoQResolutionWithoutIP(t *testing.T) {
	t.Log("Testing DoQ resolution by first resolving dns.alidns.com with DoH")

	// 首先使用DoH解析 dns.alidns.com 的IP地址
	dohServerIPs, dohErrors := doh.ResolveDomainToIPsWithDoh("dns.alidns.com", AliDNSDoHURL, AliDNSDoHIP)

	if len(dohErrors) > 0 {
		t.Logf("DoH resolution for dns.alidns.com encountered errors: %v", dohErrors)
	}

	if len(dohServerIPs) == 0 {
		t.Skip("Cannot resolve dns.alidns.com with DoH - skipping DoQ test")
		return
	}

	t.Logf("DoH resolved dns.alidns.com to %d IPs", len(dohServerIPs))
	for _, ip := range dohServerIPs {
		t.Logf("  dns.alidns.com -> %s", ip.String())
	}

	// 尝试使用解析出的每个IP地址进行DoQ测试
	var successfulResults []net.IP
	var allErrors []error

	for _, serverIP := range dohServerIPs {
		if !isValidIP(serverIP) {
			t.Logf("Skipping invalid IP: %s", serverIP.String())
			continue
		}

		t.Logf("Testing DoQ with server IP: %s", serverIP.String())

		// 创建DoQ选项（使用解析出的IP）
		options := &dns_experiment.DoqDNSOptions{
			ServerURL: AliDNSDoQURL,
			ServerIP:  serverIP.String(),
			Timeout:   Timeout,
		}

		// 测试DoQ解析
		ips, errors := dns_experiment.ResolveDomainToIPsWithDoQ(TestDomain, options)

		if len(errors) > 0 {
			t.Logf("DoQ with IP %s encountered errors: %v", serverIP.String(), errors)
			allErrors = append(allErrors, errors...)
			continue
		}

		if len(ips) > 0 {
			t.Logf("DoQ with IP %s successfully resolved %s to %d IPs", serverIP.String(), TestDomain, len(ips))
			successfulResults = append(successfulResults, ips...)
			break // 找到一个工作的IP就够了
		}
	}

	// 检查是否有任何成功的DoQ解析
	if len(successfulResults) == 0 {
		if len(allErrors) > 0 {
			t.Skip("DoQ resolution failed with all resolved IPs - likely due to network environment restrictions")
		} else {
			t.Skip("DoQ resolution failed - no successful results from any IP")
		}
		return
	}

	// 验证返回的IP地址
	for _, ip := range successfulResults {
		t.Logf("DoQ (via DoH-resolved IP) resolved %s to %s", TestDomain, ip.String())
		if !isValidIP(ip) {
			t.Errorf("Invalid IP address returned: %s", ip.String())
		}
	}
}

func TestDoTConnection(t *testing.T) {
	t.Log("Testing DoT connection")

	options := &dns_experiment.DotDNSOptions{
		ServerURL: AliDNSDoTURL,
		ServerIP:  AliDNSDoTIP,
		Timeout:   Timeout,
	}

	success, responseTime, message := dns_experiment.TestDoTConnection(options, TestDomain)

	t.Logf("DoT Connection Test - Success: %v, Response Time: %dms, Message: %s",
		success, responseTime, message)

	if !success {
		t.Errorf("DoT connection test failed: %s", message)
	}

	if responseTime <= 0 {
		t.Errorf("Invalid response time: %dms", responseTime)
	}
}

func TestDoQConnection(t *testing.T) {
	t.Log("Testing DoQ connection")

	options := &dns_experiment.DoqDNSOptions{
		ServerURL: AliDNSDoQURL,
		ServerIP:  AliDNSDoQIP,
		Timeout:   Timeout,
	}

	success, responseTime, message := dns_experiment.TestDoQConnection(options, TestDomain)

	t.Logf("DoQ Connection Test - Success: %v, Response Time: %dms, Message: %s",
		success, responseTime, message)

	if !success {
		t.Errorf("DoQ connection test failed: %s", message)
	}

	if responseTime <= 0 {
		t.Errorf("Invalid response time: %dms", responseTime)
	}
}

func TestDoQConnectionViaDoH(t *testing.T) {
	t.Log("Testing DoQ connection by first resolving dns.alidns.com with DoH")

	// 首先使用DoH解析 dns.alidns.com 的IP地址
	dohServerIPs, dohErrors := doh.ResolveDomainToIPsWithDoh("dns.alidns.com", AliDNSDoHURL, AliDNSDoHIP)

	if len(dohErrors) > 0 {
		t.Logf("DoH resolution for dns.alidns.com encountered errors: %v", dohErrors)
	}

	if len(dohServerIPs) == 0 {
		t.Skip("Cannot resolve dns.alidns.com with DoH - skipping DoQ connection test")
		return
	}

	t.Logf("DoH resolved dns.alidns.com to %d IPs", len(dohServerIPs))

	// 尝试使用解析出的每个IP地址进行DoQ连接测试
	var successfulConnection bool
	var lastError string

	for _, serverIP := range dohServerIPs {
		if !isValidIP(serverIP) {
			t.Logf("Skipping invalid IP: %s", serverIP.String())
			continue
		}

		t.Logf("Testing DoQ connection with server IP: %s", serverIP.String())

		// 创建DoQ选项（使用解析出的IP）
		options := &dns_experiment.DoqDNSOptions{
			ServerURL: AliDNSDoQURL,
			ServerIP:  serverIP.String(),
			Timeout:   Timeout,
		}

		// 测试DoQ连接
		success, responseTime, message := dns_experiment.TestDoQConnection(options, TestDomain)

		t.Logf("DoQ Connection Test with IP %s - Success: %v, Response Time: %dms, Message: %s",
			serverIP.String(), success, responseTime, message)

		if success {
			t.Logf("DoQ connection successful with IP %s, response time: %dms", serverIP.String(), responseTime)
			successfulConnection = true
			break // 找到一个工作的连接就够了
		} else {
			lastError = message
			t.Logf("DoQ connection failed with IP %s: %s", serverIP.String(), message)
		}
	}

	if !successfulConnection {
		if lastError != "" {
			t.Skip("DoQ connection failed with all resolved IPs - likely due to network environment restrictions")
		} else {
			t.Skip("DoQ connection failed - no successful connections from any IP")
		}
		return
	}
}

// 综合测试：比较所有DNS协议的解析结果
func TestAllDNSProtocolsComparison(t *testing.T) {
	t.Log("Testing comparison of all DNS protocols")

	testDomain := TestDomain
	results := make(map[string][]net.IP)

	// DoH测试
	if ips, _ := doh.ResolveDomainToIPsWithDoh(testDomain, AliDNSDoHURL, AliDNSDoHIP); len(ips) > 0 {
		results["DoH"] = ips
		t.Logf("DoH resolved %d IPs", len(ips))
	}

	// DoH3测试
	if ips, _ := doh.ResolveDomainToIPsWithDoh3(testDomain, AliDNSDoH3URL); len(ips) > 0 {
		results["DoH3"] = ips
		t.Logf("DoH3 resolved %d IPs", len(ips))
	}

	// DoT测试
	dotOptions := &dns_experiment.DotDNSOptions{
		ServerURL: AliDNSDoTURL,
		ServerIP:  AliDNSDoTIP,
		Timeout:   Timeout,
	}
	if ips, _ := dns_experiment.ResolveDomainToIPsWithDoT(testDomain, dotOptions); len(ips) > 0 {
		results["DoT"] = ips
		t.Logf("DoT resolved %d IPs", len(ips))
	}

	// DoQ测试
	doqOptions := &dns_experiment.DoqDNSOptions{
		ServerURL: AliDNSDoQURL,
		ServerIP:  AliDNSDoQIP,
		Timeout:   Timeout,
	}
	if ips, _ := dns_experiment.ResolveDomainToIPsWithDoQ(testDomain, doqOptions); len(ips) > 0 {
		results["DoQ"] = ips
		t.Logf("DoQ resolved %d IPs", len(ips))
	}

	// 验证结果
	if len(results) == 0 {
		t.Fatal("No DNS protocol returned any results")
	}

	t.Logf("DNS Resolution Summary:")
	for protocol, ips := range results {
		t.Logf("  %s: %v", protocol, formatIPs(ips))
	}
}

// 辅助函数：验证IP地址是否有效
func isValidIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	// 检查是否为有效的外网IP（排除内网IP）
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return false
	}

	if ip4 := ip.To4(); ip4 != nil {
		// 排除私有IPv4地址
		privateIPs := []string{
			"10.0.0.0/8",
			"172.16.0.0/12",
			"192.168.0.0/16",
		}

		for _, privateIP := range privateIPs {
			_, privateNet, _ := net.ParseCIDR(privateIP)
			if privateNet.Contains(ip4) {
				return false
			}
		}
		return true
	}

	// 对于IPv6，简单地检查是否不是本地地址
	return !ip.IsLoopback()
}

// 辅助函数：格式化IP地址列表
func formatIPs(ips []net.IP) []string {
	var result []string
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result
}

// 基准测试：比较不同DNS协议的性能
func BenchmarkDoHResolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = doh.ResolveDomainToIPsWithDoh(TestDomain, AliDNSDoHURL, AliDNSDoHIP)
	}
}

func BenchmarkDoH3Resolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = doh.ResolveDomainToIPsWithDoh3(TestDomain, AliDNSDoH3URL)
	}
}

func BenchmarkDoTResolution(b *testing.B) {
	options := &dns_experiment.DotDNSOptions{
		ServerURL: AliDNSDoTURL,
		ServerIP:  AliDNSDoTIP,
		Timeout:   Timeout,
	}

	for i := 0; i < b.N; i++ {
		_, _ = dns_experiment.ResolveDomainToIPsWithDoT(TestDomain, options)
	}
}

func BenchmarkDoQResolution(b *testing.B) {
	options := &dns_experiment.DoqDNSOptions{
		ServerURL: AliDNSDoQURL,
		ServerIP:  AliDNSDoQIP,
		Timeout:   Timeout,
	}

	for i := 0; i < b.N; i++ {
		_, _ = dns_experiment.ResolveDomainToIPsWithDoQ(TestDomain, options)
	}
}