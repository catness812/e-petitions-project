package config

import (
	"log"
	"os"

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

var (
	Cfg         Config
	errOccurred bool
)

func LoadConfig() {
	data, err := os.ReadFile("./mail_service/config.yml")
	if err != nil {
		errOccurred = true
	}

	if errOccurred {
		data, err = os.ReadFile("../mail_service/config.yml")
		if err == nil {
			errOccurred = false
		}
	}

	if errOccurred {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		log.Fatalf("Failed to unmarshal YAML data: %v", err)
	}
}
