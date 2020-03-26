package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string // 文件保存的路径
	fileName    string // 文件保存的文件名
	maxFileSize int64  // 最大文件的存储能力
	fileObj     *os.File
	errfileObj  *os.File
	//errFileName string //对于err级别往下的日志，在单独记录
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		filePath:    fp,
		fileName:    fn,
		Level:       logLevel,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

// initFile 按照文件路径和文件名把文件打开
func (f *FileLogger) initFile() error {
	// 记录普通日志的文件
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file faild, err:%v \n", err)
		return err
	}

	// 记录错误级别日志的文件
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file faild, err:%v \n", err)
		return err
	}

	f.fileObj = fileObj
	f.errfileObj = errfileObj
	return nil
}

func (f *FileLogger) checkSize(file *os.File) bool {
	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("file stat faild %v", err)
		return false
	}
	return stat.Size() >= f.maxFileSize

}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errfileObj.Close()
}

// enable 打印的日志级别 要 大于等于 创建log对象时的级别
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}
func (f *FileLogger) split(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err2 := file.Stat()
	if err2 != nil {
		fmt.Println("file stat faild")
		return nil, err2
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	logNameNew := fmt.Sprintf("%s.bak%s", logName, nowStr)

	// 需要切割
	// 1，关闭当前文件
	file.Close()
	// 2，做备份，根据时间戳改名  xx.log ->  xx.log.bak20060102150405000

	os.Rename(logName, logNameNew)

	// 3，打开一个新文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file faild %v/n", err)
		return nil, err
	}
	// 4，将打开的新文件对象 赋值给f.fileObj
	return fileObj, nil
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		nowStr := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		lvString := logLvString(lv)
		if f.checkSize(f.fileObj) {
			file, err := f.split(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = file
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n", nowStr, lvString, fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			// 如果记录的日志级别大于ERROR，再单独记录在 err文件
			if f.checkSize(f.errfileObj) {
				file, err := f.split(f.errfileObj)
				if err != nil {
					return
				}
				f.errfileObj = file
			}
			fmt.Fprintf(f.errfileObj, "[%s] [%s] [%s:%s:%d] %s \n", nowStr, lvString, fileName, funcName, lineNo, msg)
		}
	}
}

// Dubug 调试日志
func (f *FileLogger) Dubug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Waring(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
