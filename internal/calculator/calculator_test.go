package calculator

import (
    "errors"
    "testing"

    "github.com/chiyonn/callot/internal/config"
    "github.com/chiyonn/callot/internal/constants"
)

// MockRateProvider for testing
type MockRateProvider struct {
    rates map[string]float64
    err   error
}

func NewMockRateProvider() *MockRateProvider {
    return &MockRateProvider{
        rates: make(map[string]float64),
    }
}

func (m *MockRateProvider) SetRate(label string, rate float64) {
    m.rates[label] = rate
}

func (m *MockRateProvider) SetError(err error) {
    m.err = err
}

func (m *MockRateProvider) GetRate(label string) (float64, error) {
    if m.err != nil {
        return 0, m.err
    }
    if rate, ok := m.rates[label]; ok {
        return rate, nil
    }
    return 0, errors.New("rate not found")
}

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

    if calc.riskTolerance != constants.DefaultRiskTolerance {
        t.Fatalf("default riskTolerance expected %f got %f", constants.DefaultRiskTolerance, calc.riskTolerance)
    }
    expected := float64(conf.Margin) * constants.DefaultRiskTolerance
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

func TestNewWithProvider(t *testing.T) {
    conf := &config.Config{Margin: 100000, RiskTolerance: 0.02, SelectedPair: "USDJPY"}
    mockProvider := NewMockRateProvider()
    calc := NewWithProvider(conf, mockProvider)

    if calc.rateProvider != mockProvider {
        t.Fatal("expected calculator to use provided rate provider")
    }
    if calc.riskTolerance != 0.02 {
        t.Fatalf("riskTolerance expected 0.02 got %f", calc.riskTolerance)
    }
}

func TestPrintAllowedLoss_JPYQuoted(t *testing.T) {
    conf := &config.Config{
        Margin:          400000,
        RiskTolerance:   0.016,
        SelectedPair:    "USDJPY",
        LossCutPips:     20,
        TakeProfitRatio: 2,
    }
    
    calc := New(conf)
    // Test that it runs without error
    // Actual output testing would require capturing stdout
    calc.PrintAllowedLoss()
}

func TestPrintAllowedLoss_NonJPYQuoted(t *testing.T) {
    conf := &config.Config{
        Margin:          400000,
        RiskTolerance:   0.016,
        SelectedPair:    "EURUSD",
        LossCutPips:     20,
        TakeProfitRatio: 2,
    }
    
    mockProvider := NewMockRateProvider()
    mockProvider.SetRate("EURUSD", 1.1)
    mockProvider.SetRate("USDJPY", 150.0)
    
    calc := NewWithProvider(conf, mockProvider)
    // Test that it runs without error
    calc.PrintAllowedLoss()
}

func TestPrintAllowedLoss_RateProviderError(t *testing.T) {
    conf := &config.Config{
        Margin:          400000,
        RiskTolerance:   0.016,
        SelectedPair:    "EURUSD",
        LossCutPips:     20,
        TakeProfitRatio: 2,
    }
    
    mockProvider := NewMockRateProvider()
    mockProvider.SetError(errors.New("rate service unavailable"))
    
    calc := NewWithProvider(conf, mockProvider)
    // Test that it handles error gracefully
    calc.PrintAllowedLoss()
}

func TestPrintAllowedLoss_NilPair(t *testing.T) {
    conf := &config.Config{
        Margin:          400000,
        RiskTolerance:   0.016,
        SelectedPair:    "INVALID",
        LossCutPips:     20,
        TakeProfitRatio: 2,
    }
    
    calc := New(conf)
    // Test that it handles nil pair gracefully
    calc.PrintAllowedLoss()
}

func TestPrintAllowedLoss_ZeroLossCutAmount(t *testing.T) {
    conf := &config.Config{
        Margin:          400000,
        RiskTolerance:   0.016,
        SelectedPair:    "USDJPY",
        LossCutPips:     0,
        TakeProfitRatio: 2,
    }
    
    calc := New(conf)
    // Test that it handles zero loss cut amount gracefully
    calc.PrintAllowedLoss()
}