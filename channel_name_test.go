package logger

import (
	"testing"
)

func TestChannelName(t *testing.T) {
	channel := ChannelName("test_channel")
	if channel.GetName() != "test_channel" {
		t.Errorf("Expecting 'test_channel' got '%s'", channel.GetName())
	}
	if channel.IsExcluded() != false {
		t.Errorf("Expecting channel %s to be not a excluding channel", channel.GetName())
	}
	if channel.String() != "test_channel" {
		t.Errorf("Expecting 'test_channel' got %s", channel)
	}
}

func TestChannelName_excluded(t *testing.T) {
	channel := ChannelName("!test_channel")
	if channel.GetName() != "test_channel" {
		t.Errorf("Expecting 'test_channel' got '%s'", channel.GetName())
	}
	if channel.IsExcluded() == false {
		t.Errorf("Expecting channel %s to be a excluding channel", channel.GetName())
	}
	if channel.String() != "test_channel (excluded)" {
		t.Errorf("Expecting 'test_channel (excluded)' got %s", channel)
	}
}

func TestChannelName_empty(t *testing.T) {

	channel := ChannelName("")

	if false == channel.IsEmpty() {
		t.Errorf("Expecting channel name to be empty got got '%s'", channel)
	}

	if true == channel.IsExcluded() {
		t.Errorf("Expecting empty channel name to be excluded got %t", channel.IsExcluded())
	}

	if "" != channel.GetName() {
		t.Errorf("Expecting '' got %s", channel.GetName())
	}

	if "" != channel.String() {
		t.Errorf("Expecting '' got %s", channel.String())
	}
}
