package cmd

import (
  "fmt"
  "os"
  "strings"

  "github.com/chiyonn/callot/internal/config"
  "github.com/spf13/cobra"
)

var addPairCmd = &cobra.Command{
  Use:   "add-pair <symbol>",
  Short: "Add a currency pair (e.g., USDJPY)",
  Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    symbol := strings.ToUpper(args[0])
    if len(symbol) < 6 {
      fmt.Println("Invalid currency pair format. Example: USDJPY")
      os.Exit(1)
    }

    conf, err := config.Load()
    if err != nil {
      fmt.Println("Failed to load config:", err)
      os.Exit(1)
    }

    for _, p := range conf.Pairs {
      if p == symbol {
        fmt.Printf("Pair %s already exists.\n", symbol)
        return
      }
    }

    conf.Pairs = append(conf.Pairs, symbol)

    if err := config.Save(conf); err != nil {
      fmt.Println("Failed to save config:", err)
      os.Exit(1)
    }

    fmt.Printf("Pair %s added successfully.\n", symbol)
  },
}

func init() {
  configCmd.AddCommand(addPairCmd)
}

