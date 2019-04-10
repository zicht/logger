package logger

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	var err Errors

	if err.GetError() != nil {
		t.Fatalf("expected to get nil")
	}

	err.append(errors.New("hello"))

	if err.GetError() == nil {
		t.Fatalf("expected not to get nil")
	}

	if s := err.Error(); s != "1 error occurred:\n\thello\n" {
		t.Fatalf("expected:\n%q\nGot\n%q", "1 error occurred:\n\thello\n", s)
	}

	err.append(errors.New("world"))

	if s := err.Error(); s != "2 errors occurred:\n\thello\n\tworld\n" {
		t.Fatalf("expected:\n%q\nGot\n%q", "2 errors occurred:\n\thello\n\tworld\n", s)
	}
}
