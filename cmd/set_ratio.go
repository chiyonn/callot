package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  appErrors "github.com/chiyonn/callot/internal/errors"
  "github.com/spf13/cobra"
)

var setRatioCmd = &cobra.Command{
  Use:   "set-ratio <value>",
  Short: "Set default take-profit ratio (e.g. 2 or 3)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    ratio, err := strconv.Atoi(args[0])
    if err != nil || ratio < 1 {
      fmt.Println(appErrors.NewValidationError("Please enter a positive integer (1 or higher)"))
      os.Exit(1)
    }

    err = updateConfig(func(conf *config.Config) error {
      conf.TakeProfitRatio = ratio
      return nil
    })

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("Default take-profit ratio set to %d\n", ratio)
  },
}

func init() {
  configCmd.AddCommand(setRatioCmd)
}

