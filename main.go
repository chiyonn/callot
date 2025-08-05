package main

import (
  "fmt"
  "os"

  "github.com/chiyonn/callot/cmd"
  "github.com/chiyonn/callot/internal/calculator"
  "github.com/chiyonn/callot/internal/config"
  "github.com/chiyonn/callot/internal/constants"
  "github.com/chiyonn/callot/internal/validation"
  "github.com/manifoldco/promptui"
)

func main() {
  if len(os.Args) > 1 {
    runCommand(os.Args[1:])
    return
  }

  cfg, err := config.Load()
  if err != nil {
    fmt.Println("Failed to load config:", err)
    os.Exit(1)
  }

  runInteractive(cfg)
}

func runInteractive(cfg *config.Config) {
  if len(cfg.Pairs) == 0 {
    fmt.Println("No currency pairs found. Please add with: callot config add-pair <symbol>")
    os.Exit(1)
  }

  cfg.SelectedPair = selectCurrencyPair(cfg)
  cfg.LossCutPips = promptLossCutPips()
  cfg.TakeProfitRatio = promptTakeProfitRatio(cfg)

  calc := calculator.New(cfg)
  calc.PrintAllowedLoss()
}

func selectCurrencyPair(cfg *config.Config) string {
  prompt := promptui.Select{
    Label: "Select a currency pair",
    Items: cfg.Pairs,
  }

  _, result, err := prompt.Run()
  if err != nil {
    fmt.Println("Prompt cancelled:", err)
    os.Exit(1)
  }

  return result
}

func promptLossCutPips() int {
  validator := validation.New()
  prompt := promptui.Prompt{
    Label:    "Enter loss-cut width (in pips)",
    Validate: func(input string) error {
      _, err := validator.PositiveInt(input)
      return err
    },
  }

  input, err := prompt.Run()
  if err != nil {
    fmt.Println("Prompt cancelled:", err)
    os.Exit(1)
  }

  val, _ := validator.PositiveInt(input)
  return val
}

func promptTakeProfitRatio(cfg *config.Config) int {
  defaultRatio := constants.DefaultTakeProfitRatio
  if cfg.TakeProfitRatio > 0 {
    defaultRatio = cfg.TakeProfitRatio
  }

  validator := validation.New()
  prompt := promptui.Prompt{
    Label:   fmt.Sprintf("Enter take-profit ratio (default: %d)", defaultRatio),
    Default: fmt.Sprintf("%d", defaultRatio),
    Validate: func(input string) error {
      _, err := validator.TakeProfitRatio(input)
      return err
    },
  }

  input, err := prompt.Run()
  if err != nil {
    fmt.Println("Prompt cancelled:", err)
    os.Exit(1)
  }

  val, _ := validator.TakeProfitRatio(input)
  return val
}


func runCommand(args []string) {
  os.Args = append([]string{os.Args[0]}, args...)
  cmd.Execute()
}
