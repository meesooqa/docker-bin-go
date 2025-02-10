package calc

import (
	"errors"
)

// Divide divides two numbers and returns the result.
// Parameters:
//
//	a - dividend (numerator)
//	b - divisor (denominator)
//
// Returns:
//
//	float64 - result of division if successful
//	error - error message if division by zero is attempted
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
