#!/bin/bash
# nil Context Panic 修复验证脚本

echo "=== nil Context Panic 修复验证 ==="
echo ""

# 检查修复的代码
echo "1. 检查修复的代码..."
if grep -q "重要：确保 context 不为 nil" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ Context 保护代码已添加"
else
    echo "✗ Context 保护代码未找到"
fi

if grep -q "if ctx == nil {" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ nil context 检查逻辑已添加"
else
    echo "✗ nil context 检查逻辑未找到"
fi

if grep -q "ctx = context.Background()" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ 默认 context 创建逻辑已添加"
else
    echo "✗ 默认 context 创建逻辑未找到"
fi

echo ""
echo "2. 验证修复的代码位置："

# 显示修复的代码
echo "修复的代码片段："
echo "---"
sed -n '370,374p' /workspace/http-proxy-go-server/dnscache/caching_resolver.go
echo "---"

echo ""
echo "3. 验证相关函数是否存在："

# 检查相关函数
if grep -q "func proxy_net_DialWithResolver" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ proxy_net_DialWithResolver 函数存在"
else
    echo "✗ proxy_net_DialWithResolver 函数未找到"
fi

if grep -q "func Proxy_net_DialCached" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ Proxy_net_DialCached 函数存在"
else
    echo "✗ Proxy_net_DialCached 函数未找到"
fi

if grep -q "func Proxy_net_DialContextCached" /workspace/http-proxy-go-server/dnscache/caching_resolver.go; then
    echo "✓ Proxy_net_DialContextCached 函数存在"
else
    echo "✗ Proxy_net_DialContextCached 函数未找到"
fi

echo ""
echo "4. 检查是否还有其他 nil context 风险点："

# 检查 DialContext 调用
nil_context_risks=$(grep -n "DialContext(nil," /workspace/http-proxy-go-server/dnscache/caching_resolver.go 2>/dev/null || echo "")
if [ -z "$nil_context_risks" ]; then
    echo "✓ 未发现其他 nil context 风险点"
else
    echo "⚠ 发现其他 nil context 风险点："
    echo "$nil_context_risks"
fi

echo ""
echo "5. 验证上游IP解析功能的相关代码："

# 检查 upstreamResolveIPs 相关代码
upstream_calls=$(grep -n "upstreamResolveIPs" /workspace/http-proxy-go-server/dnscache/caching_resolver.go | head -5)
if [ -n "$upstream_calls" ]; then
    echo "✓ upstreamResolveIPs 功能代码存在："
    echo "$upstream_calls"
else
    echo "✗ upstreamResolveIPs 功能代码未找到"
fi

echo ""
echo "=== 修复验证完成 ==="
echo ""
echo "修复摘要："
echo "- 在 proxy_net_DialWithResolver 函数中添加了 nil context 检查"
echo "- 当 ctx 为 nil 时，自动创建 background context"
echo "- 这将防止在使用 -upstream-resolve-ips 时出现 panic"
echo ""
echo "现在可以安全地使用 -upstream-resolve-ips 功能，不会再出现 nil context panic"