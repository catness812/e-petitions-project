package config

import (
	"os"

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
	data, err := os.ReadFile("./mail_service/config.yml")
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
