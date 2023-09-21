package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	PetitionPort string `yaml:"petition_port"`
	UserPort     string `yaml:"user_port"`
	SecurityPort string `yaml:"security_port"`
	HttpPort     string `yaml:"http_port"`
}

func LoadConfig() Config {
	var cfg Config
	data, err := os.ReadFile("gateway/config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return cfg
}
