package logger

const (
	ALL       int16 = 255
	EMERGENCY int16 = 128
	ALERT     int16 = 64
	CRITICAL  int16 = 32
	ERROR     int16 = 16
	WARNING   int16 = 8
	NOTICE    int16 = 4
	INFO      int16 = 2
	DEBUG     int16 = 1
)

type logger struct {
	name     string
	level    int16
	handlers []HandlerInterface
}

type FormatterInterface interface {
	Format(name string, level string, message string) string
}

type HandlerInterface interface {
	Write(name string, level string, message string)
	SetFormatter(Formatter FormatterInterface)
}

type Handler struct {
	Formatter FormatterInterface
}

func (h Handler) SetFormatter(Formatter FormatterInterface) {
	h.Formatter = Formatter
}

func (h Handler) GetFormatter() FormatterInterface {
	return h.Formatter
}

type Formatter struct {
	Line string
}

func New(name string, level int16, handlers ...HandlerInterface) *logger {
	return &logger{name: name, level: level, handlers: handlers}
}

func (l *logger) log(level int16, message string) {
	if level == (level & l.level) {
		for _, handler := range l.handlers {
			handler.Write(l.name, levelToString(level), message)
		}
	}
}

func (l *logger) Emergency(message string) {
	l.log(EMERGENCY, message)
}

func (l *logger) Alert(message string) {
	l.log(ALERT, message)
}

func (l *logger) Critical(message string) {
	l.log(CRITICAL, message)
}

func (l *logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *logger) Warning(message string) {
	l.log(WARNING, message)
}

func (l *logger) Notice(message string) {
	l.log(NOTICE, message)
}

func (l *logger) Info(message string) {
	l.log(INFO, message)
}

func (l *logger) Debug(message string) {
	l.log(DEBUG, message)
}

func levelToString(level int16) (name string) {
	switch level {
	case EMERGENCY:
		return string("EMERGENCY")
	case ALERT:
		return string("ALERT")
	case CRITICAL:
		return string("CRITICAL")
	case ERROR:
		return string("ERROR")
	case WARNING:
		return string("WARNING")
	case NOTICE:
		return string("NOTICE")
	case INFO:
		return string("INFO")
	case DEBUG:
		return string("DEBUG")
	default:
		return string("UNKNOWN")
	}
}
