# logger
First go project to see how it all works, was missing a logger so desided to make this logger (inspired by monolog).

# install:
go get -u github.com/pbergman/logger

# expample:
```
package main

import (
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/handlers"
)

func main() {
	logger := logger.New("foo", logger.ALL^logger.DEBUG, printHandler.New())
	logger.Debug("foo") // Will not be printed
	logger.Info("foo")  // Will be printed
}
```
