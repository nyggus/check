/* Package check offers functions to check various conditions, including

* check if all boolean values are true (in a slice and in a map)
* check if any boolean value is true (in a slice and in a map)
* check which of boolean values is true (in a slice and in a map)
* check if all elements of a slice have the same value
* check if all keys of a map have the same value
* check if two slices are the same (either taking into account the ordering of the slices or ignoring it)
* check if two maps are the same
* check if a value is in a slice
* check if several values are in a slice
* check if any of several values are in a slice
* check which of several values are in a slice
* check if a key-value pair is among a map's key-values
* check if a value is among a map's values
* check if several values from a slice are in a map
* check if any value from a slice are in a map
* check which values from a slice are in a map
* check if all key-value pairs from a map are in a map
* check if any of key-value pairs from a map are in a map
* check which key-value pairs from a map are in a map

In addition to these checks, the package enables one to create a unique slice (that is, with unique values) out of a slice.

The package works with the following slices:

* []string
* []int
* []float64
* []bool

and maps:

* map[string]string
* map[string]int
* map[string]float64
* map[string]bool
* map[int]string
* map[int]int
* map[int]float64
* map[int]bool

Floats are compared using an epsilon value, meaning that two floats are considered equal when their absolute difference is less than or equal to epsilon.
Since using floats as map keys is not recommended, the package does not with with such maps.
*/

package check

import (
	"math"
	"sort"
)

// AreEqualSlicesFloat64 compares two float64 slices.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// The function ignores sorting, so compares both values and sorting of the slices.
// Thus, if the slices have the same values but different orders, they are not considered the same.
// When both slices has zero length, true is returned.
// If you want the function to sort the slices first, use AreEqualSortedSlicesFloat64.
func AreEqualSlicesFloat64(Slice1, Slice2 []float64, Epsilon float64) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	for i := range Slice1 {
		if math.Abs(Slice1[i]-Slice2[i]) > Epsilon {
			return false
		}
	}
	return true
}

// AreEqualSortedSlicesFloat64 compares two float64 slices.
// The Epsilon parameter sets the accuracy of the comparison of two floats.
// The function takes into account sorting, meaning that if they have the same values, they are considered the same
// even if they are differently ordered. When both slices has zero length, true is returned.
func AreEqualSortedSlicesFloat64(Slice1, Slice2 []float64, Epsilon float64) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	if !sort.Float64sAreSorted(Slice1) {
		sort.Float64s(Slice1)
	}
	if !sort.Float64sAreSorted(Slice2) {
		sort.Float64s(Slice2)
	}
	for i := range Slice1 {
		if math.Abs(Slice1[i]-Slice2[i]) > Epsilon {
			return false
		}
	}
	return true
}

// AreEqualSlicesInt compares two int slices. The function ignores sorting, so compares both values and sorting of the slices.
// If you want the function to sort the slices first, use AreEqualSortedSlicesInt instead.
// When both slices has zero length, true is returned.
func AreEqualSlicesInt(Slice1, Slice2 []int) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	for i := range Slice1 {
		if Slice1[i] != Slice2[i] {
			return false
		}
	}
	return true
}

// AreEqualSortedSlicesInt compares two int slices.
// The function takes into account sorting, meaning that if they have the same values, they are considered the same
// even if they are differently ordered. When both slices has zero length, true is returned.
func AreEqualSortedSlicesInt(Slice1, Slice2 []int) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	if !sort.IntsAreSorted(Slice1) {
		sort.Ints(Slice1)
	}
	if !sort.IntsAreSorted(Slice2) {
		sort.Ints(Slice2)
	}
	for i := range Slice1 {
		if Slice1[i] != Slice2[i] {
			return false
		}
	}
	return true
}

// AreEqualSlicesString compares two string slices. The function ignores sorting, so compares both values and sorting of the slices.
// If you want the function to sort the slices, use AreEqualSortedSlicesString instead.
// When both slices has zero length, true is returned.
func AreEqualSlicesString(Slice1, Slice2 []string) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	for i := range Slice1 {
		if Slice1[i] != Slice2[i] {
			return false
		}
	}
	return true
}

// AreEqualSortedSlicesString compares two string slices.
// The function takes into account sorting, meaning that if they have the same values, they are considered the same
// even if they are differently ordered. When both slices has zero length, true is returned.
func AreEqualSortedSlicesString(Slice1, Slice2 []string) bool {
	if len(Slice1) == 0 && len(Slice2) == 0 {
		return true
	}
	if len(Slice1) != len(Slice2) {
		return false
	}
	if !sort.StringsAreSorted(Slice1) {
		sort.Strings(Slice1)
	}
	if !sort.StringsAreSorted(Slice2) {
		sort.Strings(Slice2)
	}
	for i := range Slice1 {
		if Slice1[i] != Slice2[i] {
			return false
		}
	}
	return true
}
