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
	bubble 		bool
	writer      io.Writer
}

func NewWriterHandler(name string, writer io.Writer, level logger.LogLevel, channels ...logger.ChannelName) *WriterHandler {
	cn := new(logger.ChannelNames)
	for _, c := range channels {
		cn.AddChannel(c)
	}
	return &WriterHandler{name: name, channels: cn, level: level, bubble: true, writer: writer}
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

func (w *WriterHandler) Handle(record logger.Record) bool {

	if err := w.GetFormatter().Format(record, w.writer); err != nil {
		panic("Handler: Failed to format message, " + err.Error())
	}

	return w.bubble == false
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
	return w.level <= record.Level
}

func (w *WriterHandler) SetBubble(bubble bool) {
	w.bubble = bubble
}

func (w *WriterHandler) GetBubble() bool {
	return w.bubble
}