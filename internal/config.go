package config

import (
    "encoding/json"
    "errors"
    "os"
    "path/filepath"
)

type Config struct {
    Margin int `json:"margin"`
}

var defaultPath = filepath.Join(os.Getenv("HOME"), ".config", "callot", "config.json")

func Load() (*Config, error) {
    data, err := os.ReadFile(defaultPath)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return &Config{Margin: 0}, nil
        }
        return nil, err
    }
    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}

func Save(cfg *Config) error {
    dir := filepath.Dir(defaultPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }
    data, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(defaultPath, data, 0644)
}

