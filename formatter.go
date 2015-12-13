package logger

type FormatterInterface interface {
	Format(name string, level string, message MessageInterface) string
}

type Formatter struct {
	FormatLine string
}
