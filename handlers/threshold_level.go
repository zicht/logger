package handlers

import (
	"github.com/pbergman/logger"
	"io"
	"sort"
)

// ThresholdLevelHandler is a logger that will buffer the record until
// the given threshold is reached or exceeded. stop_buffering is false
// it will only process the logs in the buffer and then buffer again
type ThresholdLevelHandler struct {
	handler logger.HandlerInterface
	buffer  []*logger.Record
	// when true it will stop buffering and just process
	// records when threshold level is reached when false
	// it will buffer again till threshold is reached again
	stop_buffering bool
	is_buffering   bool
	Handler
}

func NewThresholdLevelHandler(handler logger.HandlerInterface, threshold logger.LogLevel, buffSize int, channels ...logger.ChannelName) *ThresholdLevelHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &ThresholdLevelHandler{
		handler,
		make([]*logger.Record, 0, buffSize),
		true,
		true,
		Handler{
			channels:   cn,
			level:      threshold,
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (f *ThresholdLevelHandler) Support(record logger.Record) bool {
	return true
}

func (f *ThresholdLevelHandler) Handle(record *logger.Record) bool {
	if f.processors.Len() > 0 {
		for _, i := range f.processors.Keys() {
			(*f.processors)[i](record)
		}
	}
	if f.is_buffering {
		// shift first element
		if cap(f.buffer) == len(f.buffer) {
			f.buffer = append(f.buffer[:0], f.buffer[1:]...)
		}
		f.buffer = append(f.buffer, record)
		if record.Level >= f.level {
			if f.stop_buffering {
				f.is_buffering = false
			}
			f.bufferWalk(func(r *logger.Record) {
				f.handler.Handle(r)
			})
			f.buffer = f.buffer[:0]
		}
	} else {
		f.handler.Handle(record)
	}

	return f.bubble
}

func (f *ThresholdLevelHandler) bufferKeys() []int {
	var keys []int
	for k := range f.buffer {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (f *ThresholdLevelHandler) bufferWalk(c func(r *logger.Record)) {
	for _, i := range f.bufferKeys() {
		c(f.buffer[i])
	}
}

func (f *ThresholdLevelHandler) Clear() {
	f.is_buffering = true
	f.buffer = f.buffer[:0]
}

func (f *ThresholdLevelHandler) SetStopBuffering(v bool) {
	if !v && !f.is_buffering {
		// turn buffer back on
		f.is_buffering = true
	}
	f.stop_buffering = v
}

func (f *ThresholdLevelHandler) IsStopBuffering() bool {
	return f.stop_buffering
}

func (f *ThresholdLevelHandler) IsBuffering() bool {
	return f.is_buffering
}

func (f *ThresholdLevelHandler) Close() error {
	if closer, ok := f.handler.(io.Closer); ok {
		return closer.Close()
	} else {
		return nil
	}
}
