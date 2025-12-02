# 命令日志记录系统设计文档

## 概述

为三个测试文件（proxy_test.go、proxy_doh_test.go、websocket_socks5_test.go）中的所有
exec.Command 命令添加统一的日志记录功能。

## 当前分析

### 发现的 exec.Command 使用情况

#### proxy_test.go (7 个实例)

1. `go build -o main.exe ../cmd/` - 编译代理服务器
2. `./main.exe` - 启动代理服务器
3. `curl -v -I http://www.baidu.com -x http://localhost:8080` - 测试 1
4. `curl -v -I -L http://www.so.com -x http://localhost:8080` - 测试 2
5. `curl -v -I https://www.baidu.com -x http://localhost:8080` - 测试 3
6. `taskkill /F /IM go.exe` - 终止 go 进程
7. `netstat -ano | findstr :1080` - 查找 1080 端口进程

#### proxy_doh_test.go (7 个实例)

1. `go build -o main.exe ../cmd/` - 编译代理服务器
2. `./main.exe -dohurl ...` - 启动 DOH 代理服务器
3. `curl -v -I http://www.baidu.com -x http://localhost:8080` - 测试 1
4. `curl -v -I -L http://www.so.com -x http://localhost:8080` - 测试 2
5. `curl -v -I https://www.baidu.com -x http://localhost:8080` - 测试 3
6. `taskkill /F /IM go.exe` - 终止 go 进程
7. `netstat -ano | findstr :1080` - 查找 1080 端口进程

#### websocket_socks5_test.go (7 个实例)

1. `go build -o main.exe ../cmd/` - 编译代理服务器
2. `./main.exe -mode server -protocol websocket -addr :8080` - 启动 WebSocket
   服务器
3. `./main.exe -mode server -protocol socks5 -addr :18080 -upstream-type websocket -upstream-address ws://localhost:8080` -
   启动 SOCKS5 服务器
4. `curl -v -I http://www.baidu.com -x socks5://localhost:18080` - 测试 1
5. `curl -v -I https://www.baidu.com -x socks5://localhost:18080` - 测试 2
6. `taskkill /F /IM go.exe` - 终止 go 进程
7. ProcessManager.go 中的 `taskkill /F /T /PID ...` - 终止子进程

## 设计方案

### 1. 日志文件设计

- **文件名**: `command_execution_log.log`
- **位置**: 与测试文件同目录
- **格式**: UTF-8 编码

### 2. 日志格式

```
[时间戳] [命令类型] 命令内容
执行结果: [成功/失败]
进程PID: [PID]
执行时间: [耗时]
输出: [命令输出]
错误: [错误信息，如有]
---
```

### 3. 命令类型分类

- **BUILD**: go build 命令
- **SERVER**: 代理服务器启动命令
- **TEST**: curl 测试命令
- **SYSTEM**: 系统管理命令 (taskkill, netstat 等)
- **PROCESS**: 进程管理命令

### 4. 实现方案

#### 步骤 1: 扩展 ProcessManager

在`ProcessManager.go`中添加命令日志记录功能：

- 添加日志文件写入方法
- 添加命令执行日志记录方法
- 添加命令执行结果记录方法

#### 步骤 2: 修改测试文件

修改三个测试文件中的所有`exec.Command`调用：

- 在命令执行前记录命令内容
- 在命令执行后记录执行结果
- 确保不影响现有功能

#### 步骤 3: 日志记录时机

- **命令创建时**: 记录完整的命令字符串
- **命令启动时**: 记录进程 PID 和启动时间
- **命令完成时**: 记录执行结果、输出和耗时

### 5. 具体实现细节

#### 日志记录函数设计

```go
// LogCommand 记录命令执行
func (pm *ProcessManager) LogCommand(cmd *exec.Cmd, cmdType string) error

// LogCommandResult 记录命令执行结果
func (pm *ProcessManager) LogCommandResult(cmd *exec.Cmd, err error, output string) error
```

#### 使用示例

```go
// 原代码
buildCmd := exec.Command("go", "build", "-o", "main.exe", "../cmd/")
buildCmd.Stdout = multiWriter
if err := buildCmd.Run(); err != nil {
    t.Fatalf("编译代理服务器失败: %v", err)
}

// 修改后
buildCmd := exec.Command("go", "build", "-o", "main.exe", "../cmd/")
buildCmd.Stdout = multiWriter
processManager.LogCommand(buildCmd, "BUILD")
if err := buildCmd.Run(); err != nil {
    processManager.LogCommandResult(buildCmd, err, "")
    t.Fatalf("编译代理服务器失败: %v", err)
}
processManager.LogCommandResult(buildCmd, nil, "")
```

### 6. 日志文件管理

- 每次测试运行时追加新日志
- 保留历史日志，便于调试
- 日志文件不会过大，因为单次测试命令数量有限

### 7. 兼容性考虑

- 不影响现有测试功能
- 不改变现有输出格式
- 日志记录失败时不影响测试执行
- 跨平台兼容（Windows/Linux/Mac）

## 预期效果

1. **完整的命令执行追踪**: 所有 exec.Command 的执行都被记录
2. **详细的执行信息**: 包括命令内容、执行时间、结果、输出等
3. **统一的日志格式**: 便于阅读和分析
4. **调试友好**: 当测试失败时，可以查看详细的命令执行日志
5. **性能影响小**: 日志记录不会显著影响测试执行速度

## 实施计划

1. 扩展 ProcessManager 添加日志功能
2. 修改 proxy_test.go 添加日志记录
3. 修改 proxy_doh_test.go 添加日志记录
4. 修改 websocket_socks5_test.go 添加日志记录
5. 测试验证功能正常
6. 优化和调整日志格式
