package golog

import (
	"bytes"
	"log"
	"reflect"
	"sync"
	"testing"
)

func TestDefaultGoLog(t *testing.T) {
	tests := []struct {
		name string
		want IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultGoLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultGoLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_Alert(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Alert(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_AlertF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.AlertF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_AlertV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.AlertV(tt.args.v...)
		})
	}
}

func TestGoLog_Clone(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			if got := l.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_Critical(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Critical(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_CriticalF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.CriticalF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_CriticalV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.CriticalV(tt.args.v...)
		})
	}
}

func TestGoLog_Debug(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Debug(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_DebugF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.DebugF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_DebugV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.DebugV(tt.args.v...)
		})
	}
}

func TestGoLog_Emergency(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Emergency(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_EmergencyF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.EmergencyF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_EmergencyV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.EmergencyV(tt.args.v...)
		})
	}
}

func TestGoLog_Error(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Error(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_ErrorF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.ErrorF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_ErrorV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.ErrorV(tt.args.v...)
		})
	}
}

func TestGoLog_Info(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Info(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_InfoF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.InfoF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_InfoV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.InfoV(tt.args.v...)
		})
	}
}

func TestGoLog_Log(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		level   string
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Log(tt.args.level, tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_Notice(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Notice(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_NoticeF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.NoticeF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_NoticeV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.NoticeV(tt.args.v...)
		})
	}
}

func TestGoLog_Panic(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Panic(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_PanicF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.PanicF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_PanicV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.PanicV(tt.args.v...)
		})
	}
}

func TestGoLog_SetFlags(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		flag int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			if got := l.SetFlags(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_SetLevelColorFunc(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		levelColorFunc func(level, s string) string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			if got := l.SetLevelColorFunc(tt.args.levelColorFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLevelColorFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_SetOutput(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		wantW  string
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			w := &bytes.Buffer{}
			got := l.SetOutput(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("SetOutput() gotW = %v, want %v", gotW, tt.wantW)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_SetPrefix(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			if got := l.SetPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoLog_Warning(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.Warning(tt.args.message, tt.args.context)
		})
	}
}

func TestGoLog_WarningF(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.WarningF(tt.args.format, tt.args.v...)
		})
	}
}

func TestGoLog_WarningV(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			l.WarningV(tt.args.v...)
		})
	}
}

func TestGoLog_WithContext(t *testing.T) {
	type fields struct {
		l              *log.Logger
		contexts       []any
		isDisplayColor bool
		levelColorFunc func(Level, s string) string
		mu             sync.Mutex
	}
	type args struct {
		context any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GoLog{
				l:              tt.fields.l,
				contexts:       tt.fields.contexts,
				isDisplayColor: tt.fields.isDisplayColor,
				levelColorFunc: tt.fields.levelColorFunc,
				mu:             tt.fields.mu,
			}
			if got := l.WithContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
