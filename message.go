package logger

type cm struct {
	m string
	c map[string]interface{}
}

func ContextMessage(message string, context map[string]interface{}) *cm {
	return &cm{message, context}
}