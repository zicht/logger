# logger

This was my first project in go, to see how it all works and decided to make this because a was missing a logger that
could dispatch multiple handlers and processors (like monolog for php).


If you want to add the file name and line number of error messages you could use the TraceProcessor :

```
	processor := processors.NewTraceProcessor(logger.DEBUG)
	Logger = logger.NewLogger("some_name", handlers.NewWriterHandler(os.Stdout, level.DEBUG))
	Logger.AddProcessor(processor.Process)
```

Handler are used to define how the messages are handled. You can register multiple handler to write to file and stdout
for example with different levels. So for example wite to file with lever debug and one to stdout with level warning:

```
	Logger = logger.NewLogger(
		"some_name",
		handlers.NewWriterHandler(os.Stdout, level.INFO),
		handlers.NewFileHandler("file.ext", logger.DEBUG),
	)

```
You could also create your own one by implementing the HandlerInterface

also see docs/example for some usages