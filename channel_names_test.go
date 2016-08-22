package logger

import (
	"testing"
)

func TestChannels(t *testing.T) {

	channels := new(ChannelNames)

	if false == channels.Support(ChannelName("000")) {
		t.Error("Channel '000' should be supported.")
	}

	channels.AddChannel(ChannelName("001"))
	channels.AddChannel(ChannelName("002"))

	if channels.Support(ChannelName("000")) {
		t.Error("Channel '000' should not be supported.")
	}

	if channels.Len() != 2 {
		t.Errorf("Expecting to have 2 channels got %d", channels.Len())
	}

	if index := channels.FindChannel("001"); index != 0 {
		t.Errorf("Expecting to have channel '001' the index 0 but got %d", index)
	}

	if index := channels.FindChannel("002"); index != 1 {
		t.Errorf("Expecting to have channel '002' the index 1 but got %d", index)
	}

	if err := channels.AddChannel(ChannelName("!002")); err == nil {
		t.Error("Expectin error not te be nil!")
	}

	channels.SetChannels(ChannelName("003"))

	if channels.Len() != 1 {
		t.Errorf("Expecting to have 1 channel got %d", channels.Len())
	}

	if index := channels.FindChannel("001"); index != -1 {
		t.Error("Expecting not to find channel '001'")
	}

	if index := channels.FindChannel("003"); index != 0 {
		t.Errorf("Expecting to have channel '003' the index 0 but got %d", index)
	}

	channels.AddChannel(ChannelName("!005"))

	if channels.Support(ChannelName("005")) {
		t.Error("Channel '005' should not be supported.")
	}

}
