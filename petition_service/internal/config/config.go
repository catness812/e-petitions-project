package config

import (
	"os"

	"github.com/gookit/slog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	GrpcPort    int         `yaml:"grpc_port"`
	Database    Postgres    `yaml:"postgres"`
	Broker      RabbitMQ    `yaml:"rabbit"`
	UserService UserService `yaml:"user_service"`
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

var Cfg Config

func LoadConfig() {
	var errOccurred bool

	data, err := os.ReadFile("./petition_service/config.yml")
	if err != nil {
		errOccurred = true
	}

	if errOccurred {
		data, err = os.ReadFile("../petition_service/config.yml")
		if err == nil {
			errOccurred = false
		}
	}

	if errOccurred {
		slog.Error("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		slog.Error("Failed to unmarshal YAML data: %v", err)
	}
}
