package main

import (
	"flag"
	"fmt"

	"github.com/ThinCrunch/crunchy-streams/internal/commands"
	"github.com/ThinCrunch/crunchy-streams/internal/config"
	"github.com/ThinCrunch/crunchy-streams/internal/io"
	"github.com/ThinCrunch/crunchy-streams/internal/requests"
)

func printHelp() {
	fmt.Println("Commands: list, watch, token")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\nTo edit the groups of followed channels, edit the follows-GROUP.txt file in the config directory.")
}

func main() {
	cfg := config.Load()
	requestService := requests.NewService(*cfg.TwitchSecrets)
	commandService := commands.NewCommandService(*requestService, *cfg.CliOptions, cfg.ConfigDir)

	switch cfg.CliOptions.Command {
	case "list":
		commandService.ListLiveUsersInfo()
		break
	case "watch":
		commandService.OpenStream()
	case "token":
		answer := io.AskForString("Are you sure you want to generate a new token? [y/n]: ")
		if answer != "y" && answer != "Y" {
			break
		}
		requestService.GetNewToken()
		break
	default:
		printHelp()
	}
}
