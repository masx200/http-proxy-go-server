//go:build windows
// +build windows

package tests   // 换成你实际包名

import "syscall"

// NewSysProcAttr 返回 Windows 下需要的新进程组标志
func NewSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}
