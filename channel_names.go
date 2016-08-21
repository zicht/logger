package logger

type ChannelNames []ChannelName

func (c *ChannelNames) Len() int {
	return len(*c)
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

func (c *ChannelNames) AddChannel(channel ChannelName) {
	*c = append(*c, channel)
}
