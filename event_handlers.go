package main

import (
	"fmt"

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
		role, err := getRoleFromReaction(config, session, event.MessageReaction)
		if err != nil {
			log.Errorf("Error finding role: %s", err)
			return
		}

		if role == nil {
			log.Infof("Role not found")
			return
		}

		err = session.GuildMemberRoleAdd(event.GuildID, event.UserID, role.ID)
		if err != nil {
			log.Errorf("Couldn't add role %s to user %s in guild %s: %s", role.ID, event.UserID, event.GuildID, err)
		} else {
			log.Infof("Role %s added succesfully to user %s in guild %s", role.ID, event.UserID, event.GuildID)
		}
	}
}

func getRemoveReactionHandler(config *Config) func(*discordgo.Session, *discordgo.MessageReactionRemove) {
	return func(session *discordgo.Session, event *discordgo.MessageReactionRemove) {
		role, err := getRoleFromReaction(config, session, event.MessageReaction)
		if err != nil {
			log.Errorf("Error finding role: %s", err)
			return
		}

		if role == nil {
			log.Infof("Role not found")
			return
		}

		err = session.GuildMemberRoleRemove(event.GuildID, event.UserID, role.ID)
		if err != nil {
			log.Errorf("Couldn't remove role %s to user %s in guild %s: %s", role.ID, event.UserID, event.GuildID, err)
		} else {
			log.Infof("Role %s removed succesfully to user %s in guild %s", role.ID, event.UserID, event.GuildID)
		}
	}
}

func getRoleFromReaction(config *Config, session *discordgo.Session, event *discordgo.MessageReaction) (*discordgo.Role, error) {
	if event.UserID == session.State.User.ID {
		log.Warn("Reaction deleted by bot user. Shouldn't occur.")
		return nil, nil
	}

	log.Debugf("Got remove reaction event for %s channel, %s guild, %s message", event.ChannelID, event.GuildID, event.MessageID)
	channel, err := session.Channel(event.ChannelID)
	if err != nil {
		return nil, fmt.Errorf("Error finding channel %s: %s", event.ChannelID, err)
	}
	message, err := session.ChannelMessage(event.ChannelID, event.MessageID)
	if err != nil {
		return nil, fmt.Errorf("Error finding message %s: %s", event.MessageID, err)
	}
	autorole := FindAutoroles(config.botConfig.Autoroles, channel, message, event.Emoji.Name)
	if autorole == nil {
		log.Debugf("No autorole matches channel=%s, message=%s", event.ChannelID, event.MessageID)
		return nil, nil
	}
	return FindRole(session, event.GuildID, autorole.Role)
}
