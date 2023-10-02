package config

import (
	"github.com/gookit/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GrpcPort int      `yaml:"grpc_port"`
	Database Postgres `yaml:"postgres"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

var Cfg Config

func init() {
	wd, err := os.Getwd()
	if err != nil {
		slog.Fatalf("Failed to get working directory: %v", err)
	}
	configPath := filepath.Join(wd, "config.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Fatalf("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data: %v", err)
	}
}
