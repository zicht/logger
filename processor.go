package logger

import "sort"

type ProcessorInterface interface {
	AddProcessor(Processor)
	GetProcessors() *Processors
}

type Processor func(record *Record)
type Processors []Processor

func (p *Processors) Len() int {
	return len(*p)
}

func (p *Processors) Keys() []int {
	keys := []int{}
	for i := range *p {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	return keys
}
