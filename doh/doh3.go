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
	"sync/atomic"
	"time"

	"github.com/miekg/dns"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

// h3ClientEntry 缓存的 H3 客户端条目，持有长生命周期的 quic.Transport
// 避免每次 DoH3 请求都创建新的 UDP socket 和 QUIC Transport
type h3ClientEntry struct {
	client        *http.Client
	quicTransport *quic.Transport
	h3Transport   *http3.Transport
	mu            sync.Mutex
	closed        atomic.Bool // 标记是否已关闭
	createdAt     time.Time  // 创建时间，用于超时检查
}

// Close 关闭 H3 客户端条目，释放所有资源
func (e *h3ClientEntry) Close() {
	if e.closed.Swap(true) {
		return // 已经关闭过了
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.client != nil {
		// 关闭 HTTP 客户端（会关闭底层的 transport）
		if t, ok := e.client.Transport.(*http3.Transport); ok {
			t.Close()
		}
	}
	if e.quicTransport != nil {
		e.quicTransport.Close()
	}
	if e.h3Transport != nil {
		e.h3Transport.Close()
	}
}

// IsStale 检查客户端是否过期（超过指定时间未使用）
func (e *h3ClientEntry) IsStale(maxAge time.Duration) bool {
	return time.Since(e.createdAt) > maxAge
}

// h3ClientCache 缓存按 (dohurl+dohip) 为 key 的 H3 客户端
var h3ClientCache sync.Map

// h3CacheCleaner 定期清理过期客户端的 goroutine
var h3CacheCleanerStop chan struct{}
var h3CacheCleanerWG sync.WaitGroup

// StartH3CacheCleaner 启动定期清理过期 H3 客户端的 goroutine
func StartH3CacheCleaner(interval time.Duration, maxAge time.Duration) {
	if h3CacheCleanerStop != nil {
		return // 已经启动
	}
	h3CacheCleanerStop = make(chan struct{})
	h3CacheCleanerWG.Add(1)
	go func() {
		defer h3CacheCleanerWG.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				CleanupStaleH3Clients(maxAge)
			case <-h3CacheCleanerStop:
				return
			}
		}
	}()
	log.Printf("H3 客户端缓存清理器已启动，间隔: %v, 最大存活时间: %v", interval, maxAge)
}

// StopH3CacheCleaner 停止缓存清理 goroutine
func StopH3CacheCleaner() {
	if h3CacheCleanerStop != nil {
		close(h3CacheCleanerStop)
		h3CacheCleanerStop = nil
		h3CacheCleanerWG.Wait()
		log.Println("H3 客户端缓存清理器已停止")
	}
}

// CleanupStaleH3Clients 清理所有过期的 H3 客户端
func CleanupStaleH3Clients(maxAge time.Duration) {
	cleaned := 0
	h3ClientCache.Range(func(key, value interface{}) bool {
		entry := value.(*h3ClientEntry)
		if entry.IsStale(maxAge) {
			entry.Close()
			h3ClientCache.Delete(key)
			cleaned++
		}
		return true
	})
	if cleaned > 0 {
		log.Printf("已清理 %d 个过期的 H3 客户端", cleaned)
	}
}

// CloseH3ClientCache 关闭并清理所有 H3 客户端缓存
// 必须在程序退出时调用，防止 goroutine 泄漏
func CloseH3ClientCache() {
	// 先停止清理器
	StopH3CacheCleaner()

	// 关闭并删除所有缓存的客户端
	closed := 0
	h3ClientCache.Range(func(key, value interface{}) bool {
		entry := value.(*h3ClientEntry)
		entry.Close()
		h3ClientCache.Delete(key)
		closed++
		return true
	})
	log.Printf("H3 客户端缓存已关闭，共清理 %d 个客户端", closed)
}

// GetH3ClientCacheStats 返回缓存统计信息（用于调试）
func GetH3ClientCacheStats() (count int, entries []string) {
	h3ClientCache.Range(func(key, value interface{}) bool {
		count++
		entries = append(entries, key.(string))
		return true
	})
	return count, entries
}

// getOrCreateH3Client 获取或创建一个可复用的 HTTP/3 客户端
// 关键：quic.Transport（底层 UDP socket）被复用，彻底消除 goroutine 泄漏
func getOrCreateH3Client(dohurl string, dohip string) *h3ClientEntry {
	cacheKey := dohurl + "|" + dohip
	if v, ok := h3ClientCache.Load(cacheKey); ok {
		entry := v.(*h3ClientEntry)
		// 检查是否已被关闭，如果是则删除并重新创建
		if entry.closed.Load() {
			entry.Close()
			h3ClientCache.Delete(cacheKey)
		} else {
			return entry
		}
	}

	entry := createH3Client(dohurl, dohip)
	if actual, loaded := h3ClientCache.LoadOrStore(cacheKey, entry); loaded {
		// 另一个 goroutine 先创建了，关闭我们创建的并使用已有的
		entry.Close()
		return actual.(*h3ClientEntry)
	}
	return entry
}

// createH3Client 创建一个新的 HTTP/3 客户端（持有永久 quic.Transport）
func createH3Client(dohurl string, dohip string) *h3ClientEntry {
	entry := &h3ClientEntry{
		createdAt: time.Now(),
	}

	udpConn, err := net.ListenUDP("udp", nil)
	if err != nil {
		// 降级：返回一个不含 quic.Transport 的基础 entry
		log.Println("H3客户端：创建UDP socket失败", err)
		h3tr := &http3.Transport{}
		entry.client = &http.Client{Transport: h3tr, Timeout: 30 * time.Second}
		entry.h3Transport = h3tr
		return entry
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

	entry.client = &http.Client{Transport: h3tr, Timeout: 30 * time.Second}
	entry.quicTransport = qt
	entry.h3Transport = h3tr
	return entry
}

// doHTTP3ClientCached 使用缓存的 H3 客户端执行 DoH3 查询
func doHTTP3ClientCached(msg *dns.Msg, dohttp3ServerURL string, dohip ...string) (*dns.Msg, error) {
	// 使用更短的默认超时，防止 QUIC 连接挂起导致 goroutine 泄漏
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	// 获取 H3 客户端，如果获取失败则降级到普通 HTTP
	entry := getOrCreateH3Client(dohttp3ServerURL, ipKey)
	if entry == nil || entry.client == nil {
		return nil, errors.New("failed to get H3 client")
	}

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
		// 如果连接失败，尝试关闭并重建客户端
		if entry.closed.Load() {
			// 客户端已被关闭，清理缓存
			h3ClientCache.Delete(dohttp3ServerURL + "|" + ipKey)
		}
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
