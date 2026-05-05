package dns_experiment

import (
	// "context"
	// "time"
	"context"
	"crypto/tls"
	"errors"
	"net"

	// "time"
	// "crypto/tls"
	// "fmt"
	// "io"
	// "net"
	// "net/http"
	// "crypto/tls"
	//	"fmt"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	doq "github.com/masx200/doq-go/doq"
	"github.com/masx200/http-proxy-go-server/utils"
	print_experiment "github.com/masx200/http3-reverse-proxy-server-experiment/print"
	"github.com/miekg/dns"
)

// dohClientCache 缓存已创建的 http.Client，按 (serverURL+dohip) 为 key 复用连接池
// 避免每次 DNS 查询都创建新的 Transport，导致 goroutine 泄漏
var dohClientCache sync.Map

// dohClientCacheEntry 缓存条目，带有最后访问时间用于 LRU 淘汰
type dohClientCacheEntry struct {
	client   *http.Client
	lastUsed time.Time
	mu       sync.Mutex
}

// getDohClientCacheKey 生成缓存 key
func getDohClientCacheKey(dohServerURL, dohip string, hasProxy bool) string {
	proxyPart := "noproxy"
	if hasProxy {
		proxyPart = "proxy"
	}
	return dohServerURL + "|" + dohip + "|" + proxyPart
}

// DNSQueryHTTPS 执行DNS查询以获取HTTPS服务记录。
//
// 参数:
// - domain: 需要查询的域名。
// - port: 目标端口，如果不为"443"，则会构建特定端口的查询域名。
// - DOHServer: DNS-over-HTTPS服务器地址。
//
// 返回值:
// - []dns.SVCB: 查询到的HTTPS服务记录列表。
// - error: 查询过程中发生的任何错误。
func DNSQueryHTTPS(domain string, port string, DOHServer string, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) ([]dns.SVCB, error) {
	var msg = new(dns.Msg)
	var service_domain = domain
	if port != "443" {
		service_domain = fmt.Sprintf("_%s._https.", port) + domain
	}
	msg.SetQuestion(service_domain+".", dns.TypeHTTPS)

	resp, err := DohClient(msg, DOHServer, "", Proxy, tranportConfigurations...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.Rcode != dns.RcodeSuccess {
		log.Printf("DNS query failed: %s ", dns.RcodeToString[resp.Rcode]+" "+DOHServer+"\n")
		return nil, errors.New(
			"DNS query failed:  " + dns.RcodeToString[resp.Rcode] + " " + DOHServer)
	}
	if len(resp.Answer) == 0 {
		log.Println(DOHServer + "-No HTTPS records found")
		return nil, errors.New(
			"No HTTPS records found" + " " + DOHServer)
	}
	log.Println(DOHServer + "-" + resp.String())
	var result []dns.SVCB
	for _, answer := range resp.Answer {
		log.Println(answer)
		if a, ok := answer.(*dns.HTTPS); ok {
			log.Printf(DOHServer+"-https record for %s: \n", domain)
			result = append(result, a.SVCB)

		}
	}
	if len(result) == 0 {
		log.Println(DOHServer + "-No HTTPS records found")
		return nil, errors.New(
			"No HTTPS records found" + " " + DOHServer)
	}
	return result, nil

}
func setPortIfMissing(rawURL string) (string, error) {
	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// 检查是否存在端口
	if parsedURL.Port() == "" {
		// 如果不存在端口，则将端口设置为853
		parsedURL.Host = fmt.Sprintf("%s:853", parsedURL.Host)
	}

	// 返回修改后的URL字符串
	return parsedURL.String(), nil
}

// DohClient 是一个通过DOH（DNS over HTTPs）协议与DNS服务器进行通信的函数。
//
// 参数:
// msg: 代表DNS查询消息的dns.Msg对象。
// dohServer: 代表DOH服务器的URL字符串。
//
// 返回值:
// r: 代表DNS应答消息的dns.Msg对象。
// err: 如果过程中发生错误，则返回错误信息。
func DohClient(msg *dns.Msg, dohServerURL string, dohip string, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) (r *dns.Msg, err error) {
	// var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	/* 为了doh的缓存,需要设置id为0 ,可以缓存*/
	msg.Id = 0
	body, err := msg.Pack()
	if err != nil {
		log.Println(dohServerURL, err)
		return nil, err
	}
	req, err := http.NewRequest("POST", dohServerURL, strings.NewReader(string(body)))
	// req, err := http.NewRequestWithContext(ctx, "POST", dohServerURL, strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/dns-message")
	//http request doh

	// 获取或创建缓存的 http.Client，避免每次调用都创建新的 Transport/连接池
	// 相同的 (serverURL + dohip + hasProxy) 组合复用同一个 client
	client := getOrCreateDohClient(dohServerURL, dohip, Proxy, tranportConfigurations...)

	log.Println("开始发起doh请求", dohServerURL)
	res, err := client.Do(req) //Post(dohServerURL, "application/dns-message", strings.NewReader(string(body)))
	if err != nil {
		log.Println(dohServerURL, err)
		return nil, err
	}
	//res.status check
	if res.StatusCode != 200 {
		log.Println(dohServerURL, "http status code is not 200  "+fmt.Sprintf("status code is %d", res.StatusCode))
		return nil, errors.New("http status code is not 200 " + fmt.Sprintf("status code is %d", res.StatusCode))
	}

	//check content-type
	if res.Header.Get("Content-Type") != "application/dns-message" {
		log.Println(dohServerURL, "content-type is not application/dns-message "+res.Header.Get("Content-Type"))
		return nil, errors.New(dohServerURL + "content-type is not application/dns-message " + res.Header.Get("Content-Type"))
	}
	//利用ioutil包读取百度服务器返回的数据
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close() //一定要记得关闭连接
	if err != nil {
		log.Println(dohServerURL, err)
		return nil, err
	}
	// log.Printf("%s", data)
	resp := &dns.Msg{}
	err = resp.Unpack(data)
	if err != nil {
		log.Println(dohServerURL, err)
		return nil, err
	}
	return resp, nil
}

func PrintResponse(resp *http.Response) {
	print_experiment.PrintResponse(resp)
}

// DoQClient 是一个通过DOQ（DNS over QUIC）协议与DNS服务器进行通信的函数。
//
// 参数:
// msg 是一个包含DNS查询信息的dns.Msg结构体指针。
// dohServerURL 是一个字符串，表示DOQ服务器的URL。
//
// 返回值:
// 返回一个包含DNS应答信息的dns.Msg结构体指针和一个错误信息。
// 如果成功，错误信息为nil；如果发生错误，则返回相应的错误信息。
func DoQClient(msg *dns.Msg, doQServerURL string) (qA *dns.Msg, err error) {
	log.Println("doQServerURL", doQServerURL)
	urlWithPort, err := setPortIfMissing(doQServerURL)
	if err != nil {
		log.Println(doQServerURL, "Error:", err)
		return nil, err
	}
	doQServerURL = urlWithPort
	if !strings.HasPrefix(doQServerURL, "quic://") {
		log.Println(doQServerURL, "DOQ server URL must start with 'quic://'")
		return nil, errors.New(doQServerURL + "DOQ server URL must start with 'quic://'")
	}
	// 从DOH服务器URL中提取服务器名称和端口信息。
	serverName, port, err := ExtractDOQServerDetails(doQServerURL)
	if err != nil {
		log.Println(doQServerURL, err) // 记录提取详情时的错误
		return nil, err                // 如果有错误，返回nil和错误信息
	}
	var addr = fmt.Sprintf("%s:%s", serverName, port) // 格式化服务器地址
	log.Println("addr", addr)
	// 创建一个DOQ客户端
	client := doq.NewClient(addr)
	// 发送DNS查询并获取应答
	respA, err := client.Send(context.Background(), msg)
	if err != nil {
		log.Println(doQServerURL, err) // 记录发送时的错误
		return nil, err                // 如果有错误，返回nil和错误信息
	}
	return respA, err // 返回DNS应答和可能的错误信息
}

// ExtractDOQServerDetails takes a DOQ server URL and returns the server name and port as separate strings.
func ExtractDOQServerDetails(doqServer string) (string, string, error) {
	parts := strings.Split(doqServer, "://")
	if len(parts) != 2 {
		return "", "", errors.New("invalid DOQ server format")
	}

	serverWithPort := parts[1]
	serverParts := strings.Split(serverWithPort, ":")
	if len(serverParts) != 2 {
		return "", "", errors.New("invalid server details, missing port")
	}

	serverName := serverParts[0]
	port := serverParts[1]
	return serverName, port, nil
}

// DoTClient 是一个通过DOH（DNS over HTTPS）协议与DNS服务器进行通信的函数。
// msg: 包含要发送的DNS查询信息的dns.Msg对象。
// doTServerURL: DOH服务器的URL，用于指定通信的目标DNS服务器。
// 返回值 qA: 发送查询后收到的应答消息，为dns.Msg对象。
// 返回值 err: 如果在进行DNS查询过程中遇到错误，则返回错误信息。
func DoTClient(msg *dns.Msg, doTServerURL string) (qA *dns.Msg, err error) {
	log.Println("doTServerURL", doTServerURL)
	urlWithPort, err := setPortIfMissing(doTServerURL)
	if err != nil {
		log.Println(doTServerURL, "Error:", err)
		return nil, err
	}
	doTServerURL = urlWithPort
	if !strings.HasPrefix(doTServerURL, "tls://") {
		log.Println(doTServerURL, "DOT server URL must start with 'tls://'")
		return nil, errors.New(doTServerURL + "DOT server URL must start with 'tls://'")
	}
	// 从DOH服务器URL中解析出服务器名称和端口。
	serverName, port, err := ExtractDOQServerDetails(doTServerURL)
	if err != nil {
		log.Println(doTServerURL, err) // 记录解析服务器详情时的错误
		return nil, err                // 如果解析出错，则返回nil和错误信息
	}
	var addr = fmt.Sprintf("%s:%s", serverName, port) // 拼接服务器的地址信息
	log.Println("addr", addr)
	// 创建一个支持TCP-TLS的DOQ客户端实例。
	client := new(dns.Client)
	client.Net = "tcp-tls"
	// 向指定的DNS服务器发送查询请求，并接收应答。
	respA, _, err := client.Exchange(msg, addr)
	if err != nil {
		log.Println(doTServerURL, err) // 记录发送时的错误
		return nil, err                // 如果有错误，返回nil和错误信息
	}
	return respA, err // 返回查询应答和可能存在的错误信息
}

// getOrCreateDohClient 获取或创建一个可复用的 http.Client。
// 对相同的 (serverURL, dohip, proxy) 组合复用同一个 client，
// 让底层 Transport 的连接池充分工作，避免 goroutine 泄漏。
func getOrCreateDohClient(dohServerURL, dohip string, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) *http.Client {
	cacheKey := getDohClientCacheKey(dohServerURL, dohip, Proxy != nil)

	if v, ok := dohClientCache.Load(cacheKey); ok {
		entry := v.(*dohClientCacheEntry)
		entry.mu.Lock()
		entry.lastUsed = time.Now()
		entry.mu.Unlock()
		return entry.client
	}

	// 创建新的 client（仅在首次使用时执行）
	var client *http.Client
	if len(dohip) > 0 {
		serverIP := dohip
		transport := &http.Transport{
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     90 * time.Second,
		}
		var DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			log.Println("DialTLSContext", "dialing", network, "to", addr)
			address, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			parsedURL, err := url.Parse(dohServerURL)
			if err != nil {
				return nil, err
			}
			originalPort := parsedURL.Port()
			if originalPort == "" {
				if parsedURL.Scheme == "https" {
					originalPort = "443"
				} else {
					originalPort = "80"
				}
			}
			newAddr := net.JoinHostPort(serverIP, port)
			if transport.Proxy != nil && port != originalPort {
				newAddr = addr
				log.Println("DialTLSContext detected proxy, using original addr:", addr)
			}
			dialer := &net.Dialer{}
			log.Println("DialTLSContext", "dialing", network, "to", newAddr)
			conn, err := dialer.DialContext(ctx, network, newAddr)
			if err != nil {
				return nil, err
			}
			tlsConfig := &tls.Config{
				ServerName: address,
			}
			tlsConn := tls.Client(conn, tlsConfig)
			err = tlsConn.HandshakeContext(ctx)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return tlsConn, nil
		}
		transport.DialTLSContext = DialTLSContext
		var DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			var host, _, err = net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			if utils.IsLoopbackIP(host) {
				var dialer = &net.Dialer{}
				return dialer.DialContext(ctx, network, addr)
			}
			log.Println("DialContext", "dialing", network, "to", addr)
			_, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			parsedURL, err := url.Parse(dohServerURL)
			if err != nil {
				return nil, err
			}
			originalPort := parsedURL.Port()
			if originalPort == "" {
				if parsedURL.Scheme == "https" {
					originalPort = "443"
				} else {
					originalPort = "80"
				}
			}
			newAddr := net.JoinHostPort(serverIP, port)
			dialer := &net.Dialer{}
			if transport.Proxy != nil && port != originalPort {
				log.Println("DialContext detected proxy, using original addr:", addr)
				return dialer.DialContext(ctx, network, addr)
			}
			log.Println("dialContext", "dialing", network, "to", newAddr)
			return dialer.DialContext(ctx, network, newAddr)
		}
		transport.DialContext = DialContext
		if Proxy != nil {
			transport.Proxy = Proxy
		}
		for _, f := range tranportConfigurations {
			transport = f(transport)
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		}
	} else {
		// 无指定 dohip：创建独立的默认 transport，不修改全局 http.DefaultClient
		transport := &http.Transport{
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     90 * time.Second,
		}
		if Proxy != nil {
			transport.Proxy = Proxy
		}
		for _, f := range tranportConfigurations {
			transport = f(transport)
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		}
	}

	entry := &dohClientCacheEntry{
		client:   client,
		lastUsed: time.Now(),
	}

	// 使用 LoadOrStore 保证并发安全：若已有其他 goroutine 先存入，使用已有的
	if actual, loaded := dohClientCache.LoadOrStore(cacheKey, entry); loaded {
		return actual.(*dohClientCacheEntry).client
	}
	return client
}
