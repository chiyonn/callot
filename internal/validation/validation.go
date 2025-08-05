package validation

import (
	"errors"
	"strconv"

	"github.com/chiyonn/callot/internal/constants"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) PositiveInt(input string) (int, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("must be a valid integer")
	}
	if n <= 0 {
		return 0, errors.New("must be a positive number")
	}
	return n, nil
}

func (v *Validator) TakeProfitRatio(input string) (int, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("must be a valid integer")
	}
	if n < 1 {
		return 0, errors.New("must be greater than or equal to 1")
	}
	return n, nil
}

func (v *Validator) PositiveFloat(input string) (float64, error) {
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("must be a valid number")
	}
	if f <= 0 {
		return 0, errors.New("must be a positive number")
	}
	return f, nil
}

func (v *Validator) CurrencyPair(pair string) error {
	if len(pair) != constants.CurrencyPairLength {
		return errors.New("must be exactly 6 characters (e.g., USDJPY)")
	}
	return nil
}