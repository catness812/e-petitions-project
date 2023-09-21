package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GrpcPort int      `yaml:"grpc_port"`
	Database Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	RedisDB  int    `yaml:"db"`
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
