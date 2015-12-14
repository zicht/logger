package formatters

import (
	"github.com/pbergman/logger"
	"text/template"
	"io"
)

type lineFormatter struct {
	logger.Formatter
}

func NewLineFormatter() *lineFormatter {
	return &lineFormatter{logger.Formatter{
		FormatLine: "[{{ printf \"%d-%02d-%02d %02d:%02d:%02d.%09d\" .time.Year .time.Month .time.Day .time.Hour .time.Minute .time.Second .time.Nanosecond }}] {{ .name }}.{{ .level }}: {{ .message }} {{ .extra | json }}\n",
	}}
}

func (f *lineFormatter) Execute(name string, w io.Writer, data interface{}) {
	tmpl, err := f.InitTemplate(name).Parse(f.FormatLine);
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}

func (f *lineFormatter) GetTemplate(name string) *template.Template {
	tmpl, err := f.InitTemplate(name).Parse(f.FormatLine);
	if err != nil {
		panic(err)
	}
	return tmpl;
}