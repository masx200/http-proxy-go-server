package doh

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	dns_experiment "github.com/masx200/http-proxy-go-server/dns_experiment"
	"github.com/miekg/dns"
)

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
	dnstypes := "A,AAAA"
	responses, errors := Dohnslookup(domain, dnstypes, dohurl, dohip, tranportConfigurations...)
	if len(responses) == 0 && len(errors) > 0 {
		return nil, errors
	}
	var ips []net.IP
	for _, response := range responses {
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
	// 将 []net.IP 转换为 []string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	// 打印日志
	log.Println("dns resolved " + domain + " ips:[" + strings.Join(ipStrings, ",") + "]")

	return ips, nil
}
