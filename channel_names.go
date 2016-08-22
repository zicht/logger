package logger

import (
	"errors"
	"fmt"
)

type ChannelNames []ChannelName

func (c *ChannelNames) Len() int {
	return len(*c)
}

func (c *ChannelNames) Support(name ChannelName) bool {
	// noting defined so support all channels
	if c.Len() <= 0 {
		return true
	}
	if i := c.FindChannel(name.GetName()); i >= 0 {
		return !(*c)[i].IsExcluded()
	} else {
		// not in stack so not supported
		return false
	}
}

func (c *ChannelNames) FindChannel(name string) int {
	for index, channel := range *c {
		if channel.GetName() == name {
			return index
		}
	}
	return -1
}

func (c *ChannelNames) SetChannels(channels ...ChannelName) {
	*c = (*c)[:0]

	for _, channel := range channels {
		c.AddChannel(channel)
	}
}

func (c *ChannelNames) AddChannel(channel ChannelName) error {
	for _, o := range *c {
		if o.GetName() == channel.GetName() {
			return errors.New(fmt.Sprintf("A channel with name %s is allready registered.", o))
		}
	}
	*c = append(*c, channel)
	return nil
}
