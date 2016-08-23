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

func TestChannels_excluded(t *testing.T) {

	channels := new(ChannelNames)
	channels.AddChannel(ChannelName("!001"))
	channels.AddChannel(ChannelName("!002"))

	if true == channels.Support(ChannelName("001")) {
		t.Error("Channel should not support channel 001")
	}

	if true == channels.Support(ChannelName("002")) {
		t.Error("Channel should not support channel 002")
	}

	if false == channels.Support(ChannelName("test")) {
		t.Error("Channel should support channel test")
	}

	if err := channels.AddChannel(ChannelName("003")); err == nil {
		t.Error("Expecting to get error while adding channel '003'")
	} else {
		if err.Error() != "Unsupported channel '003' should be either 'Include all, except [channel]...' or 'Include only [channel]...'" {
			t.Errorf("Expecting: 'Unsupported channel '003' should be either 'Include all, except [channel]...' or 'Include only [channel]...' got: %s", err.Error())
		}
	}

	(*channels) = append(*channels, ChannelName("003"))

	if err := channels.AddChannel(ChannelName("!004")); err == nil {
		t.Error("Expecting to get error while adding channel '!004'")
	} else {
		if err.Error() != "Unsupported channel '003' should be either 'Include all, except [channel]...' or 'Include only [channel]...'" {
			t.Errorf("Expecting: 'Unsupported channel '003' should be either 'Include all, except [channel]...' or 'Include only [channel]...' got: %s", err.Error())
		}
	}
}
