package logger

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func TestWriterHandler_IsHandling(t *testing.T) {
	file, err := os.OpenFile(os.DevNull, os.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	var all = LogLevel(255)
	shouldHandle(t, NewWriterHandler(file, Debug, false).(*writerHandler), all)
	shouldHandle(t, NewWriterHandler(file, Info, false).(*writerHandler), all^Debug)
	shouldHandle(t, NewWriterHandler(file, Notice, false).(*writerHandler), all^Debug^Info)
	shouldHandle(t, NewWriterHandler(file, Warning, false).(*writerHandler), all^Debug^Info^Notice)
	shouldHandle(t, NewWriterHandler(file, Error, false).(*writerHandler), all^Debug^Info^Notice^Warning)
	shouldHandle(t, NewWriterHandler(file, Critical, false).(*writerHandler), all^Debug^Info^Notice^Warning^Error)
	shouldHandle(t, NewWriterHandler(file, Alert, false).(*writerHandler), all^Debug^Info^Notice^Warning^Error^Critical)
	shouldHandle(t, NewWriterHandler(file, Emergency, false).(*writerHandler), all^Debug^Info^Notice^Warning^Error^Critical^Alert)

}

func shouldHandle(t *testing.T, handler *writerHandler, level LogLevel) {
	r := &Record{}
	for i := 1; i <= 128; i <<= 1 {
		r.Level = LogLevel(i)
		if i == (i & int(level)) {
			if !handler.IsHandling(r) {
				t.Fatalf("expected to %s to be handld for %s", r.Level, handler.level)
			}
		} else {
			if handler.IsHandling(r) {
				t.Fatalf("expected to %s not to be handld for %s", r.Level, handler.level)
			}
		}
	}
}

func TestWriterHandler_Handle(t *testing.T) {
	buf := new(bytes.Buffer)
	handler := NewWriterHandler(buf, Info, true)
	handler.(FormatHandlerInterface).SetFormatter(F(func(r *Record) ([]byte, error) {
		return []byte(r.Level.String()), nil
	}))
	if handler.Handle(&Record{Level: Debug}) {
		t.Fatalf("Should have returned false")
	}
	if !handler.Handle(&Record{Level: Info}) {
		t.Fatalf("Should have returned true")
	}
	if o := buf.String(); o != "INFO" {
		t.Fatalf("expected 'INFO' got '%s'", o)
	}
	buf.Truncate(0)
	handler.(FormatHandlerInterface).SetFormatter(F(func(r *Record) ([]byte, error) {
		return nil, errors.New("foo")
	}))
	if !handler.Handle(&Record{Level: Info}) {
		t.Fatalf("Should have returned true")
	}
	if o := buf.String(); o != "err: foo" {
		t.Fatalf("expected 'err: foo' got '%s'", o)
	}
}

type closeWrapper struct {
	io.Writer
	isClosed bool
}

func (c *closeWrapper) Close() error {
	c.isClosed = true
	return nil
}

func TestNewWriterHandler(t *testing.T) {
	file, err := os.OpenFile(os.DevNull, os.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	wrapped := closeWrapper{Writer: file}

	handler := NewWriterHandler(wrapped, Debug, true)
	handler.(*writerHandler).Close()
	if wrapped.isClosed {
		t.Fatal("expected not to be closed")
	}
	handler = NewWriteCloserHandler(&wrapped, Debug, true)
	handler.(*writerHandler).Close()
	if !wrapped.isClosed {
		t.Fatal("expected to be closed")
	}
}
