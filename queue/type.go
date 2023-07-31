package queue

import (
	"time"
)

type IQueue interface {
	Push(job Job)
	Size() int
	Pop() Job
	Run()
}

// Retry 失败重试
func Retry(times int, fn func() error, sleepMilliseconds ...int) error {
	if err := fn(); err != nil {
		if times--; times > 0 {
			var sleep int
			if len(sleepMilliseconds) == 1 {
				sleep = sleepMilliseconds[0]
			}
			if sleep > 0 {
				time.Sleep(time.Duration(sleep) * time.Millisecond)
			}
			return Retry(times, fn, sleep)
		}
		return err
	}
	return nil
}
