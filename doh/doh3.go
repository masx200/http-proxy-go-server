package doh

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	h3_experiment "github.com/masx200/http3-reverse-proxy-server-experiment/h3"
	"github.com/miekg/dns"
)

func Doh3nslookup(domain string, dnstype string, dohurl string, dohip ...string) ([]*dns.Msg, []error) {
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

				res, err := h3_experiment.DoHTTP3Client(msg, dohurl, dohip...)
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
func ResolveDomainToIPsWithDoh3(domain string, dohurl string, dohip ...string) ([]net.IP, []error) { // 使用 A 和 AAAA 记录类型查询域名
	dnstypes := "A,AAAA"
	responses, errors := Doh3nslookup(domain, dnstypes, dohurl, dohip...)
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
