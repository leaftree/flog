package flog

import (
	"encoding/json"
	"fmt"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

type Level int

const (
	INFO Level = iota
	WARN
	EROR
	maxFuncNameLength = 30
)

const (
	colorTerminalGreen  = "\x1b[1;32m"
	colorTerminalYellow = "\x1b[1;33m"
	colorTerminalRed    = "\x1b[1;31m"
	colorTerminalNormal = "\x1b[0m"

	colorHtmlGreen  = "<font color=green>"
	colorHtmlYellow = "<font color=yellow>"
	colorHtmlRed    = "<font color=red>"
	colorHtmlNormal = "</font>"
)

var LogLevelName map[Level]string

func init() {
	/*
		LogLevelName = map[Level]string{
			INFO: colorTerminalGreen + "INFO" + colorTerminalNormal,
			WARN: colorTerminalYellow + "WARN" + colorTerminalNormal,
			EROR: colorTerminalRed + "EROR" + colorTerminalNormal,
		}
	*/
	if isatty(syscall.Stdout) {
		LogLevelName = map[Level]string{
			INFO: colorTerminalGreen + "INFO" + colorTerminalNormal,
			WARN: colorTerminalYellow + "WARN" + colorTerminalNormal,
			EROR: colorTerminalRed + "EROR" + colorTerminalNormal,
		}
	} else {
		LogLevelName = map[Level]string{
			INFO: colorHtmlGreen + "INFO" + colorHtmlNormal,
			WARN: colorHtmlYellow + "WARN" + colorHtmlNormal,
			EROR: colorHtmlRed + "EROR" + colorHtmlNormal,
		}
	}
}

func isatty(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	return err == nil
}

func Info(args ...interface{}) {
	writeLog(INFO, fmt.Sprintln(args...))
}
func Infof(arg0 string, args ...interface{}) {
	writeLog(INFO, fmt.Sprintf(arg0, args...)+"\n")
}

func Warn(args ...interface{}) {
	writeLog(WARN, fmt.Sprintln(args...))
}
func Warnf(arg0 string, args ...interface{}) {
	writeLog(WARN, fmt.Sprintf(arg0, args...)+"\n")
}

func Error(args ...interface{}) {
	writeLog(EROR, fmt.Sprintln(args...))
}
func Errorf(arg0 string, args ...interface{}) {
	writeLog(EROR, fmt.Sprintf(arg0, args...)+"\n")
}

func JsonIndent(args ...interface{}) {
	data := ""
	for _, arg := range args {
		val, _ := json.MarshalIndent(arg, "", "\t")
		if data == "" {
			data = string(val)
		} else {
			data += "\n\n" + string(val)
		}
	}
	writeLog(INFO, data+"\n")
}

func Json(args ...interface{}) {
	data := ""
	for _, arg := range args {
		val, _ := json.Marshal(arg)
		if data == "" {
			data = string(val)
		} else {
			data += "\n\n" + string(val)
		}
	}
	writeLog(INFO, data+"\n")
}

func currentTimestamp() string {
	var (
		nTime       = time.Now()
		nYear       = nTime.Year()
		nMonth      = nTime.Month()
		nDay        = nTime.Day()
		nHour       = nTime.Hour()
		nMinute     = nTime.Minute()
		nSecond     = nTime.Second()
		nMicrSecond = nTime.Nanosecond() / 1e6
	)

	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d.%03d",
		nYear, nMonth, nDay, nHour, nMinute, nSecond, nMicrSecond)
}

func sourceCodeInfo() string {
	if pc, _, lineno, ok := runtime.Caller(3); ok {
		src := runtime.FuncForPC(pc).Name()
		dlen := len(src) - maxFuncNameLength
		if dlen > 0 {
			src = "..." + src[dlen:]
		}
		return fmt.Sprintf("[%s:%d]", src, lineno)
	}
	return ""
}

func writeLog(lvl Level, msg string) {
	fmt.Printf("[%s] %s %s %s",
		LogLevelName[lvl], currentTimestamp(), sourceCodeInfo(), msg)
}
