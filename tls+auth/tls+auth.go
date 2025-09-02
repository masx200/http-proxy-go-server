package tls_auth

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/masx200/http-proxy-go-server/auth"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"github.com/masx200/http-proxy-go-server/options"
)

func Tls_auth(server_cert string, server_key, hostname string, port int, username, password string, proxyoptions options.ProxyOptions,tranportConfigurations ...func(*http.Transport) *http.Transport) {

	cert, err := tls.LoadX509KeyPair(server_cert, server_key)
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", hostname+":"+fmt.Sprint(port), config)
	// tcp 连接，监听 8080 端口
	// l, err := net.Listen("tcp", ":8080")
	var l = ln
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Proxy server started on port %s", l.Addr())
	xh := http_server.GenerateRandomLoopbackIP()
	x1 := http_server.GenerateRandomIntPort()
	var upstreamAddress string = xh + ":" + fmt.Sprint(rune(x1))
	go http_server.Http(xh, x1, proxyoptions, username, password,tranportConfigurations...)
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		// go handle(client, username, password)
		go auth.Handle(client, username, password, upstreamAddress, proxyoptions,
		
		
		tranportConfigurations...)
	}
}
