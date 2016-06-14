package debug

import (
	"runtime"
	"strings"
)

type Trace struct {
	Line        int
	FileName    string
	FuncName    string
}

func (t *Trace) FuncNameShort() string {
	parts := strings.Split(t.FuncName, "/")
	return parts[len(parts) - 1]
}

func NewTrace(depth int) *Trace {

	pc, file, line, _ := runtime.Caller(depth)

	return &Trace{
		Line:       line,
		FileName:   file,
		FuncName:   runtime.FuncForPC(pc).Name(),
	}

}
