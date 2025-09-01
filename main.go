package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/masx200/http-proxy-go-server/auth"
	"github.com/masx200/http-proxy-go-server/options"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/tls"
	tls_auth "github.com/masx200/http-proxy-go-server/tls+auth"
)

type multiString []string

func (m *multiString) String() string {
	return "[" + strings.Join(*m, ", ") + "]"
}

func (m *multiString) Set(value string) error {
	*m = append(*m, value)
	return nil
}

// Config 结构体用于JSON配置文件
type Config struct {
	Hostname   string   `json:"hostname"`
	Port       int      `json:"port"`
	ServerCert string   `json:"server_cert"`
	ServerKey  string   `json:"server_key"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Dohurls    []string `json:"dohurls"`
	Dohips     []string `json:"dohips"`
	Dohalpns   []string `json:"dohalpns"`
}

// loadConfig 从JSON文件加载配置
func loadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	// 添加配置文件参数
	configFile := flag.String("config", "", "JSON配置文件路径")

	// 自定义字符串切片类型，实现 flag.Value 接口
	var (
		dohurls  multiString
		dohips   multiString
		dohalpns multiString
	)
	// 注册可重复参数
	flag.Var(&dohurls, "dohurl", "DOH URL (可重复),支持http协议和https协议")
	flag.Var(&dohips, "dohip", "DOH IP (可重复),支持ipv4地址和ipv6地址")
	flag.Var(&dohalpns, "dohalpn", "DOH alpn (可重复),支持h2协议和h3协议")

	var (
		hostname    = flag.String("hostname", "0.0.0.0", "an String value for hostname")
		port        = flag.Int("port", 8080, "TCP port to listen on")
		server_cert = flag.String("server_cert", "", "tls server cert")
		server_key  = flag.String("server_key", "", "tls server key")
		username    = flag.String("username", "", "username")
		password    = flag.String("password", "", "password")
	)
	flag.Parse()

	// 如果指定了配置文件，则从配置文件读取参数
	var config *Config
	var err error
	if *configFile != "" {
		config, err = loadConfig(*configFile)
		if err != nil {
			fmt.Printf("读取配置文件失败: %v\n", err)
			os.Exit(1)
		}
		// 使用配置文件的值覆盖命令行参数的默认值
		if config.Hostname != "" {
			*hostname = config.Hostname
		}
		if config.Port != 0 {
			*port = config.Port
		}
		if config.ServerCert != "" {
			*server_cert = config.ServerCert
		}
		if config.ServerKey != "" {
			*server_key = config.ServerKey
		}
		if config.Username != "" {
			*username = config.Username
		}
		if config.Password != "" {
			*password = config.Password
		}
		if len(config.Dohurls) > 0 {
			dohurls = multiString(config.Dohurls)
		}
		if len(config.Dohips) > 0 {
			dohips = multiString(config.Dohips)
		}
		if len(config.Dohalpns) > 0 {
			dohalpns = multiString(config.Dohalpns)
		}
	}
	fmt.Println("dohalpn:", dohalpns.String())
	//parse cmd flags
	fmt.Println(
		"hostname:", *hostname)
	fmt.Println(
		"port:", *port)
	fmt.Println(
		"server_cert:", *server_cert)
	fmt.Println(
		"server_key:", *server_key)
	fmt.Println(
		"username:", *username)
	fmt.Println(
		"password:", *password)
	fmt.Println(
		"dohurl:", dohurls.String())
	fmt.Println("dohip:", dohips.String())
	var proxyoptions = options.ProxyOptions{}

	for i, dohurl := range dohurls {

		var dohip string
		if len(dohips) > i {
			dohip = dohips[i]
		} else {
			dohip = ""
		}
		var dohalpn string
		if len(dohalpns) > i {
			dohalpn = dohalpns[i]
		} else {
			dohalpn = ""
		}

		proxyoptions = append(proxyoptions, options.ProxyOption{Dohurl: dohurl, Dohip: dohip, Dohalpn: dohalpn})
	}
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password, proxyoptions)
		return
	}
	// if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
	// 	tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password)
	// 	return
	// }
	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		auth.Auth(*hostname, *port, *username, *password, proxyoptions)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls.Tls(*server_cert, *server_key, *hostname, *port, proxyoptions)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		simple.Simple(*hostname, *port, proxyoptions)
		return
	}

}
