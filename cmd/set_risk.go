package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  "github.com/spf13/cobra"
)

var setRiskCmd = &cobra.Command{
  Use:   "set-risk <percentage>",
  Short: "Set risk tolerance in percent (e.g., 1.6 means 1.6%)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    percent, err := strconv.ParseFloat(args[0], 64)
    if err != nil || percent <= 0 {
      fmt.Println("Please provide a valid positive percentage (e.g. 1.6)")
      os.Exit(1)
    }

    conf, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      os.Exit(1)
    }

    conf.RiskTolerance = percent / 100.0

    if err := config.Save(conf); err != nil {
      fmt.Println("Failed to save config:", err)
      os.Exit(1)
    }

    fmt.Printf("Risk tolerance set to %.2f%% (%.4f internally)\n", percent, conf.RiskTolerance)
  },
}

func init() {
  configCmd.AddCommand(setRiskCmd)
}

