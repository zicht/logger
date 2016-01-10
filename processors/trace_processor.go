package processors

import (
	"github.com/pbergman/logger"
	"path"
)

type TraceProcessor struct {
	level int16
}

func NewTraceProcessor(level int16) *TraceProcessor {
	return &TraceProcessor{level}
}

func (t *TraceProcessor) Process(record *logger.Record) {
	if record.GetLevel() >= t.level {
		context := record.GetContext()
		context["file"] = path.Base(record.GetTrace().FileName)
		context["line"] = record.GetTrace().Line
	}
}
