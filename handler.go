package logger

type HandlerInterface interface {
	// Unique name of handler
	GetName() string
	// Should return true when a
	// record can be processed
	// by this handler
	Support(Record) bool
	// main method for processing record
	// shou;d return false when propagate
	// has to stop after handling
	Handle(Record) bool
	// set internal formatter for
	// processing the record, by
	// default should use line formatter
	GetFormatter() FormatterInterface
	SetFormatter(FormatterInterface)
	// Internal channel helpers
	GetChannels() *ChannelNames
	HasChannels() bool
	SetChannels(*ChannelNames)
}

