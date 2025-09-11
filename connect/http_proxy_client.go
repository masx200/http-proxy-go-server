package connect

import (
	"context"
	"fmt"
	"github.com/masx200/socks5-websocket-proxy-golang/pkg/interfaces"
	"log"
	"net"
	"net/url"
	"strconv"
)

// HttpProxyClient HTTP代理客户端实现
type HttpProxyClient struct {
	config        interfaces.ClientConfig
	conn          net.Conn
	closed        bool
	closeCallback func()
}

// NewHttpProxyClient 创建HTTP代理客户端
func NewHttpProxyClient(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
	return &HttpProxyClient{
		config: config,
		closed: false,
	}, nil
}

// Connect 连接到目标主机
func (c *HttpProxyClient) Connect(host string, port int) error {

	log.Println("Connecting to", host, port, "via HTTP proxy", c.config.ServerAddr)

	if c.conn != nil {
		c.conn.Close()
	}

	// 构建代理URL
	proxyURL, err := url.Parse(c.config.ServerAddr)
	if err != nil {
		return err
	}

	// 添加认证信息
	if c.config.Username != "" && c.config.Password != "" {
		proxyURL.User = url.UserPassword(c.config.Username, c.config.Password)
	}

	// 构建目标地址
	targetAddr := net.JoinHostPort(host, strconv.Itoa(port))

	// 使用ConnectViaHttpProxy建立连接
	conn, err := ConnectViaHttpProxy(proxyURL, targetAddr)
	if err != nil {
		return err
	}

	c.conn = conn
	c.closed = false
	return nil
}

// ForwardData 转发数据
func (c *HttpProxyClient) ForwardData(conn net.Conn) error {
	if c.conn == nil || c.closed {
		return net.ErrClosed
	}

	// 实现数据转发逻辑
	go func() {
		defer conn.Close()
		defer c.conn.Close()
		buf := make([]byte, 32*1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {

				if c.closeCallback != nil {
					c.closeCallback()
				}
				return
			}
			_, err = c.conn.Write(buf[:n])
			if err != nil {

				if c.closeCallback != nil {
					c.closeCallback()
				}
				return
			}
		}
	}()

	buf := make([]byte, 32*1024)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {

			if c.closeCallback != nil {
				c.closeCallback()
			}
			return err
		}
		_, err = conn.Write(buf[:n])
		if err != nil {

			if c.closeCallback != nil {
				c.closeCallback()
			}
			return err
		}
	}
}

// Close 关闭连接
func (c *HttpProxyClient) Close() error {
	if c.conn != nil && !c.closed {
		c.closed = true
		err := c.conn.Close()
		if c.closeCallback != nil {
			c.closeCallback()
		}
		return err
	}
	return nil
}

// SetConnectionClosedCallback 设置连接关闭回调
func (c *HttpProxyClient) SetConnectionClosedCallback(callback func()) error {
	c.closeCallback = callback
	return nil
}

// NetConn 获取底层网络连接
func (c *HttpProxyClient) NetConn() net.Conn {
	return c.conn
}

// DialContext 实现DialContext方法
func (c *HttpProxyClient) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	if c.conn == nil || c.closed {
		// 解析地址
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}

		// 转换端口为整数
		portInt := 0
		_, err = fmt.Sscanf(port, "%d", &portInt)
		if err != nil {
			return nil, err
		}

		// 连接到目标
		err = c.Connect(host, portInt)
		if err != nil {
			return nil, err
		}
	}
	return c.conn, nil
}

// CreateHttpProxyClient 创建HTTP代理客户端的工厂函数
var CreateHttpProxyClient interfaces.CreateClientFunc = func(config interfaces.ClientConfig) (interfaces.ProxyClient, error) {
	return NewHttpProxyClient(config)
}
