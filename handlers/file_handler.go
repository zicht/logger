package handlers

import (
	"os"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type fileHandler struct {
	logger.Handler
	file *os.File
}

// NewFileHandler will write all records to give file, if not exist will create file
func NewFileHandler(filename string, level int16) *fileHandler  {
	file, err := os.OpenFile(filename, getFlags(filename) , 0660)
	if (err != nil) {
		panic(err)
	}
	return &fileHandler {
		logger.Handler{
			Level: level,
			Formatter: formatters.NewLineFormatter(),
		},
		file,
	}
}

func (f *fileHandler) Write(name string, level string, message logger.MessageInterface) {
	f.GetFormatter().Execute(name, f.file, f.CreateDataMap(message, name, level))
}

func getFlags(file string) int {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return  os.O_RDWR|os.O_CREATE|os.O_EXCL
		}
	}
	return os.O_RDWR|os.O_EXCL|os.O_APPEND
}