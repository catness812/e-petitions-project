package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PetitionPort string `yaml:"petition_port"`
	UserPort     string `yaml:"user_port"`
	SecurityPort string `yaml:"security_port"`
	HttpPort     string `yaml:"http_port"`
}

func LoadConfig() Config {
	var cfg Config
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return cfg
}
