package simple

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	// "regexp"
	"strings"

	http_server "github.com/masx200/http-proxy-go-server/http"
)

func Simple(hostname string, port int) {
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

		go Handle(client, upstreamAddress)
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

// 		go Handle(client)
// 	}
// }

func Handle(client net.Conn, httpUpstreamAddress string) {
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
			log.Println(err)
			return
		}
	}
	fmt.Println("address:" + address)
	//获得了请求的 host 和 port，向服务端发起 tcp 连接
	fmt.Println("upstreamAddress:" + httpUpstreamAddress)
	var upstreamAddress string
	if method == "CONNECT" {
		upstreamAddress = address
	} else {
		upstreamAddress = httpUpstreamAddress
	}
	server, err := net.Dial("tcp", upstreamAddress)
	if err != nil {
		log.Println(err)
		fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
		return
	}
	//如果使用 https 协议，需先向客户端表示连接建立完毕
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		var requestLine = string(b[:bytes.IndexByte(b[:], '\n')+1])
		//如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端
		clienthost, port, err := net.SplitHostPort(client.RemoteAddr().String())
		if err != nil {
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
