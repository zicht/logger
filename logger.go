// Package logger, it a is a simple logger inspired by monolog
// it can register multiple handlers to write the records to.
package logger

import "time"

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
	Emergency(message MessageInterface)
	Alert(message MessageInterface)
	Critical(message MessageInterface)
	Error(message MessageInterface)
	Warning(message MessageInterface)
	Notice(message MessageInterface)
	Info(message MessageInterface)
	Debug(message MessageInterface)
}

type logger struct {
	name       string
	handlers   []HandlerInterface
	processors []func(context map[string]interface{})
}

func NewLogger(name string, handlers ...HandlerInterface) *logger {
	return &logger{name: name, handlers: handlers}
}

func (l *logger) log(level int16, message MessageInterface) {

	message.SetTime(time.Now())

	for _, processor := range l.processors {
		processor(message.GetContext())
	}

	for _, handler := range l.handlers {
		if handler.Support(level) {
			handler.Write(l.name, levelToString(level), message)
		}
	}
}

func (l *logger) AddProcessor(processor func(context map[string]interface{})) {
	l.processors = append(l.processors, processor)
}

func (l *logger) AddHandler(handler HandlerInterface) {
	l.handlers = append(l.handlers, handler)
}

func (l *logger) Emergency(message MessageInterface) {
	l.log(EMERGENCY, message)
}

func (l *logger) Alert(message MessageInterface) {
	l.log(ALERT, message)
}

func (l *logger) Critical(message MessageInterface) {
	l.log(CRITICAL, message)
}

func (l *logger) Error(message MessageInterface) {
	l.log(ERROR, message)
}

func (l *logger) Warning(message MessageInterface) {
	l.log(WARNING, message)
}

func (l *logger) Notice(message MessageInterface) {
	l.log(NOTICE, message)
}

func (l *logger) Info(message MessageInterface) {
	l.log(INFO, message)
}

func (l *logger) Debug(message MessageInterface) {
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
