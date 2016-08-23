package logger

type ChannelName string

func (c ChannelName) String() string {
	switch {
	case c.IsEmpty():
		return ""
	case c.IsExcluded():
		return c.GetName() + " (excluded)"
	default:
		return c.GetName()
	}
}

func (c ChannelName) GetName() string {
	if c.IsEmpty() {
		return ""
	}
	if c[0] == '!' {
		return string(c[1:])
	} else {
		return string(c)
	}
}

func (c ChannelName) IsEmpty() bool {
	return c == ""
}

func (c ChannelName) IsExcluded() bool {
	if c.IsEmpty() {
		return false
	} else {
		return c[0] == byte(33) // starts with !
	}
}
