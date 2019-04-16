package util

import (
	"math"
)

// IsValidPoints ...
func IsValidPoints(val float64) bool {
	return IsIntegral(math.Abs(val) / 0.5)
}

// IsIntegral ...
func IsIntegral(val float64) bool {
	return val == float64(int(val))
}
