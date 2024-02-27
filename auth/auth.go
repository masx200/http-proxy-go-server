package auth

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	// "net/url"
	"strings"

	"github.com/masx200/http-proxy-go-server/simple"
)

func Auth(hostname string, port int, username, password string) {
	// tcp 连接，监听 8080 端口
	l, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Proxy server started on port %s", l.Addr())
	xh := http_server.GenerateRandomLoopbackIP()
	x1 := http_server.GenerateRandomIntPort()
	var upstreamAddress string = xh + ":" + fmt.Sprint(rune(x1))
	go http_server.Http(xh, x1)
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go Handle(client, username, password, upstreamAddress)
	}
}

// func Main() {
// 	// tcp 连接，监听 8080 端口
// 	l, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// 死循环，每当遇到连接时，调用 handle
// 	for {
// 		client, err := l.Accept()
// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		go Handle(client, "username", "password")
// 	}
// }

func Handle(client net.Conn, username, password string, httpUpstreamAddress string) {
	if client == nil {
		return
	}
	defer client.Close()

	log.Printf("remote addr: %v\n", client.RemoteAddr())

	// 用来存放客户端数据的缓冲区
	var b [10240]byte
	//从客户端获取数据
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	var method, URL, address string
	// 从客户端数据读入 method，url
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &URL)
	fmt.Println(string(b[:bytes.IndexByte(b[:], '\n')]))
	// hostPortURL, err := url.Parse(URL)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// 检查 Proxy-Authorization 头
	proxyAuth := ""
	for _, line := range strings.Split(string(b[:n]), "\n") {
		if strings.HasPrefix(line, "Proxy-Authorization:") {
			proxyAuth = strings.TrimSpace(strings.TrimPrefix(line, "Proxy-Authorization:"))
			break
		}
	}

	// 验证身份
	if !isAuthenticated(proxyAuth, username, password) {
		fmt.Fprint(client, "HTTP/1.1 407 Proxy Authentication Required\r\nProxy-Authenticate: Basic realm=\"Proxy\"\r\n\r\n")
		fmt.Println("身份验证失败")
		return
	}
	fmt.Println("身份验证成功")
	// 如果方法是 CONNECT，则为 https 协议
	if method == "CONNECT" {
		// address = hostPortURL.Scheme + ":" + hostPortURL.Opaque
		var line = string(b[:bytes.IndexByte(b[:], '\n')])
		address = simple.ExtractAddressFromConnectRequestLine(line)
	} else { //否则为 http 协议
		// address = hostPortURL.Host
		// // 如果 host 不带端口，则默认为 80
		// if !strings.Contains(hostPortURL.Host, ":") { //host 不带端口， 默认 80
		// 	address = hostPortURL.Host + ":80"
		// }
		var line = string(b[:bytes.IndexByte(b[:], '\n')])

		// hostPortURL, err := url.Parse(line[7+1 : len(line)-9-1])
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		// address = hostPortURL.Host
		// // 如果 host 不带端口，则默认为 80
		// if !strings.Contains(hostPortURL.Host, ":") { //host 不带端口， 默认 80
		// 	address = hostPortURL.Host + ":80"
		// }
		address, err = simple.ExtractAddressFromOtherRequestLine(line)
		if err != nil {
			log.Println(err)
			return
		}
	}
	fmt.Println("address:" + address)
	var upstreamAddress string
	if method == "CONNECT" {
		upstreamAddress = address
	} else {
		upstreamAddress = httpUpstreamAddress
	}
	//获得了请求的 host 和 port，向服务端发起 tcp 连接
	server, err := net.Dial("tcp", upstreamAddress)
	if err != nil {
		log.Println(err)
		return
	}
	//如果使用 https 协议，需先向客户端表示连接建立完毕
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else { //如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端

		req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(b[:n])))
		if err != nil {
			log.Println("Error parsing request:", err)
			return
		}
		req.Header.Del("Proxy-Authorization")
		// server.Write(b[:n])
		fmt.Println(req.RequestURI)
		var requestTarget = req.RequestURI
		u, err := url.Parse(requestTarget)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to parse url: %w", err))
			return
		}
		/* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
		req.RequestURI = u.RequestURI()
		fmt.Println(req.RequestURI)
		err = req.Write(server)
		if err != nil {
			log.Println("Error writing request to server:", err)
			return
		}
	}

	//将客户端的请求转发至服务端，将服务端的响应转发给客户端。io.Copy 为阻塞函数，文件描述符不关闭就不停止
	go io.Copy(server, client)
	io.Copy(client, server)
}

func isAuthenticated(proxyAuth, expectedUsername, expectedPassword string) bool {
	if !strings.HasPrefix(proxyAuth, "Basic ") {
		return false
	}

	auth := strings.TrimPrefix(proxyAuth, "Basic ")
	decodedAuth, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return false
	}

	username, password, ok := strings.Cut(string(decodedAuth), ":")
	if !ok {
		return false
	}

	return username == expectedUsername && password == expectedPassword
}
