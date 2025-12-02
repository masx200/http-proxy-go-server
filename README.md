# http-proxy-go-server

`http-proxy-go-server` 是一个基于 Golang（Go 语言）实现的 HTTP
代理服务器。这个服务器允许你通过指定的 IP 地址和端口作为 HTTP
代理来转发网络请求。从提供的命令行参数来看，它的主要功能和配置选项如下：

## 命令行参数

| 参数                    | 类型   | 默认值             | 描述                                    |
| ----------------------- | ------ | ------------------ | --------------------------------------- |
| `-config`               | string | -                  | JSON配置文件路径                        |
| `-hostname`             | string | `0.0.0.0`          | 服务器绑定的主机名                      |
| `-port`                 | int    | `8080`             | TCP监听端口                             |
| `-username`             | string | -                  | 代理服务器用户名                        |
| `-password`             | string | -                  | 代理服务器密码                          |
| `-server_cert`          | string | -                  | TLS服务器证书文件路径                   |
| `-server_key`           | string | -                  | TLS服务器私钥文件路径                   |
| `-dohurl`               | value  | -                  | DOH服务器URL（可重复）                  |
| `-dohip`                | value  | -                  | DOH服务器IP地址（可重复）               |
| `-dohalpn`              | value  | -                  | DOH ALPN协议（可重复，支持h2和h3）      |
| `-upstream-type`        | string | -                  | 上游代理类型（websocket、socks5、http） |
| `-upstream-address`     | string | -                  | 上游代理地址                            |
| `-upstream-username`    | string | -                  | 上游代理用户名                          |
| `-upstream-password`    | string | -                  | 上游代理密码                            |
| `-upstream-resolve-ips` | bool   | `false`            | 解析上游代理域名为IP地址以绕过DNS污染   |
| `-cache-enabled`        | bool   | `true`             | 启用DNS缓存                             |
| `-cache-file`           | string | `./dns_cache.json` | DNS缓存文件路径                         |
| `-cache-ttl`            | string | `10m`              | DNS缓存TTL（生存时间）                  |
| `-cache-save-interval`  | string | `30s`              | DNS缓存全量保存间隔                     |
| `-cache-aof-enabled`    | bool   | `true`             | 启用DNS缓存AOF（增量持久化）            |
| `-cache-aof-file`       | string | `./dns_cache.aof`  | DNS缓存AOF文件路径                      |
| `-cache-aof-interval`   | string | `1s`               | DNS缓存AOF增量保存间隔                  |

1. `-config string`：指定 JSON 配置文件路径，可以通过配置文件设置所有参数。

2. `-hostname string`：设置服务器绑定的主机名，默认为
   "0.0.0.0"，表示服务器将监听所有可用的网络接口。

3. `-password string`：设置访问代理服务器所需的密码，用于基本身份验证。

4. `-port int`：设置服务器监听的 TCP 端口号，默认为 8080。

5. `-server_cert string`：设置 HTTPS 服务所需的 TLS
   服务器证书文件路径。如果提供了此选项，服务器将以安全模式运行（HTTPS）。

6. `-server_key string`：设置 HTTPS 服务所需的 TLS
   私钥文件路径，与服务器证书配套使用。

7. `-upstream-address string`：设置上游代理地址，支持 WebSocket、SOCKS5 和 HTTP
   协议，例如：

   - WebSocket: `ws://127.0.0.1:1081`
   - SOCKS5: `socks5://127.0.0.1:1080`
   - HTTP: `http://127.0.0.1:8080`

8. `-upstream-password string`：设置上游代理的密码。

9. `-upstream-type string`：设置上游代理类型，支持 `websocket`、`socks5` 和
   `http` 三种类型。

10. `-upstream-username string`：设置上游代理的用户名。

11. `-username string`：设置访问代理服务器所需的用户名，同样用于基本身份验证。

12. `-upstream-resolve-ips`：启用或禁用上游代理域名解析为IP地址功能，默认为禁用。启用后，系统会在连接上游代理之前先解析其域名为IP地址，然后依次尝试连接每个解析出的IP地址，直到连接成功为止。这对于解决上游代理存在DNS污染的情况非常有用。

13. `-cache-enabled`：启用或禁用DNS缓存功能，默认为启用。启用后可以显著提高DNS解析性能并减少对外部DNS服务器的请求次数。

13. `-cache-file string`：指定DNS缓存文件的存储路径，默认为
    "./dns_cache.json"。缓存会在程序启动时自动加载，在运行时定期保存，并在程序关闭时保存最新状态。

14. `-cache-ttl string`：设置DNS缓存的TTL（生存时间），默认为
    "10m"（10分钟）。支持的时间格式包括：5m、10m、1h 等。

15. `-cache-save-interval string`：设置DNS缓存的自动全量保存间隔，默认为
    "30s"（30秒）。系统会定期将完整缓存保存到文件中，以防止数据丢失。

16. `-cache-aof-enabled`：启用或禁用DNS缓存AOF（Append Only
    File）增量持久化功能，默认为启用。AOF模式可以实现更频繁的数据保存，提高数据安全性。

17. `-cache-aof-file string`：指定DNS缓存AOF文件的存储路径，默认为
    "./dns_cache.aof"。AOF文件采用JSONL（JSON
    Lines）格式，记录所有的DNS查询操作。

18. `-cache-aof-interval string`：设置DNS缓存AOF的增量保存间隔，默认为
    "1s"（1秒）。系统会以指定间隔将DNS查询操作追加到AOF文件中，实现近乎实时的数据持久化。

总结来说，`http-proxy-go-server` 提供了一个功能丰富的代理服务器，支持：

- 基本认证和 TLS 加密
- DOH (DNS over HTTPS) 支持
- WebSocket、SOCKS5 和 HTTP 上游代理
- **DNS 缓存功能** - 提高DNS解析性能，减少外部DNS请求
- **AOF 增量持久化** - Redis风格的增量日志，实现秒级数据持久化
- 灵活的配置方式（命令行参数和 JSON 配置文件）
  用户可以根据需要调整监听地址、端口、认证凭据、上游代理以及是否启用加密通信等配置项。

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
  "upstream_resolve_ips": false,
  "doh": [
    {
      "ip": "223.5.5.5",
      "alpn": "h2",
      "url": "https://dns.alidns.com/dns-query"
    }
  ],
  "dns_cache": {
    "enabled": true,
    "file": "./dns_cache.json",
    "ttl": "10m",
    "save_interval": "30s",
    "aof_enabled": true,
    "aof_file": "./dns_cache.aof",
    "aof_interval": "1s"
  }
}
```

### 配置项说明

- `hostname`: 服务器绑定的主机名，默认为 "0.0.0.0"
- `port`: 服务器监听的 TCP 端口号，默认为 8080
- `server_cert`: HTTPS 服务所需的 TLS 服务器证书文件路径
- `server_key`: HTTPS 服务所需的 TLS 私钥文件路径
- `username`: 访问代理服务器所需的用户名
- `password`: 访问代理服务器所需的密码
- `upstream_resolve_ips`: 是否启用上游代理域名解析为IP地址功能，默认为 false。启用后会在连接上游代理之前先解析其域名为IP地址，然后依次尝试连接每个解析出的IP地址，直到连接成功为止。这对于解决上游代理存在DNS污染的情况非常有用。
- `doh`: DOH 配置对象数组，每个对象包含以下字段：
  - `ip`: DOH 服务器 IP 地址，支持 ipv4 和 ipv6 地址
  - `alpn`: DOH ALPN 协议，支持 h2 和 h3 协议
  - `url`: DOH 服务器 URL，支持 http 和 https 协议
- `dns_cache`: DNS 缓存配置对象，包含以下字段：
  - `enabled`: 是否启用DNS缓存，默认为 true
  - `file`: DNS缓存文件路径，默认为 "./dns_cache.json"
  - `ttl`: DNS缓存TTL（生存时间），默认为 "10m"
  - `save_interval`: DNS缓存全量保存间隔，默认为 "30s"
  - `aof_enabled`: 是否启用AOF增量持久化，默认为 true
  - `aof_file`: AOF文件路径，默认为 "./dns_cache.aof"
  - `aof_interval`: AOF增量保存间隔，默认为 "1s"

### 使用配置文件

```bash
# 使用配置文件启动服务器
go run -v ./cmd/ -config config.json

# 配置文件和命令行参数可以混合使用，命令行参数会覆盖配置文件中的对应值
go run -v ./cmd/ -config config.json -port 9090 -username admin
```

### 配置文件优先级

配置文件中的参数会被用作默认值，但命令行参数会覆盖配置文件中的对应值。这样可以灵活地在基础配置上进行个性化调整。

## example

```bash
"/root/http-proxy-go-server/main" -dohurl "https://************" -dohip  "************" -port 58888 -username admin -password "*******************"  -server_cert "************"  -server_key "************"
```

```bash
# 使用配置文件的示例
go run -v ./cmd/ -config config.json
```

```bash
# 使用上游IP解析功能的示例（解决DNS污染问题）
go run -v ./cmd/ -upstream-resolve-ips=true -upstream-type http -upstream-address http://proxy.example.com:8080
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
- 实现 HTTP 数据包解析获取目标服务器域名和端口
- 支持 CONNECT 方法和其他 HTTP 方法的解析
- 调用 socks5_websocket_proxy_golang_websocket.NewWebSocketClient(wsConfig)
- 将解析的参数传递给相应的代理函数

#### ✅ 9. 实现配置文件中用户名密码覆盖代理 URL 中的用户名密码功能

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
5. **HTTP 数据包解析** - 解析目标服务器域名和端口
6. **用户名密码覆盖** - 支持配置文件中的认证信息覆盖 URL 中的认证信息

### 使用示例

#### 命令行参数使用

```bash
# 使用WebSocket上游代理
go run -v ./cmd/ -upstream-type websocket -upstream-address ws://127.0.0.1:1081 -upstream-username user -upstream-password pass
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

- `-upstream-type`: 代理类型 (websocket, socks5, http)
- `-upstream-address`: 上游代理地址 (如
  ws://127.0.0.1:1081、socks5://127.0.0.1:1080 或 http://127.0.0.1:8080)
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

## SOCKS5 代理支持

### 项目目标

为 HTTP 代理服务器添加 SOCKS5 协议作为上游代理的支持，包括：

1. 命令行参数支持 SOCKS5 上游代理
2. 配置文件支持 SOCKS5 上游代理类型
3. 集成 socks5-websocket-proxy-golang 客户端的 SOCKS5 功能

### 已完成任务

#### ✅ 1. 理解现有实现

- 分析 main.go 中的现有代理选择逻辑
- 理解 SOCKS5 代理客户端的集成方式
- 验证配置结构对 SOCKS5 的支持

#### ✅ 2. 验证配置结构支持

- 确认 UpStream 结构体支持 SOCKS5 代理配置
- 验证 SOCKS5_PROXY、SOCKS5_USERNAME、SOCKS5_PASSWORD 字段
- 确认配置文件解析逻辑正确处理 SOCKS5 配置

#### ✅ 3. 命令行参数支持

- 添加 `-socks5-proxy` 参数支持 SOCKS5 代理地址
- 添加 `-socks5-username` 参数支持 SOCKS5 用户名
- 添加 `-socks5-password` 参数支持 SOCKS5 密码
- 集成到 main 函数的参数解析逻辑

#### ✅ 4. 集成 SOCKS5 客户端

- 集成 socks5-websocket-proxy-golang 库的 SOCKS5 功能
- 实现 SOCKS5 代理客户端配置
- 创建与标准库兼容的 SOCKS5 代理函数
- 实现 SOCKS5 连接管理

#### ✅ 5. 修改代理选择逻辑

- 更新 ProxySelector 函数以支持 SOCKS5 类型
- 修改 transportConfigurations 构建逻辑
- 实现 SOCKS5 代理的 URL 解析
- 添加 SOCKS5 代理的特殊处理逻辑

#### ✅ 6. 配置示例和文档

- 修改 config.json 添加 SOCKS5 上游示例
- 更新 README.md 文档说明 SOCKS5 功能
- 添加 SOCKS5 使用示例和说明

#### ✅ 7. 用户名密码覆盖功能

- 实现 SOCKS5 代理的用户名密码覆盖功能
- 支持配置文件中的认证信息覆盖 URL 中的认证信息
- 集成到现有的覆盖逻辑中

### 技术要点

#### 核心文件修改

1. **main.go** - 添加 SOCKS5 命令行参数和主要逻辑
2. **config.json** - 更新配置示例
3. **README.md** - 更新文档
4. **simple.go** - 实现 SOCKS5 代理支持
5. **auth.go** - 实现 SOCKS5 代理支持

#### 关键技术实现

1. **代理类型识别** - 在 ProxySelector 中添加 SOCKS5 类型判断
2. **SOCKS5 客户端集成** - 使用 socks5-websocket-proxy-golang 库
3. **配置扩展** - 保持向后兼容的同时添加 SOCKS5 功能
4. **连接管理** - 确保 SOCKS5 连接的正确建立和关闭
5. **协议前缀处理** - 确保 ServerAddr
   包含正确的协议前缀（socks5://、tcp://、tls://、socks5s://）
6. **用户名密码覆盖** - 支持配置文件中的认证信息覆盖 URL 中的认证信息

### 使用示例

#### 命令行参数使用

```bash
# 使用SOCKS5上游代理
go run -v ./cmd/ -upstream-type socks5 -upstream-address socks5://127.0.0.1:1080 -upstream-username user -upstream-password pass

# 使用SOCKS5 over TLS上游代理
go run -v ./cmd/ -upstream-type socks5 -upstream-address socks5s://127.0.0.1:1080 -upstream-username user -upstream-password pass
```

#### 配置文件使用

```json
{
  "upstreams": {
    "socks5_proxy": {
      "type": "socks5",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "socks5_proxy": "socks5://127.0.0.1:1080",
      "socks5_username": "user",
      "socks5_password": "pass"
    }
  },
  "rules": [
    {
      "filter": "socks5_filter",
      "upstream": "socks5_proxy"
    }
  ],
  "filters": {
    "socks5_filter": {
      "patterns": ["*"]
    }
  }
}
```

### 新增命令行参数

- `-upstream-type`: 代理类型 (websocket, socks5, http)
- `-upstream-address`: 上游代理地址 (如 socks5://127.0.0.1:1080 或
  http://127.0.0.1:8080)
- `-upstream-username`: 上游代理用户名
- `-upstream-password`: 上游代理密码

### 配置结构扩展

```go
type UpStream struct {
    TYPE        string   `json:"type"`
    HTTP_PROXY  string   `json:"http_proxy"`
    HTTPS_PROXY string   `json:"https_proxy"`
    BypassList  []string `json:"bypass_list"`
    // WebSocket支持
    WS_PROXY    string   `json:"ws_proxy"`      // WebSocket代理地址
    WS_USERNAME string   `json:"ws_username"`   // WebSocket代理用户名
    WS_PASSWORD string   `json:"ws_password"`   // WebSocket代理密码
    // SOCKS5支持
    SOCKS5_PROXY    string   `json:"socks5_proxy"`      // SOCKS5代理地址
    SOCKS5_USERNAME string   `json:"socks5_username"`   // SOCKS5代理用户名
    SOCKS5_PASSWORD string   `json:"socks5_password"`   // SOCKS5代理密码
}
```

### 支持的协议前缀

SOCKS5 代理支持以下协议前缀：

- `socks5://`: 标准 SOCKS5 协议
- `tcp://`: TCP 连接（默认）
- `tls://`: TLS 加密的 SOCKS5 连接
- `socks5s://`: SOCKS5 over TLS（与 tls:// 等效）

如果未指定协议前缀，系统会自动添加 `tcp://` 前缀。

### 成功标准

1. **功能完整性**: 能够通过命令行参数指定 SOCKS5 上游代理
2. **配置灵活性**: 能够通过配置文件配置 SOCKS5 上游代理
3. **代理转发能力**: SOCKS5 代理能够正常转发 HTTP/HTTPS 请求
4. **兼容性保证**: 保持现有功能的完整性和向后兼容性
5. **文档完备性**: 提供完整的文档和使用示例
6. **稳定性**: SOCKS5 连接稳定，错误处理完善
7. **性能**: 连接建立和转发性能满足基本使用需求

### 注意事项

- SOCKS5 代理支持多种协议前缀，确保使用正确的格式
- SOCKS5 代理需要正确的用户名密码认证（如果需要）
- 需要考虑连接复用和池化管理以提高性能
- 确保 SOCKS5 代理能够正确处理 HTTP/HTTPS 请求的转发
- 协议前缀验证确保 ServerAddr 字段符合要求格式

## HTTP 代理支持

### 项目目标

为 HTTP 代理服务器添加 HTTP 协议作为上游代理的支持，包括：

1. 命令行参数支持 HTTP 上游代理
2. 配置文件支持 HTTP 上游代理类型
3. 集成标准 HTTP 代理客户端功能

### 已完成任务

#### ✅ 1. 理解现有实现

- 分析 main.go 中的现有代理选择逻辑
- 理解 HTTP 代理客户端的集成方式
- 验证配置结构对 HTTP 代理的支持

#### ✅ 2. 验证配置结构支持

- 确认 UpStream 结构体支持 HTTP 代理配置
- 验证 HTTP_PROXY、HTTPS_PROXY 字段
- 确认配置文件解析逻辑正确处理 HTTP 代理配置

#### ✅ 3. 命令行参数支持

- 添加 `-upstream-type` 参数支持 HTTP 代理类型
- 添加 `-upstream-address` 参数支持 HTTP 代理地址
- 添加 `-upstream-username` 参数支持 HTTP 用户名
- 添加 `-upstream-password` 参数支持 HTTP 密码
- 集成到 main 函数的参数解析逻辑

#### ✅ 4. 集成 HTTP 客户端

- 使用标准库的 HTTP 代理功能
- 实现 HTTP 代理客户端配置
- 创建与标准库兼容的 HTTP 代理函数
- 实现 HTTP 连接管理

#### ✅ 5. 修改代理选择逻辑

- 更新 ProxySelector 函数以支持 HTTP 类型
- 修改 transportConfigurations 构建逻辑
- 实现 HTTP 代理的 URL 解析
- 添加 HTTP 代理的特殊处理逻辑

#### ✅ 6. 配置示例和文档

- 修改 config.json 添加 HTTP 上游示例
- 更新 README.md 文档说明 HTTP 功能
- 添加 HTTP 使用示例和说明

#### ✅ 7. 用户名密码覆盖功能

- 实现 HTTP 代理的用户名密码覆盖功能
- 支持配置文件中的认证信息覆盖 URL 中的认证信息
- 集成到现有的覆盖逻辑中

### 技术要点

#### 核心文件修改

1. **main.go** - 添加 HTTP 命令行参数和主要逻辑
2. **config.json** - 更新配置示例
3. **README.md** - 更新文档
4. **simple.go** - 实现 HTTP 代理支持
5. **auth.go** - 实现 HTTP 代理支持

#### 关键技术实现

1. **代理类型识别** - 在 ProxySelector 中添加 HTTP 类型判断
2. **HTTP 客户端集成** - 使用标准库的 HTTP 代理功能
3. **配置扩展** - 保持向后兼容的同时添加 HTTP 功能
4. **连接管理** - 确保 HTTP 连接的正确建立和关闭
5. **URL 解析** - 确保 HTTP 代理地址正确解析
6. **用户名密码覆盖** - 支持配置文件中的认证信息覆盖 URL 中的认证信息

### 使用示例

#### 命令行参数使用

```bash
# 使用HTTP上游代理
go run -v ./cmd/ -upstream-type http -upstream-address http://127.0.0.1:8080 -upstream-username user -upstream-password pass

# 使用带认证的HTTP上游代理
go run -v ./cmd/ -upstream-type http -upstream-address http://user:pass@127.0.0.1:8080
```

#### 配置文件使用

```json
{
  "upstreams": {
    "http_proxy": {
      "type": "http",
      "http_proxy": "http://127.0.0.1:8080",
      "https_proxy": "http://127.0.0.1:8080",
      "bypass_list": []
    }
  },
  "rules": [
    {
      "filter": "http_filter",
      "upstream": "http_proxy"
    }
  ],
  "filters": {
    "http_filter": {
      "patterns": ["*"]
    }
  }
}
```

### 新增命令行参数

- `-upstream-type`: 代理类型 (websocket, socks5, http)
- `-upstream-address`: 上游代理地址 (如 http://127.0.0.1:8080)
- `-upstream-username`: 上游代理用户名
- `-upstream-password`: 上游代理密码

### 配置结构扩展

```go
type UpStream struct {
    TYPE        string   `json:"type"`
    HTTP_PROXY  string   `json:"http_proxy"`
    HTTPS_PROXY string   `json:"https_proxy"`
    BypassList  []string `json:"bypass_list"`
    // WebSocket支持
    WS_PROXY    string   `json:"ws_proxy"      // WebSocket代理地址
    WS_USERNAME string   `json:"ws_username"   // WebSocket代理用户名
    WS_PASSWORD string   `json:"ws_password"   // WebSocket代理密码
    // SOCKS5支持
    SOCKS5_PROXY    string   `json:"socks5_proxy"      // SOCKS5代理地址
    SOCKS5_USERNAME string   `json:"socks5_username"   // SOCKS5代理用户名
    SOCKS5_PASSWORD string   `json:"socks5_password"   // SOCKS5代理密码
}
```

### 支持的协议前缀

HTTP 代理支持以下协议前缀：

- `http://`: 标准 HTTP 协议
- `https://`: HTTPS 协议（如果上游代理支持）

### 成功标准

1. **功能完整性**: 能够通过命令行参数指定 HTTP 上游代理
2. **配置灵活性**: 能够通过配置文件配置 HTTP 上游代理
3. **代理转发能力**: HTTP 代理能够正常转发 HTTP/HTTPS 请求
4. **兼容性保证**: 保持现有功能的完整性和向后兼容性
5. **文档完备性**: 提供完整的文档和使用示例
6. **稳定性**: HTTP 连接稳定，错误处理完善
7. **性能**: 连接建立和转发性能满足基本使用需求

### 注意事项

- HTTP 代理支持标准的 HTTP 协议前缀
- HTTP 代理需要正确的用户名密码认证（如果需要）
- 需要考虑连接复用和池化管理以提高性能
- 确保 HTTP 代理能够正确处理 HTTP/HTTPS 请求的转发
- 用户名密码可以通过 URL 或命令行参数指定

## DNS 缓存功能

### 功能概述

`http-proxy-go-server`
现已集成智能DNS缓存功能，显著提高DNS解析性能并减少对外部DNS服务器的请求频率。该功能使用
PatrickMN/go-cache 库实现，具有以下特点：

- **内存缓存**：高速的内存DNS缓存，默认TTL为10分钟
- **文件持久化**：缓存自动保存到文件，程序重启后可恢复
- **原子操作**：使用临时文件确保缓存数据完整性
- **定期保存**：后台定期保存缓存，防止数据丢失
- **优雅关闭**：程序退出时自动保存最新缓存状态

### 核心特性

#### 1. 高性能缓存

- 支持 IPv4 和 IPv6 地址缓存
- 智能域名标准化处理
- 支持多种DNS记录类型（resolve, lookupip）
- 线程安全的并发操作

#### 2. 文件持久化

- JSON格式的缓存文件，便于调试和查看
- 自动过滤过期缓存项
- 原子写入操作，避免数据损坏
- 自动创建缓存目录

#### 3. 灵活配置

- 可通过命令行参数自定义缓存行为
- 支持完全禁用缓存功能
- 可调整TTL和保存间隔

### 命令行配置

```bash
# 启用DNS缓存（默认）
go run ./cmd/ -cache-enabled

# 自定义缓存文件路径
go run ./cmd/ -cache-file /var/cache/dns_cache.json

# 设置缓存TTL为5分钟
go run ./cmd/ -cache-ttl 5m

# 设置缓存保存间隔为1分钟
go run ./cmd/ -cache-save-interval 1m

# 禁用DNS缓存
go run ./cmd/ -cache-enabled=false
```

### 配置文件支持

当前版本的DNS缓存主要通过命令行参数配置，未来版本将添加JSON配置文件支持。

### 缓存文件格式

DNS缓存以JSON格式存储，示例如下：

```json
{
  "RESOLVE:example.com": {
    "value": "93.184.216.34",
    "expiration": "2024-01-15T10:30:00Z"
  },
  "LOOKUPIP:tcp:google.com": {
    "value": ["142.250.191.142", "142.250.191.78"],
    "expiration": "2024-01-15T10:30:00Z"
  }
}
```

### 性能优化

DNS缓存功能能够带来以下性能提升：

1. **减少DNS查询延迟**：缓存的域名解析几乎是瞬时响应
2. **降低网络负载**：减少对外部DNS服务器的请求
3. **提高连接建立速度**：更快的域名解析意味着更快的服务器连接
4. **改善DoH性能**：特别在使用DNS over HTTPS时效果明显

### 监控和调试

程序运行时会输出DNS缓存相关的日志信息：

```
DNS缓存已启用，文件: ./dns_cache.json, TTL: 10m0s
DNS cache hit for resolve: example.com
DNS cache set for lookupip: google.com (tcp) -> [142.250.191.142 142.250.191.78]
```

### 注意事项

1. **缓存失效**：域名IP地址变更时，需要等待TTL过期或重启程序
2. **文件权限**：确保程序对缓存文件和目录有读写权限
3. **磁盘空间**：长期运行可能产生较大的缓存文件，建议定期清理
4. **隐私考虑**：缓存文件包含DNS查询历史，注意文件安全性

## 上游IP解析功能

### 功能概述

`-upstream-resolve-ips` 参数提供了一个强大的功能来解决上游代理的DNS污染问题。当启用此功能时：

1. **域名解析**：系统会在连接上游代理之前，使用现有的DNS/DoH基础设施解析上游代理的域名为IP地址
2. **顺序连接**：如果域名解析出多个IP地址，系统会依次尝试连接每个IP地址，直到连接成功为止
3. **回退机制**：如果IP地址解析失败或所有IP地址都无法连接，系统会自动回退到传统的域名连接方式
4. **兼容性**：该功能与所有上游代理类型（HTTP、SOCKS5、WebSocket）完全兼容

### 使用场景

此功能特别适用于以下场景：
- 上游代理服务器存在DNS污染或被DNS劫持
- 上游代理域名被网络策略阻拦或解析错误
- 需要绕过DNS限制直接连接上游代理IP地址
- 上游代理有多个IP地址，需要自动故障转移

### 配置示例

#### 命令行参数使用
```bash
# 启用上游IP解析功能
go run -v ./cmd/ -upstream-resolve-ips=true -upstream-type http -upstream-address http://proxy.example.com:8080

# 与WebSocket上游代理配合使用
go run -v ./cmd/ -upstream-resolve-ips=true -upstream-type websocket -upstream-address ws://proxy.example.com:1081

# 与SOCKS5上游代理配合使用
go run -v ./cmd/ -upstream-resolve-ips=true -upstream-type socks5 -upstream-address socks5://proxy.example.com:1080
```

#### 配置文件使用
```json
{
  "upstream_resolve_ips": true,
  "upstreams": {
    "http_proxy": {
      "type": "http",
      "http_proxy": "http://proxy.example.com:8080",
      "https_proxy": "http://proxy.example.com:8080"
    }
  }
}
```

### 技术实现

该功能通过以下技术实现：
- 使用现有的`Proxy_net_DialContext`函数进行网络连接
- 集成DNS缓存和DoH解析基础设施
- 实现顺序连接和故障转移逻辑
- 保持与现有配置系统的完全兼容性
- 提供详细的连接日志和错误处理

### 性能优势

- **绕过DNS污染**：直接连接IP地址，避免被污染的DNS响应
- **自动故障转移**：当某个IP地址不可用时，自动尝试其他IP地址
- **连接优化**：优先尝试响应最快的IP地址
- **减少连接延迟**：避免DNS解析超时和重试

### 安全注意事项

- 确保上游代理IP地址的可靠性和信任度
- 定期验证IP地址的有效性
- 考虑IP地址变更对服务的影响
- 建议配合DoH使用以确保DNS解析的准确性

## 测试项目

```bash
go test -v -p 1 -parallel 1 ./...
```
