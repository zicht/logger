package logger

type contextMessage struct {
	m string
	c map[string]interface{}
}

func (c contextMessage) getContext() map[string]interface{} {
	if c.c == nil || len(c.c) == 0 {
		return nil
	}
	return c.c
}

func (c contextMessage) GetLogMessage() (string, map[string]interface{}) {
	return c.m, c.getContext()
}

// Message is an helper function for when you want to add context to an message
func Message(message string, context map[string]interface{}) LogMessageInterface {
	return &contextMessage{message, context}
}

// LogMessageInterface is an interface that an struct can implement
// so that you pass the object directly to the logger
//
// example:
//
// type RequestLogger struct {
//  http.Request
// }
//
// func (r RequestLogger) GetLogMessage() (string, map[string]interface{}) {
// 	return fmt.Sprintf("[%s] '%s'", r.Method, r.URL.Path), nil
// }
//
// func main() {
//  ...
//  log.Debug(RequestLogger(request))
// }
type LogMessageInterface interface {
	GetLogMessage() (string, map[string]interface{})
}
