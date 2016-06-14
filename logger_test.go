package logger

import (
	"github.com/pbergman/logger/handlers"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"github.com/pbergman/logger/processors"
	"io"
	"os"
	"testing"
)

func ExampleNewLogger() {
	handler := handlers.NewWriterHandler(os.Stdout, level.DEBUG)
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }}\n")
	message := messages.NewMessage("DEBUG")
	log := NewLogger("foo")
	log.AddHandler(handler)
	logAll(log, message)
	handler.Level = level.INFO
	message = messages.NewMessage("INFO")
	logAll(log, message)
	handler.Level = level.NOTICE
	message = messages.NewMessage("NOTICE")
	logAll(log, message)
	handler.Level = level.WARNING
	message = messages.NewMessage("WARNING")
	logAll(log, message)
	handler.Level = level.ERROR
	message = messages.NewMessage("ERROR")
	logAll(log, message)
	handler.Level = level.CRITICAL
	message = messages.NewMessage("CRITICAL")
	logAll(log, message)
	handler.Level = level.ALERT
	message = messages.NewMessage("ALERT")
	logAll(log, message)
	handler.Level = level.EMERGENCY
	message = messages.NewMessage("EMERGENCY")
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
	//foo.DEBUG: INFO
	//foo.EMERGENCY: NOTICE
	//foo.ALERT: NOTICE
	//foo.CRITICAL: NOTICE
	//foo.ERROR: NOTICE
	//foo.WARNING: NOTICE
	//foo.NOTICE: NOTICE
	//foo.INFO: NOTICE
	//foo.DEBUG: NOTICE
	//foo.EMERGENCY: WARNING
	//foo.ALERT: WARNING
	//foo.CRITICAL: WARNING
	//foo.ERROR: WARNING
	//foo.WARNING: WARNING
	//foo.NOTICE: WARNING
	//foo.INFO: WARNING
	//foo.DEBUG: WARNING
	//foo.EMERGENCY: ERROR
	//foo.ALERT: ERROR
	//foo.CRITICAL: ERROR
	//foo.ERROR: ERROR
	//foo.WARNING: ERROR
	//foo.NOTICE: ERROR
	//foo.INFO: ERROR
	//foo.DEBUG: ERROR
	//foo.EMERGENCY: CRITICAL
	//foo.ALERT: CRITICAL
	//foo.CRITICAL: CRITICAL
	//foo.ERROR: CRITICAL
	//foo.WARNING: CRITICAL
	//foo.NOTICE: CRITICAL
	//foo.INFO: CRITICAL
	//foo.DEBUG: CRITICAL
	//foo.EMERGENCY: ALERT
	//foo.ALERT: ALERT
	//foo.CRITICAL: ALERT
	//foo.ERROR: ALERT
	//foo.WARNING: ALERT
	//foo.NOTICE: ALERT
	//foo.INFO: ALERT
	//foo.DEBUG: ALERT
	//foo.EMERGENCY: EMERGENCY
	//foo.ALERT: EMERGENCY
	//foo.CRITICAL: EMERGENCY
	//foo.ERROR: EMERGENCY
	//foo.WARNING: EMERGENCY
	//foo.NOTICE: EMERGENCY
	//foo.INFO: EMERGENCY
	//foo.DEBUG: EMERGENCY
}

func ExampleAddProcessor() {
	handler := handlers.NewWriterHandler(os.Stdout, level.DEBUG)
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }} {{ json false .extra }}\n")
	trace := processors.NewTraceProcessor(level.INFO)
	logger := NewLogger("test", handler)
	logger.AddProcessor(trace.Process)
	logger.Debug("foo")
	logger.Info("foo")
	// Output:
	//test.DEBUG: foo {}
	//test.INFO: foo {"file":"logger_test.go","line":117}
}

//Example Basic illustration of using logger
func Example() {
	// A logger that will display all levels to out.txt and from level WARNING or higher to stderr
	log := NewLogger(
		"test",
		handlers.NewWriterHandler(os.Stdout, level.WARNING),
		handlers.NewFileHandler("out.txt", level.DEBUG),
	)
	log.Debug("this would only be displayed in file")
	log.Warning("this would be displayed in file and on stderr")
}

func ExampleMappedWriters() {
	handler := handlers.NewMappedWriterHandler(map[level.LogLevel]io.Writer{level.INFO: os.Stdout, level.ERROR: os.Stderr})
	// set custom line because its hard to test time in output :)
	handler.GetFormatter().SetFormatLine("{{ .name }}.{{ .level }}: {{ .message }}\n")
	message := messages.NewMessage("MAPPED MESSAGE")
	logger := NewLogger("test", handler)
	logAll(logger, message)
	//foo.DEBUG: MAPPED MESSAGE
	//foo.INFO: MAPPED MESSAGE
	//foo.EMERGENCY: MAPPED MESSAGE
	//foo.ALERT: MAPPED MESSAGE
	//foo.CRITICAL: MAPPED MESSAGE
	//foo.ERROR: MAPPED MESSAGE
}

func BenchmarkTrace(b *testing.B) {
	b.StartTimer()
	logger := NewLogger("test", handlers.NewFileHandler("/dev/null", level.DEBUG))
	for i := 0; i < b.N; i++ {
		logger.Debug("hello")
	}
	b.StopTimer()
}

func BenchmarkNoTrace(b *testing.B) {
	b.StartTimer()
	logger := NewLogger("test", handlers.NewFileHandler("/dev/null", level.DEBUG))
	logger.trace = false
	for i := 0; i < b.N; i++ {
		logger.Debug("hello")
	}
	b.StopTimer()
}

func logAll(l LoggerInterface, m *messages.Record) {
	l.Emergency(m)
	l.Alert(m)
	l.Critical(m)
	l.Error(m)
	l.Warning(m)
	l.Notice(m)
	l.Info(m)
	l.Debug(m)
}
