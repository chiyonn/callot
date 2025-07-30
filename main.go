package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/chiyonn/callot/cmd"
  "github.com/chiyonn/callot/internal/calculator"
  "github.com/chiyonn/callot/internal/config"
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
  prompt := promptui.Prompt{
    Label:    "Enter loss-cut width (in pips)",
    Validate: validatePositiveInt,
  }

  input, err := prompt.Run()
  if err != nil {
    fmt.Println("Prompt cancelled:", err)
    os.Exit(1)
  }

  val, _ := strconv.Atoi(input)
  return val
}

func promptTakeProfitRatio(cfg *config.Config) int {
  defaultRatio := 2
  if cfg.TakeProfitRatio > 0 {
    defaultRatio = cfg.TakeProfitRatio
  }

  prompt := promptui.Prompt{
    Label:   fmt.Sprintf("Enter take-profit ratio (default: %d)", defaultRatio),
    Default: fmt.Sprintf("%d", defaultRatio),
    Validate: validateTakeProfitRatio,
  }

  input, err := prompt.Run()
  if err != nil {
    fmt.Println("Prompt cancelled:", err)
    os.Exit(1)
  }

  val, _ := strconv.Atoi(input)
  return val
}

func validatePositiveInt(input string) error {
  n, err := strconv.Atoi(input)
  if err != nil || n <= 0 {
    return fmt.Errorf("Please enter a positive number")
  }
  return nil
}

func validateTakeProfitRatio(input string) error {
  n, err := strconv.Atoi(input)
  if err != nil || n < 1 {
    return fmt.Errorf("Please enter a positive integer greater than or equal to 1")
  }
  return nil
}

func runCommand(args []string) {
  os.Args = append([]string{os.Args[0]}, args...)
  cmd.Execute()
}
