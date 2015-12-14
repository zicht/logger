package handlers

import (
	"bytes"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type bufferHandler struct {
	logger.Handler
	buffer *bytes.Buffer
}

func NewBufferHandler(level int16) *bufferHandler   {
	var buff bytes.Buffer
	return &bufferHandler {
		logger.Handler{
			Level: 		level,
			Formatter: 	formatters.NewLineFormatter(),
		},
		&buff,
	}
}

func (h *bufferHandler) Write(name string, level string, message logger.MessageInterface) {
	h.GetFormatter().Execute(name, h.buffer, h.CreateDataMap(message, name, level));
}

func (h *bufferHandler) GetBuffer() *bytes.Buffer {
	return h.buffer
}