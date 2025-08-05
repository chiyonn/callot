package interfaces

// ExchangeRateProvider defines the interface for obtaining exchange rates
type ExchangeRateProvider interface {
	GetRate(label string) (float64, error)
}