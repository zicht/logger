package logger

import "sync"

type recordBuffer struct {
	buf  []*Record
	lock sync.RWMutex
}

func newRecordBuffer(max int) *recordBuffer {
	return &recordBuffer{buf: make([]*Record, 0, max)}
}

func (b *recordBuffer) push(r *Record) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if cap(b.buf) == len(b.buf) {
		b.buf = append(b.buf[:0], b.buf[1:]...)
	}
	b.buf = append(b.buf, r)
}

func (b *recordBuffer) shift() *Record {
	b.lock.Lock()
	defer b.lock.Unlock()
	var record *Record
	if len(b.buf) > 0 {
		record, b.buf = b.buf[0], b.buf[1:]
	}
	return record
}

func (b *recordBuffer) len() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return len(b.buf)
}

func (b *recordBuffer) valid() bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return len(b.buf) > 0
}
