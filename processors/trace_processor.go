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
	if record.Level >= t.level {
		context := record.Extra
		context["file"] = path.Base(record.Trace.FileName)
		context["func"] = record.Trace.FuncNameShort()
 		context["line"] = record.Trace.Line
	}
}
