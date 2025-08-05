package validation

import (
	"testing"
)

func TestPositiveInt(t *testing.T) {
	v := New()

	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid positive", "10", 10, false},
		{"valid large number", "1000", 1000, false},
		{"zero", "0", 0, true},
		{"negative", "-5", 0, true},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
		{"decimal", "10.5", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := v.PositiveInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PositiveInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PositiveInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTakeProfitRatio(t *testing.T) {
	v := New()

	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid ratio 1", "1", 1, false},
		{"valid ratio 2", "2", 2, false},
		{"valid ratio 10", "10", 10, false},
		{"zero", "0", 0, true},
		{"negative", "-1", 0, true},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := v.TakeProfitRatio(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TakeProfitRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TakeProfitRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositiveFloat(t *testing.T) {
	v := New()

	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"valid integer", "10", 10.0, false},
		{"valid decimal", "1.6", 1.6, false},
		{"valid small decimal", "0.01", 0.01, false},
		{"zero", "0", 0, true},
		{"negative", "-1.5", 0, true},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := v.PositiveFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PositiveFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PositiveFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrencyPair(t *testing.T) {
	v := New()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid USDJPY", "USDJPY", false},
		{"valid EURUSD", "EURUSD", false},
		{"valid GBPJPY", "GBPJPY", false},
		{"too short", "USD", true},
		{"too long", "USDJPYX", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := v.CurrencyPair(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("CurrencyPair() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}