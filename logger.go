package logger

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type (
	LoggerInterface interface {
		Emergency(message interface{})
		Alert(message interface{})
		Critical(message interface{})
		Error(message interface{})
		Warning(message interface{})
		Notice(message interface{})
		Info(message interface{})
		Debug(message interface{})
	}
	Logger struct {
		name       string
		channels   *Channels
		handlers   *Handlers
		processors *Processors
		mutex      sync.RWMutex
	}
	Channels struct {
		c map[string]*Channel
		m sync.RWMutex
	}
)

func NewLogger(name string, handlers ...HandlerInterface) *Logger {
	handler := new(Handlers)
	for _, h := range handlers {
		(*handler) = append(*handler, h)
	}
	channels := new(Channels)
	channels.c = make(map[string]*Channel, 1)
	channels.c[name] = nil
	return &Logger{
		name:       name,
		channels:   channels,
		processors: new(Processors),
		handlers:   handler,
	}
}

func (l *Logger) AddProcessor(processor Processor) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	(*l.processors) = append(*l.processors, processor)
}

func (l *Logger) GetProcessors() *Processors {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.processors
}

func (l *Logger) AddHandler(handler HandlerInterface) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	(*l.handlers) = append(*l.handlers, handler)
	return nil
}

func (l *Logger) GetHandlers() *Handlers {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.handlers
}

func (l *Logger) Get(name string) LoggerInterface {
	l.channels.m.Lock()
	defer l.channels.m.Unlock()
	var channel *Channel
	var exist bool
	if channel, exist = l.channels.c[name]; !exist || channel == nil {
		channel = &Channel{
			logger: l,
			name:   ChannelName(name),
		}
		l.channels.c[name] = channel
	}
	return channel
}

func (l *Logger) handle(record *Record) {

	if l.processors.Len() > 0 {
		for _, k := range l.processors.Keys() {
			(*l.processors)[k](record)
		}
	}
	if l.handlers.Len() > 0 {
		for _, k := range l.handlers.Keys() {
			if (*l.handlers)[k].HasChannels() && false == (*l.handlers)[k].GetChannels().Support(record.Channel) {
				continue
			}
			if (*l.handlers)[k].Support(*record) {
				if !(*l.handlers)[k].Handle(record) {
					break
				}
			}
		}
	}
}

func (l *Logger) log(level LogLevel, channel ChannelName, message interface{}) {

	l.mutex.Lock()
	defer l.mutex.Unlock()

	var record *Record

	if value, ok := message.(LogMessageInterface); ok {
		message, context := value.GetLogMessage()
		record = &Record{Message: message, Context: context}
	} else {
		switch value := message.(type) {
		case *Record:
			record = value
		case Record:
			record = &value
		case *cm:
			record = &Record{Message: value.m, Context: value.c}
		case string:
			record = &Record{Message: value}
		case error:
			record = &Record{Message: value.Error()}
		default:
			record = &Record{Message: fmt.Sprint(value)}
		}
	}

	record.Time = time.Now()
	record.Level = level
	record.Channel = channel
	l.handle(record)
}

func (l *Logger) Close() error {
	var errStack *ErrorStack = nil
	if l.handlers.Len() > 0 {
		for _, k := range l.handlers.Keys() {
			if closer, ok := (*l.handlers)[k].(io.Closer); ok {
				if err := closer.Close(); err != nil {
					if errStack == nil {
						errStack = new(ErrorStack)
					}
					errStack.Add(err)
				}
			}
		}
	}
	// explicit return nil else returns 0 in logger_test.go:224
	if errStack != nil {
		return errStack
	} else {
		return nil
	}
}

func (l *Logger) Emergency(message interface{}) {
	l.Get(l.name).Emergency(message)
}

func (l *Logger) Alert(message interface{}) {
	l.Get(l.name).Alert(message)
}

func (l *Logger) Critical(message interface{}) {
	l.Get(l.name).Critical(message)
}

func (l *Logger) Error(message interface{}) {
	l.Get(l.name).Error(message)
}

func (l *Logger) Warning(message interface{}) {
	l.Get(l.name).Warning(message)
}

func (l *Logger) Notice(message interface{}) {
	l.Get(l.name).Notice(message)
}

func (l *Logger) Info(message interface{}) {
	l.Get(l.name).Info(message)
}

func (l *Logger) Debug(message interface{}) {
	l.Get(l.name).Debug(message)
}
