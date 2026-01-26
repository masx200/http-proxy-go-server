# --ip-priority 参数说明

## 概述

`--ip-priority` 参数用于配置代理服务器在连接目标主机时，如何从解析出的多个 IP
地址中选择最佳的连接地址。

该参数特别适用于以下场景：

- 目标域名同时解析出 IPv4 和 IPv6 地址
- 需要优先使用特定 IP 版本进行连接
- 在不同网络环境下灵活调整 IP 连接策略

## 参数值

`--ip-priority` 支持以下三种策略：

| 参数值   | 说明      | 优先级策略                                                |
| -------- | --------- | --------------------------------------------------------- |
| `ipv4`   | IPv4 优先 | 优先从 IPv4 地址中随机选择，如果没有 IPv4 地址则使用 IPv6 |
| `ipv6`   | IPv6 优先 | 优先从 IPv6 地址中随机选择，如果没有 IPv6 地址则使用 IPv4 |
| `random` | 随机选择  | 从所有 IPv4 和 IPv6 地址中随机选择                        |

**默认值**: `ipv4`

## 工作原理

### 1. DNS 解析阶段

当启用 `--upstream-resolve-ips` 参数时，代理服务器会使用 DoH (DNS over HTTPS)
解析目标域名，并获取所有可用的 IP 地址。

示例：

```bash
# 解析 example.com 可能返回：
- IPv4: 93.184.216.34:443, 93.184.216.35:443
- IPv6: 2606:2800:220:1:248:1893:25c8:1946:443, 2606:2800:220:1:248:1893:25c8:1947:443
```

### 2. IP 地址分离

代理服务器会将解析出的地址按照 IP 版本分为两组：

- **IPv4 组**: 所有 IPv4 地址
- **IPv6 组**: 所有 IPv6 地址

### 3. 地址选择策略

根据 `--ip-priority` 参数的设置，从对应的地址组中**随机选择**一个地址进行连接：

#### ipv4 模式

```go
if ipPriority == "ipv6" {
    if len(ipv6Addrs) > 0 {
        candidateAddrs = ipv6Addrs  // 优先使用 IPv6
    } else {
        candidateAddrs = ipv4Addrs  // 回退到 IPv4
    }
}
if ipPriority == "ipv4" {
    if len(ipv4Addrs) > 0 {
        candidateAddrs = ipv4Addrs  // 优先使用 IPv4
    } else {
        candidateAddrs = ipv6Addrs  // 回退到 IPv6

    }
}
```

#### ipv6 模式

```go
if ipPriority == "ipv6" {
    if len(ipv6Addrs) > 0 {
        candidateAddrs = ipv6Addrs  // 优先使用 IPv6
    } else {
        candidateAddrs = ipv4Addrs  // 回退到 IPv4
    }
}
if ipPriority == "ipv4" {
    if len(ipv4Addrs) > 0 {
        candidateAddrs = ipv4Addrs  // 优先使用 IPv4
    } else {
        candidateAddrs = ipv6Addrs  // 回退到 IPv6

    }
}
```

#### random 模式

```go
if ipPriority == "random" {
    candidateAddrs = addrs  // 使用所有地址
}
```

#### 随机选择

```go
// 从候选地址中随机选择一个
candidateAddrs = shuffle(candidateAddrs)
selectedAddr := candidateAddrs[rand.Intn(len(candidateAddrs))]
```

## 使用示例

### 基本用法

```bash
# IPv4 优先（默认）
./http-proxy-server --ip-priority ipv4

# IPv6 优先
./http-proxy-server --ip-priority ipv6

# 随机选择
./http-proxy-server --ip-priority random
```

### 与其他参数配合使用

```bash
# 启用上游 IP 解析并优先使用 IPv6
./http-proxy-server \
  --upstream-resolve-ips \
  --ip-priority ipv6 \
  --proxy socks5://127.0.0.1:1080

# 使用 DoH 服务器进行 DNS 解析，优先使用 IPv4
./http-proxy-server \
  --upstream-resolve-ips \
  --ip-priority ipv4 \
  --doh-url https://1.1.1.1/dns-query \
  --doh-ip 1.1.1.1
```

## 实际应用场景

### 场景 1: IPv6 网络环境

在纯 IPv6 或 IPv6 优先的网络环境中使用：

```bash
./http-proxy-server \
  --ip-priority ipv6 \
  --upstream-resolve-ips
```

**优势**:

- 避免 IPv4 连接的超时或失败
- 充分利用 IPv6 的高速连接
- 减少 NAT 转换延迟

### 场景 2: IPv4 兼容性优先

在需要确保与旧系统兼容的环境中使用：

```bash
./http-proxy-server \
  --ip-priority ipv4 \
  --upstream-resolve-ips
```

**优势**:

- 确保与仅支持 IPv4 的服务的兼容性
- 避免某些防火墙对 IPv6 的限制
- 更稳定的连接体验

### 场景 3: 负载均衡和容错

使用随机模式实现负载分散：

```bash
./http-proxy-server \
  --ip-priority random \
  --upstream-resolve-ips
```

**优势**:

- 自动分散流量到不同 IP
- 提高整体连接成功率
- 避免单个 IP 过载

## 技术细节

### 修改的函数

`--ip-priority` 参数影响以下模块的轮询函数：

| 模块    | 函数                                          | 文件                                     |
| ------- | --------------------------------------------- | ---------------------------------------- |
| auth    | `resolveTargetAddressForAuthWithRoundRobin`   | `auth/auth.go`                           |
| connect | `resolveTargetAddressForHttpWithRoundRobin`   | `connect/resolvetargetaddressforhttp.go` |
| cmd     | `resolveTargetAddressWithRoundRobin`          | `cmd/resolvetargetaddress.go`            |
| simple  | `resolveTargetAddressForSimpleWithRoundRobin` | `simple/simple.go`                       |
| http    | `resolveTargetAddressForAuthWithRoundRobin`   | `http/http.go`                           |

### 与 `--upstream-resolve-ips` 的关系

`--ip-priority` 参数需要与 `--upstream-resolve-ips` 参数配合使用才能生效：

- `--upstream-resolve-ips`: 启用上游 IP 解析功能
- `--ip-priority`: 控制如何从解析的 IP 中选择

如果不启用 `--upstream-resolve-ips`，则 `--ip-priority`
参数不会生效，代理将使用系统默认的 DNS 解析。

### 随机性说明

所有选择都使用 **真正的随机算法** (`rand.Intn`)，而不是确定性哈希。这意味着：

- 同一个域名在不同连接中可能选择不同的 IP
- 实现了真正的负载分散
- 相同域名不会总是选择相同的 IP

### 回退机制

所有模式都包含智能回退机制：

- **ipv4 模式**: 如果没有 IPv4 地址，自动使用 IPv6
- **ipv6 模式**: 如果没有 IPv6 地址，自动使用 IPv4
- **random 模式**: 使用所有可用地址

这确保了即使在特定 IP 版本不可用的情况下，代理仍能正常工作。

## 日志输出

启用 `--ip-priority` 后，代理会输出详细的日志信息：

```
IPv4 priority: selecting from 2 IPv4 addresses
RoundRobin (priority=ipv4) selected address 93.184.216.34:443 from [93.184.216.34:443 93.184.216.35:443 2606:2800:220:1:248:1893:25c8:1946:443] for target example.com:443
```

日志包含以下信息：

- 使用的优先级策略
- 候选地址的数量
- 最终选择的地址
- 所有可用的地址列表

## 常见问题

### Q1: 为什么默认使用 IPv4 优先？

**A**: IPv4 仍然是最广泛支持的 IP 版本，大多数服务都提供 IPv4 连接。使用 IPv4
作为默认值可以确保最大的兼容性。

### Q2: 如何判断应该使用哪个优先级？

**A**:

- 使用 `ipv6`: 如果你的网络环境支持 IPv6 且目标服务有 IPv6 地址
- 使用 `ipv4`: 如果需要确保兼容性或网络主要使用 IPv4
- 使用 `random`: 如果需要负载均衡或不确定哪个 IP 版本更好

### Q3: IPv6 连接失败会发生什么？

**A**: 系统会自动回退到 IPv4 连接（如果 IPv4 可用），确保代理服务的可用性。

### Q4: 这个参数是否影响所有类型的代理？

**A**: 是的，`--ip-priority` 影响所有代理类型：

- SOCKS5 代理
- HTTP/HTTPS 代理
- WebSocket 代理
- 认证代理

### Q5: 如何验证当前使用的 IP 版本？

**A**: 查看代理日志，会显示选择的地址。IPv6 地址通常包含冒号（`:`），IPv4
地址使用点号（`.`）。

## 性能影响

### 优势

- **减少连接延迟**: 优先使用更快的 IP 版本
- **提高连接成功率**: 智能回退机制
- **负载均衡**: 分散流量到多个 IP

### 注意事项

- **DNS 解析延迟**: 首次连接可能需要额外的 DNS 解析时间
- **地址缓存**: DNS 结果会被缓存，减少后续解析延迟
- **网络依赖**: 需要稳定的网络连接进行 DNS 查询

## 最佳实践

1. **测试环境**: 先在测试环境中验证不同优先级策略的效果
2. **监控日志**: 观察代理日志，了解实际使用的 IP 版本
3. **灵活调整**: 根据实际网络环境调整优先级策略
4. **配合 DoH**: 建议与 `--doh-url` 参数配合使用，提高 DNS 解析质量

## 版本历史

- **v1.0**: 初始版本，支持 IPv4/IPv6/random 三种策略
- **v0.x**: 之前版本使用哈希算法，相同域名总是选择相同 IP

## 相关参数

| 参数                     | 说明                 |
| ------------------------ | -------------------- |
| `--upstream-resolve-ips` | 启用上游 IP 解析功能 |
| `--doh-url`              | 配置 DoH 服务器地址  |
| `--doh-ip`               | 配置 DoH 服务器 IP   |
| `--dot-url`              | 配置 DoT 服务器地址  |
| `--dot-ip`               | 配置 DoT 服务器 IP   |

## 参考资源

- [IPv6 介绍](https://en.wikipedia.org/wiki/IPv6)
- [DNS over HTTPS](https://en.wikipedia.org/wiki/DNS_over_HTTPS)
- [IPv4 vs IPv6 性能对比](https://www.cloudflare.com/learning/network-layer/what-is-ipv6/)
