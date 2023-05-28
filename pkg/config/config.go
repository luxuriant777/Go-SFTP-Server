package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Folders []string `json:"folders"`
}

func NewConfig(configFilePath string) (*Config, error) {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := &Config{}
	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
