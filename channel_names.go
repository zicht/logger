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

	if (*c)[0].IsExcluded() {
		// all except...
		return c.FindChannel(name.GetName()) == -1
	} else {
		// Include only
		return c.FindChannel(name.GetName()) >= 0
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
	var is_excluded bool
	var err func(string) error = func(name string) error {
		return errors.New(fmt.Sprintf(
			"Unsupported channel '%s' should be either 'Include all, except [channel]...' or 'Include only [channel]...'",
			name,
		))
	}

	for i := 0; i < c.Len(); i++ {
		if (*c)[i].GetName() == channel.GetName() {
			return errors.New(fmt.Sprintf("A channel with name %s is allready registered.", (*c)[i]))
		}

		if i == 0 {
			is_excluded = (*c)[i].IsExcluded()
		} else {
			if is_excluded != (*c)[i].IsExcluded() {
				return err((*c)[i].GetName())
			}
		}
	}

	if c.Len() > 0 && is_excluded != channel.IsExcluded() {
		return err(channel.GetName())
	}

	*c = append(*c, channel)
	return nil
}
