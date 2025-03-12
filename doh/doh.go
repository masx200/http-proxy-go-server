package doh

import (
	"fmt"
	"net"
	"strings"
	"sync"

	dns_experiment "github.com/masx200/http3-reverse-proxy-server-experiment/dns"
	"github.com/miekg/dns"
)

func Dohnslookup(domain string, dnstype string, dohurl string, dohip ...string) []*dns.Msg {
	fmt.Println("domain:", domain, "dnstype:", dnstype, "dohurl:", dohurl)
	//results := make([]*dns.Msg, 0)
	var results = make([]*dns.Msg, 0)
	var wg sync.WaitGroup
	//mutex
	var mutex sync.Mutex
	for _, d := range strings.Split(domain, ",") {
		for _, t := range strings.Split(dnstype, ",") {
			wg.Add(1)
			go func(d string, t string) {
				defer wg.Done()
				fmt.Println("domain:", d, "dnstype:", t, "dohurl:", dohurl)
				var msg = &dns.Msg{}
				msg.SetQuestion(d+".", dns.StringToType[t])
				fmt.Println(msg.String())

				res, err := dns_experiment.DohClient(msg, dohurl, dohip...)
				if err != nil {
					fmt.Println(err)
					return

				}
				fmt.Println(res.String())
				mutex.Lock()

				defer mutex.Unlock()
				results = append(results, res)
			}(d, t)

		}
	}
	wg.Wait()
	return results
}
func ResolveDomainToIPsWithDoh(domain string, dohurl string, dohip ...string) ([]net.IP, error) { // 使用 A 和 AAAA 记录类型查询域名
	dnstypes := "A,AAAA"
	responses := Dohnslookup(domain, dnstypes, dohurl, dohip...)

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
		return nil, fmt.Errorf("no IP addresses found for domain %s", domain)
	}

	return ips, nil
}
