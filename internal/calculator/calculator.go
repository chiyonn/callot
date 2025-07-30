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
  cfg            *config.Config
  riskTolerance  float64
  allowedLossJPY float64
  pair           *model.CurrencyPair
}

func New(cfg *config.Config) *Calculator {
  c := &Calculator{
    cfg: cfg,
  }

  if cfg.RiskTolerance > 0 {
    c.riskTolerance = cfg.RiskTolerance
  } else {
    c.riskTolerance = 0.01 // fallback default
  }

  c.allowedLossJPY = float64(cfg.Margin) * c.riskTolerance

  pair, err := model.NewCurrencyPair(cfg.SelectedPair)
  if err != nil {
    fmt.Println("Invalid currency pair:", err)
    return c
  }

  c.pair = pair
  return c
}

func (c *Calculator) PrintAllowedLoss() {
  if c.pair == nil {
    fmt.Println("Currency pair not initialized.")
    return
  }

  pipValue := c.pair.PipValue()
  lossCutPips := c.cfg.LossCutPips
  lossCutCurrency := float64(lossCutPips) * pipValue

  if lossCutCurrency == 0 {
    fmt.Println("Invalid pips value.")
    return
  }

  var tradableVolume float64
  var convertedVolumeJPY float64
  var rate float64
  var quoteLabel string

  tradableVolume = c.allowedLossJPY / lossCutCurrency

  if c.pair.IsJPYQuoted() {
    fmt.Println("========== Risk Summary (JPY-quoted) ==========")
  } else {
    quoteCurrency := c.pair.Quote
    if quoteCurrency == "USD" {
      quoteLabel = "USDJPY"
    } else {
      quoteLabel = quoteCurrency + "JPY"
    }

    rate = promptExchangeRate(quoteLabel)
    convertedVolumeJPY = tradableVolume / rate

    fmt.Printf("========== Risk Summary (%s-quoted) ==========\n", quoteCurrency)
  }

  fmt.Printf("Currency Pair     : %s\n", c.pair)
  fmt.Printf("Allowed Loss      : %.0f JPY (%.2f%% of %d JPY)\n", c.allowedLossJPY, c.riskTolerance*100, c.cfg.Margin)
  fmt.Printf("Loss-cut Width    : %d pips (%.4f %s)\n", lossCutPips, lossCutCurrency, c.pair.Quote)
  fmt.Printf("Tradable Volume   : %.2f units\n", tradableVolume)

  if !c.pair.IsJPYQuoted() {
    fmt.Printf("→ In JPY units    : %.2f (using rate %.3f for %s)\n", convertedVolumeJPY, rate, quoteLabel)
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

