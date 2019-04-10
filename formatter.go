package logger

import (
	"bytes"
	"encoding/json"
	"sync"
	"text/template"
)

// FormatterInterface are the methods that an struct should
// implement for it to be able to act as formatter that can
// can convert a Record to desired output format
type FormatterInterface interface {
	Format(*Record) ([]byte, error)
}

// F is an shorthand/wrapper for adding unanimous
// function as formatter to the processor
//
// example:
//
// processor.setFormatter(F(func(p *Record) ([]byte, error) {
//     ...
// }))
//
type F func(*Record) ([]byte, error)

// Format implements FormatterInterface by calling the wrapped function
func (f F) Format(r *Record) ([]byte, error) {
	return f(r)
}

type FormatHandlerInterface interface {
	SetFormatter(formatter FormatterInterface)
	GetFormatter() FormatterInterface
}

// FormatHandler is abstract type that can be embed to implement the
// FormatHandlerInterface (see processors)
type FormatHandler struct {
	formatter FormatterInterface
}

func (f *FormatHandler) SetFormatter(formatter FormatterInterface) {
	f.formatter = formatter
}

func (f *FormatHandler) GetFormatter() FormatterInterface {
	return f.formatter
}

func NewFormatHandler(formatter FormatterInterface) (FormatHandler, error) {
	var err error
	if formatter == nil {
		formatter, err = NewLineFormatter("[{{ .Time.Format \"2006-01-02 15:04:05.000000\" }}] {{ .Name }}.{{ .Level }}: {{ .Message }}{{ if .Context}} {{ json .Context }}{{end}}\n")
	}
	return FormatHandler{formatter: formatter}, err
}

type buf struct {
	c func(interface{})
	*bytes.Buffer
}

func (c *buf) Close() error {
	c.Truncate(0)
	c.c(c)
	return nil
}

type textTemplateFormatter struct {
	tmpl *template.Template
	pool *sync.Pool
}

func (t *textTemplateFormatter) Format(record *Record) ([]byte, error) {
	buf := t.pool.Get().(*buf)
	defer buf.Close()
	if err := t.tmpl.Execute(buf, record); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewTextTemplateFormatter(tmpl *template.Template) (FormatterInterface, error) {
	fmt := &textTemplateFormatter{
		pool: new(sync.Pool),
		tmpl: tmpl,
	}
	fmt.pool.New = func() interface{} {
		return &buf{Buffer: new(bytes.Buffer), c: fmt.pool.Put}
	}

	return fmt, nil
}

func NewLineFormatter(line string) (FormatterInterface, error) {
	tmpl := template.New("fmt")
	tmpl.Funcs(template.FuncMap{
		"json": func(v interface{}) string {
			j, _ := json.Marshal(v)
			return string(j)
		},
	})
	if _, err := tmpl.Parse(line); err != nil {
		return nil, err
	}
	return NewTextTemplateFormatter(tmpl)
}
