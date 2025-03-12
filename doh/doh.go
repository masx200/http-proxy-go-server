package doh

import (
	"fmt"
	"strings"
	"sync"

	"github.com/miekg/dns"
	dns_experiment "github.com/masx200/http3-reverse-proxy-server-experiment/dns"
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
