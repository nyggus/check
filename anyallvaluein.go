package check

import (
	"math"
)

// AllValuesInIntSlice checks if all values of one int slice are in another slice.
// When the first slice is empty, it returns true. When the second slice is empty, it returns false.
func AllValuesInIntSlice(Slice1, Slice2 []int) bool {
	if len(Slice1) == 0 {
		return true
	}
	if len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if !IsValueInIntSlice(x, Slice2) {
			return false
		}
	}
	return true
}

// AllValuesInStringSlice checks if all values of one string slice are in another slice.
// When the first slice is empty, it returns true. When the second slice is empty, it returns false.
func AllValuesInStringSlice(Slice1, Slice2 []string) bool {
	if len(Slice1) == 0 {
		return true
	}
	if len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if !IsValueInStringSlice(x, Slice2) {
			return false
		}
	}
	return true
}

// AllValuesInFloat64Slice checks if all values of one float64 slice are in another slice.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When the first slice is empty, it returns true. When the second slice is empty, it returns false.
func AllValuesInFloat64Slice(Slice1, Slice2 []float64, Epsilon float64) bool {
	if len(Slice1) == 0 {
		return true
	}
	if len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if !IsValueInFloat64Slice(x, Slice2, Epsilon) {
			return false
		}
	}
	return true
}

// AnyValueInIntSlice checks if any of the values of one int slice is in another slice.
// When the first or the second (or both) slice is empty, it returns false.
func AnyValueInIntSlice(Slice1, Slice2 []int) bool {
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if IsValueInIntSlice(x, Slice2) {
			return true
		}
	}
	return false
}

// WhichValuesInIntSlice checks which values of one int slice are in another slice.
// When the first or the second (or both) slice is empty, it returns false and an empty map.
// The function returns a tuple with a map with values from Slice1 as keys and their indices from Slice2 as the map's values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInIntSlice(Slice1, Slice2 []int) (map[int][]int, bool) {
	values := make(map[int][]int)
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return values, false
	}

	// Remove dulpicated elements from Slice 1 (not from Slice 2!)
	if !IsUniqueIntSlice(Slice1) {
		Slice1 = UniqueIntSlice(Slice1)
	}

	for _, valueInSlice1 := range Slice1 {
		if IsValueInIntSlice(valueInSlice1, Slice2) {
			for index, valueInSlice2 := range Slice2 {
				if valueInSlice1 == valueInSlice2 && !IsValueInIntSlice(index, values[valueInSlice1]) {
					values[valueInSlice1] = append(values[valueInSlice1], index)
				}
			}
		}
	}
	return values, len(values) > 0
}

// AnyValueInStringSlice checks if any of the values of one string slice is in another slice.
// When the first or the second (or both) slice is empty, it returns false.
func AnyValueInStringSlice(Slice1, Slice2 []string) bool {
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if IsValueInStringSlice(x, Slice2) {
			return true
		}
	}
	return false
}

// WhichValuesInStringSlice checks which values of one string slice are in another slice.
// When the first or the second (or both) slice is empty, it returns false.
// The function returns a tuple with a map with values from Slice1 as keys and their indices from Slice2 as the map's values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInStringSlice(Slice1, Slice2 []string) (map[string][]int, bool) {
	values := make(map[string][]int)
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return values, false
	}

	// Remove dulpicated elements from Slice 1 (not from Slice 2!)
	if !IsUniqueStringSlice(Slice1) {
		Slice1 = UniqueStringSlice(Slice1)
	}

	for _, valueInSlice1 := range Slice1 {
		if IsValueInStringSlice(valueInSlice1, Slice2) {
			for index, valueInSlice2 := range Slice2 {
				if valueInSlice1 == valueInSlice2 && !IsValueInIntSlice(index, values[valueInSlice1]) {
					values[valueInSlice1] = append(values[valueInSlice1], index)
				}
			}
		}
	}
	return values, len(values) > 0
}

// AnyValueInFloat64Slice checks if any of the values of one float64 slice is in another slice.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When the first or the second (or both) slice is empty, it returns false.
func AnyValueInFloat64Slice(Slice1, Slice2 []float64, Epsilon float64) bool {
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return false
	}
	for _, x := range Slice1 {
		if IsValueInFloat64Slice(x, Slice2, Epsilon) {
			return true
		}
	}
	return false
}

// WhichValuesInFloat64Slice checks which values of one float64 slice are in another slice.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When the first or the second (or both) slice is empty, it returns false.
// The function returns a tuple with a map with values from Slice1 as keys and their indices from Slice2 as the map's values,//
// and a boolean value (true if the returned map is not empty).
func WhichValuesInFloat64Slice(Slice1, Slice2 []float64, Epsilon float64) (map[float64][]int, bool) {
	values := make(map[float64][]int)
	if len(Slice1) == 0 || len(Slice2) == 0 {
		return values, false
	}

	// Remove dulpicated elements from Slice 1 (not from Slice 2!)
	if !IsUniqueFloat64Slice(Slice1, Epsilon) {
		Slice1 = UniqueFloat64Slice(Slice1, Epsilon)
	}

	for _, valueInSlice1 := range Slice1 {
		if IsValueInFloat64Slice(valueInSlice1, Slice2, Epsilon) {
			for index, valueInSlice2 := range Slice2 {
				if math.Abs(valueInSlice1-valueInSlice2) <= Epsilon && !IsValueInIntSlice(index, values[valueInSlice1]) {
					values[valueInSlice1] = append(values[valueInSlice1], index)
				}
			}
		}
	}
	return values, len(values) > 0
}

// AnyValueInMapIntInt checks if any of the values of an int slice is a value of map[int]int.
// When either the slice or the map is empty, it returns false.
func AnyValueInMapIntInt(Slice []int, Map map[int]int) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntInt(x, Map); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapIntInt checks if all values of an int slice are values of map[int]int.
// When either the slice or the map is empty, it returns false.
func AllValuesInMapIntInt(Slice []int, Map map[int]int) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntInt(x, Map); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapIntInt checks which values of an int slice are values of map[int]int.
// When either the slice or the map is empty, it returns false.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInMapIntInt(Slice []int, Map map[int]int) (map[int][]int, bool) {
	values := make(map[int][]int)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[int][]int{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapIntInt(x, Map); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}

// AnyValueInMapIntString checks if any of the values of a string slice is a value of map[int]string.
// When either the slice or the map is empty, it returns false.
func AnyValueInMapIntString(Slice []string, Map map[int]string) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntString(x, Map); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapIntString checks if all values of a string slice are in values of map[int]string.
// When either the slice or the map is empty, it returns false.
func AllValuesInMapIntString(Slice []string, Map map[int]string) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntString(x, Map); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapIntString checks which values of a string slice are values of map[int]string.
// When either the slice or the map is empty, it returns false.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInMapIntString(Slice []string, Map map[int]string) (map[string][]int, bool) {
	values := make(map[string][]int)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[string][]int{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapIntString(x, Map); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}

// AnyValueInMapIntFloat64 checks if any of the values of a float64 slice is in map[int]float64.
// When either the slice or the map is empty, it returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AnyValueInMapIntFloat64(Slice []float64, Map map[int]float64, Epsilon float64) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntFloat64(x, Map, Epsilon); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapIntFloat64 checks if all values of a float64 slice are values of map[int]float64.
// When either the slice or the map is empty, it returns false.
func AllValuesInMapIntFloat64(Slice []float64, Map map[int]float64, Epsilon float64) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapIntFloat64(x, Map, Epsilon); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapIntFloat64 checks which values of a float64 slice are values of map[int]float64.
// When either the slice or the map is empty, it returns false. The Epsilon parameter sets the accuracy of the comparison of two floats.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
// BEWARE! Do note that we work with floats, so it's safest to round them before using this function,
// since they will be keys of a returned map. Thus, this
// WhichValuesInMapIntFloat64({[]float64{.01002, .01, .2}, map[int]float64{1: .01, 2: .011}, .0001)
// will return the following map: map[float64][]int{.01: {1}, .01002: {1}}}
func WhichValuesInMapIntFloat64(Slice []float64, Map map[int]float64, Epsilon float64) (map[float64][]int, bool) {
	values := make(map[float64][]int)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[float64][]int{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapIntFloat64(x, Map, Epsilon); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}

// AnyValueInMapStringInt checks if any of the values of a string slice is a value of map[string]int.
// When either the slice or the map is empty, it returns false.
func AnyValueInMapStringInt(Slice []int, Map map[string]int) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringInt(x, Map); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapStringInt checks if all values of a string slice are values of map[string]int.
// When either the slice or the map is empty, it returns false.
func AllValuesInMapStringInt(Slice []int, Map map[string]int) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringInt(x, Map); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapStringInt checks which values of an int slice are values of map[string]int.
// When either the slice or the map is empty, it returns false.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInMapStringInt(Slice []int, Map map[string]int) (map[int][]string, bool) {
	values := make(map[int][]string)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[int][]string{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapStringInt(x, Map); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}

// AnyValueInMapStringString checks if any of the values of a string slice is a value of map[string]string.
// When either the slice or the map is empty, it returns false.
func AnyValueInMapStringString(Slice []string, Map map[string]string) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringString(x, Map); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapStringString checks if all values of a string slice are values of map[string]string.
// When either the slice or the map is empty, it returns false.
func AllValuesInMapStringString(Slice []string, Map map[string]string) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringString(x, Map); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapStringString checks which values of a string slice are values of map[string]string.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When either the slice or the map is empty, it returns false.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
func WhichValuesInMapStringString(Slice []string, Map map[string]string) (map[string][]string, bool) {
	values := make(map[string][]string)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[string][]string{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapStringString(x, Map); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}

// AnyValueInMapStringFloat64 checks if any of the values of a float64 slice is a value of map[string]float64.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When either the slice or the map is empty, it returns false.
func AnyValueInMapStringFloat64(Slice []float64, Map map[string]float64, Epsilon float64) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringFloat64(x, Map, Epsilon); ok {
			return true
		}
	}
	return false
}

// AllValuesInMapStringFloat64 checks if all values of a float64 slice are values of map[string]float64.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When either the slice or the map is empty, it returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AllValuesInMapStringFloat64(Slice []float64, Map map[string]float64, Epsilon float64) bool {
	if len(Slice) == 0 || len(Map) == 0 {
		return false
	}
	for _, x := range Slice {
		if _, ok := IsValueInMapStringFloat64(x, Map, Epsilon); !ok {
			return false
		}
	}
	return true
}

// WhichValuesInMapStringFloat64 checks which values of an int slice are values of map[int]float64.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// When either the slice or the map is empty, it returns false.
// The function returns a tuple with a map with values from Slice as keys and slices of keys in the map that have these values,
// and a boolean value (true if the returned map is not empty).
// BEWARE! Do note that we work with floats, so it's safest to round them before using this function. Thus, this
// WhichValuesInMapStringFloat64({[]float64{.01002, .01, .2}, map[string]float64{"a": .01, "b": .011}, .0001)
// will return the following map: map[float64][]string{.01: {"a"}, .01002: {"a"}}}
func WhichValuesInMapStringFloat64(Slice []float64, Map map[string]float64, Epsilon float64) (map[float64][]string, bool) {
	values := make(map[float64][]string)
	var exists bool

	if len(Slice) == 0 || len(Map) == 0 {
		return map[float64][]string{}, false
	}
	for _, x := range Slice {
		if keys, ok := IsValueInMapStringFloat64(x, Map, Epsilon); ok {
			exists = true
			values[x] = keys
		}
	}
	return values, exists
}
