package logger

type HandlerInterface interface {
	Write(name string, level string, message MessageInterface)
	SetFormatter(Formatter FormatterInterface)
	Support(level int16) bool
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
