package config

import (
	"os"

	"github.com/joho/godotenv"
)

type TwitchSecrets struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
}

func loadSecrets() (*TwitchSecrets, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &TwitchSecrets{
		ClientId:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		AccessToken:  os.Getenv("TWITCH_ACCESS_TOKEN"),
	}, nil
}
