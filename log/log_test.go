package log

import (
	"os"
	"testing"
)

// 测试设置日志等级
func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set error log level")
	}
	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set info log level")
	}
}
