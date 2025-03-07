package logger

import (
	"fmt"
	"os"
)

func ExampleChannel() {
	handler := defaultHandler(DEBUG, os.Stdout)
	logger := NewLogger("main", handler)
	channel := logger.Get("foo")

	levels := [9]int{100, 200, 250, 300, 400, 500, 550, 600, 199}

	for _, l := range levels {
		handler.level = LogLevel(l)
		logAll(channel, getRecord(fmt.Sprintf("Exmaple level %s", LogLevel(l)), ""))
	}

	// Output:
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET INFO}
	// {foo Exmaple level DEBUG <nil> 2016-01-02 10:20:30 +0100 CET DEBUG}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {foo Exmaple level INFO <nil> 2016-01-02 10:20:30 +0100 CET INFO}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {foo Exmaple level NOTICE <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {foo Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level WARNING <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {foo Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level ERROR <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level CRITICAL <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level ALERT <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level ALERT <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level EMERGENCY <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET EMERGENCY}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET ALERT}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET CRITICAL}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET ERROR}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET WARNING}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET NOTICE}
	// {foo Exmaple level UNKNOWN <nil> 2016-01-02 10:20:30 +0100 CET INFO}

}
