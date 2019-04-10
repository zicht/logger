package logger

type ProcessorInterface interface {
	Process(record *Record)
}

// P is an shorthand/wrapper for adding unanimous
// function as processor to the logger
//
// example:
//
// log.PushProcessor(P(func(p *Record){
//     ...
// }))
//
type P func(record *Record)

// Process implements ProcessorInterface by calling the wrapped function
func (p P) Process(record *Record) {
	p(record)
}

type processor struct {
	processors []ProcessorInterface
}

func (p *processor) PushProcessor(c ProcessorInterface) {
	p.processors = append(p.processors, c)
}

func (p *processor) PopProcessors() (processor ProcessorInterface) {
	if size := len(p.processors); size == 0 {
		return nil
	} else {
		processor, p.processors = p.processors[size-1], p.processors[:size-1]
		return
	}
}

func (p *processor) process(r *Record) {
	for i, c := 0, len(p.processors); i < c; i++ {
		p.processors[i].Process(r)
	}
}
