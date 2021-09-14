package check

import "math"

// IsValueInMapStringString checks if a string (X) is among the map's values.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapStringString(X string, Map map[string]string) ([]string, bool) {
	var exists bool
	keys := make([]string, 0)

	if len(Map) == 0 {
		return []string{}, false
	}
	for key, value := range Map {
		if X == value {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// IsValueInMapStringInt checks if an int (X) is among the map's values.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapStringInt(X int, Map map[string]int) ([]string, bool) {
	var exists bool
	keys := make([]string, 0)

	if len(Map) == 0 {
		return []string{}, false
	}
	for key, value := range Map {
		if X == value {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// IsValueInMapStringFloat64 checks if a foat64 (X) is among the map's values.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapStringFloat64(X float64, Map map[string]float64, Epsilon float64) ([]string, bool) {
	var exists bool
	keys := make([]string, 0)

	if len(Map) == 0 {
		return []string{}, false
	}
	for key, value := range Map {
		if math.Abs(X-value) <= Epsilon {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// IsValueInMapIntString checks if an int (X) is among the map's values.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapIntString(X string, Map map[int]string) ([]int, bool) {
	var exists bool
	keys := make([]int, 0)

	if len(Map) == 0 {
		return []int{}, false
	}
	for key, value := range Map {
		if X == value {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// IsValueInMapIntInt checks if an int (X) is among the map's values.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapIntInt(X int, Map map[int]int) ([]int, bool) {
	var exists bool
	keys := make([]int, 0)

	if len(Map) == 0 {
		return []int{}, false
	}
	for key, value := range Map {
		if X == value {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// IsValueInMapIntFloat64 checks if a float64 (X) is among the map's values.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// Returns a tuple of slice providing the map's keys that have this value, and a bool value (true if the returned slice is not empty).
func IsValueInMapIntFloat64(X float64, Map map[int]float64, Epsilon float64) ([]int, bool) {
	var exists bool
	keys := make([]int, 0)

	if len(Map) == 0 {
		return []int{}, false
	}
	for key, value := range Map {
		if math.Abs(X-value) <= Epsilon {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}
