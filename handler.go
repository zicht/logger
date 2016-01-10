package logger

type HandlerInterface interface {
	Write(name string, level string, message MessageInterface)
	SetFormatter(Formatter FormatterInterface)
	Support(level int16) bool
	CreateDataMap(message MessageInterface, name string, level string) map[string]interface{}
}

type Handler struct {
	Level     int16
	Formatter FormatterInterface
}

func (h *Handler) SetFormatter(Formatter FormatterInterface) {
	h.Formatter = Formatter
}

func (h Handler) GetFormatter() FormatterInterface {
	return h.Formatter
}

func (h Handler) Support(level int16) bool {
	return h.Level <= level
}

func (h Handler) CreateDataMap(message MessageInterface, name string, level string) map[string]interface{} {
	return map[string]interface{}{
		"message": message.GetMessage(),
		"extra":   message.GetContext(),
		"trace":   message.GetTrace(),
		"time":    message.GetTime(),
		"name":    name,
		"level":   level,
	}
}
