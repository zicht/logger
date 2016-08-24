package handlers

import (
	"github.com/pbergman/logger"
)

// ThresholdLevelHandler is a logger that will buffer the record until
// the given threshold is reached or exceeded. stop_buffering is false
// it will only process the logs in the buffer and then buffer again
type ThresholdLevelHandler struct {
	Handler
	threshold
}

func NewThresholdLevelHandler(handler logger.HandlerInterface, level logger.LogLevel, buffSize int, channels ...logger.ChannelName) *ThresholdLevelHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &ThresholdLevelHandler{
		Handler{
			channels:   cn,
			level:      level,
			bubble:     true,
			processors: new(logger.Processors),
		},
		threshold{
			handler,
			make([]*logger.Record, 0, buffSize),
			true,
			true,
		},
	}
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
