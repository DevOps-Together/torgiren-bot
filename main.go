package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := parseFlags()
	err := setupLogging(config)
	if err != nil {
		log.Error("error setting up logging,", err)
		return
	}

	err = loadConfigFile(config)
	if err != nil {
		log.Error("error loading config file,", err)
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.token)
	if err != nil {
		log.Error("error creating Discord session,", err)
		return
	}

	// Register connectHandler func as a callback for guilds created event
	dg.AddHandler(getGuildCreateHandler(config))
	// Register the messageCreate func as a callback for Reactions:
	dg.AddHandler(getAddReactionHandler(config))
	dg.AddHandler(getRemoveReactionHandler(config))

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessageReactions | discordgo.IntentsGuilds

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error("error opening connection,", err)
		return
	}

	// Cleanly close down the Discord session.
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
