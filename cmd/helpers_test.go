package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/chiyonn/callot/internal/config"
	appErrors "github.com/chiyonn/callot/internal/errors"
)

func TestUpdateConfig(t *testing.T) {
	// Setup test environment
	tmp := t.TempDir()
	configPath := filepath.Join(tmp, "test_config.json")
	os.Setenv("CALLOT_CONFIG_PATH", configPath)
	defer os.Unsetenv("CALLOT_CONFIG_PATH")

	// Create initial config
	initialCfg := &config.Config{
		Margin: 100000,
		Pairs:  []string{"USDJPY"},
	}
	if err := config.Save(initialCfg); err != nil {
		t.Fatalf("Failed to save initial config: %v", err)
	}

	t.Run("successful update", func(t *testing.T) {
		err := updateConfig(func(conf *config.Config) error {
			conf.Margin = 200000
			return nil
		})

		if err != nil {
			t.Errorf("updateConfig() error = %v, want nil", err)
		}

		// Verify the update
		loaded, err := config.Load()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}

		if loaded.Margin != 200000 {
			t.Errorf("Margin = %d, want %d", loaded.Margin, 200000)
		}
	})

	t.Run("update function returns error", func(t *testing.T) {
		expectedErr := errors.New("update failed")
		err := updateConfig(func(conf *config.Config) error {
			return expectedErr
		})

		if err != expectedErr {
			t.Errorf("updateConfig() error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("config load failure", func(t *testing.T) {
		// Temporarily set invalid config path
		os.Setenv("CALLOT_CONFIG_PATH", "/invalid/path/config.json")
		defer os.Setenv("CALLOT_CONFIG_PATH", configPath)

		err := updateConfig(func(conf *config.Config) error {
			return nil
		})

		if err == nil {
			t.Error("updateConfig() expected error for invalid path, got nil")
		}

		appErr, ok := err.(*appErrors.AppError)
		if !ok {
			t.Errorf("Expected AppError, got %T", err)
		} else if appErr.Code != 1 {
			t.Errorf("Expected error code 1, got %d", appErr.Code)
		}
	})
}

func TestExitWithError(t *testing.T) {
	t.Run("handles AppError", func(t *testing.T) {
		err := appErrors.NewConfigError("test error")
		// Note: We can't test os.Exit behavior directly
		// This test just ensures the function doesn't panic
		exitWithError(err)
	})

	t.Run("handles generic error", func(t *testing.T) {
		err := errors.New("generic error")
		// Note: We can't test os.Exit behavior directly
		// This test just ensures the function doesn't panic
		exitWithError(err)
	})
}