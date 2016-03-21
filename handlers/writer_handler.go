package handlers

import (
	"io"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/formatters"
	"github.com/pbergman/logger/messages"
)

type WriterHandler struct {
	writer io.Writer
	Handler
}

func (w *WriterHandler) Support(level level.LogLevel) bool {
	return w.Level <= level
}

// NewWriterHandler can write to a object implementing io.Writer interface
func NewWriterHandler(writer io.Writer, level level.LogLevel) *WriterHandler {
	return &WriterHandler{writer, Handler{Level: level, Formatter: formatters.NewLineFormatter()}}
}

func (h *WriterHandler) Write(name string, level level.LogLevel, message messages.MessageInterface) {
	h.GetFormatter().Execute(name, h.writer, h.CreateDataMap(message, name, level))
}
