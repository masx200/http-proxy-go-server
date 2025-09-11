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

除了命令行参数外，`http-proxy-go-server` 还支持通过 JSON 配置文件进行配置。使用
`-config` 参数指定配置文件路径。

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
  "doh": [
    {
      "ip": "223.5.5.5",
      "alpn": "h2",
      "url": "https://dns.alidns.com/dns-query"
    }
  ]
}
```

### 配置项说明

- `hostname`: 服务器绑定的主机名，默认为 "0.0.0.0"
- `port`: 服务器监听的 TCP 端口号，默认为 8080
- `server_cert`: HTTPS 服务所需的 TLS 服务器证书文件路径
- `server_key`: HTTPS 服务所需的 TLS 私钥文件路径
- `username`: 访问代理服务器所需的用户名
- `password`: 访问代理服务器所需的密码
- `doh`: DOH 配置对象数组，每个对象包含以下字段：
  - `ip`: DOH 服务器 IP 地址，支持 ipv4 和 ipv6 地址
  - `alpn`: DOH ALPN 协议，支持 h2 和 h3 协议
  - `url`: DOH 服务器 URL，支持 http 和 https 协议

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
"/root/http-proxy-go-server/main" -dohurl "https://************" -dohip  "************" -port 58888 -username admin -password "*******************"  -server_cert "************"  -server_key "************"
```

```bash
# 使用配置文件的示例
go run main.go -config config.json
```

## 使用 curl 测试

```
curl -x http://127.0.0.1:8080 http://www.baidu.com
```

## WebSocket 代理支持任务总结

### 项目目标

为 HTTP 代理服务器添加 WebSocket 协议作为上游代理的支持，包括：

1. 命令行参数支持 WebSocket 上游代理
2. 配置文件支持 WebSocket 上游代理类型
3. 集成 socks5-websocket-proxy-golang 客户端

### 已完成任务

#### ✅ 1. 理解现有架构

- 分析 main.go 中的代理选择逻辑
- 理解 ProxySelector 函数的工作原理
- 分析 transportConfigurations 的构建过程
- 理解 UpStream 结构体和相关配置

#### ✅ 2. 设计 WebSocket 代理支持

- 定义 WebSocket 上游代理的配置结构
- 设计 WebSocket 代理客户端接口
- 规划与现有代理系统的集成方案

#### ✅ 3. 修改配置结构

- 扩展 UpStream 结构体，添加 WebSocket 相关字段
- 添加 WS_PROXY、WS_USERNAME、WS_PASSWORD 字段
- 更新 Config 结构体以支持新配置
- 修改配置文件解析逻辑

#### ✅ 4. 添加命令行参数支持

- 添加 `-upstream-type` 参数
- 添加 `-upstream-address` 参数
- 添加 `-upstream-username` 参数
- 添加 `-upstream-password` 参数
- 修改 main 函数中的参数解析逻辑

#### ✅ 5. 集成 WebSocket 客户端

- 添加 socks5-websocket-proxy-golang 依赖
- 实现 WebSocket 代理客户端包装器
- 创建与标准库兼容的代理函数
- 实现 WebSocket 连接管理

#### ✅ 6. 修改代理选择逻辑

- 更新 ProxySelector 函数以支持 WebSocket 类型
- 修改 transportConfigurations 构建逻辑
- 实现 WebSocket 代理的 URL 解析
- 添加 WebSocket 代理的特殊处理逻辑

#### ✅ 7. 更新配置文件示例

- 修改 config.json 添加 WebSocket 上游示例
- 更新 README.md 文档说明新功能
- 添加使用示例和说明

#### ✅ 8. 在 simple.Handle 和 auth.Handle 函数中实现 WebSocket 代理支持

- 修改 simple.Handle 函数支持 WebSocket 代理连接
- 修改 auth.Handle 函数支持 WebSocket 代理连接
- 添加 WebSocket 代理相关的导入包
- 实现HTTP数据包解析获取目标服务器域名和端口
- 支持CONNECT方法和其他HTTP方法的解析
- 调用 socks5_websocket_proxy_golang_websocket.NewWebSocketClient(wsConfig)
- 将解析的参数传递给相应的代理函数

#### ✅ 9. 实现配置文件中用户名密码覆盖代理URL中的用户名密码功能

- 添加 overrideProxyURLCredentials 辅助函数
- 支持仅覆盖用户名、仅覆盖密码或同时覆盖两者的场景
- 在 SelectProxyURLWithCIDR 函数中集成覆盖功能
- 支持所有代理类型的用户名密码覆盖

### 技术要点

#### 核心文件修改

1. **main.go** - 添加命令行参数和主要逻辑
2. **config.json** - 更新配置示例
3. **README.md** - 更新文档
4. **simple.go** - 实现 WebSocket 代理支持
5. **auth.go** - 实现 WebSocket 代理支持

#### 关键技术实现

1. **代理类型识别** - 在 ProxySelector 中添加 WebSocket 类型判断
2. **WebSocket 客户端集成** - 使用 socks5-websocket-proxy-golang 库
3. **配置扩展** - 保持向后兼容的同时添加新功能
4. **连接管理** - 确保 WebSocket 连接的正确建立和关闭
5. **HTTP数据包解析** - 解析目标服务器域名和端口
6. **用户名密码覆盖** - 支持配置文件中的认证信息覆盖URL中的认证信息

### 使用示例

#### 命令行参数使用

```bash
# 使用WebSocket上游代理
go run main.go -upstream-type websocket -upstream-address ws://127.0.0.1:1081 -upstream-username user -upstream-password pass
```

#### 配置文件使用

```json
{
  "upstreams": {
    "websocket_proxy": {
      "type": "websocket",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "ws_proxy": "ws://127.0.0.1:1081",
      "ws_username": "user",
      "ws_password": "pass"
    }
  },
  "rules": [
    {
      "filter": "websocket_filter",
      "upstream": "websocket_proxy"
    }
  ],
  "filters": {
    "websocket_filter": {
      "patterns": ["*"]
    }
  }
}
```

### 新增命令行参数

- `-upstream-type`: 代理类型 (websocket)
- `-upstream-address`: 上游代理地址 (如 ws://127.0.0.1:1081)
- `-upstream-username`: 上游代理用户名
- `-upstream-password`: 上游代理密码

### 配置结构扩展

```go
type UpStream struct {

  TYPE        string   `json:"type"`
    HTTP_PROXY  string   `json:"http_proxy"`
    HTTPS_PROXY string   `json:"https_proxy"`
    BypassList  []string `json:"bypass_list"`
    // 新增WebSocket支持
    WS_PROXY    string   `json:"ws_proxy"`      // WebSocket代理地址
    WS_USERNAME string   `json:"ws_username"`   // WebSocket代理用户名
    WS_PASSWORD string   `json:"ws_password"`   // WebSocket代理密码
}
```

### 成功标准

1. **功能完整性**: 能够通过命令行参数指定 WebSocket 上游代理
2. **配置灵活性**: 能够通过配置文件配置 WebSocket 上游代理
3. **代理转发能力**: WebSocket 代理能够正常转发 HTTP/HTTPS 请求
4. **兼容性保证**: 保持现有功能的完整性和向后兼容性
5. **文档完备性**: 提供完整的文档和使用示例
6. **稳定性**: WebSocket 连接稳定，错误处理完善
7. **性能**: 连接建立和转发性能满足基本使用需求

### 注意事项

- Golang 标准库的 `http.Transport.Proxy` 不支持 WebSocket 协议，需要自定义实现
- WebSocket 连接需要特殊的握手协议和消息帧处理
- 需要考虑连接复用和池化管理以提高性能
- 确保 WebSocket 代理能够正确处理 HTTP/HTTPS 请求的转发
