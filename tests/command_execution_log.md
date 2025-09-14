=== 命令执行日志文件 ===
创建时间: 2025-09-14 01:11:09

[2025-09-14 01:11:09.910] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 34080
执行时间: 1.765625s
---
[2025-09-14 01:11:14.808] [CURL] C:\Program Files\Git\mingw64\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 24868
执行时间: 15.625ms
输出: * Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD http://www.baidu.com/ HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.14.1
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
< Date: Sat, 13 Sep 2025 17:10:58 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12015824289897902877
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:10:58 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12015824289897902877


* Connection #0 to host localhost left intact
---
[2025-09-14 01:11:14.863] [CURL] C:\Program Files\Git\mingw64\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 18128
执行时间: 62.5ms
输出: * Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD http://www.so.com/ HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.14.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
* Request completely sent off
< HTTP/1.1 302 Found
< Connection: keep-alive
< Content-Type: text/html
< Date: Sat, 13 Sep 2025 17:10:58 GMT
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
*   Trying [::1]:8080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.so.com:443
> CONNECT www.so.com:443 HTTP/1.1
> Host: www.so.com:443
> User-Agent: curl/8.14.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* schannel: disabled automatic use of client certificate
* ALPN: curl offers http/1.1
* ALPN: server did not agree on a protocol. Uses default.
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.14.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:10:58 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=84vqk56dhqeubch5064cn656h2; expires=Sat, 13-Sep-2025 17:20:58 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=97430541FB75C91233EAB13DDBF3C603.1757783458830; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:10:58 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:10:58 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=84vqk56dhqeubch5064cn656h2; expires=Sat, 13-Sep-2025 17:20:58 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=97430541FB75C91233EAB13DDBF3C603.1757783458830; Max-Age=63072000; Domain=so.com; Path=/


* Connection #1 to host localhost left intact
---
[2025-09-14 01:11:15.010] [CURL] C:\Program Files\Git\mingw64\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 34780
执行时间: 15.625ms
输出: * Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.14.1
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 200 Connection established
< 
* CONNECT phase completed
* CONNECT tunnel established, response 200
* schannel: disabled automatic use of client certificate
* ALPN: curl offers http/1.1
* ALPN: server accepted http/1.1
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.baidu.com
> User-Agent: curl/8.14.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
< Connection: keep-alive
< Content-Length: 277
< Content-Type: text/html
< Date: Sat, 13 Sep 2025 17:10:58 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9543423590586223627
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:10:58 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9543423590586223627


* Connection #0 to host localhost left intact
---
[2025-09-14 01:12:51.894] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 38040
执行时间: 1.671875s
---
[2025-09-14 01:12:56.403] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 38244
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11904522441405846353
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:12:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11904522441405846353
---
[2025-09-14 01:12:56.451] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 39256
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:40 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:12:40 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=u7rkth21uvv3duhkmkjm6qmbd5; expires=Sat, 13-Sep-2025 17:22:40 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=6170AB1455AED11D15FAE188FC94C44D.1757783560460; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:12:40 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:12:40 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=u7rkth21uvv3duhkmkjm6qmbd5; expires=Sat, 13-Sep-2025 17:22:40 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=6170AB1455AED11D15FAE188FC94C44D.1757783560460; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:12:56.653] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 8264
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11432493217373461392
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
Date: Sat, 13 Sep 2025 17:12:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11432493217373461392
---
[2025-09-14 01:13:02.077] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 34428
执行时间: 1.546875s
---
[2025-09-14 01:13:06.647] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 5324
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:50 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11746253304547802643
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:12:50 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11746253304547802643
---
[2025-09-14 01:13:06.703] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 33732
执行时间: 62.5ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:50 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:12:50 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=pjk7cijiklh1b6vcstsl4fmpo5; expires=Sat, 13-Sep-2025 17:22:50 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=C3BB93653A7BC9CAA0D0318D4227F966.1757783570656; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:12:50 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:12:50 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=pjk7cijiklh1b6vcstsl4fmpo5; expires=Sat, 13-Sep-2025 17:22:50 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=C3BB93653A7BC9CAA0D0318D4227F966.1757783570656; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:13:06.859] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 40252
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:12:50 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8561368074000356417
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
Date: Sat, 13 Sep 2025 17:12:50 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8561368074000356417
---
[2025-09-14 01:13:12.348] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 失败
进程PID: 31344
执行时间: 937.5ms
错误: exit status 1
---
[2025-09-14 01:13:47.416] [CLEANUP] C:\Windows\System32\taskkill.exe taskkill /F /IM go.exe
执行结果: 成功
进程PID: 19792
执行时间: 15.625ms
---
[2025-09-14 01:17:15.956] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 41240
执行时间: 1.5625s
---
[2025-09-14 01:17:20.480] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 37200
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:04 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9737617530607727559
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:17:04 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9737617530607727559
---
[2025-09-14 01:17:20.538] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 22412
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:04 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:17:04 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=npdhn2kul70hnoqd82mucnc0a2; expires=Sat, 13-Sep-2025 17:27:04 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=E623F560E48838B819A01113EEE661D5.1757783824533; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:17:04 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:17:04 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=npdhn2kul70hnoqd82mucnc0a2; expires=Sat, 13-Sep-2025 17:27:04 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=E623F560E48838B819A01113EEE661D5.1757783824533; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:17:20.722] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 38504
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:04 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8834924252019451636
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
Date: Sat, 13 Sep 2025 17:17:04 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8834924252019451636
---
[2025-09-14 01:17:26.091] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 13820
执行时间: 1.59375s
---
[2025-09-14 01:17:30.634] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 37064
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12100033192509180451
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:17:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12100033192509180451
---
[2025-09-14 01:17:30.712] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 12068
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:14 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:17:14 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=5nv5p084panelt3s4aqjj57194; expires=Sat, 13-Sep-2025 17:27:14 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=BBBA60E0B56B3FD83C80F1F05FC51B05.1757783834690; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:17:14 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:17:14 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=5nv5p084panelt3s4aqjj57194; expires=Sat, 13-Sep-2025 17:27:14 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=BBBA60E0B56B3FD83C80F1F05FC51B05.1757783834690; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:17:30.867] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 2512
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:17:14 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12102339856219504458
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
Date: Sat, 13 Sep 2025 17:17:14 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12102339856219504458
---
[2025-09-14 01:17:36.296] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 成功
进程PID: 40644
执行时间: 687.5ms
---
[2025-09-14 01:17:37.022] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 30476
执行时间: 1.09375s
---
[2025-09-14 01:17:38.370] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
执行结果: 成功
进程PID: 6188
---
[2025-09-14 01:17:39.379] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800
执行结果: 成功
进程PID: 40152
---
[2025-09-14 01:17:42.524] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 9220
执行时间: 0s
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
< Date: Sat, 13 Sep 2025 17:17:26 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11772985825386956629
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:17:26 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11772985825386956629
---
[2025-09-14 01:17:42.595] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 39488
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
< Date: Sat, 13 Sep 2025 17:17:26 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12156699822735828105
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
Date: Sat, 13 Sep 2025 17:17:26 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12156699822735828105
---
[2025-09-14 01:34:23.195] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 40160
执行时间: 1.453125s
---
[2025-09-14 01:34:27.575] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 40420
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:11 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11934198097074938811
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:34:11 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11934198097074938811
---
[2025-09-14 01:34:27.611] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 30312
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:11 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:34:11 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=3rvojbed0s5kb1p8tdeotrt9p7; expires=Sat, 13-Sep-2025 17:44:11 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=636511C24762E526A6550101C017DA31.1757784851592; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:34:11 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:34:11 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=3rvojbed0s5kb1p8tdeotrt9p7; expires=Sat, 13-Sep-2025 17:44:11 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=636511C24762E526A6550101C017DA31.1757784851592; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:34:27.765] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 25580
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:11 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8979890500803855582
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
Date: Sat, 13 Sep 2025 17:34:11 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8979890500803855582
---
[2025-09-14 01:34:33.126] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 15072
执行时间: 1.296875s
---
[2025-09-14 01:34:37.480] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 35136
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:21 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11876608259157673922
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:34:21 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11876608259157673922
---
[2025-09-14 01:34:37.528] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 39084
执行时间: 0s
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:21 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:34:21 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=csl6aomcbm7ikjpv4u54o4u7q3; expires=Sat, 13-Sep-2025 17:44:21 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=B688562E1C3FB44F237401EBAFBDBDB6.1757784861534; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:34:21 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:34:21 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=csl6aomcbm7ikjpv4u54o4u7q3; expires=Sat, 13-Sep-2025 17:44:21 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=B688562E1C3FB44F237401EBAFBDBDB6.1757784861534; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:34:37.683] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 14936
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:34:21 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9244652294484637390
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
Date: Sat, 13 Sep 2025 17:34:21 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9244652294484637390
---
[2025-09-14 01:34:43.104] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 成功
进程PID: 33460
执行时间: 1.1875s
---
[2025-09-14 01:34:43.524] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 39640
执行时间: 1.53125s
---
[2025-09-14 01:34:44.880] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
执行结果: 成功
进程PID: 34280
---
[2025-09-14 01:34:45.885] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800
执行结果: 成功
进程PID: 33052
---
[2025-09-14 01:34:49.008] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 39748
执行时间: 0s
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
< Date: Sat, 13 Sep 2025 17:34:32 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9313797059694367087
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:34:32 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9313797059694367087
---
[2025-09-14 01:34:49.063] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 35768
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
< Date: Sat, 13 Sep 2025 17:34:33 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8709760409625575372
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
Date: Sat, 13 Sep 2025 17:34:33 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8709760409625575372
---
[2025-09-14 01:48:05.035] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 35688
执行时间: 1.265625s
---
[2025-09-14 01:48:09.567] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 41060
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:47:53 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12505682812067172896
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:47:53 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12505682812067172896
---
[2025-09-14 01:48:09.614] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 43000
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:47:53 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 17:47:53 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=vhitpfdcq1b6ttvlu4n2bbea17; expires=Sat, 13-Sep-2025 17:57:53 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=39036690505BF072E6C471E9B22F21C2.1757785673632; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 17:47:53 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 17:47:53 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=vhitpfdcq1b6ttvlu4n2bbea17; expires=Sat, 13-Sep-2025 17:57:53 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=39036690505BF072E6C471E9B22F21C2.1757785673632; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 01:48:09.765] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 42148
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 17:47:53 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12274553630881881398
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
Date: Sat, 13 Sep 2025 17:47:53 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12274553630881881398
---
[2025-09-14 02:20:36] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 39368
执行时间: 2025-09-14 02:20:36
输出: 
错误: 无
---
[2025-09-14 02:20:36] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 42248
执行时间: 2025-09-14 02:20:36
输出: 
错误: 无
---
[2025-09-14 02:20:39] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 42960
执行时间: 2025-09-14 02:20:39
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Bdqid: 0xfdc4897900c726ea
< Connection: keep-alive
< Content-Length: 658699
< Content-Type: text/html; charset=utf-8
< Date: Sat, 13 Sep 2025 18:20:23 GMT
< Server: BWS/1.1
< Set-Cookie: BIDUPSID=7F5B3D3550A28BF2C91917C3F8F453D6; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: PSTM=1757787623; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BDSVRTM=1; path=/
< Set-Cookie: BD_HOME=1; path=/
< Set-Cookie: BAIDUID=7F5B3D3550A28BF2C91917C3F8F453D6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=7F5B3D3550A28BF2C91917C3F8F453D6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Tr_id: super_0xfdc4897900c726ea
< Traceid: 1757787623062381671418285891539828156138
< Vary: Accept-Encoding
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0  643k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Bdpagetype: 1
Bdqid: 0xfdc4897900c726ea
Connection: keep-alive
Content-Length: 658699
Content-Type: text/html; charset=utf-8
Date: Sat, 13 Sep 2025 18:20:23 GMT
Server: BWS/1.1
Set-Cookie: BIDUPSID=7F5B3D3550A28BF2C91917C3F8F453D6; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: PSTM=1757787623; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BDSVRTM=1; path=/
Set-Cookie: BD_HOME=1; path=/
Set-Cookie: BAIDUID=7F5B3D3550A28BF2C91917C3F8F453D6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=7F5B3D3550A28BF2C91917C3F8F453D6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Tr_id: super_0xfdc4897900c726ea
Traceid: 1757787623062381671418285891539828156138
Vary: Accept-Encoding
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 02:20:39] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 43132
执行时间: 2025-09-14 02:20:39
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:20:23 GMT
< Location: https://www.so.com/
< Server: openresty
< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
* Ignoring the response-body
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 18:20:23 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=jb20oj6b4vgae9e4523a2bkje6; expires=Sat, 13-Sep-2025 18:30:23 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=3F495C7BCCE9885DD10B41D1861D9E0C.1757787623308; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:20:23 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 18:20:23 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=jb20oj6b4vgae9e4523a2bkje6; expires=Sat, 13-Sep-2025 18:30:23 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=3F495C7BCCE9885DD10B41D1861D9E0C.1757787623308; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 02:20:39] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 43456
执行时间: 2025-09-14 02:20:39
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
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
< Date: Sat, 13 Sep 2025 18:20:23 GMT
< Pragma: no-cache
< Server: BWS/1.1
< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
< Set-Cookie: PSTM=1757787623; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BAIDUID=0205C3F326F58965F60BF9FA2041014F:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=0205C3F326F58965F60BF9FA2041014F:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Traceid: 1757787623060710503413220537192592970188
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
Date: Sat, 13 Sep 2025 18:20:23 GMT
Pragma: no-cache
Server: BWS/1.1
Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
Set-Cookie: PSTM=1757787623; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BAIDUID=0205C3F326F58965F60BF9FA2041014F:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=0205C3F326F58965F60BF9FA2041014F:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Traceid: 1757787623060710503413220537192592970188
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 02:22:04] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 42440
执行时间: 2025-09-14 02:22:05
输出: 
错误: 无
---
[2025-09-14 02:22:05] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 20200
执行时间: 2025-09-14 02:22:05
输出: 
错误: 无
---
[2025-09-14 02:22:07] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 33648
执行时间: 2025-09-14 02:22:07
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Bdqid: 0xa514c97b0055a18a
< Connection: keep-alive
< Content-Length: 657929
< Content-Type: text/html; charset=utf-8
< Date: Sat, 13 Sep 2025 18:21:51 GMT
< Server: BWS/1.1
< Set-Cookie: BIDUPSID=439A0349AAEB294C1AC7E56500FB86E6; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: PSTM=1757787711; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BDSVRTM=0; path=/
< Set-Cookie: BD_HOME=1; path=/
< Set-Cookie: BAIDUID=439A0349AAEB294C1AC7E56500FB86E6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=439A0349AAEB294C1AC7E56500FB86E6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Tr_id: super_0xa514c97b0055a18a
< Traceid: 1757787711064046285811895354045916094858
< Vary: Accept-Encoding
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0  642k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0  642k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Bdpagetype: 1
Bdqid: 0xa514c97b0055a18a
Connection: keep-alive
Content-Length: 657929
Content-Type: text/html; charset=utf-8
Date: Sat, 13 Sep 2025 18:21:51 GMT
Server: BWS/1.1
Set-Cookie: BIDUPSID=439A0349AAEB294C1AC7E56500FB86E6; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: PSTM=1757787711; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BDSVRTM=0; path=/
Set-Cookie: BD_HOME=1; path=/
Set-Cookie: BAIDUID=439A0349AAEB294C1AC7E56500FB86E6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=439A0349AAEB294C1AC7E56500FB86E6:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Tr_id: super_0xa514c97b0055a18a
Traceid: 1757787711064046285811895354045916094858
Vary: Accept-Encoding
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 02:22:07] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 16280
执行时间: 2025-09-14 02:22:07
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:21:51 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 18:21:51 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=ra89nul0t04qr44bmj21nds1c5; expires=Sat, 13-Sep-2025 18:31:51 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=2F73D1BEFE0C7BA86D52A12A9E27C408.1757787711518; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:21:51 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 18:21:51 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=ra89nul0t04qr44bmj21nds1c5; expires=Sat, 13-Sep-2025 18:31:51 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=2F73D1BEFE0C7BA86D52A12A9E27C408.1757787711518; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 02:22:07] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 33760
执行时间: 2025-09-14 02:22:08
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
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
< Date: Sat, 13 Sep 2025 18:21:51 GMT
< Pragma: no-cache
< Server: BWS/1.1
< Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
< Set-Cookie: PSTM=1757787711; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BAIDUID=064F578BA36D4F9DEFA8A01C88FD9272:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=064F578BA36D4F9DEFA8A01C88FD9272:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Traceid: 1757787711069144986611300951943915573172
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
Date: Sat, 13 Sep 2025 18:21:51 GMT
Pragma: no-cache
Server: BWS/1.1
Set-Cookie: BD_NOT_HTTPS=1; path=/; Max-Age=300
Set-Cookie: PSTM=1757787711; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BAIDUID=064F578BA36D4F9DEFA8A01C88FD9272:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=064F578BA36D4F9DEFA8A01C88FD9272:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Traceid: 1757787711069144986611300951943915573172
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 02:22:10.193] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 8260
执行时间: 1.96875s
---
[2025-09-14 02:22:14.572] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 42620
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:21:58 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12033132020282947310
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:21:58 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12033132020282947310
---
[2025-09-14 02:22:14.624] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 43208
执行时间: 62.5ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:21:58 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 18:21:58 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=l8snd1bk4rs1o8fscd0tiq5737; expires=Sat, 13-Sep-2025 18:31:58 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=861E163FBDC25862610EE1F9E9FB11D9.1757787718678; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:21:58 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 18:21:58 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=l8snd1bk4rs1o8fscd0tiq5737; expires=Sat, 13-Sep-2025 18:31:58 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=861E163FBDC25862610EE1F9E9FB11D9.1757787718678; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 02:22:14.790] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 43748
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:21:58 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12245756777065127796
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
Date: Sat, 13 Sep 2025 18:21:58 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12245756777065127796
---
[2025-09-14 02:22:20.289] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 成功
进程PID: 36504
执行时间: 1.0625s
---
[2025-09-14 02:22:20.713] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 42232
执行时间: 1.359375s
---
[2025-09-14 02:22:22.136] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
执行结果: 成功
进程PID: 40796
---
[2025-09-14 02:22:23.141] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800
执行结果: 成功
进程PID: 3332
---
[2025-09-14 02:22:26.269] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 40640
执行时间: 31.25ms
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
< Date: Sat, 13 Sep 2025 18:22:10 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:22:10 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
---
[2025-09-14 02:22:26.350] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 40592
执行时间: 46.875ms
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
< HTTP/1.1 302 Found
< Location: https://www.baidu.com/error.html
< Server: bfe
< Date: Sat, 13 Sep 2025 18:22:10 GMT
< Content-Type: text/plain; charset=utf-8
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 302 Found
Location: https://www.baidu.com/error.html
Server: bfe
Date: Sat, 13 Sep 2025 18:22:10 GMT
Content-Type: text/plain; charset=utf-8
---
[2025-09-14 02:30:46] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 12612
执行时间: 2025-09-14 02:30:47
输出: 
错误: 无
---
[2025-09-14 02:30:47] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 9032
执行时间: 2025-09-14 02:30:47
输出: 
错误: 无
---
[2025-09-14 02:30:50] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 40876
执行时间: 2025-09-14 02:30:50
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Bdqid: 0xe0133e8a005f5f8e
< Connection: keep-alive
< Content-Length: 658582
< Content-Type: text/html; charset=utf-8
< Date: Sat, 13 Sep 2025 18:30:34 GMT
< Server: BWS/1.1
< Set-Cookie: BIDUPSID=905356C84770B07D2EF17F732DB46C2C; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: PSTM=1757788234; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BDSVRTM=1; path=/
< Set-Cookie: BD_HOME=1; path=/
< Set-Cookie: BAIDUID=905356C84770B07D2EF17F732DB46C2C:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=905356C84770B07D2EF17F732DB46C2C:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Tr_id: super_0xe0133e8a005f5f8e
< Traceid: 1757788234059642266616146317851486019470
< Vary: Accept-Encoding
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0  643k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Bdpagetype: 1
Bdqid: 0xe0133e8a005f5f8e
Connection: keep-alive
Content-Length: 658582
Content-Type: text/html; charset=utf-8
Date: Sat, 13 Sep 2025 18:30:34 GMT
Server: BWS/1.1
Set-Cookie: BIDUPSID=905356C84770B07D2EF17F732DB46C2C; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: PSTM=1757788234; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BDSVRTM=1; path=/
Set-Cookie: BD_HOME=1; path=/
Set-Cookie: BAIDUID=905356C84770B07D2EF17F732DB46C2C:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=905356C84770B07D2EF17F732DB46C2C:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Tr_id: super_0xe0133e8a005f5f8e
Traceid: 1757788234059642266616146317851486019470
Vary: Accept-Encoding
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 02:30:50] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 38144
执行时间: 2025-09-14 02:30:50
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:30:34 GMT
< Location: https://www.so.com/
< Server: openresty
< Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/
* Ignoring the response-body
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
* Clear auth, redirects to port from 80 to 443
* Issue another request to this URL: 'https://www.so.com/'
* Hostname localhost was found in DNS cache
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sat, 13 Sep 2025 18:30:34 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=3g2fdj0rmsqtuuapn9gjd337i0; expires=Sat, 13-Sep-2025 18:40:34 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=98CC9D9EEEA0D356B7C9D158944C31D3.1757788234348; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sat, 13 Sep 2025 18:30:34 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sat, 13 Sep 2025 18:30:34 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=3g2fdj0rmsqtuuapn9gjd337i0; expires=Sat, 13-Sep-2025 18:40:34 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=98CC9D9EEEA0D356B7C9D158944C31D3.1757788234348; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 02:30:50] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 35796
执行时间: 2025-09-14 02:30:50
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sat, 13 Sep 2025 18:30:34 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11591513613023076032
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
Date: Sat, 13 Sep 2025 18:30:34 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11591513613023076032


错误: 无
---
[2025-09-14 13:18:32] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 33068
执行时间: 2025-09-14 13:18:33
输出: 
错误: 无
---
[2025-09-14 13:18:33] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 8124
执行时间: 2025-09-14 13:18:33
输出: 
错误: 无
---
[2025-09-14 13:18:36] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 29784
执行时间: 2025-09-14 13:18:36
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:18:19 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11348829032737148809
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:18:19 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11348829032737148809


错误: 无
---
[2025-09-14 13:18:36] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 10900
执行时间: 2025-09-14 13:18:36
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:18:19 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 05:18:19 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=tbs1rhih159tcbrrvltlhp0en6; expires=Sun, 14-Sep-2025 05:28:19 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=0BF31A78875438FA5861D1CE3497F24F.1757827099493; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:18:19 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 05:18:19 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=tbs1rhih159tcbrrvltlhp0en6; expires=Sun, 14-Sep-2025 05:28:19 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=0BF31A78875438FA5861D1CE3497F24F.1757827099493; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 13:18:36] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 9044
执行时间: 2025-09-14 13:18:37
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:18:20 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_12679302909386435646
< 
  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0  0   277    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:18:20 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_12679302909386435646


错误: 无
---
[2025-09-14 13:18:39.652] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 13920
执行时间: 1.34375s
---
[2025-09-14 13:19:35] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 24104
执行时间: 2025-09-14 13:19:36
输出: 
错误: 无
---
[2025-09-14 13:19:36] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 32724
执行时间: 2025-09-14 13:19:36
输出: 
错误: 无
---
[2025-09-14 13:19:38] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 19448
执行时间: 2025-09-14 13:19:38
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Bdqid: 0x8d2d1e370159b432
< Connection: keep-alive
< Content-Length: 656716
< Content-Type: text/html; charset=utf-8
< Date: Sun, 14 Sep 2025 05:19:21 GMT
< Server: BWS/1.1
< Set-Cookie: BIDUPSID=F78DA84A737B7CDEAD8E4BED9000B1AC; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: PSTM=1757827161; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
< Set-Cookie: BDSVRTM=1; path=/
< Set-Cookie: BD_HOME=1; path=/
< Set-Cookie: BAIDUID=F78DA84A737B7CDEAD8E4BED9000B1AC:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
< Set-Cookie: BAIDUID_BFESS=F78DA84A737B7CDEAD8E4BED9000B1AC:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
< Tr_id: super_0x8d2d1e370159b432
< Traceid: 1757827161055716660210172820354894509106
< Vary: Accept-Encoding
< X-Ua-Compatible: IE=Edge,chrome=1
< X-Xss-Protection: 1;mode=block
< 
  0  641k    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Bdpagetype: 1
Bdqid: 0x8d2d1e370159b432
Connection: keep-alive
Content-Length: 656716
Content-Type: text/html; charset=utf-8
Date: Sun, 14 Sep 2025 05:19:21 GMT
Server: BWS/1.1
Set-Cookie: BIDUPSID=F78DA84A737B7CDEAD8E4BED9000B1AC; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: PSTM=1757827161; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
Set-Cookie: BDSVRTM=1; path=/
Set-Cookie: BD_HOME=1; path=/
Set-Cookie: BAIDUID=F78DA84A737B7CDEAD8E4BED9000B1AC:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000
Set-Cookie: BAIDUID_BFESS=F78DA84A737B7CDEAD8E4BED9000B1AC:FG=1; Path=/; Domain=baidu.com; Max-Age=31536000; Secure; SameSite=None
Tr_id: super_0x8d2d1e370159b432
Traceid: 1757827161055716660210172820354894509106
Vary: Accept-Encoding
X-Ua-Compatible: IE=Edge,chrome=1
X-Xss-Protection: 1;mode=block


错误: 无
---
[2025-09-14 13:19:38] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 20368
执行时间: 2025-09-14 13:19:39
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:19:21 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 05:19:22 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=77jvasp20quq1ru2o1qmlvsnm5; expires=Sun, 14-Sep-2025 05:29:22 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=2DBD3B5F8ABAF1FDDA8FC1EA9F991070.1757827162079; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:19:21 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 05:19:22 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=77jvasp20quq1ru2o1qmlvsnm5; expires=Sun, 14-Sep-2025 05:29:22 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=2DBD3B5F8ABAF1FDDA8FC1EA9F991070.1757827162079; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 13:19:39] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 2808
执行时间: 2025-09-14 13:19:39
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0< HTTP/1.1 200 Connection established
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
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:19:22 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_9316876147546200600
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
Date: Sun, 14 Sep 2025 05:19:22 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_9316876147546200600


错误: 无
---
[2025-09-14 13:19:41.347] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 4012
执行时间: 1.4375s
---
[2025-09-14 13:19:45.722] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 2024
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:19:28 GMT
< Etag: "575e1f6f-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:23 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11368953956720787410
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:19:28 GMT
Etag: "575e1f6f-115"
Last-Modified: Mon, 13 Jun 2016 02:50:23 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11368953956720787410
---
[2025-09-14 13:19:45.826] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 26784
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:19:28 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 05:19:28 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=lrv2akrbvi4sdm65luojsrqpr4; expires=Sun, 14-Sep-2025 05:29:28 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=D86BCA2101ECCD196D24E1B2FFDA0B49.1757827168949; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:19:28 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 05:19:28 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=lrv2akrbvi4sdm65luojsrqpr4; expires=Sun, 14-Sep-2025 05:29:28 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=D86BCA2101ECCD196D24E1B2FFDA0B49.1757827168949; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 13:19:46.067] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 31532
执行时间: 46.875ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:19:29 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11363261849440015699
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
Date: Sun, 14 Sep 2025 05:19:29 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11363261849440015699
---
[2025-09-14 13:19:51.259] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 成功
进程PID: 33536
执行时间: 1.140625s
---
[2025-09-14 13:19:51.634] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 8472
执行时间: 1.453125s
---
[2025-09-14 13:19:52.808] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
执行结果: 成功
进程PID: 33472
---
[2025-09-14 13:19:53.814] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800
执行结果: 成功
进程PID: 33356
---
[2025-09-14 13:19:56.948] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 20708
执行时间: 0s
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
< Date: Sun, 14 Sep 2025 05:19:39 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11563180422438728794
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:19:39 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11563180422438728794
---
[2025-09-14 13:19:57.028] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 2804
执行时间: 46.875ms
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
< Date: Sun, 14 Sep 2025 05:19:40 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8614859276286411894
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
Date: Sun, 14 Sep 2025 05:19:40 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8614859276286411894
---
[2025-09-14 13:23:18] [BUILD] go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 5244
执行时间: 2025-09-14 13:23:19
输出: 
错误: 无
---
[2025-09-14 13:23:19] [SERVER] ./main.exe -dohurl https://dns.alidns.com/dns-query -dohip 223.5.5.5 -dohip 223.6.6.6 -dohurl https://dns.alidns.com/dns-query -dohalpn h2 -dohalpn h3
执行结果: 成功
进程PID: 27016
执行时间: 2025-09-14 13:23:19
输出: 
错误: 无
---
[2025-09-14 13:23:21] [TEST] curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 28140
执行时间: 2025-09-14 13:23:21
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:04 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11297032285556119187
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:23:04 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11297032285556119187


错误: 无
---
[2025-09-14 13:23:21] [TEST] curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 3360
执行时间: 2025-09-14 13:23:22
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:04 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 05:23:05 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=iakbm859li3nbsnt52lijnn144; expires=Sun, 14-Sep-2025 05:33:05 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=4027D14438BE46C0EC1501C316358E4C.1757827385089; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:23:04 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 05:23:05 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=iakbm859li3nbsnt52lijnn144; expires=Sun, 14-Sep-2025 05:33:05 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=4027D14438BE46C0EC1501C316358E4C.1757827385089; Max-Age=63072000; Domain=so.com; Path=/


错误: 无
---
[2025-09-14 13:23:22] [TEST] curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 14252
执行时间: 2025-09-14 13:23:22
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* CONNECT tunnel: HTTP/1.1 negotiated
* allocate connect buffer
* Establish HTTP proxy tunnel to www.baidu.com:443
> CONNECT www.baidu.com:443 HTTP/1.1
> Host: www.baidu.com:443
> User-Agent: curl/8.12.1
> Proxy-Connection: Keep-Alive
> 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0< HTTP/1.1 200 Connection established
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
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:05 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11860450412020180202
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
Date: Sun, 14 Sep 2025 05:23:05 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11860450412020180202


错误: 无
---
[2025-09-14 13:23:24.267] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 3456
执行时间: 1.484375s
---
[2025-09-14 13:23:28.675] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 33236
执行时间: 15.625ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:11 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11857946248499194976
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:23:11 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11857946248499194976
---
[2025-09-14 13:23:28.748] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I -L http://www.so.com -x http://localhost:8080
执行结果: 成功
进程PID: 6428
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:11 GMT
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
*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: www.so.com
> User-Agent: curl/8.12.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Server: openresty
< Date: Sun, 14 Sep 2025 05:23:11 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
< Vary: Accept-Encoding
< Set-Cookie: _S=2k3ve7kqbf6cf3n5qc1qv2tk52; expires=Sun, 14-Sep-2025 05:33:11 GMT; Max-Age=600; path=/
< Expires: Thu, 19 Nov 1981 08:52:00 GMT
< Cache-Control: no-store, no-cache, must-revalidate
< Pragma: no-cache
< php-waf-rep: -
< Set-Cookie: QiHooGUID=CFE82C5919B3DD1EE8BB81C6D79CFC2F.1757827391841; Max-Age=63072000; Domain=so.com; Path=/
< 
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #1 to host localhost left intact
HTTP/1.1 302 Found
Connection: keep-alive
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:23:11 GMT
Location: https://www.so.com/
Server: openresty
Set-Cookie: QiHooGUID=; Max-Age=63072000; Domain=so.com; Path=/

HTTP/1.1 200 Connection established

HTTP/1.1 200 OK
Server: openresty
Date: Sun, 14 Sep 2025 05:23:11 GMT
Content-Type: text/html; charset=UTF-8
Connection: keep-alive
Vary: Accept-Encoding
Set-Cookie: _S=2k3ve7kqbf6cf3n5qc1qv2tk52; expires=Sun, 14-Sep-2025 05:33:11 GMT; Max-Age=600; path=/
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Cache-Control: no-store, no-cache, must-revalidate
Pragma: no-cache
php-waf-rep: -
Set-Cookie: QiHooGUID=CFE82C5919B3DD1EE8BB81C6D79CFC2F.1757827391841; Max-Age=63072000; Domain=so.com; Path=/
---
[2025-09-14 13:23:28.928] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:8080
执行结果: 成功
进程PID: 29448
执行时间: 31.25ms
输出: Note: Using embedded CA bundle, for proxies (233263 bytes)
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying [::1]:8080...
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
* Connected to localhost (::1) port 8080
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
< Date: Sun, 14 Sep 2025 05:23:12 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_8956369090804054251
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
Date: Sun, 14 Sep 2025 05:23:12 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_8956369090804054251
---
[2025-09-14 13:23:34.163] [BUILD] C:\Program Files\Go\bin\go.exe go build -o socks5-websocket-proxy-golang.exe github.com/masx200/socks5-websocket-proxy-golang/cmd
执行结果: 成功
进程PID: 30228
执行时间: 1.21875s
---
[2025-09-14 13:23:34.562] [BUILD] C:\Program Files\Go\bin\go.exe go build -o main.exe ../cmd/main.go
执行结果: 成功
进程PID: 19828
执行时间: 1.484375s
---
[2025-09-14 13:23:36.073] [WEBSOCKET] ./socks5-websocket-proxy-golang.exe -mode server -protocol websocket -addr :38800
执行结果: 成功
进程PID: 29276
---
[2025-09-14 13:23:37.079] [HTTP] ./main.exe -port 10810 -upstream-type websocket -upstream-address ws://localhost:38800
执行结果: 成功
进程PID: 20824
---
[2025-09-14 13:23:40.211] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I http://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 32308
执行时间: 0s
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
< Date: Sun, 14 Sep 2025 05:23:23 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11312453395660935113
< 
  0   277    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
* Connection #0 to host localhost left intact
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: keep-alive
Content-Length: 277
Content-Type: text/html
Date: Sun, 14 Sep 2025 05:23:23 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11312453395660935113
---
[2025-09-14 13:23:40.283] [CURL] D:\迅雷下载\curl-8.12.1_4-win64-mingw\bin\curl.exe curl -v -I https://www.baidu.com -x http://localhost:10810
执行结果: 成功
进程PID: 29760
执行时间: 0s
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
< Date: Sun, 14 Sep 2025 05:23:23 GMT
< Etag: "575e1f60-115"
< Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
< Pragma: no-cache
< Server: bfe/1.0.8.18
< Tr_id: bfe_11528341636127742104
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
Date: Sun, 14 Sep 2025 05:23:23 GMT
Etag: "575e1f60-115"
Last-Modified: Mon, 13 Jun 2016 02:50:08 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
Tr_id: bfe_11528341636127742104
---
