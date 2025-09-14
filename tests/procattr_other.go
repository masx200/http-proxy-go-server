//go:build !windows
// +build !windows

package tests

import "syscall"

// NewSysProcAttr 在非 Windows 平台返回 nil，表示不需要额外属性
func NewSysProcAttr() *syscall.SysProcAttr {
	return nil
}
