package simple

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	// "regexp"

	"github.com/masx200/http-proxy-go-server/connect"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

func Simple(hostname string, port int, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) {
	// tcp 连接，监听 8080 端口
	l, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Proxy server started on port %s", l.Addr())
	xh := http_server.GenerateRandomLoopbackIP()
	x1 := http_server.GenerateRandomIntPort()
	var upstreamAddress string = xh + ":" + fmt.Sprint(rune(x1))
	go http_server.Http(xh, x1, proxyoptions, "", "", tranportConfigurations...)
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
			return
		}

		go Handle(client, upstreamAddress, proxyoptions, tranportConfigurations...)
	}
}
func CheckShouldUseProxy(upstreamAddress string, tranportConfigurations ...func(*http.Transport) *http.Transport) (*url.URL, error) {
	fmt.Println("开始检查CheckShouldUseProxy", upstreamAddress)
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

func Handle(client net.Conn, httpUpstreamAddress string, proxyoptions options.ProxyOptions, tranportConfigurations ...func(*http.Transport) *http.Transport) {
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
	fmt.Println(string(b[:bytes.IndexByte(b[:], '\n')]))
	// hostPortURL, err := url.Parse(URL)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// 如果方法是 CONNECT，则为 https 协议
	if method == "CONNECT" {
		// hostPortURL, err := url.Parse(URL)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		var line = string(b[:bytes.IndexByte(b[:], '\n')])
		address = ExtractAddressFromConnectRequestLine(line)

		//hostPortURL.Scheme + ":" + hostPortURL.Opaque
	} else { //否则为 http 协议
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
		address, err = ExtractAddressFromOtherRequestLine(line)
		if err != nil {
			fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
			log.Println(err)
			return
		}
	}
	fmt.Println("address:" + address)
	//获得了请求的 host 和 port，向服务端发起 tcp 连接

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
				fmt.Printf("WebSocket ForwardData error: %v\n", err)
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
		// fmt.Println("upstreamAddress:" + httpUpstreamAddress)
		server, err = options.Proxy_net_Dial("tcp", upstreamAddress, proxyoptions, tranportConfigurations...) //net.Dial("tcp", upstreamAddress)

		//	for _, err := range errors {
		//		if err != nil {
		//			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
		//			log.Println(err)
		//			return
		//		}
		//	}
		if err != nil {
			log.Println(err)
			fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
			return
		}
		log.Println("连接成功：" + upstreamAddress)
	}
	//如果使用 https 协议，需先向客户端表示连接建立完毕
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		var requestLine = string(b[:bytes.IndexByte(b[:], '\n')+1])
		//如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端
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
		var headers map[string]string = map[string]string{"Forwarded": forwarded}
		shouldReturn := WriteRequestLineAndHeadersWithRequestURI(requestLine, server, n, b, headers)
		if shouldReturn {
			fmt.Fprint(client, "HTTP/1.1 500 Internal Server Error\r\n\r\n")
			return
		}
	}

	//将客户端的请求转发至服务端，将服务端的响应转发给客户端。io.Copy 为阻塞函数，文件描述符不关闭就不停止
	go io.Copy(server, client)
	io.Copy(client, server)
}

// WriteRequestLineAndHeadersWithRequestURI 将请求行和头部信息写入服务器连接
func WriteRequestLineAndHeadersWithRequestURI(requestLine string, server net.Conn, n int, b [10240]byte, headers map[string]string) bool {
	/*有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
	output, err := RemoveURLPartsLeaveMethodRequestURIVersion(requestLine)
	if err != nil {
		log.Println(err)
		return true
	}
	log.Println("simple Handle", "header:")
	for k, v := range headers {
		// fmt.Println("key:", k)
		log.Println("simple Handle", k, ":", v)
	}
	fmt.Println(output)
	server.Write([]byte(output))

	for k, v := range headers {
		server.Write([]byte(k + ": " + v + "\r\n"))
		fmt.Println(string([]byte(k + ": " + v + "\r\n")))
	}
	// server.Write()
	// server.Write([]byte("\r\n"))
	server.Write(b[len(requestLine):n])
	// server.Write(append([]byte(output), b[len(requestLine):n]...))
	fmt.Println(string(b[len(requestLine):n]))
	return false
}

func ExtractAddressFromOtherRequestLine(line string) (string, error) {
	var address string
	domain, port, err := ExtractDomainAndPort(line)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	} else {
		fmt.Printf("Domain: %s, Port: %s\n", domain, port)
		address = domain + ":" + port
	}
	return address, nil
}

func ExtractAddressFromConnectRequestLine(line string) string {

	return line[7+1 : len(line)-9-1]
}
func RemoveURLPartsLeaveMethodRequestURIVersion(requestLine string) (string, error) {
	/* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
	// 正则表达式用于匹配 http(s)://[domain(:port)] 部分
	parts := strings.SplitN(requestLine, " ", 3)
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid http request line")
	}

	// 获取请求目标
	requestTarget := parts[1]

	// 解析URL
	u, err := url.Parse(requestTarget)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}
	var cleanedUrl = parts[0] + " " + u.RequestURI() + " " + parts[2]
	/* "GET / HTTP/1.1" */
	return cleanedUrl, nil
}
func ExtractDomainAndPort(requestLine string) (string, string, error) {
	/* "GET http://speedtest.cn/ HTTP/1.1" */
	// 分割字符串以获取URL部分
	parts := strings.SplitN(requestLine, " ", 3)
	if len(parts) < 3 {
		return "", "", fmt.Errorf("invalid http request line")
	}

	// 获取请求目标
	requestTarget := parts[1]

	// 解析URL
	u, err := url.Parse(requestTarget)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse url: %w", err)
	}

	// 提取域名
	domain := u.Hostname()

	// 提取端口
	port := u.Port()
	if port == "" {
		// 如果端口未指定，则根据协议使用默认端口
		if u.Scheme == "http" {
			port = "80"
		} else if u.Scheme == "https" {
			port = "443"
		}
	}
	if IsIPv6(domain) {
		domain = "[" + domain + "]"
	}

	/* 需要识别ipv6地址 */
	/* Domain: speedtest.cn, Port: 80 */
	return domain, port, nil
}
func IsIPv6(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil && ip.To16() != nil && ip.To4() == nil
}
