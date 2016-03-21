package handlers

import (
	"bytes"
	"github.com/pbergman/logger/formatters"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
)

// The data structure that inherets the logger handler and has
// a buffer to write the records to
type BufferHandler struct {
	Handler
	Buffer *bytes.Buffer
}

// NewBufferHandler initialize a new handler that writes the records to
// the buffer and can be fetched later. (fo example getBuffer.String())
func NewBufferHandler(level level.LogLevel) *BufferHandler {
	return &BufferHandler{
		Handler{
			Level:     level,
			Formatter: formatters.NewLineFormatter(),
		},
		bytes.NewBuffer(nil),
	}
}

func (b BufferHandler) Support(level level.LogLevel) bool {
	return b.Level <= level
}

// Write records to buffer
func (h *BufferHandler) Write(name string, level level.LogLevel, message messages.MessageInterface) {
	h.GetFormatter().Execute(name, h.Buffer, h.CreateDataMap(message, name, level))
}

// Will return interanl used buffer
func (h *BufferHandler) GetBuffer() *bytes.Buffer {
	return h.Buffer
}
