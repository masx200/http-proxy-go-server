package tests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	log.Println("测试开始")

	exitCode := m.Run()
	log.Println("测试结束")

	os.Exit(exitCode)
}
func TestRun(t *testing.T) {

	var testsfuncs = map[string]func(t *testing.T, logfilename string){
		"RunMainWebSocket": RunMainWebSocket,
		"RunMainDOH":       RunMainDOH,
		"RunMainDEFAULT":   RunMainDEFAULT,
	}

	for name, fn := range testsfuncs {
		log.Println("测试开始", name)
		t1 := time.Now()
		var millisecond = t1.Nanosecond() / 1e6

		now := time.Now().Format("2006_01_02_15_0_05")
		logfilename := name + "_" + now + "_" + fmt.Sprintf("%d", millisecond) + ".log"

		log.Println("日志文件", logfilename)
		t.Run(name, func(t *testing.T) {
			fn(t, logfilename)
		})
		log.Println("测试结束", name)
	}

}
