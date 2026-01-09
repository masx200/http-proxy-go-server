# Bug 修复计划：HTTP 上游代理不应启动内部 HTTP 代理服务器

## 问题概述

当配置了 HTTP 上游代理时，内部 HTTP 代理服务器（位于
127.88.236.251:15944）被不必要地启动了。相反，请求应该直接转发到上游 HTTP
代理服务器，类似于 SOCKS5 上游代理在“SOCKS5 直连模式”下的工作方式。

## 根本原因

[simple/simple.go](simple/simple.go:37-61) 和 [auth/auth.go](auth/auth.go:43-67)
都只检查 SOCKS5 上游代理来决定是否启动内部 HTTP 代理服务器：

```go
useSocks5Directly = strings.HasPrefix(proxyURL.String(), "socks5://")
```

该逻辑不检查 HTTP/HTTPS 上游代理，因此即使不需要也会启动内部 HTTP 代理服务器。

## 解决方案

添加 HTTP 上游代理检测，以便在配置 HTTP 上游代理时绕过内部 HTTP
代理服务器，类似于现有的 SOCKS5 直连模式。

## 需修改的文件

1. **[simple/simple.go](simple/simple.go)**: 第 37-61 行
2. **[auth/auth.go](auth/auth.go)**: 第 43-67 行

## 实现变更

### 变更 1：simple/simple.go (第 37-61 行)

**当前代码：**

```go
// 检查是否使用SOCKS5上游代理
var useSocks5Directly bool
var upstreamAddress string

if Proxy != nil {
    // 创建一个测试请求来检查上游代理类型
    testReq, _ := http.NewRequest("GET", "http://test", nil)
    if proxyURL, err := Proxy(testReq); err == nil && proxyURL != nil {
        useSocks5Directly = strings.HasPrefix(proxyURL.String(), "socks5://")
        if useSocks5Directly {
            log.Printf("SOCKS5 upstream detected, will handle HTTP requests directly via SOCKS5")
        }
    }
}

// 只有在非SOCKS5上游时才启动HTTP代理服务器
if !useSocks5Directly {
    xh := http_server.GenerateRandomLoopbackIP()
    x1 := http_server.GenerateRandomIntPort()
    upstreamAddress = xh + ":" + fmt.Sprint(x1)
    go http_server.Http(xh, x1, proxyoptions, dnsCache, "", "", upstreamResolveIPs, Proxy, tranportConfigurations...)
    log.Printf("Started HTTP proxy server for upstream routing at %s", upstreamAddress)
} else {
    log.Printf("SOCKS5 upstream mode: bypassing HTTP proxy server for direct SOCKS5 routing")
}
```

**新代码：**

```go
// 检查是否使用SOCKS5或HTTP上游代理
var useSocks5Directly bool
var useHttpUpstreamDirectly bool
var upstreamAddress string

if Proxy != nil {
    // 创建一个测试请求来检查上游代理类型
    testReq, _ := http.NewRequest("GET", "http://test", nil)
    if proxyURL, err := Proxy(testReq); err == nil && proxyURL != nil {
        proxyScheme := proxyURL.String()
        useSocks5Directly = strings.HasPrefix(proxyScheme, "socks5://")
        useHttpUpstreamDirectly = strings.HasPrefix(proxyScheme, "http://") || strings.HasPrefix(proxyScheme, "https://")

        if useSocks5Directly {
            log.Printf("SOCKS5 upstream detected, will handle HTTP requests directly via SOCKS5")
        } else if useHttpUpstreamDirectly {
            log.Printf("HTTP upstream detected, will handle requests directly via HTTP proxy (bypassing internal HTTP proxy server)")
        }
    }
}

// 只有在非SOCKS5上游且非HTTP上游时才启动HTTP代理服务器
if !useSocks5Directly && !useHttpUpstreamDirectly {
    xh := http_server.GenerateRandomLoopbackIP()
    x1 := http_server.GenerateRandomIntPort()
    upstreamAddress = xh + ":" + fmt.Sprint(x1)
    go http_server.Http(xh, x1, proxyoptions, dnsCache, "", "", upstreamResolveIPs, Proxy, tranportConfigurations...)
    log.Printf("Started HTTP proxy server for upstream routing at %s", upstreamAddress)
} else {
    if useSocks5Directly {
        log.Printf("SOCKS5 upstream mode: bypassing HTTP proxy server for direct SOCKS5 routing")
    } else if useHttpUpstreamDirectly {
        log.Printf("HTTP upstream mode: bypassing internal HTTP proxy server for direct HTTP proxy routing")
    }
}
```

### 变更 2：auth/auth.go (第 43-67 行)

应用与 simple/simple.go 中相同的更改，唯一的区别是日志消息以及 `Http()`
函数调用包含用户名/密码参数。

**当前代码：**

```go
// 检查是否使用SOCKS5上游代理
var useSocks5Directly bool
var upstreamAddress string

if Proxy != nil {
    // 创建一个测试请求来检查上游代理类型
    testReq, _ := http.NewRequest("GET", "http://test", nil)
    if proxyURL, err := Proxy(testReq); err == nil && proxyURL != nil {
        useSocks5Directly = strings.HasPrefix(proxyURL.String(), "socks5://")
        if useSocks5Directly {
            log.Printf("SOCKS5 upstream detected, will handle HTTP requests directly via SOCKS5")
        }
    }
}

// 只有在非SOCKS5上游时才启动HTTP代理服务器
if !useSocks5Directly {
    xh := http_server.GenerateRandomLoopbackIP()
    x1 := http_server.GenerateRandomIntPort()
    upstreamAddress = xh + ":" + fmt.Sprint(rune(x1))
    go http_server.Http(xh, x1, proxyoptions, dnsCache, username, password, upstreamResolveIPs, Proxy, tranportConfigurations...)
    log.Printf("Started HTTP proxy server for upstream routing at %s", upstreamAddress)
} else {
    log.Printf("SOCKS5 upstream mode: bypassing HTTP proxy server for direct SOCKS5 routing")
}
```

**新代码：**

```go
// 检查是否使用SOCKS5或HTTP上游代理
var useSocks5Directly bool
var useHttpUpstreamDirectly bool
var upstreamAddress string

if Proxy != nil {
    // 创建一个测试请求来检查上游代理类型
    testReq, _ := http.NewRequest("GET", "http://test", nil)
    if proxyURL, err := Proxy(testReq); err == nil && proxyURL != nil {
        proxyScheme := proxyURL.String()
        useSocks5Directly = strings.HasPrefix(proxyScheme, "socks5://")
        useHttpUpstreamDirectly = strings.HasPrefix(proxyScheme, "http://") || strings.HasPrefix(proxyScheme, "https://")

        if useSocks5Directly {
            log.Printf("SOCKS5 upstream detected, will handle HTTP requests directly via SOCKS5")
        } else if useHttpUpstreamDirectly {
            log.Printf("HTTP upstream detected, will handle requests directly via HTTP proxy (bypassing internal HTTP proxy server)")
        }
    }
}

// 只有在非SOCKS5上游且非HTTP上游时才启动HTTP代理服务器
if !useSocks5Directly && !useHttpUpstreamDirectly {
    xh := http_server.GenerateRandomLoopbackIP()
    x1 := http_server.GenerateRandomIntPort()
    upstreamAddress = xh + ":" + fmt.Sprint(rune(x1))
    go http_server.Http(xh, x1, proxyoptions, dnsCache, username, password, upstreamResolveIPs, Proxy, tranportConfigurations...)
    log.Printf("Started HTTP proxy server for upstream routing at %s", upstreamAddress)
} else {
    if useSocks5Directly {
        log.Printf("SOCKS5 upstream mode: bypassing HTTP proxy server for direct SOCKS5 routing")
    } else if useHttpUpstreamDirectly {
        log.Printf("HTTP upstream mode: bypassing internal HTTP proxy server for direct HTTP proxy routing")
    }
}
```

## 请求路由逻辑

两个模块中现有的请求路由逻辑（simple.go 中的第 147-161 行和 auth.go 中的第
183-197 行）已经正确处理了这一点：

- 对于 **CONNECT 请求**：始终直接连接到目标地址（simple.go 中的第 148-149 行）
- 对于 **HTTP 请求**：当 `httpUpstreamAddress == ""` 时（当我们绕过内部 HTTP
  代理服务器时会发生这种情况），直接连接到目标地址（simple.go 中的第 152-155
  行）

simple.go 第 246 行和 auth.go 第 282 行现有的代理使用条件将正确处理路由：

```go
} else if proxyURL != nil && (method == "CONNECT" || (method != "CONNECT" && httpUpstreamAddress == "")) {
```

这意味着：

- 对于带有 HTTP 上游的 CONNECT 请求：使用
  `connect.ConnectViaHttpProxy()`（simple.go 中的第 360-367 行）
- 对于带有 HTTP 上游的 HTTP 请求：通过 `dnscache.Proxy_net_DialCached()`
  直接连接（simple.go 中的第 371-385 行）

## 修复后的预期行为

当配置了 HTTP 上游并启动代理时：

```bash
./main -upstream-address http://192.168.31.245:58877 -upstream-type http -upstream-username=admin -upstream-password=***
```

**修复前：**

```
Proxy server started on port [::]:57788
Random IP: 127.88.236.251
Random integer: 15944
Proxy server started on port 127.88.236.251:15944  <-- 不必要的内部 HTTP 代理服务器
```

**修复后：**

```
Proxy server started on port [::]:57788
HTTP upstream detected, will handle requests directly via HTTP proxy (bypassing internal HTTP proxy server)
HTTP upstream mode: bypassing internal HTTP proxy server for direct HTTP proxy routing
```

## 验证步骤

1. **测试 HTTP 上游代理配置：**
   ```bash
   go build -o main.exe ./cmd/
   ./main -port 57788 -upstream-address http://192.168.31.245:58877 -upstream-type http -upstream-username=admin -upstream-password=***
   ```

2. **验证没有启动内部 HTTP 代理服务器：**
   - 检查日志中是否有“HTTP upstream detected”消息
   - 检查日志中是否有“HTTP upstream mode: bypassing internal HTTP proxy
     server”消息
   - 验证主代理服务器启动后**没有**“Proxy server started on port
     127.x.x.x:xxxx”消息

3. **测试实际代理功能：**
   ```bash
   # 测试 HTTP 请求
   curl -x http://localhost:57788 http://example.com

   # 测试 HTTPS 请求 (CONNECT)
   curl -x http://localhost:57788 https://example.com
   ```

4. **验证流量流向正确：**
   - 请求应直接转发到位于 http://192.168.31.245:58877 的上游 HTTP 代理
   - 不应涉及中间的内部 HTTP 代理服务器

5. **测试 SOCKS5 上游仍然有效：**
   ```bash
   ./main -port 57788 -upstream-address socks5://127.0.0.1:1080 -upstream-type socks socks5-username=user -socks5-password=pass
   ```
   - 验证 SOCKS5 直连模式仍然正常工作

6. **测试 WebSocket 上游仍然有效：**
   ```bash
   ./main -port 57788 -upstream-address ws://127.0.0.1:1081 -upstream-type websocket -ws-username=user -ws-password=pass
   ```
   - 验证内部 HTTP 代理服务器**已**启动（WebSocket 上游仍然需要它）

## 总结

此修复向 [simple/simple.go](simple/simple.go) 和 [auth/auth.go](auth/auth.go)
添加了 HTTP 上游代理检测，防止在配置 HTTP 上游代理时不必要地启动内部 HTTP
代理服务器。请求将直接转发到上游 HTTP 代理，从而提高性能并减少资源使用，这与处理
SOCKS5 上游代理的方式一致。

更改是微小的，且仅限于两个模块中的启动逻辑，保持了与所有现有上游代理类型（SOCKS5、WebSocket
以及现在的 HTTP）的向后兼容性。
