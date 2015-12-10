package lineFormatter

import (
	"fmt"
	"time"

	"github.com/pbergman/logger"
)

type lineFormatter struct {
	logger.Formatter
}

func New() *lineFormatter {
	return &lineFormatter{logger.Formatter{Line: "[%s] %s.%s %s\n"}}
}

func (f *lineFormatter) getTime() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (f *lineFormatter) Format(name string, level string, message string) string {
	return fmt.Sprintf(f.Line, f.getTime(), name, level, message)
}
