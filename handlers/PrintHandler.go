package printHandler

import (
	"fmt"

	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type printHandler struct {
	logger.Handler
}

func New() *printHandler {
	return &printHandler{logger.Handler{lineFormatter.New()}}
}

func (h printHandler) SetFormatter(formatter logger.FormatterInterface) {
	h.SetFormatter(formatter)
}

func (h printHandler) Write(name string, level string, message string) {
	fmt.Printf(h.GetFormatter().Format(name, level, message))
}
