package handlers

import (
	"github.com/zicht/logger"
	"github.com/zicht/logger/formatters"
)

type Handler struct {
	channels   *logger.ChannelNames
	formatter  logger.FormatterInterface
	level      logger.LogLevel
	bubble     bool
	processors *logger.Processors
}

func (h *Handler) AddProcessor(processor logger.Processor) {
	(*h.processors) = append(*h.processors, processor)
}

func (h *Handler) GetProcessors() *logger.Processors {
	return h.processors
}

func (h *Handler) HasChannels() bool {
	return h.channels != nil && h.channels.Len() > 0
}

func (h *Handler) SetChannels(c *logger.ChannelNames) {
	h.channels = c
}

func (h *Handler) GetChannels() *logger.ChannelNames {
	if h.channels == nil {
		h.channels = new(logger.ChannelNames)
	}

	return h.channels
}

func (h *Handler) GetLevel() logger.LogLevel {
	return h.level
}

func (h *Handler) GetFormatter() logger.FormatterInterface {
	if h.formatter == nil {
		h.formatter = formatters.NewLineFormatter()
	}

	return h.formatter
}

func (h *Handler) SetFormatter(f logger.FormatterInterface) {
	h.formatter = f
}

func (h *Handler) SetBubble(bubble bool) {
	h.bubble = bubble
}

func (h *Handler) GetBubble() bool {
	return h.bubble
}

func (h *Handler) processRecord(record *logger.Record) {
	if h.processors.Len() > 0 {
		for _, i := range h.processors.Keys() {
			(*h.processors)[i](record)
		}
	}
}
