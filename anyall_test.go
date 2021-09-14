package check

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	tests := []struct {
		conditions []bool
		expected   bool
	}{
		{[]bool{true, true}, true},
		{[]bool{true, false}, true},
		{[]bool{false, false}, false},
		{[]bool{false, false, false}, false},
		{[]bool{true}, true},
		{[]bool{}, false},
	}
	for _, test := range tests {
		if Any(test.conditions) != test.expected {
			t.Errorf("Any(%v) should be %v", test.conditions, test.expected)
		}
	}
}

func ExampleAny() {
	fmt.Println(Any([]bool{true, false, false}))
	fmt.Println(Any([]bool{false, false, false}))
	fmt.Println(Any([]bool{true}))
	fmt.Println(Any([]bool{}))
	// Output:
	//
	// true
	// false
	// true
	// false
}

func TestWhichInMapInt(t *testing.T) {
	tests := []struct {
		conditions   map[int]bool
		expectedBool bool
		expectedKeys []int
	}{
		{map[int]bool{}, false, []int{}},
		{map[int]bool{1: true, 2: true}, true, []int{1, 2}},
		{map[int]bool{1: true, 2: false}, true, []int{1}},
		{map[int]bool{1: false, 2: false}, false, []int{}},
		{map[int]bool{1: false, 2: false, 20: false}, false, []int{}},
		{map[int]bool{1: true, 2: true, 20: true, 5: true, 7: true, 88: true, 256: true, 10000: true}, true, []int{1, 2, 20, 5, 7, 88, 256, 10000}},
		{map[int]bool{1: false, 2: false, 20: false, 5: false, 7: false, 88: false, 256: false, 10000: true}, true, []int{10000}},
	}
	for _, test := range tests {
		keys, ok := WhichInMapInt(test.conditions)
		if ok != test.expectedBool || !AreEqualSortedSlicesInt(keys, test.expectedKeys) {
			t.Errorf("WhichInMapInt(%v) should be %v with keys %v", test.conditions, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichInMapInt() {
	keys, ok := WhichInMapInt(map[int]bool{1: true, 2: false})
	fmt.Println(ok, keys)
	keys, ok = WhichInMapInt(map[int]bool{1: false, 2: false})
	fmt.Println(ok, keys)
	// Output:
	// true [1]
	// false []
}

func TestAnyInMapInt(t *testing.T) {
	tests := []struct {
		conditions   map[int]bool
		expectedBool bool
	}{
		{map[int]bool{}, false},
		{map[int]bool{1: true, 2: true}, true},
		{map[int]bool{1: true, 2: false}, true},
		{map[int]bool{1: false, 2: false}, false},
		{map[int]bool{1: false, 2: false, 20: false}, false},
		{map[int]bool{1: true, 2: true, 20: true, 5: true, 7: true, 88: true, 256: true, 10000: true}, true},
		{map[int]bool{1: false, 2: false, 20: false, 5: false, 7: false, 88: false, 256: false, 10000: true}, true},
	}
	for _, test := range tests {
		if AnyInMapInt(test.conditions) != test.expectedBool {
			t.Errorf("AnyInMapInt(%v) should be %v", test.conditions, test.expectedBool)
		}
	}
}

func ExampleAnyInMapInt() {
	fmt.Println(AnyInMapInt(map[int]bool{1: true, 2: false}))
	fmt.Println(AnyInMapInt(map[int]bool{1: false, 2: false}))
	// Output:
	// true
	// false
}

func TestWhichInMapString(t *testing.T) {
	tests := []struct {
		conditions   map[string]bool
		expectedBool bool
		expectedKeys []string
	}{
		{map[string]bool{}, false, []string{}},
		{map[string]bool{"1": true, "2": true}, true, []string{"1", "2"}},
		{map[string]bool{"1": true, "2": false}, true, []string{"1"}},
		{map[string]bool{"1": false, "2": false}, false, []string{}},
		{map[string]bool{"1": false, "2": false, "20": false}, false, []string{}},
		{map[string]bool{
			"1": true, "2": true, "20": true, "5": true, "7": true, "88": true, "256": true, "10000": true},
			true,
			[]string{"1", "2", "20", "5", "7", "88", "256", "10000"},
		},
		{map[string]bool{"1": false, "2": false, "20": false, "5": false, "7": false, "88": false, "256": false, "10000": true}, true, []string{"10000"}},
	}
	for _, test := range tests {
		keys, ok := WhichInMapString(test.conditions)
		if ok != test.expectedBool || !AreEqualSortedSlicesString(keys, test.expectedKeys) {
			t.Errorf("WhichInMapInt(%v) should be %v with keys %v", test.conditions, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichInMapString() {
	keys, ok := WhichInMapString(map[string]bool{"key1": true, "key2": false})
	fmt.Println(ok, keys)
	keys, ok = WhichInMapString(map[string]bool{"key1": false, "key2": false})
	fmt.Println(ok, keys)
	// Output:
	// true [key1]
	// false []
}

func TestAnyInMapString(t *testing.T) {
	tests := []struct {
		conditions   map[string]bool
		expectedBool bool
	}{
		{map[string]bool{}, false},
		{map[string]bool{"1": true, "2": true}, true},
		{map[string]bool{"1": true, "2": false}, true},
		{map[string]bool{"1": false, "2": false}, false},
		{map[string]bool{"1": false, "2": false, "20": false}, false},
		{map[string]bool{
			"1": true, "2": true, "20": true, "5": true, "7": true, "88": true, "256": true, "10000": true},
			true,
		},
		{map[string]bool{
			"1": false, "2": false, "20": false, "5": false, "7": false, "88": false, "256": false, "10000": true},
			true,
		},
	}
	for _, test := range tests {
		if AnyInMapString(test.conditions) != test.expectedBool {
			t.Errorf("AnyInMapString(%v) should be %v", test.conditions, test.expectedBool)
		}
	}
}

func ExampleAnyInMapString() {
	fmt.Println(AnyInMapString(map[string]bool{"key1": true, "key2": false}))
	fmt.Println(AnyInMapString(map[string]bool{"key1": false, "key2": false}))
	// Output:
	// true
	// false
}

func TestAll(t *testing.T) {
	tests := []struct {
		conditions []bool
		expected   bool
	}{
		{[]bool{true, true}, true},
		{[]bool{true, false}, false},
		{[]bool{false, false}, false},
		{[]bool{false, false, false}, false},
		{[]bool{false}, false},
		{[]bool{true}, true},
		{[]bool{}, false},
	}
	for _, test := range tests {
		if All(test.conditions) != test.expected {
			t.Errorf("All(%v) should be %v", test.conditions, test.expected)
		}
	}
}

func ExampleAll() {
	fmt.Println(All([]bool{true, false, false}))
	fmt.Println(All([]bool{true, true, true}))
	fmt.Println(All([]bool{false, false, false}))
	fmt.Println(All([]bool{}))
	// Output:
	//
	// false
	// true
	// false
	// false
}

func TestAllInMapInt(t *testing.T) {
	tests := []struct {
		conditions   map[int]bool
		expectedBool bool
	}{
		{map[int]bool{}, false},
		{map[int]bool{1: true, 2: true}, true},
		{map[int]bool{1: true, 2: false}, false},
		{map[int]bool{1: false, 2: false}, false},
		{map[int]bool{1: false, 2: false, 20: false}, false},
		{map[int]bool{1: true, 2: true, 20: true, 5: true, 7: true, 88: true, 256: true, 10000: true}, true},
		{map[int]bool{1: false, 2: false, 20: false, 5: false, 7: false, 88: false, 256: false, 10000: true}, false},
	}
	for _, test := range tests {
		is := AllInMapInt(test.conditions)
		if is != test.expectedBool {
			t.Errorf("WhichInMapInt(%v) should be %v", test.conditions, test.expectedBool)
		}
	}
}

func ExampleAllInMapInt() {
	fmt.Println(AllInMapInt(map[int]bool{1: true, 2: true}))
	fmt.Println(AllInMapInt(map[int]bool{1: true, 2: false}))
	fmt.Println(AllInMapInt(map[int]bool{1: false, 2: false}))
	// Output:
	// true
	// false
	// false
}

func TestAllInMapString(t *testing.T) {
	tests := []struct {
		conditions   map[string]bool
		expectedBool bool
	}{
		{map[string]bool{}, false},
		{map[string]bool{"1": true, "2": true}, true},
		{map[string]bool{"1": true, "2": false}, false},
		{map[string]bool{"1": false, "2": false}, false},
		{map[string]bool{"1": false, "2": false, "20": false}, false},
		{map[string]bool{"1": true, "2": true, "20": true, "5": true, "7": true, "88": true, "256": true, "10000": true}, true},
		{map[string]bool{"1": false, "2": false, "20": false, "5": false, "7": false, "88": false, "256": false, "10000": true}, false},
	}
	for _, test := range tests {
		is := AllInMapString(test.conditions)
		if is != test.expectedBool {
			t.Errorf("WhichInMapInt(%v) should be %v", test.conditions, test.expectedBool)
		}
	}
}

func ExampleAllInMapString() {
	fmt.Println(AllInMapString(map[string]bool{"key1": true, "key2": true}))
	fmt.Println(AllInMapString(map[string]bool{"key1": true, "key2": false}))
	fmt.Println(AllInMapString(map[string]bool{"key1": false, "key2": false}))
	// Output:
	// true
	// false
	// false
}
