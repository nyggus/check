package check

import (
	"fmt"
	"testing"
)

func TestAllValuesInIntSlice(t *testing.T) {
	tests := []struct {
		slice1   []int
		slice2   []int
		expected bool
	}{
		{[]int{10}, []int{10, 20}, true},
		{[]int{10}, []int{11, 20}, false},
		{[]int{}, []int{10, 20}, true},
		{[]int{}, []int{}, true},
		{[]int{10}, []int{}, false},
		{[]int{10, 20, 30}, []int{10, 20, 30}, true},
		{[]int{10, 10, 10}, []int{10, 20, 30}, true},
		{[]int{10, 20, 30}, []int{10, 20, 31}, false},
	}
	for _, test := range tests {
		if AllValuesInIntSlice(test.slice1, test.slice2) != test.expected {
			t.Errorf("AllValuesInIntSlice(%v, %v) should be %v", test.slice1, test.slice2, test.expected)
		}
	}
}

func ExampleAllValuesInIntSlice() {
	fmt.Println(AllValuesInIntSlice([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	fmt.Println(AllValuesInIntSlice([]int{1, 2, 3}, []int{1, 3, 4, 5}))
	// Output:
	// true
	// false
}

func TestAllValuesInStringSlice(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected bool
	}{
		{[]string{"10"}, []string{"10", "20"}, true},
		{[]string{"10"}, []string{"11", "20"}, false},
		{[]string{}, []string{"10", "20"}, true},
		{[]string{}, []string{}, true},
		{[]string{"10"}, []string{}, false},
		{[]string{"10", "20", "30"}, []string{"10", "20", "30"}, true},
		{[]string{"10", "10", "10"}, []string{"10", "20", "30"}, true},
		{[]string{"10", "20", "30"}, []string{"10", "20", "31"}, false},
		{[]string{"10 ", "20", "30"}, []string{"10", "20", "30"}, false},
	}
	for _, test := range tests {
		if AllValuesInStringSlice(test.slice1, test.slice2) != test.expected {
			t.Errorf("AllValuesInStringSlice(%v, %v) should be %v", test.slice1, test.slice2, test.expected)
		}
	}
}

func ExampleAllValuesInStringSlice() {
	fmt.Println(AllValuesInStringSlice([]string{"a", "b"}, []string{"a", "b", "c"}))
	fmt.Println(AllValuesInStringSlice([]string{"a", "b"}, []string{"a", "c", "d"}))
	// Output:
	// true
	// false
}

func TestAllValuesInFloat64Slice(t *testing.T) {
	tests := []struct {
		slice1   []float64
		slice2   []float64
		epsilon  float64
		expected bool
	}{
		{[]float64{10}, []float64{10, 20}, .00000001, true},
		{[]float64{10}, []float64{11, 20}, .00000001, false},
		{[]float64{}, []float64{10, 20}, .00000001, true},
		{[]float64{}, []float64{}, .00000001, true},
		{[]float64{10}, []float64{}, .00000001, false},
		{[]float64{10, 20, 30}, []float64{10, 20, 30}, .00000001, true},
		{[]float64{10, 10, 10}, []float64{10, 20, 30}, .00000001, true},
		{[]float64{.10, .20, .30}, []float64{.10, .20, .31}, .00000001, false},
		{[]float64{.10, .20, .30}, []float64{.10, .20, .31}, 1, true},
	}
	for _, test := range tests {
		if AllValuesInFloat64Slice(test.slice1, test.slice2, test.epsilon) != test.expected {
			t.Errorf("AllValuesInFloat64Slice(%v, %v) should be %v", test.slice1, test.slice2, test.expected)
		}
	}
}

func ExampleAllValuesInFloat64Slice() {
	fmt.Println(AllValuesInFloat64Slice([]float64{1, 2, 3}, []float64{1, 2, 3, 4}, 0))
	fmt.Println(AllValuesInFloat64Slice([]float64{.001, .002, .0022}, []float64{.001, .002}, 0))
	fmt.Println(AllValuesInFloat64Slice([]float64{.001, .002, .0022}, []float64{.001, .002}, .001))
	// Output:
	// true
	// false
	// true
}

func TestAnyValueInIntSlice(t *testing.T) {
	tests := []struct {
		slice1   []int
		slice2   []int
		expected bool
	}{
		{[]int{10, 20, 70}, []int{10, 20}, true},
		{[]int{10}, []int{11, 20}, false},
		{[]int{10, 11, 15, 10, 2000}, []int{12, 22}, false},
		{[]int{10, 11, 15, 10, 2000}, []int{11, 22}, true},
		{[]int{}, []int{10, 20}, false},
		{[]int{}, []int{}, false},
		{[]int{10}, []int{}, false},
		{[]int{10, 20, 30}, []int{10}, true},
		{[]int{10, 10, 10}, []int{20, 20, 30}, false},
		{[]int{10, 10, 10}, []int{20, 20, 10}, true},
		{[]int{1000000, 2000000, 3000000}, []int{1000000, 20, 31}, true},
	}
	for _, test := range tests {
		if AnyValueInIntSlice(test.slice1, test.slice2) != test.expected {
			t.Errorf("AnyValueInIntSlice(%v, %v) should be %v", test.slice1, test.slice2, test.expected)
		}
	}
}

func ExampleAnyValueInIntSlice() {
	fmt.Println(AnyValueInIntSlice([]int{10, 20}, []int{10, 30, 50}))
	fmt.Println(AnyValueInIntSlice([]int{10, 11, 15, 10, 2000}, []int{12, 22}))
	fmt.Println(AnyValueInIntSlice([]int{10, 10, 10}, []int{20, 20, 10}))
	// Output:
	// true
	// false
	// true
}

func TestWhichValuesInIntSlice(t *testing.T) {
	tests := []struct {
		slice1         []int
		slice2         []int
		expectedBool   bool
		expectedValues map[int][]int
	}{
		{[]int{10, 20, 70}, []int{10, 20}, true, map[int][]int{10: {0}, 20: {1}}},
		{[]int{10, 20, 70}, []int{10, 20, 10, 20}, true, map[int][]int{10: {0, 2}, 20: {1, 3}}},
		{[]int{10}, []int{11, 20}, false, map[int][]int{}},
		{[]int{10, 11, 15, 10, 2000}, []int{12, 22}, false, map[int][]int{}},
		{[]int{10, 11, 15, 10, 2000}, []int{11, 22}, true, map[int][]int{11: {0}}},
		{[]int{}, []int{10, 20}, false, map[int][]int{}},
		{[]int{}, []int{}, false, map[int][]int{}},
		{[]int{10}, []int{}, false, map[int][]int{}},
		{[]int{10, 20, 30}, []int{10}, true, map[int][]int{10: {0}}},
		{[]int{10, 10, 10}, []int{20, 20, 30}, false, map[int][]int{}},
		{[]int{10, 10, 10}, []int{20, 20, 10}, true, map[int][]int{10: {2}}},
		{[]int{1000000, 2000000, 3000000}, []int{1000000, 20, 31}, true, map[int][]int{1000000: {0}}},
		{[]int{1, 2, 3}, []int{1, 1, 2, 1, 1, 1, 2, 7, 33, 2, 2, 67, 90, 2}, true, map[int][]int{1: {0, 1, 3, 4, 5}, 2: {2, 6, 9, 10, 13}}},
		{[]int{1, 2, 3}, []int{1, 1, 2, 1, 1, 1, 2, 7, 33, 12, 33, 67, 90}, true, map[int][]int{1: {0, 1, 3, 4, 5}, 2: {2, 6}}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInIntSlice(test.slice1, test.slice2)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInIntSlice(%v, %v) should be %v with values of %v, not %v with %v",
				test.slice1, test.slice2, test.expectedBool, test.expectedValues, ok, values)
		}
		if ok {
			for value, sliceWithIndices := range values {
				if IsValueInIntSlice(value, test.slice2) {
					if !AreEqualSlicesInt(sliceWithIndices, test.expectedValues[value]) {
						t.Errorf(
							"WhichValuesInIntSlice(%v, %v) should be %v with values of %v, not %v with %v",
							test.slice1, test.slice2, test.expectedBool, test.expectedValues, ok, values)
					}
				}
			}
		}
	}
}

func ExampleWhichValuesInIntSlice() {
	values, ok := WhichValuesInIntSlice([]int{10, 20}, []int{10, 10, 50})
	fmt.Println(ok, values)
	values, ok = WhichValuesInIntSlice([]int{10, 11, 15, 10, 2000}, []int{12, 22})
	fmt.Println(ok, values)
	values, ok = WhichValuesInIntSlice([]int{10, 10, 10, 20}, []int{20, 20, 10})
	fmt.Println(ok, values)
	// Output:
	// true map[10:[0 1]]
	// false map[]
	// true map[10:[2] 20:[0 1]]
}

func TestAnyValueInStringSlice(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected bool
	}{
		{[]string{"10"}, []string{"10", "20"}, true},
		{[]string{"10"}, []string{"11", "20"}, false},
		{[]string{"10", "10", "10"}, []string{"11", "20"}, false},
		{[]string{"10", "10", "10", "20"}, []string{"11", "20"}, true},
		{[]string{}, []string{"10", "20"}, false},
		{[]string{}, []string{}, false},
		{[]string{"10"}, []string{}, false},
		{[]string{"10", "20", "30"}, []string{"10", "20", "30"}, true},
		{[]string{"10", "10", "10"}, []string{"10", "20", "30"}, true},
		{[]string{"11", "21", "31"}, []string{"10", "20", "30"}, false},
		{[]string{"10 ", "21", "31"}, []string{"10", "20", "30"}, false},
		{[]string{"10 ", "21", "31"}, []string{"10 ", "20", "30"}, true},
	}
	for _, test := range tests {
		if AnyValueInStringSlice(test.slice1, test.slice2) != test.expected {
			t.Errorf("AnyValueInStringSlice(%v, %v) should be %v", test.slice1, test.slice2, test.expected)
		}
	}
}

func ExampleAnyValueInStringSlice() {
	fmt.Println(AnyValueInStringSlice([]string{"10"}, []string{"10", "20"}))
	fmt.Println(AnyValueInStringSlice([]string{"10 ", "21", "31"}, []string{"10", "20", "30"}))
	fmt.Println(AnyValueInStringSlice([]string{"10 ", "21", "31"}, []string{"10 ", "20", "30"}))
	// Output:
	// true
	// false
	// true
}

func TestWhichValuesInStringSlice(t *testing.T) {
	tests := []struct {
		slice1         []string
		slice2         []string
		expectedBool   bool
		expectedValues map[string][]int
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b"}, true, map[string][]int{"a": {0}, "b": {1}}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "a", "b"}, true, map[string][]int{"a": {0, 2}, "b": {1, 3}}},
		{[]string{"a"}, []string{"z", "b"}, false, map[string][]int{}},
		{[]string{"a", "b", "c"}, []string{"Shout", "Bamalama"}, false, map[string][]int{}},
		{[]string{"a", "b", "c", "d"}, []string{"c"}, true, map[string][]int{"c": {0}}},
		{[]string{}, []string{"a", "b", "c"}, false, map[string][]int{}},
		{[]string{}, []string{}, false, map[string][]int{}},
		{[]string{"a", "b", "c"}, []string{}, false, map[string][]int{}},
		{[]string{"a", "a", "a"}, []string{"z", "c", "c"}, false, map[string][]int{}},
		{[]string{"a", "a", "W"}, []string{"A", "b", "W"}, true, map[string][]int{"W": {2}}},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "b", "a", "a", "a", "a", "a", "a", "a", "a", "DDD"},
			true,
			map[string][]int{"a": {0, 3, 4, 5, 6, 7, 8, 9, 10}, "b": {1, 2}},
		},
	}
	for _, test := range tests {
		values, ok := WhichValuesInStringSlice(test.slice1, test.slice2)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInStringSlice(%v, %v) should be %v with values of %v, not %v with %v",
				test.slice1, test.slice2, test.expectedBool, test.expectedValues, ok, values)
		}
		if ok {
			for value, sliceWithIndices := range values {
				if IsValueInStringSlice(value, test.slice2) {
					if !AreEqualSlicesInt(sliceWithIndices, test.expectedValues[value]) {
						t.Errorf(
							"WhichValuesInStringSlice(%v, %v) should be %v with values of %v, not %v with %v",
							test.slice1, test.slice2, test.expectedBool, test.expectedValues, ok, values)
					}
				}
			}
		}
	}
}

func ExampleWhichValuesInStringSlice() {
	values, ok := WhichValuesInStringSlice([]string{"10", "Shout Bamalama!"}, []string{"10", "Shout Bamalama!", "50"})
	fmt.Println(ok, values)
	values, ok = WhichValuesInStringSlice([]string{"Sing", "a", "sing", "about", "the", "sun!"}, []string{"sing", "sun", "the"})
	fmt.Println(ok, values)
	values, ok = WhichValuesInStringSlice([]string{"a", "a", "b"}, []string{"A", "A", "B"})
	fmt.Println(ok, values)
	// Output:
	// true map[10:[0] Shout Bamalama!:[1]]
	// true map[sing:[0] the:[2]]
	// false map[]
}

func TestAnyValueInFloat64Slice(t *testing.T) {
	tests := []struct {
		slice1   []float64
		slice2   []float64
		epsilon  float64
		expected bool
	}{
		{[]float64{10, 20, 70}, []float64{10, 20}, .00000001, true},
		{[]float64{10}, []float64{11, 20}, .00000001, false},
		{[]float64{10, 11, 15, 10, 2000}, []float64{12, 22}, .00000001, false},
		{[]float64{10, 11, 15, 10, 2000}, []float64{11, 22}, .00000001, true},
		{[]float64{}, []float64{10, 20}, .00000001, false},
		{[]float64{}, []float64{}, .00000001, false},
		{[]float64{10}, []float64{}, .00000001, false},
		{[]float64{10, 20, 30}, []float64{10}, .00000001, true},
		{[]float64{.1, .2, .2}, []float64{.2, .11, .3}, .00000001, true},
		{[]float64{.1, .2, .2}, []float64{.3, .11, .3}, .1, true},
		{[]float64{.1, .2, .2}, []float64{.3, .11, .3}, 0, false},
		{[]float64{1000000, 2000000, 3000000}, []float64{1000000, 20, 31}, 1, true},
	}
	for _, test := range tests {
		if AnyValueInFloat64Slice(test.slice1, test.slice2, test.epsilon) != test.expected {
			t.Errorf("TestAnyValueInFloat64Slice(%v, %v, %v) should be %v", test.slice1, test.slice2, test.epsilon, test.expected)
		}
	}
}

func ExampleAnyValueInFloat64Slice() {
	fmt.Println(AnyValueInFloat64Slice([]float64{10, 20, 70}, []float64{10, 20}, .00000001))
	fmt.Println(AnyValueInFloat64Slice([]float64{.1, .2, .2}, []float64{.3, .11, .3}, .00000001))
	fmt.Println(AnyValueInFloat64Slice([]float64{.1, .2, .2}, []float64{.3, .11, .3}, .1))
	// Output:
	// true
	// false
	// true
}

func TestWhichValuesInFloat64Slice(t *testing.T) {
	tests := []struct {
		slice1         []float64
		slice2         []float64
		epsilon        float64
		expectedBool   bool
		expectedValues map[float64][]int
	}{
		{[]float64{.10, .20, .70}, []float64{.10, .20}, 0, true, map[float64][]int{.10: {0}, .20: {1}}},
		{[]float64{.10, .20, .70}, []float64{.10, .20, .10, .20}, 0, true, map[float64][]int{.10: {0, 2}, .20: {1, 3}}},
		{[]float64{.10}, []float64{.11, 2.}, 0, false, map[float64][]int{}},
		{[]float64{.10, .11, .15, .10, 20.}, []float64{.12, .22}, 0, false, map[float64][]int{}},
		{[]float64{.10, .11, .15, .10, 20.}, []float64{.11, .22}, 0, true, map[float64][]int{.11: {0}}},
		{[]float64{}, []float64{.10, .20}, 0, false, map[float64][]int{}},
		{[]float64{}, []float64{}, 0, false, map[float64][]int{}},
		{[]float64{.10}, []float64{}, 0, false, map[float64][]int{}},
		{[]float64{.10, .20, .30}, []float64{.10}, 0, true, map[float64][]int{.10: {0}}},
		{[]float64{.10, .10, .10}, []float64{.20, .20, .30}, 0, false, map[float64][]int{}},
		{[]float64{.10, .10, .10}, []float64{.20, .20, .10}, 0, true, map[float64][]int{.10: {2}}},
		{[]float64{.1, 3.2, 3.2}, []float64{.3, .11, .3}, .1, true, map[float64][]int{.1: {1}}},
		{[]float64{1000000, 2000000, 3000000}, []float64{1000000, 20, 31}, 0, true, map[float64][]int{1000000: {0}}},
		{
			[]float64{.001, .002, .003},
			[]float64{.001, .001, .002, .001, .001, .001, .002, .007, .0033, .002, .002, .0067, .009, .002},
			0,
			true,
			map[float64][]int{.001: {0, 1, 3, 4, 5}, .002: {2, 6, 9, 10, 13}},
		},
		// Note the below tests; they check various scenarios with the Epsilon
		{[]float64{.10, .20, .30}, []float64{.10, .101}, 0, true, map[float64][]int{.10: {0}}},
		{[]float64{.10, .20, .30}, []float64{.10, .101}, .1, true, map[float64][]int{.10: {0, 1}}},
		{[]float64{.10, .20, .30, .101}, []float64{.10, .101}, .1, true, map[float64][]int{.10: {0, 1}}},
		{[]float64{.10, .20, .30, .101}, []float64{.10, .101}, 0, true, map[float64][]int{.10: {0}, .101: {1}}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInFloat64Slice(test.slice1, test.slice2, test.epsilon)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInFloat64Slice(%v, %v, %v) should be %v with values of %v, not %v with %v",
				test.slice1, test.slice2, test.epsilon, test.expectedBool, test.expectedValues, ok, values)
		}
		if ok {
			for value, sliceWithIndices := range values {
				if IsValueInFloat64Slice(value, test.slice2, test.epsilon) {
					if !AreEqualSlicesInt(sliceWithIndices, test.expectedValues[value]) {
						t.Errorf(
							"WhichValuesInFloat64Slice(%v, %v, %v) should be %v with values of %v, not %v with %v",
							test.slice1, test.slice2, test.epsilon, test.expectedBool, test.expectedValues, ok, values)
					}
				}
			}
		}
	}
}

func ExampleWhichValuesInFloat64Slice() {
	values, ok := WhichValuesInFloat64Slice([]float64{.10, .20, .70}, []float64{.10, .20}, 0)
	fmt.Println(ok, values)
	values, ok = WhichValuesInFloat64Slice([]float64{.1, .2, .2}, []float64{.3, .11, .3}, .00000001)
	fmt.Println(ok, values)
	values, ok = WhichValuesInFloat64Slice([]float64{.10, .20, .30, .101}, []float64{.10, .101}, .1)
	fmt.Println(ok, values)
	// Output:
	// true map[0.1:[0] 0.2:[1]]
	// false map[]
	// true map[0.1:[0 1]]
}

func TestAnyValueInMapIntInt(t *testing.T) {
	tests := []struct {
		slice    []int
		Map      map[int]int
		expected bool
	}{
		{[]int{10, 20, 70}, map[int]int{1: 10, 2: 20}, true},
		{[]int{10, 20, 70}, map[int]int{1: 30, 2: 40}, false},
		{[]int{10, 20, 70}, map[int]int{}, false},
		{[]int{}, map[int]int{1: 10, 2: 20}, false},
		{[]int{1, 2, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}, true},
	}
	for _, test := range tests {
		if AnyValueInMapIntInt(test.slice, test.Map) != test.expected {
			t.Errorf("AnyValueInMapIntInt(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAnyValueInMapIntInt() {
	fmt.Println(AnyValueInMapIntInt([]int{10, 20, 70}, map[int]int{1: 10, 2: 20}))
	fmt.Println(AnyValueInMapIntInt([]int{10, 20, 70}, map[int]int{1: 30, 2: 40}))
	fmt.Println(AnyValueInMapIntInt([]int{1, 2, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}))
	// Output:
	// true
	// false
	// true
}

func TestWhichValuesInMapIntInt(t *testing.T) { //TODO
	tests := []struct {
		slice          []int
		Map            map[int]int
		expectedBool   bool
		expectedValues map[int][]int
	}{
		{[]int{10, 20, 70}, map[int]int{1: 10, 2: 20}, true, map[int][]int{10: {1}, 20: {2}}},
		{[]int{10, 20, 70}, map[int]int{1: 10, 2: 20, 3: 10, 4: 10, 5: 20, 6: 30, 7: 10}, true, map[int][]int{10: {1, 3, 4, 7}, 20: {2, 5}}},
		{[]int{10, 20, 70}, map[int]int{1: 10, 2: 10, 3: 10, 4: 10, 5: 10, 6: 10, 7: 10}, true, map[int][]int{10: {1, 2, 3, 4, 5, 6, 7}}},
		{[]int{10, 20, 70}, map[int]int{}, false, map[int][]int{10: {}, 20: {}}},
		{[]int{}, map[int]int{1: 10, 2: 20}, false, map[int][]int{}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapIntInt(test.slice, test.Map)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapIntInt(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesInt(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapIntInt(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}

func ExampleWhichValuesInMapIntInt() {
	values, ok := WhichValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 10, 2: 20})
	fmt.Println(ok, values)
	values, ok = WhichValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 11, 2: 21})
	fmt.Println(ok, values)
	// Output:
	// true map[10:[1] 20:[2]]
	// false map[]
}

func TestAllValuesInMapIntInt(t *testing.T) {
	tests := []struct {
		slice    []int
		Map      map[int]int
		expected bool
	}{
		{[]int{10, 20, 70}, map[int]int{1: 10, 2: 20}, false},
		{[]int{10, 20, 70}, map[int]int{1: 30, 2: 40}, false},
		{[]int{10, 20, 70}, map[int]int{}, false},
		{[]int{10, 20, 70}, map[int]int{1: 20, 2: 10, 5: 70}, true},
		{[]int{}, map[int]int{1: 10, 2: 20}, false},
		{[]int{1, 2, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}, false},
		{[]int{10, 1, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}, true},
	}
	for _, test := range tests {
		if AllValuesInMapIntInt(test.slice, test.Map) != test.expected {
			t.Errorf("AllValuesInMapIntInt(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAllValuesInMapIntInt() {
	fmt.Println(AllValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 10, 2: 20}))
	fmt.Println(AllValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 20, 2: 10, 5: 70}))
	fmt.Println(AllValuesInMapIntInt([]int{10, 1, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}))
	// Output:
	// false
	// true
	// true
}

func TestAnyValueInMapIntString(t *testing.T) {
	tests := []struct {
		slice    []string
		Map      map[int]string
		expected bool
	}{
		{[]string{"a", "b"}, map[int]string{1: "a", 2: "g"}, true},
		{[]string{"a", "b"}, map[int]string{1: "v", 2: "g"}, false},
		{[]string{"a", "b"}, map[int]string{}, false},
		{[]string{}, map[int]string{1: "v", 2: "g"}, false},
		{[]string{"a", "b"}, map[int]string{1: "b", 2: "b"}, true},
		{[]string{"a", "b"}, map[int]string{1: "a", 2: "b"}, true},
	}
	for _, test := range tests {
		if AnyValueInMapIntString(test.slice, test.Map) != test.expected {
			t.Errorf("AnyValueInMapIntString(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAnyValueInMapIntString() {
	fmt.Println(AnyValueInMapIntString([]string{"a", "b"}, map[int]string{1: "v", 2: "g"}))
	fmt.Println(AnyValueInMapIntString([]string{"a", "b"}, map[int]string{1: "a", 2: "g"}))
	// Output:
	// false
	// true
}

func TestAllValuesInMapIntString(t *testing.T) {
	tests := []struct {
		slice    []string
		Map      map[int]string
		expected bool
	}{
		{[]string{"a", "b"}, map[int]string{1: "a", 2: "b"}, true},
		{[]string{"a", "b"}, map[int]string{1: "a", 2: "g"}, false},
		{[]string{"a", "b"}, map[int]string{}, false},
		{[]string{}, map[int]string{1: "v", 2: "g"}, false},
		{[]string{"a", "b", "C"}, map[int]string{1: "a", 2: "b", 3: "C"}, true},
		{[]string{"a", "C"}, map[int]string{1: "a", 2: "b"}, false},
	}
	for _, test := range tests {
		if AllValuesInMapIntString(test.slice, test.Map) != test.expected {
			t.Errorf("AllValuesInMapIntString(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAllValuesInMapIntString() {
	fmt.Println(AllValuesInMapIntString([]string{"a", "b", "C"}, map[int]string{1: "a", 2: "b", 3: "C"}))
	fmt.Println(AllValuesInMapIntString([]string{"a", "C"}, map[int]string{1: "a", 2: "b"}))
	// Output:
	// true
	// false
}

func TestWhichValuesInMapIntString(t *testing.T) {
	tests := []struct {
		slice          []string
		Map            map[int]string
		expectedBool   bool
		expectedValues map[string][]int
	}{
		{[]string{"a", "b", "c"}, map[int]string{1: "a", 2: "b"}, true, map[string][]int{"a": {1}, "b": {2}}},
		{[]string{"a", "b", "c"}, map[int]string{1: "a", 2: "b", 3: "a", 4: "a", 5: "b", 6: "c", 7: "a"}, true, map[string][]int{"a": {1, 3, 4, 7}, "b": {2, 5}, "c": {6}}},
		{[]string{"a", "b", "c"}, map[int]string{1: "a", 2: "a", 3: "a", 4: "a", 5: "a", 6: "a", 7: "a"}, true, map[string][]int{"a": {1, 2, 3, 4, 5, 6, 7}}},
		{[]string{"a", "b", "c"}, map[int]string{}, false, map[string][]int{"a": {}, "b": {}}},
		{[]string{}, map[int]string{1: "a", 2: "b"}, false, map[string][]int{}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapIntString(test.slice, test.Map)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapIntString(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesInt(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapIntString(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}

func ExampleWhichValuesInMapIntString() {
	values, ok := WhichValuesInMapIntString([]string{"a", "b", "c"}, map[int]string{1: "a", 2: "b"})
	fmt.Println(ok, values)
	values, ok = WhichValuesInMapIntString([]string{"a", "b", "c"}, map[int]string{1: "a", 2: "a", 3: "b", 50: "a", 99: "a", 7: "b"})
	fmt.Println(ok, AreEqualSortedSlicesInt([]int{1, 2, 50, 99}, values["a"]), AreEqualSortedSlicesInt([]int{3, 7}, values["b"]))
	values, ok = WhichValuesInMapIntString([]string{"a", "b", "c"}, map[int]string{1: "F", 2: "H"})
	fmt.Println(ok, values)
	// Output:
	// true map[a:[1] b:[2]]
	// true true true
	// false map[]
}

func TestAnyValueInMapIntFloat64(t *testing.T) {
	tests := []struct {
		slice    []float64
		Map      map[int]float64
		epsilon  float64
		expected bool
	}{
		{[]float64{}, map[int]float64{1: .01, 2: .2}, 0, false},
		{[]float64{.01, .2}, map[int]float64{}, 0, false},
		{[]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0, true},
		{[]float64{.01, .2}, map[int]float64{1: .011, 2: .21}, 0, false},
		{[]float64{.01, .2}, map[int]float64{1: .011, 2: .21}, .01, true},
		{[]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, 0, false},
		{[]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, .01, true},
		{[]float64{.001, .02, 70.8}, map[int]float64{1: .0011, 2: .021}, .01, true},
	}
	for _, test := range tests {
		if AnyValueInMapIntFloat64(test.slice, test.Map, test.epsilon) != test.expected {
			t.Errorf("AnyValueInMapIntFloat64(%v, %v, %v) should be %v", test.slice, test.Map, test.epsilon, test.expected)
		}
	}
}
func ExampleAnyValueInMapIntFloat64() {
	fmt.Println(AnyValueInMapIntFloat64([]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, .01))
	fmt.Println(AnyValueInMapIntFloat64([]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, 0))
	// Output:
	// true
	// false
}

func TestAllValuesInMapIntFloat64(t *testing.T) {
	tests := []struct {
		slice    []float64
		Map      map[int]float64
		epsilon  float64
		expected bool
	}{
		{[]float64{}, map[int]float64{1: .01, 2: .2}, 0, false},
		{[]float64{.01, .2}, map[int]float64{}, 0, false},
		{[]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0, true},
		{[]float64{.01, .2}, map[int]float64{1: .011, 2: .21}, 0, false},
		{[]float64{.01, .2}, map[int]float64{1: .011, 2: .21}, .01, true},
		{[]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, 0, false},
		{[]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, .01, true},
		{[]float64{.001, .02, 70.8}, map[int]float64{1: .0011, 2: .021}, .01, false},
	}
	for _, test := range tests {
		if AllValuesInMapIntFloat64(test.slice, test.Map, test.epsilon) != test.expected {
			t.Errorf("AllValuesInMapIntFloat64(%v, %v, %v) should be %v", test.slice, test.Map, test.epsilon, test.expected)
		}
	}
}
func ExampleAllValuesInMapIntFloat64() {
	fmt.Println(AllValuesInMapIntFloat64([]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0))
	fmt.Println(AllValuesInMapIntFloat64([]float64{.001, .02, 70.8}, map[int]float64{1: .0011, 2: .021}, .01))
	// Output:
	// true
	// false
}

func TestWhichValuesInMapIntFloat64(t *testing.T) {
	tests := []struct {
		slice          []float64
		Map            map[int]float64
		epsilon        float64
		expectedBool   bool
		expectedValues map[float64][]int
	}{
		{[]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0, true, map[float64][]int{.01: {1}, .2: {2}}},
		{[]float64{.01, .2}, map[int]float64{1: .01, 2: .01}, 0, true, map[float64][]int{.01: {1, 2}}},
		{[]float64{.01002, .01, .2}, map[int]float64{1: .01, 2: .011}, .0001, true, map[float64][]int{.01: {1}, .01002: {1}}},
		{[]float64{.01, .2}, map[int]float64{}, 0, false, map[float64][]int{}},
		{[]float64{}, map[int]float64{}, 0, false, map[float64][]int{}},
		{
			[]float64{.01, .2, .3},
			map[int]float64{1: .01, 2: .01, 5: .01, 7: .01, 11: .2, 12: .01, 90: .2, 100: 1.1},
			0,
			true,
			map[float64][]int{.01: {1, 2, 5, 7, 12}, .2: {11, 90}}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapIntFloat64(test.slice, test.Map, test.epsilon)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapIntFloat64(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesInt(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapIntFloat64(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}
func ExampleWhichValuesInMapIntFloat64() {
	values, ok := WhichValuesInMapIntFloat64([]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0)
	fmt.Println(ok, values)
	// Watch this!
	values, ok = WhichValuesInMapIntFloat64([]float64{.01002, .01, .2}, map[int]float64{1: .01, 2: .011}, .0001)
	fmt.Println(ok, values)
	// Output:
	// true map[0.01:[1] 0.2:[2]]
	// true map[0.01:[1] 0.01002:[1]]
}

func TestAnyValueInMapStringInt(t *testing.T) {
	tests := []struct {
		slice    []int
		Map      map[string]int
		expected bool
	}{
		{[]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}, true},
		{[]int{10, 20, 70}, map[string]int{"1": 30, "2": 40}, false},
		{[]int{10, 20, 70}, map[string]int{}, false},
		{[]int{}, map[string]int{"1": 10, "2": 20}, false},
		{[]int{1, 2, 7000}, map[string]int{"1": 10, "2": 10, "3": 1, "5": 1, "100": 7000}, true},
	}
	for _, test := range tests {
		if AnyValueInMapStringInt(test.slice, test.Map) != test.expected {
			t.Errorf("AnyValueInMapStringInt(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAnyValueInMapStringInt() {
	fmt.Println(AnyValueInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}))
	fmt.Println(AnyValueInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 30, "2": 40}))
	fmt.Println(AnyValueInMapStringInt([]int{1, 2, 7000}, map[string]int{"1": 10, "2": 10, "3": 1, "5": 1, "100": 7000}))
	// Output:
	// true
	// false
	// true
}

func TestAllValuesInMapStringInt(t *testing.T) {
	tests := []struct {
		slice    []int
		Map      map[string]int
		expected bool
	}{
		{[]int{10, 20, 70}, map[string]int{"1": 10, "2": 20, "a": 70}, true},
		{[]int{10, 20, 70}, map[string]int{"1": 30, "2": 40}, false},
		{[]int{10, 20, 70}, map[string]int{}, false},
		{[]int{}, map[string]int{"1": 10, "2": 20}, false},
		{[]int{1, 2, 7000}, map[string]int{"1": 10, "2": 10, "3": 1, "5": 1, "100": 7000}, false},
		{[]int{1, 2, 7000}, map[string]int{"1": 1, "2": 10, "3": 1, "5": 2, "100": 7000}, true},
	}
	for _, test := range tests {
		if AllValuesInMapStringInt(test.slice, test.Map) != test.expected {
			t.Errorf("AllValuesInMapStringInt(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAllValuesInMapStringInt() {
	fmt.Println(AllValuesInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 10, "2": 20, "a": 70}))
	fmt.Println(AllValuesInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}))
	fmt.Println(AllValuesInMapStringInt([]int{1, 2, 7000}, map[string]int{"1": 10, "2": 10, "3": 1, "5": 2, "100": 7000}))
	// Output:
	// true
	// false
	// true
}

func TestWhichValuesInMapStringInt(t *testing.T) {
	tests := []struct {
		slice          []int
		Map            map[string]int
		expectedBool   bool
		expectedValues map[int][]string
	}{
		{[]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}, true, map[int][]string{10: {"1"}, 20: {"2"}}},
		{[]int{10, 20, 70}, map[string]int{"1": 10, "2": 20, "3": 10, "4": 10, "5": 20, "6": 30, "7": 10}, true, map[int][]string{10: {"1", "3", "4", "7"}, 20: {"2", "5"}}},
		{[]int{10, 20, 70}, map[string]int{"1": 10, "2": 10, "3": 10, "4": 10, "5": 10, "6": 10, "7": 10}, true, map[int][]string{10: {"1", "2", "3", "4", "5", "6", "7"}}},
		{[]int{10, 20, 70}, map[string]int{}, false, map[int][]string{10: {}, 20: {}}},
		{[]int{}, map[string]int{"1": 10, "2": 20}, false, map[int][]string{}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapStringInt(test.slice, test.Map)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapStringInt(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesString(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapStringInt(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}

func ExampleWhichValuesInMapStringInt() {
	values, ok := WhichValuesInMapStringInt([]int{10, 20, 7}, map[string]int{"a": 10, "G": 7, "H": 190})
	fmt.Println(ok, values)
	values, ok = WhichValuesInMapStringInt([]int{10, 20, 7}, map[string]int{"a": 10, "G": 7, "H": 190, "ZZ": 10, "St": 10})
	fmt.Println(ok, values[7], len(values[10]))
	values, ok = WhichValuesInMapStringInt([]int{10, 20, 7}, map[string]int{"a": 110, "G": 17})
	fmt.Println(ok, values)
	// Output:
	// true map[7:[G] 10:[a]]
	// true [G] 3
	// false map[]
}

func TestAnyValueInMapStringString(t *testing.T) {
	tests := []struct {
		slice    []string
		Map      map[string]string
		expected bool
	}{
		{[]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"}, true},
		{[]string{"a", "b", "c"}, map[string]string{"1": "nothing", "2": "else", "3": "matters"}, false},
		{[]string{"a", "b", "c"}, map[string]string{}, false},
		{[]string{}, map[string]string{"1": "", "2": ""}, false},
		{[]string{"a", "b", "z"}, map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"}, true},
	}
	for _, test := range tests {
		if AnyValueInMapStringString(test.slice, test.Map) != test.expected {
			t.Errorf("AnyValueInMapStringString(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAnyValueInMapStringString() {
	fmt.Println(AnyValueInMapStringString([]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"}))
	fmt.Println(AnyValueInMapStringString([]string{"a", "b", "c"}, map[string]string{"1": "nothing", "2": "else", "3": "matters"}))
	fmt.Println(AnyValueInMapStringString([]string{"a", "b", "z"}, map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"}))
	// Output:
	// true
	// false
	// true
}

func TestAllValuesInMapStringString(t *testing.T) {
	tests := []struct {
		slice    []string
		Map      map[string]string
		expected bool
	}{
		{[]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"}, false},
		{[]string{"a", "b", "c"}, map[string]string{"1": "nothing", "2": "else", "3": "matters"}, false},
		{[]string{"a", "b", "c"}, map[string]string{}, false},
		{[]string{}, map[string]string{"1": "", "2": ""}, false},
		{[]string{"a", "b", "z"}, map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"}, true},
	}
	for _, test := range tests {
		if AllValuesInMapStringString(test.slice, test.Map) != test.expected {
			t.Errorf("AllValuesInMapStringString(%v, %v) should be %v", test.slice, test.Map, test.expected)
		}
	}
}

func ExampleAllValuesInMapStringString() {
	fmt.Println(AllValuesInMapStringString([]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"}))
	fmt.Println(AllValuesInMapStringString([]string{"a", "b", "c"}, map[string]string{"1": "nothing", "2": "else", "3": "matters"}))
	fmt.Println(AllValuesInMapStringString([]string{"a", "b", "z"}, map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"}))
	// Output:
	// false
	// false
	// true
}

func TestWhichValuesInMapStringString(t *testing.T) { //TODO
	tests := []struct {
		slice          []string
		Map            map[string]string
		expectedBool   bool
		expectedValues map[string][]string
	}{
		{[]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"}, true, map[string][]string{"a": {"1"}, "b": {"2"}}},
		{[]string{}, map[string]string{"1": "a", "2": "b"}, false, map[string][]string{}},
		{[]string{"a", "b", "c"}, map[string]string{}, false, map[string][]string{}},
		{
			[]string{"a", "Nothing like that, forget it", "c"},
			map[string]string{"1": "a", "2": "b", "Whatever": "a", "Something else": "a", "Zigi": "Nothing like that, forget it"},
			true,
			map[string][]string{"a": {"1", "Whatever", "Something else"}, "Nothing like that, forget it": {"Zigi"}},
		},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapStringString(test.slice, test.Map)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapStringString(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesString(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapStringString(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}

func ExampleWhichValuesInMapStringString() {
	values, ok := WhichValuesInMapStringString([]string{"a", "b", "c"}, map[string]string{"1": "a", "2": "b"})
	fmt.Println(ok, values)
	values, ok = WhichValuesInMapStringString(
		[]string{"a", "Nothing like that, forget it", "c"},
		map[string]string{"1": "a", "2": "b", "Whatever": "a", "Something else": "a", "Zigi": "Nothing like that, forget it"},
	)
	fmt.Println(ok, AreEqualSortedSlicesString([]string{"1", "Whatever", "Something else"}, values["a"]))
	// Output:
	// true map[a:[1] b:[2]]
	// true true
}

func TestAnyValueInMapStringFloat64(t *testing.T) {
	tests := []struct {
		slice    []float64
		Map      map[string]float64
		epsilon  float64
		expected bool
	}{
		{[]float64{}, map[string]float64{"1": .01, "2": .2}, 0, false},
		{[]float64{.01, .2}, map[string]float64{}, 0, false},
		{[]float64{.01, .2}, map[string]float64{"1": .01, "2": .2}, 0, true},
		{[]float64{.01, .2}, map[string]float64{"1": .011, "2": .21}, 0, false},
		{[]float64{.01, .2}, map[string]float64{"1": .011, "2": .21}, .01, true},
		{[]float64{.001, .02}, map[string]float64{"1": .0011, "2": .021}, 0, false},
		{[]float64{.001, .02}, map[string]float64{"1": .0011, "2": .021}, .01, true},
		{[]float64{.001, .02, 70.8}, map[string]float64{"1": .0011, "2": .021}, .01, true},
	}
	for _, test := range tests {
		if AnyValueInMapStringFloat64(test.slice, test.Map, test.epsilon) != test.expected {
			t.Errorf("AnyValueInMapStringFloat64(%v, %v, %v) should be %v", test.slice, test.Map, test.epsilon, test.expected)
		}
	}
}
func ExampleAnyValueInMapStringFloat64() {
	fmt.Println(AnyValueInMapStringFloat64([]float64{.01, .2}, map[string]float64{"1": .01, "2": .2}, 0))
	fmt.Println(AnyValueInMapStringFloat64([]float64{.001, .02, 70.8}, map[string]float64{"1": .0011, "2": .021}, .01))
	fmt.Println(AnyValueInMapStringFloat64([]float64{.01, .02, 70.8}, map[string]float64{"1": .0111, "2": .0211}, .001))
	// Output:
	// true
	// true
	// false
}

func TestAllValuesInMapStringFloat64(t *testing.T) {
	tests := []struct {
		slice    []float64
		Map      map[string]float64
		epsilon  float64
		expected bool
	}{
		{[]float64{}, map[string]float64{"1": .01, "2": .2}, 0, false},
		{[]float64{.01, .2}, map[string]float64{}, 0, false},
		{[]float64{.01, .2}, map[string]float64{"1": .01, "2": .2}, 0, true},
		{[]float64{.01, .2}, map[string]float64{"1": .011, "2": .21}, 0, false},
		{[]float64{.01, .2}, map[string]float64{"1": .011, "2": .21}, .01, true},
		{[]float64{.001, .02}, map[string]float64{"1": .0011, "2": .021}, 0, false},
		{[]float64{.001, .02}, map[string]float64{"1": .0011, "2": .021}, .01, true},
		{[]float64{.001, .02, 70.8}, map[string]float64{"1": .0011, "2": .021}, .01, false},
	}
	for _, test := range tests {
		if AllValuesInMapStringFloat64(test.slice, test.Map, test.epsilon) != test.expected {
			t.Errorf("AllValuesInMapStringFloat64(%v, %v, %v) should be %v", test.slice, test.Map, test.epsilon, test.expected)
		}
	}
}
func ExampleAllValuesInMapStringFloat64() {
	fmt.Println(AllValuesInMapStringFloat64([]float64{.01, .2}, map[string]float64{"1": .01, "2": .2}, 0))
	fmt.Println(AllValuesInMapStringFloat64([]float64{.001, .02, 70.8}, map[string]float64{"1": .0011, "2": .021}, .01))
	// Output:
	// true
	// false
}

func TestWhichValuesInMapStringFloat64(t *testing.T) {
	tests := []struct {
		slice          []float64
		Map            map[string]float64
		epsilon        float64
		expectedBool   bool
		expectedValues map[float64][]string
	}{
		{[]float64{.01, .2}, map[string]float64{"a": .01, "b": .2}, 0, true, map[float64][]string{.01: {"a"}, .2: {"b"}}},
		{[]float64{.01, .2}, map[string]float64{"a": .01, "b": .01}, 0, true, map[float64][]string{.01: {"a", "b"}}},
		{[]float64{.01002, .01, .2}, map[string]float64{"a": .01, "b": .011}, .0001, true, map[float64][]string{.01: {"a"}, .01002: {"a"}}},
		{[]float64{.01, .2}, map[string]float64{}, 0, false, map[float64][]string{}},
		{[]float64{}, map[string]float64{}, 0, false, map[float64][]string{}},
		{
			[]float64{.01, .2, .3},
			map[string]float64{"a": .01, "b": .01, "c": .01, "G": .01, "Heja": .2, "Gotcha": .01, "Blame me": .2, "Hud": 1.1},
			0,
			true,
			map[float64][]string{.01: {"a", "b", "c", "G", "Gotcha"}, .2: {"Heja", "Blame me"}}},
	}
	for _, test := range tests {
		values, ok := WhichValuesInMapStringFloat64(test.slice, test.Map, test.epsilon)
		if ok != test.expectedBool {
			t.Errorf(
				"WhichValuesInMapStringFloat64(%v, %v) should be %v with expected values of %v, not %v with %v",
				test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
			)
		}
		for key, value := range values {
			if !AreEqualSortedSlicesString(value, test.expectedValues[key]) {
				t.Errorf(
					"WhichValuesInMapStringFloat64(%v, %v) should be %v with expected values of %v, not %v with %v",
					test.slice, test.Map, test.expectedBool, test.expectedValues, ok, values,
				)
			}
		}
	}
}
func ExampleWhichValuesInMapStringFloat64() {
	values, ok := WhichValuesInMapStringFloat64([]float64{.01, .2}, map[string]float64{"a": .01, "b": .2}, 0)
	fmt.Println(ok, values)
	// Watch this!
	values, ok = WhichValuesInMapStringFloat64([]float64{.01002, .01, .2}, map[string]float64{"a": .01, "b": .011}, .0001)
	fmt.Println(ok, values)
	// Output:
	// true map[0.01:[a] 0.2:[b]]
	// true map[0.01:[a] 0.01002:[a]]
}
