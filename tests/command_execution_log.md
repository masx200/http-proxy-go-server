=== 命令执行日志文件 === 创建时间: 2025-09-14 16:33:04

开始运行命令... [2025-09-14 16:33:04.709] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:33:04.709] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 失败
进程PID: 28168
执行时间: 703.125ms
错误: exit status 1
---
```

开始运行命令... [2025-09-14 16:33:32.448] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:33:32.449] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 失败
进程PID: 4892
执行时间: 937.5ms
错误: exit status 1
---
```

开始运行命令... [2025-09-14 16:34:02.786] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:34:02.786] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 35120
执行时间: 1.453125s
---
```

开始运行命令... [2025-09-14 16:34:03.653] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:34:03.653] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 26084
执行时间: 1.4375s
---
```

开始运行命令... [2025-09-14 16:34:05.134] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:34:05.134] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 17468
---
```

开始运行命令... [2025-09-14 16:34:06.140] [执行命令] ./main.exe -port 10810
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:34:06.140] [HTTP] ./main.exe -port 10810
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 19940
---
```

开始运行命令... [2025-09-14 16:34:09.307] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:10810

开始运行命令... [2025-09-14 16:34:09.307] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:10810

```
执行结果: 成功
进程PID: 24872
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
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
< Date: Sun, 14 Sep 2025 08:33:52 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11219641377149344054
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:33:52 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11219641377149344054
---
```

开始运行命令... [2025-09-14 16:34:09.383] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

开始运行命令... [2025-09-14 16:34:09.383] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

```
执行结果: 成功
进程PID: 7744
执行时间: 62.5ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
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
< Date: Sun, 14 Sep 2025 08:33:52 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9324045590230278964
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
Date: Sun, 14 Sep 2025 08:33:52 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9324045590230278964
---
```

开始运行命令... [2025-09-14 16:34:13.635] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

## [2025-09-14 16:34:13] [BUILD] go build -o main.exe ../cmd/main.go 执行结果: 成功 进程PID: 3932 执行时间: 2025-09-14 16:34:14 输出: 错误: 无

开始运行命令... [2025-09-14 16:34:14.928] [执行命令] ./main.exe --port 18080
-dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6
-dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3

## [2025-09-14 16:34:14] [SERVER] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 执行结果: 成功 进程PID: 34016 执行时间: 2025-09-14 16:34:14 输出: 错误: 无
