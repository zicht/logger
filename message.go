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
	GetTime() time.Time
}

type record struct {
	message string
	extra	map[string]interface{}
	time	time.Time
}

func (r *record) AddContext(name string, value interface{}) {
	r.extra[name] = value
}

func (r *record) SetContext(c map[string]interface{}) {
	r.extra = c
}

func (m record) GetContext() map[string]interface{} {
	return m.extra
}

func (r *record) SetTime(t time.Time) {
	r.time = t
}

func (r record) GetTime() time.Time {
	return r.time
}

func (r record) GetMessage() string {
	return r.message
}

// Create a message without any context,
// context can be added later by calling
// AddContext and SetContext methods
func NewMessage(m string) *record {
	return &record{
		message: m,
		extra:	 make(map[string]interface{}),
		time:	 time.Now(),
	}
}

// Create a message with context
func NewContextMessage(m string, context map[string]interface{}) *record {
	return &record{
		message: m,
		extra:	 context,
		time:	 time.Now(),
	}
}
