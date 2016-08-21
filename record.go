package logger

import "time"

type Record struct {
	Channel ChannelName     // channel who pushed the record
	Message string          // log message
	Context interface{}     // log context
	Time    time.Time       // time when message is pushed
	Level   LogLevel        // level of the log message
}
