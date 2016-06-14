// Package logger, it a is a simple logger inspired by monolog
// it can register multiple handlers to write the records to.
package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
	"github.com/pbergman/logger/handlers"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"github.com/pbergman/logger/debug"
)

type TraceableLoggerInterface interface {
	SetTrace(b bool)
	GetTraceDepth() int
	SetTraceDepth(v int)
}

type LoggerInterface interface {
	Emergency(message interface{})
	Alert(message interface{})
	Critical(message interface{})
	Error(message interface{})
	Warning(message interface{})
	Notice(message interface{})
	Info(message interface{})
	Debug(message interface{})
	AddProcessor(processor func(record *messages.Record))
	GetProcessors() []func(record *messages.Record)
	AddHandler(handler handlers.HandlerInterface)
	GetHandlers() []handlers.HandlerInterface
	CheckError(err error)
	CheckWarning(err error)
}

type Logger struct {
	name       string
	handlers   []handlers.HandlerInterface
	processors []func(record *messages.Record)
	mutex      sync.Mutex
	status     int
	trace      bool
	traceDepth int
}

func NewLogger(name string, handlers ...handlers.HandlerInterface) *Logger {
	return &Logger{
		name:       name,
		handlers:   handlers,
		trace:      true,
		traceDepth: 3,
	}
}

// Main function that will call the handlers and processors
func (l *Logger) log(level level.LogLevel, m interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var message *messages.Record

	switch t := m.(type) {
	case *messages.Record:
		message = t
	case error:
		message = new(messages.Record)
		message.Message = t.Error()
		message.Extra   = make(map[string]interface{}, 0)
	case string:
		message = new(messages.Record)
		message.Message = t
		message.Extra   = make(map[string]interface{}, 0)
	default:
		message = new(messages.Record)
		message.Message = fmt.Sprintf("%#v", t)
		message.Extra   = make(map[string]interface{}, 0)
	}

	if message.Time.IsZero() {
		message.Time = time.Now()
	}

	if message.Level == 0 {
		message.Level = level
	}

	if l.trace {
		message.Trace = debug.NewTrace(l.traceDepth)
	}

	for _, processor := range l.processors {
		processor(message)
	}

	for _, handler := range l.handlers {
		if handler.Support(message.Level) {
			handler.Write(l.name, level, message)
		}
	}
}

// AddProcessor add a record processor to stack that
// can edit the extra records of all messages
func (l *Logger) AddProcessor(processor func(record *messages.Record)) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.processors = append(l.processors, processor)
}

func (l *Logger) GetProcessors() []func(record *messages.Record) {
	return l.processors
}

// AddHandler adds a handler to the stack for outputting the messages
func (l *Logger) AddHandler(handler handlers.HandlerInterface) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.handlers = append(l.handlers, handler)
}

func (l *Logger) GetHandlers() []handlers.HandlerInterface {
	return l.handlers
}

// CheckError check if error is not nil if
// so it will print error message and exit
func (l *Logger) CheckError(err error) {
	if err != nil {
		l.log(level.ERROR, err)
		os.Exit(1)
	}
}

// CheckWarning is same as CheckError but will
// print warning message and will not exit
func (l *Logger) CheckWarning(err error) {
	if err != nil {
		l.log(level.WARNING, err)
	}
}

// Emergency will dispatch a log event of severity Emergency
func (l *Logger) Emergency(message interface{}) {
	l.log(level.EMERGENCY, message)
}

// Alert will dispatch a log event of severity Alert
func (l *Logger) Alert(message interface{}) {
	l.log(level.ALERT, message)
}

// Critical will dispatch a log event of severity Critical
func (l *Logger) Critical(message interface{}) {
	l.log(level.CRITICAL, message)
}

// Error will dispatch a log event of severity Error
func (l *Logger) Error(message interface{}) {
	l.log(level.ERROR, message)
}

// Warning will dispatch a log event of severity Warning
func (l *Logger) Warning(message interface{}) {
	l.log(level.WARNING, message)
}

// Notice will dispatch a log event of severity Notice
func (l *Logger) Notice(message interface{}) {
	l.log(level.NOTICE, message)
}

// Info will dispatch a log event of severity Info
func (l *Logger) Info(message interface{}) {
	l.log(level.INFO, message)
}

// Debug will dispatch a log event of severity Debug
func (l *Logger) Debug(message interface{}) {
	l.log(level.DEBUG, message)
}

func (l *Logger) SetTrace(b bool) {
	l.trace = b
}

func (l *Logger) GetTraceDepth() int {
	return l.traceDepth
}

func (l *Logger) SetTraceDepth(v int) {
	l.traceDepth = v
}