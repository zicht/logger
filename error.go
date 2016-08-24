package logger

import (
	"sort"
)

type ErrorStack []error

func (e *ErrorStack) Add(err error) {
	(*e) = append(*e, err)
}

func (e *ErrorStack) Len() int {
	return len(*e)
}

func (e *ErrorStack) keys() []int {
	keys := []int{}
	for k := range *e {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (e *ErrorStack) Error() string {
	if e.Len() > 0 {
		var str string
		var keys []int = e.keys()
		var last int = keys[len(keys)-1]
		for _, i := range keys {
			str += (*e)[i].Error()
			if i != last {
				str += "\n"
			}
		}
		return str
	} else {
		return ""
	}
}
