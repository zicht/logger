package handlers

import (
	"github.com/pbergman/logger"
	"io"
)

type MappedWriterHandler struct {
	writers map[logger.LogLevel]io.Writer
	Handler
}

// NewMappedWriterHandler can map different levels to write to different io.writers
func NewMappedWriterHandler(writers map[logger.LogLevel]io.Writer, channels ...logger.ChannelName) *MappedWriterHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &MappedWriterHandler{
		writers,
		Handler{
			channels:   cn,
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (h *MappedWriterHandler) Support(recored logger.Record) bool {
	for l, _ := range h.writers {
		if l <= recored.Level {
			return true
		}
	}
	return false
}

func (h *MappedWriterHandler) Handle(record *logger.Record) bool {

	var level logger.LogLevel

	for writerLevel, _ := range h.writers {
		if writerLevel <= record.Level {
			level = writerLevel
		}
	}

	h.processRecord(record)

	if buf, err := h.GetFormatter().Format(*record); err != nil {
		panic("Handler: Failed to format message, " + err.Error())
	} else {
		if _, err := h.writers[level].Write(buf); err != nil {
			panic("Handler: Failed to write message, " + err.Error())
		}
	}

	return h.bubble
}

func (h *MappedWriterHandler) Close() error {
	for _, w := range h.writers {
		if closer, ok := w.(io.Closer); ok {
			closer.Close()
		}
	}
	return nil
}
