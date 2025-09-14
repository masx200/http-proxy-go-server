=== 命令执行日志文件 === 创建时间: 2025-09-14 16:55:19

```
开始运行命令...

[2025-09-14 16:55:19.474] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```

```
开始运行命令...

[2025-09-14 16:55:19.474] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
```

```
执行结果: 成功
进程PID: 35824
执行时间: 1.203125s
---
```

```
开始运行命令...

[2025-09-14 16:55:20.010] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```

```
开始运行命令...

[2025-09-14 16:55:20.010] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```

```
执行结果: 成功
进程PID: 2524
执行时间: 1.453125s
---
```

```
开始运行命令...

[2025-09-14 16:55:21.576] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :28081
```

```
开始运行命令...

[2025-09-14 16:55:21.576] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :28081
```

```
执行结果: 成功
进程PID: 7240
---
```

```
开始运行命令...

[2025-09-14 16:55:22.581] [执行命令] ./main.exe -port 18080 -upstream-type websocket -upstream-address ws://localhost:28081
```

```
开始运行命令...

[2025-09-14 16:55:22.582] [HTTP] ./main.exe -port 18080 -upstream-type websocket -upstream-address ws://localhost:28081
```

```
执行结果: 成功
进程PID: 34644
---
```

```
开始运行命令...

[2025-09-14 16:55:25.754] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080
```

```
开始运行命令...

[2025-09-14 16:55:25.754] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080
```

```
执行结果: 成功
进程PID: 23748
执行时间: 31.25ms
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
< Date: Sun, 14 Sep 2025 08:55:08 GMT
< Location: https://www.baidu.com/error.html
< Server: bfe
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 302 Found
Content-Type: text/plain; charset=utf-8
Date: Sun, 14 Sep 2025 08:55:08 GMT
Location: https://www.baidu.com/error.html
Server: bfe
---
```

```
开始运行命令...

[2025-09-14 16:55:25.808] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080
```

```
开始运行命令...

[2025-09-14 16:55:25.808] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080
```

```
执行结果: 成功
进程PID: 35076
执行时间: 46.875ms
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
< Date: Sun, 14 Sep 2025 08:55:09 GMT
< Etag: "575e1f59-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:01 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_10920487150993261998
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
Date: Sun, 14 Sep 2025 08:55:09 GMT
Etag: "575e1f59-115"
Last-Modified: Mon, 13 Jun 2016 02:50:01 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_10920487150993261998
---
```

```
开始运行命令...

[2025-09-14 16:55:30.803] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```

## [2025-09-14 16:55:30] [BUILD] go build -o main.exe ../cmd/main.go 执行结果: 成功 进程PID: 31668 执行时间: 2025-09-14 16:55:32 输出: 错误: 无

```
开始运行命令...

[2025-09-14 16:55:32.236] [执行命令] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
```

## [2025-09-14 16:55:32] [SERVER] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3 执行结果: 成功 进程PID: 8260 执行时间: 2025-09-14 16:55:32 输出: 错误: 无

```
开始运行命令...

[2025-09-14 16:55:34.399] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080
```

[2025-09-14 16:55:34] [TEST] curl -v -I http://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程PID: 26708 执行时间: 2025-09-14
16:55:34 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:18080...
- Connected to localhost (::1) port 18080
- using HTTP/1.x

> HEAD http://www.baidu.com/ HTTP/1.1

> Host: www.baidu.com

> User-Agent: curl/8.12.1

> Accept: _/_

> Proxy-Connection: Keep-Alive

- Request completely sent off < HTTP/1.1 200 OK

< Bdpagetype: 1

< Bdqid: 0xe5948354000e21ca

< Connection: keep-alive

< Content-Length: 654044

< Content-Type: text/html; charset=utf-8

< Date: Sun, 14 Sep 2025 08:55:17 GMT

< Server: BWS/1.1

< Set-Cookie: BIDUPSID=0FCEF93E075B03DF681092096D0FD943; expires=Thu, 31-Dec-37
23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: PSTM=1757840117; expires=Thu, 31-Dec-37 23:55:55 GMT;
max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: BDSVRTM=0; path=/

< Set-Cookie: BD_HOME=1; path=/

< Set-Cookie: BAIDUID=0FCEF93E075B03DF681092096D0FD943:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000

< Set-Cookie: BAIDUID_BFESS=0FCEF93E075B03DF681092096D0FD943:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None

< Tr_id: super_0xe5948354000e21ca

< Traceid: 1757840117062381671416542991728040092106

< Vary: Accept-Encoding

< X-Ua-Compatible: IE=Edge,chrome=1

< X-Xss-Protection: 1;mode=block

<

0 638k 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #0 to host localhost left intact HTTP/1.1 200 OK Bdpagetype: 1
  Bdqid: 0xe5948354000e21ca Connection: keep-alive Content-Length: 654044
  Content-Type: text/html; charset=utf-8 Date: Sun, 14 Sep 2025 08:55:17 GMT
  Server: BWS/1.1 Set-Cookie: BIDUPSID=0FCEF93E075B03DF681092096D0FD943;
  expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/;
  domain=.baidu.com Set-Cookie: PSTM=1757840117; expires=Thu, 31-Dec-37 23:55:55
  GMT; max-age=2147483647; path=/; domain=.baidu.com Set-Cookie: BDSVRTM=0;
  path=/ Set-Cookie: BD_HOME=1; path=/ Set-Cookie:
  BAIDUID=0FCEF93E075B03DF681092096D0FD943:FG=1; Path=/; Domain=baidu.com;
  Max-Age=31536000 Set-Cookie:
  BAIDUID_BFESS=0FCEF93E075B03DF681092096D0FD943:FG=1; Path=/; Domain=baidu.com;
  Max-Age=31536000; Secure; SameSite=None Tr_id: super_0xe5948354000e21ca
  Traceid: 1757840117062381671416542991728040092106 Vary: Accept-Encoding
  X-Ua-Compatible: IE=Edge,chrome=1 X-Xss-Protection: 1;mode=block

## 错误: 无

```
开始运行命令...

[2025-09-14 16:55:34.544] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080
```

[2025-09-14 16:55:34] [TEST] curl -v -I -L http://www.so.com -x
http://localhost:18080 执行结果: 成功 进程PID: 12528 执行时间: 2025-09-14
16:55:34 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:18080...
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

< Date: Sun, 14 Sep 2025 08:55:17 GMT

< Location: https://www.so.com/

< Server: openresty

< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

- Ignoring the response-body < 0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0
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
- subject: CN=*.so.com
- start date: Aug 28 00:00:00 2025 GMT
- expire date: Sep 28 23:59:59 2026 GMT
- subjectAltName: host "www.so.com" matched cert's "*.so.com"
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

< Date: Sun, 14 Sep 2025 08:55:18 GMT

< Content-Type: text/html; charset=UTF-8

< Connection: keep-alive

< Vary: Accept-Encoding

< Set-Cookie: _S=a0d7u7icjav3pfq8qq7mgb3ar4; expires=Sun, 14-Sep-2025 09:05:18
GMT; Max-Age=600; path=/

< Expires: Thu, 19 Nov 1981 08:52:00 GMT

< Cache-Control: no-store, no-cache, must-revalidate

< Pragma: no-cache

< php-waf-rep: -

< Set-Cookie: QiHooGUID=4F05710127FF1D0BFC953133B161A7CD.1757840118024;
Max-Age=63072000; Domain=so.com; Path=/

<

0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0

- Connection #1 to host localhost left intact HTTP/1.1 302 Found Connection:
  keep-alive Content-Type: text/html Date: Sun, 14 Sep 2025 08:55:17 GMT
  Location: https://www.so.com/ Server: openresty Set-Cookie: QiHooGUID=;
  Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK Server: openresty Date: Sun, 14 Sep 2025 08:55:18 GMT
Content-Type: text/html; charset=UTF-8 Connection: keep-alive Vary:
Accept-Encoding Set-Cookie: _S=a0d7u7icjav3pfq8qq7mgb3ar4; expires=Sun,
14-Sep-2025 09:05:18 GMT; Max-Age=600; path=/ Expires: Thu, 19 Nov 1981 08:52:00
GMT Cache-Control: no-store, no-cache, must-revalidate Pragma: no-cache
php-waf-rep: - Set-Cookie:
QiHooGUID=4F05710127FF1D0BFC953133B161A7CD.1757840118024; Max-Age=63072000;
Domain=so.com; Path=/

## 错误: 无

```
开始运行命令...

[2025-09-14 16:55:34.859] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080
```

[2025-09-14 16:55:34] [TEST] curl -v -I https://www.baidu.com -x
http://localhost:18080 执行结果: 成功 进程PID: 32436 执行时间: 2025-09-14
16:55:34 输出: Note: Using embedded CA bundle, for proxies (233263 bytes)

- Host localhost:18080 was resolved.
- IPv6: ::1
- IPv4: 127.0.0.1 % Total % Received % Xferd Average Speed Time Time Time
  Current Dload Upload Total Spent Left Speed 0 0 0 0 0 0 0 0 --:--:-- --:--:--
  --:--:-- 0* Trying [::1]:18080...
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
- subjectAltName: host "www.baidu.com" matched cert's "*.baidu.com"
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

< Date: Sun, 14 Sep 2025 08:55:18 GMT

< Pragma: no-cache

< Server: BWS/1.1

< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300

< Set-Cookie: PSTM=1757840118; expires=Thu, 31-Dec-37 23:55:55 GMT;
max-age=2147483647; path=/; domain=.baidu.com

< Set-Cookie: BAIDUID=74D6201DC2729E2989E6A09B008AEDB0:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000

< Set-Cookie: BAIDUID_BFESS=74D6201DC2729E2989E6A09B008AEDB0:FG=1; Path=/;
Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None

< Traceid: 1757840118070829261815665225417760814983

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
2025 08:55:18 GMT Pragma: no-cache Server: BWS/1.1 Set-Cookie: BD_NOT_HTTPS=1;
path=/; Max-Age=300 Set-Cookie: PSTM=1757840118; expires=Thu, 31-Dec-37 23:55:55
GMT; max-age=2147483647; path=/; domain=.baidu.com Set-Cookie:
BAIDUID=74D6201DC2729E2989E6A09B008AEDB0:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000 Set-Cookie:
BAIDUID_BFESS=74D6201DC2729E2989E6A09B008AEDB0:FG=1; Path=/; Domain=baidu.com;
Max-Age=31536000; Secure; SameSite=None Traceid:
1757840118070829261815665225417760814983 X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block

## 错误: 无

```
开始运行命令...

[2025-09-14 16:55:38.141] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```

```
开始运行命令...

[2025-09-14 16:55:38.141] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
```

```
执行结果: 成功
进程PID: 21128
执行时间: 1.234375s
---
```

```
开始运行命令...

[2025-09-14 16:55:39.481] [执行命令] ./main.exe --port 18080
```

```
开始运行命令...

[2025-09-14 16:55:42.593] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080
```

```
开始运行命令...

[2025-09-14 16:55:42.593] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080
```

```
执行结果: 成功
进程PID: 24944
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
< Date: Sun, 14 Sep 2025 08:55:25 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9719762814691189749
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:55:25 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9719762814691189749
---
```

```
开始运行命令...

[2025-09-14 16:55:42.669] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080
```

```
开始运行命令...

[2025-09-14 16:55:42.670] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080
```

```
执行结果: 成功
进程PID: 19472
执行时间: 31.25ms
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
< Date: Sun, 14 Sep 2025 08:55:25 GMT
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
< Date: Sun, 14 Sep 2025 08:55:26 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=ls85dgmigb8ioa0mrpab3nmgd0; expires=Sun, 14-Sep-2025 09:05:26 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=3D4C059963653A037BB7B14731AB692E.1757840126068; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 08:55:25 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 08:55:26 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=ls85dgmigb8ioa0mrpab3nmgd0; expires=Sun, 14-Sep-2025 09:05:26 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=3D4C059963653A037BB7B14731AB692E.1757840126068; Max-Age=63072000; Domain=so.com; Path=/
---
```

```
开始运行命令...

[2025-09-14 16:55:42.920] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080
```

```
开始运行命令...

[2025-09-14 16:55:42.920] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080
```

```
执行结果: 成功
进程PID: 25724
执行时间: 62.5ms
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
< HTTP/1.1 302 Found
< Location: https://www.baidu.com/error.html
< Server: bfe
< Date: Sun, 14 Sep 2025 08:55:26 GMT
< Content-Type: text/plain; charset=utf-8
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 302 Found
Location: https://www.baidu.com/error.html
Server: bfe
Date: Sun, 14 Sep 2025 08:55:26 GMT
Content-Type: text/plain; charset=utf-8
---
```
