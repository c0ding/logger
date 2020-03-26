package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 自定义日志库

type LogLevel uint16

const (
	UNKOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// parseLogLevel 解析str 转成 LogLevel
func parseLogLevel(levelStr string) (LogLevel, error) {

	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKOWN, err
	}
}

func logLvString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.caller() faild ")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	lineNo = line
	funcName = strings.Split(funcName, ".")[1]
	return
}
