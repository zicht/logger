package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestNewFormatHandler(t *testing.T) {
	formatter, err := NewFormatHandler(nil)

	if err != nil {
		t.Fatal(err.Error())
	}

	out, err := formatter.GetFormatter().Format(&Record{
		Name:    "test",
		Message: "hello world",
		Context: make(map[string]interface{}),
		Time:    time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local),
		Level:   Debug,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	if string(out) != "[2016-01-02 10:20:30.000000] test.DEBUG: hello world\n" {
		t.Fatalf("\nExpected:\n%q\nGot:\n%q", "[2016-01-02 10:20:30.000000] test.DEBUG: hello world\n", string(out))
	}

	record := &Record{
		Name:    "test",
		Message: "hello world",
		Context: map[string]interface{}{"memory": 1024},
		Time:    time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local),
		Level:   Error,
	}

	out, err = formatter.GetFormatter().Format(record)

	if err != nil {
		t.Fatal(err.Error())
	}

	if string(out) != "[2016-01-02 10:20:30.000000] test.ERROR: hello world {\"memory\":1024}\n" {
		t.Fatalf("\nExpected:\n%q\nGot:\n%q", "[2016-01-02 10:20:30.000000] test.ERROR: hello world {\"memory\":1024}\n", string(out))
	}

	formatter.SetFormatter(F(func(r *Record) ([]byte, error) {
		return nil, errors.New("test")
	}))

	_, err = formatter.formatter.Format(record)

	if err == nil {
		t.Fatal("expected an error")
	}

}

func ExampleNewFormatHandler() {
	formatter, _ := NewFormatHandler(F(func(r *Record) ([]byte, error) {
		return json.Marshal(r)
	}))
	record := &Record{
		Name:    "test",
		Message: "hello world",
		Context: make(map[string]interface{}),
		Time:    time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local),
		Level:   Debug,
	}
	out, _ := formatter.formatter.Format(record)
	fmt.Println(string(out))
	// Output:
	// {"Name":"test","Message":"hello world","Context":{},"Time":"2016-01-02T10:20:30+01:00","Level":128}
}

func TestNewLineFormatter_error(t *testing.T) {
	_, err := NewLineFormatter("{{")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTextTemplateFormatter_Format(t *testing.T) {
	handler, _ := NewFormatHandler(nil)
	if _, err := handler.formatter.Format(nil); err == nil {
		t.Fatal("expected error")
	}
}