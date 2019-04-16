Logger

A logger library that is loosely based on [monolog](https://github.com/Seldaek/monolog) and can register processors and handlers to process the records.

```
file, _ := os.OpenFile("/var/log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)

logger := logger.NewLogger(
    "main",
    logger.NewWriterHandler(
        s.Stdout,
        logger.Info,
        true,
    ),
    logger.NewThresholdHandler(
        logger.NewWriteCloserHandler(
            file,
            logger.Debug,
            false,
        ),
        10,
        logger.Error,
        false,
    ),
)

// will close file
defer logger.Close()

logger.Debug("foo") // will print foo
```

see [godoc.org](https://godoc.org/github.com/pbergman/logger) for docs and examples.