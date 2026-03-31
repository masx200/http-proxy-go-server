# HTTP Proxy Go Server 内存泄漏测试快速指南

## 📋 目录

1. [快速开始](#快速开始)
2. [测试方法](#测试方法)
3. [监控工具](#监控工具)
4. [分析技巧](#分析技巧)
5. [故障排查](#故障排查)

---

## 🚀 快速开始

### 1. 准备环境

#### 安装依赖
```bash
# 安装leaktest库
go get go.uber.org/tools/leaktest

# 验证Go版本（建议1.21+）
go version
```

#### 构建项目
```bash
cd http-proxy-go-server
go build -o http-proxy-go-server ./cmd/
```

### 2. 运行基础测试

#### 单元测试
```bash
# 运行所有测试
go test ./tests/ -v

# 运行内存泄漏测试
go test ./tests/ -run TestMemory -v

# 运行基准测试
go test ./tests/ -bench=. -benchmem
```

#### 泄漏检测
```bash
# 检测goroutine泄漏
go test ./tests/ -run TestNoGoroutineLeak -v

# 检测内存泄漏
go test ./tests/ -run TestDNSCacheMemoryLeaks -v
```

---

## 🧪 测试方法

### 自动化测试

#### 1. 内存泄漏测试
```bash
# 运行专门的内存泄漏测试
go test -v -run TestMemoryLeak ./tests/

# 运行压力测试
go test -v -run TestMemoryStress ./tests/

# 长时间运行测试
go test -v -run TestLongRunning ./tests/ -timeout 30m
```

#### 2. 基准性能测试
```bash
# 基准测试并生成内存profile
go test -bench=. -benchmem -memprofile=mem.prof ./tests/

# 分析内存profile
go tool pprof mem.prof

# 在pprof交互界面中:
(pprof) top        # 查看最大内存分配者
(pprof) list FuncName  # 查看特定函数
(pprof) web        # 生成图形化报告
```

### 手动测试

#### 1. 启动带pprof的代理服务器

**Linux/Mac:**
```bash
# 使用启动脚本
chmod +x scripts/start_with_pprof.sh
./scripts/start_with_pprof.sh

# 或手动启动
go run ./cmd/ -hostname 0.0.0.0 -port 8080 &
```

**Windows:**
```powershell
# 直接运行
go run ./cmd/ -hostname 0.0.0.0 -port 8080
```

#### 2. 运行监控工具

**Linux/Mac:**
```bash
# 启动内存监控
chmod +x scripts/monitor_memory.sh
./scripts/monitor_memory.sh http-proxy-go-server

# 查看生成的CSV文件
cat memory_usage_*.csv
```

**Windows:**
```powershell
# 启动内存监控
.\scripts\monitor_memory.ps1 http-proxy-go-server

# 查看生成的CSV文件
Import-Csv memory_usage_*.csv | Format-Table -AutoSize
```

---

## 📊 监控工具

### 1. pprof 内存分析

#### 获取内存profile
```bash
# 获取当前内存状态
curl http://localhost:6060/debug/pprof/heap > heap.prof

# 比较30秒前后的内存差异
curl http://localhost:6060/debug/pprof/heap > heap1.prof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap2.prof
go tool pprof -base heap1.prof heap2.prof
```

#### 分析内存分配
```bash
# 启动pprof交互界面
go tool pprof heap.prof

# 常用命令:
(pprof) top                    # Top内存分配者
(pprof) top -cum               # 累积内存分配
(pprof) list <function>        # 查看函数详情
(pprof) web                    # 生成图形化报告
(pprof) pdf                    # 生成PDF报告
(pprof) peek <regex>           # 查看匹配的函数
```

### 2. Goroutine 分析

#### 检查goroutine数量
```bash
# 简单查看goroutine数量
curl http://localhost:6060/debug/pprof/goroutine?debug=1

# 详细goroutine信息
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# 在pprof中分析
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### 3. 实时监控

#### 使用脚本监控
```bash
# Linux/Mac - 持续监控内存
watch -n 5 'ps aux | grep http-proxy-go-server'

# Windows PowerShell - 监控特定进程
while($true) { Get-Process -Name "*http-proxy*" | Select-Object ProcessName, CPU, WorkingSet64; Start-Sleep -Seconds 5 }
```

#### 自定义监控
```bash
# 查看文件描述符（Linux）
lsof -p $(pgrep http-proxy-go-server) | wc -l

# 查看网络连接
netstat -anp | grep $(pgrep http-proxy-go-server)

# 查看线程数
ps -o nlwp $(pgrep http-proxy-go-server)
```

---

## 🔍 分析技巧

### 1. 识别内存泄漏模式

#### 正常内存使用模式
```
启动 → 稳定增长 → GC清理 → 回到基线 → 循环
```

#### 内存泄漏模式
```
启动 → 持续增长 → GC无法有效清理 → 内存持续增长
```

### 2. 重点检查区域

#### DNS缓存系统
```bash
# 测试DNS缓存内存
go test -v -run TestDNSCacheMemoryLeaks ./tests/

# 检查缓存大小
curl http://localhost:6060/debug/pprof/heap | grep DNSCache
```

#### HTTP连接管理
```bash
# 测试连接泄漏
go test -v -run TestHTTPConnectionMemoryLeaks ./tests/

# 检查文件描述符
lsof -p $(pgrep http-proxy-go-server) | grep -c "TCP"
```

#### Goroutine泄漏
```bash
# 检查goroutine数量
curl http://localhost:6060/debug/pprof/goroutine?debug=1 | grep "^goroutine" | wc -l

# 分析goroutine栈
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### 3. 性能分析

#### CPU分析
```bash
# 采集30秒CPU profile
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof

# 分析CPU使用
go tool pprof cpu.prof
(pprof) top        # 查看CPU热点
(pprof) web        # 生成火焰图
```

#### 内存分配分析
```bash
# 查看内存分配速率
go tool pprof http://localhost:6060/debug/pprof/alloc_objects

# 查看分配的对象
(pprof) top        # 查看最多分配的对象
(pprof) list FuncName  # 查看函数分配详情
```

---

## 🛠️ 故障排查

### 问题1: 内存持续增长

#### 诊断步骤
```bash
# 1. 确认内存增长
./scripts/monitor_memory.sh http-proxy-go-server

# 2. 获取内存profile
curl http://localhost:6060/debug/pprof/heap > heap.prof

# 3. 分析内存分配
go tool pprof heap.prof
(pprof) top
```

#### 可能原因
- DNS缓存无限制增长
- HTTP连接未正确关闭
- Goroutine泄漏
- 缓存未设置TTL

### 问题2: Goroutine数量过多

#### 诊断步骤
```bash
# 1. 检查goroutine数量
curl http://localhost:6060/debug/pprof/goroutine?debug=1 > goroutine.txt

# 2. 分析goroutine状态
grep "blocked" goroutine.txt | wc -l
grep "chan receive" goroutine.txt | wc -l

# 3. 在pprof中分析
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

#### 可能原因
- Channel阻塞
- 死锁
- 无限循环的goroutine
- Context未正确取消

### 问题3: pprof无法访问

#### 解决方案
```bash
# 1. 确认pprof端口
netstat -tuln | grep 6060

# 2. 检查防火墙
sudo ufw allow 6060  # Ubuntu/Debian
sudo firewall-cmd --add-port=6060/tcp  # CentOS/RHEL

# 3. 确认pprof已导入
grep -r "net/http/pprof" ./cmd/
```

---

## 📈 预防措施

### 代码审查检查清单

#### 资源管理
- [ ] 所有HTTP响应都使用`defer resp.Body.Close()`
- [ ] 文件操作后正确关闭文件
- [ ] 定时器使用`defer ticker.Stop()`
- [ ] Context正确使用和取消

#### Goroutine管理
- [ ] 每个goroutine都有明确的退出条件
- [ ] 使用context控制goroutine生命周期
- [ ] 避免在goroutine中永久阻塞
- [ ] 定期检查goroutine数量

#### 缓存管理
- [ ] 设置合理的TTL
- [ ] 实现缓存大小限制
- [ ] 定期清理过期数据
- [ ] 监控缓存内存使用

### 测试策略

#### 单元测试
```bash
# 为关键组件编写内存测试
go test -v -run TestMemory ./tests/

# 确保测试覆盖率
go test -cover ./tests/
```

#### 集成测试
```bash
# 模拟真实使用场景
go test -v -run TestProxyLongRunning ./tests/

# 并发测试
go test -v -run TestConcurrent ./tests/
```

#### 压力测试
```bash
# 大量并发请求
go test -v -run TestMemoryStress ./tests/

# 长时间运行
go test -v -timeout 1h -run TestLongRunning ./tests/
```

---

## 📞 获取帮助

### 文档资源
- [Go官方pprof文档](https://pkg.go.dev/net/http/pprof)
- [Go性能分析最佳实践](https://go.dev/doc/diagnostics)
- [项目内存泄漏测试指南](./memory-leak-testing-guide.md)

### 常用命令速查

```bash
# 测试相关
go test ./tests/ -v                    # 运行所有测试
go test -bench=. -benchmem ./tests/   # 基准测试
go test -race ./tests/                # 竞态检测

# 分析相关
go tool pprof heap.prof                # 分析内存
go tool pprof cpu.prof                 # 分析CPU
go tool trace trace.out                # 分析执行轨迹

# 监控相关
./scripts/monitor_memory.sh            # Linux/Mac监控
.\scripts\monitor_memory.ps1           # Windows监控
```

### 报告问题
如果发现内存泄漏问题，请提供以下信息：
1. 内存使用监控数据（CSV文件）
2. pprof分析结果
3. 重现步骤
4. Go版本和操作系统信息

---

## 🎯 快速参考

### 内存泄漏检测流程
```bash
# 1. 启动带pprof的代理服务器
./scripts/start_with_pprof.sh

# 2. 在另一个终端启动监控
./scripts/monitor_memory.sh http-proxy-go-server

# 3. 运行测试
go test -v -run TestMemoryStress ./tests/

# 4. 分析内存profile
curl http://localhost:6060/debug/pprof/heap > heap.prof
go tool pprof heap.prof

# 5. 查看分析结果
(pprof) top
(pprof) web
```

### 日常监控建议
- 每次代码变更后运行内存测试
- 在测试环境运行长时间稳定性测试
- 定期检查生产环境的内存使用情况
- 设置内存使用告警阈值

---

**最后更新**: 2025-01-31
**维护者**: http-proxy-go-server团队