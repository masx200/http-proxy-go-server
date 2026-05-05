package doh

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

// h3ClientEntry 缓存的 H3 客户端条目，持有长生命周期的 quic.Transport
// 避免每次 DoH3 请求都创建新的 UDP socket 和 QUIC Transport
type h3ClientEntry struct {
	client       *http.Client
	quicTransport *quic.Transport
	h3Transport  *http3.Transport
	mu           sync.Mutex
}

// h3ClientCache 缓存按 (dohurl+dohip) 为 key 的 H3 客户端
var h3ClientCache sync.Map

// getOrCreateH3Client 获取或创建一个可复用的 HTTP/3 客户端
// 关键：quic.Transport（底层 UDP socket）被复用，彻底消除 goroutine 泄漏
func getOrCreateH3Client(dohurl string, dohip string) *h3ClientEntry {
	cacheKey := dohurl + "|" + dohip
	if v, ok := h3ClientCache.Load(cacheKey); ok {
		return v.(*h3ClientEntry)
	}

	entry := createH3Client(dohurl, dohip)
	if actual, loaded := h3ClientCache.LoadOrStore(cacheKey, entry); loaded {
		// 另一个 goroutine 先创建了，关闭我们创建的并使用已有的
		entry.quicTransport.Close()
		entry.h3Transport.Close()
		return actual.(*h3ClientEntry)
	}
	return entry
}

// createH3Client 创建一个新的 HTTP/3 客户端（持有永久 quic.Transport）
func createH3Client(dohurl string, dohip string) *h3ClientEntry {
	udpConn, err := net.ListenUDP("udp", nil)
	if err != nil {
		// 降级：返回一个不含 quic.Transport 的基础 entry
		log.Println("H3客户端：创建UDP socket失败", err)
		h3tr := &http3.Transport{}
		return &h3ClientEntry{
			client:      &http.Client{Transport: h3tr, Timeout: 30 * time.Second},
			h3Transport: h3tr,
		}
	}

	qt := &quic.Transport{Conn: udpConn}

	// 选择拨号函数
	var dialFunc func(ctx context.Context, addr string, tlsConf *tls.Config, quicConf *quic.Config) (*quic.Conn, error)
	if dohip != "" {
		ip := dohip
		dialFunc = func(ctx context.Context, addr string, tlsConf *tls.Config, quicConf *quic.Config) (*quic.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			addr2 := net.JoinHostPort(ip, port)
			a, err := net.ResolveUDPAddr("udp", addr2)
			if err != nil {
				return nil, err
			}
			conn, err := qt.DialEarly(ctx, a, tlsConf, quicConf)
			if err != nil {
				log.Println("H3连接失败", host, port)
				return nil, err
			}
			log.Println("H3连接成功(复用Transport)", host, port, conn.LocalAddr(), conn.RemoteAddr())
			return conn, nil
		}
	} else {
		dialFunc = func(ctx context.Context, addr string, tlsConf *tls.Config, quicConf *quic.Config) (*quic.Conn, error) {
			a, err := net.ResolveUDPAddr("udp", addr)
			if err != nil {
				return nil, err
			}
			conn, err := qt.DialEarly(ctx, a, tlsConf, quicConf)
			if err != nil {
				log.Println("H3连接失败(无IP)", addr)
				return nil, err
			}
			log.Println("H3连接成功(复用Transport,无IP)", addr, conn.LocalAddr(), conn.RemoteAddr())
			return conn, nil
		}
	}

	h3tr := &http3.Transport{
		Dial: dialFunc,
	}

	return &h3ClientEntry{
		client:        &http.Client{Transport: h3tr, Timeout: 30 * time.Second},
		quicTransport: qt,
		h3Transport:   h3tr,
	}
}

// doHTTP3ClientCached 使用缓存的 H3 客户端执行 DoH3 查询
func doHTTP3ClientCached(msg *dns.Msg, dohttp3ServerURL string, dohip ...string) (*dns.Msg, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg.Id = 0
	body, err := msg.Pack()
	if err != nil {
		log.Println(dohttp3ServerURL, err)
		return nil, err
	}

	ipKey := ""
	if len(dohip) > 0 {
		ipKey = dohip[0]
	}
	entry := getOrCreateH3Client(dohttp3ServerURL, ipKey)

	req, err := http.NewRequestWithContext(ctx, "POST", dohttp3ServerURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/dns-message")

	entry.mu.Lock()
	client := entry.client
	entry.mu.Unlock()

	res, err := client.Do(req)
	if err != nil {
		log.Println(dohttp3ServerURL, err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println(dohttp3ServerURL, "http status code is not 200", fmt.Sprintf("status code is %d", res.StatusCode))
		return nil, errors.New("http status code is not 200 " + fmt.Sprintf("status code is %d", res.StatusCode))
	}

	if res.Header.Get("Content-Type") != "application/dns-message" {
		log.Println(dohttp3ServerURL, "content-type is not application/dns-message "+res.Header.Get("Content-Type"))
		return nil, errors.New("content-type is not application/dns-message " + res.Header.Get("Content-Type"))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(dohttp3ServerURL, err)
		return nil, err
	}

	resp := &dns.Msg{}
	err = resp.Unpack(data)
	if err != nil {
		log.Println(dohttp3ServerURL, err)
		return nil, err
	}
	return resp, nil
}

func Doh3nslookup(domain string, dnstype string, dohurl string, dohip ...string) ([]*dns.Msg, []error) {
	log.Println("domain:", domain, "dnstype:", dnstype, "dohurl:", dohurl)
	var errs = make([]error, 0)
	var results = make([]*dns.Msg, 0)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, d := range strings.Split(domain, ",") {
		for _, t := range strings.Split(dnstype, ",") {
			wg.Add(1)
			go func(d string, t string) {
				defer wg.Done()
				log.Println("domain:", d, "dnstype:", t, "dohurl:", dohurl)
				var msg = &dns.Msg{}
				msg.SetQuestion(d+".", dns.StringToType[t])

				// 使用缓存的 H3 客户端，不再每次创建新的 quic.Transport
				res, err := doHTTP3ClientCached(msg, dohurl, dohip...)
				mutex.Lock()
				defer mutex.Unlock()
				if err != nil {
					log.Println(err)
					errs = append(errs, err)
					return
				}
				results = append(results, res)
			}(d, t)
		}
	}
	wg.Wait()
	return results, errs
}

func ResolveDomainToIPsWithDoh3(domain string, dohurl string, dohip ...string) ([]net.IP, []error) { // 使用 A 和 AAAA 记录类型查询域名

	dnstypes := "A,AAAA"
	responses, errors := Doh3nslookup(domain, dnstypes, dohurl, dohip...)
	if len(responses) == 0 && len(errors) > 0 {
		return nil, errors
	}
	var ips []net.IP
	for _, response := range responses {
		for _, record := range response.Answer {
			switch r := record.(type) {
			case *dns.A:
				ips = append(ips, r.A)
			case *dns.AAAA:
				ips = append(ips, r.AAAA)
			}
		}
	}

	if len(ips) == 0 {
		return nil, []error{fmt.Errorf("no IP addresses found for domain %s", domain)}
	}

	// 将 []net.IP 转换为 []string
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	// 打印日志
	log.Println("dns resolved " + domain + " ips:[" + strings.Join(ipStrings, ",") + "]")

	return ips, nil
}
