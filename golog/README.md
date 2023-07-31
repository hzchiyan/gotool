
## Basic Usage

~~~
package main


import "github.com/hzchiyan/gotool/golog"

func main() {
	l := golog.DefaultGoLog()
	l.Emergency("Emergency", nil)
	golog.EmergencyF("Emergency %s", "test v")
	golog.EmergencyV("Emergency", "test v")
	l.Alert("Alert", nil)
	golog.AlertF("Alert %s", "test v")
	golog.AlertV("Alert", "test v")
	l.Critical("Critical", nil)
	golog.CriticalF("Critical %s", "test v")
	golog.CriticalV("Critical", "test v")
	l.Warning("Warning", nil)
	golog.WarningF("Warning %s", "test v")
	golog.WarningV("Warning", "test v")
	l.Notice("Notice", nil)
	golog.NoticeF("Notice %s", "test v")
	golog.NoticeV("Notice", "test v")
	l.Info("Info", nil)
	golog.InfoF("Info %s", "test v")
	golog.InfoV("Info", "test v")
	l.Debug("Debug", nil)
	golog.DebugF("Debug %s", "test v")
	golog.DebugV("Debug", "test v")
	l.Error("Error", nil)
	golog.ErrorF("Error  %s", "test v")
	golog.ErrorV("Error", "v1", "v2")
	l.Panic("Panic", nil)
	golog.PanicF("Panic %s", "test v")
	golog.PanicV("Panic", "test v")
}
~~~

> This context is not another context. This context means that any data type can be passed
>
> You can use the WithContext method to carry any data type
>
> When using the WithContext method, the data will be loaded into the memory. It is better not to print the contexts to the parameters next time, but to use Clone for coherent operations.
>
> For example
>
> ```
> golog.WithContext("this is context data").DebugV("this is debug with context data")
> golog.Clone().DebugF("No need to carry content data")
> golog.DebugV("Continue to carry content data")
> ```
