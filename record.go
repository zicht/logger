package logger

import "time"

type Record struct {
	Name    string
	Message string
	Context map[string]interface{}
	Time    time.Time
	Level   LogLevel
}
