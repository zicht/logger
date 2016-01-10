package logger

import (
	"time"
)

type MessageInterface interface {
	AddContext(name string, value interface{})
	SetContext(context map[string]interface{})
	SetTime(t time.Time)
	GetContext() map[string]interface{}
	GetMessage() string
	GetTrace() *Trace
	GetTime() time.Time
}

type Record struct {
	message string
	extra   map[string]interface{}
	time    time.Time
	level   int16
	trace   *Trace
}

func (r *Record) AddContext(name string, value interface{}) {
	r.extra[name] = value
}

func (r *Record) SetContext(c map[string]interface{}) {
	r.extra = c
}

func (r *Record) SetLevel(l int16) {
	r.level = l
}

func (m Record) GetLevel() int16 {
	return m.level
}

func (r *Record) SetTrace(t *Trace) {
	r.trace = t
}

func (r *Record) HasTrace() bool {
	return r.trace != nil
}

func (r *Record) GetTrace() *Trace {
	return r.trace
}

func (m *Record) GetContext() map[string]interface{} {
	return m.extra
}

func (r *Record) SetTime(t time.Time) {
	r.time = t
}

func (r *Record) GetTime() time.Time {
	return r.time
}

func (r *Record) GetMessage() string {
	return r.message
}

// NewMessage creates a message without any context, context
// can be added later by calling AddContext and SetContext methods
func NewMessage(m string) *Record {
	return &Record{
		message: m,
		extra:   make(map[string]interface{}),
	}
}

// NewContextMessage creates a message with context
func NewContextMessage(m string, context map[string]interface{}) *Record {
	return &Record{
		message: m,
		extra:   context,
	}
}
