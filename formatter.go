package logger

import (
	"io"
)

type FormatterInterface interface {
	Format(Record, io.Writer) error
}
