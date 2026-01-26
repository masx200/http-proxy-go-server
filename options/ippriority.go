package options

import (
	"net"
	"strings"
)

// IPPriority 定义 IP 优先级策略类型
type IPPriority string

const (
	IPPv4Priority  IPPriority = "ipv4"  // IPv4 优先
	IPPv6Priority  IPPriority = "ipv6"  // IPv6 优先
	IPRandomPriority IPPriority = "random" // IPv4 和 IPv6 随机
)

// String 实现 Stringer 接口
func (p IPPriority) String() string {
	return string(p)
}

// ParseIPPriority 解析 IP 优先级策略字符串
func ParseIPPriority(s string) IPPriority {
	switch strings.ToLower(s) {
	case "ipv4":
		return IPPv4Priority
	case "ipv6":
		return IPPv6Priority
	case "random":
		return IPRandomPriority
	default:
		return IPRandomPriority // 默认为 IPv4 优先
	}
}

// SortAddressesByPriority 根据 IP 优先级策略对地址进行排序
// 返回排序后的地址列表
func SortAddressesByPriority(addrs []string, priority IPPriority) []string {
	if len(addrs) == 0 {
		return addrs
	}

	var ipv4Addrs []string
	var ipv6Addrs []string

	// 分离 IPv4 和 IPv6 地址
	for _, addr := range addrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			// 如果无法解析，保持原样
			ipv4Addrs = append(ipv4Addrs, addr)
			continue
		}
		if ip := net.ParseIP(host); ip != nil {
			if ip.To4() != nil {
				ipv4Addrs = append(ipv4Addrs, addr)
			} else {
				ipv6Addrs = append(ipv6Addrs, addr)
			}
		} else {
			// 如果不是 IP 地址，保持原样
			ipv4Addrs = append(ipv4Addrs, addr)
		}
	}

	// 根据优先级策略排序
	switch priority {
	case IPPv6Priority:
		// IPv6 优先
		result := append(ipv6Addrs, ipv4Addrs...)
		return result
	case IPRandomPriority:
		// 随机混合（使用 Shuffle）
		mixed := append(ipv4Addrs, ipv6Addrs...)
		return Shuffle(mixed)
	case IPPv4Priority:
		// IPv4 优先（默认）
		fallthrough
	default:
		result := append(ipv4Addrs, ipv6Addrs...)
		return result
	}
}
