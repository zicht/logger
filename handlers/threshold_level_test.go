package handlers

import (
	"bytes"
	"errors"
	"github.com/pbergman/logger"
	"io"
	"testing"
	"os"
	"github.com/pbergman/logger/formatters"
)

type formatter struct{}

func (f formatter) Format(r logger.Record, w io.Writer) error {
	_, e := w.Write([]byte(r.Level.String() + "\n"))
	return e
}

func TestThresholdLevel(t *testing.T) {
	buffer := new(bytes.Buffer)
	record := getRecord("bar", logger.DEBUG, logger.ChannelName("main"))
	handler := NewThresholdLevelHandler(NewWriterHandler(buffer, logger.DEBUG), logger.ERROR, 5)
	handler.handler.SetFormatter(&formatter{})
	if true != handler.Support(record) {
		t.Errorf("Expecting to support record %#v", record)
	}
	if true != handler.IsStopBuffering() {
		t.Error("Expecting default stop buffering to be true")
	}
	if true != handler.IsBuffering() {
		t.Error("Expecting default buffering to be true")
	}
	handler.Handle(&record)
	if s := buffer.Len(); s > 0 {
		t.Error("Expecting handler not to be called.")
	}
	for i := 0; i < 10; i++ {
		handler.Handle(&record)
	}
	if s := len(handler.buffer); s != 5 {
		t.Errorf("Expecting buffer size not to exceed 5, size: %d, cap %d.", len(handler.buffer), cap(handler.buffer))
	}
	nr := getRecord("foo", logger.ERROR, logger.ChannelName("main"))
	handler.Handle(&nr)
	if str := buffer.String(); str != "DEBUG\nDEBUG\nDEBUG\nDEBUG\nERROR\n" {
		t.Errorf("Expecting: 'DEBUG\nDEBUG\nDEBUG\nDEBUG\nERROR\n' got: %s", str)
	}
	handler.Handle(&record)
	if s := len(handler.buffer); s > 0 {
		t.Errorf("Expecting record not to be bufferd (s:%d).", s)
	}
	handler.SetStopBuffering(false)
	for i := 0; i < 10; i++ {
		handler.Handle(&record)
	}
	if s := len(handler.buffer); s != 5 {
		t.Errorf("Expecting buffer size not to exceed 5, size: %d, cap %d.", len(handler.buffer), cap(handler.buffer))
	}
	handler.Clear()
	if s := len(handler.buffer); s != 0 {
		t.Errorf("Expecting buffer size to be 0, size: %d, cap %d.", len(handler.buffer), cap(handler.buffer))
	}
}

func TestThresholdLevel_close(t *testing.T) {
	writer := &testWriter{}
	handler := NewThresholdLevelHandler(NewWriterHandler(writer, logger.DEBUG), logger.ERROR, 5)

	if err := handler.Close(); err != nil {
		t.Errorf("Expecting to get nil error got %#v", err)
	}
	writer.e = errors.New("foo")
	if err := handler.Close(); err == nil {
		t.Error("Expecting to get a error")
	} else {
		if str := err.Error(); str != "foo" {
			t.Errorf("Expecting 'foo' got: %s", str)
		}
	}
	// no io.Closer handler
	handler.handler = &stub_handler{}
	if err := handler.Close(); err != nil {
		t.Errorf("Expecting to get nil error got %#v", err)
	}
}

func TestThresholdLevel_processor(t *testing.T) {
	buff := new(bytes.Buffer)
	record := getRecord("bar", logger.WARNING, logger.ChannelName("main"))
	handler := NewThresholdLevelHandler(NewWriterHandler(buff, logger.DEBUG), logger.ERROR, 5)
	handler.AddProcessor(func(r *logger.Record) {
		r.Channel = logger.ChannelName("foo")
	})

	if handler.GetProcessors().Len() <= 0 {
		t.Errorf("Expecting to have 1 processor got %d", handler.GetProcessors().Len())
	}

	handler.Handle(&record)

	if record.Channel.GetName() != "foo" {
		t.Errorf("Expecting record to have channel name 'foo' got: %s", record.Channel.GetName())
	}
}

func TestThresholdLevel_channel(t *testing.T) {
	buff := new(bytes.Buffer)
	record := getRecord("bar", logger.WARNING, logger.ChannelName("main"))
	handler := NewThresholdLevelHandler(NewWriterHandler(buff, logger.DEBUG), logger.ERROR, 5, logger.ChannelName("!main"))

	if true == handler.GetChannels().Support(record.Channel) {
		t.Errorf("Handler should not support channel %s (handler: %s)", record.Channel.GetName(), (*handler.channels)[handler.channels.FindChannel("main")])
	}

	if false == handler.GetChannels().Support(logger.ChannelName("test")) {
		t.Errorf("Handler should support channel %s (handler: %s)", record.Channel.GetName(), (*handler.channels)[handler.channels.FindChannel("main")])
	}
}

func ExampleThresholdLevelHandler_no_output() {
	handler := NewWriterHandler(os.Stdout, logger.DEBUG)
	handler.SetFormatter(formatters.NewCustomLineFormatter("{{.Channel | printf \"%-4s\" }} [{{ .Level | printf \"%-8s\" }}] :: {{ .Message }}\n"))
	logwriter := logger.NewLogger("app", NewThresholdLevelHandler(handler,logger.CRITICAL, 10))
	// buffer again when threshold is reached
	(*logwriter.GetHandlers())[0].(*ThresholdLevelHandler).SetStopBuffering(false)
	// some random logging in channels app, bar and main
	logwriter.Error("foo")
	logwriter.Warning("foo")
	logwriter.Notice("foo")
	logwriter.Info("foo")
	logwriter.Debug("foo")
	// Output:
}

func ExampleThresholdLevelHandler() {
	handler := NewWriterHandler(os.Stdout, logger.INFO)
	handler.SetFormatter(formatters.NewCustomLineFormatter("{{.Channel | printf \"%-4s\" }} [{{ .Level | printf \"%-8s\" }}] :: {{ .Message }}\n"))
	logwriter := logger.NewLogger("app", NewThresholdLevelHandler(handler,logger.CRITICAL, 10))
	// buffer again when threshold is reached
	(*logwriter.GetHandlers())[0].(*ThresholdLevelHandler).SetStopBuffering(false)
	// some random logging in channels app, bar and main
	logwriter.Error("foo")
	logwriter.Error("foo")
	logwriter.Warning("foo")
	logwriter.Warning("foo")
	logwriter.Notice("foo")
	logwriter.Notice("foo")
	logwriter.Info("foo")
	logwriter.Info("foo")
	logwriter.Debug("foo")
	logwriter.Debug("foo")
	// should output now
	logwriter.Critical("foo")
	// Output:
	// app  [ERROR   ] :: foo
	// app  [WARNING ] :: foo
	// app  [WARNING ] :: foo
	// app  [NOTICE  ] :: foo
	// app  [NOTICE  ] :: foo
	// app  [INFO    ] :: foo
	// app  [INFO    ] :: foo
	// app  [CRITICAL] :: foo

}

type stub_handler struct{}

func (h stub_handler) GetLevel() logger.LogLevel                        { return logger.DEBUG }
func (h stub_handler) GetFormatter() logger.FormatterInterface          { return nil }
func (h stub_handler) SetFormatter(formatter logger.FormatterInterface) {}
func (h stub_handler) GetChannels() *logger.ChannelNames                { return nil }
func (h stub_handler) HasChannels() bool                                { return false }
func (h stub_handler) SetChannels(c *logger.ChannelNames)               {}
func (h stub_handler) Support(record logger.Record) bool                { return true }
func (h stub_handler) Handle(record *logger.Record) bool                { return true }
func (h stub_handler) AddProcessor(logger.Processor)                    {}
func (h stub_handler) GetProcessors() *logger.Processors                { return nil }
