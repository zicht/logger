package logger

import "time"

type MessageInterface interface {
	AddContext(name string, value interface{})
	SetContext(context map[string]interface{})
	SetTime(t time.Time)
	GetContext() map[string]interface{}
	GetMessage() string
	GetTime() time.Time
}

type contextMessage struct {
	message string
	context map[string]interface{}
	time    time.Time
}

func (c *contextMessage) AddContext(name string, value interface{}) {
	c.context[name] = value
}

func (c *contextMessage) SetContext(context map[string]interface{}) {
	c.context = context
}

func (c *contextMessage) SetTime(t time.Time) {
	c.time = t
}

func (c contextMessage) GetTime() time.Time {
	return c.time
}

func (c contextMessage) GetContext() map[string]interface{} {
	return c.context
}

func (c contextMessage) GetMessage() string {
	return c.message
}

func NewMessage(message string) *contextMessage {
	return &contextMessage{message: message, context: make(map[string]interface{}, 1), time: time.Now()}
}

func NewContextMessage(message string, context map[string]interface{}) *contextMessage {
	return &contextMessage{message: message, context: context, time: time.Now()}
}
