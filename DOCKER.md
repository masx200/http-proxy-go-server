# Docker 部署指南

## 📦 获取Docker镜像

本项目的Docker镜像自动构建并发布到GitHub Container Registry
(ghcr.io)，使用Git哈希值作为版本标签。

### 🏷️ 镜像标签说明

每个自动构建的镜像都会创建以下标签：

| 标签类型      | 说明                  | 示例              | 使用场景     |
| ------------- | --------------------- | ----------------- | ------------ |
| **Git短哈希** | 基于Git提交的短哈希值 | `cc8b437`         | 生产环境推荐 |
| **分支名**    | 基于分支名            | `main`            | 测试环境     |
| **完整哈希**  | 基于完整Git哈希       | `hash-cc8b437...` | 精确定位版本 |
| **时间戳**    | 构建时间戳            | `20241115_143022` | 临时测试     |
| **Latest**    | 最新发布版本          | `latest`          | 仅发布版本   |

## 🚀 快速开始

### 1. 基本HTTP代理

```bash
# 使用最新的主分支版本
docker run -d -p 8080:8080 --name http-proxy \
  ghcr.io/masx200/http-proxy-go-server:main

# 带基本认证
docker run -d -p 8080:8080 --name http-proxy \
  ghcr.io/masx200/http-proxy-go-server:cc8b437 \
  -hostname 0.0.0.0 -port 8080 \
  -username admin -password secret
```

### 2. 带DNS缓存和DoH的代理

```bash
# 创建数据目录
mkdir -p ./dns_cache

# 运行带完整功能的代理
docker run -d -p 8080:8080 --name http-proxy \
  -v $(pwd)/dns_cache:/app/cache \
  ghcr.io/masx200/http-proxy-go-server:cc8b437 \
  -hostname 0.0.0.0 -port 8080 \
  -username admin -password secret \
  -dohurl https://dns.google/dns-query \
  -dohip 8.8.8.8 \
  -dohalpn h2 \
  -cache-enabled true \
  -cache-file /app/cache/dns_cache.json \
  -cache-aof-enabled true \
  -cache-aof-file /app/cache/dns_cache.aof \
  -cache-aof-interval 1s
```

### 3. 使用配置文件

```bash
# 创建配置文件
cat > config.json << EOF
{
  "hostname": "0.0.0.0",
  "port": 8080,
  "username": "admin",
  "password": "secret",
  "dns_cache": {
    "enabled": true,
    "aof_enabled": true,
    "aof_interval": "1s"
  },
  "doh": [
    {
      "ip": "8.8.8.8",
      "alpn": "h2",
      "url": "https://dns.google/dns-query"
    }
  ]
}
EOF

# 运行容器
docker run -d -p 8080:8080 --name http-proxy \
  -v $(pwd)/config.json:/app/config.json \
  -v $(pwd)/dns_cache:/app/cache \
  ghcr.io/masx200/http-proxy-go-server:cc8b437 \
  -config /app/config.json
```

## ⚙️ 配置选项

### DNS缓存配置

```bash
# 禁用AOF持久化
-cache-aof-enabled false

# 自定义保存间隔
-cache-aof-interval 5s  # 每5秒增量保存
-cache-save-interval 60s  # 每60秒全量保存

# 自定义TTL
-cache-ttl 30m  # 缓存30分钟
```

### DoH配置

```bash
# 多个DoH服务器
-dohurl https://dns.google/dns-query \
-dohip 8.8.8.8 \
-dohalpn h2

-dohurl https://dns.alidns.com/dns-query \
-dohip 223.5.5.5 \
-dohalpn h3
```

### 上游代理配置

```bash
# WebSocket上游
-upstream-type websocket \
-upstream-address ws://127.0.0.1:1081 \
-upstream-username user \
-upstream-password pass

# SOCKS5上游
-upstream-type socks5 \
-upstream-address socks5://127.0.0.1:1080 \
-upstream-username user \
-upstream-password pass
```

## 🔍 验证部署

### 1. 检查容器状态

```bash
# 查看容器日志
docker logs http-proxy

# 检查容器状态
docker ps | grep http-proxy
```

### 2. 测试代理功能

```bash
# 基本HTTP测试
curl -x http://localhost:8080 http://httpbin.org/ip

# HTTPS测试
curl -x http://localhost:8080 https://httpbin.org/ip

# 带认证测试
curl -x admin:secret@localhost:8080 http://httpbin.org/ip
```

### 3. 验证DNS缓存

```bash
# 进入容器查看缓存文件
docker exec http-proxy ls -la /app/cache/
docker exec http-proxy cat /app/cache/dns_cache.aof
```

## 📊 监控和维护

### 查看日志

```bash
# 实时查看日志
docker logs -f http-proxy

# 查看DNS缓存相关日志
docker logs http-proxy | grep "dns cache"
```

### 数据备份

```bash
# 备份DNS缓存数据
docker cp http-proxy:/app/cache ./backup_cache_$(date +%Y%m%d)

# 恢复缓存数据
docker cp ./backup_cache_20241115 http-proxy:/app/cache
```

### 性能监控

```bash
# 查看容器资源使用情况
docker stats http-proxy

# 查看容器详细信息
docker inspect http-proxy
```

## 🔧 高级配置

### 1. 多架构支持

镜像支持 `linux/amd64` 和 `linux/arm64` 架构：

```bash
# 拉取特定架构的镜像
docker pull --platform linux/amd64 ghcr.io/masx200/http-proxy-go-server:cc8b437
docker pull --platform linux/arm64 ghcr.io/masx200/http-proxy-go-server:cc8b437
```

### 2. 自定义网络

```bash
# 创建自定义网络
docker network create proxy-network

# 运行容器在自定义网络中
docker run -d --network proxy-network --name http-proxy \
  -p 8080:8080 \
  ghcr.io/masx200/http-proxy-go-server:cc8b437
```

### 3. 环境变量配置

```bash
# 使用环境变量设置默认值
docker run -d -p 8080:8080 --name http-proxy \
  -e PROXY_USERNAME=admin \
  -e PROXY_PASSWORD=secret \
  -e DOH_URL=https://dns.google/dns-query \
  ghcr.io/masx200/http-proxy-go-server:cc8b437
```

## 🛠️ 故障排除

### 常见问题

1. **容器启动失败**
   ```bash
   # 检查容器日志
   docker logs http-proxy

   # 检查端口是否被占用
   netstat -tlnp | grep :8080
   ```

2. **DNS缓存不工作**
   ```bash
   # 检查缓存目录权限
   docker exec http-proxy ls -la /app/cache/

   # 重新创建缓存目录
   docker exec http-proxy rm -rf /app/cache/*
   ```

3. **DoH连接失败**
   ```bash
   # 检查网络连接
   docker exec http-proxy nslookup google.com

   # 测试DoH服务器连通性
   docker exec http-proxy curl -v https://dns.google/dns-query
   ```

### 获取帮助

如果遇到问题，请：

1. 查看 [GitHub Issues](https://github.com/masx200/http-proxy-go-server/issues)
2. 检查最新的[文档](https://github.com/masx200/http-proxy-go-server)
3. 提交新的Issue并包含详细的错误信息和日志
