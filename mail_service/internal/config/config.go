package config

import (
	"os"
	"path/filepath"

	"github.com/gookit/slog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Smtp   smtp   `yaml:"smtp"`
	Rabbit rabbit `yaml:"rabbit"`
}

type rabbit struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type smtp struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() *Config {
	var cfg *Config
	wd, err := os.Getwd()
	if err != nil {
		slog.Fatalf("Failed to get working directory: %v", err)
	}
	configPath := filepath.Join(wd, "config.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		data, err = os.ReadFile("../mail_service/config.yml")
		if err != nil {
			slog.Fatalf("Failed to read configuration file: %v", err)
		}
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data: %v", err)
	}

	return cfg
}
