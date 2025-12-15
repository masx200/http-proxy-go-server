package utils

import (
	"log"
	"net"
	"net/http"
	"net/url"
)

func CheckShouldUseProxy(upstreamAddress string, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) (*url.URL, error) {

	var addr = upstreamAddress
	var host, _, err = net.SplitHostPort(addr)
	if err != nil {
		if addrErr, ok := err.(*net.AddrError); ok && addrErr.Err == "missing port in address" {
			host = addr // 整个字符串就是 host
		} else {
			return nil, err
		}
	}
	if IsLoopbackIP(host) {
		return nil, nil
	}

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
	// if t, ok := transport.(*http.Transport); ok {

	var proxy = Proxy
	if proxy != nil {
		req, err := http.NewRequest("GET", "https://"+upstreamAddress, nil)
		if err != nil {
			return nil, err
		}
		return proxy(req)
	} else {
		return nil, nil
	}
	// }
	return nil, nil
}
