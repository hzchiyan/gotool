package queue

type Job struct {
	Job func() error
	//失败重试次数
	MaxTries int
}
