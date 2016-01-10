package logger

import (
	"sync"
)

type queue struct {
	size  int
	queue chan *Record
	wg    sync.WaitGroup
}

func NewQueue(size int) *queue {
	return &queue{
		size:  size,
		queue: make(chan *Record, size),
	}
}

func (q *queue) Len() int {
	return len(q.queue)
}

func (q *queue) Valid() bool {
	return q.Len() > 0
}

func (q *queue) Push(r *Record) {
	for q.Len() >= q.size {
		q.Pop()
	}
	q.queue <- r
}

func (q *queue) Pop() *Record {
	if q.Len() > 0 {
		record := <-q.queue
		return record
	} else {
		return nil
	}
}
