package logger

import (
	"text/template"
	"encoding/json"
	"io"
)

type FormatterInterface interface {
	Execute(name string, w io.Writer, data interface{})
	SetFormatLine(l string)
}

type Formatter struct {
	FormatLine string
}

func (f *Formatter) SetFormatLine(l string) {
	f.FormatLine = l
}

func (f *Formatter) InitTemplate(name string) *template.Template {
	tmpl := template.New(name);
	tmpl.Funcs(template.FuncMap{
		"json": func(v interface {}) string {
			j, _ := json.Marshal(v)
			return string(j)
		},
	})
	return tmpl;
}
