package hosts

import (
	"bufio"
	"net"
	"os"
	"runtime"
	"strings"
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
		return nil, os.ErrNotExist
	}

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
