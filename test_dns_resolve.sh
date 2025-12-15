#!/bin/bash

echo "=== DNS解析日志修复测试 ==="
echo "测试-upstream-resolve-ips参数是否会产生DNS解析日志"
echo

# 启动SOCKS5测试服务器（后台运行）
echo "1. 启动SOCKS5测试服务器..."
go run -v ./tests/socks5_server/socks5_test_server_standalone.go &
SOCKS5_PID=$!
echo "SOCKS5服务器PID: $SOCKS5_PID"
sleep 3

# 启动HTTP代理服务器（后台运行）
echo "2. 启动带-upstream-resolve-ips的HTTP代理服务器..."
./main.exe -upstream-resolve-ips \
    -upstream-type socks5 \
    -upstream-username g7envpwz14b0u55 \
    -upstream-password juvytdsdzc225pq \
    -upstream-address socks5://127.0.0.1:44444 \
    -dohurl https://dns.alidns.com/dns-query \
    -dohip 223.5.5.5 \
    -dohalpn h3 \
    -hostname 127.0.0.1 \
    -port 8080 &
PROXY_PID=$!
echo "代理服务器PID: $PROXY_PID"
sleep 3

echo "3. 发送测试请求..."
echo "测试HTTP请求（应该显示DNS解析日志）:"
curl -v -I -X GET -x http://localhost:8080 http://www.baidu.com 2>&1 | grep -E "(upstream-resolve-ips|Resolving|resolved)"

echo
echo "测试HTTPS请求（应该显示DNS解析日志）:"
curl -v -I -X GET -x http://localhost:8080 https://www.baidu.com 2>&1 | grep -E "(upstream-resolve-ips|Resolving|resolved)"

echo
echo "4. 清理进程..."
kill $PROXY_PID $SOCKS5_PID 2>/dev/null
wait 2>/dev/null

echo "=== 测试完成 ==="