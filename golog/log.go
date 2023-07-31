package golog

import (
	"io"
)

var goLog = DefaultGoLog()

func Reset() {
	goLog = DefaultGoLog()
}
func Clone() IGoLog {
	return goLog.Clone()
}
func SetOutput(w io.Writer) IGoLog {
	return goLog.SetOutput(w)
}
func SetPrefix(s string) IGoLog {
	return goLog.SetPrefix(s)
}
func WithContext(context any) IGoLog {
	return goLog.WithContext(context)
}

func SetLevelColorFunc(LevelColor func(Level, s string) string) IGoLog {
	return goLog.SetLevelColorFunc(LevelColor)
}
func Emergency(message string, context any) {
	goLog.Emergency(message, context)
}
func EmergencyF(format string, v ...any) {
	goLog.EmergencyF(format, v...)
}

func EmergencyV(v ...any) {
	goLog.EmergencyV(v...)
}

func Alert(message string, context any) {
	goLog.Alert(message, context)
}
func AlertF(format string, v ...any) {
	goLog.AlertF(format, v...)
}
func AlertV(v ...any) {
	goLog.AlertV(v...)
}

func Critical(message string, context any) {
	goLog.Critical(message, context)
}
func CriticalF(format string, v ...any) {
	goLog.CriticalF(format, v...)
}
func CriticalV(v ...any) {
	goLog.CriticalV(v...)
}

func Error(message string, context any) {
	goLog.Error(message, context)
}
func ErrorF(format string, v ...any) {
	goLog.ErrorF(format, v...)
}
func ErrorV(v ...any) {
	goLog.ErrorV(v...)
}

func Warning(message string, context any) {
	goLog.Warning(message, context)
}
func WarningF(format string, v ...any) {
	goLog.WarningF(format, v...)
}
func WarningV(v ...any) {
	goLog.WarningV(v...)
}
func Notice(message string, context any) {
	goLog.Notice(message, context)
}
func NoticeF(format string, v ...any) {
	goLog.NoticeF(format, v...)
}
func NoticeV(v ...any) {
	goLog.NoticeV(v...)
}
func Info(message string, context any) {
	goLog.Info(message, context)
}
func InfoF(format string, v ...any) {
	goLog.InfoF(format, v...)
}
func InfoV(v ...any) {
	goLog.InfoV(v...)
}
func Debug(message string, context any) {
	goLog.Debug(message, context)
}
func DebugF(format string, v ...any) {
	goLog.DebugF(format, v...)
}
func DebugV(v ...any) {
	goLog.DebugV(v...)
}
func Panic(message string, context any) {
	goLog.Panic(message, context)
}
func PanicF(format string, v ...any) {
	goLog.PanicF(format, v...)
}
func PanicV(v ...any) {
	goLog.PanicV(v...)
}
func Log(level, message string, context any) {
	goLog.Log(level, message, context)
}
