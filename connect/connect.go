package connect

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"strings"
)

func ConnectViaHttpProxy(proxyURL *url.URL) (net.Conn, error) {
	// 解析代理服务器地址
	proxyAddr := proxyURL.Host
	if !strings.Contains(proxyAddr, ":") {
		proxyAddr += ":80" // 默认端口
	}

	// 建立到代理服务器的TCP连接
	conn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to proxy: %v", err)
	}

	// 从proxyURL中获取目标地址（这里假设proxyURL的Path或Opaque包含目标地址）
	targetAddr := proxyURL.Path
	if targetAddr == "" {
		targetAddr = proxyURL.Opaque
	}
	if targetAddr != "" && targetAddr[0] == '/' {
		targetAddr = targetAddr[1:]
	}

	if targetAddr == "" {
		conn.Close()
		return nil, fmt.Errorf("target address not specified in proxy URL")
	}

	// 确保目标地址包含端口
	if !strings.Contains(targetAddr, ":") {
		if strings.HasPrefix(proxyURL.String(), "https://") {
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
	_, err = fmt.Sscanf(respLine, "HTTP/1.%*d %d", &statusCode)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to parse status code: %v", err)
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
