package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var logLevel = LevelInfo
var logFlag = LogFlagPrefix | LogFlagColor | LogFlagTime

const LogFlagPrefix = 1 << 0
const LogFlagColor = 1 << 1
const LogFlagTime = 1 << 2
const LogFlagCaller = 1 << 3

var _levelColor = map[int]func(format string, a ...interface{}) string{
	LevelDebug:  color.HiBlueString,
	LevelInfo:   color.HiCyanString,
	LevelNotice: color.HiGreenString,
	LevelWarn:   color.HiYellowString,
	LevelError:  color.HiRedString,
}

func init() {
	if os.Getenv("DEBUG") == "1" {
		SelLevel(LevelDebug)
	}
}

func SelLevel(level int) {
	logLevel = level
}

func SelLevelString(level string) {
	logLevel = stringToLevel(level)
}

func SelFlag(flag int) {
	logFlag = flag
}

const LevelDebug = 0
const LevelInfo = 1
const LevelNotice = 2
const LevelWarn = 3
const LevelError = 4

func stringToLevel(str string) int {
	switch str {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "notice":
		return LevelNotice
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	}
	return logLevel
}

func Debug(format string, v ...interface{}) {
	logf(LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	logf(LevelInfo, format, v...)
}

func Notice(format string, v ...interface{}) {
	logf(LevelNotice, format, v...)
}

func Warn(format string, v ...interface{}) {
	logf(LevelWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	logf(LevelError, format, v...)
}

func logf(level int, format string, v ...interface{}) {
	if level < logLevel {
		return
	}
	out := strings.Builder{}

	if logFlag&LogFlagPrefix == LogFlagPrefix {
		switch level {
		case LevelDebug:
			out.WriteString("[DEBUG] ")
		case LevelInfo:
			out.WriteString("[INFO] ")
		case LevelNotice:
			out.WriteString("[NOTICE] ")
		case LevelWarn:
			out.WriteString("[WARN] ")
		case LevelError:
			out.WriteString("[ERROR] ")
		default:
			out.WriteString("[-] ")
		}
	}
	if logFlag&LogFlagTime == LogFlagTime {
		now := time.Now().Format("15:04:05.000")
		out.WriteString(now)
		out.WriteString(" ")
	}

	if logFlag&LogFlagCaller == LogFlagCaller {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		out.WriteString(filepath.Base(file))
		out.WriteString(":")
		out.WriteString(strconv.FormatInt(int64(line), 10))
		out.WriteString(" ")
	}
	out.WriteString(fmt.Sprintf(format, v...))
	output := out.String()
	if logFlag&LogFlagColor == LogFlagColor {
		if foo := _levelColor[level]; foo != nil {
			output = foo(out.String())
		}
	}
	fmt.Println(output)
}
