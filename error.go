package logger

import (
	"strconv"
)

type Errors []error

func (e *Errors) append(n error) {
	*e = append(*e, n)
}

func (e Errors) Error() string {
	var err string
	if s := len(e); s == 1 {
		err = "1 error occurred:"
	} else {
		err = strconv.Itoa(s) + " errors occurred:"
	}
	for i, c := 0, len(e); i < c; i++ {
		err += "\n\t" + e[i].Error()
	}
	err += "\n"
	return err
}

func (e *Errors) GetError() error {
	if nil == e || len(*e) == 0 {
		return nil
	}
	return e
}
