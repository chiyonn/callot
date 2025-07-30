package config

import (
    "path/filepath"
    "reflect"
    "testing"
)

func TestSaveAndLoad(t *testing.T) {
    tmp := t.TempDir()
    configFilePath = filepath.Join(tmp, "config.json")

    original := &Config{
        Margin:          100,
        Pairs:           []string{"USDJPY"},
        SelectedPair:    "USDJPY",
        LossCutPips:     10,
        TakeProfitRatio: 2,
        RiskTolerance:   0.02,
    }

    if err := Save(original); err != nil {
        t.Fatalf("save failed: %v", err)
    }

    loaded, err := Load()
    if err != nil {
        t.Fatalf("load failed: %v", err)
    }

    if !reflect.DeepEqual(original, loaded) {
        t.Fatalf("loaded config mismatch: %+v vs %+v", original, loaded)
    }
}

func TestLoadDefaultWhenMissing(t *testing.T) {
    tmp := t.TempDir()
    configFilePath = filepath.Join(tmp, "config.json")

    cfg, err := Load()
    if err != nil {
        t.Fatalf("load failed: %v", err)
    }

    if cfg.Margin != 0 || len(cfg.Pairs) != 0 {
        t.Fatalf("unexpected defaults: %+v", cfg)
    }
}
