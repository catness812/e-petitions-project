package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GrpcPort int      `yaml:"grpc_port"`
	Database Postgres `yaml:"database"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

var (
	configFile string
	Cfg        Config
)

func init() {
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		log.Fatalf("Failed to unmarshal YAML data: %v", err)
	}
}