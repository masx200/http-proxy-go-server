package http

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"maps"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masx200/http-proxy-go-server/dnscache"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	socks5_websocket_proxy_golang_websocket "github.com/masx200/socks5-websocket-proxy-golang/pkg/websocket"

	// "github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/utils"
)

func startsWithHTTP(s string) bool {
	return strings.HasPrefix(s, "http://")
}

// 辅助函数：将ForwardedBy列表转换为集合（set），用于快速判断重复项
func setFromForwardedBy(forwardedByList []ForwardedBy) map[string]bool {
	set := make(map[string]bool)
	for _, fb := range forwardedByList {
		set[fb.Identifier] = true
	}
	return set
}

type ForwardedBy struct {
	Identifier string
}

// parseForwardedHeader 解析 "Forwarded" HTTP 头部信息，返回一个 ForwardedBy 结构体切片。
// header: 代表被转发的请求的 "Forwarded" 头部字符串。
// 返回值: 一个包含所有转发标识的 ForwardedBy 结构体切片，以及可能发生的错误。
func parseForwardedHeader(header string) ([]ForwardedBy, error) {
	var forwardedByList []ForwardedBy
	parts := strings.Split(header, ", ")

	for _, part := range parts {
		for _, param := range strings.Split(part, ";") {
			param = strings.TrimSpace(param)
			if !strings.HasPrefix(param, "by=") {
				continue
			}

			// 分离 by 参数的值
			value := strings.TrimPrefix(param, "by=")
			// host, port, err := net.SplitHostPort(value)
			// if err != nil {
			// 如果没有端口信息，host 就是整个值
			var host = value
			// port = ""
			// }

			forwardedBy := ForwardedBy{
				Identifier: host,
				// Port:       port,
			}

			// 检查是否重复
			// isDuplicate := false
			// for _, existing := range forwardedByList {
			// 	if existing.Identifier == forwardedBy.Identifier && existing.Port == forwardedBy.Port {
			// 		isDuplicate = true
			// 		break
			// 	}
			// }
			// if !isDuplicate {
			forwardedByList = append(forwardedByList, forwardedBy)
			// }
		}
	}

	return forwardedByList, nil
}
func proxyHandler(w http.ResponseWriter, r *http.Request, LocalAddr string, proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, username, password string, upstreamResolveIPs bool, ipPriority options.IPPriority, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) error {
	log.Println("method:", r.Method)
	log.Println("url:", r.URL)
	log.Println("host:", r.Host)
	log.Println("proxyHandler", "header:")
	/*/* 这里删除除了第一次请求的 Proxy-Authorization  删除代理认证信息 */

	if username != "" && password != "" {
		var Proxy_Authorization = r.Header.Get("Proxy-Authorization")
		if !isAuthenticated(Proxy_Authorization, username, password) {
			var body = "407 Proxy Authentication Required"
			// fmt.Fprint(client, "HTTP/1.1 407 Proxy Authentication Required\r\ncontent-length: "+strconv.Itoa(len(body))+"\r\nProxy-Authenticate: Basic realm=\"Proxy\"\r\n\r\n")
			// fmt.Fprint(client, body)
			w.Header().Set("Proxy-Authenticate", "Basic realm=\"Proxy\"")
			w.Header().Set("content-length", strconv.Itoa(len(body)))
			w.WriteHeader(407)
			w.Write([]byte(body))
			//fmt.Fprintln(w, "407 Proxy Authentication Required")
			log.Println("身份验证失败")
			//w.Close()
			return nil
		}
		log.Println("身份验证成功")
	}

	r.Header.Del("Proxy-Authorization")
	clienthost, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("clienthost:", clienthost)
	log.Println("clientport:", port)
	forwarded := fmt.Sprintf(
		"for=%s;by=%s;host=%s;proto=%s",
		clienthost, // 代理自己的标识或IP地址
		LocalAddr,  // 代理的标识
		r.Host,     // 原始请求的目标主机名
		"http",     // 或者 "https" 根据实际协议
	)
	r.Header.Add("Forwarded", forwarded)
	for k, v := range r.Header {
		// log.Println("key:", k)
		log.Println("proxyHandler", k, ":", strings.Join(v, ","))
	}
	forwardedHeader := strings.Join(r.Header.Values("Forwarded"), ", ")
	log.Println("forwardedHeader:", forwardedHeader)
	forwardedByList, err := parseForwardedHeader(forwardedHeader)
	log.Println("forwardedByList:", forwardedByList)
	if len(forwardedByList) != len(setFromForwardedBy(forwardedByList)) {
		w.WriteHeader(508)
		fmt.Fprintln(w, "508 Loop Detected")
		log.Println("Duplicate 'by' identifiers found in 'Forwarded' header.")
		return nil
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing 'Forwarded' header: %v", err)
		return err
	}
	targetUrl := "http://" + r.Host + r.RequestURI
	/*r.URL可能是http://开头,也可能只有路径  */
	if startsWithHTTP(r.URL.String()) {
		targetUrl = r.URL.String()
	}
	// 这里假设目标服务器都是HTTP的，实际情况可能需要处理HTTPS
	log.Println("targetUrl:", targetUrl)
	// 创建一个使用了代理的客户端
	defer r.Body.Close()
	/* 请求body的问题 */
	// bodyBytes, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// log.Println("body:", string(bodyBytes))
	transport := &http.Transport{
		ForceAttemptHTTP2: true,
		// 自定义 DialContext 函数
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {

			var host, _, err = net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			if utils.IsLoopbackIP(host) {
				var dialer = &net.Dialer{}
				return dialer.DialContext(ctx, network, addr)
			}
			// 解析出原地址中的端口
			hostname, _, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}

			if IsIP(hostname) {
				dialer := &net.Dialer{}
				//				// 发起连接
				return dialer.DialContext(ctx, network, addr)
			}

			// 如果启用了上游IP解析，先解析目标地址
			targetAddr := addr
			if upstreamResolveIPs {
				log.Printf("upstream-resolve-ips enabled, resolving target address %s before connection", targetAddr)

				resolvedAddrs, err := resolveTargetAddressForAuth(targetAddr, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, ipPriority, tranportConfigurations...)
				if err != nil {
					log.Printf("Failed to resolve target address %s: %v, using original", targetAddr, err)
					resolvedAddrs = []string{targetAddr}
				}

				// 使用轮询从解析的地址中选择一个
				resolvedAddr := resolveTargetAddressForAuthWithRoundRobin(resolvedAddrs, targetAddr, ipPriority)

				if upstreamResolveIPs && resolvedAddr != targetAddr {
					log.Printf("Using resolved address %s instead of original %s", resolvedAddr, targetAddr)
					targetAddr = resolvedAddr
				}
			}

			return dnscache.Proxy_net_DialContextCached(ctx, network, targetAddr, proxyoptions, dnsCache, upstreamResolveIPs, Proxy, tranportConfigurations...)
		},
		DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {

			//				// 解析出原地址中的端口
			hostname, _, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}

			// 如果启用了上游IP解析，先解析目标地址
			targetAddr := addr
			if upstreamResolveIPs {
				log.Printf("upstream-resolve-ips enabled, resolving TLS target address %s before connection", targetAddr)

				resolvedAddrs, err := resolveTargetAddressForAuth(targetAddr, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, ipPriority, tranportConfigurations...)
				if err != nil {
					log.Printf("Failed to resolve TLS target address %s: %v, using original", targetAddr, err)
					resolvedAddrs = []string{targetAddr}
				}

				// 使用轮询从解析的地址中选择一个
				resolvedAddr := resolveTargetAddressForAuthWithRoundRobin(resolvedAddrs, targetAddr, ipPriority)

				if upstreamResolveIPs && resolvedAddr != targetAddr {
					log.Printf("Using resolved TLS address %s instead of original %s", resolvedAddr, targetAddr)
					targetAddr = resolvedAddr

					// 如果地址被解析为IP，需要更新hostname用于TLS配置
					resolvedHostname, _, err := net.SplitHostPort(resolvedAddr)
					if err == nil {
						// 只有当解析后的hostname不是IP时才更新ServerName
						// 如果是IP地址，保持原来的hostname作为SNI
						if net.ParseIP(resolvedHostname) == nil {
							hostname = resolvedHostname
						}
					}
				}
			}

			//				// 用指定的 IP 地址和原端口创建新地址
			//				newAddr := net.JoinHostPort(serverIP, port)
			//				// 创建 net.Dialer 实例
			//				dialer := &net.Dialer{}
			//				// 发起连接
			conn, err := dnscache.Proxy_net_DialContextCached(ctx, network, targetAddr, proxyoptions, dnsCache, upstreamResolveIPs, Proxy, tranportConfigurations...) //dialer.DialContext(ctx, network, newAddr)
			if err != nil {
				return nil, err
			}
			//			var address = addr
			tlsConfig := &tls.Config{
				ServerName: hostname,
			}
			// 创建 TLS 连接
			tlsConn := tls.Client(conn, tlsConfig)
			// 进行 TLS 握手
			err = tlsConn.HandshakeContext(ctx)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return tlsConn, nil
		},
	}
	client := &http.Client{ /* Transport: newTransport("http://your_proxy_address:port") */
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse /* 不进入重定向 */
		},

		/* Jar: jar */} // 替换为你的代理服务器地址和端口
	for _, f := range tranportConfigurations {
		transport = f(transport)
	}
	if len(proxyoptions) > 0 {

		client.Transport = transport
	}
	/* 流式处理,防止内存溢出 */
	proxyReq, err := http.NewRequest(r.Method, targetUrl, r.Body /* bytes.NewReader(bodyBytes) */)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	proxyUrl, err := utils.CheckShouldUseProxy(proxyReq.Host, Proxy, tranportConfigurations...)
	if err != nil {
		log.Println(err)
		return err
	}
	if proxyUrl != nil && (proxyUrl.Scheme == "ws" || proxyUrl.Scheme == "wss") {

		log.Println("使用代理：" + proxyUrl.String())

		if client.Transport == nil {
			client.Transport = http.DefaultTransport
		}

		if transport, ok := client.Transport.(*http.Transport); ok {
			transport.Proxy = nil

			log.Println("已经修改了代理为websocket", proxyUrl.String())
			var DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
				var host, _, err = net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}
				if utils.IsLoopbackIP(host) {
					var dialer = &net.Dialer{}
					return dialer.DialContext(ctx, network, addr)
				}
				log.Println("使用代理：" + proxyUrl.String())

				log.Println("network,addr", network, addr)
				return websocketDialContext(ctx, network, addr, proxyUrl, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, ipPriority)
			}
			transport.DialContext = DialContext
			transport.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {

				//				// 解析出原地址中的端口
				hostname, _, err := net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}
				//				// 用指定的 IP 地址和原端口创建新地址
				//				newAddr := net.JoinHostPort(serverIP, port)
				//				// 创建 net.Dialer 实例
				//				dialer := &net.Dialer{}
				//				// 发起连接
				conn, err := DialContext(ctx, network, addr)
				if err != nil {
					return nil, err
				}
				//			var address = addr
				tlsConfig := &tls.Config{
					ServerName: hostname,
				}
				// 创建 TLS 连接
				tlsConn := tls.Client(conn, tlsConfig)
				// 进行 TLS 握手
				err = tlsConn.HandshakeContext(ctx)
				if err != nil {
					conn.Close()
					return nil, err
				}
				return tlsConn, nil
			}
		}
	}

	proxyReq.Header = r.Header.Clone()
	proxyReq.ContentLength = r.ContentLength
	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return err
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	// Copy headers from the response to the client's response.
	maps.Copy(w.Header(), resp.Header)

	// Copy the response body back to the client.
	/* bodyBytes2, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} */
	/* 流式处理,防止内存溢出 */
	if _, err := io.Copy(w, resp.Body /* bytes.NewReader(bodyBytes2) */); err != nil {
		log.Println("Error writing response:", err)
	}
	return nil
}

func Http(hostname string, port int, proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, username, password string, upstreamResolveIPs bool, ipPriority options.IPPriority, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	/* /* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
	// 监听本地8080端口
	listener, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Proxy server started on port %s", listener.Addr())
	var LocalAddr = listener.Addr().String()
	engine.Use(func(c *gin.Context) {
		var w = c.Writer
		var r = c.Request
		err := proxyHandler(w, r /* jar, */, LocalAddr, proxyoptions, dnsCache, username, password, upstreamResolveIPs, ipPriority, Proxy, tranportConfigurations...)
		c.Abort()

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	// 设置自定义处理器
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//engine.Handler().ServeHTTP(w, r)
	//})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		engine.Handler().ServeHTTP(w, r)
	})
	// 开始服务
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
func GenerateRandomLoopbackIP() string {
	// Check if running on Windows
	if runtime.GOOS == "windows" {
		// For Windows compatibility, use 127.0.0.1 instead of random 127.x.x.x
		// Windows networking may not properly handle all 127.x.x.x addresses
		ip := net.IPv4(127, 0, 0, 1)
		log.Println("Using Windows-compatible loopback IP:", ip.String())
		return ip.String()
	} else {
		// For Linux/macOS, use the original random IP generation
		randomIP := generateRandomIP()
		log.Println("Random IP:", randomIP)
		return randomIP.String()
	}
}

func generateRandomIP() net.IP {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	ip := net.IPv4(
		byte(127 /* +rand.Intn(1) */),
		byte(rd.Intn(256)),
		byte(rd.Intn(256)),
		byte(rd.Intn(256)),
	)
	return ip
}

func GenerateRandomIntPort() int {
	randomInt := generateRandomInt()
	log.Println("Random integer:", randomInt)
	return randomInt
}

func generateRandomInt() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	minport := 10000
	maxport := 65535
	return rd.Intn(maxport-minport+1) + minport
}

// IsIP 判断给定的字符串是否是有效的 IPv4 或 IPv6 地址。
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
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
func websocketDialContext(ctx context.Context, network, addr string, proxyUrl *url.URL, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, ipPriority options.IPPriority) (net.Conn, error) {
	// 解析目标地址
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		// 如果没有端口，尝试添加默认端口
		if network == "tcp" {
			if strings.Contains(addr, ":") {
				// IPv6地址
				addr = "[" + addr + "]:80"
			} else {
				// 域名或IPv4地址
				addr = addr + ":80"
			}
			host, port, err = net.SplitHostPort(addr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse address %s: %v", addr, err)
			}
		} else {
			return nil, fmt.Errorf("failed to parse address %s: %v", addr, err)
		}
	}

	// 转换端口号为整数
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to parse port %s: %v", port, err)
	}

	// 创建WebSocket客户端配置
	wsConfig := interfaces.ClientConfig{
		Username:   proxyUrl.User.Username(),
		Password:   "",
		ServerAddr: proxyUrl.String(),
		Protocol:   "websocket",
		Timeout:    30 * time.Second,
	}
	if proxyUrl.User != nil {
		wsConfig.Username = proxyUrl.User.Username()
		wsConfig.Password, _ = proxyUrl.User.Password()
	}

	// 详细打印wsConfig的每个字段
	log.Println("WebSocket Config Details:")
	log.Println("host, portNum", host, portNum)
	log.Printf("  Username: %s", wsConfig.Username)
	log.Printf("  Password: %s", wsConfig.Password)
	log.Printf("  ServerAddr: %s", wsConfig.ServerAddr)
	log.Printf("  Protocol: %s", wsConfig.Protocol)
	log.Printf("  Timeout: %v", wsConfig.Timeout)

	// 创建WebSocket客户端
	websocketClient := socks5_websocket_proxy_golang_websocket.NewWebSocketClient(wsConfig)
	log.Println("host, portNum", host, portNum)
	// 连接到目标主机
	// 如果启用了DNS解析，先解析目标地址
	targetAddr := net.JoinHostPort(host, strconv.Itoa(portNum))

	if upstreamResolveIPs {
		log.Printf("upstream-resolve-ips enabled, resolving WebSocket target address %s before connection", targetAddr)
	}

	resolvedAddrs, err := resolveTargetAddressForAuth(targetAddr, Proxy, proxyoptions, dnsCache, upstreamResolveIPs, ipPriority)

	if err != nil {
		log.Printf("Failed to resolve WebSocket target address %s: %v, using original", targetAddr, err)

		resolvedAddrs = []string{targetAddr}
	}

	// 使用轮询从解析的地址中选择一个
	//
	resolvedAddr := resolveTargetAddressForAuthWithRoundRobin(resolvedAddrs, targetAddr, ipPriority)

	if upstreamResolveIPs && resolvedAddr != targetAddr {

		log.Printf("WebSocket: Using resolved address %s instead of original %s", resolvedAddr, targetAddr)

		// 重新解析解析后的地址以获取正确的host和port用于连接
		//
		//
		resolvedHost, resolvedPort, err := net.SplitHostPort(resolvedAddr)

		if err != nil {
			return nil, fmt.Errorf("failed to parse resolved address %s: %v", resolvedAddr, err)
		}

		host = resolvedHost

		portNum, err = strconv.Atoi(resolvedPort)

		if err != nil {
			return nil, fmt.Errorf("failed to parse resolved port %s: %v", resolvedPort, err)
		}
	}
	err = websocketClient.Connect(host, portNum)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s:%d via WebSocket proxy: %v", host, portNum, err)
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

	// 返回客户端连接
	return clientConn, nil
}

// resolveTargetAddressForAuth 解析目标地址的域名为IP地址（用于auth模块）
func resolveTargetAddressForAuth(addr string, Proxy func(*http.Request) (*url.URL, error), proxyoptions options.ProxyOptionsDNSSLICE, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, ipPriority options.IPPriority, transportConfigurations ...func(*http.Transport) *http.Transport) ([]string, error) {
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

	log.Printf("Resolving target address %s using DoH infrastructure", host)

	// 使用DoH解析
	resolver := dnscache.CreateHostsAndDohResolverCachedSimple(proxyoptions, dnsCache, Proxy, transportConfigurations...)
	ips, err := resolver.LookupIP(context.Background(), "tcp", host)
	if err != nil {
		log.Printf("DoH resolution failed for target %s: %v", host, err)
		return []string{addr}, err
	}

	if len(ips) == 0 {
		log.Printf("No IP addresses resolved for target %s", host)
		return []string{addr}, fmt.Errorf("no IP addresses resolved for target %s", host)
	}

	// 收集所有解析出的IP地址
	var resolvedAddrs []string
	for _, ip := range ips {
		resolvedAddr := net.JoinHostPort(ip.String(), port)
		resolvedAddrs = append(resolvedAddrs, resolvedAddr)
	}

	// 根据 IP 优先级策略排序
	resolvedAddrs = options.SortAddressesByPriority(resolvedAddrs, ipPriority)

	// 统计 IPv4 和 IPv6 地址数量
	var ipv4Count, ipv6Count int
	for _, addr := range resolvedAddrs {
		host, _, _ := net.SplitHostPort(addr)
		if ip := net.ParseIP(host); ip != nil {
			if ip.To4() != nil {
				ipv4Count++
			} else {
				ipv6Count++
			}
		}
	}

	log.Printf("Resolved target address %s to %d IP addresses (IPv4: %d, IPv6: %d, priority: %s): %v",
		addr, len(resolvedAddrs), ipv4Count, ipv6Count, ipPriority, resolvedAddrs)

	return resolvedAddrs, nil
}

// resolveTargetAddressForAuthWithRoundRobin 从解析的IP数组中根据优先级策略选择一个地址（auth模块使用）
func resolveTargetAddressForAuthWithRoundRobin(addrs []string, target string, ipPriority options.IPPriority) string {
	if len(addrs) == 0 {
		return target
	}

	if len(addrs) == 1 {
		return addrs[0]
	}

	var ipv4Addrs []string
	var ipv6Addrs []string

	// 分离 IPv4 和 IPv6 地址
	for _, addr := range addrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			// 如果无法解析，作为 IPv4 处理
			ipv4Addrs = append(ipv4Addrs, addr)
			continue
		}
		if ip := net.ParseIP(host); ip != nil {
			if ip.To4() != nil {
				ipv4Addrs = append(ipv4Addrs, addr)
			} else {
				ipv6Addrs = append(ipv6Addrs, addr)
			}
		} else {
			// 如果不是 IP 地址，作为 IPv4 处理
			ipv4Addrs = append(ipv4Addrs, addr)
		}
	}

	var candidateAddrs []string

	// 根据优先级策略选择候选地址列表
	switch ipPriority {
	case options.IPv6Priority:
		if len(ipv6Addrs) > 0 {
			candidateAddrs = ipv6Addrs
		} else {
			candidateAddrs = ipv4Addrs
		}
		log.Printf("IPv6 priority: selecting from %d IPv6 addresses", len(ipv6Addrs))
	case options.IPv4Priority:
		if len(ipv4Addrs) > 0 {
			candidateAddrs = ipv4Addrs
		} else {
			candidateAddrs = ipv6Addrs
		}
		log.Printf("IPv4 priority: selecting from %d IPv4 addresses", len(ipv4Addrs))
	case options.IPRandomPriority:
		candidateAddrs = addrs
		log.Printf("Random priority: selecting from all %d addresses", len(addrs))
	default:
		// 默认 IPv4 优先
		if len(ipv4Addrs) > 0 {
			candidateAddrs = ipv4Addrs
		} else {
			candidateAddrs = ipv6Addrs
		}
	}

	// 如果候选地址为空，使用所有地址
	if len(candidateAddrs) == 0 {
		candidateAddrs = addrs
	}

	// 从候选地址中随机选择一个
	candidateAddrs = shuffleAuth(candidateAddrs)
	selectedAddr := candidateAddrs[rand.Intn(len(candidateAddrs))]

	log.Printf("SOCKS5 RoundRobin (priority=%s) selected address %s from %v for target %s", ipPriority, selectedAddr, addrs, target)

	return selectedAddr
}

// shuffleAuth 对切片进行随机排序（auth模块使用）
func shuffleAuth[T any](slice []T) []T {
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为随机种子
	rand1.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
