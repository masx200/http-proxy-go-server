# http-proxy-go-server

`http-proxy-go-server` 是一个基于 Golang（Go 语言）实现的 HTTP
代理服务器。这个服务器允许你通过指定的 IP 地址和端口作为 HTTP
代理来转发网络请求。从提供的命令行参数来看，它的主要功能和配置选项如下：

## usage

http-proxy-go-server

```text
-dohalpn value
        DOH alpn (可重复),支持h2协议和h3协议
-dohip value
        DOH IP (可重复),支持ipv4地址和ipv6地址
-dohurl value
        DOH URL (可重复),支持http协议和https协议
-hostname string
        an String value for hostname (default "0.0.0.0")
-password string
        password
-port int
        TCP port to listen on (default 8080)
-server_cert string
        tls server cert
-server_key string
        tls server key
-username string
        username
```

1. `-hostname string`：设置服务器绑定的主机名，默认为
   "0.0.0.0"，表示服务器将监听所有可用的网络接口。

2. `-password string`：设置访问代理服务器所需的密码，用于基本身份验证。

3. `-port int`：设置服务器监听的 TCP 端口号，默认为 8080。

4. `-server_cert string`：设置 HTTPS 服务所需的 TLS
   服务器证书文件路径。如果提供了此选项，服务器将以安全模式运行（HTTPS）。

5. `-server_key string`：设置 HTTPS 服务所需的 TLS
   私钥文件路径，与服务器证书配套使用。

6. `-username string`：设置访问代理服务器所需的用户名，同样用于基本身份验证。

总结来说，`http-proxy-go-server` 提供了一个可配置的、支持基本认证且可以运行在
HTTP 或 HTTPS
模式下的代理服务器。用户可以根据需要调整其监听地址、端口、认证凭据以及是否启用加密通信等配置项。

## JSON 配置文件

除了命令行参数外，`http-proxy-go-server` 还支持通过 JSON 配置文件进行配置。使用 `-config` 参数指定配置文件路径。

### 配置文件结构

JSON 配置文件支持以下参数：

```json
{
  "hostname": "0.0.0.0",
  "port": 8080,
  "server_cert": "",
  "server_key": "",
  "username": "",
  "password": "",
  "dohurls": ["https://dns.alidns.com/dns-query"],
  "dohips": ["223.5.5.5"],
  "dohalpns": ["h2"]
}
```

### 配置项说明

- `hostname`: 服务器绑定的主机名，默认为 "0.0.0.0"
- `port`: 服务器监听的 TCP 端口号，默认为 8080
- `server_cert`: HTTPS 服务所需的 TLS 服务器证书文件路径
- `server_key`: HTTPS 服务所需的 TLS 私钥文件路径
- `username`: 访问代理服务器所需的用户名
- `password`: 访问代理服务器所需的密码
- `dohurls`: DOH URL 数组，支持 http 和 https 协议
- `dohips`: DOH IP 数组，支持 ipv4 和 ipv6 地址
- `dohalpns`: DOH alpn 数组，支持 h2 和 h3 协议

### 使用配置文件

```bash
# 使用配置文件启动服务器
go run main.go -config config.json

# 配置文件和命令行参数可以混合使用，命令行参数会覆盖配置文件中的对应值
go run main.go -config config.json -port 9090 -username admin
```

### 配置文件优先级

配置文件中的参数会被用作默认值，但命令行参数会覆盖配置文件中的对应值。这样可以灵活地在基础配置上进行个性化调整。

## example

```bash
"/root/http-proxy-go-server/main" -dohurl "https://******************************" -dohip  "************" -port 58888 -username admin -password "*************************************"  -server_cert "**********************************************"  -server_key "**********************************************"
```

```bash
# 使用配置文件的示例
go run main.go -config config.json
```
