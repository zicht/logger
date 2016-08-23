## Logger

This a simple multi channel logger than register different channels and processors similar as monolog (php).  

Each channel can be bind to a channel(s), excluded or write to all handlers.

```
logger := logger.NewLogger(
    "main",
    handlers.NewWriterHandler(
        "main_handler", o
        s.Stdout, 
        logger.INFO, 
        logger.ChannelName("redis"),
    ),     
)

logger.Debug("foo")                     // will print nothing
logger.Register("redis")                // register channel
logger.MustGet("redis").Debug("foo")    // will print log message 

```
see [godoc.org](https://godoc.org/github.com/pbergman/logger) for docs