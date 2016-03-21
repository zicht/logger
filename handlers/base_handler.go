package handlers

import (
	"github.com/pbergman/logger/formatters"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
)

type HandlerInterface interface {
	Write(name string, level level.LogLevel, message messages.MessageInterface)
	SetFormatter(Formatter formatters.FormatterInterface)
	Support(level level.LogLevel) bool
	CreateDataMap(message messages.MessageInterface, name string, level level.LogLevel) map[string]interface{}
}

type Handler struct {
	Level     level.LogLevel
	Formatter formatters.FormatterInterface
}

func (h *Handler) SetFormatter(Formatter formatters.FormatterInterface) {
	h.Formatter = Formatter
}

func (h Handler) GetFormatter() formatters.FormatterInterface {
	return h.Formatter
}

func (h Handler) CreateDataMap(message messages.MessageInterface, name string, level level.LogLevel) map[string]interface{} {
	return map[string]interface{}{
		"message": message.GetMessage(),
		"extra":   message.GetContext(),
		"trace":   message.GetTrace(),
		"time":    message.GetTime(),
		"name":    name,
		"level":   level,
	}
}
