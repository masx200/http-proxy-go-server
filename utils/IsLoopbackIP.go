package utils

import "net"

func IsLoopbackIP(ipStr string) (isLoopback bool) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	// 检查是否为 IPv4
	ip = ip.To4()
	if ip == nil {
		return false // 是 IP，但不是 IPv4
	}

	// 判断是否在 127.0.0.0/8 范围内
	_, loopbackNet, err := net.ParseCIDR("127.0.0.0/8")
	if err != nil {
		return false
	}
	return loopbackNet.Contains(ip)
}
