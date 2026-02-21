package chat

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Nickname string `json:"nickname"`
}

func configDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, ".config", "dopogoto")
}

var configPath = filepath.Join(configDir(), "config.json")

// LoadConfig loads the nickname config, or returns a default.
func LoadConfig() Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{Nickname: GenerateAnonName()}
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil || cfg.Nickname == "" {
		return Config{Nickname: GenerateAnonName()}
	}
	return cfg
}

// SaveConfig persists the nickname config.
func SaveConfig(cfg Config) error {
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
