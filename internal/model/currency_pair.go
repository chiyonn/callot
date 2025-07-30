package model

import (
  "errors"
  "fmt"
  "strings"
)

type CurrencyPair struct {
  Base  string // e.g. USD
  Quote string // e.g. JPY
}

func NewCurrencyPair(pair string) (*CurrencyPair, error) {
  if len(pair) != 6 {
    return nil, errors.New("invalid currency pair format (must be 6 letters)")
  }

  return &CurrencyPair{
    Base:  strings.ToUpper(pair[:3]),
    Quote: strings.ToUpper(pair[3:]),
  }, nil
}

func (cp *CurrencyPair) IsJPYQuoted() bool {
  return cp.Quote == "JPY"
}

func (cp *CurrencyPair) IsUSDQuoted() bool {
  return cp.Quote == "USD"
}

func (cp *CurrencyPair) PipValue() float64 {
  if cp.IsJPYQuoted() {
    return 0.01
  }
  return 0.0001
}

func (cp *CurrencyPair) String() string {
  return fmt.Sprintf("%s/%s", cp.Base, cp.Quote)
}

