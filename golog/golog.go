package golog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"sync"
)

type GoLog struct {
	l              *log.Logger
	contexts       []any
	isDisplayColor bool
	levelColorFunc func(Level, s string) string
	mu             sync.Mutex
}

func DefaultGoLog() IGoLog {
	g := GoLog{l: log.Default()}
	return g.Clone()
}

func (l GoLog) Clone() IGoLog {
	l.contexts = nil
	l.isDisplayColor = false
	l.SetLevelColorFunc(defaultLevelColorFunc)
	l.mu = sync.Mutex{}
	return &l
}
func (l *GoLog) SetOutput(w io.Writer) IGoLog {
	l.l.SetOutput(w)
	return l
}
func (l *GoLog) SetPrefix(prefix string) IGoLog {
	l.l.SetPrefix(prefix)
	return l
}
func (l *GoLog) SetFlags(flag int) IGoLog {
	l.l.SetFlags(flag)
	return l
}
func (l *GoLog) WithContext(context any) IGoLog {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.contexts = append(l.contexts, context)
	return l
}

func (l *GoLog) SetLevelColorFunc(levelColorFunc func(level, s string) string) IGoLog {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.levelColorFunc = levelColorFunc
	return l
}
func (l *GoLog) Emergency(message string, context any) {
	l.Log(EMERGENCY, message, context)
}

func (l *GoLog) EmergencyF(format string, v ...any) {
	l.Emergency(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) EmergencyV(v ...any) {
	l.Emergency(fmt.Sprint(v...), nil)
}

func (l *GoLog) Alert(message string, context any) {
	l.Log(ALERT, message, context)
}
func (l *GoLog) AlertF(format string, v ...any) {
	l.Alert(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) AlertV(v ...any) {
	l.Alert(fmt.Sprint(v...), nil)
}

func (l *GoLog) Critical(message string, context any) {
	l.Log(CRITICAL, message, context)
}
func (l *GoLog) CriticalF(format string, v ...any) {
	l.Critical(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) CriticalV(v ...any) {
	l.Critical(fmt.Sprint(v...), nil)
}

func (l *GoLog) Error(message string, context any) {
	l.Log(ERROR, message, context)
}
func (l *GoLog) ErrorF(format string, v ...any) {
	l.Error(fmt.Sprintf(format, v...), nil)
}

func (l *GoLog) ErrorV(v ...any) {
	l.Critical(fmt.Sprint(v...), nil)
}

func (l *GoLog) Warning(message string, context any) {
	l.Log(WARNING, message, context)
}
func (l *GoLog) WarningF(format string, v ...any) {
	l.Warning(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) WarningV(v ...any) {
	l.Warning(fmt.Sprint(v...), nil)
}
func (l *GoLog) Notice(message string, context any) {
	l.Log(NOTICE, message, context)
}
func (l *GoLog) NoticeF(format string, v ...any) {
	l.Notice(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) NoticeV(v ...any) {
	l.Notice(fmt.Sprint(v...), nil)
}
func (l *GoLog) Info(message string, context any) {
	l.Log(INFO, message, context)
}
func (l *GoLog) InfoF(format string, v ...any) {
	l.Info(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) InfoV(v ...any) {
	l.Info(fmt.Sprint(v...), nil)
}
func (l *GoLog) Debug(message string, context any) {
	l.Log(DEBUG, message, context)
}
func (l *GoLog) DebugF(format string, v ...any) {
	l.Debug(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) DebugV(v ...any) {
	l.Debug(fmt.Sprint(v...), nil)
}

func (l *GoLog) Panic(message string, context any) {
	l.Log(PANIC, message, context)
}
func (l *GoLog) PanicF(format string, v ...any) {
	l.Panic(fmt.Sprintf(format, v...), nil)
}
func (l *GoLog) PanicV(v ...any) {
	l.Panic(fmt.Sprint(v...), nil)
}
func (l *GoLog) Log(level, message string, context any) {
	l.WithContext(context)
	s := level
	s += " " + message
	for _, ctx := range l.contexts {
		if ctx != nil {
			s += " " + fmt.Sprintf("%#v", ctx)
		}
	}
	//输出堆栈信息
	if level == ERROR {
		s += string(debug.Stack())
	}
	if l.levelColorFunc != nil && l.l.Writer() == os.Stderr {
		s = l.levelColorFunc(level, s)
	}
	_ = l.l.Output(2, s)
	if level == ERROR {
		os.Exit(1)
	}
	if level == PANIC {
		panic(s)
	}
}
