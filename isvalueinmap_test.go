package check

import (
	"fmt"
	"testing"
)

func TestIsValueInMapStringString(t *testing.T) {
	tests := []struct {
		x            string
		Map          map[string]string
		expectedBool bool
		expectedKeys []string
	}{
		{"a", map[string]string{"a": "a", "b": "b", "c": "c"}, true, []string{"a"}},
		{"b", map[string]string{"a": "a", "b": "b", "c": "c"}, true, []string{"b"}},
		{"c", map[string]string{"a": "a", "b": "b", "c": "c"}, true, []string{"c"}},
		{"d", map[string]string{"a": "a", "b": "b", "c": "c"}, false, []string{}},
		{"a", map[string]string{"a": "a", "b": "a", "c": "c"}, true, []string{"a", "b"}},
		{"a", map[string]string{}, false, []string{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapStringString(test.x, test.Map)
		if ok != test.expectedBool || !AreEqualSortedSlicesString(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapStringString(%v, %v) should be %v with keys %v", test.x, test.Map, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapStringString() {
	fmt.Println(IsValueInMapStringString("c", map[string]string{"a": "a", "b": "b", "c": "c"}))
	fmt.Println(IsValueInMapStringString("d", map[string]string{"a": "a", "b": "b", "c": "c"}))
	// Output:
	// [c] true
	// [] false
}

func TestIsValueInMapStringInt(t *testing.T) {
	tests := []struct {
		x            int
		Map          map[string]int
		expectedBool bool
		expectedKeys []string
	}{
		{1, map[string]int{"a": 1, "b": 1, "c": 2}, true, []string{"a", "b"}},
		{2, map[string]int{"a": 1, "b": 1, "c": 2}, true, []string{"c"}},
		{3, map[string]int{"a": 1, "b": 1, "c": 2}, false, []string{}},
		{1, map[string]int{}, false, []string{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapStringInt(test.x, test.Map)
		if ok != test.expectedBool || !AreEqualSortedSlicesString(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapStringInt(%v, %v) should be %v with keys %v", test.x, test.Map, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapStringInt() {
	fmt.Println(IsValueInMapStringInt(2, map[string]int{"a": 1, "b": 1, "c": 2}))
	fmt.Println(IsValueInMapStringInt(3, map[string]int{"a": 1, "b": 1, "c": 2}))
	// Output:
	// [c] true
	// [] false
}

func TestIsValueInMapStringFloat64(t *testing.T) {
	tests := []struct {
		x            float64
		Map          map[string]float64
		epsilon      float64
		expectedBool bool
		expectedKeys []string
	}{
		{1, map[string]float64{"a": 1, "b": 1, "c": 2}, 0, true, []string{"a", "b"}},
		{2, map[string]float64{"a": 1, "b": 1, "c": 2}, 0, true, []string{"c"}},
		{3, map[string]float64{"a": 1, "b": 1, "c": 2}, 0, false, []string{}},
		{.01, map[string]float64{"a": .01, "b": .02, "c": .02}, 0, true, []string{"a"}},
		{.01, map[string]float64{"a": .011, "b": .02, "c": .02}, 0, false, []string{}},
		{.01, map[string]float64{"a": .011, "b": .02, "c": .02}, 0.1, true, []string{"a", "b", "c"}},
		{1, map[string]float64{}, 0, false, []string{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapStringFloat64(test.x, test.Map, test.epsilon)
		if ok != test.expectedBool || !AreEqualSortedSlicesString(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapStringFloat64(%v, %v, %v) should be %v with keys %v", test.x, test.Map, test.epsilon, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapStringFloat64() {
	fmt.Println(IsValueInMapStringFloat64(.01, map[string]float64{"a": .01, "b": .02, "c": .02}, 0))
	fmt.Println(IsValueInMapStringFloat64(.01, map[string]float64{"a": .011, "b": .02, "c": .02}, 0))
	keys, ok := IsValueInMapStringFloat64(.01, map[string]float64{"a": .011, "b": .02, "c": .02}, 0.1)
	fmt.Println(ok, len(keys))
	// Output:
	// [a] true
	// [] false
	// true 3
}

func TestIsValueInMapIntString(t *testing.T) {
	tests := []struct {
		x            string
		Map          map[int]string
		expectedBool bool
		expectedKeys []int
	}{
		{"a", map[int]string{1: "a", 2: "b", 3: "c"}, true, []int{1}},
		{"b", map[int]string{1: "a", 2: "b", 3: "c"}, true, []int{2}},
		{"c", map[int]string{1: "a", 2: "b", 3: "c"}, true, []int{3}},
		{"d", map[int]string{1: "a", 2: "b", 3: "c"}, false, []int{}},
		{"a", map[int]string{1: "a", 2: "a", 3: "c"}, true, []int{1, 2}},
		{"a", map[int]string{}, false, []int{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapIntString(test.x, test.Map)
		if ok != test.expectedBool || !AreEqualSortedSlicesInt(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapIntString(%v, %v) should be %v with keys %v", test.x, test.Map, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapIntString() {
	keys, ok := IsValueInMapIntString("a", map[int]string{1: "a", 2: "b", 3: "a"})
	fmt.Println(ok, len(keys))
	keys, ok = IsValueInMapIntString("a", map[int]string{1: "c", 2: "b", 3: "d"})
	fmt.Println(ok, len(keys))
	// Output:
	// true 2
	// false 0
}

func TestIsValueInMapIntInt(t *testing.T) {
	tests := []struct {
		x            int
		Map          map[int]int
		expectedBool bool
		expectedKeys []int
	}{
		{1, map[int]int{1: 1, 2: 1, 3: 2}, true, []int{1, 2}},
		{2, map[int]int{1: 1, 2: 1, 3: 2}, true, []int{3}},
		{3, map[int]int{1: 1, 2: 1, 3: 2}, false, []int{}},
		{1, map[int]int{}, false, []int{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapIntInt(test.x, test.Map)
		if ok != test.expectedBool || !AreEqualSortedSlicesInt(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapIntInt(%v, %v) should be %v with keys %v", test.x, test.Map, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapIntInt() {
	fmt.Println(IsValueInMapIntInt(2, map[int]int{1: 1, 2: 1, 3: 2}))
	fmt.Println(IsValueInMapIntInt(3, map[int]int{1: 1, 2: 1, 3: 2}))
	// Output:
	// [3] true
	// [] false
}

func TestIsValueInMapIntFloat64(t *testing.T) {
	tests := []struct {
		x            float64
		Map          map[int]float64
		epsilon      float64
		expectedBool bool
		expectedKeys []int
	}{
		{1, map[int]float64{1: 1, 2: 1, 3: 2}, 0, true, []int{1, 2}},
		{2, map[int]float64{1: 1, 2: 1, 3: 2}, 0, true, []int{3}},
		{3, map[int]float64{1: 1, 2: 1, 3: 2}, 0, false, []int{}},
		{.01, map[int]float64{1: .01, 2: .02, 3: .02}, 0, true, []int{1}},
		{.01, map[int]float64{1: .011, 2: .02, 3: .02}, 0, false, []int{}},
		{.01, map[int]float64{1: .011, 2: .02, 3: .02}, 0.1, true, []int{1, 2, 3}},
		{1, map[int]float64{}, 0, false, []int{}},
	}
	for _, test := range tests {
		keys, ok := IsValueInMapIntFloat64(test.x, test.Map, test.epsilon)
		if ok != test.expectedBool || !AreEqualSortedSlicesInt(keys, test.expectedKeys) {
			t.Errorf("IsValueInMapIntFloat64(%v, %v, %v) should be %v with keys %v", test.x, test.Map, test.epsilon, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleIsValueInMapIntFloat64() {
	keys, ok := IsValueInMapIntFloat64(.01, map[int]float64{1: .01, 2: .02, 3: .02}, 0)
	fmt.Println(ok, len(keys))
	keys, ok = IsValueInMapIntFloat64(.01, map[int]float64{1: .011, 2: .02, 3: .02}, 0)
	fmt.Println(ok, len(keys))
	keys, ok = IsValueInMapIntFloat64(.01, map[int]float64{1: .011, 2: .02, 3: .02}, 0.1)
	fmt.Println(ok, len(keys))
	// Output:
	// true 1
	// false 0
	// true 3
}
