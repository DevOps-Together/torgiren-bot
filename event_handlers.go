package main

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Connect handler checks if guilds have proper setup
func getGuildCreateHandler(config *Config) func(*discordgo.Session, *discordgo.GuildCreate) {
	return func(session *discordgo.Session, event *discordgo.GuildCreate) {
		guildSetup(session, event.Guild, config)
	}
}

func getAddReactionHandler(config *Config) func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
		log.Debug("MessageReactionAdd event:", event)
	}
}

func getRemoveReactionHandler(config *Config) func(*discordgo.Session, *discordgo.MessageReactionRemove) {
	return func(session *discordgo.Session, event *discordgo.MessageReactionRemove) {
		log.Debug("MessageReactionAdd event:", event)
	}
}
