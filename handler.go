package logger

import "sort"

type (
	HandlerInterface interface {
		// Should return true when a
		// record can be processed
		// by this handler
		Support(Record) bool
		// main method for processing record
		// shou;d return false when propagate
		// has to stop after handling
		Handle(*Record) bool
		// set internal formatter for
		// processing the record, by
		// default should use line formatter
		GetFormatter() FormatterInterface
		SetFormatter(FormatterInterface)
		// Internal channel helpers
		GetChannels() *ChannelNames
		HasChannels() bool
		SetChannels(*ChannelNames)
		// Add/Get processor methods
		ProcessorInterface
	}
	Handlers []HandlerInterface
)

func (h *Handlers) Len() int {
	return len(*h)
}

func (h *Handlers) Keys() []int {
	keys := []int{}
	for i := range *h {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	return keys
}
