package logger

import (
	"io"
)

type HandlerInterface interface {
	GetName() string
	GetLevel() LogLevel
	GetFormatter() FormatterInterface
	SetFormatter(Formatter FormatterInterface)
	GetChannels() *ChannelNames
	HasChannels() bool
	SetChannels(ChannelNames)
	Support(Record) bool
	io.WriteCloser
}

