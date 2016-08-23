package handlers

import (
	"github.com/pbergman/logger"
	"io"
)

type WriterHandler struct {
	writer io.Writer
	Handler
}

func NewWriterHandler(name string, writer io.Writer, level logger.LogLevel, channels ...logger.ChannelName) *WriterHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &WriterHandler{
		writer,
		Handler{
			name:       name,
			channels:   cn,
			level:      level,
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (w *WriterHandler) Handle(record *logger.Record) bool {

	if w.processors.Len() > 0 {
		for _, i := range w.processors.Keys() {
			(*w.processors)[i](record)
		}
	}

	if err := w.GetFormatter().Format(*record, w.writer); err != nil {
		panic("Handler: Failed to format message, " + err.Error())
	}

	return w.bubble
}

func (w *WriterHandler) Support(record logger.Record) bool {
	return w.level <= record.Level
}
