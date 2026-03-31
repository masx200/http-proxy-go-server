#!/bin/bash

# HTTP Proxy Go Server 内存监控脚本
# 用法: ./scripts/monitor_memory.sh [进程名或PID]

set -e

# 默认进程名
PROCESS_NAME="${1:-http-proxy-go-server}"
OUTPUT_FILE="memory_usage_$(date +%Y%m%d_%H%M%S).csv"
LOG_FILE="monitor_$(date +%Y%m%d_%H%M%S).log"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "=========================================="
echo "HTTP Proxy Server 内存监控工具"
echo "=========================================="
echo "监控目标: $PROCESS_NAME"
echo "输出文件: $OUTPUT_FILE"
echo "日志文件: $LOG_FILE"
echo "=========================================="

# 查找进程
if [[ "$PROCESS_NAME" =~ ^[0-9]+$ ]]; then
    # 如果是PID
    PID=$PROCESS_NAME
    PROCESS_NAME=$(ps -p $PID -o comm= 2>/dev/null || echo "unknown")
else
    # 如果是进程名
    PID=$(pgrep -f "$PROCESS_NAME" | head -1)
fi

if [ -z "$PID" ]; then
    echo -e "${RED}错误: 未找到进程 '$PROCESS_NAME'${NC}"
    echo "请确保代理服务器正在运行"
    echo ""
    echo "提示:"
    echo "  1. 检查进程是否运行: ps aux | grep $PROCESS_NAME"
    echo "  2. 启动代理服务器: go run ./cmd/"
    echo "  3. 使用进程名监控: ./scripts/monitor_memory.sh http-proxy-go-server"
    echo "  4. 使用PID监控: ./scripts/monitor_memory.sh <PID>"
    exit 1
fi

echo -e "${GREEN}找到进程: PID=$PID, 名称=$PROCESS_NAME${NC}"
echo ""

# 检查系统命令
check_command() {
    if ! command -v $1 &> /dev/null; then
        echo -e "${YELLOW}警告: $1 未安装，某些功能可能不可用${NC}"
        return 1
    fi
    return 0
}

# 创建CSV头
echo "Timestamp,RSS_MB,VSZ_MB,CPU_Percent,NumConnections,NumGoroutines" > "$OUTPUT_FILE"

# 初始统计
INITIAL_STATS=$(ps -p $PID -o rss,vsz,%cpu --no-headers | awk '{printf "%.2f,%.2f,%.1f", $1/1024, $2/1024, $3}')
INITIAL_RSS=$(echo $INITIAL_STATS | cut -d',' -f1)

echo "开始监控... (按 Ctrl+C 停止)"
echo ""
printf "%-20s | %-10s | %-10s | %-8s | %-15s | %-15s\n" \
    "Time" "RSS(MB)" "VSZ(MB)" "CPU%" "Connections" "Goroutines"
printf "%s\n" "----------------------------------------------------------------------------------------"

# 监控循环
iteration=0
max_iterations=10000  # 防止无限循环
check_goroutine_every=10  # 每10次检查goroutine

while kill -0 $PID 2>/dev/null && [ $iteration -lt $max_iterations ]; do
    # 获取内存和CPU统计
    STATS=$(ps -p $PID -o rss,vsz,%cpu --no-headers 2>/dev/null)
    if [ $? -ne 0 ]; then
        echo -e "${RED}进程 $PID 已终止${NC}"
        break
    fi

    RSS_MB=$(echo $STATS | awk '{printf "%.2f", $1/1024}')
    VSZ_MB=$(echo $STATS | awk '{printf "%.2f", $2/1024}')
    CPU_PER=$(echo $STATS | awk '{printf "%.1f", $3}')

    # 检查网络连接数
    NUM_CONN=0
    if check_command netstat; then
        NUM_CONN=$(netstat -an 2>/dev/null | grep -c ":8080" || echo "0")
    elif check_command ss; then
        NUM_CONN=$(ss -ant 2>/dev/null | grep -c ":8080" || echo "0")
    fi

    # 检查goroutine数量（通过pprof）
    NUM_GOROUTINES="N/A"
    if [ $((iteration % $check_goroutine_every)) -eq 0 ]; then
        if check_command curl; then
            GOROUTINE_DATA=$(curl -s http://localhost:6060/debug/pprof/goroutine?debug=1 2>/dev/null)
            if [ $? -eq 0 ]; then
                NUM_GOROUTINES=$(echo "$GOROUTINE_DATA" | grep "^goroutine" | wc -l | tr -d ' ')
            fi
        fi
    fi

    # 获取当前时间
    TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')

    # 写入CSV
    echo "$TIMESTAMP,$RSS_MB,$VSZ_MB,$CPU_PER,$NUM_CONN,$NUM_GOROUTINES" >> "$OUTPUT_FILE"

    # 显示到终端
    printf "%-20s | %-10s | %-10s | %-8s | %-15s | %-15s\n" \
        "$TIMESTAMP" "$RSS_MB" "$VSZ_MB" "$CPU_PER%" "$NUM_CONN" "$NUM_GOROUTINES"

    # 检查内存是否超过阈值（相对于初始内存增长超过500MB）
    RSS_INCREASE=$(echo "$RSS_MB - $INITIAL_RSS" | bc)
    if [ $(echo "$RSS_INCREASE > 500" | bc) -eq 1 ]; then
        echo -e "${YELLOW}警告: 内存增长超过500MB (${RSS_INCREASE}MB)${NC}" | tee -a "$LOG_FILE"
    fi

    # 检查goroutine数量
    if [ "$NUM_GOROUTINES" != "N/A" ] && [ $NUM_GOROUTINES -gt 1000 ]; then
        echo -e "${YELLOW}警告: Goroutine数量过多 ($NUM_GOROUTINES)${NC}" | tee -a "$LOG_FILE"
    fi

    # 等待5秒
    sleep 5
    ((iteration++))
done

# 最终统计
echo ""
echo "=========================================="
echo "监控结束"
echo "=========================================="

if [ -f "$OUTPUT_FILE" ]; then
    echo "数据已保存到: $OUTPUT_FILE"
    echo ""

    # 生成简单的统计报告
    TOTAL_SAMPLES=$(wc -l < "$OUTPUT_FILE")
    AVG_RSS=$(awk -F',' 'NR>1 {sum+=$2; count++} END {if(count>0) print sum/count; else print 0}' "$OUTPUT_FILE")
    MAX_RSS=$(awk -F',' 'NR>1 {if($2>max) max=$2} END {print max}' "$OUTPUT_FILE")

    echo "统计信息:"
    echo "  总采样次数: $TOTAL_SAMPLES"
    echo "  平均内存: ${AVG_RSS} MB"
    echo "  最大内存: ${MAX_RSS} MB"
    echo ""
    echo "提示: 可以使用以下命令分析数据:"
    echo "  # 查看数据文件"
    echo "  cat $OUTPUT_FILE"
    echo ""
    echo "  # 在Excel/Numbers中打开CSV文件"
    echo "  # 或使用 gnuplot 生成图表"
    echo "  gnuplot -e \"set datafile separator ','; plot '$OUTPUT_FILE' using 2 with lines title 'RSS Memory'\""
fi

# 提供内存泄漏检测建议
echo ""
echo "内存泄漏检测建议:"
echo "  1. 如果内存持续增长且不释放，可能存在内存泄漏"
echo "  2. 如果RSS增长超过500MB，建议使用pprof进行详细分析"
echo "  3. 如果Goroutine数量持续增长，可能存在goroutine泄漏"
echo "  4. 运行以下命令进行详细分析:"
echo "     go tool pprof http://localhost:6060/debug/pprof/heap"
echo "     curl http://localhost:6060/debug/pprof/goroutine?debug=2"

exit 0