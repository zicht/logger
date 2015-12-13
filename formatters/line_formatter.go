package formatters

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pbergman/logger"
)

type lineFormatter struct {
	logger.Formatter
}

func NewLineFormatter() *lineFormatter {
	return &lineFormatter{logger.Formatter{FormatLine: "[%s] %s.%s: %s %s\n"}}
}

func (f *lineFormatter) formatTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (f *lineFormatter) Format(name string, level string, message logger.MessageInterface) string {
	return fmt.Sprintf(f.FormatLine, f.formatTime(message.GetTime()), name, level, message.GetMessage(), toString(message.GetContext()))
}

func toString(context map[string]interface{}) string {
	json, _ := json.Marshal(context)
	return string(json)
}
