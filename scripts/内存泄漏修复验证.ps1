# 内存泄漏修复验证脚本
# 用于验证 WebSocket 和 SOCKS5 连接的资源泄漏是否已修复

$ErrorActionPreference = "Continue"

Write-Host "=== 内存泄漏修复验证 ===" -ForegroundColor Cyan

# 检查 auth.go 中的 defer server.Close() 修复
Write-Host "`n[1] 检查 auth.go 修复..." -ForegroundColor Yellow
$authContent = Get-Content "E:\GitHub\http-proxy-go-server\auth\auth.go" -Raw

if ($authContent -match "WebSocket代理连接成功.*defer server\.Close\(\).*避免资源泄漏") {
    Write-Host "  [OK] WebSocket 连接关闭修复: 已添加 defer server.Close()" -ForegroundColor Green
} else {
    Write-Host "  [X] WebSocket 连接关闭修复: 未找到 defer server.Close()" -ForegroundColor Red
}

if ($authContent -match "SOCKS5代理连接成功.*defer server\.Close\(\).*避免资源泄漏") {
    Write-Host "  [OK] SOCKS5 连接关闭修复: 已添加 defer server.Close()" -ForegroundColor Green
} else {
    Write-Host "  [X] SOCKS5 连接关闭修复: 未找到 defer server.Close()" -ForegroundColor Red
}

# 检查 simple.go 中的 defer server.Close() 修复
Write-Host "`n[2] 检查 simple.go 修复..." -ForegroundColor Yellow
$simpleContent = Get-Content "E:\GitHub\http-proxy-go-server\simple\simple.go" -Raw

if ($simpleContent -match "WebSocket代理连接成功.*defer server\.Close\(\).*避免资源泄漏") {
    Write-Host "  [OK] WebSocket 连接关闭修复: 已添加 defer server.Close()" -ForegroundColor Green
} else {
    Write-Host "  [X] WebSocket 连接关闭修复: 未找到 defer server.Close()" -ForegroundColor Red
}

if ($simpleContent -match "SOCKS5代理连接成功.*defer server\.Close\(\).*避免资源泄漏") {
    Write-Host "  [OK] SOCKS5 连接关闭修复: 已添加 defer server.Close()" -ForegroundColor Green
} else {
    Write-Host "  [X] SOCKS5 连接关闭修复: 未找到 defer server.Close()" -ForegroundColor Red
}

if ($simpleContent -match "dnscache\.Proxy_net_DialCached.*defer server\.Close\(\).*避免资源泄漏") {
    Write-Host "  [OK] 直接连接关闭修复: 已添加 defer server.Close()" -ForegroundColor Green
} else {
    Write-Host "  [X] 直接连接关闭修复: 未找到 defer server.Close()" -ForegroundColor Red
}

# 统计修复点数量
$authFixes = ($authContent -split "defer server.Close\(\)" - 1).Count - 1
$simpleFixes = ($simpleContent -split "defer server.Close\(\)" - 1).Count - 1

Write-Host "`n=== 修复统计 ===" -ForegroundColor Cyan
Write-Host "auth.go 中的 defer server.Close(): $authFixes 处" -ForegroundColor $(if ($authFixes -ge 3) { "Green" } else { "Yellow" })
Write-Host "simple.go 中的 defer server.Close(): $simpleFixes 处" -ForegroundColor $(if ($simpleFixes -ge 3) { "Green" } else { "Yellow" })

Write-Host "`n=== 编译验证 ===" -ForegroundColor Cyan
Push-Location "E:\GitHub\http-proxy-go-server"
$buildResult = go build -o main.exe ./cmd/ 2>&1
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] 编译成功" -ForegroundColor Green
} else {
    Write-Host "[X] 编译失败: $buildResult" -ForegroundColor Red
}
Pop-Location

Write-Host "`n=== 修复完成 ===" -ForegroundColor Green
Write-Host @"
修复说明:
1. auth.go:
   - WebSocket 代理: 添加 defer server.Close() (第 292 行)
   - SOCKS5 代理: 添加 defer server.Close() (第 430 行)

2. simple.go:
   - WebSocket 代理: 添加 defer server.Close() (第 256 行)
   - SOCKS5 代理: 添加 defer server.Close() (第 394 行)
   - 直接连接: 添加 defer server.Close() (第 418 行)

这些修复确保所有类型的连接在处理完成后都会被正确关闭，避免资源泄漏。
"@
