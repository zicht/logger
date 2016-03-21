package processors

import (
	"github.com/pbergman/logger/level"
	"github.com/pbergman/logger/messages"
	"path"
)

type TraceProcessor struct {
	level level.LogLevel
}

func NewTraceProcessor(level level.LogLevel) *TraceProcessor {
	return &TraceProcessor{level}
}

func (t *TraceProcessor) Process(record *messages.Record) {
	if record.GetLevel() >= t.level {
		context := record.GetContext()
		context["file"] = path.Base(record.GetTrace().FileName)
		context["line"] = record.GetTrace().Line
	}
}
