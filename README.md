# logger

This was my first project in go, to see how it all works and decided to make this because a was missing a logger that
could dispatch multiple handlers and processors (like monolog for php).

This package has bin evolved from the first releases, and you can now pause the dispatching to ask for user input or
print something else to stdout.

so for example:

```
	Logger = logger.NewLogger("some_name")
	Logger.Pause(10) // Will buffer last 10 messages
	// so something with arguments and use logger
	// do some Logger.Debug calls...
	if verbose {
		Logger.AddHandler(handlers.NewStdoutHandler(logger.DEBUG))
	} else {
		Logger.AddHandler(handlers.NewStdoutHandler(logger.INFO))
	}
	Logger.ResumeOutput() // Will print the queue based on log level
```

You can register message processors tp add something to the message context, this could be done by register a closure
or using a predefined one.

If you want to add the file name and line number of error messages you could use the TraceProcessor :

```
	processor := processors.NewTraceProcessor(logger.DEBUG)
	Logger = logger.NewLogger("some_name", 10, handlers.NewStdoutHandler(logger.DEBUG))
	Logger.AddProcessor(processor.Process)
```

Handler are used to define how the messages are handled. You can register multiple handler to write to file and stdout
for example with different levels. So for example wite to file with lever debug and one to stdout with level warning:

```
	Logger = logger.NewLogger(
		"some_name",
		10,
		handlers.NewStdoutHandler(logger.INFO),
		handlers.NewFileHandler(logger.DEBUG),
	)

```
You could also create your own one by implementing the HandlerInterface

also see docs/example for some usages