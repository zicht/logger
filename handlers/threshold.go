package handlers

import (
	"io"
	"sort"

	"github.com/zicht/logger"
)

type BufferStrategy interface {
	StopBuffering(*logger.Record) bool
	ValidateBuffering(BufferInterface)
}

type BufferInterface interface {
	SetBuffering(v bool)
	IsBuffering() bool
}

type threshold struct {
	handler      logger.HandlerInterface
	buffer       []*logger.Record
	is_buffering bool
	Strategy     BufferStrategy
	// embedded handler
	Handler
}

func (f *threshold) Support(record logger.Record) bool {
	return true
}

func (f *threshold) bufferKeys() []int {
	var keys []int
	for k := range f.buffer {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (f *threshold) bufferWalk(c func(r *logger.Record)) {
	for _, i := range f.bufferKeys() {
		c(f.buffer[i])
	}
}

func (f *threshold) Clear() {
	f.is_buffering = true
	f.clearBuffer()
}

func (f *threshold) IsBuffering() bool {
	return f.is_buffering
}

func (f *threshold) SetBuffering(v bool) {
	f.is_buffering = v
}

func (f *threshold) clearBuffer() {
	f.buffer = f.buffer[:0]
}

func (f *threshold) addToBuffer(record *logger.Record) {
	if cap(f.buffer) == len(f.buffer) {
		f.buffer = append(f.buffer[:0], f.buffer[1:]...)
	}
	f.buffer = append(f.buffer, record)
}

func (f *threshold) Close() error {
	if closer, ok := f.handler.(io.Closer); ok {
		return closer.Close()
	} else {
		return nil
	}
}

func (f *threshold) flush() {
	f.bufferWalk(func(r *logger.Record) {
		if f.handler.Support(*r) {
			f.handler.Handle(r)
		}
	})
}

func (f *threshold) Handle(record *logger.Record) bool {
	f.processRecord(record)
	if !f.is_buffering {
		f.Strategy.ValidateBuffering(f)
	}
	if f.is_buffering {
		f.addToBuffer(record)
		if f.Strategy.StopBuffering(record) {
			f.is_buffering = false
			f.flush()
			f.clearBuffer()
		}
	} else {
		f.handler.Handle(record)
	}
	return f.bubble
}
