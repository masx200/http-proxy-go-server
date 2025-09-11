

###

# WebSocket和http级联代理测试记录

## 测试时间
2025-09-12 01:06:46

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./http-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 44364

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动http服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 http服务器进程PID: 11848

等待http服务器启动...
✅ http服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 28820, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 17:06:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11390912182963063807
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 17:06:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11390912182963063807


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 53320, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 17:06:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10799265538454071569
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
Date: Thu, 11 Sep 2025 17:06:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10799265538454071569


```

### 📋 所有进程PID记录

所有进程PID: 44364, 11848, 28820, 53320



###

# WebSocket和http级联代理测试记录

## 测试时间
2025-09-12 01:06:46

## 1. 编译代理服务器

执行命令: `go build -o main.exe ../cmd/main.go`

✅ 代理服务器编译成功

## 2. 启动WebSocket服务器（上游）

执行命令: `./http-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800`

📋 WebSocket服务器进程PID: 44364

等待WebSocket服务器启动...
✅ WebSocket服务器启动成功

## 3. 启动http服务器（下游）

执行命令: `./main.exe  -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800`

📋 http服务器进程PID: 11848

等待http服务器启动...
✅ http服务器启动成功

## 4. 测试级联代理功能

### 测试1: HTTP代理通过级联

执行命令: `curl -v -I http://www.baidu.com -x http://localhost:10810`

📋 Curl测试1进程PID: 28820, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 17:06:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11390912182963063807
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Thu, 11 Sep 2025 17:06:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11390912182963063807


```

### 测试2: HTTPS代理通过级联

执行命令: `curl -v -I https://www.baidu.com -x http://localhost:10810`

📋 Curl测试2进程PID: 53320, 退出状态码: 0

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
< Date: Thu, 11 Sep 2025 17:06:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10799265538454071569
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
Date: Thu, 11 Sep 2025 17:06:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10799265538454071569


```

### 📋 所有进程PID记录

所有进程PID: 44364, 11848, 28820, 53320

## 5. 关闭服务器

✅ 所有测试成功，正在关闭服务器进程...

🛑 正在终止WebSocket服务器进程...
✅ WebSocket服务器进程已终止

🛑 正在终止http服务器进程...
✅ http服务器进程已终止

🧹 正在清理所有子进程...
✅ 所有子进程已清理完成

🧹 已清理编译的可执行文件
### WebSocket服务器日志输出

```
2025/09/12 01:06:48 main.go:71: 启动websocket服务端，监听地址: :38800
2025/09/12 01:06:48 server.go:71: [WEBSOCKET-SERVER] Server started successfully, listening on :38800
2025/09/12 01:06:48 server.go:72: [WEBSOCKET-SERVER] Authentication enabled: false (0 users configured)
2025/09/12 01:06:48 server.go:74: [WEBSOCKET-SERVER] Upstream selector enabled: false
2025/09/12 01:06:48 server.go:75: [WEBSOCKET-SERVER] Read timeout: 30s, Write timeout: 30s
2025/09/12 01:06:48 main.go:129: websocket服务端已启动，按Ctrl+C停止
2025/09/12 01:06:50 server.go:90: url /
2025/09/12 01:06:50 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[Q+c+4+ogH63zilb5/u2r8Q==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.170.111.185] X-Proxy-Target-Port:[36602]]
2025/09/12 01:06:50 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65438 at 2025-09-12 01:06:50
2025/09/12 01:06:50 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65438
2025/09/12 01:06:50 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.170.111.185', targetPort: 36602
2025/09/12 01:06:50 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 01:06:50 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65438
2025/09/12 01:06:50 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.170.111.185:36602 from [::1]:65438
2025/09/12 01:06:50 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.170.111.185:36602 (timeout: 30s)
2025/09/12 01:06:50 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.170.111.185:36602
2025/09/12 01:06:50 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 01:06:50 server.go:90: url /
2025/09/12 01:06:50 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[5uIJQ7zM3woJDVI9QFmzmA==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 01:06:50 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:65440 at 2025-09-12 01:06:50
2025/09/12 01:06:50 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:65440
2025/09/12 01:06:50 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 80
2025/09/12 01:06:50 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 01:06:50 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:65440
2025/09/12 01:06:50 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:80 from [::1]:65440
2025/09/12 01:06:50 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:80 (timeout: 30s)
2025/09/12 01:06:50 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:80
2025/09/12 01:06:50 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 01:06:52 server.go:90: url /
2025/09/12 01:06:52 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[+3qPdYokomSTafC6uzaSXg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[127.170.111.185] X-Proxy-Target-Port:[36602]]
2025/09/12 01:06:52 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:63607 at 2025-09-12 01:06:52
2025/09/12 01:06:52 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:63607
2025/09/12 01:06:52 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: '127.170.111.185', targetPort: 36602
2025/09/12 01:06:52 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 01:06:52 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:63607
2025/09/12 01:06:52 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target 127.170.111.185:36602 from [::1]:63607
2025/09/12 01:06:52 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target 127.170.111.185:36602 (timeout: 30s)
2025/09/12 01:06:52 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target 127.170.111.185:36602
2025/09/12 01:06:52 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
2025/09/12 01:06:52 server.go:90: url /
2025/09/12 01:06:52 server.go:92: headers map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Sec-Websocket-Key:[qu+b6814lxB1XaTfE4Mq5g==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1] X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 01:06:52 server.go:95: [WEBSOCKET-CONN] New connection attempt from [::1]:63610 at 2025-09-12 01:06:52
2025/09/12 01:06:52 server.go:101: [WEBSOCKET-CONN] No authentication required for client [::1]:63610
2025/09/12 01:06:52 server.go:306: [WEBSOCKET-AUTH] Parsed auth info - username: '', password: '', targetHost: 'www.baidu.com', targetPort: 443
2025/09/12 01:06:52 server.go:173: [WEBSOCKET-AUTH] No authentication configured, allowing access for user ''
2025/09/12 01:06:52 server.go:119: [WEBSOCKET-AUTH] Authentication successful for user '' from [::1]:63610
2025/09/12 01:06:52 server.go:129: [WEBSOCKET-CONN] WebSocket connection established successfully for target www.baidu.com:443 from [::1]:63610
2025/09/12 01:06:52 server.go:227: [WEBSOCKET-UPSTREAM] Using direct connection for target www.baidu.com:443 (timeout: 30s)
2025/09/12 01:06:52 server.go:232: [WEBSOCKET-UPSTREAM] Direct connection established for target www.baidu.com:443
2025/09/12 01:06:52 server.go:316: [WEBSOCKET-FORWARD] Starting data forwarding between connections
```

### http服务器日志输出

```
2025/09/12 01:06:49 main.go:518: 代理服务器启动中...
2025/09/12 01:06:49 main.go:562: dohalpn: []
2025/09/12 01:06:49 main.go:564: hostname: 0.0.0.0
2025/09/12 01:06:49 main.go:566: port: 10810
2025/09/12 01:06:49 main.go:568: server_cert: 
2025/09/12 01:06:49 main.go:570: server_key: 
2025/09/12 01:06:49 main.go:572: username: 
2025/09/12 01:06:49 main.go:574: password: 
2025/09/12 01:06:49 main.go:576: dohurl: []
2025/09/12 01:06:49 main.go:578: dohip: []
2025/09/12 01:06:49 main.go:579: upstream-type: websocket
2025/09/12 01:06:49 main.go:580: upstream-address: ws://localhost:38800
2025/09/12 01:06:49 main.go:581: upstream-username: 
2025/09/12 01:06:49 main.go:582: upstream-password: 
2025/09/12 01:06:49 main.go:639: WebSocket代理配置已添加
2025/09/12 01:06:49 main.go:854: {
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
2025/09/12 01:06:49 simple.go:31: Proxy server started on port [::]:10810
2025/09/12 01:06:49 http.go:372: Random IP: 127.170.111.185
2025/09/12 01:06:49 http.go:390: Random integer: 36602
2025/09/12 01:06:49 http.go:342: Proxy server started on port 127.170.111.185:36602
2025/09/12 01:06:50 simple.go:57: remote addr: [::1]:65437
2025/09/12 01:06:50 simple.go:79: GET http://www.baidu.com/ HTTP/1.1
2025/09/12 01:06:50 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 01:06:50 simple.go:117: address:www.baidu.com:80
2025/09/12 01:06:50 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.170.111.185:36602
2025/09/12 01:06:50 main.go:791: ProxySelector 127.170.111.185:36602
2025/09/12 01:06:50 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 01:06:50 simple.go:179: WebSocket Config Details:
2025/09/12 01:06:50 simple.go:180: host, portNum 127.170.111.185 36602
2025/09/12 01:06:50 simple.go:181:   Username: 
2025/09/12 01:06:50 simple.go:182:   Password: 
2025/09/12 01:06:50 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 01:06:50 simple.go:184:   Protocol: websocket
2025/09/12 01:06:50 simple.go:185:   Timeout: 30s
2025/09/12 01:06:50 client.go:98: url: ws://localhost:38800
2025/09/12 01:06:50 client.go:99: headers: map[X-Proxy-Target-Host:[127.170.111.185] X-Proxy-Target-Port:[36602]]
2025/09/12 01:06:50 client.go:110: url: http://localhost:38800
2025/09/12 01:06:50 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[EZpLXEvpHG6jVBGbdWoke/bpucE=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 01:06:50 simple.go:209: WebSocket代理连接成功：127.170.111.185:36602
2025/09/12 01:06:50 simple.go:248: clienthost: ::1
2025/09/12 01:06:50 simple.go:249: clientport: 65437
2025/09/12 01:06:50 simple.go:278: simple Handle header:
2025/09/12 01:06:50 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 01:06:50 simple.go:283: GET / HTTP/1.1
2025/09/12 01:06:50 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 01:06:50 simple.go:294: Host: www.baidu.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
2025/09/12 01:06:50 http.go:90: method: GET
2025/09/12 01:06:50 http.go:91: url: /
2025/09/12 01:06:50 http.go:92: host: www.baidu.com
2025/09/12 01:06:50 http.go:93: proxyHandler header:
2025/09/12 01:06:50 http.go:120: clienthost: 127.0.0.1
2025/09/12 01:06:50 http.go:121: clientport: 65439
2025/09/12 01:06:50 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.170.111.185:36602;host=www.baidu.com;proto=http
2025/09/12 01:06:50 http.go:132: proxyHandler User-Agent : Go-http-client/1.1
2025/09/12 01:06:50 http.go:132: proxyHandler Accept-Encoding : gzip
2025/09/12 01:06:50 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.170.111.185:36602;host=www.baidu.com;proto=http
2025/09/12 01:06:50 http.go:137: forwardedByList: [{[::1]:10810} {127.170.111.185:36602}]
2025/09/12 01:06:50 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 01:06:50 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 01:06:50 main.go:791: ProxySelector www.baidu.com
2025/09/12 01:06:50 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 01:06:50 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 01:06:50 http.go:259: 已经修改了代理为websocket ws://localhost:38800
2025/09/12 01:06:50 http.go:262: 使用代理：ws://localhost:38800
2025/09/12 01:06:50 http.go:264: network,addr tcp www.baidu.com:80
2025/09/12 01:06:50 http.go:466: WebSocket Config Details:
2025/09/12 01:06:50 http.go:467: host, portNum www.baidu.com 80
2025/09/12 01:06:50 http.go:468:   Username: 
2025/09/12 01:06:50 http.go:469:   Password: 
2025/09/12 01:06:50 http.go:470:   ServerAddr: ws://localhost:38800
2025/09/12 01:06:50 http.go:471:   Protocol: websocket
2025/09/12 01:06:50 http.go:472:   Timeout: 30s
2025/09/12 01:06:50 http.go:476: host, portNum www.baidu.com 80
2025/09/12 01:06:50 client.go:98: url: ws://localhost:38800
2025/09/12 01:06:50 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[80]]
2025/09/12 01:06:50 client.go:110: url: http://localhost:38800
2025/09/12 01:06:50 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[4tfK7F5N1zDyJtuvLo5sn2orZGc=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
[GIN] 2025/09/12 - 01:06:50 | 200 |     64.6947ms |       127.0.0.1 | GET      "/"
2025/09/12 01:06:52 simple.go:57: remote addr: [::1]:63606
2025/09/12 01:06:52 simple.go:79: HEAD http://www.baidu.com/ HTTP/1.1
2025/09/12 01:06:52 simple.go:305: Domain: www.baidu.com, Port: 80
2025/09/12 01:06:52 simple.go:117: address:www.baidu.com:80
2025/09/12 01:06:52 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy 127.170.111.185:36602
2025/09/12 01:06:52 main.go:791: ProxySelector 127.170.111.185:36602
2025/09/12 01:06:52 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 01:06:52 simple.go:179: WebSocket Config Details:
2025/09/12 01:06:52 simple.go:180: host, portNum 127.170.111.185 36602
2025/09/12 01:06:52 simple.go:181:   Username: 
2025/09/12 01:06:52 simple.go:182:   Password: 
2025/09/12 01:06:52 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 01:06:52 simple.go:184:   Protocol: websocket
2025/09/12 01:06:52 simple.go:185:   Timeout: 30s
2025/09/12 01:06:52 client.go:98: url: ws://localhost:38800
2025/09/12 01:06:52 client.go:99: headers: map[X-Proxy-Target-Host:[127.170.111.185] X-Proxy-Target-Port:[36602]]
2025/09/12 01:06:52 client.go:110: url: http://localhost:38800
2025/09/12 01:06:52 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[BWpGU9wy9u8UEW/IQnA+XGwpKdE=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 01:06:52 simple.go:209: WebSocket代理连接成功：127.170.111.185:36602
2025/09/12 01:06:52 simple.go:248: clienthost: ::1
2025/09/12 01:06:52 simple.go:249: clientport: 63606
2025/09/12 01:06:52 simple.go:278: simple Handle header:
2025/09/12 01:06:52 simple.go:281: simple Handle Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 01:06:52 simple.go:283: HEAD / HTTP/1.1
2025/09/12 01:06:52 simple.go:288: Forwarded: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http
2025/09/12 01:06:52 simple.go:294: Host: www.baidu.com
User-Agent: curl/8.12.1
Accept: */*
Proxy-Connection: Keep-Alive
2025/09/12 01:06:52 http.go:90: method: HEAD
2025/09/12 01:06:52 http.go:91: url: /
2025/09/12 01:06:52 http.go:92: host: www.baidu.com
2025/09/12 01:06:52 http.go:93: proxyHandler header:
2025/09/12 01:06:52 http.go:120: clienthost: 127.0.0.1
2025/09/12 01:06:52 http.go:121: clientport: 63608
2025/09/12 01:06:52 http.go:132: proxyHandler Forwarded : for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http,for=127.0.0.1;by=127.170.111.185:36602;host=www.baidu.com;proto=http
2025/09/12 01:06:52 http.go:132: proxyHandler User-Agent : curl/8.12.1
2025/09/12 01:06:52 http.go:132: proxyHandler Accept : */*
2025/09/12 01:06:52 http.go:132: proxyHandler Proxy-Connection : Keep-Alive
2025/09/12 01:06:52 http.go:135: forwardedHeader: for=::1;by=[::1]:10810;host=www.baidu.com:80;proto=http, for=127.0.0.1;by=127.170.111.185:36602;host=www.baidu.com;proto=http
2025/09/12 01:06:52 http.go:137: forwardedByList: [{[::1]:10810} {127.170.111.185:36602}]
2025/09/12 01:06:52 http.go:155: targetUrl: http://www.baidu.com/
2025/09/12 01:06:52 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com
2025/09/12 01:06:52 main.go:791: ProxySelector www.baidu.com
2025/09/12 01:06:52 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 01:06:52 http.go:250: 使用代理：ws://localhost:38800
2025/09/12 01:06:52 http.go:259: 已经修改了代理为websocket ws://localhost:38800
[GIN] 2025/09/12 - 01:06:52 | 200 |     14.7514ms |       127.0.0.1 | HEAD     "/"
2025/09/12 01:06:52 simple.go:57: remote addr: [::1]:63609
2025/09/12 01:06:52 simple.go:79: CONNECT www.baidu.com:443 HTTP/1.1
2025/09/12 01:06:52 simple.go:117: address:www.baidu.com:443
2025/09/12 01:06:52 CheckShouldUseProxy.go:10: 开始检查CheckShouldUseProxy www.baidu.com:443
2025/09/12 01:06:52 main.go:791: ProxySelector www.baidu.com:443
2025/09/12 01:06:52 main.go:797: 选择的代理 URL: ws://localhost:38800
2025/09/12 01:06:52 simple.go:179: WebSocket Config Details:
2025/09/12 01:06:52 simple.go:180: host, portNum www.baidu.com 443
2025/09/12 01:06:52 simple.go:181:   Username: 
2025/09/12 01:06:52 simple.go:182:   Password: 
2025/09/12 01:06:52 simple.go:183:   ServerAddr: ws://localhost:38800
2025/09/12 01:06:52 simple.go:184:   Protocol: websocket
2025/09/12 01:06:52 simple.go:185:   Timeout: 30s
2025/09/12 01:06:52 client.go:98: url: ws://localhost:38800
2025/09/12 01:06:52 client.go:99: headers: map[X-Proxy-Target-Host:[www.baidu.com] X-Proxy-Target-Port:[443]]
2025/09/12 01:06:52 client.go:110: url: http://localhost:38800
2025/09/12 01:06:52 client.go:111: headers: map[Connection:[Upgrade] Sec-Websocket-Accept:[2lafXh5NGz0zN4QRpGmOMLKbYgU=] Sec-Websocket-Extensions:[permessage-deflate; server_no_context_takeover; client_no_context_takeover] Upgrade:[websocket]]
2025/09/12 01:06:52 simple.go:209: WebSocket代理连接成功：www.baidu.com:443
2025/09/12 01:06:52 http.go:493: WebSocket ForwardData error: read tcp [::1]:65440->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 01:06:52 simple.go:204: WebSocket ForwardData error: read tcp [::1]:65438->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 01:06:52 simple.go:204: WebSocket ForwardData error: read tcp [::1]:63610->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
2025/09/12 01:06:52 simple.go:204: WebSocket ForwardData error: read tcp [::1]:63607->[::1]:38800: wsarecv: An existing connection was forcibly closed by the remote host.
```

✅ 端口38800已成功释放
✅ 端口10810已成功释放
