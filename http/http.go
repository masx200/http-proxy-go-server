package http

import (
	// "bytes"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
	// "github.com/go-kit/kit/sd/etcd"
	// "net/url"
)

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
func proxyHandler(w http.ResponseWriter, r *http.Request, jar *cookiejar.Jar) {
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("host:", r.Host)
	fmt.Println("header:")

	for k, v := range r.Header {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
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
	/* 能否解决请求body的问题 */
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("body:", string(bodyBytes))
	client := &http.Client{ /* Transport: newTransport("http://your_proxy_address:port") */
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse /* 不进入重定向 */
		},

		Jar: jar} // 替换为你的代理服务器地址和端口
	proxyReq, err := http.NewRequest(r.Method, targetUrl, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyReq.Header = r.Header

	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Copy headers from the response to the client's response.
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)

	// Copy the response body back to the client.
	bodyBytes2, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(bodyBytes2); err != nil {
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
func Http(hostname string, port int) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	/* /* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
	// 监听本地8080端口
	listener, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Proxy server started on port %s", listener.Addr())

	// 设置自定义处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		proxyHandler(w, r, jar)
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
	min := 10000
	max := 65535
	return rand.Intn(max-min+1) + min
}
