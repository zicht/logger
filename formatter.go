package logger

type FormatterInterface interface {
	Format(Record) ([]byte, error)
}
