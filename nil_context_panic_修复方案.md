# nil Context Panic 错误诊断与修复方案

## 问题概述

在使用 `-upstream-resolve-ips` 功能时，程序出现 `panic: nil context`
错误，导致代理服务器崩溃。

## 错误分析

### 堆栈跟踪

```
panic: nil context

net.(*Dialer).DialContext(0xc0005435e0?, {0x0?, 0x0?}, {0x7ff65b07d6bd?, 0x11?}, {0xc000c15440?, 0xc00016c380?})
    C:/Program Files/Go/src/net/dial.go:527 +0x998
github.com/masx200/http-proxy-go-server/dnscache.proxy_net_DialWithResolver({0x0, 0x0}, {0x7ff65b07d6bd, 0x3}, {0xc000141470, 0x14}, {0xc00017b208, 0x9, 0x11}, 0x1, ...)
    D:/github/http-proxy-go-server/dnscache/caching_resolver.go:384 +0x18e6
github.com/masx200/http-proxy-go-server/dnscache.Proxy_net_DialCached({0x7ff65b07d6bd, 0x3}, {0xc000141470, 0x14}, {0xc00017b208, 0x9, 0x11}, 0x1, 0xc00016c380, {0xc000120088, ...})
    D:/github/http-proxy-go-server/dnscache/caching_resolver.go:350 +0x1d8
github.com/masx200/http-proxy-go-server/auth.Handle({0x7ff65b27a7f0, 0xc00008c870}, {0xc00000a158, 0x5}, {0xc000030270, 0x28}, {0xc000141470, 0x14}, {0xc00017b208, 0x9, ...}, ...)
    D:/github/http-proxy-go-server/auth/auth.go:248 +0xc8a
```

### 问题根源

1. **错误位置**: `dnscache/caching_resolver.go:384`
2. **调用链**: `auth.Handle()` → `dnscache.Proxy_net_DialCached()` →
   `proxy_net_DialWithResolver(nil, ...)` → `dialer.DialContext(nil, ...)`

### 详细问题

#### 1. 代码问题点

**问题代码位置：**

- `auth/auth.go:248`
- `simple/simple.go:221`
- `dnscache/caching_resolver.go:350, 384`

**具体问题：**

```go
// auth/auth.go:248
server, err = dnscache.Proxy_net_DialCached("tcp", upstreamAddress, proxyoptions, upstreamResolveIPs, dnsCache, Proxy,tranportConfigurations...)

// dnscache/caching_resolver.go:350
return proxy_net_DialWithResolver(nil, network, addr, proxyoptions, upstreamResolveIPs, dnsCache, CreateHostsAndDohResolverCached(proxyoptions, dnsCache, Proxy,tranportConfigurations...))

// dnscache/caching_resolver.go:384
connection, err1 := dialer.DialContext(ctx, network, newAddr)  // ctx is nil!
```

#### 2. 函数设计不一致

- `auth.Handle()` 和 `simple.Handle()` 没有 `context.Context` 参数
- `http.Handle()` 有 `context.Context` 参数并正确使用
  `Proxy_net_DialContextCached`
- `Proxy_net_DialCached` 内部使用 nil context，但在处理 upstreamResolveIPs
  时调用 `DialContext`

## 解决方案

### 方案一：最小修改（推荐）

在 `proxy_net_DialWithResolver` 函数中添加 context 保护：

```go
// dnscache/caching_resolver.go 第 364 行后添加
func proxy_net_DialWithResolver(ctx context.Context, network string, addr string, proxyoptions options.ProxyOptions, upstreamResolveIPs bool, dnsCache interface{}, resolver NameResolver, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) (net.Conn, error) {
    hostname, port, err := net.SplitHostPort(addr)
    if err != nil {
        return nil, err
    }

    // 重要：确保 context 不为 nil
    if ctx == nil {
        ctx = context.Background()
    }

    // 如果启用了上游IP解析功能，则使用新的解析逻辑
    if upstreamResolveIPs && len(proxyoptions) > 0 {
        // ... 其余代码保持不变
    }
}
```

### 方案二：完整修复（API 变更）

#### 步骤 1: 修改 Handle 函数签名

**auth/auth.go:**

```go
// 修改前
func Handle(client net.Conn, username, password string, httpUpstreamAddress string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) {

// 修改后  
func Handle(ctx context.Context, client net.Conn, username, password string, httpUpstreamAddress string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) {
```

**simple/simple.go:**

```go
// 修改前
func Handle(client net.Conn, httpUpstreamAddress string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) {

// 修改后
func Handle(ctx context.Context, client net.Conn, httpUpstreamAddress string, proxyoptions options.ProxyOptions, dnsCache *dnscache.DNSCache, upstreamResolveIPs bool, Proxy func(*http.Request) (*url.URL, error), tranportConfigurations ...func(*http.Transport) *http.Transport) {
```

#### 步骤 2: 修改调用处

**auth/auth.go:248** 和 **simple/simple.go:221**:

```go
// 修改前
server, err = dnscache.Proxy_net_DialCached("tcp", upstreamAddress, proxyoptions, upstreamResolveIPs, dnsCache, Proxy,tranportConfigurations...)

// 修改后
server, err = dnscache.Proxy_net_DialContextCached(ctx, "tcp", upstreamAddress, proxyoptions, dnsCache, upstreamResolveIPs, Proxy,tranportConfigurations...)
```

#### 步骤 3: 修改调用 Handle 的地方

**auth/auth.go:51:**

```go
// 修改前
go Handle(client, username, password, upstreamAddress, proxyoptions, dnsCache, upstreamResolveIPs, Proxy,tranportConfigurations...)

// 修改后
go Handle(context.Background(), client, username, password, upstreamAddress, proxyoptions, dnsCache, upstreamResolveIPs, Proxy,tranportConfigurations...)
```

**simple/simple.go 中的相应调用也需要修改**

## 临时解决方案

如果不想修改代码，可以临时禁用 `-upstream-resolve-ips` 功能：

```bash
# 启动时不使用 upstream-resolve-ips 参数
./http-proxy-go-server -auth=username:password -upstream-resolve-ips=false
```

## 建议

**推荐使用方案一**，因为：

1. 修改最小，风险最低
2. 向后兼容，不需要修改 API
3. 立即可用，不需要其他调用方修改
4. 符合 Go 语言的最佳实践

**方案二更适合长期维护**，但需要：

1. 更新所有调用 Handle 函数的地方
2. 可能需要更新依赖此 API 的其他代码
3. 进行充分的测试

## 测试建议

修复后应测试以下场景：

1. 启用 `-upstream-resolve-ips` 时的正常连接
2. 禁用 `-upstream-resolve-ips` 时的正常连接
3. 上游代理连接失败时的 fallback 机制
4. 各种代理协议（HTTP CONNECT, SOCKS5, WebSocket）

## 预防措施

建议在代码中添加 context 验证：

```go
func requireContext(ctx context.Context) context.Context {
    if ctx == nil {
        return context.Background()
    }
    return ctx
}
```

并在所有调用 `DialContext` 的地方使用此函数。
