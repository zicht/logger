package logger

import (
	"fmt"
	"sync"
	"time"
	"errors"
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
	Handlers        map[string]HandlerInterface
	Processors      []func(record *Record)
	Channels        map[string]*Channel
)

func NewLogger(name string, handlers ...HandlerInterface) *Logger {

	handler := make(Handlers, 0)

	for _, h := range handlers {
		handler[h.GetName()] = h
	}

	channels := make(Channels, 0)

	return &Logger{
		name:       name,
		channels:   &channels,
		processors: new(Processors),
		handlers:   &handler,
	}
}

func (l *Logger) AddProcessor(processor func(record *Record)) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	(*l.processors) = append(*l.processors, processor)
}

func (l *Logger) GetProcessors() *Processors {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.processors
}

// AddHandler adds a handler to the stack for outputting the messages
func (l *Logger) AddHandler(handler HandlerInterface) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if _, o := (*l.handlers)[handler.GetName()]; o {
		return errors.New("A handler with name " + handler.GetName() + " is allready registered.")
	}

	(*l.handlers)[handler.GetName()] = handler
	return nil
}

func (l *Logger) GetHandlers() *Handlers {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.handlers
}

func (l *Logger) Register(name string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if _, o := (*l.channels)[name]; !o {
		(*l.channels)[name] = nil
	} else {
		return errors.New("Channel " + name + " is allready registered.")
	}
	return nil
}

func (l *Logger) Get(name string) (LoggerInterface, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if v, o := (*l.channels)[name]; o {
		if v == nil {
			(*l.channels)[name] = &Channel{
				logger: l,
				name:   ChannelName(name),
			}
			v = (*l.channels)[name]
		}
		return v, nil
	} else {
		return nil, errors.New("Requesting a non existing channel (" + name + ")")
	}
}

func (l *Logger) handle(record Record) {

	for _, processor := range (*l.processors) {
		processor(&record)
	}

	for _, handler := range (*l.handlers) {
		if handler.HasChannels() && false == handler.GetChannels().Support(record.Channel) {
			continue
		}
		if handler.Support(record) {
			if false == handler.Handle(record) {
				break
			}
		}
	}
}

func (l *Logger) log(level LogLevel, channel ChannelName, message interface{}) {

	l.mutex.Lock()
	defer l.mutex.Unlock()

	var record Record

	switch value := message.(type) {
	case *Record:
		record = *value
	case Record:
		record = value
	case string:
		record = Record{Message: value}
	case error:
		record = Record{Message: value.Error()}
	default:
		record = Record{Message: fmt.Sprint(value)}
	}

	record.Time = time.Now()
	record.Level = level
	record.Channel = channel
	l.handle(record)
}

func (l *Logger) Emergency(message interface{}) {
	l.log(EMERGENCY, ChannelName(l.name), message)
}

func (l *Logger) Alert(message interface{}) {
	l.log(ALERT, ChannelName(l.name), message)
}

func (l *Logger) Critical(message interface{}) {
	l.log(CRITICAL, ChannelName(l.name), message)
}

func (l *Logger) Error(message interface{}) {
	l.log(ERROR, ChannelName(l.name), message)
}

func (l *Logger) Warning(message interface{}) {
	l.log(WARNING, ChannelName(l.name), message)
}

func (l *Logger) Notice(message interface{}) {
	l.log(NOTICE, ChannelName(l.name), message)
}

func (l *Logger) Info(message interface{}) {
	l.log(INFO, ChannelName(l.name), message)
}

func (l *Logger) Debug(message interface{}) {
	l.log(DEBUG, ChannelName(l.name), message)
}