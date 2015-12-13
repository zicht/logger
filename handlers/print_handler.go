package handlers

import (
	"fmt"

	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type printHandler struct {
	logger.Handler
}

func NewPrintHandler(level int16) *printHandler {
	return &printHandler{logger.Handler{Level: level, Formatter: formatters.NewLineFormatter()}}
}

func (h *printHandler) Write(name string, level string, message logger.MessageInterface) {
	fmt.Print(h.GetFormatter().Format(name, level, message))
}
