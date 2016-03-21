package messages

import (
	"sync"
)

type Queue struct {
	size  int
	queue chan *Record
	wg    sync.WaitGroup
}

func NewQueue(size int) *Queue {
	return &Queue{
		size:  size,
		queue: make(chan *Record, size),
	}
}

func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *Queue) Valid() bool {
	return q.Len() > 0
}

func (q *Queue) Push(r *Record) {
	for q.Len() >= q.size {
		q.Pop()
	}
	q.queue <- r
}

func (q *Queue) Pop() *Record {
	if q.Len() > 0 {
		record := <-q.queue
		return record
	} else {
		return nil
	}
}
