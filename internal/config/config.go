package config

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	CliOptions    *CliOptions
	TwitchSecrets *TwitchSecrets
	ConfigDir     string
}

func Load() *Config {
	cliOpts, err := getCliOptions()
	if err != nil {
		log.Fatal("Failed to load cli options: ", err)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		if cliOpts.Command == "list" {
			log.Fatal("Unable to access the configuration", err)
		}
		log.Println("Warning: Unable to access the configuration.", err)
	}
	configDir = filepath.Join(configDir, "crunchy-streams")

	twitchSecrets, err := loadSecrets()
	if err != nil {
		log.Fatal("Failed to load twitch secrets", err)
	}

	return &Config{
		CliOptions:    cliOpts,
		TwitchSecrets: twitchSecrets,
		ConfigDir:     configDir,
	}
}
