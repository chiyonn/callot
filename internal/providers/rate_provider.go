package providers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	appErrors "github.com/chiyonn/callot/internal/errors"
)

// InteractiveRateProvider prompts the user for exchange rates via stdin
type InteractiveRateProvider struct{}

func NewInteractiveRateProvider() *InteractiveRateProvider {
	return &InteractiveRateProvider{}
}

func (p *InteractiveRateProvider) GetRate(label string) (float64, error) {
	fmt.Printf("Enter current %s rate: ", label)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		rate, err := strconv.ParseFloat(input, 64)
		if err == nil && rate > 0 {
			return rate, nil
		}
		return 0, appErrors.NewValidationError("invalid rate input")
	}
	return 0, appErrors.NewIOError("failed to read input")
}