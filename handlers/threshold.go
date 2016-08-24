package handlers

import (
	"io"
	"sort"
	"github.com/pbergman/logger"
)

type threshold struct {
	handler logger.HandlerInterface
	buffer  []*logger.Record
	// when true it will stop buffering and just process
	// records when threshold level is reached when false
	// it will buffer again till threshold is reached again
	stop_buffering bool
	is_buffering   bool
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
	f.buffer = f.buffer[:0]
}

func (f *threshold) SetStopBuffering(v bool) {
	if !v && !f.is_buffering {
		// turn buffer back on
		f.is_buffering = true
	}
	f.stop_buffering = v
}

func (f *threshold) IsStopBuffering() bool {
	return f.stop_buffering
}

func (f *threshold) IsBuffering() bool {
	return f.is_buffering
}

func (f *threshold) Close() error {
	if closer, ok := f.handler.(io.Closer); ok {
		return closer.Close()
	} else {
		return nil
	}
}