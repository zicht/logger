package handlers

import (
	"github.com/pbergman/logger/formatters"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"io"
)

type MappedWriterHandler struct {
	writers map[level.LogLevel]io.Writer
	Handler
}

// NewMappedWriterHandler can map different levels to write to different io.writers
func NewMappedWriterHandler(writers map[level.LogLevel]io.Writer) *MappedWriterHandler {
	return &MappedWriterHandler{writers, Handler{Formatter: formatters.NewLineFormatter()}}
}

func (h *MappedWriterHandler) Support(level level.LogLevel) bool {
	for l, _ := range h.writers {
		if l <= level {
			return true
		}
	}
	return false
}

func (h *MappedWriterHandler) Write(name string, l level.LogLevel, message messages.MessageInterface) {
	var logLevel level.LogLevel
	for key, _ := range h.writers {
		if key <= l {
			logLevel = key
		}
	}
	h.GetFormatter().Execute(name, h.writers[logLevel], h.CreateDataMap(message, name, l))
}
