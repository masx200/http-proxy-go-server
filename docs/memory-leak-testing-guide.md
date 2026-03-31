# HTTP Proxy Go Server 内存泄漏测试指南

## 目录
1. [内存泄漏基础知识](#内存泄漏基础知识)
2. [Go语言内存泄漏常见原因](#go语言内存泄漏常见原因)
3. [测试工具安装](#测试工具安装)
4. [内存泄漏测试方法](#内存泄漏测试方法)
5. [项目特定测试策略](#项目特定测试策略)
6. [内存泄漏分析技术](#内存泄漏分析技术)
7. [常见内存泄漏场景](#常见内存泄漏场景)
8. [预防措施](#预防措施)

---

## 内存泄漏基础知识

### 什么是内存泄漏？
内存泄漏是指程序在运行过程中，由于某些原因导致已分配的内存无法被释放，最终导致程序占用的内存持续增长，严重时会导致程序崩溃或系统性能下降。

### Go语言的垃圾回收机制
Go使用垃圾回收（GC）来自动管理内存。但是，即使有GC机制，Go程序仍然可能出现内存泄漏：

1. **引用未释放**：对象仍然被引用但不再使用
2. **Goroutine泄漏**：Goroutine永久阻塞无法退出
3. **资源未关闭**：文件、网络连接等资源未正确关闭

---

## Go语言内存泄漏常见原因

### 1. Goroutine泄漏
```go
// 危险：无限循环的goroutine
go func() {
    for {
        time.Sleep(time.Second)
        // 没有退出条件
    }
}()

// 危险：阻塞的goroutine
ch := make(chan int)
go func() {
    val := <-ch  // 如果没有发送者，永久阻塞
}()
```

### 2. 连接泄漏
```go
// 危险：连接未关闭
resp, err := http.Get("http://example.com")
if err != nil {
    return err
}
// 忘记 resp.Body.Close()
```

### 3. 切片/map引用
```go
// 危险：无限增长的缓存
cache := make(map[string]interface{})
go func() {
    for {
        cache[fmt.Sprint(time.Now())] = getData()  // 永不清理
    }
}()
```

### 4. 定时器未停止
```go
// 危险：定时器未停止
ticker := time.NewTicker(time.Second)
go func() {
    for range ticker.C {
        // 处理任务
    }
}()
// 忘记 ticker.Stop()
```

---

## 测试工具安装

### 1. Go内置工具

#### pprof（CPU/Memory分析）
```bash
# Go 1.21+ 自动包含pprof
go version  # 确认版本
```

#### trace（执行追踪）
```bash
# 内置在Go工具链中
go tool trace --help
```

### 2. 第三方工具

#### go-torch（火焰图生成）
```bash
go install github.com/iovisor/gobtorch/cmd/gobtorch@latest
```

#### heapview（堆内存查看）
```bash
go install github.com/alecthomas/heapview@latest
```

#### leaktest（泄漏检测库）
```bash
go get go.uber.org/tools/leaktest
```

### 3. 系统监控工具

#### Linux/Mac
```bash
# 内存监控
top -p $(pgrep -f "http-proxy-go-server")

# 详细内存信息
ps aux | grep "http-proxy-go-server"

# 持续监控
watch -n 1 'ps aux | grep "http-proxy-go-server"'
```

#### Windows
```powershell
# PowerShell监控
Get-Process | Where-Object {$_.ProcessName -like "*http-proxy*"}

# 持续监控
while($true) { Get-Process -Name "*http-proxy*" | Select-Object ProcessName, CPU, WorkingSet64; Start-Sleep -Seconds 1 }
```

---

## 内存泄漏测试方法

### 1. 基准测试（Benchmark Testing）

#### 创建基准测试文件
创建文件 `cmd/main_benchmark_test.go`：

```go
package main

import (
    "net/http"
    "net/url"
    "testing"
    "time"
)

func BenchmarkMemoryUsage(b *testing.B) {
    // 启动代理服务器
    go func() {
        // 你的代理服务器启动代码
    }()
    time.Sleep(2 * time.Second) // 等待服务器启动

    // 创建HTTP客户端
    client := &http.Client{
        Transport: &http.Transport{
            Proxy: func(req *http.Request) (*url.URL, error) {
                return url.Parse("http://localhost:8080")
            },
        },
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resp, err := client.Get("http://example.com")
        if err != nil {
            b.Fatal(err)
        }
        resp.Body.Close()
    }
}

func BenchmarkDNSCacheMemory(b *testing.B) {
    cache := setupTestCache()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        domain := fmt.Sprintf("test%d.com", i%1000)
        cache.Get(domain)
    }
}
```

#### 运行基准测试
```bash
# 基础基准测试
go test -bench=. -benchmem ./cmd/

# 带内存分析的基准测试
go test -bench=. -benchmem -memprofile=mem.prof ./cmd/

# 分析内存profile
go tool pprof mem.prof
```

### 2. 泄漏检测测试

#### 使用leaktest库
创建文件 `tests/leak_test.go`：

```go
package tests

import (
    "testing"
    "time"
    "go.uber.org/tools/leaktest"
)

func TestNoGoroutineLeak(t *testing.T) {
    // 检测goroutine泄漏
    defer leaktest.Check(t)()

    // 启动代理服务器
    proxy := startTestProxy(t)
    defer proxy.Stop()

    // 执行一些操作
    for i := 0; i < 100; i++ {
        makeProxyRequest(t, proxy)
    }
}

func TestNoMemoryLeak(t *testing.T) {
    // 检测内存泄漏
    defer leaktest.Check(t)()

    // 测试DNS缓存
    cache := setupTestCache()

    for i := 0; i < 1000; i++ {
        cache.Set(fmt.Sprintf("domain%d.com", i), []string{"1.2.3.4"})
        cache.Get(fmt.Sprintf("domain%d.com", i))
    }
}
```

### 3. 压力测试

#### 创建压力测试脚本
创建文件 `tests/stress_test.go`：

```go
package tests

import (
    "fmt"
    "net/http"
    "sync"
    "testing"
    "time"
)

func TestProxyMemoryUnderStress(t *testing.T) {
    proxy := startTestProxy(t)
    defer proxy.Stop()

    var wg sync.WaitGroup
    numRequests := 10000
    numClients := 50

    // 分段执行，避免瞬时内存压力
    for batch := 0; batch < numRequests/numClients; batch++ {
        for i := 0; i < numClients; i++ {
            wg.Add(1)
            go func(reqNum int) {
                defer wg.Done()

                reqID := batch*numClients + reqNum
                makeRequest(t, proxy, fmt.Sprintf("http://test%d.com", reqID))

                // 每1000个请求强制GC
                if reqID % 1000 == 0 {
                    time.Sleep(100 * time.Millisecond)
                }
            }(i)
        }
        wg.Wait()

        // 每批次后检查内存
        if t.Failed() {
            break
        }
    }
}

func TestDNSCacheMemoryStress(t *testing.T) {
    cache := setupTestCache()

    const numDomains = 100000
    domains := make([]string, numDomains)

    for i := 0; i < numDomains; i++ {
        domains[i] = fmt.Sprintf("stress-test-domain-%d.com", i)
    }

    // 写入测试
    for _, domain := range domains {
        cache.Set(domain, []string{"1.2.3.4", "5.6.7.8"})
    }

    // 读取测试
    for i := 0; i < 1000; i++ {
        for _, domain := range domains {
            cache.Get(domain)
        }
        time.Sleep(10 * time.Millisecond) // 避免CPU过载
    }
}
```

---

## 项目特定测试策略

### 1. DNS缓存系统测试

#### 测试文件：`tests/dns_cache_memory_test.go`
```go
package tests

import (
    "testing"
    "time"
    "github.com/masx200/http-proxy-go-server/dnscache"
)

func TestDNSCacheMemoryLeaks(t *testing.T) {
    // 配置测试缓存
    config := &dnscache.Config{
        FilePath:     "/tmp/test_cache.json",
        AOFPath:      "/tmp/test_cache.aof",
        DefaultTTL:   5 * time.Minute,
        SaveInterval: 1 * time.Second,
        AOFInterval:  500 * time.Millisecond,
        Enabled:      true,
    }

    cache, err := dnscache.NewWithConfig(config)
    if err != nil {
        t.Fatalf("Failed to create cache: %v", err)
    }
    defer cache.Close()

    // 测试大量域名缓存
    numDomains := 10000
    for i := 0; i < numDomains; i++ {
        domain := fmt.Sprintf("memleak-test-%d.com", i)
        ips := []string{fmt.Sprintf("192.168.1.%d", i%256)}

        err := cache.Set(domain, ips)
        if err != nil {
            t.Errorf("Failed to set %s: %v", domain, err)
        }

        // 每千个域名检查一次
        if (i+1)%1000 == 0 {
            time.Sleep(100 * time.Millisecond)
        }
    }

    // 测试缓存过期和清理
    time.Sleep(2 * time.Second)

    // 验证缓存可以被正常读取
    result := cache.Get("memleak-test-0.com")
    if result == nil {
        t.Error("Expected cached result")
    }
}

func TestDNSCacheAOFMemory(t *testing.T) {
    config := &dnscache.Config{
        AOFPath:      "/tmp/test_aof.aof",
        DefaultTTL:   1 * time.Minute,
        AOFInterval:  100 * time.Millisecond,
        Enabled:      true,
    }

    cache, err := dnscache.NewWithConfig(config)
    if err != nil {
        t.Fatalf("Failed to create cache: %v", err)
    }
    defer cache.Close()

    // 测试AOF文件内存使用
    for i := 0; i < 1000; i++ {
        domain := fmt.Sprintf("aof-test-%d.com", i)
        cache.Set(domain, []string{"1.2.3.4"})
        time.Sleep(10 * time.Millisecond)
    }
}
```

### 2. 代理连接测试

#### 测试文件：`tests/proxy_connection_test.go`
```go
package tests

import (
    "bufio"
    "net/http"
    "net/url"
    "testing"
    "time"
)

func TestProxyConnectionMemoryLeak(t *testing.T) {
    // 启动测试代理服务器
    proxy := startTestProxyWithAuth(t, "testuser", "testpass")
    defer proxy.Stop()

    // 创建HTTP客户端配置
    proxyURL, _ := url.Parse("http://testuser:testpass@localhost:8080")
    client := &http.Client{
        Transport: &http.Transport{
            Proxy: http.ProxyURL(proxyURL),
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     90 * time.Second,
        },
        Timeout: 30 * time.Second,
    }

    // 发送大量请求
    for i := 0; i < 1000; i++ {
        resp, err := client.Get("http://example.com")
        if err != nil {
            t.Logf("Request %d failed: %v", i, err)
            continue
        }

        // 确保响应体被完全读取和关闭
        _, _ = bufio.NewReader(resp.Body).ReadString('\n')
        resp.Body.Close()

        // 每100个请求暂停一下
        if (i+1)%100 == 0 {
            time.Sleep(1 * time.Second)
            t.Logf("Completed %d requests", i+1)
        }
    }
}

func TestConcurrentConnectionsMemory(t *testing.T) {
    proxy := startTestProxy(t)
    defer proxy.Stop()

    const numClients = 20
    const requestsPerClient = 100

    var wg sync.WaitGroup
    for client := 0; client < numClients; client++ {
        wg.Add(1)
        go func(clientID int) {
            defer wg.Done()

            client := createTestClient(t)
            for i := 0; i < requestsPerClient; i++ {
                resp, err := client.Get(fmt.Sprintf("http://test%d.com", clientID*1000+i))
                if err == nil {
                    resp.Body.Close()
                }
                time.Sleep(50 * time.Millisecond)
            }
        }(client)
    }
    wg.Wait()
}
```

### 3. 长时间运行测试

#### 测试文件：`tests/long_running_test.go`
```go
package tests

import (
    "runtime"
    "testing"
    "time"
)

func TestProxyLongRunning(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long-running test in short mode")
    }

    proxy := startTestProxy(t)
    defer proxy.Stop()

    client := createTestClient(t)
    duration := 10 * time.Minute

    t.Log("Starting long-running memory leak test...")
    startTime := time.Now()

    // 记录初始内存
    var m1 runtime.MemStats
    runtime.ReadMemStats(&m1)

    requestCount := 0
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    for time.Since(startTime) < duration {
        select {
        case <-ticker.C:
            // 每分钟检查内存使用情况
            var m2 runtime.MemStats
            runtime.ReadMemStats(&m2)

            allocatedMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
            t.Logf("After %v: %d requests, %.2f MB allocated",
                time.Since(startTime).Round(time.Minute),
                requestCount, allocatedMB)

            // 如果内存增长超过500MB，可能存在泄漏
            if allocatedMB > 500 {
                t.Errorf("Possible memory leak detected: %.2f MB allocated", allocatedMB)
            }

            // 强制GC并检查内存是否释放
            runtime.GC()
            time.Sleep(5 * time.Second)

        default:
            // 持续发送请求
            for i := 0; i < 100; i++ {
                resp, err := client.Get(fmt.Sprintf("http://test%d.com", requestCount))
                if err == nil {
                    resp.Body.Close()
                }
                requestCount++
            }
            time.Sleep(1 * time.Second)
        }
    }

    t.Logf("Completed %d requests in %v", requestCount, duration)
}
```

---

## 内存泄漏分析技术

### 1. 使用pprof进行内存分析

#### 添加pprof支持到主程序

在 `cmd/main.go` 中添加：

```go
import (
    _ "net/http/pprof"  // 导入pprof
    "net/http"
)

func main() {
    // 在单独的goroutine中启动pprof服务器
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // ... 原有的代理服务器代码
}
```

#### 内存分析命令
```bash
# 1. 获取内存profile
curl http://localhost:6060/debug/pprof/heap > heap.prof

# 2. 分析内存profile
go tool pprof heap.prof

# 3. 在pprof交互界面中
# 查看top内存占用
(pprof) top

# 查看详细内存分配
(pprof) list <function_name>

# 生成图形化报告
(pprof) web

# 查看特定函数的内存分配
(pprof) list DNSCache.Set

# 4. 比较两个时间点的内存差异
# 第一次采样
curl http://localhost:6060/debug/pprof/heap > heap1.prof
# 运行一段时间后再次采样
curl http://localhost:6060/debug/pprof/heap > heap2.prof
# 比较
go tool pprof -base heap1.prof heap2.prof

# 5. 生成火焰图
go tool pprof -http=:8080 heap.prof
```

### 2. 使用trace进行执行追踪

#### 添加trace支持
```go
import (
    "os"
    "runtime/trace"
)

func main() {
    // 开始trace
    f, _ := os.Create("trace.out")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()

    // ... 原有代码
}
```

#### 分析trace
```bash
# 1. 生成trace文件（运行程序后）
# 2. 分析trace
go tool trace trace.out

# 3. 在浏览器中查看详细执行信息
# 将打开本地web服务器显示详细的执行轨迹
```

### 3. 实时内存监控脚本

#### 创建监控脚本：`scripts/monitor_memory.sh`
```bash
#!/bin/bash

# 内存监控脚本
PID=$(pgrep -f "http-proxy-go-server")
if [ -z "$PID" ]; then
    echo "Proxy server not running"
    exit 1
fi

echo "Monitoring proxy server (PID: $PID) memory usage..."
echo "Time,RSS(MB),VSZ(MB),CPU%" > memory_usage.csv

while kill -0 $PID 2>/dev/null; do
    STATS=$(ps -p $PID -o rss,vsz,%cpu --no-headers | awk '{printf "%.2f,%.2f,%.1f", $1/1024, $2/1024, $3}')
    TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')
    echo "$TIMESTAMP,$STATS" >> memory_usage.csv
    echo "$TIMESTAMP: $STATS MB"
    sleep 5
done

echo "Process $PID has terminated"
```

#### Windows PowerShell监控脚本：`scripts/monitor_memory.ps1`
```powershell
# PowerShell内存监控脚本
$processName = "http-proxy-go-server"
$outputFile = "memory_usage.csv"

# 创建CSV头
"Time,WorkingSetMB,PrivateMemoryMB,CPU%" | Out-File $outputFile

Write-Host "Monitoring $processName memory usage..."

while ($true) {
    $process = Get-Process -Name $processName -ErrorAction SilentlyContinue

    if ($process) {
        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $workingSet = [math]::Round($process.WorkingSet64 / 1MB, 2)
        $privateMemory = [math]::Round($process.PrivateMemorySize64 / 1MB, 2)
        $cpu = [math]::Round($process.CPU, 1)

        $line = "$timestamp,$workingSet,$privateMemory,$cpu"
        $line | Out-File $outputFile -Append
        Write-Host "$timestamp: WS=$workingSet MB, PM=$privateMemory MB, CPU=$cpu%"
    } else {
        Write-Host "Process $processName not found"
        break
    }

    Start-Sleep -Seconds 5
}
```

---

## 常见内存泄漏场景

### 1. DNS缓存相关泄漏

#### 问题场景
```go
// 危险：无限增长的缓存
func (c *DNSCache) Set(domain string, ips []string) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.cache.Set(domain, ips, cache.DefaultExpiration)
    // 没有定期清理机制，缓存可能无限增长

    return nil
}
```

#### 检测方法
```bash
# 运行DNS缓存压力测试
go test -v -run TestDNSCacheMemoryStress ./tests/

# 监控内存使用
go tool pprof http://localhost:6060/debug/pprof/heap
```

#### 解决方案
确保DNS缓存有TTL和清理机制，定期检查缓存大小。

### 2. HTTP连接泄漏

#### 问题场景
```go
// 危险：响应体未关闭
resp, err := http.Get(url)
if err != nil {
    return err
}
// 忘记 resp.Body.Close()
```

#### 检测方法
```bash
# 检查文件描述符数量
lsof -p $(pgrep http-proxy-go-server) | wc -l

# 检查网络连接
netstat -anp | grep $(pgrep http-proxy-go-server)
```

#### 解决方案
```go
// 正确的处理方式
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close() // 确保关闭响应体
```

### 3. Goroutine泄漏

#### 问题场景
```go
// 危险：goroutine无法退出
go func() {
    for {
        select {
        case data := <-ch:
            process(data)
        // 没有退出条件，永久阻塞
        }
    }
}()
```

#### 检测方法
```bash
# 检查goroutine数量
curl http://localhost:6060/debug/pprof/goroutine?debug=1

# 在代码中添加检查
import "runtime"
fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
```

#### 解决方案
```go
// 正确的处理方式
ctx, cancel := context.WithCancel(context.Background())
go func() {
    for {
        select {
        case data := <-ch:
            process(data)
        case <-ctx.Done():
            return // 正常退出
        }
    }
}()

// 在适当的时候调用 cancel()
```

### 4. 定时器泄漏

#### 问题场景
```go
// 危险：定时器未停止
ticker := time.NewTicker(1 * time.Second)
go func() {
    for range ticker.C {
        // 定期任务
    }
}()
// 忘记 ticker.Stop()
```

#### 解决方案
```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop() // 确保停止定时器

go func() {
    for range ticker.C {
        // 定期任务
    }
}()
```

---

## 预防措施

### 1. 代码审查清单

#### Goroutine管理
- [ ] 每个goroutine都有明确的退出条件
- [ ] 使用context来控制goroutine生命周期
- [ ] 避免在goroutine中使用阻塞操作而不设置超时
- [ ] 定期检查goroutine数量：`runtime.NumGoroutine()`

#### 资源管理
- [ ] 所有HTTP响应体都使用defer关闭
- [ ] 文件操作后正确关闭文件
- [ ] 网络连接设置超时时间
- [ ] 使用defer确保资源清理

#### 缓存管理
- [ ] 缓存设置合理的TTL
- [ ] 实现缓存大小限制
- [ ] 定期清理过期数据
- [ ] 监控缓存内存使用

#### 并发安全
- [ ] 共享资源正确使用锁保护
- [ ] 避免死锁和竞态条件
- [ ] 使用sync.Pool来重用对象

### 2. 测试策略

#### 单元测试
- 为每个关键组件编写内存测试
- 使用table-driven测试覆盖多种场景
- 添加基准测试监控性能

#### 集成测试
- 模拟真实使用场景
- 长时间运行测试检测缓慢泄漏
- 并发测试检测竞争条件

#### 压力测试
- 大量并发请求测试
- 极限条件下的行为测试
- 资源耗尽情况的处理

### 3. 监控和告警

#### 运行时监控
```go
// 添加到主程序
import (
    "runtime"
    "time"
)

func startMemoryMonitor() {
    ticker := time.NewTicker(5 * time.Minute)
    go func() {
        for range ticker.C {
            var m runtime.MemStats
            runtime.ReadMemStats(&m)

            log.Printf("Memory Stats: Alloc=%d MB, TotalAlloc=%d MB, Sys=%d MB, NumGC=%d",
                m.Alloc/1024/1024,
                m.TotalAlloc/1024/1024,
                m.Sys/1024/1024,
                m.NumGC)

            // 内存使用过高告警
            if m.Alloc > 500*1024*1024 { // 500MB
                log.Printf("WARNING: High memory usage: %d MB", m.Alloc/1024/1024)
            }

            // goroutine数量告警
            numGoroutines := runtime.NumGoroutine()
            if numGoroutines > 1000 {
                log.Printf("WARNING: High goroutine count: %d", numGoroutines)
            }
        }
    }()
}
```

#### 日志分析
- 记录关键操作的内存使用
- 监控日志中的内存告警
- 定期审查日志文件大小

### 4. 最佳实践

#### 编码规范
1. **资源管理**
   - 总是使用defer关闭资源
   - 设置合理的超时时间
   - 使用context管理生命周期

2. **并发编程**
   - 明确goroutine的退出条件
   - 使用带缓冲的channel避免死锁
   - 定期检查goroutine数量

3. **内存优化**
   - 重用对象而不是频繁创建
   - 使用sync.Pool减少GC压力
   - 避免不必要的内存分配

4. **测试驱动**
   - 先编写测试，再编写代码
   - 测试覆盖率不低于80%
   - 包含内存泄漏测试

---

## 快速启动指南

### 1. 基础内存测试
```bash
# 克隆仓库
cd http-proxy-go-server

# 运行基础测试
go test -v -bench=. -benchmem ./tests/

# 启动代理服务器（带pprof）
go run ./cmd/ -hostname 0.0.0.0 -port 8080

# 在另一个终端进行内存分析
go tool pprof http://localhost:6060/debug/pprof/heap
```

### 2. 泄漏检测测试
```bash
# 安装leaktest
go get go.uber.org/tools/leaktest

# 运行泄漏检测
go test -v -run TestNoGoroutineLeak ./tests/
```

### 3. 压力测试
```bash
# 运行压力测试
go test -v -run TestProxyMemoryUnderStress ./tests/

# 长时间运行测试
go test -v -run TestProxyLongRunning ./tests/

# 监控内存使用
./scripts/monitor_memory.sh
```

---

## 故障排查流程

### 发现内存泄漏时的检查步骤

1. **确认泄漏存在**
   ```bash
   # 持续监控内存使用
   watch -n 10 'ps aux | grep http-proxy-go-server'
   ```

2. **获取内存profile**
   ```bash
   curl http://localhost:6060/debug/pprof/heap > heap.prof
   go tool pprof heap.prof
   ```

3. **分析内存分配**
   ```bash
   # 查看最大内存分配者
   (pprof) top

   # 查看特定函数
   (pprof) list <function_name>
   ```

4. **检查goroutine**
   ```bash
   curl http://localhost:6060/debug/pprof/goroutine?debug=1
   ```

5. **生成火焰图**
   ```bash
   go tool pprof -http=:8080 heap.prof
   # 在浏览器中打开 http://localhost:8080
   ```

6. **修复和验证**
   - 根据分析结果修复代码
   - 重新运行测试验证修复效果
   - 监控生产环境内存使用

---

## 总结

内存泄漏测试是一个系统性的过程，需要：

1. **理解内存泄漏的常见原因**
2. **掌握测试工具的使用方法**
3. **编写针对性的测试用例**
4. **建立监控和告警机制**
5. **遵循最佳编码实践**

通过本指南，你应该能够：
- 有效地检测和诊断内存泄漏
- 使用各种工具进行内存分析
- 编写防止内存泄漏的代码
- 建立完善的测试和监控体系

记住，**预防胜于治疗**。良好的编码习惯和完善的测试是防止内存泄漏的最佳方法。

---

## 附录：有用的Go命令和工具

```bash
# 查看Go版本
go version

# 环境变量
go env

# 运行测试并显示覆盖率
go test -cover ./...

# 运行特定测试
go test -run TestName -v ./...

# 竞态检测
go test -race ./...

# 内存分析
go test -memprofile=mem.prof ./...
go tool pprof mem.prof

# CPU分析
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# 查看依赖关系
go mod graph

# 清理缓存
go clean -cache
```

## 相关资源

- [Go官方pprof文档](https://pkg.go.dev/net/http/pprof)
- [Go性能分析最佳实践](https://go.dev/doc/diagnostics)
- [内存泄漏检测技术](https://go.dev/blog/pprof)
- [Goroutine泄漏检测](https://github.com/uber-go/goleak)

---

**文档版本**: 1.0
**最后更新**: 2025-01-31
**维护者**: http-proxy-go-server团队
