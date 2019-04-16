package logger

import (
	"fmt"
	"io"
	"time"
)

func NewLogger(name string, handler ...HandlerInterface) *Logger {
	return &Logger{name: name, handlers: handlers{processor: processor{processors: make([]ProcessorInterface, 0)}, handlers: handler}}
}

type Logger struct {
	name string
	handlers
}

func (l *Logger) log(level LogLevel, name string, message interface{}) {
	var record *Record
	switch value := message.(type) {
	case LogMessageInterface:
		message, context := value.GetLogMessage()
		record = &Record{Message: message, Context: context}
	case *Record:
		record = value
	case Record:
		record = &value
	case string:
		record = &Record{Message: value}
	case error:
		record = &Record{Message: value.Error()}
	default:
		record = &Record{Message: fmt.Sprint(value)}
	}
	record.Time = time.Now()
	record.Level = level
	record.Name = name
	l.handle(record)
}

// WithName will return an new logger with the given name and shared internal handler
func (l *Logger) WithName(name string) *Logger {
	return &Logger{name: name, handlers: l.handlers}
}

func (l *Logger) Close() error {
	var err = new(Errors)
	for _, h := range l.handlers.handlers {
		if v, o := h.(io.Closer); o {
			if e := v.Close(); e != nil {
				err.append(e)
			}
		}
	}
	for _, p := range l.processors {
		if v, o := p.(io.Closer); o {
			if e := v.Close(); e != nil {
				err.append(e)
			}
		}
	}
	return err.GetError()
}

// NewWriter creates an wrapper that can be used as an io.Writer to print all given
// message to the supplied level
//
// example:
//
// buf := new(bytes.Buffer)
// writer := io.MultiWriter(buf, logger.NewWriter(logger.Debug))
//
// extending internal log:
//
// log.SetFlags(0)
// log.SetOutput(logger.NewWriter(logger.Debug))
//
func (l *Logger) NewWriter(level LogLevel) io.Writer {
	return &writer{logger: l, level: level}
}

func (l Logger) Emergency(message interface{}) { l.log(Emergency, l.name, message) }
func (l Logger) Alert(message interface{})     { l.log(Alert, l.name, message) }
func (l Logger) Critical(message interface{})  { l.log(Critical, l.name, message) }
func (l Logger) Error(message interface{})     { l.log(Error, l.name, message) }
func (l Logger) Warning(message interface{})   { l.log(Warning, l.name, message) }
func (l Logger) Notice(message interface{})    { l.log(Notice, l.name, message) }
func (l Logger) Info(message interface{})      { l.log(Info, l.name, message) }
func (l Logger) Debug(message interface{})     { l.log(Debug, l.name, message) }
