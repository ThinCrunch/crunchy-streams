package config

import (
	"os"
	"path/filepath"
)

type FileOptions struct {
	ConfigDir      string
	DataFilePrefix string
	DataFileSuffix string
}

func getFileOptions() (*FileOptions, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configDir = filepath.Join(configDir, "crunchy-streams")

	return &FileOptions{
		ConfigDir:      configDir,
		DataFilePrefix: "follows-",
		DataFileSuffix: ".txt",
	}, nil
}
