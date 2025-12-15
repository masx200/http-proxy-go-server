package simple

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	// "regexp"

	"github.com/masx200/http-proxy-go-server/connect"
	"github.com/masx200/http-proxy-go-server/dnscache"
	http_server "github.com/masx200/http-proxy-go-server/http"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/utils"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/socks5"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"
)

func Simple(hostname string, port int, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, tranportConfigurations ...func(*http.Transport) *http.Transport) {
	// tcp 连接，监听 8080 端口
	l, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Proxy server started on port %s", l.Addr())

	// 检查是否使用SOCKS5上游代理
	var useSocks5Directly bool
	var upstreamAddress string

	if Proxy != nil {
		// 创建一个测试请求来检查上游代理类型
		testReq, _ := http.NewRequest("GET", "http://test", nil)
		if proxyURL, err := Proxy(testReq); err == nil && proxyURL != nil {
			useSocks5Directly = strings.HasPrefix(proxyURL.String(), "socks5://")
			if useSocks5Directly {
				log.Printf("SOCKS5 upstream detected, will handle HTTP requests directly via SOCKS5")
			}
		}
	}

	// 只有在非SOCKS5上游时才启动HTTP代理服务器
	if !useSocks5Directly {
		xh := http_server.GenerateRandomLoopbackIP()
		x1 := http_server.GenerateRandomIntPort()
		upstreamAddress = xh + ":" + fmt.Sprint(x1)
		go http_server.Http(xh, x1, proxyoptions, dnsCache, "", "", upstreamResolveIPs, Proxy, tranportConfigurations...)
		log.Printf("Started HTTP proxy server for upstream routing at %s", upstreamAddress)
	} else {
		log.Printf("SOCKS5 upstream mode: bypassing HTTP proxy server for direct SOCKS5 routing")
	}

	// 死循环，每当遇到连接时，调用 handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
			return
		}

		go Handle(client, upstreamAddress, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, tranportConfigurations...)
	}
}
func CheckShouldUseProxy(upstreamAddress string, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) (*url.URL, error) {
	return utils.CheckShouldUseProxy(upstreamAddress, Proxy, tranportConfigurations...)
}

func Handle(client net.Conn, httpUpstreamAddress string, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, tranportConfigurations ...func(*http.Transport) *http.Transport) {
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
	newVar := bytes.IndexByte(b[:], '\n')

	if newVar == -1 {
		fmt.Fprint(client, "HTTP/1.1 400 Bad Request\r\n\r\n")
		log.Println("400 Bad Request,not http request")
		return
	}
	fmt.Sscanf(string(b[:newVar]), "%s%s", &method, &URL)
	log.Println(string(b[:bytes.IndexByte(b[:], '\n')]))
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
	log.Println("address:" + address)
	//获得了请求的 host 和 port，向服务端发起 tcp 连接

	var upstreamAddress string
	if method == "CONNECT" {
		upstreamAddress = address
	} else {
		// 对于HTTP请求，检查是否使用SOCKS5直接连接模式
		if httpUpstreamAddress == "" {
			// SOCKS5直接模式：直接连接到目标地址
			upstreamAddress = address
			log.Printf("HTTP request will be routed directly via SOCKS5 to: %s", upstreamAddress)
		} else {
			// 传统模式：通过HTTP代理服务器
			upstreamAddress = httpUpstreamAddress
			log.Printf("HTTP request will be routed via HTTP proxy server at: %s", upstreamAddress)
		}
	}
	var server net.Conn
	proxyURL, err := CheckShouldUseProxy(upstreamAddress, Proxy, tranportConfigurations...)

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
		log.Println("WebSocket Config Details:")
		log.Println("host, portNum", host, portNum)
		log.Printf("  Username: %s", wsConfig.Username)
		log.Printf("  Password: %s", wsConfig.Password)
		log.Printf("  ServerAddr: %s", wsConfig.ServerAddr)
		log.Printf("  Protocol: %s", wsConfig.Protocol)
		log.Printf("  Timeout: %v", wsConfig.Timeout)
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
	} else if proxyURL != nil && (method == "CONNECT" || (method != "CONNECT" && httpUpstreamAddress == "")) {
		// 检查是否是SOCKS5代理 (适用于CONNECT请求和SOCKS5直接模式的HTTP请求)
		if strings.HasPrefix(proxyURL.String(), "socks5://") {
			// 使用SOCKS5代理处理请求（CONNECT请求或SOCKS5直接模式的HTTP请求）
			requestType := "CONNECT"
			if method != "CONNECT" {
				requestType = "HTTP"
			}
			log.Printf("Processing %s request via SOCKS5 proxy to: %s", requestType, upstreamAddress)
			host, port, err := net.SplitHostPort(upstreamAddress)
			if err != nil {
				log.Println("failed to parse address:", err)
				fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
				return
			}

			// 转换端口号为整数
			portNum, err := strconv.Atoi(port)
			if err != nil {
				log.Println("failed to parse port:", err)
				fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
				return
			}

			// 从代理URL中提取主机和端口
			proxyHost := proxyURL.Hostname()
			proxyPort := proxyURL.Port()
			if proxyPort == "" {
				proxyPort = "1080" // SOCKS5默认端口
			}

			// 构建SOCKS5服务器地址（不包含用户名密码）
			socks5ServerAddr := fmt.Sprintf("tcp://%s:%s", proxyHost, proxyPort)

			// 创建SOCKS5客户端配置
			socks5Config := interfaces.ClientConfig{
				Username:   proxyURL.User.Username(),
				Password:   "",
				ServerAddr: socks5ServerAddr,
				Protocol:   "socks5",
				Timeout:    30 * time.Second,
			}
			if proxyURL.User != nil {
				if password, ok := proxyURL.User.Password(); ok {
					socks5Config.Password = password
				}
			}

			// 创建SOCKS5客户端
			socks5Client := socks5.NewSOCKS5Client(socks5Config)
			log.Println("SOCKS5 Config Details:")
			log.Println("host, portNum", host, portNum)
			log.Printf("  Username: %s", socks5Config.Username)
			log.Printf("  Password: %s", socks5Config.Password)
			log.Printf("  ServerAddr: %s", socks5Config.ServerAddr)
			log.Printf("  Protocol: %s", socks5Config.Protocol)
			log.Printf("  Timeout: %v", socks5Config.Timeout)

			// 如果启用了DNS解析，先解析目标地址
			targetAddr := net.JoinHostPort(host, strconv.Itoa(portNum))
			if upstreamResolveIPs {
				log.Printf("upstream-resolve-ips enabled, resolving SOCKS5 target address %s before connection", targetAddr)
			}
			resolvedAddrs, err := resolveTargetAddressForSimple(targetAddr, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, tranportConfigurations...)
			if err != nil {
				log.Printf("Failed to resolve SOCKS5 target address %s: %v, using original", targetAddr, err)
				resolvedAddrs = []string{targetAddr}
			}

			// 使用轮询从解析的地址中选择一个
			resolvedAddr := resolveTargetAddressForSimpleWithRoundRobin(resolvedAddrs, targetAddr)
			if upstreamResolveIPs && resolvedAddr != targetAddr {
				log.Printf("SOCKS5: Using resolved address %s instead of original %s", resolvedAddr, targetAddr)
				// 重新解析解析后的地址以获取正确的host和port用于连接
				resolvedHost, resolvedPort, err := net.SplitHostPort(resolvedAddr)
				if err != nil {
					log.Println("failed to parse resolved address:", err)
					fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
					return
				}
				host = resolvedHost
				portNum, err = strconv.Atoi(resolvedPort)
				if err != nil {
					log.Println("failed to parse resolved port:", err)
					fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
					return
				}
			}

			// 连接到目标主机
			err = socks5Client.Connect(host, portNum)
			if err != nil {
				log.Println("failed to connect via SOCKS5 proxy:", err)
				fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
				return
			}

			// 创建一个管道连接来处理SOCKS5数据转发
			clientConn, serverConn := net.Pipe()

			// 在goroutine中处理SOCKS5数据转发
			go func() {
				defer clientConn.Close()
				defer serverConn.Close()
				// 使用ForwardData方法处理SOCKS5连接
				err := socks5Client.ForwardData(serverConn)
				if err != nil {
					log.Printf("SOCKS5 ForwardData error: %v\n", err)
				}
			}()

			server = clientConn
			log.Printf("SOCKS5代理连接成功 (%s请求): %s", requestType, upstreamAddress)
		} else {
			// 使用HTTP代理处理CONNECT请求
			server, err = connect.ConnectViaHttpProxy(proxyURL, upstreamAddress, Proxy, proxyoptions, dnsCache, upstreamResolveIPs)
			if err != nil {
				log.Println(err)
				fmt.Fprint(client, "HTTP/1.1 502 Bad Gateway\r\n\r\n")
				return
			}
			log.Println("连接成功：" + upstreamAddress)
		}
	} else {
		// log.Println("upstreamAddress:" + httpUpstreamAddress)
		server, err = dnscache.Proxy_net_DialCached("tcp", upstreamAddress, proxyoptions, upstreamResolveIPs, dnsCache, Proxy, tranportConfigurations...) //net.Dial("tcp", upstreamAddress)

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
		// log.Println("key:", k)
		log.Println("simple Handle", k, ":", v)
	}
	log.Println(output)
	server.Write([]byte(output))

	for k, v := range headers {
		server.Write([]byte(k + ": " + v + "\r\n"))
		log.Println(string([]byte(k + ": " + v + "\r\n")))
	}
	// server.Write()
	// server.Write([]byte("\r\n"))
	server.Write(b[len(requestLine):n])
	// server.Write(append([]byte(output), b[len(requestLine):n]...))
	log.Println(string(b[len(requestLine):n]))
	return false
}

func ExtractAddressFromOtherRequestLine(line string) (string, error) {
	var address string
	domain, port, err := ExtractDomainAndPort(line)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	} else {
		log.Printf("Domain: %s, Port: %s\n", domain, port)
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

// resolveTargetAddressForSimple 解析目标地址的域名为IP地址（用于simple模块）
func resolveTargetAddressForSimple(addr string, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, transportConfigurations ...func(*http.Transport) *http.Transport) ([]string, error) {
	if !upstreamResolveIPs || len(proxyoptions) == 0 || dnsCache == nil {
		return []string{addr}, nil
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return []string{addr}, err
	}

	// 如果已经是IP地址，直接返回
	if net.ParseIP(host) != nil {
		return []string{addr}, nil
	}

	log.Printf("Resolving SOCKS5 target address %s using DoH infrastructure", host)

	// 使用DoH解析
	resolver := dnscache.CreateHostsAndDohResolverCachedSimple(proxyoptions, dnsCache, Proxy, transportConfigurations...)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for SOCKS5 target %s: %v", host, err)
		return []string{addr}, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for SOCKS5 target %s", host)
		return []string{addr}, fmt.Errorf("no IP addresses resolved for SOCKS5 target %s", host)
	}

	// 返回所有解析出的IP地址，优先使用IPv4
	var resolvedAddrs []string
	var ipv4Addrs []string
	var ipv6Addrs []string

	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		if ip.To4() != nil {
			// IPv4地址优先添加
			ipv4Addrs = append(ipv4Addrs, resolvedAddr)
		} else {
			// IPv6地址后添加
			ipv6Addrs = append(ipv6Addrs, resolvedAddr)
		}
	}

	// IPv4地址在前，IPv6地址在后
	resolvedAddrs = append(ipv4Addrs, ipv6Addrs...)

	log.Printf("Resolved SOCKS5 target address %s to %d IP addresses (IPv4: %d, IPv6: %d): %v",
		addr, len(resolvedAddrs), len(ipv4Addrs), len(ipv6Addrs), resolvedAddrs)

	return resolvedAddrs, nil
}

// resolveTargetAddressForSimpleWithRoundRobin 从解析的IP数组中轮询选择一个地址（simple模块使用）
func resolveTargetAddressForSimpleWithRoundRobin(addrs []string, target string) string {
	if len(addrs) == 0 {
		return target
	}

	if len(addrs) == 1 {
		return addrs[0]
	}

	// 分离IPv4和IPv6地址
	var ipv4Addrs []string
	var ipv6Addrs []string

	for _, addr := range addrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			continue
		}
		if ip := net.ParseIP(host); ip != nil && ip.To4() != nil {
			ipv4Addrs = append(ipv4Addrs, addr)
		} else {
			ipv6Addrs = append(ipv6Addrs, addr)
		}
	}

	// 优先选择IPv4地址，提高SOCKS5兼容性
	var candidateAddrs []string
	if len(ipv4Addrs) > 0 {
		candidateAddrs = ipv4Addrs
		log.Printf("SOCKS5 Preferring IPv4 addresses for compatibility: %v", ipv4Addrs)
	} else {
		candidateAddrs = ipv6Addrs
		log.Printf("SOCKS5 No IPv4 addresses available, using IPv6: %v", ipv6Addrs)
	}

	candidateAddrs = shuffleSimple(candidateAddrs)

	// 简单轮询：基于目标字符串哈希来选择一个相对稳定的IP
	hash := 0
	for _, c := range target {
		hash = (hash*31 + int(c)) % len(candidateAddrs)
	}

	selectedAddr := candidateAddrs[hash]
	log.Printf("SOCKS5 RoundRobin selected address %s from %v for target %s", selectedAddr, candidateAddrs, target)

	return selectedAddr
}

// shuffleSimple 对切片进行随机排序（simple模块使用）
func shuffleSimple[T any](slice []T) []T {
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
