package messages

import (
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/debug"
	"time"
)

type Record struct {
	Message string
	Extra   map[string]interface{}
	Time    time.Time
	Level   level.LogLevel
	Trace   *debug.Trace
}


// NewMessage creates a message without any context, context
// can be added later by calling AddContext and SetContext methods
func NewMessage(m string) *Record {
	return &Record{
		Message: m,
		Extra:   make(map[string]interface{}),
	}
}

// NewContextMessage creates a message with context
func NewContextMessage(m string, context map[string]interface{}) *Record {
	return &Record{
		Message: m,
		Extra:   context,
	}
}
