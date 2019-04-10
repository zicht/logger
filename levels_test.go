package logger

import (
	"testing"
)

func TestLogLevel(t *testing.T) {
	levels := map[int]string{
		0x80: "DEBUG",
		0x40: "INFO",
		0x20: "NOTICE",
		0x10: "WARNING",
		0x08: "ERROR",
		0x04: "CRITICAL",
		0x02: "ALERT",
		0x01: "EMERGENCY",
		0x00: "UNKNOWN",
	}

	for i, n := range levels {
		if LogLevel(i).String() != n {
			t.Errorf("Expecting (0x%02x) %s got %s", i, n, LogLevel(i))
		}
	}

	if v := LogLevel(Debug ^ Alert).String(); v != "ALERT" {
		t.Fatalf("expected ALERT got %s", v)
	}

	if !LogLevel(Debug ^ Alert).Match(Alert) {
		t.Fatal("expected to match Alert level")
	}

	if !LogLevel(Debug ^ Alert).Match(Debug) {
		t.Fatal("expected to match Debug level")
	}

	if LogLevel(Debug & Alert).Has(Info) {
		t.Fatalf("expected not to have Info level")
	}
}
