package handlers

import (
	"os"

	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type stderrHandler struct {
	logger.Handler
}

func NewStderrHandler(level int16) *stderrHandler  {
	return &stderrHandler {logger.Handler{Level: level, Formatter: formatters.NewLineFormatter()}}
}

func (h *stderrHandler) Write(name string, level string, message logger.MessageInterface) {
	h.GetFormatter().Execute(name, os.Stderr, h.CreateDataMap(message, name, level))
}
