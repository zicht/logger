package logger

import (
	"io"
	"time"
	"fmt"
	"os"
	"errors"
	"runtime"
	"testing"
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
	logger := NewLogger("main", defaultHandler("main_handler", DEBUG, os.Stdout))
	logger.AddProcessor(func(r *Record){
		//pc, file, line, _ := runtime.Caller(3)
		//r.Context = map[string]interface{}{
		//	"file_name":    path.Base(file),
		//	"func_name":    runtime.FuncForPC(pc).Name(),
		//	"line_nummer":  line,
		//}

		pc, _, _, _ := runtime.Caller(5)
		r.Context = map[string]interface{}{
			"func_name":    runtime.FuncForPC(pc).Name(),
		}
	})
	logger.Debug(struct {name, message string} {"Exmaple", "Trace processor example"})
	// Output:
	// {main {Exmaple Trace processor example} map[func_name:github.com/pbergman/logger.ExampleLogger_processor] 2016-01-02 10:20:30 +0100 CET DEBUG}
}

func ExampleLogger_types() {
	logger := NewLogger("main", defaultHandler("main_handler", DEBUG, os.Stdout))
	types := []interface{}{
		"Simple string message",
		errors.New("Simple error message"),
		&Record{
			Time:    test_time,
			Channel: ChannelName("main"),
			Message: "Simple record reference message",
		},
		struct {name, message string} {"foo", "Struct message"},
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

	handler := defaultHandler("main_handler", DEBUG, os.Stdout)
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
	logger.Register("foo")

	// check lazy loading
	if (*logger.channels)["foo"] != nil {
		t.Errorf("Expecting to be nil is %T", (*logger.channels)["foo"])
	}

	logger.Get("foo")

	// check lazy loading
	if (*logger.channels)["foo"] == nil {
		t.Error("Expecting to be a Chanel instance not nil")
	}


	_, ret := logger.Get("bar")

	if ret.Error() != "Requesting a non existing channel (bar)" {
		t.Errorf("Expecting error message: 'Requesting a non existing channel (bar).' Got: %s", ret)
	}

}

func TestLogger_processor(t *testing.T) {
	p := func(record *Record){}
	logger := NewLogger("main")
	logger.AddProcessor(p)
	if len(*logger.GetProcessors()) != 1 {
		t.Errorf("Expecting to have 1 processor got: %d", len(*logger.GetHandlers()))
	}
}

func ExampleLogger_handle() {

	handler1 := defaultHandler("main_handler1", DEBUG, os.Stdout)
	handler1.channels.AddChannel(ChannelName("foo"))
	handler1.channels.AddChannel(ChannelName("main"))

	handler2 := defaultHandler("main_handler2", DEBUG, os.Stdout)
	handler2.channels.AddChannel(ChannelName("bar"))
	handler2.channels.AddChannel(ChannelName("main"))
	handler2.bubble = false

	handler3 := defaultHandler("main_handler3", DEBUG, os.Stdout)
	handler3.channels.AddChannel(ChannelName("foo"))
	handler3.channels.AddChannel(ChannelName("main"))

	logger := NewLogger("main", handler1, handler2, handler3)
	logger.Register("foo")
	logger.Register("bar")

	logger1, _ := logger.Get("foo")
	logger2, _ := logger.Get("bar")

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

func TestLogger_must(t *testing.T) {
	logger := NewLogger("main")

	catch := func(name string) (err string) {

		defer func() {
			if r := recover(); r != nil {
				err = r.(error).Error()
			}
		}()

		logger.MustGet(name)

		return
	}


	if l := catch("main"); l != "" {
		t.Errorf("Expecting to get empty string got: %s", l)
	}

	if l := catch("foo"); l != "Requesting a non existing channel (foo)" {
		t.Errorf("Expecting 'Requesting a non existing channel (foo)' got: '%s'", l)
	}

}


func TestLogger_handlers(t *testing.T) {
	logger := NewLogger("main")
	logger.AddHandler(defaultHandler("main_handler1", DEBUG, os.Stdout))
	logger.AddHandler(defaultHandler("main_handler2", DEBUG, os.Stdout))

	if len(*logger.GetHandlers()) != 2 {
		t.Errorf("Expecting to have 2 handers got: %d", len(*logger.GetHandlers()))
	}

	ret := logger.AddHandler(defaultHandler("main_handler2", DEBUG, os.Stdout))

	if ret.Error() != "A handler with name main_handler2 is allready registered." {
		t.Errorf("Expecting error message: 'A handler with name main_handler2 is allready registered.' Got: %s", ret)
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

type formatter struct { f func(r Record, w io.Writer) error }
func (f formatter) 	Format(r Record, w io.Writer) (error) { return f.f(r,w) }

type handler struct {
	name        string
	level       LogLevel
	formatter   FormatterInterface
	buff        io.Writer
	handle		func(Record, *handler) bool
	channels    *ChannelNames
	bubble		bool
}

func (h handler) GetName() string { return h.name }
func (h handler) GetLevel() LogLevel { return h.level }
func (h handler) GetFormatter() FormatterInterface { return h.formatter }
func (h handler) SetFormatter(formatter FormatterInterface) { }
func (h handler) GetChannels() *ChannelNames { return h.channels }
func (h handler) HasChannels() bool { return h.channels.Len() > 0 }
func (h handler) SetChannels(c *ChannelNames) { h.channels = c }
func (h handler) Support(record Record) bool { return h.level <= record.Level }
func (h handler) Handle(record Record) bool { return h.handle(record, &h) }

func defaultHandler(name string, level LogLevel, writer io.Writer) *handler {
	return &handler{
		name,
		level,
		&formatter{
			func(r Record, w io.Writer)error {
				r.Time = test_time;
				_, e := fmt.Fprintln(w,r);
				return e
			},
		},
		writer,
		func(r Record, h *handler) bool {
			h.GetFormatter().Format(r, h.buff)
			return h.bubble
		},
		new(ChannelNames),
		true,
	}
}