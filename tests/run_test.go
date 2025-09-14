package tests

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	log.Println("测试开始")

	exitCode := m.Run()
	log.Println("测试结束")

	os.Exit(exitCode)
}
func TestRun(t *testing.T) {
	log.Println("测试开始")
	RunMainWebSocket(t)
	log.Println("测试结束")
	log.Println("测试开始")
	RunMainDOH(t)
	log.Println("测试结束")
	log.Println("测试开始")
	RunMainDEFAULT(t)
	log.Println("测试结束")
}
