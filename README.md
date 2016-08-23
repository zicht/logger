## Logger

This a simple multi channel logger than register different channels and processors similar as monolog (php).  

Echo channel can be bind to a channel(s), excluded or write to all channels.

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

	
	logger := logger.NewLogger("main", 
	    handlers.NewWriterHandler("main_handler", os.Stdout, logger.INFO)                             // will write all channels 
	    handlers.NewWriterHandler("main_handler", file01, logger.INFO, logger.ChannelName("redis"))   // will write only to redis
	    handlers.NewWriterHandler("main_handler", file02, logger.INFO, logger.ChannelName("mysql"))   // will write only to mysql
	)
	logger.Register("redis")
	logger.Register("mysql")
