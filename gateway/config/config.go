package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PetitionPort int `yaml:"petition_port"`
	UserPort     int `yaml:"user_port"`
	SecurityPort int `yaml:"security_port"`
	HttpPort     int `yaml:"http_port"`
}

func LoadConfig() *Config {
	var cfg *Config
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return cfg
}
