package handlers

import (
	"io"

	"github.com/pbergman/logger"
)

type WriterHandler struct {
	writer io.Writer
	Handler
}

func NewWriterHandler(writer io.Writer, level logger.LogLevel, channels ...logger.ChannelName) *WriterHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &WriterHandler{
		writer,
		Handler{
			channels:   cn,
			level:      level,
			bubble:     true,
			processors: new(logger.Processors),
		},
	}
}

func (w *WriterHandler) Handle(record *logger.Record) bool {

	w.processRecord(record)

	if err := w.GetFormatter().Format(*record, w.writer); err != nil {
		panic("Handler: Failed to format message, " + err.Error())
	}

	return w.bubble
}

func (w *WriterHandler) Support(record logger.Record) bool {
	return w.level <= record.Level
}

func (w *WriterHandler) Close() error {
	if closer, ok := w.writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
