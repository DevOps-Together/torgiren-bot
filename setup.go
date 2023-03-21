package main

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func guildSetup(session *discordgo.Session, guild *discordgo.Guild, config *Config) {
	log.Info("Setting up guild: ", guild.ID)

	var err error

	for i := range config.botConfig.Autoroles {
		err = setupAutorole(session, guild.ID, config.botConfig.Autoroles[i])
	}

	if err == nil {
		log.Infof("Guild %s set up successfully", guild.ID)
	} else {
		log.Errorf("Guild %s setup failed", guild.ID)
	}
}

func setupAutorole(session *discordgo.Session, guildId string, autorole *Autorole) error {
	channels, err := session.GuildChannels(guildId)
	if err != nil {
		return err
	}
	// find or crete channel
	channel := FindChannel(channels, autorole.Channel)
	if channel == nil {
		channel, err = CreateChannel(session, guildId, autorole.Channel)
		if err != nil {
			log.Errorf("Error creating channel %s in guild %s: %s", autorole.Channel, guildId, err)
			return err
		} else {
			log.Infof("Channel %s created for guild %s", channel.Name, guildId)
		}

	} else {
		log.Infof("Channel %s already exists in guild %s", channel.Name, guildId)
	}

	// find or crete message in channel
	message, err := FindChannelMessage(session, channel.ID, autorole)

	if err != nil {
		return err
	}

	if message == nil {
		message, err = session.ChannelMessageSend(channel.ID, autorole.Message)
		if err != nil {
			log.Errorf("Error creating message in channel %s in guild %s: %s", autorole.Channel, guildId, err)
			return err
		} else {
			log.Infof("Message %s created successfully in channel %s in guild %s", message.ID, autorole.Channel, guildId)
		}
	} else {
		log.Infof("Message %s found in channel %s in guild %s", message.ID, channel.Name, guildId)
		if message.Content != autorole.Message {
			log.Infof("Message content different - needs update.")
			message, err = session.ChannelMessageEdit(channel.ID, message.ID, autorole.Message)
			if err != nil {
				log.Errorf("Error updating message %s in channel %s in guild %s: %s", message.ID, autorole.Channel, guildId, err)
				return err
			}
		}
	}

	if !message.Pinned {
		err := session.ChannelMessagePin(channel.ID, message.ID)
		if err != nil {
			log.Errorf("Error pinning message %s in channel %s in guild %s: %s", message.ID, autorole.Channel, guildId, err)
			return err
		} else {
			log.Infof("Message %s pinned successfully in channel %s in guild %s", message.ID, autorole.Channel, guildId)
		}
	} else {
		log.Infof("Message %s already pinned in channel %s in guild %s", message.ID, autorole.Channel, guildId)
	}

	//find role
	role, err := FindRole(session, guildId, autorole.Role)
	if err != nil {
		log.Errorf("Couldn't find role %s for guild %s: %s", autorole.Role, guildId, err)
		return err
	} else if role == nil {
		// create role
		role, err := session.GuildRoleCreate(guildId)
		if err != nil {
			log.Errorf("Couldn't create role %s for guild %s: %s", autorole.Role, guildId, err)
			return err
		}
		_, err = session.GuildRoleEdit(guildId, role.ID, autorole.Role, 0xFFFFFF, true, 0, false)
		if err != nil {
			log.Errorf("Couldn't edit role %s for guild %s: %s", autorole.Role, guildId, err)
			return err
		}
		log.Infof("Role %s for guild %s created", autorole.Role, guildId)
	} else {
		log.Infof("Role %s for guild %s already exists", autorole.Role, guildId)
	}

	foundReaction := false
	for i := range message.Reactions {
		if message.Reactions[i].Emoji.Name == autorole.Emoji {
			log.Infof("Reaction emoji %s already exists for message %s in channel %s in guild %s already exists", autorole.Emoji, message.ID, channel.Name, guildId)
			foundReaction = true
			break
		}
	}
	if !foundReaction {
		err := session.MessageReactionAdd(message.ChannelID, message.ID, autorole.Emoji)
		if err != nil {
			log.Errorf("Reaction emoji %s creation failed for message %s in channel %s in guild %s already exists, %s", autorole.Emoji, message.ID, channel.Name, guildId, err)
		} else {
			log.Infof("Reaction emoji %s created successfully for message %s in channel %s in guild %s already exists", autorole.Emoji, message.ID, channel.Name, guildId)
		}
	}
	return nil
}

func setupLogging(config *Config) error {
	level, err := log.ParseLevel(config.logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
