package hosts

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

// ResolveDomainToIPs 根据域名解析出对应的 IP 地址列表
func ResolveDomainToIPsWithHosts(domain string) ([]net.IP, error) {
	// 根据操作系统选择 hosts 文件路径
	var hostsFilePath string
	switch runtime.GOOS {
	case "windows":
		hostsFilePath = `C:\Windows\System32\drivers\etc\hosts`
	case "linux", "darwin": // darwin 是 macOS 的 GOOS 值
		hostsFilePath = "/etc/hosts"
	default:
		hostsFilePath = "/etc/hosts"
	}

	// 解析 hosts 文件
	hostsMap, err := ParseHostsFile(hostsFilePath)
	if err != nil {
		return nil, err
	}

	// 存储解析到的 IP 地址
	var ips []net.IP

	// 遍历 hostsMap，查找所有匹配的域名（包括别名）
	for host, ipStr := range hostsMap {
		if host == domain {

			//遍历ipStr，解析为net.IP
			for _, ipStr := range ipStr {
				ip := net.ParseIP(ipStr)
				if ip != nil {
					ips = append(ips, ip)
				}
			}

		}
	}

	// 如果没有找到任何 IP 地址，返回错误
	if len(ips) == 0 {
		return nil, errors.New("hostname " + domain + " no ip found")
	}
	// 将 []net.IP 转换为 []string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	// 打印日志
	log.Println("hosts resolved " + domain + " ips:" + strings.Join(ipStrings, ","))

	return ips, nil
}

// ParseHostsFile 解析系统的 hosts 文件并返回一个映射，其中键是域名，值是对应的 IP 地址
func ParseHostsFile(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hostsMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 忽略注释和空行
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		ip := strings.TrimSpace(parts[0])
		domain := strings.TrimSpace(parts[1])
		hostsMap[domain] = append(hostsMap[domain], ip)

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hostsMap, nil
}

// ResolveDomainToIPsWithCache 解析域名到 IP 地址，如果提供了缓存则使用缓存
func ResolveDomainToIPsWithCache(domain string, dnsCache interface{}) ([]net.IP, error) {
	// 首先尝试本地hosts文件解析
	ips, err := ResolveDomainToIPsWithHosts(domain)
	if err == nil && len(ips) > 0 {
		log.Printf("Resolved domain %s to IPs via hosts file: %v", domain, ips)
		return ips, nil
	}

	// 如果hosts文件中没有找到，并且提供了DNS缓存，则使用DNS缓存解析
	if dnsCache != nil {
		log.Printf("Resolving domain %s using DNS cache/DoH infrastructure", domain)
		
		// 由于循环导入问题，我们通过类型断言来调用DNS缓存
		// 检查是否是dnscache.DNSCache类型
		if typedCache, ok := dnsCache.(interface {
			GetIPs(dnsType, domain string) ([]net.IP, bool)
			SetIPs(dnsType, domain string, ips []net.IP, ttl time.Duration)
		}); ok {
			// 尝试从缓存获取
			if cachedIPs, found := typedCache.GetIPs("doh", domain); found {
				log.Printf("DNS cache hit for domain %s: %v", domain, cachedIPs)
				return cachedIPs, nil
			}
			
			// 缓存未命中，直接返回hosts文件的结果
			// 实际的DoH解析在dnscache包的resolver中进行
			log.Printf("DNS cache miss for domain %s, fallback to hosts resolution", domain)
		}
	}

	// 如果所有方法都失败，返回hosts文件的错误或空结果
	if err != nil {
		return nil, fmt.Errorf("failed to resolve domain %s: hosts file error: %v, no cache available", domain, err)
	}
	
	log.Printf("No IPs found for domain %s in hosts file", domain)
	return []net.IP{}, fmt.Errorf("no IP addresses found for domain %s", domain)
}
