package cmd

import (
  "fmt"

  "github.com/chiyonn/callot/internal/config"
  "github.com/chiyonn/callot/internal/constants"
  "github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
  Use:   "show",
  Short: "Display current configuration",
  Run: func(cmd *cobra.Command, args []string) {
    conf, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      return
    }

    fmt.Println("Current Configuration:")
    fmt.Printf("  Margin: %d (= %d JPY)\n", conf.Margin/constants.MarginMultiplier, conf.Margin)

    if len(conf.Pairs) == 0 {
      fmt.Println("  Currency Pairs: (none)")
    } else {
      fmt.Println("  Currency Pairs:")
      for _, p := range conf.Pairs {
        fmt.Printf("    - %s\n", p)
      }
    }
  },
}

func init() {
  configCmd.AddCommand(showCmd)
}

