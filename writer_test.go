package logger

import (
	"os"
	"time"
)

func ExampleWriter() {

	logger := NewLogger("app", NewWriterHandler(os.Stdout, Notice, false))
	logger.PushProcessor(P(func(r *Record) {
		r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
	}))

	writer := logger.NewWriter(Warning)
	writer.Write([]byte("hello world"))

	// Output:
	// [2016-01-02 10:20:30.000000] app.WARNING: hello world

}
