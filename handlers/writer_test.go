package handlers

import (
	"bytes"
	"errors"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
	"os"
	"reflect"
	"testing"
	"time"
)

var test_time time.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)

func getRecord(m string, l logger.LogLevel, n logger.ChannelName) logger.Record {
	return logger.Record{
		Time:    test_time,
		Channel: n,
		Context: make(map[string]interface{}),
		Message: m,
		Level:   l,
	}
}

func TestWriter(t *testing.T) {
	record := getRecord("bar", logger.WARNING, logger.ChannelName("main"))
	buff := new(bytes.Buffer)
	handler := NewWriterHandler(buff, logger.INFO)
	if false == handler.Support(record) {
		t.Errorf("Handler should support level %s (%s > %s)", record.Level, record.Level, handler.level)
	}
	if true == handler.Support(getRecord("bar", logger.DEBUG, logger.ChannelName("main"))) {
		t.Errorf("Handler should not support level %s (%s < %s)", logger.DEBUG, logger.DEBUG, handler.level)
	}
	if err := handler.GetFormatter().Format(record, handler.writer); err != nil {
		t.Error(err.Error())
	}
	if buff.String() != "[2016-01-02 10:20:30.000000] main.WARNING: bar \n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] main.WARNING: bar' got: %s", buff.String())
	}
}

func TestWriter_processor(t *testing.T) {
	buff := new(bytes.Buffer)
	record := getRecord("bar", logger.WARNING, logger.ChannelName("main"))
	handler := NewWriterHandler(buff, logger.INFO)
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

func TestWriter_channel(t *testing.T) {
	buff := new(bytes.Buffer)
	record := getRecord("bar", logger.WARNING, logger.ChannelName("main"))
	handler := NewWriterHandler(buff, logger.INFO)

	if err := handler.GetChannels().AddChannel(logger.ChannelName("!main")); err != nil {
		t.Error(err.Error())
	}

	if true == handler.GetChannels().Support(record.Channel) {
		t.Errorf("Handler should not support channel %s (handler: %s)", record.Channel.GetName(), (*handler.channels)[handler.channels.FindChannel("main")])
	}

	if false == handler.GetChannels().Support(logger.ChannelName("test")) {
		t.Errorf("Handler should support channel %s (handler: %s)", record.Channel.GetName(), (*handler.channels)[handler.channels.FindChannel("main")])
	}
}

func TestWriter_close(t *testing.T) {
	writer := &testWriter{}
	handler := NewWriterHandler(writer, logger.INFO)

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
	// no io.Closer writer
	handler.writer = &testWriterNoClose{}
	if err := handler.Close(); err != nil {
		t.Errorf("Expecting to get nil error got %#v", err)
	}
}

type testWriterNoClose struct{ e error }

func (w testWriterNoClose) Write(p []byte) (n int, err error) { return 0, nil }

type testWriter struct{ e error }

func (w testWriter) Write(p []byte) (n int, err error) { return 0, nil }
func (w testWriter) Close() error                      { return w.e }

func TestWriter_channels_not_nil(t *testing.T) {
	handler := &WriterHandler{}
	if handler.GetChannels() == nil {
		t.Errorf("Expecting channels not to be nil")
	}
}

func TestWriter_channe(t *testing.T) {
	handler := NewWriterHandler(
		os.Stdout,
		logger.INFO,
		logger.ChannelName("foo"),
		logger.ChannelName("bar"),
		logger.ChannelName("example"),
	)

	if handler.GetChannels().Len() != 3 {
		t.Errorf("Expecting 3 channels got %d", handler.GetChannels().Len())
	}
}

func TestWriter_getters(t *testing.T) {
	writer := testWriter{}
	handler := NewWriterHandler(&writer, logger.INFO)

	if handler.GetLevel() != logger.INFO {
		t.Errorf("Expecting level 'INFO' got:  %s", logger.INFO)
	}

	formatter := formatters.NewCustomLineFormatter("{{ .Message }}")
	handler.SetFormatter(formatter)

	if !reflect.DeepEqual(formatter, handler.GetFormatter()) {
		t.Errorf("Expecting to get same formatter as was set (%p != %p)", formatter, handler.GetFormatter())
	}

	if true == handler.HasChannels() {
		t.Errorf("Expecting not to have any channels got %d", handler.GetChannels().Len())
	}

	channels := new(logger.ChannelNames)
	channels.AddChannel(logger.ChannelName("foo"))
	channels.AddChannel(logger.ChannelName("bar"))
	handler.SetChannels(channels)

	if !reflect.DeepEqual(channels, handler.GetChannels()) {
		t.Errorf("Expecting to get same channel colection as was set (%p != %p)", formatter, handler.GetFormatter())
	}

	if false == handler.HasChannels() && handler.GetChannels().Len() != 2 {
		t.Errorf("Expecting not to have 2 channels got %d", handler.GetChannels().Len())
	}

	if true != handler.GetBubble() {
		t.Errorf("Expecting default bubble to be true got %t", handler.GetBubble())
	}

	if ret := handler.Handle(&logger.Record{}); ret != true {
		t.Errorf("Expecting propagate to be true got %t", ret)
	}

	handler.SetBubble(false)

	if false != handler.GetBubble() {
		t.Errorf("Expecting default bubble to be true got %t", handler.GetBubble())
	}

	if ret := handler.Handle(&logger.Record{}); ret != false {
		t.Errorf("Expecting propagate to be false got %t", ret)
	}

	// break template formatter with bad syntax!
	handler.SetFormatter(formatters.NewCustomLineFormatter("{{*}}"))

	err := func() (err string) {

		defer func() {
			if r := recover().(string); r != "" {
				err = r
			}
		}()

		handler.Handle(&logger.Record{})

		return
	}()

	if err == "" {
		t.Errorf("Expecting to have error about template syntax.")
	}

}
