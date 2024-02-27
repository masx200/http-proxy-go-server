package http

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
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
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetUrl := "http://" + r.URL.Host + r.URL.Path // 这里假设目标服务器都是HTTP的，实际情况可能需要处理HTTPS

	// 创建一个使用了代理的客户端
	client := &http.Client{ /* Transport: newTransport("http://your_proxy_address:port") */ } // 替换为你的代理服务器地址和端口
	proxyReq, err := http.NewRequest(r.Method, targetUrl, r.Body)
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
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(bodyBytes); err != nil {
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
	/* /* 有的服务器不支持这种 "GET http://speedtest.cn/ HTTP/1.1" */
	// 监听本地8080端口
	listener, err := net.Listen("tcp", hostname+":"+fmt.Sprint(port))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Proxy server started on port %s", listener.Addr())

	// 设置自定义处理器
	http.HandleFunc("/", proxyHandler)

	// 开始服务
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
