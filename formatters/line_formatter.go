package formatters

import (
	"io"
	"text/template"
)

type lineFormatter struct {
	Formatter
}

func NewLineFormatter() *lineFormatter {
	return &lineFormatter{Formatter{
		FormatLine: "[{{ .time.Format \"2006-01-02 15:04:05.000000\" }}] {{ .name }}{{if .trace }}.{{ .trace.PackageName }}{{end}}.{{ .level }}: {{ .message }} {{ json true .extra }}\n",
	}}
}

func (f *lineFormatter) Execute(name string, w io.Writer, data interface{}) {
	f.GetTemplate(name).Execute(w, data)
}

func (f *lineFormatter) GetTemplate(name string) *template.Template {
	tmpl, err := f.initTemplate(name).Parse(f.FormatLine)
	if err != nil {
		panic(err)
	}
	return tmpl
}
