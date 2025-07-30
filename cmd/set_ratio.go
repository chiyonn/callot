package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  "github.com/spf13/cobra"
)

var setRatioCmd = &cobra.Command{
  Use:   "set-ratio <value>",
  Short: "Set default take-profit ratio (e.g. 2 or 3)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    ratio, err := strconv.Atoi(args[0])
    if err != nil || ratio < 1 {
      fmt.Println("Please enter a positive integer (1 or higher).")
      os.Exit(1)
    }

    conf, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      os.Exit(1)
    }

    conf.TakeProfitRatio = ratio

    if err := config.Save(conf); err != nil {
      fmt.Println("Failed to save config:", err)
      os.Exit(1)
    }

    fmt.Printf("Default take-profit ratio set to %d\n", ratio)
  },
}

func init() {
  configCmd.AddCommand(setRatioCmd)
}

