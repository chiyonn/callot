package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  appErrors "github.com/chiyonn/callot/internal/errors"
  "github.com/spf13/cobra"
)

var setRiskCmd = &cobra.Command{
  Use:   "set-risk <percentage>",
  Short: "Set risk tolerance in percent (e.g., 1.6 means 1.6%)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    percent, err := strconv.ParseFloat(args[0], 64)
    if err != nil || percent <= 0 {
      fmt.Println(appErrors.NewValidationError("Please provide a valid positive percentage (e.g. 1.6)"))
      os.Exit(1)
    }

    riskTolerance := percent / 100.0
    err = updateConfig(func(conf *config.Config) error {
      conf.RiskTolerance = riskTolerance
      return nil
    })

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("Risk tolerance set to %.2f%% (%.4f internally)\n", percent, riskTolerance)
  },
}

func init() {
  configCmd.AddCommand(setRiskCmd)
}

