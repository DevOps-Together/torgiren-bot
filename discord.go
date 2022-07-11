package main

import "github.com/bwmarrin/discordgo"

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

func FindMessage(messages []*discordgo.Message, needle string) *discordgo.Message {
	for i := range messages {
		if messages[i].Content == needle {
			return messages[i]
		}
	}
	return nil
}

func CreateMessage(session *discordgo.Session, channel *discordgo.Channel, message string) (*discordgo.Message, error) {
	return session.ChannelMessageSend(channel.ID, message)
}
