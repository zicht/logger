// Package logger, it a is a simple logger inspired by monolog
// it can register multiple handlers to write the records to.
package logger

import (
	"fmt"
	"sync"
	"time"
	"os"
)

const (
	// Levels as described by http://tools.ietf.org/html/rfc5424
	EMERGENCY int16 = 600
	ALERT     int16 = 550
	CRITICAL  int16 = 500
	ERROR     int16 = 400
	WARNING   int16 = 300
	NOTICE    int16 = 250
	INFO      int16 = 200
	DEBUG     int16 = 100

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
	handlers   []HandlerInterface
	processors []func(record *Record)
	mutex      sync.Mutex
	status     int
	queue      *queue
}

func NewLogger(name string, handlers ...HandlerInterface) *Logger {
	return &Logger{
		name:     name,
		handlers: handlers,
		status:   STATE_RUNNING,
	}
}

func (l *Logger) dispatch(record *Record) {
	for _, processor := range l.processors {
		processor(record)
	}

	for _, handler := range l.handlers {
		if handler.Support(record.GetLevel()) {
			handler.Write(l.name, levelToString(record.GetLevel()), record)
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
	l.queue = NewQueue(bufferSize)
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
func (l *Logger) log(level int16, m interface{}) {

	l.mutex.Lock()
	defer l.mutex.Unlock()

	message := &Record{extra: make(map[string]interface{}, 0)}

	switch m.(type) {
	case *Record:
		message.message = m.(*Record).GetMessage()
		message.extra = m.(*Record).GetContext()
	case error:
		message.message = m.(error).Error()
	case string:
		message.message = m.(string)
	default:
		message.message = fmt.Sprintf("%#v", m)
	}

	message.SetTime(time.Now())
	message.SetLevel(level)
	message.SetTrace(NewTrace())

	switch l.status {
	case STATE_RUNNING:
		l.dispatch(message)
	case STATE_PAUSED:
		l.queue.Push(message)
	}
}

// AddProcessor add a record processor to stack that
// can edit the extra records of all messages
func (l *Logger) AddProcessor(processor func(record *Record)) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.processors = append(l.processors, processor)
}

func (l *Logger) GetProcessors() []func(record *Record) {
	return l.processors
}

// AddHandler adds a handler to the stack for outputting the messages
func (l *Logger) AddHandler(handler HandlerInterface) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.handlers = append(l.handlers, handler)
}

func (l *Logger) GetHandlers() []HandlerInterface {
	return l.handlers
}

// CheckError check if error is not nil if
// so it will print error message and exit
func (l *Logger) CheckError(err error) {
	if err != nil {
		l.log(ERROR, err)
		os.Exit(1)
	}
}

// CheckWarning is same as CheckError but will
// print warning message and will not exit
func (l *Logger) CheckWarning(err error) {
	if err != nil {
		l.log(WARNING, err)
	}
}

// Emergency will dispatch a log event of severity Emergency
func (l *Logger) Emergency(message interface{}) {
	l.log(EMERGENCY, message)
}

// Alert will dispatch a log event of severity Alert
func (l *Logger) Alert(message interface{}) {
	l.log(ALERT, message)
}

// Critical will dispatch a log event of severity Critical
func (l *Logger) Critical(message interface{}) {
	l.log(CRITICAL, message)
}

// Error will dispatch a log event of severity Error
func (l *Logger) Error(message interface{}) {
	l.log(ERROR, message)
}

// Warning will dispatch a log event of severity Warning
func (l *Logger) Warning(message interface{}) {
	l.log(WARNING, message)
}

// Notice will dispatch a log event of severity Notice
func (l *Logger) Notice(message interface{}) {
	l.log(NOTICE, message)
}

// Info will dispatch a log event of severity Info
func (l *Logger) Info(message interface{}) {
	l.log(INFO, message)
}

// Debug will dispatch a log event of severity Debug
func (l *Logger) Debug(message interface{}) {
	l.log(DEBUG, message)
}

func levelToString(level int16) (name string) {
	switch level {
	case EMERGENCY:
		return string("EMERGENCY")
	case ALERT:
		return string("ALERT")
	case CRITICAL:
		return string("CRITICAL")
	case ERROR:
		return string("ERROR")
	case WARNING:
		return string("WARNING")
	case NOTICE:
		return string("NOTICE")
	case INFO:
		return string("INFO")
	case DEBUG:
		return string("DEBUG")
	default:
		return string("UNKNOWN")
	}
}
