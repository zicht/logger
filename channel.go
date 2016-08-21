package logger

type Channel struct {
	logger      *Logger
	name        ChannelName
}

func (l *Channel) Emergency(message interface{}) {
	l.logger.log(EMERGENCY, l.name.GetName(), message)
}

func (l *Channel) Alert(message interface{}) {
	l.logger.log(ALERT, l.name.GetName(), message)
}

func (l *Channel) Critical(message interface{}) {
	l.logger.log(CRITICAL, l.name.GetName(), message)
}

func (l *Channel) Error(message interface{}) {
	l.logger.log(ERROR, l.name.GetName(), message)
}

func (l *Channel) Warning(message interface{}) {
	l.logger.log(WARNING, l.name.GetName(), message)
}

func (l *Channel) Notice(message interface{}) {
	l.logger.log(NOTICE, l.name.GetName(), message)
}

func (l *Channel) Info(message interface{}) {
	l.logger.log(INFO, l.name.GetName(), message)
}

func (l *Channel) Debug(message interface{}) {
	l.logger.log(DEBUG, l.name.GetName(), message)
}
