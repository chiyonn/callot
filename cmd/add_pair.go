package cmd

import (
  "fmt"
  "os"
  "strings"

  "github.com/chiyonn/callot/internal/config"
  appErrors "github.com/chiyonn/callot/internal/errors"
  "github.com/chiyonn/callot/internal/validation"
  "github.com/spf13/cobra"
)

var addPairCmd = &cobra.Command{
  Use:   "add-pair <symbol>",
  Short: "Add a currency pair (e.g., USDJPY)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    symbol := strings.ToUpper(args[0])
    
    validator := validation.New()
    if err := validator.CurrencyPair(symbol); err != nil {
      fmt.Println(appErrors.NewValidationError(fmt.Sprintf("Invalid currency pair: %v", err)))
      os.Exit(1)
    }

    err := updateConfig(func(conf *config.Config) error {
      for _, p := range conf.Pairs {
        if p == symbol {
          return appErrors.NewValidationError(fmt.Sprintf("Pair %s already exists", symbol))
        }
      }
      conf.Pairs = append(conf.Pairs, symbol)
      return nil
    })

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("Pair %s added successfully.\n", symbol)
  },
}

func init() {
  configCmd.AddCommand(addPairCmd)
}

