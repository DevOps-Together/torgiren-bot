package main

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func guildSetup(session *discordgo.Session, guild *discordgo.Guild, config *Config) {
	log.Info("Setting up guild: ", guild.ID)

	var err error = nil

	channels, err := session.GuildChannels(guild.ID)
	if err != nil {
		log.Errorf("Guild %s set up failed. Can't get guild channels: ", err)
		return
	}
	for i := range config.botConfig.Autoroles {
		err = setupAutorole(session, channels, guild.ID, config.botConfig.Autoroles[i])
	}

	if err == nil {
		log.Infof("Guild %s set up successfully", guild.ID)
	} else {
		log.Errorf("Guild %s setup failed", guild.ID)
	}
}

func setupAutorole(session *discordgo.Session, channels []*discordgo.Channel, guildId string, autorole *Autorole) error {
	// find or crete channel
	channel := FindChannel(channels, autorole.Channel)
	if channel == nil {
		channel, err := CreateChannel(session, guildId, autorole.Channel)
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
	messages, err := session.ChannelMessages(channel.ID, 100, "", "", "")
	if err != nil {
		log.Errorf("Error getting messages from channel %s in guild %s: %s", autorole.Channel, guildId, err)
		return err
	}
	message := FindMessage(messages, autorole.Message)
	if message == nil {
		message, err = CreateMessage(session, channel, autorole.Message)
		if err != nil {
			log.Errorf("Error creating message in channel %s in guild %s: %s", autorole.Channel, guildId, err)
			return err
		} else {
			log.Infof("Message %s created successfully in channel %s in guild %s", message.ID, autorole.Channel, guildId)
		}
	} else {
		log.Infof("Message %s found in channel %s in guild %s", message.ID, channel.Name, guildId)
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
		role, err = session.GuildRoleEdit(guildId, role.ID, autorole.Role, 0xFFFFFF, true, 0, false)
		if err != nil {
			log.Errorf("Couldn't edit role %s for guild %s: %s", autorole.Role, guildId, err)
			return err
		}
		log.Infof("Role %s for guild %s created", autorole.Role, guildId)
	} else {
		log.Infof("Role %s for guild %s already exists", autorole.Role, guildId)
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
