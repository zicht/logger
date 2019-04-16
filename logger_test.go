package logger

import (
    "os"
    "time"
    "testing"
    "errors"
    "bytes"
)

type errCloser struct { err error }
func (n errCloser) Close() error { return n.err }
type errWriter struct { errCloser }
func (e errWriter) Write(p []byte) (n int, err error) { return len(p), nil}
type errProcessor struct { errCloser }
func (e errProcessor) Process(record *Record) { }

func TestLogger_Close_error(t *testing.T) {
    logger := NewLogger("app", NewWriteCloserHandler(&errWriter{errCloser{errors.New("foo")}}, Critical, false))
    logger.PushProcessor(errProcessor{errCloser{errors.New("bar")}})
    err := logger.Close()
    if s := len(*err.(*Errors)); s != 2 {
        t.Fatalf("expected 2 errors got %d", s)
    }
    if err.Error() != "2 errors occurred:\n\tfoo\n\tbar\n" {
        t.Fatalf("expected:\n%q\ngot:\n%q", "2 errors occurred:\n\tfoo\n\tbar\n", err.Error())
    }
}

func TestLogger_Log(t *testing.T) {
    buf := new(bytes.Buffer)
    logger := NewLogger("app", NewWriterHandler(buf, Debug, false))
    defer logger.Close()
    logger.PushProcessor(P(func(r *Record) {
        // set the time static so we can test output
        r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
    }))
    record := Record{
        Message: "test message",
        Level: Warning,
    }
    logger.Debug(record)
    expected := "[2016-01-02 10:20:30.000000] app.DEBUG: test message\n"
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
    buf.Truncate(0)
    logger.Debug(&record)
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
    buf.Truncate(0)
    logger.Debug(errors.New("test message"))
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
    buf.Truncate(0)
    logger.Debug(Message("test message", nil))
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
    buf.Truncate(0)
    logger.Debug(Message("test message", nil))
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
    buf.Truncate(0)
    logger.Debug(112)
    expected = "[2016-01-02 10:20:30.000000] app.DEBUG: 112\n"
    if o := buf.String(); o != expected {
        t.Fatalf("expected:\n%s\ngot:\n%s", expected, o)
    }
}

func ExampleNewLogger() {
    logger := NewLogger("app", NewWriterHandler(os.Stdout, Critical, false))
    defer logger.Close()
    logger.PushProcessor(P(func(r *Record) {
        // set the time static so we can test output
        r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
    }))
    logger.Debug("Should not be printed")
    logger.Info("Should not be printed")
    logger.Notice("Should not be printed")
    logger.Warning("Should not be printed")
    logger.Error("Should not be printed")
    logger.Critical("Should be printed")
    logger.Alert("Should be printed")
    logger.Emergency("Should be printed")
    // Output:
    // [2016-01-02 10:20:30.000000] app.CRITICAL: Should be printed
    // [2016-01-02 10:20:30.000000] app.ALERT: Should be printed
    // [2016-01-02 10:20:30.000000] app.EMERGENCY: Should be printed

}

func ExampleLogger_WithName() {
    logger := NewLogger("foo", NewWriterHandler(os.Stdout, Debug, false)).WithName("bar")
    defer logger.Close()
    logger.PushProcessor(P(func(r *Record) {
        // set the time static so we can test output
        r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
    }))
    logger.Debug("hello world")
    // Output:
    // [2016-01-02 10:20:30.000000] bar.DEBUG: hello world
}