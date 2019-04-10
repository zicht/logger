package logger

import (
	"testing"
)

type nopHandler struct {
	bubble  bool
	handle  bool
	handled bool
}

func (n nopHandler) IsHandling(r *Record) bool {
	return n.handle
}

func (n *nopHandler) Handle(r *Record) bool {
	n.handled = true
	return !n.bubble
}

func TestFormatHandler(t *testing.T) {
	handler := new(handlers)
	handler.SetHandlers(
		&nopHandler{false, false, false},
		&nopHandler{false, false, false},
	)
	if i := handler.isHandling(&Record{}); i != -1 {
		t.Fatalf("expected to have -1 returned got %d", i)
	}
	handlers := handler.GetHandlers()
	if l := len(handlers); l != 2 {
		t.Fatalf("expected to have 2 handlers now got %d", l)
	}
	handler.handle(&Record{})
	if handlers[0].(*nopHandler).handled {
		t.Fatalf("expected handler with index 0 not to be run")
	}
	if handlers[1].(*nopHandler).handled {
		t.Fatalf("expected handler with index 1 not to be run")
	}
	handler.AddHandlers(&nopHandler{true, true, false})
	if i := handler.isHandling(&Record{}); i != 2 {
		t.Fatalf("expected to have 2 returned got %d", i)
	}
	handler.AddHandlers(&nopHandler{true, true, false})
	handler.AddHandlers(&nopHandler{false, true, false})
	handler.AddHandlers(&nopHandler{false, true, false})
	handlers = handler.GetHandlers()
	if l := len(handlers); l != 6 {
		t.Fatalf("expected to have 6 handlers now got %d", l)
	}
	handler.handle(&Record{})
	if handlers[0].(*nopHandler).handled {
		t.Fatalf("expected handler with index 0 not to be run")
	}
	if handlers[1].(*nopHandler).handled {
		t.Fatalf("expected handler with index 1 not to be run")
	}
	if !handlers[2].(*nopHandler).handled {
		t.Fatalf("expected handler with index 2 to be run")
	}
	if !handlers[3].(*nopHandler).handled {
		t.Fatalf("expected handler with index 3 to be run")
	}
	if !handlers[4].(*nopHandler).handled {
		t.Fatalf("expected handler with index 4 to be run")
	}
	if handlers[5].(*nopHandler).handled {
		t.Fatalf("expected handler with index 5 not to be run")
	}
}
