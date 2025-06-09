package config

import (
	"os"
	"path/filepath"
)

func GetStoragePath() (string, error) {
	configDir, err := os.UserConfigDir() //Uses XDG_CONFIG_HOME on Linux, APPDATA on Windows, and HOME/.config on macOS
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "cli-todo")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if err := os.MkdirAll(appDir, 0755); err != nil {
			return "", err
		}
	}

	return filepath.Join(appDir, "tasks.json"), nil
}
