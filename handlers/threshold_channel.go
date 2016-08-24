package handlers

import (
	"github.com/pbergman/logger"
)

// ThresholdChannelHandler is a logger that will buffer the record until
// the given threshold is reached or exceeded. stop_buffering is false
// it will only process the logs in the buffer and then buffer again.
//
// It is similar as ThresholdLevelHandler but instead of setting a
// threshold level you set a map channels and levels that will be
// matched against the giver record
type ThresholdChannelHandler struct {
	levels  map[logger.ChannelName]logger.LogLevel
	Handler
	threshold
}

func NewThresholdChannelHandler(handler logger.HandlerInterface, levels map[logger.ChannelName]logger.LogLevel, buffSize int, channels ...logger.ChannelName) *ThresholdChannelHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &ThresholdChannelHandler{
		levels,
		Handler{
			channels:   cn,
			level:      logger.LogLevel(0),
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

func (f *ThresholdChannelHandler) Handle(record *logger.Record) bool {
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
		if val, ok := f.levels[record.Channel]; ok && record.Level >= val {
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