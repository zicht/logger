package logger

import (
	"go/parser"
	"go/token"
	"runtime"
)

type packages map[string]string

func (p *packages) Get(filename string) string {

	if *p == nil {
		*p = make(map[string]string, 0)
	}

	list := *p

	if name, ok := list[filename]; ok {
		return name
	}

	if ast, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.PackageClauseOnly); err == nil {
		list[filename] = ast.Name.Name
		*p = list
		return ast.Name.Name
	}

	return ""
}

var packages_cache *packages

func init() {
	packages_cache = new(packages)
}

type Trace struct {
	Line        int
	FileName    string
	PackageName string
}

func NewTrace() *Trace {
	if _, file, line, ok := runtime.Caller(3); ok {
		return &Trace{Line: line, FileName: file, PackageName: packages_cache.Get(file)}
	}
	return nil
}
