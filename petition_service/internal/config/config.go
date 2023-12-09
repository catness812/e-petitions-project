package config

import (
	"os"
	"path/filepath"

	"github.com/gookit/slog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	GrpcPort      int           `yaml:"grpc_port"`
	HttpPort      int           `yaml:"http_port"`
	Database      Postgres      `yaml:"postgres"`
	Broker        RabbitMQ      `yaml:"rabbit"`
	UserService   UserService   `yaml:"user_service"`
	ElasticSearch ElasticSearch `yaml:"elastic_search"`
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

type RabbitMQ struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type UserService struct {
	Port    string `yaml:"port"`
	Address string `yaml:"address"`
}

type ElasticSearch struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var Cfg Config

func LoadConfig() {
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

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		slog.Error("Failed to unmarshal YAML data: %v", err)
	}
}
