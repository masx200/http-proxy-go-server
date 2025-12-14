# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.x.x] - 2025-12-15

### Fixed
- **SOCKS5 Upstream Proxy Bug** - Fixed critical issue where SOCKS5 upstream proxies were not working for HTTPS CONNECT requests
  - Root cause: The code was trying to send HTTP CONNECT requests to SOCKS5 proxies, but SOCKS5 protocol doesn't accept HTTP CONNECT requests
  - Added dedicated SOCKS5 CONNECT request handling logic in `simple.go` and `auth.go`
  - Fixed SOCKS5 URL parsing issue that caused "too many colons in address" error
  - Now properly extracts proxy host/port from URLs and constructs clean SOCKS5 server addresses

### Technical Details
- **File Changes**:
  - `simple/simple.go`: Added SOCKS5 proxy support for CONNECT requests
  - `auth/auth.go`: Added SOCKS5 proxy support for authenticated proxy mode
  - Added imports for `github.com/masx200/socks5-websocket-proxy-golang/pkg/socks5`
- **Implementation**:
  - Detection of SOCKS5 proxy URLs (starting with `socks5://`)
  - Proper URL parsing to separate credentials from server address
  - Use of SOCKS5 client library for direct target connections
  - Bidirectional data forwarding using `net.Pipe()`
- **Backward Compatibility**: Maintained full compatibility with HTTP and WebSocket proxy types

### Verification
- Tested with SOCKS5 proxy server on localhost:44444
- Successfully established HTTPS tunnels through SOCKS5 upstream
- Confirmed both authenticated and non-authenticated proxy modes work correctly
- Verified HTTP requests continue to work through existing transport configuration

### Before
```bash
curl -x http://127.0.0.1:8080 https://www.baidu.com
# Result: HTTP/1.1 502 Bad Gateway
# Error: failed to read proxy response: EOF
# SOCKS5 Error: Unsupported SOCKS version: [67]
```

### After
```bash
curl -x http://127.0.0.1:8080 https://www.baidu.com
# Result: HTTP/1.1 200 Connection established
# Success: Full HTTPS response received
# SOCKS5 Log: SOCKS5 connection handled successfully
```

## [Previous Versions]

### Features
- Multiple upstream proxy support (HTTP, WebSocket, SOCKS5)
- DNS over HTTPS (DoH) integration
- TLS proxy modes with authentication
- Flexible routing and filtering system
- DNS caching with AOF persistence
- IPv6 support

### Supported Protocols
- HTTP/HTTPS proxy
- WebSocket proxy
- SOCKS5 proxy
- DNS over HTTPS (DoH)
- DNS over TLS (DoT)
- DNS over QUIC (DoQ)