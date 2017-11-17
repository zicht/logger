package handlers

import (
	"github.com/pbergman/logger"
)

type NoOpHandler struct {
	bubble   bool
	channels *logger.ChannelNames
	level    logger.LogLevel
}

func NewNoOpHandler(level logger.LogLevel, bubble bool, channels ...logger.ChannelName) logger.HandlerInterface {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &NoOpHandler{
		bubble:   bubble,
		channels: cn,
		level: level,
	}
}
func (h *NoOpHandler) GetChannels() *logger.ChannelNames {
	if h.channels == nil {
		h.channels = new(logger.ChannelNames)
	}
	return h.channels
}

func (h *NoOpHandler) HasChannels() bool { return h.channels != nil && h.channels.Len() > 0 }
func (h *NoOpHandler) SetChannels(c *logger.ChannelNames) { h.channels = c }
func (w *NoOpHandler) Handle(record *logger.Record) bool { return w.bubble }
func (w *NoOpHandler) Support(record logger.Record) bool { return w.level <= record.Level }
func (h *NoOpHandler) AddProcessor(processor logger.Processor) {}
func (h *NoOpHandler) GetProcessors() *logger.Processors { return nil }
func (h *NoOpHandler) GetFormatter() logger.FormatterInterface { return nil }
func (h *NoOpHandler) SetFormatter(f logger.FormatterInterface) {}