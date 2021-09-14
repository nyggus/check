package check

import (
	"math"
)

// AreEqualMapsStringFloat64 compares two maps map[string]float64.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AreEqualMapsStringFloat64(Map1, Map2 map[string]float64, Epsilon float64) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if math.Abs(Map1[i]-Map2[i]) > Epsilon {
			return false
		}
	}
	return true
}

// AreEqualMapsStringInt compares two maps map[string]int.
func AreEqualMapsStringInt(Map1, Map2 map[string]int) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if Map1[i] != Map2[i] {
			return false
		}
	}
	return true
}

// AreEqualMapsStringString compares two maps map[string]string.
func AreEqualMapsStringString(Map1, Map2 map[string]string) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if Map1[i] != Map2[i] {
			return false
		}
	}
	return true
}

// AreEqualMapsIntFloat64 compares two maps map[int]float64.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AreEqualMapsIntFloat64(Map1, Map2 map[int]float64, Epsilon float64) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if math.Abs(Map1[i]-Map2[i]) > Epsilon {
			return false
		}
	}
	return true
}

// AreEqualMapsIntInt compares two maps map[int]int.
func AreEqualMapsIntInt(Map1, Map2 map[int]int) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if Map1[i] != Map2[i] {
			return false
		}
	}
	return true
}

// AreEqualMapsIntString compares two maps map[int]string.
func AreEqualMapsIntString(Map1, Map2 map[int]string) bool {
	if len(Map1) != len(Map2) {
		return false
	}
	for i := range Map1 {
		if Map1[i] != Map2[i] {
			return false
		}
	}
	return true
}
