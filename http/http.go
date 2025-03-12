package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masx200/http-proxy-go-server/options"
)

// "bytes"
// "bytes"

// "net/http/cookiejar"

// "net/url"

// "github.com/go-kit/kit/sd/etcd"

// Create a custom transport that uses the proxy for HTTP requests.
// func newTransport(proxyAddress string) *http.Transport {
// 	proxyURL, _ := url.Parse(proxyAddress) // 注意处理错误
// 	return &http.Transport{
// 		Proxy: http.ProxyURL(proxyURL),
// 	}
// }

// ServeHTTP is a handler that forwards incoming requests to the target URL specified in the request.

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
func proxyHandler(w http.ResponseWriter, r *http.Request /*  jar *cookiejar.Jar, */, LocalAddr string, proxyoptions options.ProxyOptions) {
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("host:", r.Host)
	log.Println("proxyHandler", "header:")
	/*/* 这里删除除了第一次请求的 Proxy-Authorization  删除代理认证信息 */

	r.Header.Del("Proxy-Authorization")
	clienthost, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
		return
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
		// fmt.Println("key:", k)
		log.Println("proxyHandler", k, ":", strings.Join(v, ","))
	}
	forwardedHeader := strings.Join(r.Header.Values("Forwarded"), ", ")
	log.Println("forwardedHeader:", forwardedHeader)
	forwardedByList, err := parseForwardedHeader(forwardedHeader)
	log.Println("forwardedByList:", forwardedByList)
	if len(forwardedByList) != len(setFromForwardedBy(forwardedByList)) {
		w.WriteHeader(508)
		fmt.Fprintln(w, "Duplicate 'by' identifiers found in 'Forwarded' header.")
		log.Println("Duplicate 'by' identifiers found in 'Forwarded' header.")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing 'Forwarded' header: %v", err)
		return
	}
	targetUrl := "http://" + r.Host + r.RequestURI
	/*r.URL可能是http://开头,也可能只有路径  */
	if startsWithHTTP(r.URL.String()) {
		targetUrl = r.URL.String()
	}
	// 这里假设目标服务器都是HTTP的，实际情况可能需要处理HTTPS
	fmt.Println("targetUrl:", targetUrl)
	// 创建一个使用了代理的客户端
	defer r.Body.Close()
	/* 请求body的问题 */
	// bodyBytes, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Println("body:", string(bodyBytes))
	transport := &http.Transport{
		ForceAttemptHTTP2: true,
		// 自定义 DialContext 函数
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
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
			//				// 用指定的 IP 地址和原端口创建新地址
			//				newAddr := net.JoinHostPort(serverIP, port)
			//				// 创建 net.Dialer 实例
			//				dialer := &net.Dialer{}
			//				// 发起连接
			//				return dialer.DialContext(ctx, network, newAddr)

			return options.Proxy_net_DialContext(ctx, network, addr, proxyoptions)
		},
		DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {

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
			conn, err := options.Proxy_net_DialContext(ctx, network, addr, proxyoptions) //dialer.DialContext(ctx, network, newAddr)
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

	if len(proxyoptions) > 0 {

		client.Transport = transport
	}
	/* 流式处理,防止内存溢出 */
	proxyReq, err := http.NewRequest(r.Method, targetUrl, r.Body /* bytes.NewReader(bodyBytes) */)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyReq.Header = r.Header.Clone()

	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	// Copy headers from the response to the client's response.
	for k, v := range resp.Header {
		w.Header()[k] = v
	}

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
}

// func main() {
// 	// 监听本地8080端口
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// 	log.Printf("Proxy server started on port %s", listener.Addr())

// 	// 设置自定义处理器
// 	http.HandleFunc("/", proxyHandler)

//		// 开始服务
//		err = http.Serve(listener, nil)
//		if err != nil {
//			log.Fatal("Serve: ", err)
//		}
//	}
func Http(hostname string, port int, proxyoptions options.ProxyOptions) {

	engine := gin.Default()

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
	engine.Any("/*path", func(c *gin.Context) {
		var w = c.Writer
		var r = c.Request
		proxyHandler(w, r /* jar, */, LocalAddr, proxyoptions)

	})
	// 设置自定义处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		engine.Handler().ServeHTTP(w, r)
	})

	// 开始服务
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
func GenerateRandomLoopbackIP() string {
	rand.Seed(time.Now().UnixNano())
	randomIP := generateRandomIP()
	fmt.Println("Random IP:", randomIP)
	return randomIP.String()
}

func generateRandomIP() net.IP {
	ip := net.IPv4(
		byte(127 /* +rand.Intn(1) */),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
	)
	return ip
}

func GenerateRandomIntPort() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := generateRandomInt()
	fmt.Println("Random integer:", randomInt)
	return randomInt
}

func generateRandomInt() int {
	minport := 10000
	maxport := 65535
	return rand.Intn(maxport-minport+1) + minport
}

// IsIP 判断给定的字符串是否是有效的 IPv4 或 IPv6 地址。
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}
