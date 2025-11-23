package dns_experiment

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/ameshkov/dnscrypt/v2"
	"github.com/masx200/dnsproxy/upstream"
)

// CustomUpstreamOptions 自定义上游选项，支持指定服务器IP
type CustomUpstreamOptions struct {
	*upstream.Options
	serverIP  string
	serverURL string
}

// Clone 实现上游选项接口
func (c *CustomUpstreamOptions) Clone() upstream.UpstreamOptions {
	return &CustomUpstreamOptions{
		Options:   c.Options,
		serverIP:  c.serverIP,
		serverURL: c.serverURL,
	}
}

// SetBootstrap 实现上游选项接口
func (c *CustomUpstreamOptions) SetBootstrap(bootstrap upstream.Resolver) {
	c.Options.Bootstrap = bootstrap
}

// SetLogger 实现上游选项接口
func (c *CustomUpstreamOptions) SetLogger(logger *slog.Logger) {
	// 这里可以设置日志记录器，如果需要的话
	if c.Options != nil {
		c.Options.Logger = logger
	}
}

// GetTimeout 获取超时时间
func (c *CustomUpstreamOptions) GetTimeout() time.Duration {
	if c.Options != nil && c.Options.Timeout > 0 {
		return c.Options.Timeout
	}
	return 30 * time.Second
}

// DialTCP 自定义TCP拨号函数
func (c *CustomUpstreamOptions) DialTCP(ctx context.Context, addr string) (net.Conn, error) {
	if c.serverIP != "" {
		// 从URL中提取端口
		_, port, err := net.SplitHostPort(c.serverURL)
		if err != nil {
			// 如果解析失败，尝试使用默认端口
			if port == "" {
				port = "853" // DoT默认端口
			}
		}

		// 使用指定的服务器IP
		customAddr := net.JoinHostPort(c.serverIP, port)
		dialer := &net.Dialer{
			Timeout: c.GetTimeout(),
		}
		return dialer.DialContext(ctx, "tcp", customAddr)
	}

	// 如果没有指定服务器IP，使用默认拨号
	dialer := &net.Dialer{
		Timeout: c.GetTimeout(),
	}
	return dialer.DialContext(ctx, "tcp", addr)
}

// DialUDP 自定义UDP拨号函数（主要用于DoQ）
func (c *CustomUpstreamOptions) DialUDP(ctx context.Context, addr string) (*net.UDPConn, error) {
	if c.serverIP != "" {
		// 从URL中提取端口
		_, port, err := net.SplitHostPort(c.serverURL)
		if err != nil {
			// 如果解析失败，尝试使用默认端口
			if port == "" {
				port = "784" // DoQ默认端口
			}
		}

		// 使用指定的服务器IP
		customAddr := net.JoinHostPort(c.serverIP, port)
		dialer := &net.Dialer{
			Timeout: c.GetTimeout(),
		}

		conn, err := dialer.DialContext(ctx, "udp", customAddr)
		if err != nil {
			return nil, err
		}

		// 尝试转换为UDP连接
		if udpConn, ok := conn.(*net.UDPConn); ok {
			return udpConn, nil
		}

		// 如果不是UDP连接，关闭连接并返回错误
		conn.Close()
		return nil, fmt.Errorf("failed to create UDP connection")
	}

	// 如果没有指定服务器IP，使用默认拨号
	dialer := &net.Dialer{
		Timeout: c.GetTimeout(),
	}

	conn, err := dialer.DialContext(ctx, "udp", addr)
	if err != nil {
		return nil, err
	}

	// 尝试转换为UDP连接
	if udpConn, ok := conn.(*net.UDPConn); ok {
		return udpConn, nil
	}

	// 如果不是UDP连接，关闭连接并返回错误
	conn.Close()
	return nil, fmt.Errorf("failed to create UDP connection")
}

// 实现其他必需的接口方法
func (c *CustomUpstreamOptions) GetHTTPVersions() []upstream.HTTPVersion {
	return upstream.DefaultHTTPVersions
}

func (c *CustomUpstreamOptions) GetInsecureSkipVerify() bool {
	if c.Options != nil {
		return c.Options.InsecureSkipVerify
	}
	return false
}

func (c *CustomUpstreamOptions) GetPreferIPv6() bool {
	if c.Options != nil {
		return c.Options.PreferIPv6
	}
	return false
}

// 为了简化实现，其他方法直接返回默认值或委托给内部Options
func (c *CustomUpstreamOptions) GetVerifyServerCertificate() func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	if c.Options != nil {
		return c.Options.VerifyServerCertificate
	}
	return nil
}

func (c *CustomUpstreamOptions) GetVerifyConnection() func(state tls.ConnectionState) error {
	if c.Options != nil {
		return c.Options.VerifyConnection
	}
	return nil
}

func (c *CustomUpstreamOptions) GetVerifyDNSCryptCertificate() func(cert *dnscrypt.Cert) error {
	if c.Options != nil {
		return c.Options.VerifyDNSCryptCertificate
	}
	return nil
}

func (c *CustomUpstreamOptions) GetQUICTracer() upstream.QUICTraceFunc {
	if c.Options != nil {
		return c.Options.QUICTracer
	}
	return nil
}

func (c *CustomUpstreamOptions) GetRootCAs() *x509.CertPool {
	if c.Options != nil {
		return c.Options.RootCAs
	}
	return nil
}

func (c *CustomUpstreamOptions) GetCipherSuites() []uint16 {
	if c.Options != nil {
		return c.Options.CipherSuites
	}
	return nil
}

func (c *CustomUpstreamOptions) GetBootstrap() upstream.Resolver {
	if c.Options != nil {
		return c.Options.Bootstrap
	}
	return nil
}
