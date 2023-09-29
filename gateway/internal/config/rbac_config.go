package config

import (
	"github.com/gookit/slog"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type role struct {
	Code        string `yaml:"code"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type resource struct {
	Code        string `yaml:"code"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type permissions struct {
	Role     string `yaml:"role"`
	Resource string `yaml:"resource"`
	Allow    struct {
		Read   bool `yaml:"read"`
		Write  bool `yaml:"write"`
		Delete bool `yaml:"delete"`
	} `yaml:"allow"`
	Deny struct {
		Read   bool `yaml:"read"`
		Write  bool `yaml:"write"`
		Delete bool `yaml:"delete"`
	} `yaml:"deny"`
}

type PermissionsConfig struct {
	Roles       []role        `yaml:"roles"`
	Resources   []resource    `yaml:"resources"`
	Permissions []permissions `yaml:"permissions"`
}

func LoadConfigRBAC() *PermissionsConfig {
	var permConfig *PermissionsConfig
	wd, err := os.Getwd()
	if err != nil {
		slog.Fatalf("Failed to get working directory: %v", err)
	}
	configPath := filepath.Join(wd, "rbac.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Fatalf("Failed to read permissions configuration file: %v", err)
	}
	if err := yaml.Unmarshal(data, &permConfig); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return permConfig
}
