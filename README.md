# logger
First go project to see how it all works, was missing a logger so decided to make this logger (inspired by monolog).

It supports multiple handlers/processors and custom formatter (see tests)

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
	log := logger.NewLogger("foo", handlers.NewPrintHandler(logger.INFO))
	log.Debug("foo") // Will not be printed
	log.Info("foo")  // Will be printed
}
```
to add a process to for example add the version to the context:
```
package main

import (
	"runtime"
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/handlers"
)

func main() {
	log := logger.NewLogger("foo", handlers.NewPrintHandler(logger.INFO))
	log.Info("foo")  // Will print something like : [2015-12-12 00:18:33] foo.INFO: foo {}
	log.AddProcessor(func(context map[string]interface{}) {
		context["version"] = runtime.Version()
	})
	log.Info("foo")  // Will print something like : [2015-12-12 00:18:33] foo.INFO: foo {"version":"go1.5.1"}

}
```
