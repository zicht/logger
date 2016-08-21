package logger

type ChannelName string

func (c ChannelName) String() string {
	if c.IsExcluded() {
		return c.GetName() + " (excluded)"
	} else {
		return c.GetName()
	}
}

func (c ChannelName) GetName() string {
	if c[0] == '!' {
		return string(c[1:])
	} else {
		return string(c)
	}
}

func (c ChannelName) IsExcluded() bool {
	return  c[0] == byte(33) // starts with !
}