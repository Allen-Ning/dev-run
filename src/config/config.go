package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Repositories []string `yaml:"repositories"`
	Token        string   `yaml:"token"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
