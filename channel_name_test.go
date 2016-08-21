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