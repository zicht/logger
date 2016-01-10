package formatters

import (
	"github.com/pbergman/logger"
	"io"
	"text/template"
)

type lineFormatter struct {
	logger.Formatter
}

func NewLineFormatter() *lineFormatter {
	return &lineFormatter{logger.Formatter{
		FormatLine: "[{{ .time.Format \"2006-01-02 15:04:05.000000\" }}] {{ .name }}.{{ .trace.PackageName }}.{{ .level }}: {{ .message }} {{ json true .extra }}\n",
	}}
}

func (f *lineFormatter) Execute(name string, w io.Writer, data interface{}) {
	tmpl, err := f.InitTemplate(name).Parse(f.FormatLine)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func (f *lineFormatter) GetTemplate(name string) *template.Template {
	tmpl, err := f.InitTemplate(name).Parse(f.FormatLine)
	if err != nil {
		panic(err)
	}
	return tmpl
}
