package main

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestFindChannel(t *testing.T) {
	expected := &discordgo.Channel{
		ID:   "1",
		Name: "one",
	}
	channels := []*discordgo.Channel{
		{
			ID:   "1",
			Name: "one",
		},
		{
			ID:   "2",
			Name: "two",
		},
	}

	result := FindChannel(channels, "one")

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}

func TestFindChannelNegtive(t *testing.T) {
	var expected *discordgo.Channel = nil
	channels := []*discordgo.Channel{
		{
			ID:   "1",
			Name: "one",
		},
		{
			ID:   "2",
			Name: "two",
		},
	}

	result := FindChannel(channels, "three")

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}

func TestFindMessageInternal(t *testing.T) {
	expected := &discordgo.Message{
		ID:      "3",
		Content: "test content",
	}
	messages := []*discordgo.Message{
		{
			ID:      "1",
			Content: "one",
		},
		{
			ID:      "2",
			Content: "two",
		},
		{
			ID:      "3",
			Content: "test content",
		},
	}

	result := findMessageInternal(messages, "test content")

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}

func TestFindMessageInternalNegtive(t *testing.T) {
	var expected *discordgo.Message = nil
	messages := []*discordgo.Message{
		{
			ID:      "1",
			Content: "one",
		},
		{
			ID:      "2",
			Content: "two",
		},
	}

	result := findMessageInternal(messages, "three")

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}

func TestFindAutoroles(t *testing.T) {
	expected := &Autorole{
		Channel: "search-channel",
		Message: "test message",
		Role:    "testing-role",
		Emoji:   "Ok",
	}
	needleChannel := &discordgo.Channel{
		Name: "search-channel",
	}
	needleMessage := &discordgo.Message{
		Content: "test message",
	}
	needleEmoji := "Ok"
	autoroles := []*Autorole{
		{
			Channel: "test-channel",
			Message: "test message1",
			Role:    "testing-role2",
			Emoji:   "Ok",
		},
		{
			Channel: "search-channel",
			Message: "test message",
			Role:    "testing-role",
			Emoji:   "Ok",
		},
		{
			Channel: "wrong-one",
			Message: "test message",
			Role:    "testing-role",
			Emoji:   "Ok",
		},
	}

	result := FindAutoroles(autoroles, needleChannel, needleMessage, needleEmoji)

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}

func TestFindAutorolesNegative(t *testing.T) {
	var expected *Autorole = nil
	needleChannel := &discordgo.Channel{
		Name: "search-channel",
	}
	needleMessage := &discordgo.Message{
		Content: "test message",
	}
	needleEmoji := "Ok"
	autoroles := []*Autorole{
		{
			Channel: "test-channel1",
			Message: "test message1",
			Role:    "testing-role2",
			Emoji:   "Ok",
		},
		{
			Channel: "search-channel2",
			Message: "test message",
			Role:    "testing-role",
			Emoji:   "Ok",
		},
		{
			Channel: "wrong-one",
			Message: "test message",
			Role:    "testing-role",
			Emoji:   "Ok",
		},
	}

	result := FindAutoroles(autoroles, needleChannel, needleMessage, needleEmoji)

	assert.Equalf(t, expected, result, "Got %s expected %s", result, expected)
}
