package connect

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"
)

// CreateClientFunc 客户端创建函数类型
type CreateClientFunc func(config ClientConfig) (ProxyClient, error)

// ProxyClient 代理客户端接口
type ProxyClient interface {
	Connect(host string, port int) error
	// Authenticate(username, password string) error
	ForwardData(conn net.Conn) error
	Close() error
	SetConnectionClosedCallback(callback func()) error
	NetConn() net.Conn

	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Username   string        `json:"username"`    // 用户名
	Password   string        `json:"password"`    // 密码
	ServerAddr string        `json:"server_addr"` // 服务器地址
	Protocol   string        `json:"protocol"`    // 协议类型
	Timeout    time.Duration `json:"timeout"`     // 超时时间
}

// HttpProxyClient HTTP代理客户端实现
type HttpProxyClient struct {
	config        ClientConfig
	conn          net.Conn
	closed        bool
	closeCallback func()
}

// NewHttpProxyClient 创建HTTP代理客户端
func NewHttpProxyClient(config ClientConfig) (ProxyClient, error) {
	return &HttpProxyClient{
		config: config,
		closed: false,
	}, nil
}

// Connect 连接到目标主机
func (c *HttpProxyClient) Connect(host string, port int) error {
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
var CreateHttpProxyClient CreateClientFunc = func(config ClientConfig) (ProxyClient, error) {
	return NewHttpProxyClient(config)
}
