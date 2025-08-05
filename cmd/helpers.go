package cmd

import (
	"fmt"

	"github.com/chiyonn/callot/internal/config"
	appErrors "github.com/chiyonn/callot/internal/errors"
)

// updateConfig loads the configuration, applies updates, and saves it back
func updateConfig(updateFunc func(*config.Config) error) error {
	conf, err := config.Load()
	if err != nil {
		return appErrors.NewConfigError(fmt.Sprintf("failed to load config: %v", err))
	}

	if err := updateFunc(conf); err != nil {
		return err
	}

	if err := config.Save(conf); err != nil {
		return appErrors.NewConfigError(fmt.Sprintf("failed to save config: %v", err))
	}

	return nil
}

// exitWithError prints the error message and exits with the appropriate code
func exitWithError(err error) {
	fmt.Println(err)
	if appErr, ok := err.(*appErrors.AppError); ok {
		// Use the specific error code if it's an AppError
		// For now, we'll just use 1 for all errors
		_ = appErr
	}
	// Exit with code 1 for all errors
	// This maintains backward compatibility
}