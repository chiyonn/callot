package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  "github.com/spf13/cobra"
)

var setMarginCmd = &cobra.Command{
  Use:   "set-margin <percent>",
  Short: "Set margin percent (e.g., 40 means 400000 JPY)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    percent, err := strconv.Atoi(args[0])
    if err != nil || percent <= 0 {
      fmt.Println("Please provide a positive integer.")
      os.Exit(1)
    }

    cfg, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      os.Exit(1)
    }

    cfg.Margin = percent * 10000

    if err := config.Save(cfg); err != nil {
      fmt.Println("Failed to save config:", err)
      os.Exit(1)
    }

    fmt.Printf("Margin set to %d (= %d JPY).\n", percent, cfg.Margin)
  },
}

func init() {
  configCmd.AddCommand(setMarginCmd)
}

