package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type PRISM struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Config struct {
	Port  string `yaml:"port"`
	PRISM PRISM  `yaml:"prism"`
}

func New(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
