package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/ThinCrunch/crunchy-streams/internal/io"
)

func (c *Config) ValidateConfig() bool {
	return c.validateGroup()
}

func (c *Config) validateGroup() bool {
	group := c.CliOptions.Group
	filePrefix := c.FileOptions.DataFilePrefix
	fileSuffix := c.FileOptions.DataFileSuffix

	groupDataFileName := strings.Join([]string{filePrefix, group, fileSuffix}, "")
	dataPath := filepath.Join(c.FileOptions.ConfigDir)

	dataFileExists, err := io.FileInDirectory(groupDataFileName, dataPath)

	dirAccessFailText := "Unable to access the config directory."
	if err != nil {
		if c.CliOptions.Command == "list" {
			log.Println("Fatal:", dirAccessFailText, "\n", err)
		} else {
			log.Println("Warning:", dirAccessFailText, "\n", err)
		}
	}

	fileExistsFailText := "The followed channels file doesn't exist for the specified group. Create it first."
	if err == nil && !dataFileExists {
		if c.CliOptions.Command == "list" {
			log.Println("Fatal:", fileExistsFailText)
		} else {
			log.Println("Warning:", fileExistsFailText)
		}
	}

	return !(c.CliOptions.Command == "list" && (!dataFileExists || err != nil))
}
