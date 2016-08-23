package handlers

import (
	"github.com/pbergman/logger"
	"github.com/pbergman/logger/formatters"
)

type Handler struct {
	name       string
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

func (h *Handler) GetName() string {
	return h.name
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
