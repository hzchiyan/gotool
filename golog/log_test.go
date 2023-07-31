package golog

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAlert(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alert(tt.args.message, tt.args.context)
		})
	}
}

func TestAlertF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlertF(tt.args.format, tt.args.v...)
		})
	}
}

func TestAlertV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AlertV(tt.args.v...)
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name string
		want IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCritical(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critical(tt.args.message, tt.args.context)
		})
	}
}

func TestCriticalF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CriticalF(tt.args.format, tt.args.v...)
		})
	}
}

func TestCriticalV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CriticalV(tt.args.v...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.message, tt.args.context)
		})
	}
}

func TestDebugF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugF(tt.args.format, tt.args.v...)
		})
	}
}

func TestDebugV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugV(tt.args.v...)
		})
	}
}

func TestEmergency(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Emergency(tt.args.message, tt.args.context)
		})
	}
}

func TestEmergencyF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmergencyF(tt.args.format, tt.args.v...)
		})
	}
}

func TestEmergencyV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmergencyV(tt.args.v...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.message, tt.args.context)
		})
	}
}

func TestErrorF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorF(tt.args.format, tt.args.v...)
		})
	}
}

func TestErrorV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorV(tt.args.v...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.message, tt.args.context)
		})
	}
}

func TestInfoF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoF(tt.args.format, tt.args.v...)
		})
	}
}

func TestInfoV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoV(tt.args.v...)
		})
	}
}

func TestLog(t *testing.T) {
	type args struct {
		level   string
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Log(tt.args.level, tt.args.message, tt.args.context)
		})
	}
}

func TestNotice(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Notice(tt.args.message, tt.args.context)
		})
	}
}

func TestNoticeF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoticeF(tt.args.format, tt.args.v...)
		})
	}
}

func TestNoticeV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoticeV(tt.args.v...)
		})
	}
}

func TestPanic(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Panic(tt.args.message, tt.args.context)
		})
	}
}

func TestPanicF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PanicF(tt.args.format, tt.args.v...)
		})
	}
}

func TestPanicV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PanicV(tt.args.v...)
		})
	}
}

func TestReset(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reset()
		})
	}
}

func TestSetLevelColorFunc(t *testing.T) {
	type args struct {
		LevelColor func(Level, s string) string
	}
	tests := []struct {
		name string
		args args
		want IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetLevelColorFunc(tt.args.LevelColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLevelColorFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOutput(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
		want  IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got := SetOutput(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("SetOutput() gotW = %v, want %v", gotW, tt.wantW)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPrefix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPrefix(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarning(t *testing.T) {
	type args struct {
		message string
		context any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warning(tt.args.message, tt.args.context)
		})
	}
}

func TestWarningF(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningF(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarningV(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarningV(tt.args.v...)
		})
	}
}

func TestWithContext(t *testing.T) {
	type args struct {
		context any
	}
	tests := []struct {
		name string
		args args
		want IGoLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
