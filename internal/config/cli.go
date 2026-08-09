package config

import (
	"flag"
	"os"
	"strings"
)

type CliOptions struct {
	Interactive bool
	NoChat      bool
	Ytdlp       bool
	Group       string
	Command     string
	Args        []string
}

func getCliOptions() (*CliOptions, error) {
	interactive := flag.Bool("i", false, "[i]nteractive mode")
	noChat := flag.Bool("c", false, "Do not open [c]hat alongside the stream")
	ytdlp := flag.Bool("y", false, "Use [y]t-dlp instead of streamlink")
	group := flag.String("g", "default", "[g]roup of followed channels (-g=GROUP)")

	flags, nonFlags := separateFlags(os.Args[1:])
	err := flag.CommandLine.Parse(flags)
	if err != nil {
		return nil, err
	}

	result := CliOptions{
		Interactive: *interactive,
		NoChat:      *noChat,
		Ytdlp:       *ytdlp,
		Group:       *group,
	}

	if len(nonFlags) >= 1 {
		result.Command = nonFlags[0]
		result.Args = nonFlags[1:]
	}

	return &result, nil
}

func separateFlags(args []string) ([]string, []string) {
	var flags []string
	var nonFlags []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
			continue
		}
		nonFlags = append(nonFlags, arg)
	}

	return flags, nonFlags
}
