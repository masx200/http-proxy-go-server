

###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:47:37

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 53224

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 50760

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 51124, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:47:30 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11354468487813765212
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:47:30 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11354468487813765212


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 54732, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:47:30 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10687389693549337146
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:47:30 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10687389693549337146


```

### 📋 所有进程PID记录

所有进程PID: 53224, 50760, 51124, 54732



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:47:37

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :8080`

📋 WebSocket服务器进程PID: 53224

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:8080`

📋 SOCKS5服务器进程PID: 50760

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 51124, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:47:30 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11354468487813765212
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:47:30 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11354468487813765212


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 54732, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:47:30 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10687389693549337146
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:47:30 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10687389693549337146


```

### 📋 所有进程PID记录

所有进程PID: 53224, 50760, 51124, 54732

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:47:38 main.go:71: 启动websocket服务端，监听地址: :8080
2025/09/12 00:47:38 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :8080
2025/09/12 00:47:38 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:47:38 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:47:38 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:47:38 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:47:40 server.go:90: url /
2025/09/12 00:47:40 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[wKpsccq+7zuHY7YWQmsMCg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.180.192.61] X-Proxy-Target-Port:[54324]]
2025/09/12 00:47:40 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39107 at 2025-09-12 00:47:40
2025/09/12 00:47:40 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39107
2025/09/12 00:47:40 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.180.192.61', targetPort: 54324
2025/09/12 00:47:40 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:47:40 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39107
2025/09/12 00:47:40 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.180.192.61:54324 from [::1]:39107
2025/09/12 00:47:40 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.180.192.61:54324 (timeout: 30s)
2025/09/12 00:47:40 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.180.192.61:54324
2025/09/12 00:47:40 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:47:40 server.go:90: url /
2025/09/12 00:47:40 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[IcTiDrPllqzLaGq1+Si+Lw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:47:40 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39109 at 2025-09-12 00:47:40
2025/09/12 00:47:40 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39109
2025/09/12 00:47:40 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 80
2025/09/12 00:47:40 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:47:40 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39109
2025/09/12 00:47:40 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:80 from [::1]:39109
2025/09/12 00:47:40 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:80 (timeout: 30s)
2025/09/12 00:47:40 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:80
2025/09/12 00:47:40 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:47:42 server.go:90: url /
2025/09/12 00:47:42 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[HDyoT82GqHR626/uMMX7og==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.180.192.61] X-Proxy-Target-Port:[54324]]
2025/09/12 00:47:42 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39114 at 2025-09-12 00:47:42
2025/09/12 00:47:42 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39114
2025/09/12 00:47:42 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.180.192.61', targetPort: 54324
2025/09/12 00:47:42 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:47:42 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39114
2025/09/12 00:47:42 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.180.192.61:54324 from [::1]:39114
2025/09/12 00:47:42 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.180.192.61:54324 (timeout: 30s)
2025/09/12 00:47:42 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.180.192.61:54324
2025/09/12 00:47:42 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:47:42 server.go:90: url /
2025/09/12 00:47:42 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[PqeVxrFVOF1J28S5Z3N4sg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:47:42 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39117 at 2025-09-12 00:47:42
2025/09/12 00:47:42 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39117
2025/09/12 00:47:42 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 443
2025/09/12 00:47:42 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:47:42 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39117
2025/09/12 00:47:42 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:443 from [::1]:39117
2025/09/12 00:47:42 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:443 (timeout: 30s)
2025/09/12 00:47:42 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:443
2025/09/12 00:47:42 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:47:39 main.go:518: 代理服务器启动中...
2025/09/12 00:47:39 main.go:562: dohalpn: []
2025/09/12 00:47:39 main.go:564: hostname: 0.0.0.0
2025/09/12 00:47:39 main.go:566: port: 10810
2025/09/12 00:47:39 main.go:568: server_cert: 
2025/09/12 00:47:39 main.go:570: server_key: 
2025/09/12 00:47:39 main.go:572: username: 
2025/09/12 00:47:39 main.go:574: password: 
2025/09/12 00:47:39 main.go:576: dohurl: []
2025/09/12 00:47:39 main.go:578: dohip: []
2025/09/12 00:47:39 main.go:579: upstream-type: websocket
2025/09/12 00:47:39 main.go:580: upstream-address: ws://localhost:8080
2025/09/12 00:47:39 main.go:581: upstream-username: 
2025/09/12 00:47:39 main.go:582: upstream-password: 
2025/09/12 00:47:39 main.go:639: WebSocket代理配置已添加
2025/09/12 00:47:39 main.go:854: {
  "hostname": "",
  "port": 0,
  "server_cert": "",
  "server_key": "",
  "username": "",
  "password": "",
  "doh": null,
  "upstreams": {
    "websocket_upstream": {
      "type": "websocket",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "ws_proxy": "ws://localhost:8080",
      "ws_username": "",
      "ws_password": "",
      "http_username": "",
      "http_password": "",
      "socks5_proxy": "",
      "socks5_username": "",
      "socks5_password": ""
    }
  },
  "rules": [
    {
      "filter": "websocket_filter",
      "upstream": "websocket_upstream"
    }
  ],
  "filters": {
    "websocket_filter": {
      "patterns": [
        "*"
      ]
    }
  }
}
2025/09/12 00:47:39 simple.go:31: Proxy server started on port [::]:10810
2025/09/12 00:47:39 http.go:372: Random IP: 127.180.192.61
2025/09/12 00:47:39 http.go:390: Random integer: 54324
2025/09/12 00:47:39 http.go:342: Proxy server started on port 127.180.192.61:54324
2025/09/12 00:47:40 simple.go:57: remote addr: [::1]:39106
2025/09/12 00:47:40 simple.go:79: GET http://www.baidu.com/ HTTP/1.1
2025/09/12 00:47:40 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:47:40 simple.go:117: address:www.baidu.com:80
2025/09/12 00:47:40 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.180.192.61:54324
2025/09/12 00:47:40 main.go:791: ProxySelector 127.180.192.61:54324
2025/09/12 00:47:40 main.go:797: 选择的代理 URL: ws://localhost:8080
2025/09/12 00:47:40 simple.go:179: WebSocket Config Details:
2025/09/12 00:47:40 simple.go:180: host, portNum 127.180.192.61 54324
2025/09/12 00:47:40 simple.go:181:   Username: 
2025/09/12 00:47:40 simple.go:182:   Password: 
2025/09/12 00:47:40 simple.go:183:   ServerAddr: ws://localhost:8080
2025/09/12 00:47:40 simple.go:184:   Protocol: websocket
2025/09/12 00:47:40 simple.go:185:   Timeout: 30s
2025/09/12 00:47:40 client.go:98: url: ws://localhost:8080
2025/09/12 00:47:40 client.go:99: headers: map[X-Proxy-Target-Host:[127.180.192.61] X-Proxy-Target-Port:[54324]]
2025/09/12 00:47:40 client.go:110: url: http://localhost:8080
2025/09/12 00:47:40 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[2IAb0wAV/6h3C385KPrYpfFF7e8=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:47:40 simple.go:209: WebSocket代理连接成功：127.180.192.61:54324
2025/09/12 00:47:40 simple.go:248: clienthost: ::1
2025/09/12 00:47:40 simple.go:249: clientport: 39106
2025/09/12 00:47:40 simple.go:278: simple Handle header:
2025/09/12 00:47:40 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:47:40 simple.go:283: GET / HTTP/1.1
2025/09/12 00:47:40 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:47:40 simple.go:294: Host: www.baidu.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
2025/09/12 00:47:40 http.go:90: method: GET
2025/09/12 00:47:40 http.go:91: url: /
2025/09/12 00:47:40 http.go:92: host: www.baidu.com
2025/09/12 00:47:40 http.go:93: proxyHandler header:
2025/09/12 00:47:40 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:47:40 http.go:121: clientport: 39108
2025/09/12 00:47:40 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.180.192.61:54324;host=www.baidu.com;proto=http
2025/09/12 00:47:40 http.go:132: proxyHandler User-Agent : Go-http-client/1.1
2025/09/12 00:47:40 http.go:132: proxyHandler Accept-Encoding : gzip
2025/09/12 00:47:40 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.180.192.61:54324;host=www.baidu.com;proto=http
2025/09/12 00:47:40 http.go:137: forwardedByList: [{[::1]:10810} {127.180.192.61:54324}]
2025/09/12 00:47:40 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:47:40 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:47:40 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:47:40 main.go:797: 选择的代理 URL: ws://localhost:8080
2025/09/12 00:47:40 http.go:250: 使用代理：ws://localhost:8080
2025/09/12 00:47:40 http.go:259: 已经修改了代理为websocket ws://localhost:8080
2025/09/12 00:47:40 http.go:262: 使用代理：ws://localhost:8080
2025/09/12 00:47:40 http.go:264: network,addr tcp www.baidu.com:80
2025/09/12 00:47:40 http.go:466: WebSocket Config Details:
2025/09/12 00:47:40 http.go:467: host, portNum www.baidu.com 80
2025/09/12 00:47:40 http.go:468:   Username: 
2025/09/12 00:47:40 http.go:469:   Password: 
2025/09/12 00:47:40 http.go:470:   ServerAddr: ws://localhost:8080
2025/09/12 00:47:40 http.go:471:   Protocol: websocket
2025/09/12 00:47:40 http.go:472:   Timeout: 30s
2025/09/12 00:47:40 http.go:476: host, portNum www.baidu.com 80
2025/09/12 00:47:40 client.go:98: url: ws://localhost:8080
2025/09/12 00:47:40 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:47:40 client.go:110: url: http://localhost:8080
2025/09/12 00:47:40 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[8M8QzQg1t3l1H76gzsUAxVtN1Xs=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
[GIN] 2025/09/12 - 00:47:40 | 200 |     49.3889ms |       127.0.0.1 | GET      "/"
2025/09/12 00:47:42 simple.go:57: remote addr: [::1]:39113
2025/09/12 00:47:42 simple.go:79: HEAD http://www.baidu.com/ HTTP/1.1
2025/09/12 00:47:42 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:47:42 simple.go:117: address:www.baidu.com:80
2025/09/12 00:47:42 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.180.192.61:54324
2025/09/12 00:47:42 main.go:791: ProxySelector 127.180.192.61:54324
2025/09/12 00:47:42 main.go:797: 选择的代理 URL: ws://localhost:8080
2025/09/12 00:47:42 simple.go:179: WebSocket Config Details:
2025/09/12 00:47:42 simple.go:180: host, portNum 127.180.192.61 54324
2025/09/12 00:47:42 simple.go:181:   Username: 
2025/09/12 00:47:42 simple.go:182:   Password: 
2025/09/12 00:47:42 simple.go:183:   ServerAddr: ws://localhost:8080
2025/09/12 00:47:42 simple.go:184:   Protocol: websocket
2025/09/12 00:47:42 simple.go:185:   Timeout: 30s
2025/09/12 00:47:42 client.go:98: url: ws://localhost:8080
2025/09/12 00:47:42 client.go:99: headers: map[X-Proxy-Target-Host:[127.180.192.61] X-Proxy-Target-Port:[54324]]
2025/09/12 00:47:42 client.go:110: url: http://localhost:8080
2025/09/12 00:47:42 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[mkZ0e8z7BiXK0U2Qw1p5yTKjAdM=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:47:42 simple.go:209: WebSocket代理连接成功：127.180.192.61:54324
2025/09/12 00:47:42 simple.go:248: clienthost: ::1
2025/09/12 00:47:42 simple.go:249: clientport: 39113
2025/09/12 00:47:42 simple.go:278: simple Handle header:
2025/09/12 00:47:42 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:47:42 simple.go:283: HEAD / HTTP/1.1
2025/09/12 00:47:42 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:47:42 simple.go:294: Host: www.baidu.com
User-Agent: curl/8.12.1
Accept: */*
Proxy-Connection: Keep-Alive
2025/09/12 00:47:42 http.go:90: method: HEAD
2025/09/12 00:47:42 http.go:91: url: /
2025/09/12 00:47:42 http.go:92: host: www.baidu.com
2025/09/12 00:47:42 http.go:93: proxyHandler header:
2025/09/12 00:47:42 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:47:42 http.go:121: clientport: 39115
2025/09/12 00:47:42 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.180.192.61:54324;host=www.baidu.com;proto=http
2025/09/12 00:47:42 http.go:132: proxyHandler User-Agent : curl/8.12.1
2025/09/12 00:47:42 http.go:132: proxyHandler Accept : */*
2025/09/12 00:47:42 http.go:132: proxyHandler Proxy-Connection : Keep-Alive
2025/09/12 00:47:42 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.180.192.61:54324;host=www.baidu.com;proto=http
2025/09/12 00:47:42 http.go:137: forwardedByList: [{[::1]:10810} {127.180.192.61:54324}]
2025/09/12 00:47:42 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:47:42 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:47:42 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:47:42 main.go:797: 选择的代理 URL: ws://localhost:8080
2025/09/12 00:47:42 http.go:250: 使用代理：ws://localhost:8080
2025/09/12 00:47:42 http.go:259: 已经修改了代理为websocket ws://localhost:8080
[GIN] 2025/09/12 - 00:47:42 | 200 |     12.9499ms |       127.0.0.1 | HEAD     "/"
2025/09/12 00:47:42 simple.go:57: remote addr: [::1]:39116
2025/09/12 00:47:42 simple.go:79: CONNECT www.baidu.com:443 HTTP/1.1
2025/09/12 00:47:42 simple.go:117: address:www.baidu.com:443
2025/09/12 00:47:42 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com:443
2025/09/12 00:47:42 main.go:791: ProxySelector www.baidu.com:443
2025/09/12 00:47:42 main.go:797: 选择的代理 URL: ws://localhost:8080
2025/09/12 00:47:42 simple.go:179: WebSocket Config Details:
2025/09/12 00:47:42 simple.go:180: host, portNum www.baidu.com 443
2025/09/12 00:47:42 simple.go:181:   Username: 
2025/09/12 00:47:42 simple.go:182:   Password: 
2025/09/12 00:47:42 simple.go:183:   ServerAddr: ws://localhost:8080
2025/09/12 00:47:42 simple.go:184:   Protocol: websocket
2025/09/12 00:47:42 simple.go:185:   Timeout: 30s
2025/09/12 00:47:42 client.go:98: url: ws://localhost:8080
2025/09/12 00:47:42 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:47:42 client.go:110: url: http://localhost:8080
2025/09/12 00:47:42 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[DWcmplcG44MNvgqX92Ca/6skiF0=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:47:42 simple.go:209: WebSocket代理连接成功：www.baidu.com:443
2025/09/12 00:47:42 http.go:493: WebSocket ForwardData error: read tcp [::1]:39109->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:47:42 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39107->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:47:42 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39114->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:47:42 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39117->[::1]:8080: wsarecv: An existing connection was forcibly closed by the remote host.
```

✅ 端口8080已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:48:20

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 29992

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 42368

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 32548, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:48:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11107466872888363946
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:48:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11107466872888363946


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 49220, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:48:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11809174807241000235
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:48:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11809174807241000235


```

### 📋 所有进程PID记录

所有进程PID: 29992, 42368, 32548, 49220



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:48:20

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 29992

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 42368

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 32548, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:48:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11107466872888363946
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:48:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11107466872888363946


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 49220, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:48:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11809174807241000235
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:48:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11809174807241000235


```

### 📋 所有进程PID记录

所有进程PID: 29992, 42368, 32548, 49220

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:48:22 main.go:71: 启动websocket服务端，监听地址: :38800
2025/09/12 00:48:22 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :38800
2025/09/12 00:48:22 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:48:22 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:48:22 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:48:22 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:48:24 server.go:90: url /
2025/09/12 00:48:24 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[WutrLnANGPN5lP18sBGNUA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.51.65.105] X-Proxy-Target-Port:[17811]]
2025/09/12 00:48:24 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39364 at 2025-09-12 00:48:24
2025/09/12 00:48:24 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39364
2025/09/12 00:48:24 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.51.65.105', targetPort: 17811
2025/09/12 00:48:24 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:48:24 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39364
2025/09/12 00:48:24 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.51.65.105:17811 from [::1]:39364
2025/09/12 00:48:24 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.51.65.105:17811 (timeout: 30s)
2025/09/12 00:48:24 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.51.65.105:17811
2025/09/12 00:48:24 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:48:24 server.go:90: url /
2025/09/12 00:48:24 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[Ntyhr9jc6xhdRZE6EXpByg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:48:24 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39366 at 2025-09-12 00:48:24
2025/09/12 00:48:24 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39366
2025/09/12 00:48:24 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 80
2025/09/12 00:48:24 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:48:24 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39366
2025/09/12 00:48:24 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:80 from [::1]:39366
2025/09/12 00:48:24 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:80 (timeout: 30s)
2025/09/12 00:48:24 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:80
2025/09/12 00:48:24 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:48:26 server.go:90: url /
2025/09/12 00:48:26 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[9RDRWJVeGIqjoOg64tIf/g==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.51.65.105] X-Proxy-Target-Port:[17811]]
2025/09/12 00:48:26 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39378 at 2025-09-12 00:48:26
2025/09/12 00:48:26 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39378
2025/09/12 00:48:26 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.51.65.105', targetPort: 17811
2025/09/12 00:48:26 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:48:26 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39378
2025/09/12 00:48:26 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.51.65.105:17811 from [::1]:39378
2025/09/12 00:48:26 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.51.65.105:17811 (timeout: 30s)
2025/09/12 00:48:26 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.51.65.105:17811
2025/09/12 00:48:26 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:48:26 server.go:90: url /
2025/09/12 00:48:26 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[fclFny8kt3sOvd4wAsVgYA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:48:26 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:39381 at 2025-09-12 00:48:26
2025/09/12 00:48:26 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:39381
2025/09/12 00:48:26 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 443
2025/09/12 00:48:26 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:48:26 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:39381
2025/09/12 00:48:26 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:443 from [::1]:39381
2025/09/12 00:48:26 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:443 (timeout: 30s)
2025/09/12 00:48:26 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:443
2025/09/12 00:48:26 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:48:23 main.go:518: 代理服务器启动中...
2025/09/12 00:48:23 main.go:562: dohalpn: []
2025/09/12 00:48:23 main.go:564: hostname: 0.0.0.0
2025/09/12 00:48:23 main.go:566: port: 10810
2025/09/12 00:48:23 main.go:568: server_cert: 
2025/09/12 00:48:23 main.go:570: server_key: 
2025/09/12 00:48:23 main.go:572: username: 
2025/09/12 00:48:23 main.go:574: password: 
2025/09/12 00:48:23 main.go:576: dohurl: []
2025/09/12 00:48:23 main.go:578: dohip: []
2025/09/12 00:48:23 main.go:579: upstream-type: websocket
2025/09/12 00:48:23 main.go:580: upstream-address: ws://localhost:38800
2025/09/12 00:48:23 main.go:581: upstream-username: 
2025/09/12 00:48:23 main.go:582: upstream-password: 
2025/09/12 00:48:23 main.go:639: WebSocket代理配置已添加
2025/09/12 00:48:23 main.go:854: {
  "hostname": "",
  "port": 0,
  "server_cert": "",
  "server_key": "",
  "username": "",
  "password": "",
  "doh": null,
  "upstreams": {
    "websocket_upstream": {
      "type": "websocket",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "ws_proxy": "ws://localhost:38800",
      "ws_username": "",
      "ws_password": "",
      "http_username": "",
      "http_password": "",
      "socks5_proxy": "",
      "socks5_username": "",
      "socks5_password": ""
    }
  },
  "rules": [
    {
      "filter": "websocket_filter",
      "upstream": "websocket_upstream"
    }
  ],
  "filters": {
    "websocket_filter": {
      "patterns": [
        "*"
      ]
    }
  }
}
2025/09/12 00:48:23 simple.go:31: Proxy server started on port [::]:10810
2025/09/12 00:48:23 http.go:372: Random IP: 127.51.65.105
2025/09/12 00:48:23 http.go:390: Random integer: 17811
2025/09/12 00:48:23 http.go:342: Proxy server started on port 127.51.65.105:17811
2025/09/12 00:48:24 simple.go:57: remote addr: [::1]:39363
2025/09/12 00:48:24 simple.go:79: GET http://www.baidu.com/ HTTP/1.1
2025/09/12 00:48:24 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:48:24 simple.go:117: address:www.baidu.com:80
2025/09/12 00:48:24 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.51.65.105:17811
2025/09/12 00:48:24 main.go:791: ProxySelector 127.51.65.105:17811
2025/09/12 00:48:24 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:48:24 simple.go:179: WebSocket Config Details:
2025/09/12 00:48:24 simple.go:180: host, portNum 127.51.65.105 17811
2025/09/12 00:48:24 simple.go:181:   Username: 
2025/09/12 00:48:24 simple.go:182:   Password: 
2025/09/12 00:48:24 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:48:24 simple.go:184:   Protocol: websocket
2025/09/12 00:48:24 simple.go:185:   Timeout: 30s
2025/09/12 00:48:24 client.go:98: url: ws://localhost:38800
2025/09/12 00:48:24 client.go:99: headers: map[X-Proxy-Target-Host:[127.51.65.105] X-Proxy-Target-Port:[17811]]
2025/09/12 00:48:24 client.go:110: url: http://localhost:38800
2025/09/12 00:48:24 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[ZaN9mQaWvsoOQq7FTPOZv82uhjI=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:48:24 simple.go:209: WebSocket代理连接成功：127.51.65.105:17811
2025/09/12 00:48:24 simple.go:248: clienthost: ::1
2025/09/12 00:48:24 simple.go:249: clientport: 39363
2025/09/12 00:48:24 simple.go:278: simple Handle header:
2025/09/12 00:48:24 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:48:24 simple.go:283: GET / HTTP/1.1
2025/09/12 00:48:24 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:48:24 simple.go:294: Host: www.baidu.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
2025/09/12 00:48:24 http.go:90: method: GET
2025/09/12 00:48:24 http.go:91: url: /
2025/09/12 00:48:24 http.go:92: host: www.baidu.com
2025/09/12 00:48:24 http.go:93: proxyHandler header:
2025/09/12 00:48:24 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:48:24 http.go:121: clientport: 39365
2025/09/12 00:48:24 http.go:132: proxyHandler Accept-Encoding : gzip
2025/09/12 00:48:24 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.51.65.105:17811;host=www.baidu.com;proto=http
2025/09/12 00:48:24 http.go:132: proxyHandler User-Agent : Go-http-client/1.1
2025/09/12 00:48:24 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.51.65.105:17811;host=www.baidu.com;proto=http
2025/09/12 00:48:24 http.go:137: forwardedByList: [{[::1]:10810} {127.51.65.105:17811}]
2025/09/12 00:48:24 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:48:24 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:48:24 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:48:24 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:48:24 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:48:24 http.go:259: 已经修改了代理为websocket ws://localhost:38800
2025/09/12 00:48:24 http.go:262: 使用代理：ws://localhost:38800
2025/09/12 00:48:24 http.go:264: network,addr tcp www.baidu.com:80
2025/09/12 00:48:24 http.go:466: WebSocket Config Details:
2025/09/12 00:48:24 http.go:467: host, portNum www.baidu.com 80
2025/09/12 00:48:24 http.go:468:   Username: 
2025/09/12 00:48:24 http.go:469:   Password: 
2025/09/12 00:48:24 http.go:470:   ServerAddr: ws://localhost:38800
2025/09/12 00:48:24 http.go:471:   Protocol: websocket
2025/09/12 00:48:24 http.go:472:   Timeout: 30s
2025/09/12 00:48:24 http.go:476: host, portNum www.baidu.com 80
2025/09/12 00:48:24 client.go:98: url: ws://localhost:38800
2025/09/12 00:48:24 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:48:24 client.go:110: url: http://localhost:38800
2025/09/12 00:48:24 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[zcYnXdzciOZdCoco7JQ1dioK4vw=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
[GIN] 2025/09/12 - 00:48:24 | 200 |     56.8347ms |       127.0.0.1 | GET      "/"
2025/09/12 00:48:26 simple.go:57: remote addr: [::1]:39377
2025/09/12 00:48:26 simple.go:79: HEAD http://www.baidu.com/ HTTP/1.1
2025/09/12 00:48:26 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:48:26 simple.go:117: address:www.baidu.com:80
2025/09/12 00:48:26 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.51.65.105:17811
2025/09/12 00:48:26 main.go:791: ProxySelector 127.51.65.105:17811
2025/09/12 00:48:26 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:48:26 simple.go:179: WebSocket Config Details:
2025/09/12 00:48:26 simple.go:180: host, portNum 127.51.65.105 17811
2025/09/12 00:48:26 simple.go:181:   Username: 
2025/09/12 00:48:26 simple.go:182:   Password: 
2025/09/12 00:48:26 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:48:26 simple.go:184:   Protocol: websocket
2025/09/12 00:48:26 simple.go:185:   Timeout: 30s
2025/09/12 00:48:26 client.go:98: url: ws://localhost:38800
2025/09/12 00:48:26 client.go:99: headers: map[X-Proxy-Target-Host:[127.51.65.105] X-Proxy-Target-Port:[17811]]
2025/09/12 00:48:26 client.go:110: url: http://localhost:38800
2025/09/12 00:48:26 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[otrxlSU4V+eTGBHP28FWjLLAEWc=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:48:26 simple.go:209: WebSocket代理连接成功：127.51.65.105:17811
2025/09/12 00:48:26 simple.go:248: clienthost: ::1
2025/09/12 00:48:26 simple.go:249: clientport: 39377
2025/09/12 00:48:26 simple.go:278: simple Handle header:
2025/09/12 00:48:26 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:48:26 simple.go:283: HEAD / HTTP/1.1
2025/09/12 00:48:26 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:48:26 simple.go:294: Host: www.baidu.com
User-Agent: curl/8.12.1
Accept: */*
Proxy-Connection: Keep-Alive
2025/09/12 00:48:26 http.go:90: method: HEAD
2025/09/12 00:48:26 http.go:91: url: /
2025/09/12 00:48:26 http.go:92: host: www.baidu.com
2025/09/12 00:48:26 http.go:93: proxyHandler header:
2025/09/12 00:48:26 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:48:26 http.go:121: clientport: 39379
2025/09/12 00:48:26 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.51.65.105:17811;host=www.baidu.com;proto=http
2025/09/12 00:48:26 http.go:132: proxyHandler User-Agent : curl/8.12.1
2025/09/12 00:48:26 http.go:132: proxyHandler Accept : */*
2025/09/12 00:48:26 http.go:132: proxyHandler Proxy-Connection : Keep-Alive
2025/09/12 00:48:26 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.51.65.105:17811;host=www.baidu.com;proto=http
2025/09/12 00:48:26 http.go:137: forwardedByList: [{[::1]:10810} {127.51.65.105:17811}]
2025/09/12 00:48:26 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:48:26 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:48:26 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:48:26 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:48:26 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:48:26 http.go:259: 已经修改了代理为websocket ws://localhost:38800
[GIN] 2025/09/12 - 00:48:26 | 200 |     14.8668ms |       127.0.0.1 | HEAD     "/"
2025/09/12 00:48:26 simple.go:57: remote addr: [::1]:39380
2025/09/12 00:48:26 simple.go:79: CONNECT www.baidu.com:443 HTTP/1.1
2025/09/12 00:48:26 simple.go:117: address:www.baidu.com:443
2025/09/12 00:48:26 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com:443
2025/09/12 00:48:26 main.go:791: ProxySelector www.baidu.com:443
2025/09/12 00:48:26 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:48:26 simple.go:179: WebSocket Config Details:
2025/09/12 00:48:26 simple.go:180: host, portNum www.baidu.com 443
2025/09/12 00:48:26 simple.go:181:   Username: 
2025/09/12 00:48:26 simple.go:182:   Password: 
2025/09/12 00:48:26 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:48:26 simple.go:184:   Protocol: websocket
2025/09/12 00:48:26 simple.go:185:   Timeout: 30s
2025/09/12 00:48:26 client.go:98: url: ws://localhost:38800
2025/09/12 00:48:26 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:48:26 client.go:110: url: http://localhost:38800
2025/09/12 00:48:26 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[jRFd21piuiT7frWI1MvJcL/qJMo=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:48:26 simple.go:209: WebSocket代理连接成功：www.baidu.com:443
2025/09/12 00:48:26 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39378->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:48:26 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39381->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:48:26 simple.go:204: WebSocket ForwardData error: read tcp [::1]:39364->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:48:26 http.go:493: WebSocket ForwardData error: read tcp [::1]:39366->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:49:08

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 53096

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 29328

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 53840, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:49:02 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11259716535820513827
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:49:02 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11259716535820513827


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 44208, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:49:02 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11168092801432514417
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:49:02 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11168092801432514417


```

### 📋 所有进程PID记录

所有进程PID: 53096, 29328, 53840, 44208



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:49:08

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 53096

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 29328

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 53840, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:49:02 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11259716535820513827
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:49:02 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11259716535820513827


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 44208, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:49:02 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11168092801432514417
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:49:02 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11168092801432514417


```

### 📋 所有进程PID记录

所有进程PID: 53096, 29328, 53840, 44208

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
```

### SOCKS5服务器日志输出

```
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:50:55

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 42920

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 28300

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 50416, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:50:49 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11103280684161771938
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:50:49 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11103280684161771938


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 18956, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:50:49 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11003105441950832344
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:50:49 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11003105441950832344


```

### 📋 所有进程PID记录

所有进程PID: 42920, 28300, 50416, 18956



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:50:55

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 42920

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 28300

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 50416, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:50:49 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11103280684161771938
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:50:49 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11103280684161771938


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 18956, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:50:49 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11003105441950832344
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:50:49 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11003105441950832344


```

### 📋 所有进程PID记录

所有进程PID: 42920, 28300, 50416, 18956

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
```

### SOCKS5服务器日志输出

```
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:51:12

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 18520

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 11044

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 47748, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:51:05 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11515057590084730121
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:51:05 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11515057590084730121


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 30992, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:51:05 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11201437050654097518
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:51:05 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11201437050654097518


```

### 📋 所有进程PID记录

所有进程PID: 18520, 11044, 47748, 30992



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:51:12

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 18520

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 11044

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 47748, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:51:05 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11515057590084730121
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:51:05 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11515057590084730121


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 30992, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:51:05 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11201437050654097518
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:51:05 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11201437050654097518


```

### 📋 所有进程PID记录

所有进程PID: 18520, 11044, 47748, 30992

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:51:13 main.go:71: 启动websocket服务端，监听地址: :38800
2025/09/12 00:51:13 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :38800
2025/09/12 00:51:13 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:51:13 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:51:13 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:51:13 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:51:15 server.go:90: url /
2025/09/12 00:51:15 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[MFVH1w8d9tKAIVjy1YBsKQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.29.80.86] X-Proxy-Target-Port:[26801]]
2025/09/12 00:51:15 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:40573 at 2025-09-12 00:51:15
2025/09/12 00:51:15 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:40573
2025/09/12 00:51:15 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.29.80.86', targetPort: 26801
2025/09/12 00:51:15 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:51:15 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:40573
2025/09/12 00:51:15 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.29.80.86:26801 from [::1]:40573
2025/09/12 00:51:15 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.29.80.86:26801 (timeout: 30s)
2025/09/12 00:51:15 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.29.80.86:26801
2025/09/12 00:51:15 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:51:15 server.go:90: url /
2025/09/12 00:51:15 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[FzwIR0iyHnHha8gTYPdMxQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:51:15 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:40575 at 2025-09-12 00:51:15
2025/09/12 00:51:15 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:40575
2025/09/12 00:51:15 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 80
2025/09/12 00:51:15 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:51:15 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:40575
2025/09/12 00:51:15 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:80 from [::1]:40575
2025/09/12 00:51:15 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:80 (timeout: 30s)
2025/09/12 00:51:15 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:80
2025/09/12 00:51:15 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:51:17 server.go:90: url /
2025/09/12 00:51:17 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[qJxlvL49JxS7dQZJjGkisQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.29.80.86] X-Proxy-Target-Port:[26801]]
2025/09/12 00:51:17 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:40625 at 2025-09-12 00:51:17
2025/09/12 00:51:17 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:40625
2025/09/12 00:51:17 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.29.80.86', targetPort: 26801
2025/09/12 00:51:17 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:51:17 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:40625
2025/09/12 00:51:17 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.29.80.86:26801 from [::1]:40625
2025/09/12 00:51:17 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.29.80.86:26801 (timeout: 30s)
2025/09/12 00:51:17 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.29.80.86:26801
2025/09/12 00:51:17 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:51:17 server.go:90: url /
2025/09/12 00:51:17 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[qnHGVtqAL/y86BbUKJI23g==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:51:17 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:40628 at 2025-09-12 00:51:17
2025/09/12 00:51:17 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:40628
2025/09/12 00:51:17 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 443
2025/09/12 00:51:17 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:51:17 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:40628
2025/09/12 00:51:17 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:443 from [::1]:40628
2025/09/12 00:51:17 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:443 (timeout: 30s)
2025/09/12 00:51:17 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:443
2025/09/12 00:51:17 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:51:14 main.go:518: 代理服务器启动中...
2025/09/12 00:51:14 main.go:562: dohalpn: []
2025/09/12 00:51:14 main.go:564: hostname: 0.0.0.0
2025/09/12 00:51:14 main.go:566: port: 10810
2025/09/12 00:51:14 main.go:568: server_cert: 
2025/09/12 00:51:14 main.go:570: server_key: 
2025/09/12 00:51:14 main.go:572: username: 
2025/09/12 00:51:14 main.go:574: password: 
2025/09/12 00:51:14 main.go:576: dohurl: []
2025/09/12 00:51:14 main.go:578: dohip: []
2025/09/12 00:51:14 main.go:579: upstream-type: websocket
2025/09/12 00:51:14 main.go:580: upstream-address: ws://localhost:38800
2025/09/12 00:51:14 main.go:581: upstream-username: 
2025/09/12 00:51:14 main.go:582: upstream-password: 
2025/09/12 00:51:14 main.go:639: WebSocket代理配置已添加
2025/09/12 00:51:14 main.go:854: {
  "hostname": "",
  "port": 0,
  "server_cert": "",
  "server_key": "",
  "username": "",
  "password": "",
  "doh": null,
  "upstreams": {
    "websocket_upstream": {
      "type": "websocket",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "ws_proxy": "ws://localhost:38800",
      "ws_username": "",
      "ws_password": "",
      "http_username": "",
      "http_password": "",
      "socks5_proxy": "",
      "socks5_username": "",
      "socks5_password": ""
    }
  },
  "rules": [
    {
      "filter": "websocket_filter",
      "upstream": "websocket_upstream"
    }
  ],
  "filters": {
    "websocket_filter": {
      "patterns": [
        "*"
      ]
    }
  }
}
2025/09/12 00:51:14 simple.go:31: Proxy server started on port [::]:10810
2025/09/12 00:51:14 http.go:372: Random IP: 127.29.80.86
2025/09/12 00:51:14 http.go:390: Random integer: 26801
2025/09/12 00:51:14 http.go:342: Proxy server started on port 127.29.80.86:26801
2025/09/12 00:51:15 simple.go:57: remote addr: [::1]:40572
2025/09/12 00:51:15 simple.go:79: GET http://www.baidu.com/ HTTP/1.1
2025/09/12 00:51:15 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:51:15 simple.go:117: address:www.baidu.com:80
2025/09/12 00:51:15 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.29.80.86:26801
2025/09/12 00:51:15 main.go:791: ProxySelector 127.29.80.86:26801
2025/09/12 00:51:15 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:51:15 simple.go:179: WebSocket Config Details:
2025/09/12 00:51:15 simple.go:180: host, portNum 127.29.80.86 26801
2025/09/12 00:51:15 simple.go:181:   Username: 
2025/09/12 00:51:15 simple.go:182:   Password: 
2025/09/12 00:51:15 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:51:15 simple.go:184:   Protocol: websocket
2025/09/12 00:51:15 simple.go:185:   Timeout: 30s
2025/09/12 00:51:15 client.go:98: url: ws://localhost:38800
2025/09/12 00:51:15 client.go:99: headers: map[X-Proxy-Target-Host:[127.29.80.86] X-Proxy-Target-Port:[26801]]
2025/09/12 00:51:15 client.go:110: url: http://localhost:38800
2025/09/12 00:51:15 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[qYvnYuEw1GqdXmcHSLfL7RzyHJ4=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:51:15 simple.go:209: WebSocket代理连接成功：127.29.80.86:26801
2025/09/12 00:51:15 simple.go:248: clienthost: ::1
2025/09/12 00:51:15 simple.go:249: clientport: 40572
2025/09/12 00:51:15 simple.go:278: simple Handle header:
2025/09/12 00:51:15 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:51:15 simple.go:283: GET / HTTP/1.1
2025/09/12 00:51:15 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:51:15 simple.go:294: Host: www.baidu.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
2025/09/12 00:51:15 http.go:90: method: GET
2025/09/12 00:51:15 http.go:91: url: /
2025/09/12 00:51:15 http.go:92: host: www.baidu.com
2025/09/12 00:51:15 http.go:93: proxyHandler header:
2025/09/12 00:51:15 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:51:15 http.go:121: clientport: 40574
2025/09/12 00:51:15 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.29.80.86:26801;host=www.baidu.com;proto=http
2025/09/12 00:51:15 http.go:132: proxyHandler User-Agent : Go-http-client/1.1
2025/09/12 00:51:15 http.go:132: proxyHandler Accept-Encoding : gzip
2025/09/12 00:51:15 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.29.80.86:26801;host=www.baidu.com;proto=http
2025/09/12 00:51:15 http.go:137: forwardedByList: [{[::1]:10810} {127.29.80.86:26801}]
2025/09/12 00:51:15 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:51:15 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:51:15 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:51:15 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:51:15 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:51:15 http.go:259: 已经修改了代理为websocket ws://localhost:38800
2025/09/12 00:51:15 http.go:262: 使用代理：ws://localhost:38800
2025/09/12 00:51:15 http.go:264: network,addr tcp www.baidu.com:80
2025/09/12 00:51:15 http.go:466: WebSocket Config Details:
2025/09/12 00:51:15 http.go:467: host, portNum www.baidu.com 80
2025/09/12 00:51:15 http.go:468:   Username: 
2025/09/12 00:51:15 http.go:469:   Password: 
2025/09/12 00:51:15 http.go:470:   ServerAddr: ws://localhost:38800
2025/09/12 00:51:15 http.go:471:   Protocol: websocket
2025/09/12 00:51:15 http.go:472:   Timeout: 30s
2025/09/12 00:51:15 http.go:476: host, portNum www.baidu.com 80
2025/09/12 00:51:15 client.go:98: url: ws://localhost:38800
2025/09/12 00:51:15 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:51:15 client.go:110: url: http://localhost:38800
2025/09/12 00:51:15 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[mrYBj6dQzDNV+0cJkogelOTX16o=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
[GIN] 2025/09/12 - 00:51:15 | 200 |     49.5107ms |       127.0.0.1 | GET      "/"
2025/09/12 00:51:17 simple.go:57: remote addr: [::1]:40624
2025/09/12 00:51:17 simple.go:79: HEAD http://www.baidu.com/ HTTP/1.1
2025/09/12 00:51:17 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:51:17 simple.go:117: address:www.baidu.com:80
2025/09/12 00:51:17 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.29.80.86:26801
2025/09/12 00:51:17 main.go:791: ProxySelector 127.29.80.86:26801
2025/09/12 00:51:17 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:51:17 simple.go:179: WebSocket Config Details:
2025/09/12 00:51:17 simple.go:180: host, portNum 127.29.80.86 26801
2025/09/12 00:51:17 simple.go:181:   Username: 
2025/09/12 00:51:17 simple.go:182:   Password: 
2025/09/12 00:51:17 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:51:17 simple.go:184:   Protocol: websocket
2025/09/12 00:51:17 simple.go:185:   Timeout: 30s
2025/09/12 00:51:17 client.go:98: url: ws://localhost:38800
2025/09/12 00:51:17 client.go:99: headers: map[X-Proxy-Target-Host:[127.29.80.86] X-Proxy-Target-Port:[26801]]
2025/09/12 00:51:17 client.go:110: url: http://localhost:38800
2025/09/12 00:51:17 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[YNFwCMPe7La79XhjXZXhBk6ekKI=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:51:17 simple.go:209: WebSocket代理连接成功：127.29.80.86:26801
2025/09/12 00:51:17 simple.go:248: clienthost: ::1
2025/09/12 00:51:17 simple.go:249: clientport: 40624
2025/09/12 00:51:17 simple.go:278: simple Handle header:
2025/09/12 00:51:17 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:51:17 simple.go:283: HEAD / HTTP/1.1
2025/09/12 00:51:17 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:51:17 simple.go:294: Host: www.baidu.com
User-Agent: curl/8.12.1
Accept: */*
Proxy-Connection: Keep-Alive
2025/09/12 00:51:17 http.go:90: method: HEAD
2025/09/12 00:51:17 http.go:91: url: /
2025/09/12 00:51:17 http.go:92: host: www.baidu.com
2025/09/12 00:51:17 http.go:93: proxyHandler header:
2025/09/12 00:51:17 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:51:17 http.go:121: clientport: 40626
2025/09/12 00:51:17 http.go:132: proxyHandler User-Agent : curl/8.12.1
2025/09/12 00:51:17 http.go:132: proxyHandler Accept : */*
2025/09/12 00:51:17 http.go:132: proxyHandler Proxy-Connection : Keep-Alive
2025/09/12 00:51:17 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.29.80.86:26801;host=www.baidu.com;proto=http
2025/09/12 00:51:17 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.29.80.86:26801;host=www.baidu.com;proto=http
2025/09/12 00:51:17 http.go:137: forwardedByList: [{[::1]:10810} {127.29.80.86:26801}]
2025/09/12 00:51:17 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:51:17 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:51:17 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:51:17 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:51:17 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:51:17 http.go:259: 已经修改了代理为websocket ws://localhost:38800
[GIN] 2025/09/12 - 00:51:17 | 200 |      17.265ms |       127.0.0.1 | HEAD     "/"
2025/09/12 00:51:17 simple.go:57: remote addr: [::1]:40627
2025/09/12 00:51:17 simple.go:79: CONNECT www.baidu.com:443 HTTP/1.1
2025/09/12 00:51:17 simple.go:117: address:www.baidu.com:443
2025/09/12 00:51:17 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com:443
2025/09/12 00:51:17 main.go:791: ProxySelector www.baidu.com:443
2025/09/12 00:51:17 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:51:17 simple.go:179: WebSocket Config Details:
2025/09/12 00:51:17 simple.go:180: host, portNum www.baidu.com 443
2025/09/12 00:51:17 simple.go:181:   Username: 
2025/09/12 00:51:17 simple.go:182:   Password: 
2025/09/12 00:51:17 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:51:17 simple.go:184:   Protocol: websocket
2025/09/12 00:51:17 simple.go:185:   Timeout: 30s
2025/09/12 00:51:17 client.go:98: url: ws://localhost:38800
2025/09/12 00:51:17 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:51:17 client.go:110: url: http://localhost:38800
2025/09/12 00:51:17 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[Dnn/xHmYZ+hLQcjtFnnBym1c7+U=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:51:17 simple.go:209: WebSocket代理连接成功：www.baidu.com:443
2025/09/12 00:51:17 simple.go:204: WebSocket ForwardData error: read tcp [::1]:40573->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:51:17 http.go:493: WebSocket ForwardData error: read tcp [::1]:40575->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:51:17 simple.go:204: WebSocket ForwardData error: read tcp [::1]:40625->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:51:17 simple.go:204: WebSocket ForwardData error: read tcp [::1]:40628->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放


###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:52:53

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 54848

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 21080

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 48600, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:52:47 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11175283375068861257
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:52:47 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11175283375068861257


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 45504, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:52:47 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11537178552325439668
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:52:47 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11537178552325439668


```

### 📋 所有进程PID记录

所有进程PID: 54848, 21080, 48600, 45504



###

# WebSocket和SOCKS5级联代理测试记录

## 测试时间
2025-09-12 00:52:53

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 54848

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动SOCKS5服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 SOCKS5服务器进程PID: 21080

等待SOCKS5服务器启动...
✅ SOCKS5服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 48600, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:52:47 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11175283375068861257
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:52:47 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11175283375068861257


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 45504, 退出状态码: 0

✅ 测试成功

输出结果:
```
Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [308 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [102 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [4771 bytes data]
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
{ [333 bytes data]
* TLSv1.2 (IN), TLS handshake, Server finished (14):
{ [4 bytes data]
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
} [70 bytes data]
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
} [1 bytes data]
* TLSv1.2 (OUT), TLS handshake, Finished (20):
} [16 bytes data]
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
{ [1 bytes data]
* TLSv1.2 (IN), TLS handshake, Finished (20):
{ [16 bytes data]
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
* ALPN: server accepted http/1.1
* Server certificate:
*  subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science Technology Co., Ltd; CN=baidu.com
*  start date: Jul  9 07:01:02 2025 GMT
*  expire date: Aug 10 07:01:01 2026 GMT
*  subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
*  issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* Connected to localhost (::1) port 10810
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Thu, 11 Sep 2025 16:52:47 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11537178552325439668
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 16:52:47 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11537178552325439668


```

### 📋 所有进程PID记录

所有进程PID: 54848, 21080, 48600, 45504

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止SOCKS5服务器进程...
✅ SOCKS5服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 00:52:54 main.go:71: 启动websocket服务端，监听地址: :38800
2025/09/12 00:52:54 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :38800
2025/09/12 00:52:54 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 00:52:54 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 00:52:54 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 00:52:54 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 00:52:56 server.go:90: url /
2025/09/12 00:52:56 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[abLBxYaii5xSb/GElC5hBg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.214.169.149] X-Proxy-Target-Port:[46470]]
2025/09/12 00:52:56 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:41057 at 2025-09-12 00:52:56
2025/09/12 00:52:56 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:41057
2025/09/12 00:52:56 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.214.169.149', targetPort: 46470
2025/09/12 00:52:56 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:52:56 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:41057
2025/09/12 00:52:56 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.214.169.149:46470 from [::1]:41057
2025/09/12 00:52:56 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.214.169.149:46470 (timeout: 30s)
2025/09/12 00:52:56 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.214.169.149:46470
2025/09/12 00:52:56 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:52:56 server.go:90: url /
2025/09/12 00:52:56 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[KQBkvqmhosafI+xH0bAWEQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:52:56 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:41059 at 2025-09-12 00:52:56
2025/09/12 00:52:56 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:41059
2025/09/12 00:52:56 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 80
2025/09/12 00:52:56 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:52:56 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:41059
2025/09/12 00:52:56 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:80 from [::1]:41059
2025/09/12 00:52:56 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:80 (timeout: 30s)
2025/09/12 00:52:56 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:80
2025/09/12 00:52:56 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:52:58 server.go:90: url /
2025/09/12 00:52:58 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[BIM5HdAE89XbokdVaOcmRQ==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.214.169.149] X-Proxy-Target-Port:[46470]]
2025/09/12 00:52:58 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:41075 at 2025-09-12 00:52:58
2025/09/12 00:52:58 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:41075
2025/09/12 00:52:58 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.214.169.149', targetPort: 46470
2025/09/12 00:52:58 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:52:58 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:41075
2025/09/12 00:52:58 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.214.169.149:46470 from [::1]:41075
2025/09/12 00:52:58 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.214.169.149:46470 (timeout: 30s)
2025/09/12 00:52:58 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.214.169.149:46470
2025/09/12 00:52:58 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 00:52:59 server.go:90: url /
2025/09/12 00:52:59 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[UOmjObnlBdrshf1X/hOrLw==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:52:59 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:41078 at 2025-09-12 00:52:59
2025/09/12 00:52:59 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:41078
2025/09/12 00:52:59 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 443
2025/09/12 00:52:59 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 00:52:59 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:41078
2025/09/12 00:52:59 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:443 from [::1]:41078
2025/09/12 00:52:59 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:443 (timeout: 30s)
2025/09/12 00:52:59 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:443
2025/09/12 00:52:59 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### SOCKS5服务器日志输出

```
2025/09/12 00:52:55 main.go:518: 代理服务器启动中...
2025/09/12 00:52:55 main.go:562: dohalpn: []
2025/09/12 00:52:55 main.go:564: hostname: 0.0.0.0
2025/09/12 00:52:55 main.go:566: port: 10810
2025/09/12 00:52:55 main.go:568: server_cert: 
2025/09/12 00:52:55 main.go:570: server_key: 
2025/09/12 00:52:55 main.go:572: username: 
2025/09/12 00:52:55 main.go:574: password: 
2025/09/12 00:52:55 main.go:576: dohurl: []
2025/09/12 00:52:55 main.go:578: dohip: []
2025/09/12 00:52:55 main.go:579: upstream-type: websocket
2025/09/12 00:52:55 main.go:580: upstream-address: ws://localhost:38800
2025/09/12 00:52:55 main.go:581: upstream-username: 
2025/09/12 00:52:55 main.go:582: upstream-password: 
2025/09/12 00:52:55 main.go:639: WebSocket代理配置已添加
2025/09/12 00:52:55 main.go:854: {
  "hostname": "",
  "port": 0,
  "server_cert": "",
  "server_key": "",
  "username": "",
  "password": "",
  "doh": null,
  "upstreams": {
    "websocket_upstream": {
      "type": "websocket",
      "http_proxy": "",
      "https_proxy": "",
      "bypass_list": [],
      "ws_proxy": "ws://localhost:38800",
      "ws_username": "",
      "ws_password": "",
      "http_username": "",
      "http_password": "",
      "socks5_proxy": "",
      "socks5_username": "",
      "socks5_password": ""
    }
  },
  "rules": [
    {
      "filter": "websocket_filter",
      "upstream": "websocket_upstream"
    }
  ],
  "filters": {
    "websocket_filter": {
      "patterns": [
        "*"
      ]
    }
  }
}
2025/09/12 00:52:55 simple.go:31: Proxy server started on port [::]:10810
2025/09/12 00:52:55 http.go:372: Random IP: 127.214.169.149
2025/09/12 00:52:55 http.go:390: Random integer: 46470
2025/09/12 00:52:55 http.go:342: Proxy server started on port 127.214.169.149:46470
2025/09/12 00:52:56 simple.go:57: remote addr: [::1]:41056
2025/09/12 00:52:56 simple.go:79: GET http://www.baidu.com/ HTTP/1.1
2025/09/12 00:52:56 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:52:56 simple.go:117: address:www.baidu.com:80
2025/09/12 00:52:56 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.214.169.149:46470
2025/09/12 00:52:56 main.go:791: ProxySelector 127.214.169.149:46470
2025/09/12 00:52:56 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:52:56 simple.go:179: WebSocket Config Details:
2025/09/12 00:52:56 simple.go:180: host, portNum 127.214.169.149 46470
2025/09/12 00:52:56 simple.go:181:   Username: 
2025/09/12 00:52:56 simple.go:182:   Password: 
2025/09/12 00:52:56 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:52:56 simple.go:184:   Protocol: websocket
2025/09/12 00:52:56 simple.go:185:   Timeout: 30s
2025/09/12 00:52:56 client.go:98: url: ws://localhost:38800
2025/09/12 00:52:56 client.go:99: headers: map[X-Proxy-Target-Host:[127.214.169.149] X-Proxy-Target-Port:[46470]]
2025/09/12 00:52:56 client.go:110: url: http://localhost:38800
2025/09/12 00:52:56 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[zrbg/dGrGtwaoq2UTUcYGec0obM=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:52:56 simple.go:209: WebSocket代理连接成功：127.214.169.149:46470
2025/09/12 00:52:56 simple.go:248: clienthost: ::1
2025/09/12 00:52:56 simple.go:249: clientport: 41056
2025/09/12 00:52:56 simple.go:278: simple Handle header:
2025/09/12 00:52:56 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:52:56 simple.go:283: GET / HTTP/1.1
2025/09/12 00:52:56 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:52:56 simple.go:294: Host: www.baidu.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
2025/09/12 00:52:56 http.go:90: method: GET
2025/09/12 00:52:56 http.go:91: url: /
2025/09/12 00:52:56 http.go:92: host: www.baidu.com
2025/09/12 00:52:56 http.go:93: proxyHandler header:
2025/09/12 00:52:56 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:52:56 http.go:121: clientport: 41058
2025/09/12 00:52:56 http.go:132: proxyHandler Accept-Encoding : gzip
2025/09/12 00:52:56 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.214.169.149:46470;host=www.baidu.com;proto=http
2025/09/12 00:52:56 http.go:132: proxyHandler User-Agent : Go-http-client/1.1
2025/09/12 00:52:56 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.214.169.149:46470;host=www.baidu.com;proto=http
2025/09/12 00:52:56 http.go:137: forwardedByList: [{[::1]:10810} {127.214.169.149:46470}]
2025/09/12 00:52:56 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:52:56 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:52:56 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:52:56 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:52:56 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:52:56 http.go:259: 已经修改了代理为websocket ws://localhost:38800
2025/09/12 00:52:56 http.go:262: 使用代理：ws://localhost:38800
2025/09/12 00:52:56 http.go:264: network,addr tcp www.baidu.com:80
2025/09/12 00:52:56 http.go:466: WebSocket Config Details:
2025/09/12 00:52:56 http.go:467: host, portNum www.baidu.com 80
2025/09/12 00:52:56 http.go:468:   Username: 
2025/09/12 00:52:56 http.go:469:   Password: 
2025/09/12 00:52:56 http.go:470:   ServerAddr: ws://localhost:38800
2025/09/12 00:52:56 http.go:471:   Protocol: websocket
2025/09/12 00:52:56 http.go:472:   Timeout: 30s
2025/09/12 00:52:56 http.go:476: host, portNum www.baidu.com 80
2025/09/12 00:52:56 client.go:98: url: ws://localhost:38800
2025/09/12 00:52:56 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 00:52:56 client.go:110: url: http://localhost:38800
2025/09/12 00:52:56 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[WdcqFPmlNE71Ui3OlbcrBehqqVc=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
[GIN] 2025/09/12 - 00:52:56 | 200 |     60.5102ms |       127.0.0.1 | GET      "/"
2025/09/12 00:52:58 simple.go:57: remote addr: [::1]:41074
2025/09/12 00:52:58 simple.go:79: HEAD http://www.baidu.com/ HTTP/1.1
2025/09/12 00:52:58 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 00:52:58 simple.go:117: address:www.baidu.com:80
2025/09/12 00:52:58 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.214.169.149:46470
2025/09/12 00:52:58 main.go:791: ProxySelector 127.214.169.149:46470
2025/09/12 00:52:58 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:52:58 simple.go:179: WebSocket Config Details:
2025/09/12 00:52:58 simple.go:180: host, portNum 127.214.169.149 46470
2025/09/12 00:52:58 simple.go:181:   Username: 
2025/09/12 00:52:58 simple.go:182:   Password: 
2025/09/12 00:52:58 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:52:58 simple.go:184:   Protocol: websocket
2025/09/12 00:52:58 simple.go:185:   Timeout: 30s
2025/09/12 00:52:58 client.go:98: url: ws://localhost:38800
2025/09/12 00:52:58 client.go:99: headers: map[X-Proxy-Target-Host:[127.214.169.149] X-Proxy-Target-Port:[46470]]
2025/09/12 00:52:58 client.go:110: url: http://localhost:38800
2025/09/12 00:52:58 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[CAKMjemStom1nntfVuZfBOsm3zQ=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:52:58 simple.go:209: WebSocket代理连接成功：127.214.169.149:46470
2025/09/12 00:52:58 simple.go:248: clienthost: ::1
2025/09/12 00:52:58 simple.go:249: clientport: 41074
2025/09/12 00:52:58 simple.go:278: simple Handle header:
2025/09/12 00:52:58 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:52:58 simple.go:283: HEAD / HTTP/1.1
2025/09/12 00:52:58 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 00:52:58 simple.go:294: Host: www.baidu.com
User-Agent: curl/8.12.1
Accept: */*
Proxy-Connection: Keep-Alive
2025/09/12 00:52:58 http.go:90: method: HEAD
2025/09/12 00:52:58 http.go:91: url: /
2025/09/12 00:52:58 http.go:92: host: www.baidu.com
2025/09/12 00:52:58 http.go:93: proxyHandler header:
2025/09/12 00:52:58 http.go:120: clienthost: 127.0.0.1
2025/09/12 00:52:58 http.go:121: clientport: 41076
2025/09/12 00:52:58 http.go:132: proxyHandler User-Agent : curl/8.12.1
2025/09/12 00:52:58 http.go:132: proxyHandler Accept : */*
2025/09/12 00:52:58 http.go:132: proxyHandler Proxy-Connection : Keep-Alive
2025/09/12 00:52:58 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.214.169.149:46470;host=www.baidu.com;proto=http
2025/09/12 00:52:58 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.214.169.149:46470;host=www.baidu.com;proto=http
2025/09/12 00:52:58 http.go:137: forwardedByList: [{[::1]:10810} {127.214.169.149:46470}]
2025/09/12 00:52:58 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 00:52:58 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 00:52:58 main.go:791: ProxySelector www.baidu.com
2025/09/12 00:52:58 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:52:58 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 00:52:58 http.go:259: 已经修改了代理为websocket ws://localhost:38800
[GIN] 2025/09/12 - 00:52:59 | 200 |     90.6868ms |       127.0.0.1 | HEAD     "/"
2025/09/12 00:52:59 simple.go:57: remote addr: [::1]:41077
2025/09/12 00:52:59 simple.go:79: CONNECT www.baidu.com:443 HTTP/1.1
2025/09/12 00:52:59 simple.go:117: address:www.baidu.com:443
2025/09/12 00:52:59 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com:443
2025/09/12 00:52:59 main.go:791: ProxySelector www.baidu.com:443
2025/09/12 00:52:59 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 00:52:59 simple.go:179: WebSocket Config Details:
2025/09/12 00:52:59 simple.go:180: host, portNum www.baidu.com 443
2025/09/12 00:52:59 simple.go:181:   Username: 
2025/09/12 00:52:59 simple.go:182:   Password: 
2025/09/12 00:52:59 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 00:52:59 simple.go:184:   Protocol: websocket
2025/09/12 00:52:59 simple.go:185:   Timeout: 30s
2025/09/12 00:52:59 client.go:98: url: ws://localhost:38800
2025/09/12 00:52:59 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 00:52:59 client.go:110: url: http://localhost:38800
2025/09/12 00:52:59 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[UTP4AdDGfuJ2qNDxpkDGv0V6FFY=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 00:52:59 simple.go:209: WebSocket代理连接成功：www.baidu.com:443
2025/09/12 00:52:59 http.go:493: WebSocket ForwardData error: read tcp [::1]:41059->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:52:59 simple.go:204: WebSocket ForwardData error: read tcp [::1]:41075->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:52:59 simple.go:204: WebSocket ForwardData error: read tcp [::1]:41057->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 00:52:59 simple.go:204: WebSocket ForwardData error: read tcp [::1]:41078->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放
