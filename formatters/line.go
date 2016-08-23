package formatters

import (
	"encoding/json"
	"github.com/pbergman/logger"
	"io"
	"text/template"
)

type lineFormatter struct {
	line string
	tmpl *template.Template
}

func NewLineFormatter() *lineFormatter {
	return NewCustomLineFormatter(
		"[{{ .Time.Format \"2006-01-02 15:04:05.000000\" }}] {{ .Channel.GetName }}.{{ .Level }}: {{ .Message }} {{ json true .Context }}\n",
	)
}

func NewCustomLineFormatter(line string) *lineFormatter {
	return &lineFormatter{line: line}
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

func (l *lineFormatter) Format(record logger.Record, writer io.Writer) (err error) {

	defer func() {
		if e, o := recover().(error); o {
			err = e
		}
	}()

	return l.GetTemplate().Execute(writer, record)
}
