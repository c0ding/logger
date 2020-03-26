package mylogger

import (
	"fmt"
	"time"
)

// 往终端上写日志 相关内容

// Logger is 自定义日志库
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 是Logger的 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{Level: level}
}

// enable 打印的日志级别 要 大于等于 创建log对象时的级别
func (c ConsoleLogger) enable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		nowStr := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		lvString := logLvString(lv)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", nowStr, lvString, fileName, funcName, lineNo, msg)
	}
}

// Dubug 调试日志
func (c ConsoleLogger) Dubug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)

}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Waring(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
