package calculator

import (
    "testing"

    "github.com/chiyonn/callot/internal/config"
)

func TestNewCalculator(t *testing.T) {
    conf := &config.Config{Margin: 100000, RiskTolerance: 0.02, SelectedPair: "USDJPY"}
    calc := New(conf)

    if calc.riskTolerance != 0.02 {
        t.Fatalf("riskTolerance expected 0.02 got %f", calc.riskTolerance)
    }
    expected := float64(conf.Margin) * 0.02
    if calc.maxLossJPY != expected {
        t.Fatalf("maxLossJPY expected %f got %f", expected, calc.maxLossJPY)
    }
    if calc.pair == nil || calc.pair.Base != "USD" {
        t.Fatalf("unexpected pair: %+v", calc.pair)
    }
}

func TestNewCalculatorDefaultRisk(t *testing.T) {
    conf := &config.Config{Margin: 200000, SelectedPair: "EURUSD"}
    calc := New(conf)

    if calc.riskTolerance != 0.01 {
        t.Fatalf("default riskTolerance expected 0.01 got %f", calc.riskTolerance)
    }
    expected := float64(conf.Margin) * 0.01
    if calc.maxLossJPY != expected {
        t.Fatalf("maxLossJPY expected %f got %f", expected, calc.maxLossJPY)
    }
}

func TestNewCalculatorInvalidPair(t *testing.T) {
    conf := &config.Config{Margin: 100000, SelectedPair: "BAD"}
    calc := New(conf)
    if calc.pair != nil {
        t.Fatalf("expected nil pair for invalid input")
    }
}
