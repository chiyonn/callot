package config

import (
  "encoding/json"
  "errors"
  "os"
  "path/filepath"
)

type Config struct {
  Margin          int      `json:"margin"`
  Pairs           []string `json:"pairs"`
  SelectedPair    string   `json:"selectedPair"`
  LossCutPips     int      `json:"lossCutPips"`
  TakeProfitRatio int      `json:"takeProfitRatio"`
  RiskTolerance   float64  `json:"riskTolerance"`

  RateAPIEndpoint string   `json:"rateAPIEndpoint"`
  PrimaryCurrency string   `json:"primaryCurrency"`
}

// GetConfigPath returns the configuration file path
// It checks CALLOT_CONFIG_PATH environment variable first
func GetConfigPath() string {
	if path := os.Getenv("CALLOT_CONFIG_PATH"); path != "" {
		return path
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "callot", "config.json")
}

func Load() (*Config, error) {
  configPath := GetConfigPath()
  data, err := os.ReadFile(configPath)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return &Config{Margin: 0, Pairs: []string{}}, nil // initialize with empty list
    }
    return nil, err
  }

  var cfg Config
  if err := json.Unmarshal(data, &cfg); err != nil {
    return nil, err
  }

  // ensure Pairs is not nil even if omitted in the file
  if cfg.Pairs == nil {
    cfg.Pairs = []string{}
  }

  return &cfg, nil
}

func Save(cfg *Config) error {
  configPath := GetConfigPath()
  dir := filepath.Dir(configPath)
  if err := os.MkdirAll(dir, 0755); err != nil {
    return err
  }

  data, err := json.MarshalIndent(cfg, "", "  ")
  if err != nil {
    return err
  }

  return os.WriteFile(configPath, data, 0644)
}
