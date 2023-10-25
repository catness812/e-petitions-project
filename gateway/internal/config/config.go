package config

import (
	"github.com/gookit/slog"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	PetitionPort string `yaml:"petition_port"`
	UserPort     string `yaml:"user_port"`
	SecurityPort string `yaml:"security_port"`
	HttpPort     string `yaml:"http_port"`
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
		slog.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return cfg

}
