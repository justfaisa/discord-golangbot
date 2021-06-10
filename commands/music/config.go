package main

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DiscordToken                  string
	CommandPrefix                 string
      YoutubeToken                  string
	MaximumAudioDurationInSeconds int
	MaximumQueueSize              int
	BotAdmins                     []string
}

func (config *Config) IsUserBotAdmin(userId string) bool {
	for _, adminId := range config.BotAdmins {
		if adminId == userId {
			return true
		}
	}
	return false
}

var cfg *Config

// Load initializes the configuration
func Load() {
	cfg = &Config{
		DiscordToken:  strings.TrimSpace(os.Getenv("BOT_TOKEN")),
		CommandPrefix: strings.TrimSpace(os.Getenv("COMMAND_PREFIX")),
	}
	if len(cfg.DiscordToken) == 0 {
		panic("environment variable 'BOT_TOKEN' must not be empty")
	}
	if len(cfg.CommandPrefix) != 1 {
		cfg.CommandPrefix = "!"
	}
        if len(cfg.YoutubeToken) == "" {
                panic("environment variable 'YOUTUBE_TOKEN' must not be empty")
        }
	maximumAudioDurationInSeconds, err := strconv.Atoi(strings.TrimSpace(os.Getenv("MAXIMUM_AUDIO_DURATION_IN_SECONDS")))
	if err != nil {
		cfg.MaximumAudioDurationInSeconds = 480
	} else {
		cfg.MaximumAudioDurationInSeconds = maximumAudioDurationInSeconds
	}
	maximumQueueSize, err := strconv.Atoi(strings.TrimSpace(os.Getenv("MAXIMUM_QUEUE_SIZE")))
	if err != nil {
		cfg.MaximumQueueSize = 10
	} else {
		cfg.MaximumQueueSize = maximumQueueSize
	}
	botAdmins := strings.TrimSpace(os.Getenv("BOT_ADMINS"))
	if len(botAdmins) > 0 {
		cfg.BotAdmins = strings.Split(botAdmins, ",")
	}
}

// Get returns the configuration
func Get() *Config {
	return cfg
}
