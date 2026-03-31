# HTTP Proxy Go Server - pprof 性能分析使用指南

## 📋 概述

本服务器现已内置 pprof 性能分析功能，可以帮助您进行内存泄漏检测、CPU 性能分析和运行时性能调优。

## 🚀 快速开始

### 1. 启用 pprof 功能

#### 方法一：使用命令行参数

```bash
# 启动代理服务器并启用 pprof（默认监听 127.0.0.1:6060）
go run ./cmd/ -enable-pprof

# 自定义 pprof 端口
go run ./cmd/ -enable-pprof -pprof-port 8080

# 允许外部访问（使用 0.0.0.0）
go run ./cmd/ -enable-pprof -pprof-addr 0.0.0.0 -pprof-port 6060

# 完整示例
go run ./cmd/ \
  -hostname 0.0.0.0 \
  -port 8080 \
  -enable-pprof \
  -pprof-port 6060 \
  -username admin \
  -password secret
```

#### 方法二：使用配置文件

在您的 JSON 配置文件中添加（如果支持）：

```json
{
  "hostname": "0.0.0.0",
  "port": 8080,
  "enable_pprof": true,
  "pprof_port": 6060,
  "pprof_addr": "127.0.0.1"
}
```

### 2. 验证 pprof 是否启动

```bash
# 检查 pprof 端点是否可访问
curl http://localhost:6060/debug/pprof/

# 查看可用的 pprof 类型
curl http://localhost:6060/debug/pprof/index
```

## 📊 pprof 功能详解

### 可用的性能分析端点

| 端点 | 说明 | 用途 |
|------|------|------|
| `/debug/pprof/` | pprof 首页 | 查看所有可用的 profile |
| `/debug/pprof/heap` | 堆内存分析 | 检测内存泄漏、内存分配 |
| `/debug/pprof/goroutine` | Goroutine 分析 | 检测 goroutine 泄漏 |
| `/debug/pprof/threadcreate` | 线程创建分析 | 查看线程创建情况 |
| `/debug/pprof/block` | 阻塞操作分析 | 查看同步原语阻塞 |
| `/debug/pprof/mutex` | 互斥锁分析 | 查看锁竞争情况 |
| `/debug/pprof/profile` | CPU profile | CPU 性能分析 |
| `/debug/pprof/trace` | 执行追踪 | 详细的执行轨迹分析 |

## 🔍 使用场景

### 1. 内存泄漏检测

#### 获取内存快照
```bash
# 获取当前内存状态
curl http://localhost:6060/debug/pprof/heap > heap.prof

# 分析内存
go tool pprof heap.prof
```

#### 比较内存变化
```bash
# 获取第一个快照
curl http://localhost:6060/debug/pprof/heap > heap1.prof

# 等待一段时间（例如运行测试）
sleep 30

# 获取第二个快照
curl http://localhost:6060/debug/pprof/heap > heap2.prof

# 比较差异
go tool pprof -base heap1.prof heap2.prof
```

#### 交互式分析命令
```bash
$ go tool pprof heap.prof

# 查看内存占用排名
(pprof) top

# 查看累积内存分配
(pprof) top -cum

# 查看特定函数的内存分配
(pprof) list DNSCache.Set

# 生成图形化报告（需要 graphviz）
(pprof) web

# 生成 PDF 报告
(pprof) pdf

# 查看内存分配来源
(pprof) tree
```

### 2. Goroutine 泄漏检测

#### 检查 Goroutine 数量
```bash
# 查看当前 goroutine 数量
curl http://localhost:6060/debug/pprof/goroutine?debug=1

# 分析 goroutine 泄漏
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

#### 查找阻塞的 Goroutine
```bash
# 查看所有 goroutine 的详细信息
curl http://localhost:6060/debug/pprof/goroutine?debug=2 > goroutine.txt

# 分析阻塞的 goroutine
grep "blocked" goroutine.txt
grep "chan receive" goroutine.txt
grep "semacquire" goroutine.txt
```

### 3. CPU 性能分析

#### 采集 CPU Profile
```bash
# 采集 30 秒的 CPU profile
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof

# 分析 CPU 使用情况
go tool pprof cpu.prof

# 查看 CPU 热点函数
(pprof) top

# 生成火焰图
(pprof) web
```

### 4. 执行追踪

#### 生成 Trace 文件
```bash
# 采集 5 秒的执行追踪
curl http://localhost:6060/debug/pprof/trace?seconds=5 > trace.out

# 在浏览器中查看
go tool trace trace.out
```

## 🛠️ 常用命令速查

### 内存分析
```bash
# 实时查看内存状态
go tool pprof http://localhost:6060/debug/pprof/heap

# 保存内存快照
curl http://localhost:6060/debug/pprof/heap > heap_$(date +%Y%m%d_%H%M%S).prof

# 比较两个时间点的内存差异
go tool pprof -base heap1.prof heap2.prof
```

### Goroutine 分析
```bash
# 查看 goroutine 状态
curl http://localhost:6060/debug/pprof/goroutine?debug=1

# 分析 goroutine 泄漏
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### CPU 分析
```bash
# CPU 性能分析（30秒）
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof
go tool pprof cpu.prof
```

### 一键分析脚本
```bash
#!/bin/bash
# analyze_pprof.sh

echo "正在采集内存 profile..."
curl http://localhost:6060/debug/pprof/heap > heap.prof

echo "正在分析内存..."
go tool pprof -http=:9999 heap.prof

echo "浏览器访问: http://localhost:9999"
```

## 📈 性能测试工作流

### 完整的内存泄漏检测流程

```bash
# 1. 启动带 pprof 的代理服务器
go run ./cmd/ -enable-pprof -port 8080 &

# 2. 在另一个终端运行内存监控
./scripts/monitor_memory.sh http-proxy-go-server

# 3. 获取初始内存快照
curl http://localhost:6060/debug/pprof/heap > heap_initial.prof

# 4. 运行测试或负载
go test -v -run TestMemoryStress ./tests/

# 5. 等待一段时间后获取最终快照
sleep 60
curl http://localhost:6060/debug/pprof/heap > heap_final.prof

# 6. 比较内存变化
go tool pprof -base heap_initial.prof heap_final.prof

# 7. 生成可视化报告
go tool pprof -http=:8081 heap_final.prof
```

### 压力测试 + 性能分析

```bash
# 1. 启动服务器
go run ./cmd/ -enable-pprof -port 8080 &

# 2. 启动 CPU 分析（后台运行）
curl http://localhost:6060/debug/pprof/profile?seconds=60 > cpu.prof &

# 3. 运行压力测试
ab -n 10000 -c 100 http://localhost:8080/

# 4. 等待 CPU 分析完成
wait

# 5. 分析 CPU 使用
go tool pprof -http=:8082 cpu.prof
```

## 🔧 高级技巧

### 1. 定期采集内存快照
```bash
#!/bin/bash
# snapshot_memory.sh

while true; do
    timestamp=$(date +%Y%m%d_%H%M%S)
    curl -s http://localhost:6060/debug/pprof/heap > "heap_${timestamp}.prof"
    echo "采集内存快照: heap_${timestamp}.prof"

    # 显示当前内存使用
    curl -s http://localhost:6060/debug/pprof/heap | grep "heap_alloc" || true

    sleep 300  # 每5分钟采集一次
done
```

### 2. 监控 Goroutine 数量趋势
```bash
#!/bin/bash
# monitor_goroutines.sh

while true; do
    count=$(curl -s http://localhost:6060/debug/pprof/goroutine?debug=1 | grep "^goroutine" | wc -l)
    timestamp=$(date '+%Y-%m-%d %H:%M:%S')
    echo "$timestamp - Goroutines: $count"

    if [ $count -gt 1000 ]; then
        echo "警告: Goroutine 数量过多 ($count)"
        curl -s http://localhost:6060/debug/pprof/goroutine?debug=2 > "goroutine_alert_$(date +%Y%m%d_%H%M%S).txt"
    fi

    sleep 10
done
```

### 3. 自动化性能报告
```bash
#!/bin/bash
# generate_report.sh

REPORT_DIR="pprof_reports_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$REPORT_DIR"

# 采集各种 profile
echo "正在采集性能数据..."
curl -s http://localhost:6060/debug/pprof/heap > "$REPORT_DIR/heap.prof"
curl -s http://localhost:6060/debug/pprof/goroutine > "$REPORT_DIR/goroutine.prof"
curl -s http://localhost:6060/debug/pprof/profile?seconds=30 > "$REPORT_DIR/cpu.prof"

# 生成报告
echo "正在生成报告..."
go tool pprof -png "$REPORT_DIR/heap.prof" > "$REPORT_DIR/heap.png" 2>/dev/null
go tool pprof -png "$REPORT_DIR/goroutine.prof" > "$REPORT_DIR/goroutine.png" 2>/dev/null
go tool pprof -png "$REPORT_DIR/cpu.prof" > "$REPORT_DIR/cpu.png" 2>/dev/null

echo "报告已生成到: $REPORT_DIR"
ls -lh "$REPORT_DIR/"
```

## ⚠️ 注意事项

### 安全考虑

1. **默认绑定地址**: pprof 默认绑定到 `127.0.0.1`，仅允许本地访问
2. **生产环境**: 如果需要在生产环境使用，建议：
   - 使用防火墙限制访问
   - 添加认证机制
   - 仅在需要时启用
3. **敏感信息**: pprof 数据可能包含敏感的运行时信息，请妥善保护

### 性能影响

1. **CPU Profile**: 采集 CPU profile 会对性能产生一定影响（约 10-15%）
2. **内存开销**: pprof 本身的内存开销很小（< 1MB）
3. **建议**: 仅在需要分析和调试时启用，避免在生产环境持续运行

### 限制

1. **图形化报告**: 生成 PDF/PNG 需要 `graphviz` 已安装
2. **火焰图**: 需要浏览器支持 SVG
3. **Trace 分析**: trace 文件可能很大，注意磁盘空间

## 📚 相关资源

- [Go 官方 pprof 文档](https://pkg.go.dev/net/http/pprof)
- [Go 性能分析最佳实践](https://go.dev/doc/diagnostics)
- [pprof GitHub 仓库](https://github.com/google/pprof)
- [项目内存泄漏测试指南](./memory-leak-testing-guide.md)

## 🆘 故障排查

### pprof 无法访问
```bash
# 检查端口是否监听
netstat -tuln | grep 6060

# 检查防火墙
sudo ufw allow 6060  # Ubuntu/Debian

# 确认 pprof 已启用
ps aux | grep http-proxy-go-server
```

### 采集超时
```bash
# 增加超时时间
curl --max-time 300 http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof
```

### 图形化报告生成失败
```bash
# 安装 graphviz
# Ubuntu/Debian
sudo apt-get install graphviz

# CentOS/RHEL
sudo yum install graphviz

# macOS
brew install graphviz
```

---

**最后更新**: 2025-01-31
**版本**: 1.0.0