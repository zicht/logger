package logger

import (
	"strconv"
	"testing"
)

func TestRecordBuffer(t *testing.T) {
	buf := newRecordBuffer(10)

	if buf.len() != 0 {
		t.Fatalf("expected size to be 0 got %d", buf.len())
	}

	buf.push(&Record{Name: "0"})
	buf.push(&Record{Name: "1"})
	buf.push(&Record{Name: "2"})
	buf.push(&Record{Name: "3"})
	buf.push(&Record{Name: "4"})
	buf.push(&Record{Name: "5"})
	buf.push(&Record{Name: "6"})
	buf.push(&Record{Name: "7"})
	buf.push(&Record{Name: "8"})
	buf.push(&Record{Name: "9"})
	buf.push(&Record{Name: "10"})

	if buf.len() != 10 {
		t.Fatalf("expected size to be 10 got %d", buf.len())
	}

	i := 0

	for buf.valid() {
		i++
		if r := buf.shift(); r.Name != strconv.Itoa(i) {
			t.Fatalf("expected %d got %s", i, r.Name)
		}
	}

	if i != 10 {
		t.Fatalf("expected 10 iteration got %d", i)
	}
}
