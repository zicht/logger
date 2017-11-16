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
	levels 		 map[logger.ChannelName]logger.LogLevel
}

func NewThresholdChannelHandler(handler logger.HandlerInterface, levels map[logger.ChannelName]logger.LogLevel, buffSize int, channels ...logger.ChannelName) logger.HandlerInterface {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &threshold{
		Strategy: &ThresholdChannelHandler{levels},
		handler: handler,
		buffer: make([]*logger.Record, 0, buffSize),
		is_buffering: true,
		Handler: Handler{
			channels:   cn,
			level:      logger.LogLevel(0),
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (f *ThresholdChannelHandler) StopBuffering(record *logger.Record) bool {
	if val, ok := f.levels[record.Channel]; ok {
		return record.Level >= val
	}
	return false
}

func (f *ThresholdChannelHandler) ValidateBuffering(buf BufferInterface) {
	buf.SetBuffering(true)
}
