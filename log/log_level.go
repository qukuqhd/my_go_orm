package log

import (
	"io/ioutil"
	"os"
)

// 日志级别枚举
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// 设置日志级别
func SetLevel(level int) {
	//使用互斥锁保证并发安全
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
