package stringutil_test

import (
	"encoding/json"
	"fmt"

	"github.com/pbergman/logger"
	"github.com/pbergman/logger/handlers"
)

type formatter struct {
	logger.Formatter
}

func New() *formatter {
	return &formatter{logger.Formatter{FormatLine: "%s.%s: %s %s\n"}}
}

func (f *formatter) Format(name string, level string, message logger.MessageInterface) string {
	return fmt.Sprintf(f.FormatLine, name, level, message.GetMessage(), toString(message.GetContext()))
}

func toString(context map[string]interface{}) string {
	json, _ := json.Marshal(context)
	return string(json)
}

func ExampleOutputDebug() {
	handler := handlers.NewPrintHandler(logger.DEBUG)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
	// foo.ERROR: foo {}
	// foo.WARNING: foo {}
	// foo.NOTICE: foo {}
	// foo.INFO: foo {}
	// foo.DEBUG: foo {}
}

func ExampleOutputInfo() {
	handler := handlers.NewPrintHandler(logger.INFO)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
	// foo.ERROR: foo {}
	// foo.WARNING: foo {}
	// foo.NOTICE: foo {}
	// foo.INFO: foo {}
}

func ExampleOutputNotice() {
	handler := handlers.NewPrintHandler(logger.NOTICE)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
	// foo.ERROR: foo {}
	// foo.WARNING: foo {}
	// foo.NOTICE: foo {}
}

func ExampleOutputWarning() {
	handler := handlers.NewPrintHandler(logger.WARNING)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
	// foo.ERROR: foo {}
	// foo.WARNING: foo {}
}

func ExampleOutputError() {
	handler := handlers.NewPrintHandler(logger.ERROR)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
	// foo.ERROR: foo {}
}

func ExampleOutputCritical() {
	handler := handlers.NewPrintHandler(logger.CRITICAL)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
	// foo.CRITICAL: foo {}
}

func ExampleOutputAlert() {
	handler := handlers.NewPrintHandler(logger.ALERT)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
}

func ExampleOutputEmerency() {
	handler := handlers.NewPrintHandler(logger.EMERGENCY)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
}

func ExampleMultipleOutput() {
	handler := handlers.NewPrintHandler(logger.EMERGENCY)
	handler.SetFormatter(New())
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo", handler, handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.EMERGENCY: foo {}
}

func ExampleContext() {
	handler := handlers.NewPrintHandler(logger.EMERGENCY)
	handler.SetFormatter(New())
	message := logger.NewContextMessage("bar", map[string]interface{}{"bar": "foo"})
	log := logger.NewLogger("foo", handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: bar {"bar":"foo"}
}

func ExampleProcessor() {
	handler := handlers.NewPrintHandler(logger.EMERGENCY)
	handler.SetFormatter(New())
	message := logger.NewContextMessage("bar", map[string]interface{}{"bar": "foo"})
	log := logger.NewLogger("foo", handler)
	log.AddProcessor(func(context map[string]interface{}) {
		context["version"] = "1.0.0"
	})
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: bar {"bar":"foo","version":"1.0.0"}
}

func logAll(l logger.LoggerInterface, m logger.MessageInterface) {
	l.Emergency(m)
	l.Alert(m)
	l.Critical(m)
	l.Error(m)
	l.Warning(m)
	l.Notice(m)
	l.Info(m)
	l.Debug(m)
}
