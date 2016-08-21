package formatters

import (
	"testing"
	"time"
	"bytes"
	"github.com/pbergman/logger"
	"text/template"
)

var test_time time.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)

func TestLine(t *testing.T) {
	record := logger.Record{
		Time:    test_time,
		Channel: logger.ChannelName("foo"),
		Context: map[string]string{},
		Message: "bar",
		Level:   logger.CRITICAL,
	}
	buff := new(bytes.Buffer)
	fm := NewLineFormatter()
	if err := fm.Format(record, buff); err != nil {
		t.Error(err)
	}
	if buff.String() != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar \n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] foo.CRITICAL: bar' got: %s", buff.String())
	}

	// check for not null printing
	buff.Truncate(0)
	record.Context = nil
	if err := fm.Format(record, buff); err != nil {
		t.Error(err)
	}
	if buff.String() != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar \n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] foo.CRITICAL: bar' got: %s", buff.String())
	}
	// check context printing
	buff.Truncate(0)
	record.Context = map[string]string{"test": "test_context"}
	if err := fm.Format(record, buff); err != nil {
		t.Error(err)
	}
	if buff.String() != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar {\"test\":\"test_context\"}\n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] foo.CRITICAL: bar {\"test\":\"test_context\"}' got: %s", buff.String())
	}
}

func TestLine_exec_error(t *testing.T) {
	record := logger.Record{
		Time:    test_time,
		Channel: logger.ChannelName("foo"),
		Message: "bar",
		Level:   logger.CRITICAL,
	}
	buff := new(bytes.Buffer)
	fm := NewCustomLineFormatter("{{.Bar}}")
	err := fm.Format(record, buff)
	if _, o := err.(template.ExecError); !o {
		t.Errorf("Expecting template.ExecError got %T", err)
	}
}

func TestLine_error(t *testing.T) {
	record := logger.Record{
		Time:    test_time,
		Channel: logger.ChannelName("foo"),
		Message: "bar",
		Level:   logger.CRITICAL,
	}
	buff := new(bytes.Buffer)
	fm := NewCustomLineFormatter("{{*}}")
	err := fm.Format(record, buff)

	if err.Error() != "template: line_formatter:1: unexpected \"*\" in command" {
		t.Errorf("Unexpected error message: ", err.Error())

	}

}
