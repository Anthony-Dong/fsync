package logs

import "testing"

func Test_SelLevel(t *testing.T) {
	SelLevel(LevelDebug)
	doPrint()

	SelLevel(LevelInfo)
	doPrint()

	SelLevel(LevelNotice)
	doPrint()

	SelLevel(LevelWarn)
	doPrint()

	SelLevel(LevelError)
	doPrint()
}
func doPrint() {
	Debug("123456")
	Info("hello")
	Notice("china")
	Warn("欢迎")
	Error("中国")
}
