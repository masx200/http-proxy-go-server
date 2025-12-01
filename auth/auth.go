package auth

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/masx200/http-proxy-go-server/connect"
	"github.com/masx200/http-proxy-go-server/dnscache"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

func CheckShouldUseProxy(upstreamAddress string, tranportConfigurations ...func(*http.Transport) *http.Transport) (*url.URL, error) {

	return simple.CheckShouldUseProxy(upstreamAddress, tranportConfigurations...)
}

// options.ProxyOptions
func Auth(hostname string, port int, username, password string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, tranportConfigurations ...func(*http.Transport) *http.Transport) {
	// tcp 连接，监听 8080 端口
	l, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Proxy server started on port %s", l.Addr())
	xh := http_server.GenerateRandomLoopbackIP()
	x1 := http_server.GenerateRandomIntPort()
	var upstreamAddress string = xh + ":" + fmt.Sprint(rune(x1))
	go http_server.Http(xh, x1, proxyoptions, dnsCache, username, password, upstreamResolveIPs, tranportConfigurations...)
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
			return
		}

		go Handle(client, username, password, upstreamAddress, proxyoptions, dnsCache, tranportConfigurations...)
	}
}

func Handle(client net.Conn, username, password string, httpUpstreamAddress string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache,
	tranportConfigurations ...func(*http.Transport) *http.Transport) {
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
		fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
		return
	}

	var method, URL, address string
	// 从客户端数据读入 method，url
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &URL)
	log.Println(string(b[:bytes.IndexByte(b[:], '\n')]))
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
		/* var body = "407 Proxy Authentication Required"
		fmt.Fprint(client, "HTTP/1.1 407 Proxy Authentication Required\r\ncontent-length: "+strconv.Itoa(len(body))+"\r\nProxy-Authenticate: Basic realm=\"Proxy\"\r\n\r\n")
		fmt.Fprint(client, body)
		*/
		// 创建一个新的 HTTP 响应
		resp := &http.Response{
			StatusCode: 407,
			Status:     "407 Proxy Authentication Required",
			Header: http.Header{
				"Content-Length":     []string{strconv.Itoa(len("407 Proxy Authentication Required"))},
				"Proxy-Authenticate": []string{"Basic realm=\"Proxy\""},
			},
			Body:          io.NopCloser(strings.NewReader("407 Proxy Authentication Required")),
			ContentLength: int64(len("407 Proxy Authentication Required")),
			ProtoMajor:    1,
			ProtoMinor:    1,
		}
		// 将响应写入客户端连接
		resp.Write(client)
		log.Println("身份验证失败")
		return
	}
	log.Println("身份验证成功")
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
			fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
			return
		}
	}
	log.Println("address:" + address)
	var upstreamAddress string
	if method == "CONNECT" {
		upstreamAddress = address
	} else {
		upstreamAddress = httpUpstreamAddress
	}
	var server net.Conn
	proxyURL, err := CheckShouldUseProxy(upstreamAddress, tranportConfigurations...)

	if err != nil {
		log.Println(err)
		return
	}
	// 检查是否需要使用WebSocket代理
	if proxyURL != nil && (strings.HasPrefix(proxyURL.String(), "ws://") || strings.HasPrefix(proxyURL.String(), "wss://")) {
		// 解析目标地址
		host, port, err := net.SplitHostPort(upstreamAddress)
		if err != nil {
			// 如果没有端口，尝试添加默认端口
			if strings.Contains(upstreamAddress, ":") {
				// IPv6地址
				upstreamAddress = "[" + upstreamAddress + "]:80"
			} else {
				// 域名或IPv4地址
				upstreamAddress = upstreamAddress + ":80"
			}
			host, port, err = net.SplitHostPort(upstreamAddress)
			if err != nil {
				log.Println("failed to parse address:", err)
				fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
				return
			}
		}

		// 转换端口号为整数
		portNum, err := strconv.Atoi(port)
		if err != nil {
			log.Println("failed to parse port:", err)
			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
			return
		}

		// 创建WebSocket客户端配置
		wsConfig := interfaces.ClientConfig{
			Username:   proxyURL.User.Username(),
			Password:   "",
			ServerAddr: proxyURL.String(),
			Protocol:   "websocket",
			Timeout:    30 * time.Second,
		}
		log.Println("WebSocket Config Details:")
		log.Println("host, portNum", host, portNum)
		log.Printf("  Username: %s", wsConfig.Username)
		log.Printf("  Password: %s", wsConfig.Password)
		log.Printf("  ServerAddr: %s", wsConfig.ServerAddr)
		log.Printf("  Protocol: %s", wsConfig.Protocol)
		log.Printf("  Timeout: %v", wsConfig.Timeout)
		if proxyURL.User != nil {
			if password, ok := proxyURL.User.Password(); ok {
				wsConfig.Password = password
			}
		}

		// 创建WebSocket客户端
		websocketClient := socks5_websocket_proxy_golang_websocket.NewWebSocketClient(wsConfig)

		// 连接到目标主机
		err = websocketClient.Connect(host, portNum)
		if err != nil {
			log.Println("failed to connect via WebSocket proxy:", err)
			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
			return
		}

		// 创建一个管道连接来处理WebSocket数据转发
		clientConn, serverConn := net.Pipe()

		// 在goroutine中处理WebSocket数据转发
		go func() {
			defer clientConn.Close()
			defer serverConn.Close()
			// 使用ForwardData方法处理WebSocket连接
			err := websocketClient.ForwardData(serverConn)
			if err != nil {
				log.Printf("WebSocket ForwardData error: %v\n", err)
			}
		}()

		server = clientConn
		log.Println("WebSocket代理连接成功：" + upstreamAddress)
	} else if method == "CONNECT" && proxyURL != nil {
		server, err = connect.ConnectViaHttpProxy(proxyURL, upstreamAddress)
		if err != nil {
			log.Println(err)
			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
			return
		}
		log.Println("连接成功：" + upstreamAddress)
	} else {
		server, err = dnscache.Proxy_net_DialCached("tcp", upstreamAddress, proxyoptions, upstreamResolveIPs, dnsCache, tranportConfigurations...) // net.Dial("tcp", upstreamAddress)
		if err != nil {
			log.Println(err)
			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
			return
		}
		log.Println("连接成功：" + upstreamAddress)
	}
	//获得了请求的 host 和 port，向服务端发起 tcp 连接

	//	for _, err := range errors {
	//		if err != nil {
	//			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
	//			log.Println(err)
	//			return
	//		}
	//	}
	//如果使用 https 协议，需先向客户端表示连接建立完毕
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else { //如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端

		req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(b[:n])))
		if err != nil {
			fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
			log.Println("Error parsing request:", err)
			return
		}
		/* 这里只能删除第一次请求的 Proxy-Authorization */
		//req.Header.Del("Proxy-Authorization")
		clienthost, port, err := net.SplitHostPort(client.RemoteAddr().String())
		if err != nil {
			fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
			log.Println(err)
			return
		}
		log.Println("clienthost:", clienthost)
		log.Println("clientport:", port)
		forwarded := fmt.Sprintf(
			"for=%s;by=%s;host=%s;proto=%s",
			clienthost,                  // 代理自己的标识或IP地址
			client.LocalAddr().String(), // 代理的标识
			address,                     // 原始请求的目标主机名
			"http",                      // 或者 "https" 根据实际协议
		)
		req.Header.Add("Forwarded", forwarded)
		log.Println("auth Handle", "header:")

		for k, v := range req.Header {
			// log.Println("key:", k)
			log.Println("auth Handle", k, ":", strings.Join(v, ""))
		}
		// server.Write(b[:n])
		log.Println(req.RequestURI)
		var requestTarget = req.RequestURI
		u, err := url.Parse(requestTarget)
		if err != nil {
			log.Println(fmt.Errorf("failed to parse url: %w", err))
			fmt.Fprint(client, "HTTP/1.1 500 Internal Server Error\r\n\r\n")
			return
		}
		/* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
		req.RequestURI = u.RequestURI()
		log.Println(req.RequestURI)
		req.Header = req.Header.Clone()
		err = req.Write(server)
		if err != nil {
			log.Println("Error writing request to server:", err)
			fmt.Fprint(client, "HTTP/1.1 500 Internal Server Error\r\n\r\n")
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
