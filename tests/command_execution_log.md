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

开始运行命令... [2025-09-14 16:34:06.140] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:34:06.140] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 19940
---
```

开始运行命令... [2025-09-14 16:34:09.307] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:34:09.307] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 24872
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
18080
  0     0    0     0    0     0     18080   0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:10810...
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

开始运行命令... [2025-09-14 16:34:09.383] [执行命令]18080
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

开始运行命令... [2025-09-14 16:34:09.383] [CURL]18080
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

```
执行结果: 成功
进程PID: 7744
执行时间: 62.5ms18080
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current18080
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
*   Certificate level 0: Public key 18080? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
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

## [2025-09-14 16:34:13] [BUILD] go build -o main.exe ../cmd/main.go 执行结果: 成功 进程 PID: 3932 执行时间: 2025-09-14 16:34:14 输出: 错误: 无

开始运行命令... [2025-09-14 16:34:14.928] [执行命令] ./main.exe --port 18080
-dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6
-dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3

## [2025-09-14 16:34:14] [SERVER] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 执行结果: 成功 进程 PID: 34016 执行时间: 2025-09-14 16:34:14 输出: 错误: 无

开始运行命令... [2025-09-14 16:36:02.203] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:36:02.204] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 32772
执行时间: 1.8125s
---
```

开始运行命令... [2025-09-14 16:36:03.034] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:36:03.034] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 22532
执行时间: 1.6875s
---
```

开始运行命令... [2025-09-14 16:36:04.996] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:36:04.996] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 913218080
---
```

开始运行命令... [2025-09-14 16:36:06.001] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:36:06.001] [HTTP] ./main.exe -port 10810
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 2825618080
---
```

开始运行命令... [2025-09-14 16:36:09.154] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:36:09.154] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:10810

```
执行结果: 成功18080
进程PID: 3676
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:10810 was resolved.
* IPv6: ::118080
* IPv4: 127.0.0.118080
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

< Date: Sun, 14 Sep 2025 08:35:52 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache
18080
< Server: bfe/1.0.8.18

< Tr_id: bfe_11552078580505659851
18080
<


  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: pr18080, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:35:52 GMT
Etag: "575e1f60-115"18080
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11552078580505659851
---
```

开始运行命令... [2025-09-14 16:36:09.223] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

开始运行命令... [2025-09-14 16:36:09.223] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:10810

```
执行结果: 成功
进程PID: 12576
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
} [308 bytes data]18080
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

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
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

< Date: Sun, 14 Sep 2025 08:35:52 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_12104553899071472023

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
Date: Sun, 14 Sep 2025 08:35:52 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12104553899071472023
---
```

开始运行命令... [2025-09-14 16:36:13.601] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

[2025-09-14 16:36:13] [BUILD] go build -o main.exe ../cmd/main.go 执行结果: 成功
进程 PID: 20632 执行时间: 2025-09-14 16:36:14 输出: 错误: 无

---

开始运行命令... [2025-09-14 16:36:14.942] [执行命令] ./main.exe --port 18080
-dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6
-dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3

[2025-09-14 16:36:14] [SERVER] ./main.exe --port 18080 -dohurl
https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl
https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 执行结果: 成功 进程
PID: 30296 执行时间: 2025-09-14 16:36:14 输出: 错误: 无

---

开始运行命令... [2025-09-14 16:36:17.115] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

[2025-09-14 16:36:17] [TEST] curl -v -I http://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 6644 执行时间: 2025-09-14
16:36:17 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD http://www.baidu.com/ HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

> Proxy-Connection: Keep-Alive

- Request completely sent off < HTTP/1.1 200 OK

< Bdpagetype: 1

< Bdqid: 0xd9660f850008ae5c

< Connection: keep-alive

< Content-Length: 654184

< Content-Type: text/html; charset=utf-8

< Date: Sun, 14 Sep 2025 08:36:00 GMT

< Server: BWS/1.1

< Set-Cookie: BIDUPSID=85E2B56578ACC1982A8E7DC8C9689113; expires=Thu, 31-Dec-37
23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: PSTM=1757838960; expires=Thu, 31-Dec-37 23:55:55 GMT;
max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: BDSVRTM=1; path=/

< Set-Cookie: BD_HOME=1; path=/

< Set-Cookie: BAIDUID=85E2B56578ACC1982A8E7DC8C9689113:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000

< Set-Cookie: BAIDUID_BFESS=85E2B56578ACC1982A8E7DC8C9689113:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None

< Tr_id: super_0xd9660f850008ae5c

< Traceid: 1757838960070829261815665225417760484956

< Vary: Accept-Encoding

< X-Ua-Compatible: IE=Edge,chrome=1

< X-Xss-Protection: 1;mode=block

<

0 638k 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Bdpagetype: 1
  Bdqid: 0xd9660f850008ae5c Connection: keep-alive Content-Length: 654184
  Content-Type: text/html; charset=utf-8 Date: Sun, 14 Sep 2025 08:36:00 GMT
  Server: BWS/1.1 Set-Cookie: BIDUPSID=85E2B56578ACC1982A8E7DC8C9689113;
  expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/;
  domain=.baidu.com Set-Cookie: PSTM=1757838960; expires=Thu, 31-Dec-37 23:55:55
  GMT; max-age=2147483647; path=/; domain=.baidu.com Set-Cookie: BDSVRTM=1;
  path=/ Set-Cookie: BD_HOME=1; path=/ Set-Cookie:
  BAIDUID=85E2B56578ACC1982A8E7DC8C9689113:FG=1; Path=/; Domain=baidu.com;
  Max-Age=31536000 Set-Cookie:
  BAIDUID_BFESS=85E2B56578ACC1982A8E7DC8C9689113:FG=1; Path=/; Domain=baidu.com;
  Max-Age=31536000; Secure; SameSite=None Tr_id: super_0xd9660f850008ae5c
  Traceid: 1757838960070829261815665225417760484956 Vary: Accept-Encoding
  X-Ua-Compatible: IE=Edge,chrome=1 X-Xss-Protection: 1;mode=block

## 错误: 无

开始运行命令... [2025-09-14 16:36:17.246] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

[2025-09-14 16:36:17] [TEST] curl -v -I -L http://www.so.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 4540 执行时间: 2025-09-14
16:36:17 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD http://www.so.com/ HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: _/_

> Proxy-Connection: Keep-Alive

- Request completely sent off < HTTP/1.1 302 Found

< Connection: keep-alive

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:36:00 GMT

< Location: https://www.so.com/

< Server: openresty

< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

- Ignoring the response-body <

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0

- Connection #0 to host localhost left intact
- Clear auth, redirects to port from 80 to 443
- Issue another request to this URL: 'https://www.so.com/'
- Hostname localhost was found in DNS cache
- Trying [::1]:18080...
- CONNECT tunnel: HTTP/1.1 negotiated
- allocate connect buffer
- Establish HTTP proxy tunnel to www.so.com:443
  > CONNECT www.so.com:443 HTTP/1.1

> Host: www.so.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

< HTTP/1.1 200 Connection established

<

- CONNECT phase completed
- CONNECT tunnel established, response 200
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [305 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [93 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [5077 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server did not agree on a protocol. Uses default.
- Server certificate:
- subject: CN=\*.so.com
- start date: Aug 28 00:00:00 2025 GMT
- expire date: Sep 28 23:59:59 2026 GMT
- subjectAltName: host "www.so.com" matched cert's "\*.so.com"
- issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA [Run by the Issuer]
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha384WithRSAEncryption
- Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using
  sha384WithRSAEncryption
- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD / HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Server: openresty

< Date: Sun, 14 Sep 2025 08:36:00 GMT

< Content-Type: text/html; charset=UTF-8

< Connection: keep-alive

< Vary: Accept-Encoding

< Set-Cookie: \_S=gs1jl5436h8ghcl1ebvs1a8a47; expires=Sun, 14-Sep-2025 08:46:00
GMT; Max-Age=600; path=/

< Expires: Thu, 19 Nov 1981 08:52:00 GMT

< Cache-Control: no-store, no-cache, must-revalidate

< Pragma: no-cache

< php-waf-rep: -

< Set-Cookie: QiHooGUID=3C8B5EF932A65A6EE662613043B99A8F.1757838960597;
Max-Age=63072000; Domain=so.com; Path=/

<

0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #1 to host localhost left intact HTTP/1.1 302 Found Connection:
  keep-alive Content-Type: text/html Date: Sun, 14 Sep 2025 08:36:00 GMT
  Location: https://www.so.com/ Server: openresty Set-Cookie: QiHooGUID=;
  Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK Server: openresty Date: Sun, 14 Sep 2025 08:36:00 GMT
Content-Type: text/html; charset=UTF-8 Connection: keep-alive Vary:
Accept-Encoding Set-Cookie: \_S=gs1jl5436h8ghcl1ebvs1a8a47; expires=Sun,
14-Sep-2025 08:46:00 GMT; Max-Age=600; path=/ Expires: Thu, 19 Nov 1981 08:52:00
GMT Cache-Control: no-store, no-cache, must-revalidate Pragma: no-cache
php-waf-rep: - Set-Cookie:
QiHooGUID=3C8B5EF932A65A6EE662613043B99A8F.1757838960597; Max-Age=63072000;
Domain=so.com; Path=/

## 错误: 无

开始运行命令... [2025-09-14 16:36:17.493] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

[2025-09-14 16:36:17] [TEST] curl -v -I https://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 29368 执行时间: 2025-09-14
16:36:17 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- CONNECT tunnel: HTTP/1.1 negotiated
- allocate connect buffer
- Establish HTTP proxy tunnel to www.baidu.com:443
  > CONNECT www.baidu.com:443 HTTP/1.1

> Host: www.baidu.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

< HTTP/1.1 200 Connection established

<

- CONNECT phase completed
- CONNECT tunnel established, response 200
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [308 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [102 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [4771 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server accepted http/1.1
- Server certificate:
- subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science
  Technology Co., Ltd; CN=baidu.com
- start date: Jul 9 07:01:02 2025 GMT
- expire date: Aug 10 07:01:01 2026 GMT
- subjectAltName: host "www.baidu.com" matched cert's "\*.baidu.com"
- issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: no-cache

< Connection: keep-alive

< Content-Length: 227

< Content-Security-Policy: frame-ancestors 'self' https://chat.baidu.com
http://mirror-chat.baidu.com https://fj-chat.baidu.com
https://hba-chat.baidu.com https://hbe-chat.baidu.com
https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com
https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com
https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com
https://platform-openai.now.baidu.com;

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:36:00 GMT

< Pragma: no-cache

< Server: BWS/1.1

< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300

< Set-Cookie: PSTM=1757838960; expires=Thu, 31-Dec-37 23:55:55 GMT;
max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: BAIDUID=29F0BD33381FB0886EEE60600E18C2B5:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000

< Set-Cookie: BAIDUID_BFESS=29F0BD33381FB0886EEE60600E18C2B5:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None

< Traceid: 1757838960044555879413448146751056025995

< X-Ua-Compatible: IE=Edge,chrome=1

< X-Xss-Protection: 1;mode=block

<

0 227 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 Connection
  established

HTTP/1.1 200 OK Accept-Ranges: bytes Cache-Control: no-cache Connection:
keep-alive Content-Length: 227 Content-Security-Policy: frame-ancestors 'self'
https://chat.baidu.com http://mirror-chat.baidu.com https://fj-chat.baidu.com
https://hba-chat.baidu.com https://hbe-chat.baidu.com
https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com
https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com
https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com
https://platform-openai.now.baidu.com; Content-Type: text/html Date: Sun, 14 Sep
2025 08:36:00 GMT Pragma: no-cache Server: BWS/1.1 Set-Cookie: BD_NOT_HTTPS=1;
path=/; Max-Age=300 Set-Cookie: PSTM=1757838960; expires=Thu, 31-Dec-37 23:55:55
GMT; max-age=2147483647; path=/; domain=.baidu.com Set-Cookie:
BAIDUID=29F0BD33381FB0886EEE60600E18C2B5:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000 Set-Cookie:
BAIDUID_BFESS=29F0BD33381FB0886EEE60600E18C2B5:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000; Secure; SameSite=None Traceid:
1757838960044555879413448146751056025995 X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block

## 错误: 无

开始运行命令... [2025-09-14 16:36:20.863] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:36:20.863] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 11568
执行时间: 1.671875s
---
```

开始运行命令... [2025-09-14 16:36:22.209] [执行命令] ./main.exe --port 18080

开始运行命令... [2025-09-14 16:36:25.316] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:36:25.316] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 15852
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
* Connected to localhost (::1) port 18080
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: */*

> Proxy-Connection: Keep-Alive

>

* Request completely sent off
< HTTP/1.1 302 Found

< Content-Type: text/plain; charset=utf-8

< Date: Sun, 14 Sep 2025 08:36:09 GMT

< Location: https://www.baidu.com/error.html

< Server: bfe

<


  0     0    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
  0     0    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 302 Found
Content-Type: text/plain; charset=utf-8
Date: Sun, 14 Sep 2025 08:36:09 GMT
Location: https://www.baidu.com/error.html
Server: bfe
---
```

开始运行命令... [2025-09-14 16:36:26.419] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:36:26.419] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 18732
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
* Connected to localhost (::1) port 18080
* using HTTP/1.x
> HEAD http://www.so.com/ HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: */*

> Proxy-Connection: Keep-Alive

>

* Request completely sent off
< HTTP/1.1 302 Found

< Connection: keep-alive

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:36:09 GMT

< Location: https://www.so.com/

< Server: openresty

< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

* Ignoring the response-body
<


  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:18080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.so.com:443
> CONNECT www.so.com:443 HTTP/1.1

> Host: www.so.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

>

< HTTP/1.1 200 Connection established

<

* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [305 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [93 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [5077 bytes data]
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
* ALPN: server did not agree on a protocol. Uses default.
* Server certificate:
*  subject: CN=*.so.com
*  start date: Aug 28 00:00:00 2025 GMT
*  expire date: Sep 28 23:59:59 2026 GMT
*  subjectAltName: host "www.so.com" matched cert's "*.so.com"
*  issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA  [Run by the Issuer]
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha384WithRSAEncryption
*   Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using sha384WithRSAEncryption
* Connected to localhost (::1) port 18080
* using HTTP/1.x
> HEAD / HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: */*

>

* Request completely sent off
< HTTP/1.1 200 OK

< Server: openresty

< Date: Sun, 14 Sep 2025 08:36:09 GMT

< Content-Type: text/html; charset=UTF-8

< Connection: keep-alive

< Vary: Accept-Encoding

< Set-Cookie: _S=rr4uq1q9b48mfkog0ta85hnp71; expires=Sun, 14-Sep-2025 08:46:09 GMT; Max-Age=600; path=/

< Expires: Thu, 19 Nov 1981 08:52:00 GMT

< Cache-Control: no-store, no-cache, must-revalidate

< Pragma: no-cache

< php-waf-rep: -

< Set-Cookie: QiHooGUID=8C45361319A2FACBEEACF15FA9CE86BC.1757838969776; Max-Age=63072000; Domain=so.com; Path=/

<


  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:36:09 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 08:36:09 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=rr4uq1q9b48mfkog0ta85hnp71; expires=Sun, 14-Sep-2025 08:46:09 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=8C45361319A2FACBEEACF15FA9CE86BC.1757838969776; Max-Age=63072000; Domain=so.com; Path=/
---
```

开始运行命令... [2025-09-14 16:36:26.643] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:36:26.643] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 32840
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
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
* Connected to localhost (::1) port 18080
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

< Date: Sun, 14 Sep 2025 08:36:09 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_9275929367889434610

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
Date: Sun, 14 Sep 2025 08:36:09 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9275929367889434610
---
```

开始运行命令... [2025-09-14 16:36:57.631] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:36:57.631] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 23128
执行时间: 1.234375s
---
```

开始运行命令... [2025-09-14 16:36:58.027] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:36:58.027] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 24744
执行时间: 1.140625s
---
```

开始运行命令... [2025-09-14 16:36:59.235] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:36:59.235] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 25996
---
```

开始运行命令... [2025-09-14 16:37:00.240] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:37:00.240] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 34092
---
```

开始运行命令... [2025-09-14 16:37:20.121] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:37:20.121] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 30488
执行时间: 1.5s
---
```

开始运行命令... [2025-09-14 16:37:20.545] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:37:20.545] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 34764
执行时间: 890.625ms
---
```

开始运行命令... [2025-09-14 16:37:21.738] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:37:21.738] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 21192
---
```

开始运行命令... [2025-09-14 16:37:22.744] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:37:22.744] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 11752
---
```

开始运行命令... [2025-09-14 16:37:37.608] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:37:37.608] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 6672
执行时间: 1.15625s
---
```

开始运行命令... [2025-09-14 16:37:38.025] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:37:38.025] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 33596
执行时间: 1.140625s
---
```

开始运行命令... [2025-09-14 16:37:39.239] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:37:39.239] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 34932
---
```

开始运行命令... [2025-09-14 16:37:40.245] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:37:40.245] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 6644
---
```

开始运行命令... [2025-09-14 16:38:39.226] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:38:39.226] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 10552
执行时间: 1.03125s
---
```

开始运行命令... [2025-09-14 16:38:39.661] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:38:39.661] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 33512
执行时间: 1.0625s
---
```

开始运行命令... [2025-09-14 16:38:40.938] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

开始运行命令... [2025-09-14 16:38:40.938] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:18080

```
执行结果: 成功
进程PID: 25556
---
```

开始运行命令... [2025-09-14 16:38:41.943] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

开始运行命令... [2025-09-14 16:38:41.943] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:18080

```
执行结果: 成功
进程PID: 33204
---
```

开始运行命令... [2025-09-14 16:42:30.599] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:42:30.599] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 35696
执行时间: 1.1875s
---
```

开始运行命令... [2025-09-14 16:42:30.998] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:42:30.998] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 30912
执行时间: 1.59375s
---
```

开始运行命令... [2025-09-14 16:42:32.368] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:28080

开始运行命令... [2025-09-14 16:42:32.368] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:28080

```
执行结果: 成功
进程PID: 26156
---
```

开始运行命令... [2025-09-14 16:42:32.374] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:28080

开始运行命令... [2025-09-14 16:42:32.374] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:28080

```
执行结果: 成功
进程PID: 12320
---
```

开始运行命令... [2025-09-14 16:44:12.168] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

开始运行命令... [2025-09-14 16:44:12.168] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o socks5-websocket-proxy-golang.exe
github.com/masx200/socks5-websocket-proxy-golang/cmd

```
执行结果: 成功
进程PID: 32584
执行时间: 1.453125s
---
```

开始运行命令... [2025-09-14 16:44:12.606] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:44:12.606] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 29460
执行时间: 1.453125s
---
```

开始运行命令... [2025-09-14 16:44:14.058] [执行命令]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:28081

开始运行命令... [2025-09-14 16:44:14.058] [WEBSOCKET]
./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr
:28081

```
执行结果: 成功
进程PID: 31592
---
```

开始运行命令... [2025-09-14 16:44:15.064] [执行命令] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:28081

开始运行命令... [2025-09-14 16:44:15.064] [HTTP] ./main.exe -port 18080
-upstream-type websocket -upstream-address ws://localhost:28081

```
执行结果: 成功
进程PID: 31948
---
```

开始运行命令... [2025-09-14 16:44:18.245] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:44:18.245] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 26772
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
* Connected to localhost (::1) port 18080
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

< Date: Sun, 14 Sep 2025 08:44:01 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_10901102181150578156

<


  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:44:01 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10901102181150578156
---
```

开始运行命令... [2025-09-14 16:44:18.313] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:44:18.313] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 15992
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
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
* Connected to localhost (::1) port 18080
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

< Date: Sun, 14 Sep 2025 08:44:01 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_11268909660045715536

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
Date: Sun, 14 Sep 2025 08:44:01 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11268909660045715536
---
```

开始运行命令... [2025-09-14 16:44:22.666] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

[2025-09-14 16:44:22] [BUILD] go build -o main.exe ../cmd/main.go 执行结果: 成功
进程 PID: 29736 执行时间: 2025-09-14 16:44:23 输出: 错误: 无

---

开始运行命令... [2025-09-14 16:44:23.915] [执行命令] ./main.exe --port 18080
-dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6
-dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3

[2025-09-14 16:44:23] [SERVER] ./main.exe --port 18080 -dohurl
https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl
https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 执行结果: 成功 进程
PID: 26348 执行时间: 2025-09-14 16:44:23 输出: 错误: 无

---

开始运行命令... [2025-09-14 16:44:26.083] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

[2025-09-14 16:44:26] [TEST] curl -v -I http://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 6904 执行时间: 2025-09-14
16:44:27 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD http://www.baidu.com/ HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

> Proxy-Connection: Keep-Alive

- Request completely sent off

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0< HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform

< Connection: keep-alive

< Content-Length: 277

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:44:10 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_12574121221855285883

<

0 277 0 0 0 0 0 0 --:--:-- 0:00:01 --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Accept-Ranges:
  bytes Cache-Control: private, no-cache, no-store, proxy-revalidate,
  no-transform Connection: keep-alive Content-Length: 277 Content-Type:
  text/html Date: Sun, 14 Sep 2025 08:44:10 GMT Etag: "575e1f60-115"
  Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT Pragma: no-cache Server:
  bfe/1.0.8.18 Tr_id: bfe_12574121221855285883

## 错误: 无

开始运行命令... [2025-09-14 16:44:27.246] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

[2025-09-14 16:44:27] [TEST] curl -v -I -L http://www.so.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 17724 执行时间: 2025-09-14
16:44:27 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- Connected to localhost (::1) port 18080
- using HTTP/1.x

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0> HEAD http://www.so.com/ HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: _/_

> Proxy-Connection: Keep-Alive

- Request completely sent off < HTTP/1.1 302 Found

< Connection: keep-alive

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:44:10 GMT

< Location: https://www.so.com/

< Server: openresty

< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

- Ignoring the response-body <

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact
- Clear auth, redirects to port from 80 to 443
- Issue another request to this URL: 'https://www.so.com/'
- Hostname localhost was found in DNS cache
- Trying [::1]:18080...
- CONNECT tunnel: HTTP/1.1 negotiated
- allocate connect buffer
- Establish HTTP proxy tunnel to www.so.com:443
  > CONNECT www.so.com:443 HTTP/1.1

> Host: www.so.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

< HTTP/1.1 200 Connection established

<

- CONNECT phase completed
- CONNECT tunnel established, response 200
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [305 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [93 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [5077 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server did not agree on a protocol. Uses default.
- Server certificate:
- subject: CN=\*.so.com
- start date: Aug 28 00:00:00 2025 GMT
- expire date: Sep 28 23:59:59 2026 GMT
- subjectAltName: host "www.so.com" matched cert's "\*.so.com"
- issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA [Run by the Issuer]
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha384WithRSAEncryption
- Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using
  sha384WithRSAEncryption
- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD / HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Server: openresty

< Date: Sun, 14 Sep 2025 08:44:10 GMT

< Content-Type: text/html; charset=UTF-8

< Connection: keep-alive

< Vary: Accept-Encoding

< Set-Cookie: \_S=2qkbd5o6b2vjg6u9qt33mdidl0; expires=Sun, 14-Sep-2025 08:54:10
GMT; Max-Age=600; path=/

< Expires: Thu, 19 Nov 1981 08:52:00 GMT

< Cache-Control: no-store, no-cache, must-revalidate

< Pragma: no-cache

< php-waf-rep: -

< Set-Cookie: QiHooGUID=F2F13AC29FFE5678CAD8E1EB0704588A.1757839450692;
Max-Age=63072000; Domain=so.com; Path=/

<

0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #1 to host localhost left intact HTTP/1.1 302 Found Connection:
  keep-alive Content-Type: text/html Date: Sun, 14 Sep 2025 08:44:10 GMT
  Location: https://www.so.com/ Server: openresty Set-Cookie: QiHooGUID=;
  Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK Server: openresty Date: Sun, 14 Sep 2025 08:44:10 GMT
Content-Type: text/html; charset=UTF-8 Connection: keep-alive Vary:
Accept-Encoding Set-Cookie: \_S=2qkbd5o6b2vjg6u9qt33mdidl0; expires=Sun,
14-Sep-2025 08:54:10 GMT; Max-Age=600; path=/ Expires: Thu, 19 Nov 1981 08:52:00
GMT Cache-Control: no-store, no-cache, must-revalidate Pragma: no-cache
php-waf-rep: - Set-Cookie:
QiHooGUID=F2F13AC29FFE5678CAD8E1EB0704588A.1757839450692; Max-Age=63072000;
Domain=so.com; Path=/

## 错误: 无

开始运行命令... [2025-09-14 16:44:27.542] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

[2025-09-14 16:44:27] [TEST] curl -v -I https://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程 PID: 35172 执行时间: 2025-09-14
16:44:27 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed

  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0\* Trying [::1]:18080...

- CONNECT tunnel: HTTP/1.1 negotiated
- allocate connect buffer
- Establish HTTP proxy tunnel to www.baidu.com:443
  > CONNECT www.baidu.com:443 HTTP/1.1

> Host: www.baidu.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

< HTTP/1.1 200 Connection established

<

- CONNECT phase completed
- CONNECT tunnel established, response 200
- ALPN: curl offers h2,http/1.1
- TLSv1.3 (OUT), TLS handshake, Client hello (1): } [308 bytes data]
- CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
- CApath: none
- TLSv1.3 (IN), TLS handshake, Server hello (2): { [102 bytes data]
- TLSv1.2 (IN), TLS handshake, Certificate (11): { [4771 bytes data]
- TLSv1.2 (IN), TLS handshake, Server key exchange (12): { [333 bytes data]
- TLSv1.2 (IN), TLS handshake, Server finished (14): { [4 bytes data]
- TLSv1.2 (OUT), TLS handshake, Client key exchange (16): } [70 bytes data]
- TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1): } [1 bytes data]
- TLSv1.2 (OUT), TLS handshake, Finished (20): } [16 bytes data]
- TLSv1.2 (IN), TLS change cipher, Change cipher spec (1): { [1 bytes data]
- TLSv1.2 (IN), TLS handshake, Finished (20): { [16 bytes data]
- SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / [blank] / UNDEF
- ALPN: server accepted http/1.1
- Server certificate:
- subject: C=CN; ST=beijing; L=beijing; O=Beijing Baidu Netcom Science
  Technology Co., Ltd; CN=baidu.com
- start date: Jul 9 07:01:02 2025 GMT
- expire date: Aug 10 07:01:01 2026 GMT
- subjectAltName: host "www.baidu.com" matched cert's "\*.baidu.com"
- issuer: C=BE; O=GlobalSign nv-sa; CN=GlobalSign RSA OV SSL CA 2018
- SSL certificate verify ok.
- Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Certificate level 2: Public key type ? (2048/112 Bits/secBits), signed using
  sha256WithRSAEncryption
- Connected to localhost (::1) port 18080
- using HTTP/1.x
  > HEAD / HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

- Request completely sent off < HTTP/1.1 200 OK

< Accept-Ranges: bytes

< Cache-Control: no-cache

< Connection: keep-alive

< Content-Length: 227

< Content-Security-Policy: frame-ancestors 'self' https://chat.baidu.com
http://mirror-chat.baidu.com https://fj-chat.baidu.com
https://hba-chat.baidu.com https://hbe-chat.baidu.com
https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com
https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com
https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com
https://platform-openai.now.baidu.com;

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:44:11 GMT

< Pragma: no-cache

< Server: BWS/1.1

< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300

< Set-Cookie: PSTM=1757839451; expires=Thu, 31-Dec-37 23:55:55 GMT;
max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: BAIDUID=54274BF524C73EFD7DF41045FB6DFBCD:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000

< Set-Cookie: BAIDUID_BFESS=54274BF524C73EFD7DF41045FB6DFBCD:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None

< Traceid: 175783945105964226669383213264865584105

< X-Ua-Compatible: IE=Edge,chrome=1

< X-Xss-Protection: 1;mode=block

<

0 227 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 Connection
  established

HTTP/1.1 200 OK Accept-Ranges: bytes Cache-Control: no-cache Connection:
keep-alive Content-Length: 227 Content-Security-Policy: frame-ancestors 'self'
https://chat.baidu.com http://mirror-chat.baidu.com https://fj-chat.baidu.com
https://hba-chat.baidu.com https://hbe-chat.baidu.com
https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com
https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com
https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com
https://platform-openai.now.baidu.com; Content-Type: text/html Date: Sun, 14 Sep
2025 08:44:11 GMT Pragma: no-cache Server: BWS/1.1 Set-Cookie: BD_NOT_HTTPS=1;
path=/; Max-Age=300 Set-Cookie: PSTM=1757839451; expires=Thu, 31-Dec-37 23:55:55
GMT; max-age=2147483647; path=/; domain=.baidu.com Set-Cookie:
BAIDUID=54274BF524C73EFD7DF41045FB6DFBCD:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000 Set-Cookie:
BAIDUID_BFESS=54274BF524C73EFD7DF41045FB6DFBCD:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000; Secure; SameSite=None Traceid:
175783945105964226669383213264865584105 X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block

## 错误: 无

开始运行命令... [2025-09-14 16:44:30.967] [执行命令] C:\Program
Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go

开始运行命令... [2025-09-14 16:44:30.967] [BUILD] C:\Program Files\Go\bin\go.exe
go build -o main.exe ../cmd/main.go

```
执行结果: 成功
进程PID: 27384
执行时间: 1.5s
---
```

开始运行命令... [2025-09-14 16:44:32.208] [执行命令] ./main.exe --port 18080

开始运行命令... [2025-09-14 16:44:35.338] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:44:35.338] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
http://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 27120
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
* Connected to localhost (::1) port 18080
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

< Date: Sun, 14 Sep 2025 08:44:19 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_10315764190906471132

<


  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:44:19 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10315764190906471132
---
```

开始运行命令... [2025-09-14 16:44:36.423] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:44:36.423] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L
http://www.so.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 25820
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
* Connected to localhost (::1) port 18080
* using HTTP/1.x
> HEAD http://www.so.com/ HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: */*

> Proxy-Connection: Keep-Alive

>

* Request completely sent off
< HTTP/1.1 302 Found

< Connection: keep-alive

< Content-Type: text/html

< Date: Sun, 14 Sep 2025 08:44:19 GMT

< Location: https://www.so.com/

< Server: openresty

< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

* Ignoring the response-body
<


  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:18080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.so.com:443
> CONNECT www.so.com:443 HTTP/1.1

> Host: www.so.com:443

> User-Agent: curl/8.12.1

> Proxy-Connection: Keep-Alive

>

< HTTP/1.1 200 Connection established

<

* CONNECT phase completed
* CONNECT tunnel established, response 200
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
} [305 bytes data]
*  CAfile: D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl-ca-bundle.crt
*  CApath: none
* TLSv1.3 (IN), TLS handshake, Server hello (2):
{ [93 bytes data]
* TLSv1.2 (IN), TLS handshake, Certificate (11):
{ [5077 bytes data]
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
* ALPN: server did not agree on a protocol. Uses default.
* Server certificate:
*  subject: CN=*.so.com
*  start date: Aug 28 00:00:00 2025 GMT
*  expire date: Sep 28 23:59:59 2026 GMT
*  subjectAltName: host "www.so.com" matched cert's "*.so.com"
*  issuer: C=CN; O=WoTrus CA Limited; CN=WoTrus DV Server CA  [Run by the Issuer]
*  SSL certificate verify ok.
*   Certificate level 0: Public key type ? (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type ? (2048/112 Bits/secBits), signed using sha384WithRSAEncryption
*   Certificate level 2: Public key type ? (4096/128 Bits/secBits), signed using sha384WithRSAEncryption
* Connected to localhost (::1) port 18080
* using HTTP/1.x
> HEAD / HTTP/1.1

> Host: www.so.com

> User-Agent: curl/8.12.1

> Accept: */*

>

* Request completely sent off
< HTTP/1.1 200 OK

< Server: openresty

< Date: Sun, 14 Sep 2025 08:44:19 GMT

< Content-Type: text/html; charset=UTF-8

< Connection: keep-alive

< Vary: Accept-Encoding

< Set-Cookie: _S=jte1u7p07vcs55r2n6ov0rlb27; expires=Sun, 14-Sep-2025 08:54:19 GMT; Max-Age=600; path=/

< Expires: Thu, 19 Nov 1981 08:52:00 GMT

< Cache-Control: no-store, no-cache, must-revalidate

< Pragma: no-cache

< php-waf-rep: -

< Set-Cookie: QiHooGUID=28757436392D5484E6F1114BFC9F77F8.1757839459810; Max-Age=63072000; Domain=so.com; Path=/

<


  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:44:19 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 08:44:19 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=jte1u7p07vcs55r2n6ov0rlb27; expires=Sun, 14-Sep-2025 08:54:19 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=28757436392D5484E6F1114BFC9F77F8.1757839459810; Max-Age=63072000; Domain=so.com; Path=/
---
```

开始运行命令...

[2025-09-14 16:44:36.653] [执行命令]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

开始运行命令... [2025-09-14 16:44:36.653] [CURL]
D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I
https://www.baidu.com -x http://localhost:18080

```
执行结果: 成功
进程PID: 30052
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:18080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed

  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:18080...
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
* Connected to localhost (::1) port 18080
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

< Date: Sun, 14 Sep 2025 08:44:19 GMT

< Etag: "575e1f60-115"

< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT

< Pragma: no-cache

< Server: bfe/1.0.8.18

< Tr_id: bfe_11473488289623244232

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
Date: Sun, 14 Sep 2025 08:44:19 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11473488289623244232
---
```
