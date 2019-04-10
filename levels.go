package logger

type LogLevel uint8

const (
	Emergency LogLevel = 1 << iota
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
)

func (l LogLevel) Match(v LogLevel) bool {
	for i := 1; i <= int(v); i <<= 1 {
		if l.Has(LogLevel(i)) {
			return true
		}
	}
	return false
}

func (l LogLevel) Has(v LogLevel) bool {
	return v == (v & l)
}

func (l LogLevel) String() string {
	for i := 1; i <= 255; i <<= 1 {
		if l.Has(LogLevel(i)) {
			switch LogLevel(i) {
			case Emergency:
				return string("EMERGENCY")
			case Alert:
				return string("ALERT")
			case Critical:
				return string("CRITICAL")
			case Error:
				return string("ERROR")
			case Warning:
				return string("WARNING")
			case Notice:
				return string("NOTICE")
			case Info:
				return string("INFO")
			case Debug:
				return string("DEBUG")
			}
		}
	}
	return string("UNKNOWN")
}
