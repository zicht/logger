package handlers

import (
	"bytes"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

// The data structure that inherets the logger handler and has
// a buffer to write the records to
type bufferHandler struct {
	logger.Handler
	buffer *bytes.Buffer
}

// NewBufferHandler initialize a new handler that writes the records to
// the buffer and can be fetched later. (fo example getBuffer.String())
func NewBufferHandler(level int16) *bufferHandler   {
	return &bufferHandler {
		logger.Handler{
			Level: 		level,
			Formatter: 	formatters.NewLineFormatter(),
		},
		bytes.NewBuffer(nil),
	}
}

// Write records to buffer
func (h *bufferHandler) Write(name string, level string, message logger.MessageInterface) {
	h.GetFormatter().Execute(name, h.buffer, h.CreateDataMap(message, name, level));
}

// Will return interanl used buffer
func (h *bufferHandler) GetBuffer() *bytes.Buffer {
	return h.buffer
}