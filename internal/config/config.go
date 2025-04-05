package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName string = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return home + string(os.PathSeparator) + configFileName, nil
}

func write(cfg Config) error {
	stream, err := json.Marshal(&cfg)

	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	os.WriteFile(path, stream, 0644)
	return nil
}

func (c Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(c)
}

func Read() (Config, error) {
	var config Config

	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	stream, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal(stream, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
