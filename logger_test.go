package logger

import (
    "os"
    "time"
)

func ExampleNewLogger() {
    logger := NewLogger("app", NewWriterHandler(os.Stdout, Critical, false))
    defer logger.Close()
    logger.PushProcessor(P(func(r *Record) {
        // set the time static so we can test output
        r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
    }))
    logger.Debug("Should not be printed")
    logger.Info("Should not be printed")
    logger.Notice("Should not be printed")
    logger.Warning("Should not be printed")
    logger.Error("Should not be printed")
    logger.Critical("Should be printed")
    logger.Alert("Should be printed")
    logger.Emergency("Should be printed")
    // Output:
    // [2016-01-02 10:20:30.000000] app.CRITICAL: Should be printed
    // [2016-01-02 10:20:30.000000] app.ALERT: Should be printed
    // [2016-01-02 10:20:30.000000] app.EMERGENCY: Should be printed

}

func ExampleLogger_WithName() {
    logger := NewLogger("foo", NewWriterHandler(os.Stdout, Debug, false)).WithName("bar")
    defer logger.Close()
    logger.PushProcessor(P(func(r *Record) {
        // set the time static so we can test output
        r.Time = time.Date(2016, 1, 2, 10, 20, 30, 0, time.Local)
    }))
    logger.Debug("hello world")
    // Output:
    // [2016-01-02 10:20:30.000000] bar.DEBUG: hello world
}