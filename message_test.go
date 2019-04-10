package logger

import "testing"

func TestMessage(t *testing.T) {
	message := Message("foo", make(map[string]interface{}))
	msg, ctx := message.GetLogMessage()
	if ctx != nil {
		t.Fatalf("expected to have nil context got %T", ctx)
	}
	if "foo" != msg {
		t.Fatalf("expected foo as message got %s", msg)
	}
	ctx = map[string]interface{}{
		"bar": "hello",
		"foo": "world",
	}
	message = Message("foo", ctx)
	_, ctx = message.GetLogMessage()
	if ctx == nil {
		t.Fatalf("expected to have context got %T", ctx)
	}
	if _, o := ctx["bar"]; !o {
		t.Fatalf("expected to have bar key in ctx")
	}
	if _, o := ctx["foo"]; !o {
		t.Fatalf("expected to have foo key in ctx")
	}
}
