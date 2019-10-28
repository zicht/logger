package formatters

import (
	"testing"
	"text/template"
	"time"

	"github.com/zicht/logger"
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
	fm := NewLineFormatter()
	out, err := fm.Format(record)
	if err != nil {
		t.Error(err)
	}
	if string(out) != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar \n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] foo.CRITICAL: bar' got: %s", buff.String())
	}
	// check for not null printing
	record.Context = nil
	out, err = fm.Format(record)

	if err != nil {
		t.Error(err)
	}
	if string(out) != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar \n" {
		t.Errorf("Expecting '[2016-01-02 10:20:30.000000] foo.CRITICAL: bar' got: %s", buff.String())
	}
	// check context printing
	record.Context = map[string]string{"test": "test_context"}
	out, err = fm.Format(record)
	if err != nil {
		t.Error(err)
	}
	if string(out) != "[2016-01-02 10:20:30.000000] foo.CRITICAL: bar {\"test\":\"test_context\"}\n" {
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
	fm := NewCustomLineFormatter("{{.Bar}}")
	_, err := fm.Format(record)
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

	fm := NewCustomLineFormatter("{{*}}")
	_, err := fm.Format(record)

	if err.Error() != "template: line_formatter:1: unexpected \"*\" in command" {
		t.Errorf("Unexpected error message: ", err.Error())

	}

}
