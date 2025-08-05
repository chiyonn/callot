package providers

import (
	"errors"
	"testing"
)

// MockRateProvider is a test implementation of ExchangeRateProvider
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
	return 0, nil
}

func TestInteractiveRateProvider_GetRate(t *testing.T) {
	// Note: Testing interactive input is complex and typically requires
	// integration tests or manual testing. This test serves as a placeholder
	// and documents the expected behavior.
	
	t.Run("documentation", func(t *testing.T) {
		// The InteractiveRateProvider reads from stdin
		// Expected behavior:
		// 1. Prompts user with "Enter current {label} rate: "
		// 2. Reads input from stdin
		// 3. Returns parsed float64 if valid and positive
		// 4. Returns validation error if invalid or non-positive
		// 5. Returns I/O error if reading fails
		
		provider := NewInteractiveRateProvider()
		if provider == nil {
			t.Error("NewInteractiveRateProvider() should not return nil")
		}
	})
}

func TestMockRateProvider(t *testing.T) {
	t.Run("returns set rates", func(t *testing.T) {
		mock := NewMockRateProvider()
		mock.SetRate("USDJPY", 150.0)
		mock.SetRate("EURUSD", 1.1)
		
		rate, err := mock.GetRate("USDJPY")
		if err != nil {
			t.Errorf("GetRate() unexpected error: %v", err)
		}
		if rate != 150.0 {
			t.Errorf("GetRate() = %v, want %v", rate, 150.0)
		}
		
		rate, err = mock.GetRate("EURUSD")
		if err != nil {
			t.Errorf("GetRate() unexpected error: %v", err)
		}
		if rate != 1.1 {
			t.Errorf("GetRate() = %v, want %v", rate, 1.1)
		}
	})
	
	t.Run("returns zero for unknown rates", func(t *testing.T) {
		mock := NewMockRateProvider()
		
		rate, err := mock.GetRate("UNKNOWN")
		if err != nil {
			t.Errorf("GetRate() unexpected error: %v", err)
		}
		if rate != 0 {
			t.Errorf("GetRate() = %v, want %v", rate, 0.0)
		}
	})
	
	t.Run("returns set error", func(t *testing.T) {
		mock := NewMockRateProvider()
		expectedErr := errors.New("test error")
		mock.SetError(expectedErr)
		
		_, err := mock.GetRate("USDJPY")
		if err == nil {
			t.Error("GetRate() expected error, got nil")
		}
	})
}