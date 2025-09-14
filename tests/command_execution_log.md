=== 命令执行日志文件 ===
创建时间: 2025-09-14 15:55:27

开始运行命令...
[2025-09-14 15:55:27.927] [执行命令] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd


开始运行命令...
[2025-09-14 15:55:27.927] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd


```
执行结果: 成功
进程PID: 26084
执行时间: 1.09375s
---
```
开始运行命令...
[2025-09-14 15:55:28.307] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go


开始运行命令...
[2025-09-14 15:55:28.307] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go


```
执行结果: 成功
进程PID: 17008
执行时间: 1.28125s
---
```
开始运行命令...
[2025-09-14 15:55:29.527] [执行命令] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :18080


开始运行命令...
[2025-09-14 15:55:29.527] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :18080


```
执行结果: 成功
进程PID: 3728
---
```
开始运行命令...
[2025-09-14 15:55:30.532] [执行命令] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:18080


开始运行命令...
[2025-09-14 15:55:30.532] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:18080


```
执行结果: 成功
进程PID: 18972
---
```
开始运行命令...
[2025-09-14 15:55:34.413] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810


开始运行命令...
[2025-09-14 15:55:34.413] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810


```
执行结果: 成功
进程PID: 23192
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
< Date: Sun, 14 Sep 2025 07:55:17 GMT
< Etag: "575e1f71-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:25 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11404394141359252402
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:55:17 GMT
Etag: "575e1f71-115"
Last-Modified: Mon, 13 Jun 2016 02:50:25 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11404394141359252402
---
```
开始运行命令...
[2025-09-14 15:55:34.771] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810


开始运行命令...
[2025-09-14 15:55:34.771] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810


```
执行结果: 成功
进程PID: 29936
执行时间: 15.625ms
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
< Date: Sun, 14 Sep 2025 07:55:18 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9853218984006409667
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
Date: Sun, 14 Sep 2025 07:55:18 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9853218984006409667
---
```
开始运行命令...
[2025-09-14 15:55:39.277] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go


[2025-09-14 15:55:39] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 27108
执行时间: 2025-09-14 15:55:40
输出: 
错误: 无
---
开始运行命令...
[2025-09-14 15:55:40.528] [执行命令] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3


[2025-09-14 15:55:40] [SERVER] ./main.exe --port 18080 -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 33112
执行时间: 2025-09-14 15:55:40
输出: 
错误: 无
---
开始运行命令...
[2025-09-14 15:55:42.686] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080


[2025-09-14 15:55:42] [TEST] curl -v -I http://www.baidu.com -x http://localhost:18080
执行结果: 成功
进程PID: 19868
执行时间: 2025-09-14 15:55:42
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
< Bdpagetype: 1
< Bdqid: 0xb65773040194dd8c
< Connection: keep-alive
< Content-Length: 655066
< Content-Type: text/html; charset=utf-8
< Date: Sun, 14 Sep 2025 07:55:25 GMT
< Server: BWS/1.1
< Set-Cookie: BIDUPSID=A0B1DE5E463EE32C226ADF8924225CF9; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: PSTM=1757836525; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BDSVRTM=0; path=/
< Set-Cookie: BD_HOME=1; path=/
< Set-Cookie: BAIDUID=A0B1DE5E463EE32C226ADF8924225CF9:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=A0B1DE5E463EE32C226ADF8924225CF9:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Tr_id: super_0xb65773040194dd8c
< Traceid: 1757836525065756775413139096898920308108
< Vary: Accept-Encoding
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0  639k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Bdpagetype: 1
Bdqid: 0xb65773040194dd8c
Connection: keep-alive
Content-Length: 655066
Content-Type: text/html; charset=utf-8
Date: Sun, 14 Sep 2025 07:55:25 GMT
Server: BWS/1.1
Set-Cookie: BIDUPSID=A0B1DE5E463EE32C226ADF8924225CF9; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: PSTM=1757836525; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BDSVRTM=0; path=/
Set-Cookie: BD_HOME=1; path=/
Set-Cookie: BAIDUID=A0B1DE5E463EE32C226ADF8924225CF9:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=A0B1DE5E463EE32C226ADF8924225CF9:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Tr_id: super_0xb65773040194dd8c
Traceid: 1757836525065756775413139096898920308108
Vary: Accept-Encoding
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
开始运行命令...
[2025-09-14 15:55:42.813] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080


[2025-09-14 15:55:42] [TEST] curl -v -I -L http://www.so.com -x http://localhost:18080
执行结果: 成功
进程PID: 1952
执行时间: 2025-09-14 15:55:43
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
< Date: Sun, 14 Sep 2025 07:55:26 GMT
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
< Date: Sun, 14 Sep 2025 07:55:26 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=20kacemt888bqfspe5vttu7sp0; expires=Sun, 14-Sep-2025 08:05:26 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=3F9A7B6B42FCF153EA10D107082C81CC.1757836526234; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:55:26 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:55:26 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=20kacemt888bqfspe5vttu7sp0; expires=Sun, 14-Sep-2025 08:05:26 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=3F9A7B6B42FCF153EA10D107082C81CC.1757836526234; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
开始运行命令...
[2025-09-14 15:55:43.138] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080


[2025-09-14 15:55:43] [TEST] curl -v -I https://www.baidu.com -x http://localhost:18080
执行结果: 成功
进程PID: 9072
执行时间: 2025-09-14 15:55:43
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
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: no-cache
< Connection: keep-alive
< Content-Length: 227
< Content-Security-Policy: frame-ancestors 'self' https://chat.baidu.com http://mirror-chat.baidu.com https://fj-chat.baidu.com https://hba-chat.baidu.com https://hbe-chat.baidu.com https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com https://platform-openai.now.baidu.com;
< Content-Type: text/html
< Date: Sun, 14 Sep 2025 07:55:26 GMT
< Pragma: no-cache
< Server: BWS/1.1
< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
< Set-Cookie: PSTM=1757836526; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BAIDUID=136349B41E0572439638D48B39717146:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=136349B41E0572439638D48B39717146:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Traceid: 1757836526057971098612740736302802609058
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0   227    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: no-cache
Connection: keep-alive
Content-Length: 227
Content-Security-Policy: frame-ancestors 'self' https://chat.baidu.com http://mirror-chat.baidu.com https://fj-chat.baidu.com https://hba-chat.baidu.com https://hbe-chat.baidu.com https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com https://hnb-chat.baidu.com http://debug.baidu-int.com https://sai.baidu.com https://mcpstore.baidu.com https://mcpserver.baidu.com https://www.mcpworld.com https://platform-openai.now.baidu.com;
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:55:26 GMT
Pragma: no-cache
Server: BWS/1.1
Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
Set-Cookie: PSTM=1757836526; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BAIDUID=136349B41E0572439638D48B39717146:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=136349B41E0572439638D48B39717146:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Traceid: 1757836526057971098612740736302802609058
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
开始运行命令...
[2025-09-14 15:55:46.236] [执行命令] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go


开始运行命令...
[2025-09-14 15:55:46.236] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go


```
执行结果: 成功
进程PID: 20580
执行时间: 1.171875s
---
```
开始运行命令...
[2025-09-14 15:55:47.460] [执行命令] ./main.exe --port 18080


开始运行命令...
[2025-09-14 15:55:50.566] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080


开始运行命令...
[2025-09-14 15:55:50.566] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:18080


```
执行结果: 成功
进程PID: 30712
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
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Sun, 14 Sep 2025 07:55:33 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9105971924304972876
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:55:33 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9105971924304972876
---
```
开始运行命令...
[2025-09-14 15:55:50.654] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080


开始运行命令...
[2025-09-14 15:55:50.654] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:18080


```
执行结果: 成功
进程PID: 12104
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
< Date: Sun, 14 Sep 2025 07:55:33 GMT
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
< Date: Sun, 14 Sep 2025 07:55:33 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=7f5ktlhncmdgcuh9peidjuqtp0; expires=Sun, 14-Sep-2025 08:05:33 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=9F598D1F6716E9AC145091FBE66D5BD0.1757836533965; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 07:55:33 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 07:55:33 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=7f5ktlhncmdgcuh9peidjuqtp0; expires=Sun, 14-Sep-2025 08:05:33 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=9F598D1F6716E9AC145091FBE66D5BD0.1757836533965; Max-Age=63072000; Domain=so.com; Path=/
---
```
开始运行命令...
[2025-09-14 15:55:50.871] [执行命令] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080


开始运行命令...
[2025-09-14 15:55:50.871] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:18080


```
执行结果: 成功
进程PID: 33172
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
< Date: Sun, 14 Sep 2025 07:55:34 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8481616153485965955
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
Date: Sun, 14 Sep 2025 07:55:34 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8481616153485965955
---
```
