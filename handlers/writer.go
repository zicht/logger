package handlers

import (
	"io"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type WriterHandler struct {
	name        string
	channels    *logger.ChannelNames
	formatter   logger.FormatterInterface
	level       logger.LogLevel
	writer      io.Writer
}

func NewWriterHandler(name string, writer io.Writer, level logger.LogLevel, channels ...logger.ChannelName) *WriterHandler {

	cn := new(logger.ChannelNames)

	for _, c := range channels {
		cn.AddChannel(c)
	}

	return &WriterHandler{name: name, channels: cn, level: level, writer: writer}
}

func (w *WriterHandler) GetName() string {
	return w.name
}

func (w *WriterHandler) HasChannels() bool {
	return w.channels != nil && w.channels.Len() > 0
}

func (w *WriterHandler) SetChannels(c *logger.ChannelNames) {
	w.channels = c
}

func (w *WriterHandler) GetChannels() *logger.ChannelNames {
	if w.channels == nil {
		w.channels = new(logger.ChannelNames)
	}

	return w.channels
}


func (w *WriterHandler) GetLevel() logger.LogLevel {
	return w.level
}

func (w *WriterHandler) GetFormatter() logger.FormatterInterface {
	if w.formatter == nil {
		w.formatter = formatters.NewLineFormatter()
	}

	return w.formatter
}

func (w *WriterHandler) SetFormatter(f logger.FormatterInterface) {
	w.formatter = f
}

func (w *WriterHandler) Support(record logger.Record) bool {
	// @todo move to logger
	if w.HasChannels() {
		if index := w.channels.FindChannel(record.Channel.GetName()); index >= 0 {
			if (*w.channels)[index].IsExcluded() {
				return false
			}
		} else {
			return false
		}
	}
	return w.level <= record.Level
}

func (w *WriterHandler) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *WriterHandler) Close() error {
	if closer, ok := w.writer.(io.Closer); ok {
		return closer.Close()
	} else {
		return nil
	}
}