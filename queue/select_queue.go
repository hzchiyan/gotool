package queue

var DefaultSelectQueueLimit = 10000

type SelectQueue struct {
	limit  int
	closed bool
	// 入队job
	C chan Job
	//执行失败的job
	E chan Job
}

func NewSelectQueue(limit int) *SelectQueue {
	q := &SelectQueue{closed: false, limit: limit}
	return q
}
func (s *SelectQueue) clone() *SelectQueue {
	if s.limit == 0 {
		s.limit = DefaultSelectQueueLimit
	}
	s.C = make(chan Job, s.limit)
	s.E = make(chan Job, s.limit)
	return s
}

func (s *SelectQueue) Push(job Job) {
	s.clone()
	if s.limit > 0 {
		s.C <- job
	}
}

func (s *SelectQueue) Size() int {
	return len(s.C)
}
func (s *SelectQueue) Pop() Job {
	return <-s.C
}

func (s *SelectQueue) Run() {
	for {
		select {
		case job := <-s.C:
			err := Retry(job.MaxTries, job.Job)
			if err != nil {
				s.closed = true
				s.C <- job
				return
			}
		}
	}
}
