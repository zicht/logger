package logger

import (
	"testing"
)

func TestLogLevel(t *testing.T) {
	levels := map[int]string{
		100: "DEBUG",
		200: "INFO",
		250: "NOTICE",
		300: "WARNING",
		400: "ERROR",
		500: "CRITICAL",
		550: "ALERT",
		600: "EMERGENCY",
		199: "UNKNOWN",

	}

	for i, n := range levels {
		if LogLevel(i).String() != n {
			t.Errorf("Expecting %s got %s", LogLevel(i), n)
		}
	}

}
