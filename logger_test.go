package logger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"testing"
	"time"
)

var test_time time.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)

func getRecord(m string, n ChannelName) Record {
	return Record{
		Time:    test_time,
		Channel: n,
		Message: m,
	}
}

func ExampleLogger_processor() {
	logger := NewLogger("main", defaultHandler(DEBUG, os.Stdout))
	logger.AddProcessor(func(r *Record) {
		//pc, file, line, _ := runtime.Caller(3)
		//r.Context = map[string]interface{}{
		//	"file_name":    path.Base(file),
		//	"func_name":    runtime.FuncForPC(pc).Name(),
		//	"line_nummer":  line,
		//}

		pc, _, _, _ := runtime.Caller(5)
		r.Context = map[string]interface{}{
			"func_name": runtime.FuncForPC(pc).Name(),
		}
	})
	logger.Debug(struct{ name, message string }{"Exmaple", "Trace processor example"})
	// Output:
	// {main {Exmaple Trace processor example} map[func_name:github.com/pbergman/logger.ExampleLogger_processor] 2016-01-02 10:20:30 +0100 CET DEBUG}
}

func ExampleLogger_cm() {
	logger := NewLogger("main", defaultHandler(DEBUG, os.Stdout))
	logger.Debug(ContextMessage("Foo", map[string]interface{}{"one": 1}))

	// Output:
	// {main Foo map[one:1] 2016-01-02 10:20:30 +0100 CET DEBUG}
}

func ExampleLogger_types() {
	logger := NewLogger("main", defaultHandler(DEBUG, os.Stdout))
	types := []interface{}{
		"Simple string message",
		errors.New("Simple error message"),
		&Record{
			Time:    test_time,
			Channel: ChannelName("main"),
			Message: "Simple record reference message",
		},
		struct{ name, message string }{"foo", "Struct message"},
	}

	for _, t := range types {
		logger.Debug(t)
	}

	// Output:
	// {main Simple string message <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main Simple error message <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main Simple record reference message <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main {foo Struct message} <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}

}

func ExampleLogger_levels() {

	handler := defaultHandler(DEBUG, os.Stdout)
	logger := NewLogger("main", handler)
	levels := [9]int{100, 200, 250, 300, 400, 500, 550, 600, 199}

	for _, l := range levels {
		handler.level = LogLevel(l)
		logAll(logger, getRecord(fmt.Sprintf("Exmaple level %s", LogLevel(l)), ChannelName("main")))
	}

	// Output:
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET INFO}
	// {main Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {main Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET INFO}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {main Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {main Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {main Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level ALERT <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level ALERT <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level EMERGENCY <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {main Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET INFO}

}

func TestLogger_channel(t *testing.T) {
	logger := NewLogger("main")
	logger.Get("foo")

	// check lazy loading
	if logger.channels.c["foo"] == nil {
		t.Error("Expecting to be a Chanel instance not nil")
	}
}

func TestLogger_processor(t *testing.T) {
	p := func(record *Record) {}
	logger := NewLogger("main")
	logger.AddProcessor(p)
	if len(*logger.GetProcessors()) != 1 {
		t.Errorf("Expecting to have 1 processor got: %d", len(*logger.GetHandlers()))
	}
}

func ExampleLogger_handle() {

	handler1 := defaultHandler(DEBUG, os.Stdout)
	handler1.channels.AddChannel(ChannelName("foo"))
	handler1.channels.AddChannel(ChannelName("main"))

	handler2 := defaultHandler(DEBUG, os.Stdout)
	handler2.channels.AddChannel(ChannelName("bar"))
	handler2.channels.AddChannel(ChannelName("main"))
	handler2.bubble = false

	handler3 := defaultHandler(DEBUG, os.Stdout)
	handler3.channels.AddChannel(ChannelName("foo"))
	handler3.channels.AddChannel(ChannelName("main"))

	logger := NewLogger("main", handler1, handler2, handler3)
	logger1 := logger.Get("foo")
	logger2 := logger.Get("bar")

	logger1.Debug("one")
	logger2.Debug("two")
	logger.Debug("hello")

	// Output:
	// {foo one <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {foo one <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {bar two <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main hello <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {main hello <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
}

func TestLogger_close(t *testing.T) {

	logger := NewLogger("main")

	if err := logger.Close(); err != nil {
		t.Errorf("Expecting empty error got: %d", err)
	}

	logger.AddHandler(defaultHandler(DEBUG, os.Stdout))

	if err := logger.Close(); err != nil {
		t.Errorf("Expecting empty error got: %d", err)
	}

	(*logger.handlers)[0].(*handler).err = errors.New("foo")

	if err := logger.Close(); err == nil {
		t.Error("Expecting a error")
	} else {

		if str := err.Error(); str != "foo" {
			t.Errorf("Expecting 'foo' got %s", str)
		}
	}
}

func TestLogger_handlers(t *testing.T) {
	logger := NewLogger("main")
	logger.AddHandler(defaultHandler(DEBUG, os.Stdout))
	logger.AddHandler(defaultHandler(DEBUG, os.Stdout))

	if len(*logger.GetHandlers()) != 2 {
		t.Errorf("Expecting to have 2 handers got: %d", len(*logger.GetHandlers()))
	}
}

func logAll(l LoggerInterface, r Record) {
	l.Emergency(r)
	l.Alert(r)
	l.Critical(r)
	l.Error(r)
	l.Warning(r)
	l.Notice(r)
	l.Info(r)
	l.Debug(r)
}

type formatter struct {
	f func(r Record) ([]byte, error)
}

func (f formatter) Format(r Record) ([]byte, error) { return f.f(r) }

type handler struct {
	level     LogLevel
	formatter FormatterInterface
	buff      io.Writer
	handle    func(*Record, *handler) bool
	channels  *ChannelNames
	bubble    bool
	err       error
}

func (h handler) GetLevel() LogLevel                        { return h.level }
func (h handler) GetFormatter() FormatterInterface          { return h.formatter }
func (h handler) SetFormatter(formatter FormatterInterface) {}
func (h handler) GetChannels() *ChannelNames                { return h.channels }
func (h handler) HasChannels() bool                         { return h.channels.Len() > 0 }
func (h handler) SetChannels(c *ChannelNames)               { h.channels = c }
func (h handler) Support(record Record) bool                { return h.level <= record.Level }
func (h handler) Handle(record *Record) bool                { return h.handle(record, &h) }
func (h handler) AddProcessor(Processor)                    {}
func (h handler) GetProcessors() *Processors                { return nil }
func (h handler) Close() error                              { return h.err }

func defaultHandler(level LogLevel, writer io.Writer) *handler {
	return &handler{
		level: level,
		formatter: &formatter{
			func(r Record) ([]byte, error) {
				buf := new(bytes.Buffer)
				r.Time = test_time
				_, e := fmt.Fprintln(buf, r)
				return buf.Bytes(), e
			},
		},
		buff: writer,
		handle: func(r *Record, h *handler) bool {
			b, _ := h.GetFormatter().Format(*r)
			h.buff.Write(b)
			return h.bubble
		},
		channels: new(ChannelNames),
		bubble:   true,
	}
}
