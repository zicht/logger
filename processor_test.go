package logger

import (
	"bytes"
	"testing"
)

func TestProcessor(t *testing.T) {

	o := new(bytes.Buffer)
	p := new(processor)
	a := P(func(r *Record) { o.WriteString("A") })
	b := P(func(r *Record) { o.WriteString("B") })
	p.PushProcessor(&a)
	p.PushProcessor(&b)

	p.process(&Record{})

	if s := o.String(); s != "AB" {
		t.Fatalf("expected 'AB' got '%s'", s)
	}

	if l := len(p.processors); l != 2 {
		t.Fatalf("expected to have 2 processors got %d", l)
	}

	if f := p.PopProcessors(); f != &b {
		t.Fatalf("expected %#v got %#v", &b, f)
	}

	if l := len(p.processors); l != 1 {
		t.Fatalf("expected to have 1 processors got %d", l)
	}

	if f := p.PopProcessors(); f != &a {
		t.Fatalf("expected %#v got %#v", &a, f)
	}

	if l := len(p.processors); l != 0 {
		t.Fatalf("expected to have 0 processors got %d", l)
	}

	if c := p.PopProcessors(); c != nil {
		t.Fatalf("expected nil got %T", c)
	}
}
