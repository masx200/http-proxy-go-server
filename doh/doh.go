package doh

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/masx200/http-proxy-go-server/dnscache"
	dns_experiment "github.com/masx200/http-proxy-go-server/dns_experiment"
	"github.com/miekg/dns"
)

// 全局DNS缓存实例
var (
	globalDNSCache *dnscache.DNSCache
	cacheOnce      sync.Once
)

// CacheConfig DNS缓存配置
type CacheConfig struct {
	Enabled          bool          `json:"enabled"`
	FilePath         string        `json:"file_path"`
	AOFPath          string        `json:"aof_path"`
	DefaultTTL       time.Duration `json:"default_ttl"`
	SaveInterval     time.Duration `json:"save_interval"`
	AOFInterval      time.Duration `json:"aof_interval"`
	AOFEnabled       bool          `json:"aof_enabled"`
}

// DefaultCacheConfig 返回默认缓存配置
func DefaultCacheConfig() *CacheConfig {
	return &CacheConfig{
		Enabled:      true,
		FilePath:     "./dns_cache.json",
		AOFPath:      "./dns_cache.aof",
		DefaultTTL:   10 * time.Minute,
		SaveInterval: 30 * time.Second,
		AOFInterval:  1 * time.Second,
		AOFEnabled:   true,
	}
}

// InitDNSCache 初始化DNS缓存
func InitDNSCache(config *CacheConfig) error {
	var err error
	cacheOnce.Do(func() {
		if !config.Enabled {
			globalDNSCache = &dnscache.DNSCache{}
			return
		}

		dnscacheConfig := &dnscache.Config{
			FilePath:        config.FilePath,
			AOFPath:         config.AOFPath,
			DefaultTTL:      config.DefaultTTL,
			SaveInterval:    config.SaveInterval,
			AOFInterval:     config.AOFInterval,
			Enabled:         config.Enabled,
		}

		globalDNSCache, err = dnscache.NewWithConfig(dnscacheConfig)
		if err != nil {
			log.Printf("初始化DNS缓存失败: %v", err)
		} else {
			log.Printf("DNS缓存初始化成功，AOF: %v", config.AOFEnabled)
		}
	})
	return err
}

// GetDNSCache 获取全局DNS缓存实例
func GetDNSCache() *dnscache.DNSCache {
	return globalDNSCache
}

// CloseDNSCache 关闭DNS缓存
func CloseDNSCache() {
	if globalDNSCache != nil {
		globalDNSCache.Close()
	}
}

func Dohnslookup(domain string, dnstype string, dohurl string, dohip string, tranportConfigurations ...func(*http.Transport) *http.Transport) ([]*dns.Msg, []error) {
	log.Println("domain:", domain, "dnstype:", dnstype, "dohurl:", dohurl)
	//results := make([]*dns.Msg, 0)
	var errors = make([]error, 0)
	var results = make([]*dns.Msg, 0)
	var wg sync.WaitGroup
	//mutex
	var mutex sync.Mutex
	for _, d := range strings.Split(domain, ",") {
		for _, t := range strings.Split(dnstype, ",") {
			wg.Add(1)
			go func(d string, t string) {
				defer wg.Done()
				log.Println("domain:", d, "dnstype:", t, "dohurl:", dohurl)
				var msg = &dns.Msg{}
				msg.SetQuestion(d+".", dns.StringToType[t])
				// log.Println(msg.String())

				res, err := dns_experiment.DohClient(msg, dohurl, dohip, tranportConfigurations...)
				mutex.Lock()

				defer mutex.Unlock()
				if err != nil {
					log.Println(err)
					errors = append(errors, err)
					return

				}
				// log.Println(res.String())

				results = append(results, res)
			}(d, t)

		}
	}
	wg.Wait()
	return results, errors
}

// ResolveDomainToIPsWithDoh 使用 A 和 AAAA 记录类型查询域名，将域名解析为 IP 地址
// 参数:
//   - domain: 要解析的域名
//   - dohurl: DNS over HTTPS (DoH) 服务的 URL
//   - dohip: 可选的 DoH 服务器 IP 地址
//
// 返回值:
//   - []net.IP: 解析得到的 IP 地址列表
//   - []error: 解析过程中出现的错误列表
func ResolveDomainToIPsWithDoh(domain string, dohurl string, dohip string, tranportConfigurations ...func(*http.Transport) *http.Transport) ([]net.IP, []error) { // 使用 A 和 AAAA 记录类型查询域名
	// 首先尝试从缓存获取
	if globalDNSCache != nil {
		if cachedIPs, found := globalDNSCache.GetIPs("A", domain); found && len(cachedIPs) > 0 {
			log.Printf("dns cache hit for %s: [%s]", domain, formatIPs(cachedIPs))
			return cachedIPs, nil
		}
		if cachedIPs, found := globalDNSCache.GetIPs("AAAA", domain); found && len(cachedIPs) > 0 {
			log.Printf("dns cache hit for %s (AAAA): [%s]", domain, formatIPs(cachedIPs))
			return cachedIPs, nil
		}
	}

	dnstypes := "A,AAAA"
	responses, errors := Dohnslookup(domain, dnstypes, dohurl, dohip, tranportConfigurations...)
	if len(responses) == 0 && len(errors) > 0 {
		return nil, errors
	}
	var ips []net.IP
	var aIPs []net.IP
	var aaaaIPs []net.IP
	for _, response := range responses {
		for _, record := range response.Answer {
			switch r := record.(type) {
			case *dns.A:
				ips = append(ips, r.A)
				aIPs = append(aIPs, r.A)
			case *dns.AAAA:
				ips = append(ips, r.AAAA)
				aaaaIPs = append(aaaaIPs, r.AAAA)
			}
		}
	}

	if len(ips) == 0 {
		return nil, []error{fmt.Errorf("no IP addresses found for domain %s", domain)}
	}

	// 将结果缓存到DNS缓存中
	if globalDNSCache != nil {
		if len(aIPs) > 0 {
			globalDNSCache.SetIPs("A", domain, aIPs, 5*time.Minute)
			log.Printf("dns cached A record for %s: [%s]", domain, formatIPs(aIPs))
		}
		if len(aaaaIPs) > 0 {
			globalDNSCache.SetIPs("AAAA", domain, aaaaIPs, 5*time.Minute)
			log.Printf("dns cached AAAA record for %s: [%s]", domain, formatIPs(aaaaIPs))
		}
	}

	// 将 []net.IP 转换为 []string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	// 打印日志
	log.Println("dns resolved " + domain + " ips:[" + strings.Join(ipStrings, ",") + "]")

	return ips, nil
}

// formatIPs 格式化IP地址列表为字符串
func formatIPs(ips []net.IP) string {
	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}
	return strings.Join(ipStrings, ",")
}
