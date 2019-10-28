package handlers

import (
	"github.com/zicht/logger"
)

// ThresholdLevelHandler is a logger that will buffer the record until
// the given threshold is reached or exceeded. stop_buffering is false
// it will only process the logs in the buffer and then buffer again
type ThresholdLevelHandler struct {
	level logger.LogLevel
}

func NewThresholdLevelHandler(handler logger.HandlerInterface, level logger.LogLevel, buffSize int, channels ...logger.ChannelName) logger.HandlerInterface {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &threshold{
		handler:      handler,
		Strategy:     &ThresholdLevelHandler{level},
		buffer:       make([]*logger.Record, 0, buffSize),
		is_buffering: true,
		Handler: Handler{
			channels:   cn,
			level:      level,
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (f *ThresholdLevelHandler) StopBuffering(record *logger.Record) bool {
	return record.Level >= f.level
}

func (f *ThresholdLevelHandler) ValidateBuffering(buf BufferInterface) {
	buf.SetBuffering(true)
}
