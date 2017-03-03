package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Slack SlackConfig
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
		config = Config{SlackConfig{"", "", ""}}
	}
	return config
}

