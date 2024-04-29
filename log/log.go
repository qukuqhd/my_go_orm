package log

import (
	"log"
	"os"
	"sync"
)

// 不同的日志分级定义
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile) //错误级别 红色的字符串
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile) //普通信息 蓝色的字符串
	loggers  = []*log.Logger{errorLog, infoLog}                                            //不同级别的日志的切片
	mu       sync.Mutex                                                                    //日志互斥锁
)

// 日志调用方法
var (
	Error  = errorLog.Println //换行错误日志打印
	Errorf = errorLog.Printf  //格式化错误日志打印
	Info   = infoLog.Println  //换行普通日志打印
	Infof  = infoLog.Printf   //格式化普通日志打印
)
