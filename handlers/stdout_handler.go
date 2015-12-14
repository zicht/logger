package handlers

import (
	"os"

	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type stdoutHandler struct {
	logger.Handler
}

// NewStdoutHandler will write all records to stdout (os.stdout)
func NewStdoutHandler(level int16) *stdoutHandler  {
	return &stdoutHandler {logger.Handler{Level: level, Formatter: formatters.NewLineFormatter()}}
}

func (h *stdoutHandler) Write(name string, level string, message logger.MessageInterface) {
	h.GetFormatter().Execute(name, os.Stdout, h.CreateDataMap(message, name, level));
}
