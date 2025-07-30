package cmd

import (
  "github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
  Use:   "config",
  Short: "Manage configuration (margin, pairs, etc.)",
}

func init() {
  rootCmd.AddCommand(configCmd)
}

