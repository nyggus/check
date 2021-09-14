package check

import "math"

// IsValueInStringSlice checks if a string (X) is in a slice.
func IsValueInStringSlice(X string, Slice []string) bool {
	for _, value := range Slice {
		if X == value {
			return true
		}
	}
	return false
}

// IsValueInIntSlice checks if an int (X) is in a slice.
func IsValueInIntSlice(X int, Slice []int) bool {
	for _, value := range Slice {
		if X == value {
			return true
		}
	}
	return false
}

// IsValueInFloat64Slice checks if a float64 (X) is in a slice.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func IsValueInFloat64Slice(X float64, Slice []float64, Epsilon float64) bool {
	for _, value := range Slice {
		if math.Abs(X-value) <= Epsilon {
			return true
		}
	}
	return false
}
