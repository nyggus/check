package check

import "math"

// AllKeyValuePairsInMapStringString checks if all key-value pairs from one map[string]string are in another map[string]string.
// When any of the maps is empty, the function returns false.
func AllKeyValuePairsInMapStringString(Map1, Map2 map[string]string) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if valueMap1 != valueMap2 {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapStringString checks if any key-value pair from one map[string]string is in another map[string]string.
// When any of the maps is empty, the function returns false.
func AnyKeyValuePairInMapStringString(Map1, Map2 map[string]string) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapStringString checks which key-value pairs from one map[string]string is in another map[string]string.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
func WhichKeyValuePairsInMapStringString(Map1, Map2 map[string]string) (map[string]string, bool) {
	keys := make(map[string]string)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}

// AllKeyValuePairsInMapStringInt checks if all key-value pairs from one map[string]int are in another map[string]int.
// When any of the maps is empty, the function returns false.
func AllKeyValuePairsInMapStringInt(Map1, Map2 map[string]int) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if valueMap1 != valueMap2 {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapStringInt checks if any key-value pair from one map[string]int is in another map[string]int.
// When any of the maps is empty, the function returns false.
func AnyKeyValuePairInMapStringInt(Map1, Map2 map[string]int) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapStringInt checks which key-value pair from one map[string]int is in another map[string]int.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
func WhichKeyValuePairsInMapStringInt(Map1, Map2 map[string]int) (map[string]int, bool) {
	keys := make(map[string]int)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}

// AllKeyValuePairsInMapStringFloat64 checks if all key-value pairs from one map[string]float64 are in another map[string]float64.
// When any of the maps is empty, the function returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AllKeyValuePairsInMapStringFloat64(Map1, Map2 map[string]float64, Epsilon float64) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if math.Abs(valueMap1-valueMap2) > Epsilon {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapStringFloat64 checks if any key-value pair from one map[string]float64 is in another map[string]float64.
// When any of the maps is empty, the function returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AnyKeyValuePairInMapStringFloat64(Map1, Map2 map[string]float64, Epsilon float64) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if math.Abs(valueMap1-valueMap2) <= Epsilon {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapStringFloat64 checks if any key-value pair from one map[string]float64 is in another map[string]float64.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func WhichKeyValuePairsInMapStringFloat64(Map1, Map2 map[string]float64, Epsilon float64) (map[string]float64, bool) {
	keys := make(map[string]float64)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if math.Abs(valueMap1-valueMap2) <= Epsilon {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}

// AllKeyValuePairsInMapIntString checks if all key-value pairs from one map[int]string are in another map[int]string.
// When any of the maps is empty, the function returns false.
func AllKeyValuePairsInMapIntString(Map1, Map2 map[int]string) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if valueMap1 != valueMap2 {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapIntString checks if any key-value pair from one map[int]string is in another map[int]string.
// When any of the maps is empty, the function returns false.
func AnyKeyValuePairInMapIntString(Map1, Map2 map[int]string) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapIntString checks if any key-value pair from one map[int]string is in another map[int]string.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
func WhichKeyValuePairsInMapIntString(Map1, Map2 map[int]string) (map[int]string, bool) {
	keys := make(map[int]string)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}

// AllKeyValuePairsInMapIntInt checks if all key-value pairs from one map[int]int are in another map[int]int.
// When any of the maps is empty, the function returns false.
func AllKeyValuePairsInMapIntInt(Map1, Map2 map[int]int) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if valueMap1 != valueMap2 {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapIntInt checks if any key-value pair from one map[int]int is in another map[int]int.
// When any of the maps is empty, the function returns false.
func AnyKeyValuePairInMapIntInt(Map1, Map2 map[int]int) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapIntInt checks if any key-value pair from one map[int]int is in another map[int]int.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
func WhichKeyValuePairsInMapIntInt(Map1, Map2 map[int]int) (map[int]int, bool) {
	keys := make(map[int]int)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if valueMap1 == valueMap2 {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}

// AllKeyValuePairsInMapIntFloat64 checks if all key-value pairs from one map[int]float64 are in another map[int]float64.
// When any of the maps is empty, the function returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AllKeyValuePairsInMapIntFloat64(Map1, Map2 map[int]float64, Epsilon float64) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if !ok {
			return false
		}
		if math.Abs(valueMap1-valueMap2) > Epsilon {
			return false
		}
	}
	return true
}

// AnyKeyValuePairInMapIntFloat64 checks if any key-value pair from one map[int]float64 is in another map[int]float64.
// When any of the maps is empty, the function returns false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func AnyKeyValuePairInMapIntFloat64(Map1, Map2 map[int]float64, Epsilon float64) bool {
	if len(Map1) == 0 || len(Map2) == 0 {
		return false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if math.Abs(valueMap1-valueMap2) <= Epsilon {
				return true
			}
		}
	}
	return false
}

// WhichKeyValuePairsInMapIntFloat64 checks if any key-value pair from one map[int]float64 is in another map[int]float64.
// Returns a tuple of a map with the found key-value pairs, and true if the map is not empty.
// When any of the maps is empty, the function returns empty slice and false.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
func WhichKeyValuePairsInMapIntFloat64(Map1, Map2 map[int]float64, Epsilon float64) (map[int]float64, bool) {
	keys := make(map[int]float64)
	var exists bool

	if len(Map1) == 0 || len(Map2) == 0 {
		return keys, false
	}
	for key, valueMap1 := range Map1 {
		valueMap2, ok := Map2[key]
		if ok {
			if math.Abs(valueMap1-valueMap2) <= Epsilon {
				exists = true
				keys[key] = valueMap1
			}
		}
	}
	return keys, exists
}
