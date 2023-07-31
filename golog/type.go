package golog

import (
	"fmt"
	"io"
)

//https://www.rfc-editor.org/rfc/rfc5424
//紧急，警报，关键，错误，警告，通知，信息和调试。
const (
	EMERGENCY = "emergency"
	ALERT     = "alert"
	CRITICAL  = "critical"
	ERROR     = "error"
	WARNING   = "warning"
	NOTICE    = "notice"
	INFO      = "info"
	DEBUG     = "debug"
	PANIC     = "Panic"
)

var defaultLevelColorFunc = func(level, s string) string {
	const (
		color_black = uint8(iota + 90)
		color_red
		color_green
		color_yellow
		color_blue
		color_magenta
		color_cyan
		color_white
	)
	var color uint8
	switch level {
	case EMERGENCY:
		color = color_magenta
	case ALERT:
		color = color_blue
	case CRITICAL:
		color = color_cyan
	case ERROR:
		color = color_red
	case WARNING:
		color = color_yellow
	case NOTICE:
		color = color_green
	case INFO:
		color = color_white
	case DEBUG:
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
}

type IGoLog interface {
	SetOutput(w io.Writer) IGoLog
	SetPrefix(prefix string) IGoLog
	SetFlags(flag int) IGoLog
	WithContext(context any) IGoLog
	SetLevelColorFunc(LevelColor func(Level, s string) string) IGoLog
	Clone() IGoLog
	Emergency(message string, context any)
	EmergencyF(format string, v ...any)
	EmergencyV(v ...any)

	Alert(message string, context any)
	AlertF(format string, v ...any)
	AlertV(v ...any)

	Critical(message string, context any)
	CriticalF(format string, v ...any)
	CriticalV(v ...any)

	Error(message string, context any)
	ErrorF(format string, v ...any)
	ErrorV(v ...any)

	Warning(message string, context any)
	WarningF(format string, v ...any)
	WarningV(v ...any)

	Notice(message string, context any)
	NoticeF(format string, v ...any)
	NoticeV(v ...any)

	Info(message string, context any)
	InfoF(format string, v ...any)
	InfoV(v ...any)

	Debug(message string, context any)
	DebugF(format string, v ...any)
	DebugV(v ...any)

	Panic(message string, context any)
	PanicF(format string, v ...any)
	PanicV(v ...any)

	Log(level, message string, context any)
}
