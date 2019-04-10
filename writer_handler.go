package logger

import (
	"io"
)

type nopWriterCloser struct {
	io.Writer
}

func (nopWriterCloser) Close() error { return nil }

// NewWriterHandler is an handler that writes the log entries to the given
// writer and wrapes the writer around an nop io.WriterCloser so when the
// close method is invoked it will not close the writer when it also
// implements the io,WriterCloser
func NewWriterHandler(w io.Writer, level LogLevel, bubble bool) HandlerInterface {
	return NewWriteCloserHandler(&nopWriterCloser{w}, level, bubble)
}

// NewWriteCloserHandler is similar as the NewWriterHandler but expects an
// io.WriteCloser where the close method will be invoked when the logger
// is closed
func NewWriteCloserHandler(w io.WriteCloser, level LogLevel, bubble bool) HandlerInterface {
	formatter, _ := NewFormatHandler(nil)
	return &writerHandler{
		writer: w,
		baseHandler: baseHandler{
			level:  level,
			bubble: bubble,
		},
		FormatHandler: formatter,
		processor:     processor{processors: make([]ProcessorInterface, 0)},
	}
}

type writerHandler struct {
	writer io.WriteCloser
	baseHandler
	processor
	FormatHandler
}

func (w writerHandler) Close() error {
	return w.writer.Close()
}

func (w writerHandler) IsHandling(r *Record) bool {
	return r.Level.Match(w.level)
}

func (w writerHandler) Handle(r *Record) bool {
	if !w.IsHandling(r) {
		return false
	}
	w.process(r)
	buf, err := w.GetFormatter().Format(r)
	if err != nil {
		w.writer.Write([]byte("err: " + err.Error()))
	} else {
		w.writer.Write(buf)
	}
	return w.bubble
}
