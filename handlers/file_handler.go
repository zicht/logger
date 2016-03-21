package handlers

import (
	"github.com/pbergman/logger/formatters"
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"os"
)

type FileHandler struct {
	Handler
	File *os.File
}

// NewFileHandler will write all records to give file, if not exist will create file
func NewFileHandler(filename string, level level.LogLevel) *FileHandler {
	file, err := os.OpenFile(filename, getFlags(filename), 0660)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		Handler{
			Level:     level,
			Formatter: formatters.NewLineFormatter(),
		},
		file,
	}
}

func (f *FileHandler) Support(level level.LogLevel) bool {
	return f.Level <= level
}

func (f *FileHandler) Write(name string, level level.LogLevel, message messages.MessageInterface) {
	f.GetFormatter().Execute(name, f.File, f.CreateDataMap(message, name, level))
}

func getFlags(file string) int {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return os.O_RDWR | os.O_CREATE | os.O_EXCL
		}
	}
	return os.O_RDWR | os.O_EXCL | os.O_APPEND
}
