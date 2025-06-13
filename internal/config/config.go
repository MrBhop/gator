package config

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var output Config
	if err := json.Unmarshal(content, &output); err != nil {
		return Config{}, err
	}

	return output, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, configFileName), nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return c.write()
}

func (c *Config) write() error {
	content, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(content); err != nil {
		return err
	}

	return nil
}
