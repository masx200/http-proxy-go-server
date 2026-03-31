# HTTP Proxy Go Server 内存监控脚本 (PowerShell版本)
# 用法: .\scripts\monitor_memory.ps1 [进程名]

param(
    [string]$ProcessName = "http-proxy-go-server",
    [int]$IntervalSeconds = 5,
    [string]$OutputFile = "memory_usage_$(Get-Date -Format 'yyyyMMdd_HHmmss').csv"
)

# 错误处理
$ErrorActionPreference = "Continue"

# 颜色函数
function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White"
    )
    Write-Host $Message -ForegroundColor $Color
}

# 显示标题
Write-ColorOutput "==========================================" "Cyan"
Write-ColorOutput "HTTP Proxy Server 内存监控工具" "Cyan"
Write-ColorOutput "==========================================" "Cyan"
Write-Host "监控目标: $ProcessName"
Write-Host "输出文件: $OutputFile"
Write-Host "采样间隔: $IntervalSeconds 秒"
Write-ColorOutput "==========================================" "Cyan"
Write-Host ""

# 查找进程
$process = Get-Process -Name $ProcessName -ErrorAction SilentlyContinue

if (-not $process) {
    Write-ColorOutput "错误: 未找到进程 '$ProcessName'" "Red"
    Write-Host ""
    Write-Host "请确保代理服务器正在运行"
    Write-Host ""
    Write-Host "提示:"
    Write-Host "  1. 检查进程是否运行: Get-Process | Where-Object {$_.ProcessName -like '*http-proxy*'}"
    Write-Host "  2. 启动代理服务器: go run ./cmd/"
    Write-Host "  3. 使用进程名监控: .\scripts\monitor_memory.ps1 http-proxy-go-server"
    Write-Host "  4. 使用PID监控: .\scripts\monitor_memory.ps1 -ProcessName <PID>"
    exit 1
}

$pid = $process.Id
Write-ColorOutput "找到进程: PID=$pid, 名称=$($process.ProcessName)" "Green"
Write-Host ""

# 创建CSV头
$csvHeader = "Timestamp,WorkingSetMB,PrivateMemoryMB,VirtualMemoryMB,CPUPercent,StartTime"
$csvHeader | Out-File -FilePath $OutputFile -Encoding utf8

# 记录初始内存
$initialWS = $process.WorkingSet64 / 1MB
Write-Host "开始监控... (按 Ctrl+C 停止)"
Write-Host ""

# 显示表头
Write-Host ("{0,-20} | {1,-12} | {2,-12} | {3,-12} | {4,-8}" -f `
    "Time", "WS(MB)", "PM(MB)", "VM(MB)", "CPU%")
Write-Host ("-" * 85)

# 监控循环
$iteration = 0
$maxIterations = 10000  # 防止无限循环

try {
    while ($iteration -lt $maxIterations) {
        # 重新获取进程信息
        $process = Get-Process -Id $pid -ErrorAction SilentlyContinue

        if (-not $process) {
            Write-ColorOutput "进程 $pid 已终止" "Red"
            break
        }

        # 获取内存和CPU统计
        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $workingSetMB = [math]::Round($process.WorkingSet64 / 1MB, 2)
        $privateMemoryMB = [math]::Round($process.PrivateMemorySize64 / 1MB, 2)
        $virtualMemoryMB = [math]::Round($process.VirtualMemorySize64 / 1MB, 2)
        $cpuPercent = [math]::Round($process.CPU, 1)

        # 写入CSV
        $csvLine = "$timestamp,$workingSetMB,$privateMemoryMB,$virtualMemoryMB,$cpuPercent,$($process.StartTime)"
        $csvLine | Out-File -FilePath $OutputFile -Append -Encoding utf8

        # 显示到终端
        Write-Host ("{0,-20} | {1,-12} | {2,-12} | {3,-12} | {4,-8}" -f `
            $timestamp, $workingSetMB, $privateMemoryMB, $virtualMemoryMB, "$cpuPercent%")

        # 检查内存是否超过阈值（相对于初始内存增长超过500MB）
        $wsIncrease = $workingSetMB - $initialWS
        if ($wsIncrease -gt 500) {
            $warningMsg = "警告: 内存增长超过500MB ($wsIncrease MB)"
            Write-ColorOutput $warningMsg "Yellow"
        }

        # 检查是否有pprof端点可以获取goroutine信息
        if ($iteration % 10 -eq 0) {
            try {
                $response = Invoke-WebRequest -Uri "http://localhost:6060/debug/pprof/" -TimeoutSec 2 -ErrorAction SilentlyContinue
                if ($response.StatusCode -eq 200) {
                    Write-Host "  [pprof可用: http://localhost:6060/debug/pprof/]" -ForegroundColor Gray
                }
            } catch {
                # pprof不可用，忽略错误
            }
        }

        # 等待指定间隔
        Start-Sleep -Seconds $IntervalSeconds
        $iteration++
    }
}
finally {
    # 最终统计
    Write-Host ""
    Write-ColorOutput "==========================================" "Cyan"
    Write-ColorOutput "监控结束" "Cyan"
    Write-ColorOutput "==========================================" "Cyan"

    if (Test-Path $OutputFile) {
        Write-Host "数据已保存到: $OutputFile"
        Write-Host ""

        # 读取CSV文件并计算统计信息
        $data = Import-Csv $OutputFile
        $totalSamples = $data.Count

        if ($totalSamples -gt 1) {
            $avgWS = ($data | Measure-Object -Property WorkingSetMB -Average).Average
            $maxWS = ($data | Measure-Object -Property WorkingSetMB -Maximum).Maximum
            $avgPM = ($data | Measure-Object -Property PrivateMemoryMB -Average).Average
            $maxPM = ($data | Measure-Object -Property PrivateMemoryMB -Maximum).Maximum

            Write-Host "统计信息:"
            Write-Host "  总采样次数: $totalSamples"
            Write-Host ("  平均工作集内存: {0:N2} MB" -f $avgWS)
            Write-Host ("  最大工作集内存: {0:N2} MB" -f $maxWS)
            Write-Host ("  平均私有内存: {0:N2} MB" -f $avgPM)
            Write-Host ("  最大私有内存: {0:N2} MB" -f $maxPM)
            Write-Host ""
        }

        Write-Host "提示: 可以使用以下方法分析数据:"
        Write-Host "  1. 在Excel中打开CSV文件"
        Write-Host "  2. 使用PowerShell分析:"
        Write-Host "     `$data = Import-Csv '$OutputFile'"
        Write-Host "     `$data | Format-Table -AutoSize"
        Write-Host "  3. 生成图表 (需要安装ImportExcel模块):"
        Write-Host "     Install-Module -Name ImportExcel"
        Write-Host "     Import-Csv '$OutputFile' | Export-Excel -Path 'memory_chart.xlsx' -Show"
    }

    Write-Host ""
    Write-Host "内存泄漏检测建议:"
    Write-Host "  1. 如果内存持续增长且不释放，可能存在内存泄漏"
    Write-Host "  2. 如果工作集内存增长超过500MB，建议进行详细分析"
    Write-Host "  3. 运行以下命令进行详细分析:"
    Write-Host "     go tool pprof http://localhost:6060/debug/pprof/heap"
    Write-Host "     curl http://localhost:6060/debug/pprof/goroutine?debug=2"
}

# 退出
exit 0