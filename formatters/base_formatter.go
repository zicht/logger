package formatters

import (
	"encoding/json"
	"io"
	"text/template"
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

func (f *Formatter) initTemplate(name string) *template.Template {
	tmpl := template.New(name)
	tmpl.Funcs(template.FuncMap{
		"json": func(hide_empty bool, v interface{}) string {
			j, _ := json.Marshal(v)
			if len(string(j)) > 2 || hide_empty == false {
				return string(j)
			} else {
				return ""
			}
		},
	})
	return tmpl
}
