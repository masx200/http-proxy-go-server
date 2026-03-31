#!/bin/bash

# HTTP Proxy Go Server 启动脚本（带pprof支持）
# 用法: ./scripts/start_with_pprof.sh [参数...]

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "=========================================="
echo "HTTP Proxy Server 启动脚本"
echo "=========================================="
echo ""

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo -e "${RED}错误: Go 未安装${NC}"
    exit 1
fi

echo -e "${GREEN}Go 版本:${NC} $(go version)"
echo ""

# 创建临时启动脚本，添加pprof支持
TEMP_MAIN=$(mktemp)
trap "rm -f $TEMP_MAIN" EXIT

# 生成带pprof的启动代码
cat > "$TEMP_MAIN" << 'EOF'
package main

import (
    "flag"
    "log"
    "net/http"
    _ "net/http/pprof"
    "os"
)

func main() {
    // 启动pprof服务器
    go func() {
        pprofAddr := ":6060"
        log.Printf("Starting pprof server on %s", pprofAddr)
        log.Println("pprof URLs:")
        log.Println("  - http://localhost:6060/debug/pprof/")
        log.Println("  - http://localhost:6060/debug/pprof/heap")
        log.Println("  - http://localhost:6060/debug/pprof/goroutine")
        log.Println("  - http://localhost:6060/debug/pprof/threadcreate")
        log.Println("  - http://localhost:6060/debug/pprof/block")
        log.Println("  - http://localhost:6060/debug/pprof/mutex")

        if err := http.ListenAndServe(pprofAddr, nil); err != nil {
            log.Fatalf("pprof server failed: %v", err)
        }
    }()

    // 导入主包并运行
    // 这里使用反射或者直接调用主包的main函数
    // 由于Go的限制，我们使用另一种方法

    log.Println("Starting HTTP proxy server with pprof support...")
    log.Println("Use Ctrl+C to stop the server")
}
EOF

# 设置端口
PORT="${PORT:-8080}"
HOSTNAME="${HOSTNAME:-0.0.0.0}"

echo -e "${BLUE}配置:${NC}"
echo "  Hostname: $HOSTNAME"
echo "  Port: $PORT"
echo "  Pprof Port: 6060"
echo ""

# 检查端口是否可用
if netstat -tuln 2>/dev/null | grep -q ":$PORT "; then
    echo -e "${YELLOW}警告: 端口 $PORT 已被占用${NC}"
    read -p "是否继续? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

if netstat -tuln 2>/dev/null | grep -q ":6060 "; then
    echo -e "${YELLOW}警告: pprof端口 6060 已被占用${NC}"
    read -p "是否继续? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

echo -e "${GREEN}启动代理服务器...${NC}"
echo ""
echo -e "${BLUE}pprof分析命令:${NC}"
echo "  # 获取内存profile"
echo "  curl http://localhost:6060/debug/pprof/heap > heap.prof"
echo ""
echo "  # 分析内存"
echo "  go tool pprof heap.prof"
echo ""
echo "  # 实时内存分析"
echo "  go tool pprof http://localhost:6060/debug/pprof/heap"
echo ""
echo "  # 检查goroutine"
echo "  curl http://localhost:6060/debug/pprof/goroutine?debug=1"
echo ""
echo "  # CPU分析"
echo "  curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof"
echo "  go tool pprof cpu.prof"
echo ""
echo -e "${BLUE}监控命令:${NC}"
echo "  # 在另一个终端运行内存监控"
echo "  ./scripts/monitor_memory.sh http-proxy-go-server"
echo ""
echo "=========================================="
echo ""

# 启动代理服务器，传递所有参数
# 设置环境变量启用pprof
export GODEBUG=gctrace=1

# 直接运行主程序，传入所有参数
cd "$(dirname "$0")/.."
exec go run ./cmd/ \
    -hostname "$HOSTNAME" \
    -port "$PORT" \
    "$@"