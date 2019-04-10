package logger

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestThresholdHandler(t *testing.T) {
	recordTime := time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
	buf := new(bytes.Buffer)
	handler := NewThresholdHandler(NewWriterHandler(buf, Debug, false), 10, Error, false)
	if !handler.IsHandling(nil) {
		t.Fatal("expected to always handle records")
	}
	handler.Handle(&Record{Level: Debug, Name: "test", Time: recordTime, Message: "test message"})
	handler.Handle(&Record{Level: Info, Name: "test", Time: recordTime, Message: "test message"})
	if buf.Len() != 0 {
		t.Fatal("not records should have been dispatched")
	}
	handler.Handle(&Record{Level: Error, Name: "test", Time: recordTime, Message: "test message"})
	if buf.Len() == 0 {
		t.Fatal("records should have been dispatched")
	}
	expected := "[2016-01-02 10:20:30.000000] test.DEBUG: test message\n[2016-01-02 10:20:30.000000] test.INFO: test message\n[2016-01-02 10:20:30.000000] test.ERROR: test message\n"
	if buf.String() != expected {
		t.Fatalf("\nExpected:\n%q\nGot:\n%q", expected, buf.String())
	}
}

type testCloserWrapper struct {
	io.Writer
	isClosed bool
}

func (t *testCloserWrapper) Close() error {
	t.isClosed = true
	return nil
}

func TestThresholdHandler_Close(t *testing.T) {
	handler := NewThresholdHandler(NewWriterHandler(&testCloserWrapper{Writer: new(bytes.Buffer)}, Debug, false), 10, Error, false)
	handler.(*thresholdHandler).Close()
	if handler.(*thresholdHandler).handler.(*writerHandler).writer.(*nopWriterCloser).Writer.(*testCloserWrapper).isClosed {
		t.Fatalf("expecte to have closed wrapped handler")
	}
}

func ExampleNewThresholdHandler() {

	handler := NewThresholdHandler(NewWriterHandler(os.Stdout, Info, false), 4, Error, false)
	logger := NewLogger("app", handler)

	defer logger.Close()

	logger.PushProcessor(P(func(r *Record) {
		// set the time static so we can test output
		r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
	}))

	logger.Debug("test")
	logger.Info("test")
	logger.Error("test")

	// Output:
	// [2016-01-02 10:20:30.000000] app.INFO: test
	// [2016-01-02 10:20:30.000000] app.ERROR: test
}
