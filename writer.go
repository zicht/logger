package logger

// writer is an wrapper around an logger that implement
// the io.Writer and will log to the defined level
type writer struct {
	logger *Logger
	level  LogLevel
}

// Write implements io,Write by dispatching the given
// the bytes to the logger as an string
func (w writer) Write(b []byte) (int, error) {
	w.logger.log(w.level, w.logger.name, string(b))
	return len(b), nil
}
