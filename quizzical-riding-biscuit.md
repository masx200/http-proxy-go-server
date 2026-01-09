# Bug Fix Plan: HTTP Upstream Proxy Should Not Start Internal HTTP Proxy Server

## Problem Summary

When an HTTP upstream proxy is configured, the internal HTTP proxy server (on
127.88.236.251:15944) is being started unnecessarily. Instead, requests should
be forwarded directly to the upstream HTTP proxy server, similar to how SOCKS5
upstream proxies work in "SOCKS5 Direct Mode".

## Root Cause

Both [simple/simple.go](simple/simple.go:37-61) and
[auth/auth.go](auth/auth.go:43-67) only check for SOCKS5 upstream proxies to
decide whether to start the internal HTTP proxy server:

```go
useSocks5Directly = strings.HasPrefix(proxyURL.String(), "socks5://")
```

This logic does not check for HTTP/HTTPS upstream proxies, so the internal HTTP
proxy server is started even when it's not needed.

## Solution

Add HTTP upstream proxy detection to bypass the internal HTTP proxy server when
HTTP upstream proxies are configured, similar to the existing SOCKS5 Direct
Mode.

## Files to Modify

1. **[simple/simple.go](simple/simple.go)**: Lines 37-61
2. **[auth/auth.go](auth/auth.go)**: Lines 43-67

## Implementation Changes

### Change 1: simple/simple.go (lines 37-61)

**Current code:**

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

**New code:**

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

### Change 2: auth/auth.go (lines 43-67)

Apply the same changes as in simple/simple.go, with the only difference being
the log message and the `Http()` function call includes username/password
parameters.

**Current code:**

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

**New code:**

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

## Request Routing Logic

The existing request routing logic in both modules (lines 147-161 in simple.go
and lines 183-197 in auth.go) already handles this correctly:

- For **CONNECT requests**: Always connect directly to the target address (line
  148-149 in simple.go)
- For **HTTP requests**: When `httpUpstreamAddress == ""` (which happens when we
  bypass the internal HTTP proxy server), connect directly to the target address
  (lines 152-155 in simple.go)

The existing proxy usage condition at line 246 in simple.go and line 282 in
auth.go will handle the routing correctly:

```go
} else if proxyURL != nil && (method == "CONNECT" || (method != "CONNECT" && httpUpstreamAddress == "")) {
```

This means:

- For CONNECT requests with HTTP upstream: Use `connect.ConnectViaHttpProxy()`
  (lines 360-367 in simple.go)
- For HTTP requests with HTTP upstream: Direct connection via
  `dnscache.Proxy_net_DialCached()` (lines 371-385 in simple.go)

## Expected Behavior After Fix

When starting the proxy with an HTTP upstream configured:

```bash
./main -upstream-address http://192.168.31.245:58877 -upstream-type http -upstream-username=admin -upstream-password=***
```

**Before fix:**

```
Proxy server started on port [::]:57788
Random IP: 127.88.236.251
Random integer: 15944
Proxy server started on port 127.88.236.251:15944  <-- Unnecessary internal HTTP proxy server
```

**After fix:**

```
Proxy server started on port [::]:57788
HTTP upstream detected, will handle requests directly via HTTP proxy (bypassing internal HTTP proxy server)
HTTP upstream mode: bypassing internal HTTP proxy server for direct HTTP proxy routing
```

## Verification Steps

1. **Test HTTP upstream proxy configuration:**
   ```bash
   go build -o main.exe ./cmd/
   ./main -port 57788 -upstream-address http://192.168.31.245:58877 -upstream-type http -upstream-username=admin -upstream-password=***
   ```

2. **Verify no internal HTTP proxy server is started:**
   - Check logs for "HTTP upstream detected" message
   - Check logs for "HTTP upstream mode: bypassing internal HTTP proxy server"
     message
   - Verify there's NO "Proxy server started on port 127.x.x.x:xxxx" message
     after the main proxy server starts

3. **Test actual proxy functionality:**
   ```bash
   # Test HTTP request
   curl -x http://localhost:57788 http://example.com

   # Test HTTPS request (CONNECT)
   curl -x http://localhost:57788 https://example.com
   ```

4. **Verify traffic flows correctly:**
   - Requests should be forwarded directly to the upstream HTTP proxy at
     http://192.168.31.245:58877
   - No intermediate internal HTTP proxy server should be involved

5. **Test that SOCKS5 upstream still works:**
   ```bash
   ./main -port 57788 -upstream-address socks5://127.0.0.1:1080 -upstream-type socks socks5-username=user -socks5-password=pass
   ```
   - Verify SOCKS5 direct mode still works correctly

6. **Test that WebSocket upstream still works:**
   ```bash
   ./main -port 57788 -upstream-address ws://127.0.0.1:1081 -upstream-type websocket -ws-username=user -ws-password=pass
   ```
   - Verify internal HTTP proxy server IS started (WebSocket upstreams still
     need it)

## Summary

This fix adds HTTP upstream proxy detection to both
[simple/simple.go](simple/simple.go) and [auth/auth.go](auth/auth.go),
preventing the unnecessary startup of the internal HTTP proxy server when an
HTTP upstream proxy is configured. Requests will be forwarded directly to the
upstream HTTP proxy, improving performance and reducing resource usage,
consistent with how SOCKS5 upstream proxies are handled.

The changes are minimal and localized to the startup logic in both modules,
maintaining backward compatibility with all existing upstream proxy types
(SOCKS5, WebSocket, and now HTTP).
