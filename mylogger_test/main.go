package main

import (
	"github.com/c0ding/logger/mylogger"
	"time"
)

var log mylogger.Logger

// 测试自定义的测试库
func main() {
	//log = mylogger.NewLog("debug")
	log = mylogger.NewFileLogger("debug", "./", "demo.log", 10*1024)
	for {
		log.Dubug("ss")
		log.Info("in")
		log.Error("er %s %d", "ssss", 11)
		log.Fatal("fa")
		time.Sleep(time.Microsecond)
	}

}
