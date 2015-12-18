// Package logger, it a is a simple logger inspired by monolog
// it can register multiple handlers to write the records to.
package logger

import (
	"time"
	"reflect"
	"fmt"
	"sync"
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

type logger struct {
	name       string
	handlers   []HandlerInterface
	processors []func(context map[string]interface{})
	mutex	   sync.Mutex
}

func NewLogger(name string, handlers ...HandlerInterface) *logger {
	return &logger{name: name, handlers: handlers}
}

// Main function that will call the handlers and processors
func (l *logger) log(level int16, m interface{}) {
	message := l.createRecord(m)
	l.mutex.Lock()
	defer l.mutex.Unlock()
	for _, processor := range l.processors {
		processor(message.GetContext())
	}
	for _, handler := range l.handlers {
		if handler.Support(level) {
			handler.Write(l.name, levelToString(level), message)
		}
	}
}

// create a record struct from given argument
func (l *logger) createRecord(m interface{}) (message *record) {

	message = &record{extra: make(map[string]interface{},0)}

	switch m := m.(type) {
	case *record:
		message = m.(*record)
	case error:
		message = m.Error()
	case string:
		message.message = m.(string)
	default:
		message.message = fmt.Sprintf("%#v", m)
	}

	message.SetTime(time.Now())
	return
}

// AddProcessor add a record processor to stack that
// can edit the extra records of all messages
func (l *logger) AddProcessor(processor func(context map[string]interface{})) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.processors = append(l.processors, processor)
}

// AddHandler adds a hanlder to the stack for outputting the messages
func (l *logger) AddHandler(handler HandlerInterface) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.handlers = append(l.handlers, handler)
}

// Emergency will dispatch a log event of severity Emergency
func (l *logger) Emergency(message interface{}) {
	l.log(EMERGENCY, message)
}
// Alert will dispatch a log event of severity Alert
func (l *logger) Alert(message interface{}) {
	l.log(ALERT, message)
}
// Critical will dispatch a log event of severity Critical
func (l *logger) Critical(message interface{}) {
	l.log(CRITICAL, message)
}
// Error will dispatch a log event of severity Error
func (l *logger) Error(message interface{}) {
	l.log(ERROR, message)
}
// Warning will dispatch a log event of severity Warning
func (l *logger) Warning(message interface{}) {
	l.log(WARNING, message)
}
// Notice will dispatch a log event of severity Notice
func (l *logger) Notice(message interface{}) {
	l.log(NOTICE, message)
}
// Info will dispatch a log event of severity Info
func (l *logger) Info(message interface{}) {
	l.log(INFO, message)
}
// Debug will dispatch a log event of severity Debug
func (l *logger) Debug(message interface{}) {
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
