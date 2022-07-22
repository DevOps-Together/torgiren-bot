package main

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func FindChannel(channels []*discordgo.Channel, channelName string) *discordgo.Channel {
	for i := range channels {
		if channels[i].Name == channelName {
			return channels[i]
		}
	}
	return nil
}

func CreateChannel(session *discordgo.Session, guildId, channelName string) (*discordgo.Channel, error) {
	return session.GuildChannelCreate(guildId, channelName, discordgo.ChannelTypeGuildText)
}

func FindChannelMessage(session *discordgo.Session, channelID string, autorole *Autorole) (*discordgo.Message, error) {
	limit := 100
	messages, err := session.ChannelMessages(channelID, limit, "", "", "")
	for len(messages) > 0 {
		if err != nil {
			log.Errorf("Error getting messages from channel %s: %s", autorole.Channel, err)
			return nil, err
		}
		message := findMessageInternal(messages, autorole.Message)
		if message != nil {
			return message, nil
		}
		messages, err = session.ChannelMessages(channelID, limit, messages[len(messages)-1].ID, "", "")
	}

	return nil, nil
}

func findMessageInternal(messages []*discordgo.Message, needle string) *discordgo.Message {
	for i := range messages {
		if messages[i].Content == needle {
			return messages[i]
		}
	}
	return nil
}

func FindRole(session *discordgo.Session, guildId string, role string) (*discordgo.Role, error) {
	roles, err := session.GuildRoles(guildId)
	if err != nil {
		return nil, err
	}
	for i := range roles {
		if roles[i].Name == role {
			return roles[i], nil
		}
	}
	return nil, nil
}

func FindAutoroles(autoroles []*Autorole, channel *discordgo.Channel, message *discordgo.Message, emoji string) *Autorole {
	for i := range autoroles {
		if channel.Name == autoroles[i].Channel && message.Content == autoroles[i].Message && emoji == autoroles[i].Emoji {
			return autoroles[i]
		}
	}
	return nil
}
