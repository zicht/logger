package logger

type HandlerInterface interface {
	// should return false/true whether this
	// handler supports this record. This is
	// to do an pre check before processing
	IsHandling(*Record) bool
	// returns true when handled the record and
	// propagation should be stopped and false
	// when either not processed or propagation
	// should continue
	Handle(*Record) bool
}

type handlers struct {
	processor
	handlers []HandlerInterface
}

func (l handlers) isHandling(r *Record) int {
	for i, c := 0, len(l.handlers); i < c; i++ {
		if l.handlers[i].IsHandling(r) {
			return i
		}
	}
	return -1
}

func (l handlers) handle(r *Record) {
	if index := l.isHandling(r); index == -1 {
		return
	} else {
		l.process(r)
		for i, c := index, len(l.handlers); i < c; i++ {
			if l.handlers[i].Handle(r) {
				return
			}
		}
	}
}

func (l *handlers) AddHandlers(h HandlerInterface) {
	l.handlers = append(l.handlers, h)
}

func (l *handlers) SetHandlers(h ...HandlerInterface) {
	l.handlers = h
}

func (l *handlers) GetHandlers() []HandlerInterface {
	return l.handlers
}

type baseHandler struct {
	level  LogLevel
	bubble bool
}
