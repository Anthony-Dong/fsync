package main

import . "github.com/anthony-dong/fsync/logs"

func main() {
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
