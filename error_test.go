package logger

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {

	err := new(ErrorStack)

	if len := err.Len(); len > 0 {
		t.Errorf("Expecting to have 0 error elements got %d", len)
	}

	if str := err.Error(); str != "" {
		t.Errorf("Expecting '' got: %s", str)
	}

	err.Add(errors.New("First  error"))
	err.Add(errors.New("Second error"))

	if len := err.Len(); len != 2 {
		t.Errorf("Expecting to have 2 error elements got %d", len)
	}

	err.Add(errors.New("Extra  error"))

	if str := err.Error(); str != "First  error\nSecond error\nExtra  error" {
		t.Errorf("Expecting 'First  error\nSecond error\nExtra  error' got: %s", str)
	}
}
