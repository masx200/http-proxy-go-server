package tls_auth

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/masx200/http-proxy-go-server/auth"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"github.com/masx200/http-proxy-go-server/options"
)

func Tls_auth(server_cert string, server_key, hostname string, port int, username, password string, proxyoptions options.ProxyOptions) {

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
	go http_server.Http(xh, x1, proxyoptions, username, password)
	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		// go handle(client, username, password)
		go auth.Handle(client, username, password, upstreamAddress, proxyoptions)
	}
}

// func Main() {

// 	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	config := &tls.Config{Certificates: []tls.Certificate{cert}}
// 	ln, err := tls.Listen("tcp", ":443", config)
// 	// tcp 连接，监听 8080 端口
// 	// l, err := net.Listen("tcp", ":8080")
// 	var l = ln
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// 死循环，每当遇到连接时，调用 handle
// 	for {
// 		client, err := l.Accept()
// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		go auth.Handle(client, username, password)
// 	}
// } /*
// func handle(client net.Conn, username, password string) {
// 	if client == nil {
// 		return
// 	}
// 	defer client.Close()

// 	log.Printf("remote addr: %v\n", client.RemoteAddr())

// 	// 用来存放客户端数据的缓冲区
// 	var b [10240]byte
// 	//从客户端获取数据
// 	n, err := client.Read(b[:])
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	var method, URL, address string
// 	// 从客户端数据读入 method，url
// 	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &URL)
// 	hostPortURL, err := url.Parse(URL)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// 检查 Proxy-Authorization 头
// 	proxyAuth := ""
// 	for _, line := range strings.Split(string(b[:n]), "\n") {
// 		if strings.HasPrefix(line, "Proxy-Authorization:") {
// 			proxyAuth = strings.TrimSpace(strings.TrimPrefix(line, "Proxy-Authorization:"))
// 			break
// 		}
// 	}

// 	// 验证身份
// 	if !isAuthenticated(proxyAuth, username, password) {
// 		fmt.Fprint(client, "HTTP/1.1 407 Proxy Authentication Required\r\nProxy-Authenticate: Basic realm=\"Proxy\"\r\n\r\n")
// 		fmt.Println("身份验证失败")
// 		return
// 	}
// 	fmt.Println("身份验证成功")
// 	// 如果方法是 CONNECT，则为 https 协议
// 	if method == "CONNECT" {
// 		address = hostPortURL.Scheme + ":" + hostPortURL.Opaque
// 	} else { //否则为 http 协议
// 		address = hostPortURL.Host
// 		// 如果 host 不带端口，则默认为 80
// 		if !strings.Contains(hostPortURL.Host, ":") { //host 不带端口， 默认 80
// 			address = hostPortURL.Host + ":80"
// 		}
// 	}
// 	fmt.Println("address:" + address)
// 	//获得了请求的 host 和 port，向服务端发起 tcp 连接
// 	server, err := net.Dial("tcp", address)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	//如果使用 https 协议，需先向客户端表示连接建立完毕
// 	if method == "CONNECT" {
// 		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
// 	} else { //如果使用 http 协议，需将从客户端得到的 http 请求转发给服务端

// 		req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(b[:n])))
// 		if err != nil {
// 			log.Println("Error parsing request:", err)
// 			return
// 		}
// 		req.Header.Del("Proxy-Authorization")
// 		// server.Write(b[:n])
// 		err = req.Write(server)
// 		if err != nil {
// 			log.Println("Error writing request to server:", err)
// 			return
// 		}
// 	}

// 	//将客户端的请求转发至服务端，将服务端的响应转发给客户端。io.Copy 为阻塞函数，文件描述符不关闭就不停止
// 	go io.Copy(server, client)
// 	io.Copy(client, server)
// }

// func isAuthenticated(proxyAuth, expectedUsername, expectedPassword string) bool {
// 	if !strings.HasPrefix(proxyAuth, "Basic ") {
// 		return false
// 	}

// 	auth := strings.TrimPrefix(proxyAuth, "Basic ")
// 	decodedAuth, err := base64.StdEncoding.DecodeString(auth)
// 	if err != nil {
// 		return false
// 	}

// 	username, password, ok := strings.Cut(string(decodedAuth), ":")
// 	if !ok {
// 		return false
// 	}

// 	return username == expectedUsername && password == expectedPassword
// }
// */
