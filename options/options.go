package options

import "net"

type ProxyOptions struct {
	Dohurls []string
	Dohips  []string
}

func Proxy_net_Dial(network string, address string, proxyoptions ProxyOptions) (net.Conn, []error) {

	if len(proxyoptions.Dohurls) > 0 {
		return nil, nil
	} else {
		newVar, err1 := net.Dial(network, address)
		return newVar, []error{err1}
	}
}
