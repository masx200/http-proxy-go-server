package connect

import (
	"bufio"

	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"

	"github.com/masx200/http-proxy-go-server/dnscache"
	"github.com/masx200/http-proxy-go-server/options"
)

// ConnectViaHttpProxy 通过HTTP代理服务器建立网络连接。
// 该函数使用HTTP CONNECT方法通过代理服务器连接到目标地址。
//
// 参数:
//   - proxyURL: 代理服务器的URL，包含代理地址、协议（http/https）、认证信息等。
//
// 返回值:
//   - net.Conn: 成功时返回与目标地址建立的网络连接。
//   - error: 如果连接失败或代理响应异常，返回相应的错误信息。

func ConnectViaHttpProxy(proxyURL *url.URL, targetAddr string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool) (net.Conn, error) {
	log.Println("开始连接代理服务器", proxyURL, targetAddr)
	var scheme = proxyURL.Scheme

	// 解析代理服务器地址
	proxyAddr := proxyURL.Host
	if !strings.Contains(proxyAddr, ":") {
		if scheme == "https" {
			proxyAddr += ":443" // HTTPS默认端口
		} else {
			proxyAddr += ":80" // HTTP默认端口
		}
	}

	// 建立到代理服务器的TCP连接
	var conn net.Conn
	var err error
	if scheme == "https" {
		// 使用TLS连接
		conn, err = tls.Dial("tcp", proxyAddr, &tls.Config{
			ServerName: proxyURL.Hostname(),
		})
	} else {
		// 使用普通TCP连接
		conn, err = net.Dial("tcp", proxyAddr)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to proxy: %v", err)
	}

	// 从proxyURL中获取目标地址（这里假设proxyURL的Path或Opaque包含目标地址）
	// targetAddr := proxyURL.Path
	// if targetAddr == "" {
	// 	targetAddr = proxyURL.Opaque
	// }
	// if targetAddr != "" && targetAddr[0] == '/' {
	// 	targetAddr = targetAddr[1:]
	// }

	// if targetAddr == "" {
	// 	conn.Close()
	// 	return nil, fmt.Errorf("target address not specified in proxy URL")
	// }

	// 如果启用了DNS解析，先解析目标地址
	if upstreamResolveIPs && len(proxyoptions) > 0 && dnsCache != nil {
		resolvedAddrs, err := resolveTargetAddressForHttp(targetAddr, proxyoptions, dnsCache)
		if err != nil {
			log.Printf("Failed to resolve target address %s: %v, using original", targetAddr, err)
		} else {
			// 使用轮询从解析的地址中选择一个
			targetAddr = resolveTargetAddressForHttpWithRoundRobin(resolvedAddrs, targetAddr)
			log.Printf("Resolved HTTP proxy target address to: %s", targetAddr)
		}
	}

	// 确保目标地址包含端口
	if !strings.Contains(targetAddr, ":") {
		if scheme == "https" {
			targetAddr += ":443"
		} else {
			targetAddr += ":80"
		}
	}

	// 构造CONNECT请求
	connectReq := fmt.Sprintf("CONNECT %s HTTP/1.1\r\n", targetAddr)
	connectReq += fmt.Sprintf("Host: %s\r\n", targetAddr)

	// 添加代理认证信息（如果存在）
	if proxyURL.User != nil {
		username := proxyURL.User.Username()
		password, _ := proxyURL.User.Password()
		if username != "" && password != "" {
			auth := username + ":" + password
			connectReq += fmt.Sprintf("Proxy-Authorization: Basic %s\r\n",
				base64.StdEncoding.EncodeToString([]byte(auth)))
		}
	}

	connectReq += "\r\n"

	// 发送CONNECT请求
	_, err = conn.Write([]byte(connectReq))
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to send CONNECT request: %v", err)
	}

	// 读取代理服务器的响应
	reader := bufio.NewReader(conn)
	respLine, err := reader.ReadString('\n')
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to read proxy response: %v", err)
	}

	// 解析状态行
	if !strings.HasPrefix(respLine, "HTTP/1.") {
		conn.Close()
		return nil, fmt.Errorf("invalid HTTP response from proxy")
	}

	// 检查状态码
	var statusCode int
	_, err = fmt.Sscanf(respLine, "HTTP/1.1 %d", &statusCode)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to parse status code: %v "+respLine, err)
	}

	// 读取剩余的头部
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return nil, fmt.Errorf("failed to read headers: %v", err)
		}
		if line == "\r\n" {
			break
		}
	}

	// 检查状态码是否为200（连接建立成功）
	if statusCode != 200 {
		conn.Close()
		return nil, fmt.Errorf("proxy returned status code %d", statusCode)
	}

	// 返回连接，此时连接已经可以用于数据传输
	return conn, nil
}
