package main

import (
	"fmt"
	"github.com/hzchiyan/gotool/times"
	"time"
)

func main() {
	_ = times.SetInterval(func(args ...any) {
		fmt.Println(args)
	}, 1000, "hello world")
	_ = times.SetTimeout(func(args ...any) {
		fmt.Println(args)
	}, 1200, "I ran once")
	time.Sleep(5 * time.Second)
}
