package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	logTest()
}

// 标准库中的日志库 测试，借鉴
func logTest() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 00644)
	if err != nil {
		fmt.Printf("open file faild,err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)

	for {
		log.Println("测试日志")
		time.Sleep(time.Second * 2)
	}
}
