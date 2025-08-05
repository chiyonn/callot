package cmd

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/internal/config"
  "github.com/chiyonn/callot/internal/constants"
  appErrors "github.com/chiyonn/callot/internal/errors"
  "github.com/spf13/cobra"
)

var setMarginCmd = &cobra.Command{
  Use:   "set-margin <percent>",
  Short: "Set margin percent (e.g., 40 means 400000 JPY)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    percent, err := strconv.Atoi(args[0])
    if err != nil || percent <= 0 {
      fmt.Println(appErrors.NewValidationError("Please provide a positive integer"))
      os.Exit(1)
    }

    err = updateConfig(func(conf *config.Config) error {
      conf.Margin = percent * constants.MarginMultiplier
      return nil
    })

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("Margin set to %d (= %d JPY).\n", percent, percent*constants.MarginMultiplier)
  },
}

func init() {
  configCmd.AddCommand(setMarginCmd)
}

