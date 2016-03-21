// Package logger, it a is a simple logger inspired by monolog
// it can register multiple handlers to write the records to.
package logger

import (
	"fmt"
	"github.com/pbergman/logger/handlers"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"os"
	"sync"
	"time"
)

const (
	STATE_RUNNING = 1 << iota
	STATE_PAUSED
)

type LoggerInterface interface {
	Emergency(message interface{})
	Alert(message interface{})
	Critical(message interface{})
	Error(message interface{})
	Warning(message interface{})
	Notice(message interface{})
	Info(message interface{})
	Debug(message interface{})
}

type Logger struct {
	name       string
	handlers   []handlers.HandlerInterface
	processors []func(record *messages.Record)
	mutex      sync.Mutex
	status     int
	queue      *messages.Queue
	Trace      bool
}

func NewLogger(name string, handlers ...handlers.HandlerInterface) *Logger {
	return &Logger{
		name:     name,
		handlers: handlers,
		status:   STATE_RUNNING,
		Trace:    true,
	}
}

func (l *Logger) dispatch(record *messages.Record) {
	for _, processor := range l.processors {
		processor(record)
	}

	for _, handler := range l.handlers {
		if handler.Support(record.GetLevel()) {
			handler.Write(l.name, record.GetLevel(), record)
		}
	}
}

func (l *Logger) GetState() int {
	return l.status
}

// Pause will puase the output and capture the log records
// in a queue limited to the given buffer size, when resumed
// the records will be dispatched
func (l *Logger) Pause(bufferSize int) {
	l.mutex.Lock()
	l.queue = messages.NewQueue(bufferSize)
	l.status = STATE_PAUSED
	l.mutex.Unlock()
}

func (l *Logger) Resume() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for l.queue.Valid() {
		l.dispatch(l.queue.Pop())
	}
	l.queue = nil
	l.status = STATE_RUNNING
}

// Main function that will call the handlers and processors
func (l *Logger) log(level level.LogLevel, m interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var message = new(messages.Record)
	switch m.(type) {
	case *messages.Record:
		message.SetMessage(m.(*messages.Record).GetMessage())
		message.SetContext(m.(*messages.Record).GetContext())
	case error:
		message.SetMessage(m.(error).Error())
		message.SetContext(make(map[string]interface{}, 0))
	case string:
		message.SetMessage(m.(string))
		message.SetContext(make(map[string]interface{}, 0))
	default:
		message.SetMessage(fmt.Sprintf("%#v", m))
		message.SetContext(make(map[string]interface{}, 0))
	}

	message.SetTime(time.Now())
	message.SetLevel(level)
	if l.Trace {
		message.SetTrace(messages.NewTrace())
	}

	switch l.status {
	case STATE_RUNNING:
		l.dispatch(message)
	case STATE_PAUSED:
		l.queue.Push(message)
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
