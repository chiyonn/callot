package config

import (
    "os"
    "path/filepath"
    "testing"
)

func TestGetConfigPath(t *testing.T) {
    t.Run("uses environment variable when set", func(t *testing.T) {
        customPath := "/custom/path/config.json"
        os.Setenv("CALLOT_CONFIG_PATH", customPath)
        defer os.Unsetenv("CALLOT_CONFIG_PATH")
        
        got := GetConfigPath()
        if got != customPath {
            t.Errorf("GetConfigPath() = %v, want %v", got, customPath)
        }
    })
    
    t.Run("uses default path when env var not set", func(t *testing.T) {
        os.Unsetenv("CALLOT_CONFIG_PATH")
        
        got := GetConfigPath()
        expected := filepath.Join(os.Getenv("HOME"), ".config", "callot", "config.json")
        
        if got != expected {
            t.Errorf("GetConfigPath() = %v, want %v", got, expected)
        }
    })
    
    t.Run("empty env var uses default path", func(t *testing.T) {
        os.Setenv("CALLOT_CONFIG_PATH", "")
        defer os.Unsetenv("CALLOT_CONFIG_PATH")
        
        got := GetConfigPath()
        expected := filepath.Join(os.Getenv("HOME"), ".config", "callot", "config.json")
        
        if got != expected {
            t.Errorf("GetConfigPath() = %v, want %v", got, expected)
        }
    })
}

func TestLoadWithCustomPath(t *testing.T) {
    tmp := t.TempDir()
    customConfig := filepath.Join(tmp, "custom_config.json")
    
    os.Setenv("CALLOT_CONFIG_PATH", customConfig)
    defer os.Unsetenv("CALLOT_CONFIG_PATH")
    
    // Test that it creates default config when file doesn't exist
    cfg, err := Load()
    if err != nil {
        t.Fatalf("Load() error = %v", err)
    }
    
    if cfg.Margin != 0 || len(cfg.Pairs) != 0 {
        t.Errorf("Expected default config, got %+v", cfg)
    }
    
    // Save a config
    testCfg := &Config{
        Margin: 500000,
        Pairs:  []string{"GBPUSD"},
    }
    
    if err := Save(testCfg); err != nil {
        t.Fatalf("Save() error = %v", err)
    }
    
    // Load it back
    loaded, err := Load()
    if err != nil {
        t.Fatalf("Load() error = %v", err)
    }
    
    if loaded.Margin != testCfg.Margin || loaded.Pairs[0] != testCfg.Pairs[0] {
        t.Errorf("Loaded config mismatch: got %+v, want %+v", loaded, testCfg)
    }
}