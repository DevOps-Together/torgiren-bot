package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	token      string
	logLevel   string
	configFile string
	botConfig  *BotConfig
}

type BotConfig struct {
	Autoroles []*Autorole `yaml:"autoroles,flow"`
}

type Autorole struct {
	Channel string `yaml:"channel"`
	Message string `yaml:"message"`
	Role    string `yaml:"role"`
	Emoji   string `yaml:"emoji"`
}

func parseFlags() *Config {
	config := Config{}
	flag.StringVar(&config.token, "token", os.Getenv("TOKEN"), "Discord bot token. By default can be set using TOKEN env variable.")
	flag.StringVar(&config.logLevel, "log-level", "info", "Log level. One of: panic, fatal, error, warn, warning, info, debug, trace")
	flag.StringVar(&config.configFile, "config-file", "config.yaml", "Yaml config file path. Default set to config.yaml")
	flag.Parse()

	return &config
}

func loadConfigFile(config *Config) error {
	buf, err := ioutil.ReadFile(config.configFile)
	if err != nil {
		return err
	}

	botConfig := &BotConfig{}
	err = yaml.Unmarshal(buf, botConfig)
	if err != nil {
		return fmt.Errorf("in file %q: %v", config.configFile, err)
	}
	pp, err := PrettyPrint(botConfig)
	if err != nil {
		return fmt.Errorf("Error dumping config file contents: %s", err)
	}
	log.Tracef("Config file loaded: %s", pp)
	config.botConfig = botConfig

	return nil
}
