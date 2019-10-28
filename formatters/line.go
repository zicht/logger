package formatters

import (
	"bytes"
	"encoding/json"
	"sync"
	"text/template"

	"github.com/zicht/logger"
)

type lineFormatter struct {
	line string
	tmpl *template.Template
	lock sync.Mutex
	buf  *bytes.Buffer
}

func NewLineFormatter() *lineFormatter {
	return NewCustomLineFormatter(
		"[{{ .Time.Format \"2006-01-02 15:04:05.000000\" }}] {{ .Channel.GetName }}.{{ .Level }}: {{ .Message }} {{ json true .Context }}\n",
	)
}

func NewCustomLineFormatter(line string) *lineFormatter {
	return &lineFormatter{line: line, buf: new(bytes.Buffer)}
}

func (l *lineFormatter) GetTemplate() *template.Template {
	if l.tmpl == nil {
		l.tmpl = template.New("line_formatter")
		l.tmpl.Funcs(template.FuncMap{
			"json": func(hide_empty bool, v interface{}) string {
				if v == nil && hide_empty {
					return ""
				}
				j, _ := json.Marshal(v)
				if len(string(j)) > 2 || hide_empty == false {
					return string(j)
				} else {
					return ""
				}
			},
		})

		var err error

		if l.tmpl, err = l.tmpl.Parse(l.line); err != nil {
			panic(err)
		}

	}
	return l.tmpl
}

func (l *lineFormatter) Format(record logger.Record) (raw []byte, err error) {

	defer func() {
		if e, o := recover().(error); o {
			err = e
		}
	}()

	l.lock.Lock()

	defer l.lock.Unlock()
	defer l.buf.Truncate(0)

	if err := l.GetTemplate().Execute(l.buf, record); err != nil {
		return nil, err
	}

	return l.buf.Bytes(), err
}
