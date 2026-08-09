package commands

import (
	"fmt"
	"log"
	"os"
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
	dataPath       string
	dataFilePrefix string
}

func NewCommandService(requestService requests.Service, cliOptions config.CliOptions, configDir string) *CommandService {
	return &CommandService{
		requestService: requestService,
		cliOptions:     cliOptions,
		dataPath:       configDir,
		dataFilePrefix: "follows-",
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
	dataFileName := strings.Join([]string{s.dataFilePrefix, s.cliOptions.Group, ".txt"}, "")
	return filepath.Join(s.dataPath, dataFileName)
}

func formatStreamInfo(name string, title string, game string, viewerCount int) string {
	return fmt.Sprintf("%s (%s, %d): %s", name, game, viewerCount, title)
}

func (s *CommandService) EditFollows() {
	dataFilePath := s.getFollowsDataFilePath()

	editor := os.Getenv("EDITOR")

	cmd := exec.Command(editor, dataFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
