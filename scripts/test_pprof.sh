#!/bin/bash

# pprof 功能测试脚本
# 用于验证 pprof 功能是否正常工作

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "=========================================="
echo "pprof 功能测试脚本"
echo "=========================================="
echo ""

# 检查必要工具
check_command() {
    if ! command -v $1 &> /dev/null; then
        echo -e "${RED}错误: $1 未安装${NC}"
        return 1
    fi
    return 0
}

echo "检查必要工具..."
check_command curl || exit 1
check_command go || exit 1
echo -e "${GREEN}✓ 必要工具检查通过${NC}"
echo ""

# 配置
PPROF_ADDR="${PPROF_ADDR:-127.0.0.1:6060}"
PROXY_PORT="${PROXY_PORT:-8080}"

echo "测试配置:"
echo "  pprof 地址: $PPROF_ADDR"
echo "  代理端口: $PROXY_PORT"
echo ""

# 函数：测试 pprof 端点
test_endpoint() {
    local endpoint=$1
    local description=$2

    echo -n "测试 $description... "
    if curl -s "http://$PPROF_ADDR$endpoint" > /dev/null 2>&1; then
        echo -e "${GREEN}✓ 通过${NC}"
        return 0
    else
        echo -e "${RED}✗ 失败${NC}"
        return 1
    fi
}

# 函数：获取并分析 profile
test_profile() {
    local endpoint=$1
    local output_file=$2
    local description=$3

    echo -n "获取 $description... "
    if curl -s "http://$PPROF_ADDR$endpoint" > "$output_file" 2>/dev/null; then
        echo -e "${GREEN}✓ 成功${NC}"

        # 检查文件大小
        if [ -s "$output_file" ]; then
            size=$(du -h "$output_file" | cut -f1)
            echo "  文件大小: $size"
            return 0
        else
            echo -e "${YELLOW}  警告: 文件为空${NC}"
            return 1
        fi
    else
        echo -e "${RED}✗ 失败${NC}"
        return 1
    fi
}

# 检查 pprof 是否可访问
echo "1. 检查 pprof 服务是否运行..."
if curl -s "http://$PPROF_ADDR/debug/pprof/" > /dev/null 2>&1; then
    echo -e "${GREEN}✓ pprof 服务正在运行${NC}"
else
    echo -e "${RED}✗ pprof 服务无法访问${NC}"
    echo ""
    echo "请确保代理服务器已启用 pprof 功能："
    echo "  go run ./cmd/ -enable-pprof"
    exit 1
fi
echo ""

# 测试各个端点
echo "2. 测试 pprof 端点..."
test_endpoint "/debug/pprof/" "pprof 首页"
test_endpoint "/debug/pprof/heap" "堆内存 profile"
test_endpoint "/debug/pprof/goroutine" "Goroutine profile"
test_endpoint "/debug/pprof/threadcreate" "线程创建 profile"
test_endpoint "/debug/pprof/block" "阻塞操作 profile"
test_endpoint "/debug/pprof/mutex" "互斥锁 profile"
echo ""

# 创建临时目录保存 profile
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

echo "3. 采集并分析 profile..."
echo "保存目录: $TEMP_DIR"

# 测试内存 profile
test_profile "/debug/pprof/heap" "$TEMP_DIR/heap.prof" "内存 profile"
if [ $? -eq 0 ]; then
    echo "  分析内存 profile..."
    top_output=$(go tool pprof -top "$TEMP_DIR/heap.prof" 2>/dev/null | head -10)
    if [ $? -eq 0 ]; then
        echo "  内存占用排名:"
        echo "$top_output" | sed 's/^/    /'
    fi
fi
echo ""

# 测试 goroutine profile
test_profile "/debug/pprof/goroutine" "$TEMP_DIR/goroutine.prof" "Goroutine profile"
if [ $? -eq 0 ]; then
    goroutine_count=$(curl -s "http://$PPROF_ADDR/debug/pprof/goroutine?debug=1" | grep -c "^goroutine" || echo "N/A")
    echo "  当前 Goroutine 数量: $goroutine_count"

    if [ "$goroutine_count" != "N/A" ] && [ $goroutine_count -gt 1000 ]; then
        echo -e "${YELLOW}  警告: Goroutine 数量过多 ($goroutine_count)${NC}"
    fi
fi
echo ""

# 测试 CPU profile（需要等待）
echo "4. 测试 CPU profile 采集（5秒）..."
echo -n "  采集中... "
if curl -s "http://$PPROF_ADDR/debug/pprof/profile?seconds=5" > "$TEMP_DIR/cpu.prof" 2>/dev/null; then
    echo -e "${GREEN}✓ 完成${NC}"

    if [ -s "$TEMP_DIR/cpu.prof" ]; then
        echo "  分析 CPU profile..."
        top_output=$(go tool pprof -top "$TEMP_DIR/cpu.prof" 2>/dev/null | head -10)
        if [ $? -eq 0 ]; then
            echo "  CPU 使用排名:"
            echo "$top_output" | sed 's/^/    /'
        fi
    fi
else
    echo -e "${RED}✗ 失败${NC}"
fi
echo ""

# 性能检查建议
echo "5. 性能检查建议..."

# 检查内存使用
mem_info=$(curl -s "http://$PPROF_ADDR/debug/pprof/heap" | go tool pprof -top - 2>/dev/null | head -5)
if [ $? -eq 0 ]; then
    echo "  内存使用情况:"
    echo "$mem_info" | sed 's/^/    /'
fi
echo ""

# 检查 goroutine 数量
goroutine_debug=$(curl -s "http://$PPROF_ADDR/debug/pprof/goroutine?debug=1" 2>/dev/null)
if [ $? -eq 0 ]; then
    # 统计不同状态的 goroutine
    running=$(echo "$goroutine_debug" | grep -c "running" || echo 0)
    blocked=$(echo "$goroutine_debug" | grep -c "chan receive\|semacquire" || echo 0)

    echo "  Goroutine 状态统计:"
    echo "    运行中: $running"
    echo "    阻塞中: $blocked"

    if [ $blocked -gt 100 ]; then
        echo -e "${YELLOW}    警告: 大量 goroutine 处于阻塞状态${NC}"
    fi
fi
echo ""

# 总结
echo "=========================================="
echo "测试完成"
echo "=========================================="
echo ""
echo "快速分析命令:"
echo "  # 实时内存分析"
echo "  go tool pprof http://$PPROF_ADDR/debug/pprof/heap"
echo ""
echo "  # 查看 goroutine"
echo "  curl http://$PPROF_ADDR/debug/pprof/goroutine?debug=1"
echo ""
echo "  # 生成火焰图"
echo "  go tool pprof -http=:9999 http://$PPROF_ADDR/debug/pprof/heap"
echo ""
echo "详细使用指南: docs/PPROF_USAGE.md"

exit 0