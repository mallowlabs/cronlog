package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Commands []CommandConfig
	Slack    SlackConfig
}

type CommandConfig struct {
	Path        string
	SuccessCode int
}

func (c Config) FindCommand(path string) CommandConfig {
	for _, command := range c.Commands {
		if command.Path == path {
			return command
		}
	}
	return CommandConfig{"", 0}
}

type SlackConfig struct {
	Url      string
	Channel  string
	Username string
}

func ReadConfig(path string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		config = Config{[]CommandConfig{}, SlackConfig{"", "", ""}}
	}

	if url := os.Getenv("CRONLOG_SLACK_URL"); url != "" {
		config.Slack.Url = url
	}
	if channel := os.Getenv("CRONLOG_SLACK_CHANNEL"); channel != "" {
		config.Slack.Channel = channel
	}
	if username := os.Getenv("CRONLOG_SLACK_USERNAME"); username != "" {
		config.Slack.Username = username
	}

	return config
}
