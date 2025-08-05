package calculator

import (
  "fmt"

  "github.com/chiyonn/callot/internal/config"
  "github.com/chiyonn/callot/internal/constants"
  "github.com/chiyonn/callot/internal/interfaces"
  "github.com/chiyonn/callot/internal/model"
  "github.com/chiyonn/callot/internal/providers"
)

type Calculator struct {
  config        *config.Config
  riskTolerance float64
  maxLossJPY    float64
  pair          *model.CurrencyPair
  rateProvider  interfaces.ExchangeRateProvider
}

func New(conf *config.Config) *Calculator {
  return NewWithProvider(conf, providers.NewInteractiveRateProvider())
}

func NewWithProvider(conf *config.Config, provider interfaces.ExchangeRateProvider) *Calculator {
  calc := &Calculator{
    config:       conf,
    rateProvider: provider,
  }

  if conf.RiskTolerance > 0 {
    calc.riskTolerance = conf.RiskTolerance
  } else {
    calc.riskTolerance = constants.DefaultRiskTolerance
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

    rate, err := calc.rateProvider.GetRate(rateLabel)
    if err != nil {
      fmt.Printf("Error getting exchange rate: %v\n", err)
      return
    }
    exchangeRate = rate
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


