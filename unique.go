package check

import (
	"math"
)

// IsUniqueFloat64Slice checks if all elements of a float64 slice are unique (so the slice does not contain duplicated elements).
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// If the slice has no elements, the function returns true (since it does not contain duplicated elements).
func IsUniqueFloat64Slice(Slice []float64, Epsilon float64) bool {
	if len(Slice) == 0 {
		return true
	}
	for i, value := range Slice {
		for j := i + 1; j < len(Slice); j++ {
			if math.Abs(value-Slice[j]) <= Epsilon {
				return false
			}
		}
	}
	return true
}

// IsUniqueIntSlice checks if all elements of an int slice are unique (so the slice does not contain duplicated elements).
// If the slice has no elements, the function returns true (since it does not contain duplicated elements).
func IsUniqueIntSlice(Slice []int) bool {
	if len(Slice) == 0 {
		return true
	}
	for i, value := range Slice {
		for j := i; j < len(Slice); j++ {
			if i != j {
				if value == Slice[j] {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueStringSlice checks if all elements in a string slice are unique (so the slice does not contain duplicated elements).
// If the slice has no elements, the function returns true (since it does not contain duplicated elements).
func IsUniqueStringSlice(Slice []string) bool {
	if len(Slice) == 0 {
		return true
	}
	for i, value := range Slice {
		for j := i; j < len(Slice); j++ {
			if i != j {
				if value == Slice[j] {
					return false
				}
			}
		}
	}
	return true
}

// IsUnique checks whether a slice ([]int, []string and []float64) is unique (so it does not contain duplicated elements).
// Note that if you're using it for a float slice, the comparisons are made with Epsilon set to 0.
// If you need to use a different level of accuracy, use IsUniqueFloat64 instead.
func IsUnique(Slice interface{}) bool {
	if SliceInt, ok := Slice.([]int); ok {
		return IsUniqueIntSlice(SliceInt)
	}
	if SliceString, ok := Slice.([]string); ok {
		return IsUniqueStringSlice(SliceString)
	}
	if SliceFloat64, ok := Slice.([]float64); ok {
		return IsUniqueFloat64Slice(SliceFloat64, 0)
	}
	panic("implemented only for []int, []string and []float64")
}

// UniqueStringSlice returns a slice with unique elements of a string slice.
// If the slice has no elements, the function returns an empty slice.
func UniqueStringSlice(Slice []string) []string {
	if len(Slice) == 0 {
		return []string{}
	}
	unique := make([]string, 1)
	unique[0] = Slice[0]
	for _, str := range Slice {
		if !IsValueInStringSlice(str, unique) {
			unique = append(unique, str)
		}
	}
	return unique
}

// UniqueIntSlice returns a slice with unique elements of an int slice.
// If the slice has no elements, the function returns an empty slice.
func UniqueIntSlice(Slice []int) []int {
	if len(Slice) == 0 {
		return []int{}
	}
	unique := make([]int, 1)
	unique[0] = Slice[0]
	for _, value := range Slice {
		if !IsValueInIntSlice(value, unique) {
			unique = append(unique, value)
		}
	}
	return unique
}

// UniqueFloat64Slice returns a slice with unique elements of a float64 slice.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// If the slice has no elements, the function returns an empty slice.
func UniqueFloat64Slice(Slice []float64, Epsilon float64) []float64 {
	if len(Slice) == 0 {
		return []float64{}
	}
	unique := make([]float64, 1)
	unique[0] = Slice[0]
	for _, value := range Slice {
		if !IsValueInFloat64Slice(value, unique, Epsilon) {
			unique = append(unique, value)
		}
	}
	return unique
}

// IsUniqueMapIntFloat64 checks if all values of map[int]float64 are unique (so the map does not contain duplicated elements).
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapIntFloat64(Map map[int]float64, Epsilon float64) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if math.Abs(value-otherValue) <= Epsilon {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueMapStringFloat64 checks if all values of map[string]float64 are unique (so the map does not contain duplicated elements).
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapStringFloat64(Map map[string]float64, Epsilon float64) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if math.Abs(value-otherValue) <= Epsilon {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueMapIntInt checks if all values of map[int]int are unique (so the map does not contain duplicated elements).
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapIntInt(Map map[int]int) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if value == otherValue {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueMapStringInt checks if all values of map[string]int are unique (so the map does not contain duplicated elements).
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapStringInt(Map map[string]int) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if value == otherValue {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueMapIntString checks if all values of map[int]string are unique (so the map does not contain duplicated elements).
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapIntString(Map map[int]string) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if value == otherValue {
					return false
				}
			}
		}
	}
	return true
}

// IsUniqueMapStringString checks if all values of map[int]string are unique (so the map does not contain duplicated elements).
// Note that uniqueness here means the equality of values, since a map will never be unique in terms of key-value pairs.
func IsUniqueMapStringString(Map map[string]string) bool {
	if len(Map) == 0 {
		return true
	}
	for key1, value := range Map {
		for key2, otherValue := range Map {
			if key1 != key2 {
				if value == otherValue {
					return false
				}
			}
		}
	}
	return true
}
