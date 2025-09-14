package tests

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ProcessManager 进程管理器
type ProcessManager struct {
	processes []*exec.Cmd
	mutex     sync.Mutex
	logFile   *os.File
	logWriter *bufio.Writer
}

// NewProcessManager 创建新的进程管理器
func NewProcessManager() *ProcessManager {
	pm := &ProcessManager{
		processes: make([]*exec.Cmd, 0),
	}

	// 初始化日志文件
	pm.initLogFile()

	return pm
}
func (pm *ProcessManager) Command(name string, arg ...string) *exec.Cmd {

	cmd := exec.Command(name, arg...)

	pm.AddProcess(cmd)

	log.Printf("执行命令: %s %s", name, strings.Join(arg, " "))
	pm.LogCommand(cmd, "执行命令")
	return cmd
}

// initLogFile 初始化日志文件
func (pm *ProcessManager) initLogFile() {
	// 打开或创建日志文件
	file, err := os.OpenFile("command_execution_log.md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 如果无法打开日志文件，记录到标准错误输出
		fmt.Fprintf(os.Stderr, "无法打开命令日志文件: %v\n", err)
		return
	}

	pm.logFile = file
	pm.logWriter = bufio.NewWriter(file)

	// 写入日志文件头（如果是新文件）
	if _, err := file.Stat(); err == nil {
		// 检查文件是否为空
		info, _ := file.Stat()
		if info.Size() == 0 {
			pm.writeLog("=== 命令执行日志文件 ===\n")
			pm.writeLog(fmt.Sprintf("创建时间: %s\n", time.Now().Format("2006-01-02 15:04:05")))
			pm.writeLog("\n")
		}
	}
}

// writeLog 写入日志
func (pm *ProcessManager) writeLog(log string) {
	if pm.logWriter != nil {
		pm.logWriter.WriteString(log)
		pm.logWriter.Flush()
	}
}

// LogCommand 记录命令执行
func (pm *ProcessManager) LogCommand(cmd *exec.Cmd, cmdType string) error {
	if pm.logWriter == nil {
		return nil // 日志文件不可用，跳过记录
	}

	// 构建命令字符串
	var cmdStr string
	if cmd.Path != "" {
		cmdStr = cmd.Path
		for _, arg := range cmd.Args {
			if arg != cmd.Path {
				cmdStr += " " + arg
			}
		}
	}

	// 记录命令信息
	logEntry := fmt.Sprintf("[%s] [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05.000"),
		cmdType,
		cmdStr)

	pm.writeLog("开始运行命令...\n" + logEntry + "\n\n")
	return nil
}

// LogCommandResult 记录命令执行结果
func (pm *ProcessManager) LogCommandResult(cmd *exec.Cmd, err error, output string) {
	if pm.logWriter == nil {
		return // 日志文件不可用，跳过记录
	}

	// 确定执行结果
	result := "成功"
	if err != nil {
		result = "失败"
	}

	// 获取进程PID
	pid := "N/A"
	if cmd.Process != nil {
		pid = strconv.Itoa(cmd.Process.Pid)
	}

	// 记录执行结果
	logEntry := fmt.Sprintf("执行结果: %s\n", result)
	logEntry += fmt.Sprintf("进程PID: %s\n", pid)

	// 记录执行时间
	if cmd.ProcessState != nil {
		duration := cmd.ProcessState.SystemTime() + cmd.ProcessState.UserTime()
		logEntry += fmt.Sprintf("执行时间: %v\n", duration)
	}

	// 记录输出（如果存在）
	if output != "" {
		logEntry += fmt.Sprintf("输出: %s\n", strings.TrimSpace(output))
	}

	// 记录错误（如果存在）
	if err != nil {
		logEntry += fmt.Sprintf("错误: %s\n", err.Error())
	}

	logEntry += "---\n"
	pm.writeLog("```\n" + logEntry + "```\n")
}

// AddProcess 添加进程到管理器
func (pm *ProcessManager) AddProcess(cmd *exec.Cmd) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.processes = append(pm.processes, cmd)
}

// CleanupAll 清理所有进程
func (pm *ProcessManager) CleanupAll() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			// Windows系统下使用更强制的方式终止进程
			if runtime.GOOS == "windows" {
				// 在Windows上，我们需要终止整个进程树
				cmd.Process.Kill()
				// 等待进程退出
				cmd.Wait()

				// 尝试查找并终止子进程
				pm.killChildProcesses(cmd.Process.Pid)
			} else {
				// Unix系统下使用进程组
				cmd.Process.Kill()
				cmd.Wait()
			}
		}
	}
	pm.processes = make([]*exec.Cmd, 0)
}

// killChildProcesses 在Windows上终止子进程
func (pm *ProcessManager) killChildProcesses(parentPid int) {
	// 在Windows上使用taskkill命令终止进程树
	killCmd := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(parentPid))
	killCmd.Run() // 忽略错误，因为进程可能已经退出
}

// GetPIDs 获取所有进程的PID
func (pm *ProcessManager) GetPIDs() []string {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	var pids []string
	for _, cmd := range pm.processes {
		if cmd.Process != nil {
			pids = append(pids, strconv.Itoa(cmd.Process.Pid))
		}
	}
	return pids
}

// Close 关闭日志文件
func (pm *ProcessManager) Close() {
	if pm.logFile != nil {
		pm.logWriter.Flush()
		pm.logFile.Close()
	}
}
