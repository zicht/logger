package stringutil_test

import (
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/handlers"
)

type formatter struct {
	logger.Formatter
}

func ExampleOutputDebug() {
	handler := handlers.NewStdoutHandler(logger.DEBUG)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.INFO)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.NOTICE)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.WARNING)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.ERROR)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.CRITICAL)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
	handler := handlers.NewStdoutHandler(logger.ALERT)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.ALERT: foo {}
}

func ExampleOutputEmerency() {
	handler := handlers.NewStdoutHandler(logger.EMERGENCY)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
}

func ExampleMultipleOutput() {
	handler := handlers.NewStdoutHandler(logger.EMERGENCY)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
	message := logger.NewMessage("foo")
	log := logger.NewLogger("foo", handler, handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: foo {}
	// foo.EMERGENCY: foo {}
}

func ExampleContext() {
	handler := handlers.NewStdoutHandler(logger.EMERGENCY)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
	message := logger.NewContextMessage("bar", map[string]interface{}{"bar": "foo"})
	log := logger.NewLogger("foo", handler)
	logAll(log, message)
	// Output:
	// foo.EMERGENCY: bar {"bar":"foo"}
}

func ExampleProcessor() {
	handler := handlers.NewStdoutHandler(logger.EMERGENCY)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n")
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
