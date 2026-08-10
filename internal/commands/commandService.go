package commands

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ThinCrunch/crunchy-streams/internal/config"
	"github.com/ThinCrunch/crunchy-streams/internal/io"
	"github.com/ThinCrunch/crunchy-streams/internal/requests"
)

var Browser = "firefox-developer-edition"
var Player = "mpv"

type CommandService struct {
	requestService requests.Service
	cliOptions     config.CliOptions
	fileOptions    config.FileOptions
}

func NewCommandService(requestService requests.Service, cliOptions config.CliOptions, fileOptions config.FileOptions) *CommandService {
	return &CommandService{
		requestService: requestService,
		cliOptions:     cliOptions,
		fileOptions:    fileOptions,
	}
}

func (s *CommandService) OpenStream() {
	streamUsername := getStreamName(s.cliOptions.Interactive, s.cliOptions.Args)

	streamHostURL := "https://www.twitch.tv/"
	streamURL := fmt.Sprintf("%s%s", streamHostURL, streamUsername)

	var command *exec.Cmd
	if s.cliOptions.Ytdlp {
		command = exec.Command(Player, streamURL)
	} else {
		command = exec.Command("streamlink", "--player", Player, streamURL, "best")
	}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}

	if s.cliOptions.NoChat {
		return
	}

	chatUrl := fmt.Sprintf("https://www.twitch.tv/popout/%s/chat", streamUsername)
	err = exec.Command(Browser, "--new-window", chatUrl).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func getStreamName(interactive bool, args []string) string {
	if interactive {
		return io.AskForString("Username: ")
	}
	return args[0]
}

func (s *CommandService) ListLiveUsersInfo() {
	followedUsers := s.getFollowedUsers()
	liveFollowedUsers := s.requestService.GetLiveUsers(followedUsers)
	data := liveFollowedUsers.Data

	for _, stream := range data {
		fmt.Println(formatStreamInfo(stream.UserName, stream.Title, stream.GameName, stream.ViewerCount))
	}
}

func (s *CommandService) getFollowedUsers() []string {
	dataFilePath := s.getFollowsDataFilePath()

	return io.GetFileLines(dataFilePath)
}

func (s *CommandService) getFollowsDataFilePath() string {
	dataFileName := strings.Join([]string{s.fileOptions.DataFilePrefix, s.cliOptions.Group, s.fileOptions.DataFileSuffix}, "")
	return filepath.Join(s.fileOptions.ConfigDir, dataFileName)
}

func formatStreamInfo(name string, title string, game string, viewerCount int) string {
	return fmt.Sprintf("%s (%s, %d): %s", name, game, viewerCount, title)
}
