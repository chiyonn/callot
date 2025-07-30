package cmd

import (
  "fmt"

  "github.com/chiyonn/callot/internal/config"
  "github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
  Use:   "show",
  Short: "Display current configuration",
  Run: func(cmd *cobra.Command, args []string) {
    cfg, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      return
    }

    fmt.Println("Current Configuration:")
    fmt.Printf("  Margin: %d (= %d JPY)\n", cfg.Margin/10000, cfg.Margin)

    if len(cfg.Pairs) == 0 {
      fmt.Println("  Currency Pairs: (none)")
    } else {
      fmt.Println("  Currency Pairs:")
      for _, p := range cfg.Pairs {
        fmt.Printf("    - %s\n", p)
      }
    }
  },
}

func init() {
  configCmd.AddCommand(showCmd)
}

