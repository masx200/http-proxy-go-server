# 内存泄漏测试和 pprof 功能实现总结

## 📋 完成的工作

### 1. 核心功能实现

#### ✅ 添加 pprof 命令行参数
在 `cmd/main.go` 中添加了以下命令行参数：

- `-enable-pprof`: 启用 pprof 性能分析功能（默认：false）
- `-pprof-port`: pprof 服务器端口（默认：6060）
- `-pprof-addr`: pprof 服务器绑定地址（默认：127.0.0.1）

#### ✅ 集成 pprof 服务器
- 在 main 函数中添加了 pprof 服务器的自动启动逻辑
- pprof 在独立的 goroutine 中运行，不影响代理服务器性能
- 提供详细的日志输出，方便调试

### 2. 文档创建

#### 📄 核心文档
1. **[内存泄漏测试完整指南](./memory-leak-testing-guide.md)** - 3万字详细指南
   - 内存泄漏基础知识
   - Go 语言内存泄漏常见原因
   - 测试工具安装和使用
   - 针对项目的测试策略
   - 内存泄漏分析技术
   - 常见问题场景和解决方案

2. **[pprof 使用指南](./PPROF_USAGE.md)** - pprof 功能详细说明
   - 快速开始指南
   - 所有 pprof 端点说明
   - 内存泄漏检测步骤
   - Goroutine 泄漏检测
   - CPU 性能分析
   - 常用命令速查

3. **[内存测试快速指南](./MEMORY_TESTING_README.md)** - 快速参考
   - 快速开始步骤
   - 测试方法汇总
   - 监控工具使用
   - 故障排查流程

#### 🧪 测试文件
4. **[内存泄漏测试套件](../tests/memory_leak_test.go)**
   - DNS 缓存内存泄漏测试
   - Goroutine 泄漏检测
   - HTTP 连接内存泄漏测试
   - AOF 内存使用测试
   - 压力测试和基准测试

#### 🛠️ 实用脚本
5. **[Linux/Mac 内存监控脚本](../scripts/monitor_memory.sh)**
   - 实时内存监控
   - CSV 数据导出
   - 网络连接统计
   - Goroutine 数量监控

6. **[Windows 内存监控脚本](../scripts/monitor_memory.ps1)**
   - PowerShell 版本监控工具
   - 内存统计和告警
   - CSV 数据导出

7. **[pprof 启动脚本](../scripts/start_with_pprof.sh)**
   - 一键启动带 pprof 的代理服务器
   - 自动端口检查
   - 使用提示和命令示例

8. **[pprof 测试脚本](../scripts/test_pprof.sh)**
   - 自动化 pprof 功能测试
   - 端点可用性检查
   - Profile 采集和分析
   - 性能检查建议

## 🚀 使用方法

### 快速开始

#### 1. 启用 pprof 功能
```bash
# 启动代理服务器并启用 pprof
go run ./cmd/ -enable-pprof

# 自定义端口和地址
go run ./cmd/ -enable-pprof -pprof-port 8080 -pprof-addr 0.0.0.0
```

#### 2. 运行内存测试
```bash
# 运行所有内存泄漏测试
go test -v -run TestMemory ./tests/

# 运行压力测试
go test -v -run TestMemoryStress ./tests/

# 运行基准测试
go test -bench=. -benchmem ./tests/
```

#### 3. 监控内存使用
```bash
# Linux/Mac
chmod +x scripts/monitor_memory.sh
./scripts/monitor_memory.sh http-proxy-go-server

# Windows PowerShell
.\scripts\monitor_memory.ps1 http-proxy-go-server
```

#### 4. 分析内存泄漏
```bash
# 采集内存 profile
curl http://localhost:6060/debug/pprof/heap > heap.prof

# 分析内存
go tool pprof heap.prof

# 在交互界面中：
(pprof) top        # 查看内存占用排名
(pprof) web        # 生成图形化报告
(pprof) list FuncName  # 查看特定函数
```

## 📊 测试覆盖范围

### 内存泄漏检测
- ✅ DNS 缓存系统
- ✅ HTTP 连接管理
- ✅ Goroutine 泄漏
- ✅ AOF 文件操作
- ✅ 定时器清理

### 性能分析
- ✅ CPU 性能分析
- ✅ 内存分配分析
- ✅ Goroutine 状态分析
- ✅ 锁竞争分析
- ✅ 阻塞操作分析

### 压力测试
- ✅ 大量并发请求
- ✅ 长时间运行稳定性
- ✅ 极限条件测试
- ✅ 资源耗尽处理

## 🔍 关键特性

### 1. 非侵入式监控
- pprof 运行在独立端口
- 不影响代理服务器性能
- 可随时启用/禁用

### 2. 完整的工具链
- 内存泄漏检测
- 性能分析
- 实时监控
- 自动化测试

### 3. 跨平台支持
- Linux
- macOS
- Windows

### 4. 详细的文档
- 快速开始指南
- 详细技术文档
- 实用脚本示例
- 故障排查指南

## 📈 预期效果

### 开发阶段
- 早期发现内存泄漏
- 性能瓶颈识别
- 代码质量提升

### 测试阶段
- 自动化内存泄漏检测
- 性能回归测试
- 稳定性验证

### 生产环境
- 实时性能监控
- 问题快速定位
- 性能优化指导

## 🎯 下一步建议

### 1. 集成到 CI/CD
```yaml
# .github/workflows/memory-test.yml
name: Memory Leak Test
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Run memory tests
        run: |
          go test -v -run TestMemory ./tests/
          go test -bench=. -benchmem ./tests/
```

### 2. 定期性能审查
- 每周运行一次完整的内存泄漏测试
- 每次重大变更后进行性能分析
- 维护性能基线数据

### 3. 监控告警
```bash
# 设置内存使用告警
if [ $(ps aux | grep http-proxy-go-server | awk '{print $6}') -gt 1048576 ]; then
    echo "内存使用超过1GB"
    # 发送告警通知
fi
```

### 4. 文档维护
- 根据实际使用情况更新文档
- 添加新发现的内存泄漏场景
- 分享性能优化经验

## 📞 获取帮助

- 详细文档：[docs/PPROF_USAGE.md](./PPROF_USAGE.md)
- 测试指南：[docs/MEMORY_TESTING_README.md](./MEMORY_TESTING_README.md)
- 完整指南：[docs/memory-leak-testing-guide.md](./memory-leak-testing-guide.md)

## ✅ 验证清单

使用以下清单验证实现是否完整：

- [x] pprof 命令行参数可以正常使用
- [x] pprof 服务器可以正常启动
- [x] 所有 pprof 端点可以正常访问
- [x] 内存泄漏测试可以全部通过
- [x] 监控脚本可以正常工作
- [x] 文档清晰易懂
- [x] 脚本具有执行权限

## 🎉 总结

本次实现为 HTTP Proxy Go Server 添加了完整的内存泄漏检测和性能分析能力，包括：

1. **pprof 功能集成** - 方便的性能分析工具
2. **完整测试套件** - 自动化内存泄漏检测
3. **监控工具** - 实时内存和性能监控
4. **详细文档** - 使用指南和最佳实践

这些工具和文档将帮助开发团队：
- 早期发现和修复内存泄漏
- 优化服务器性能
- 提高代码质量
- 增强系统稳定性

---

**实现日期**: 2025-01-31
**版本**: 1.0.0
**状态**: ✅ 完成并测试通过