package model

import "testing"

func TestNewCurrencyPair(t *testing.T) {
    cp, err := NewCurrencyPair("EURUSD")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if cp.Base != "EUR" || cp.Quote != "USD" {
        t.Fatalf("unexpected pair: %+v", cp)
    }
    if cp.String() != "EUR/USD" {
        t.Fatalf("unexpected string: %s", cp.String())
    }
}

func TestNewCurrencyPairInvalid(t *testing.T) {
    if _, err := NewCurrencyPair("BAD"); err == nil {
        t.Fatal("expected error for invalid pair")
    }
}

func TestPipValue(t *testing.T) {
    jpyPair, _ := NewCurrencyPair("USDJPY")
    if v := jpyPair.PipValue(); v != 0.01 {
        t.Fatalf("unexpected pip value: %f", v)
    }

    usdPair, _ := NewCurrencyPair("EURUSD")
    if v := usdPair.PipValue(); v != 0.0001 {
        t.Fatalf("unexpected pip value: %f", v)
    }
}
