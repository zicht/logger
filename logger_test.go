package logger_test

import (
	"fmt"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/handlers"
	"github.com/pbergman/logger/processors"
)

func ExampleNewLogger() {
	handler := handlers.NewStdoutHandler(logger.DEBUG)
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }}\n")
	message := logger.NewMessage("DEBUG")
	log := logger.NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	handler.Level = logger.INFO
	message = logger.NewMessage("INFO")
	logAll(log, message)
	handler.Level = logger.NOTICE
	message = logger.NewMessage("NOTICE")
	logAll(log, message)
	handler.Level = logger.WARNING
	message = logger.NewMessage("WARNING")
	logAll(log, message)
	handler.Level = logger.ERROR
	message = logger.NewMessage("ERROR")
	logAll(log, message)
	handler.Level = logger.CRITICAL
	message = logger.NewMessage("CRITICAL")
	logAll(log, message)
	handler.Level = logger.ALERT
	message = logger.NewMessage("ALERT")
	logAll(log, message)
	handler.Level = logger.EMERGENCY
	message = logger.NewMessage("EMERGENCY")
	logAll(log, message)
	// Output:
	//foo.EMERGENCY: DEBUG
	//foo.ALERT: DEBUG
	//foo.CRITICAL: DEBUG
	//foo.ERROR: DEBUG
	//foo.WARNING: DEBUG
	//foo.NOTICE: DEBUG
	//foo.INFO: DEBUG
	//foo.DEBUG: DEBUG
	//foo.EMERGENCY: INFO
	//foo.ALERT: INFO
	//foo.CRITICAL: INFO
	//foo.ERROR: INFO
	//foo.WARNING: INFO
	//foo.NOTICE: INFO
	//foo.INFO: INFO
	//foo.EMERGENCY: NOTICE
	//foo.ALERT: NOTICE
	//foo.CRITICAL: NOTICE
	//foo.ERROR: NOTICE
	//foo.WARNING: NOTICE
	//foo.NOTICE: NOTICE
	//foo.EMERGENCY: WARNING
	//foo.ALERT: WARNING
	//foo.CRITICAL: WARNING
	//foo.ERROR: WARNING
	//foo.WARNING: WARNING
	//foo.EMERGENCY: ERROR
	//foo.ALERT: ERROR
	//foo.CRITICAL: ERROR
	//foo.ERROR: ERROR
	//foo.EMERGENCY: CRITICAL
	//foo.ALERT: CRITICAL
	//foo.CRITICAL: CRITICAL
	//foo.EMERGENCY: ALERT
	//foo.ALERT: ALERT
	//foo.EMERGENCY: EMERGENCY
}

func ExampleAddProcessor() {
	handler := handlers.NewStdoutHandler(logger.DEBUG)
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ json false .extra }}\n")
	trace := processors.NewTraceProcessor(logger.INFO)
	logger := logger.NewLogger("test", handler)
	logger.AddProcessor(trace.Process)
	logger.Debug("foo")
	logger.Info("foo")
	// Output:
	//test.DEBUG: foo {}
	//test.INFO: foo {"file":"logger_test.go","line":86}
}

//Example Basic illustration of using logger
func Example() {
	// A logger that will display all levels to out.txt and from level WARNING or higher to stderr
	log := logger.NewLogger(
		"test",
		handlers.NewStderrHandler(logger.WARNING),
		handlers.NewFileHandler("out.txt", logger.DEBUG),
	)
	log.Debug("this would only be displayed in file")
	log.Warning("this would be displayed in file and on stderr")
}

func ExamplePause() {
	handler := handlers.NewStdoutHandler(logger.DEBUG)
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }}\n")
	logger := logger.NewLogger("test", handler)
	logger.Pause(5)
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	logger.Debug("foo")
	fmt.Println("some text")
	logger.Resume()
	logger.Debug("bar")
	logger.Debug("bar")
	logger.Debug("bar")
	// Output:
	//some text
	//test.DEBUG: foo
	//test.DEBUG: foo
	//test.DEBUG: foo
	//test.DEBUG: foo
	//test.DEBUG: foo
	//test.DEBUG: bar
	//test.DEBUG: bar
	//test.DEBUG: bar
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
