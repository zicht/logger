## Logger

This is a logger lib that can log to multiple channels and defined multiple handlers and processors for output and processing the records.

```
logger := logger.NewLogger(
    "main",
    handlers.NewWriterHandler(
        s.Stdout, 
        logger.INFO, 
        logger.ChannelName("redis"),
    ),     
)

logger.Debug("foo")                     // will print nothing
logger.Register("redis")                // register channel
logger.MustGet("redis").Debug("foo")    // will print log message 

```

see [godoc.org](https://godoc.org/github.com/zicht/logger) for docs