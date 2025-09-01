package main

import (
	"strings"
	"testing"
)

func TestSelectProxyURLWithCIDR(t *testing.T) {
	// 设置测试用的upstreams
	upstreams := map[string]UpStream{
		"proxy1": {
			HTTP_PROXY:  "http://proxy1.example.com:8080",
			HTTPS_PROXY: "https://proxy1.example.com:8080",
			BypassList:  []string{"localhost", "127.0.0.1"},
		},
		"proxy2": {
			HTTP_PROXY:  "http://proxy2.example.com:8080",
			HTTPS_PROXY: "https://proxy2.example.com:8080",
			BypassList:  []string{"*.local", "192.168.1.0/24"},
		}, "proxy3": {
			HTTP_PROXY:  "http://proxy3.example.com:8080",
			HTTPS_PROXY: "https://proxy3.example.com:8080",
			BypassList:  []string{"*.local", "192.168.1.0/24"},
		},
	}

	// 设置测试用的rules
	rules := []struct {
		Pattern  string `json:"pattern"`
		Upstream string `json:"upstream"`
	}{
		{Pattern: "google.com", Upstream: "proxy1"}, // 字符串包含匹配
		{Pattern: "github.com", Upstream: "proxy1"},
		{Pattern: "192.168.1.0/24", Upstream: "proxy2"},
		{Pattern: "10.0.0.1", Upstream: "proxy2"},
		{Pattern: "*.baidu.com", Upstream: "proxy3"},
		{Pattern: "*", Upstream: "proxy1"}, // 通配符匹配所有
	}

	tests := []struct {
		name        string
		domain      string
		expectedURL string
		expectError bool
	}{{
		name:        "Google域名包含匹配",
		domain:      "www.baidu.com",
		expectedURL: "http://proxy3.example.com:8080", // 非HTTPS域名返回HTTP代理
		expectError: false,
	},
		// 域名匹配测试（使用字符串包含匹配）
		{
			name:        "Google域名包含匹配",
			domain:      "www.google.com",
			expectedURL: "http://proxy1.example.com:8080", // 非HTTPS域名返回HTTP代理
			expectError: false,
		},
		{
			name:        "GitHub精确匹配",
			domain:      "github.com",
			expectedURL: "http://proxy1.example.com:8080", // 非HTTPS域名返回HTTP代理
			expectError: false,
		},
		{
			name:        "无效域名格式-带https协议",
			domain:      "https://github.com",
			expectedURL: "",
			expectError: true, // 带协议前缀的域名应该返回错误
		},
		{
			name:        "无效域名格式-带http协议",
			domain:      "http://github.com",
			expectedURL: "",
			expectError: true, // 带协议前缀的域名应该返回错误
		},
		// IP地址CIDR测试
		{
			name:        "CIDR范围内IP匹配",
			domain:      "192.168.1.100",
			expectedURL: "https://proxy2.example.com:8080",
			expectError: false,
		},
		{
			name:        "CIDR边界IP匹配",
			domain:      "192.168.1.1",
			expectedURL: "https://proxy2.example.com:8080",
			expectError: false,
		},
		{
			name:        "CIDR外IP不匹配",
			domain:      "192.168.2.1",
			expectedURL: "http://proxy1.example.com:8080", // IP地址没有匹配到规则，应该返回错误
			expectError: true,
		},
		// 完全相同的IP地址匹配
		{
			name:        "精确IP匹配",
			domain:      "10.0.0.1",
			expectedURL: "https://proxy2.example.com:8080",
			expectError: false,
		},
		{
			name:        "不匹配的IP",
			domain:      "10.0.0.2",
			expectedURL: "http://proxy1.example.com:8080", // IP地址没有匹配到规则，应该返回错误
			expectError: true,
		},
		// 前缀匹配测试
		{
			name:        "域名前缀匹配",
			domain:      "api.github.com",
			expectedURL: "http://proxy1.example.com:8080", // 非HTTPS域名返回HTTP代理
			expectError: false,
		},
		// 通配符匹配所有（只对域名有效）
		{
			name:        "通配符匹配任意域名",
			domain:      "unknown.domain.com",
			expectedURL: "http://proxy1.example.com:8080", // 非HTTPS域名返回HTTP代理
			expectError: false,
		},
		{
			name:        "IP地址不匹配通配符",
			domain:      "8.8.8.8",
			expectedURL: "http://proxy1.example.com:8080", // IP地址不支持通配符匹配
			expectError: true,
		},
		// 无效域名格式测试
		{
			name:        "无效域名格式-包含路径",
			domain:      "github.com/path",
			expectedURL: "",
			expectError: true, // 包含路径的域名应该返回错误
		},
		{
			name:        "无效域名格式-包含端口",
			domain:      "github.com:8080",
			expectedURL: "",
			expectError: true, // 包含端口的域名应该返回错误
		},
		{
			name:        "无效域名格式-空字符串",
			domain:      "",
			expectedURL: "",
			expectError: true, // 空字符串应该返回错误
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SelectProxyURLWithCIDR(upstreams, rules, tt.domain)

			if tt.expectError {
				if err == nil {
					t.Errorf("期望错误但没有得到错误")
				} else {
					// 检查错误消息是否包含预期的内容
					if strings.Contains(tt.domain, "https://") || strings.Contains(tt.domain, "http://") {
						if !strings.Contains(err.Error(), "invalid domain format") {
							t.Errorf("期望域名格式错误，实际得到: %v", err)
						}
					}
				}
				return
			}

			if err != nil {
				t.Errorf("意外的错误: %v", err)
				return
			}

			if result != tt.expectedURL {
				t.Errorf("期望URL: %s, 实际得到: %s", tt.expectedURL, result)
			}
		})
	}
}

func TestSelectProxyURLWithCIDR_NoMatchingRule(t *testing.T) {
	// 设置没有通配符规则的测试数据
	upstreams := map[string]UpStream{
		"proxy1": {
			HTTP_PROXY:  "http://proxy1.example.com:8080",
			HTTPS_PROXY: "https://proxy1.example.com:8080",
		},
	}

	rules := []struct {
		Pattern  string `json:"pattern"`
		Upstream string `json:"upstream"`
	}{
		{Pattern: "google.com", Upstream: "proxy1"},
		{Pattern: "192.168.1.0/24", Upstream: "proxy1"},
		// 没有通配符规则
	}

	tests := []struct {
		name   string
		domain string
	}{
		{"不匹配的域名", "unknown.com"},
		{"不匹配的IP", "8.8.8.8"},
		{"CIDR外的IP", "10.0.0.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SelectProxyURLWithCIDR(upstreams, rules, tt.domain)

			if err == nil {
				t.Errorf("期望错误但没有得到错误")
			}
		})
	}
}

func TestSelectProxyURLWithCIDR_EmptyUpstream(t *testing.T) {
	// 测试空的upstreams
	upstreams := map[string]UpStream{}

	rules := []struct {
		Pattern  string `json:"pattern"`
		Upstream string `json:"upstream"`
	}{
		{Pattern: "google.com", Upstream: "proxy1"},
	}

	_, err := SelectProxyURLWithCIDR(upstreams, rules, "google.com")

	if err == nil {
		t.Errorf("期望错误但没有得到错误")
	}
}

func TestSelectProxyURLWithCIDR_Priority(t *testing.T) {
	// 测试规则优先级
	upstreams := map[string]UpStream{
		"proxy1": {
			HTTP_PROXY:  "http://proxy1.example.com:8080",
			HTTPS_PROXY: "https://proxy1.example.com:8080",
		},
		"proxy2": {
			HTTP_PROXY:  "http://proxy2.example.com:8080",
			HTTPS_PROXY: "https://proxy2.example.com:8080",
		},
	}

	// 规则按顺序，第一个匹配的规则生效
	rules := []struct {
		Pattern  string `json:"pattern"`
		Upstream string `json:"upstream"`
	}{
		{Pattern: "com", Upstream: "proxy1"}, // 字符串包含匹配
		{Pattern: "google.com", Upstream: "proxy2"},
	}

	result, err := SelectProxyURLWithCIDR(upstreams, rules, "google.com")

	if err != nil {
		t.Errorf("意外的错误: %v", err)
		return
	}

	// 应该匹配第一个规则 "com" (字符串包含匹配)
	expectedURL := "http://proxy1.example.com:8080" // 非HTTPS域名返回HTTP代理
	if result != expectedURL {
		t.Errorf("期望URL: %s, 实际得到: %s", expectedURL, result)
	}
}
