package level

const (
	// Levels as described by http://tools.ietf.org/html/rfc5424
	EMERGENCY LogLevel = 600
	ALERT     LogLevel = 550
	CRITICAL  LogLevel = 500
	ERROR     LogLevel = 400
	WARNING   LogLevel = 300
	NOTICE    LogLevel = 250
	INFO      LogLevel = 200
	DEBUG     LogLevel = 100
)

type LogLevel uint16

func (l LogLevel) String() string {
	switch l {
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
