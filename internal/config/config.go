package config

import (
	"log"
)

type Config struct {
	CliOptions    *CliOptions
	TwitchSecrets *TwitchSecrets
	FileOptions   *FileOptions
}

func Load() *Config {
	cliOpts, err := getCliOptions()
	if err != nil {
		log.Fatal("Failed to load cli options: ", err)
	}

	fileOpts, err := getFileOptions()
	if err != nil {
		if cliOpts.Command == "list" {
			log.Fatal("Unable to access the configuration", err)
		}
		log.Println("Warning: Unable to access the configuration.", err)
	}

	twitchSecrets, err := loadSecrets()
	if err != nil {
		log.Fatal("Failed to load twitch secrets", err)
	}

	return &Config{
		CliOptions:    cliOpts,
		TwitchSecrets: twitchSecrets,
		FileOptions:   fileOpts,
	}
}
