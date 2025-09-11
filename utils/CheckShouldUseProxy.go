package utils

import (
	"log"
	"net/http"
	"net/url"
)

func CheckShouldUseProxy(upstreamAddress string, tranportConfigurations ...func(*http.Transport) *http.Transport) (*url.URL, error) {
	log.Println("开始检查CheckShouldUseProxy", upstreamAddress)
	// clienthost, port, err := net.SplitHostPort(upstreamAddress)
	// if err != nil {
	// 	return nil, err
	// }

	var transport = http.DefaultTransport
	for _, f := range tranportConfigurations {
		if t, ok := transport.(*http.Transport); ok {
			transport = f(t)
		}
	}
	if t, ok := transport.(*http.Transport); ok {

		var proxy = t.Proxy
		if proxy != nil {
			req, err := http.NewRequest("GET", "https://"+upstreamAddress, nil)
			if err != nil {
				return nil, err
			}
			return proxy(req)
		} else {
			return nil, nil
		}
	}
	return nil, nil
}
