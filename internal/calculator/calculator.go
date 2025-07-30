package calculator

import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "fmt"

  "github.com/chiyonn/callot/internal/config"
  "github.com/chiyonn/callot/internal/model"
)

type Calculator struct {
  config        *config.Config
  riskTolerance float64
  maxLossJPY    float64
  pair          *model.CurrencyPair
}

func New(conf *config.Config) *Calculator {
  calc := &Calculator{
    config: conf,
  }

  if conf.RiskTolerance > 0 {
    calc.riskTolerance = conf.RiskTolerance
  } else {
    calc.riskTolerance = 0.01 // fallback default
  }

  calc.maxLossJPY = float64(conf.Margin) * calc.riskTolerance

  pair, err := model.NewCurrencyPair(conf.SelectedPair)
  if err != nil {
    fmt.Println("Invalid currency pair:", err)
    return calc
  }

  calc.pair = pair
  return calc
}

func (calc *Calculator) PrintAllowedLoss() {
  if calc.pair == nil {
    fmt.Println("Currency pair not initialized.")
    return
  }

  pipValue := calc.pair.PipValue()
  lossCutPips := calc.config.LossCutPips
  lossCutAmount := float64(lossCutPips) * pipValue

  if lossCutAmount == 0 {
    fmt.Println("Invalid pips value.")
    return
  }

  var volumeJPY float64
  var exchangeRate float64
  var rateLabel string

  volume := calc.maxLossJPY / lossCutAmount

  if calc.pair.IsJPYQuoted() {
    fmt.Println("========== Risk Summary (JPY-quoted) ==========")
  } else {
    quoteCurrency := calc.pair.Quote
    if quoteCurrency == "USD" {
      rateLabel = "USDJPY"
    } else {
      rateLabel = quoteCurrency + "JPY"
    }

    exchangeRate = promptExchangeRate(rateLabel)
    volumeJPY = volume / exchangeRate

    fmt.Printf("========== Risk Summary (%s-quoted) ==========\n", quoteCurrency)
  }

  fmt.Printf("Currency Pair     : %s\n", calc.pair)
  fmt.Printf("Allowed Loss      : %.0f JPY (%.2f%% of %d JPY)\n", calc.maxLossJPY, calc.riskTolerance*100, calc.config.Margin)
  fmt.Printf("Loss-cut Width    : %d pips (%.4f %s)\n", lossCutPips, lossCutAmount, calc.pair.Quote)
  fmt.Printf("Tradable Volume   : %.2f units\n", volume)

  if !calc.pair.IsJPYQuoted() {
    fmt.Printf("→ In JPY units    : %.2f (using rate %.3f for %s)\n", volumeJPY, exchangeRate, rateLabel)
  }

  fmt.Println("===============================================")
}

func promptExchangeRate(label string) float64 {
  fmt.Printf("Enter current %s rate: ", label)
  scanner := bufio.NewScanner(os.Stdin)
  if scanner.Scan() {
    input := strings.TrimSpace(scanner.Text())
    rate, err := strconv.ParseFloat(input, 64)
    if err == nil && rate > 0 {
      return rate
    }
  }
  fmt.Println("Invalid rate input.")
  os.Exit(1)
  return 0
}

