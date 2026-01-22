package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AuthConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

func SaveAuthConfig(cfg AuthConfig) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot resolve home directory: %w", err)
	}

	dir := filepath.Join(home, ".xplr-mq")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create config dir: %w", err)
	}

	path := filepath.Join(dir, "config.yaml")

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("cannot marshal auth config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}

	return nil
}

func LoadAuthConfig() (*AuthConfig, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("not logged in, run `xplr login`")
	}

	var cfg AuthConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
