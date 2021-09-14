package check

import (
	"fmt"
	"testing"
)

func TestAreEqualMapsStringFloat64(t *testing.T) {
	tests := []struct {
		map1     map[string]float64
		map2     map[string]float64
		epsilon  float64
		expected bool
	}{
		{map[string]float64{"1": 1.01}, map[string]float64{}, .00000001, false},
		{map[string]float64{}, map[string]float64{"1": 1.01}, .00000001, false},
		{map[string]float64{"1": 1.01}, map[string]float64{"1": 1.01}, .00000001, true},
		{map[string]float64{"1": 1.01}, map[string]float64{"1": 1.011}, .00000001, false},
		{map[string]float64{"1": 1.01}, map[string]float64{"1": 1.011}, .01, true},
		{map[string]float64{"1": 1.01, "2": 100.001, "3": 1.}, map[string]float64{"1": 1.01, "3": 1., "2": 100.001}, .00000001, true},
		{map[string]float64{"1": 1.01, "2": 100.001, "3": 1.}, map[string]float64{"1": 1.01, "3": 1., "2": 100.2}, .1, false},
	}
	for _, test := range tests {
		actual := AreEqualMapsStringFloat64(test.map1, test.map2, test.epsilon)
		if actual != test.expected {
			t.Errorf("AreEqualMapsStringFloat64(%v, %v, %v) = %v; want %v", test.map1, test.map2, test.epsilon, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsStringFloat64() {
	fmt.Println(AreEqualMapsStringFloat64(
		map[string]float64{"1": 1.01, "2": 100.001, "3": 1.},
		map[string]float64{"1": 1.1, "3": 1., "2": 100.001},
		.00000001,
	))
	fmt.Println(AreEqualMapsStringFloat64(
		map[string]float64{"1": 1.01, "2": 100.001, "3": 1.},
		map[string]float64{"1": 1.1, "3": 1., "2": 100.001},
		.1,
	))
	fmt.Println(AreEqualMapsStringFloat64(
		map[string]float64{"1": 1.01, "2": 100.001, "3": 1.},
		map[string]float64{"1": 1.01, "3": 1., "2": 100.001},
		0,
	))
	// Output:
	// false
	// true
	// true
}

func TestAreEqualMapsIntFloat64(t *testing.T) {
	tests := []struct {
		map1     map[int]float64
		map2     map[int]float64
		epsilon  float64
		expected bool
	}{
		{map[int]float64{1: 1.01}, map[int]float64{1: 1.01}, .00000001, true},
		{map[int]float64{1: 1.01}, map[int]float64{}, .00000001, false},
		{map[int]float64{}, map[int]float64{1: 1.01}, .00000001, false},
		{map[int]float64{1: 1.01}, map[int]float64{1: 1.011}, .00000001, false},
		{map[int]float64{1: 1.01}, map[int]float64{1: 1.011}, .01, true},
		{map[int]float64{1: 1.01, 2: 100.001, 3: 1.}, map[int]float64{1: 1.01, 3: 1., 2: 100.001}, .00000001, true},
		{map[int]float64{1: 1.01, 2: 100.001, 3: 1.}, map[int]float64{1: 1.01, 3: 1., 2: 100.2}, .1, false},
	}
	for _, test := range tests {
		actual := AreEqualMapsIntFloat64(test.map1, test.map2, test.epsilon)
		if actual != test.expected {
			t.Errorf("AreEqualMapsIntFloat64(%v, %v, %v) = %v; want %v", test.map1, test.map2, test.epsilon, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsIntFloat64() {
	fmt.Println(AreEqualMapsIntFloat64(
		map[int]float64{1: 1.01, 2: 100.001, 3: 1.},
		map[int]float64{1: 1.1, 3: 1., 2: 100.001},
		.00000001,
	))
	fmt.Println(AreEqualMapsIntFloat64(
		map[int]float64{1: 1.01, 2: 100.001, 3: 1.},
		map[int]float64{1: 1.1, 3: 1., 2: 100.001},
		.1,
	))
	fmt.Println(AreEqualMapsIntFloat64(
		map[int]float64{1: 1.01, 2: 100.001, 3: 1.},
		map[int]float64{1: 1.01, 3: 1., 2: 100.001},
		0,
	))
	// Output:
	// false
	// true
	// true
}

func TestAreEqualMapsStringInt(t *testing.T) {
	tests := []struct {
		map1     map[string]int
		map2     map[string]int
		expected bool
	}{
		{map[string]int{"1": 1}, map[string]int{"1": 1}, true},
		{map[string]int{"1": 1}, map[string]int{}, false},
		{map[string]int{}, map[string]int{"1": 1}, false},
		{map[string]int{"1": 2}, map[string]int{"1": 1}, false},
		{map[string]int{"1": 1}, map[string]int{"2": 1}, false},
		{map[string]int{"1": 1, "2": 100, "3": 1}, map[string]int{"1": 1, "3": 1, "2": 100}, true},
		{map[string]int{"1": 1, "2": 100, "3": 1}, map[string]int{"1": 1, "3": 1, "2": 101}, false},
	}
	for _, test := range tests {
		actual := AreEqualMapsStringInt(test.map1, test.map2)
		if actual != test.expected {
			t.Errorf("AreEqualMapsStringInt(%v, %v) = %v; want %v", test.map1, test.map2, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsStringInt() {
	fmt.Println(AreEqualMapsStringInt(
		map[string]int{"1": 1, "2": 100, "3": 1},
		map[string]int{"1": 1, "3": 1, "2": 100},
	))
	fmt.Println(AreEqualMapsStringInt(
		map[string]int{"1": 1, "2": 100, "3": 1},
		map[string]int{"1": 1, "3": 1, "2": 101},
	))
	// Output:
	// true
	// false
}

func TestAreEqualMapsStringString(t *testing.T) {
	tests := []struct {
		map1     map[string]string
		map2     map[string]string
		expected bool
	}{
		{map[string]string{"1": "1"}, map[string]string{"1": "1"}, true},
		{map[string]string{"1": "1"}, map[string]string{}, false},
		{map[string]string{}, map[string]string{"1": "1"}, false},
		{map[string]string{"1": "2"}, map[string]string{"1": "1"}, false},
		{map[string]string{"1": "1"}, map[string]string{"2": "1"}, false},
		{map[string]string{"1": "1", "2": "100", "3": "1"}, map[string]string{"1": "1", "3": "1", "2": "100"}, true},
		{map[string]string{"1": "1", "2": "100", "3": "|"}, map[string]string{"1": "1", "3": "1", "2": "101"}, false},
	}
	for _, test := range tests {
		actual := AreEqualMapsStringString(test.map1, test.map2)
		if actual != test.expected {
			t.Errorf("AreEqualMapsStringString(%v, %v) = %v; want %v", test.map1, test.map2, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsStringString() {
	fmt.Println(AreEqualMapsStringString(
		map[string]string{"1": "1", "2": "100", "3": "1"},
		map[string]string{"1": "1", "3": "1", "2": "100"},
	))
	fmt.Println(AreEqualMapsStringString(
		map[string]string{"1": "1", "2": "100", "3": "|"},
		map[string]string{"1": "1", "3": "1", "2": "101"},
	))
	// Output:
	// true
	// false
}

func TestAreEqualMapsIntInt(t *testing.T) {
	tests := []struct {
		map1     map[int]int
		map2     map[int]int
		expected bool
	}{
		{map[int]int{1: 1}, map[int]int{}, false},
		{map[int]int{}, map[int]int{1: 1}, false},
		{map[int]int{1: 1}, map[int]int{1: 1}, true},
		{map[int]int{1: 2}, map[int]int{1: 1}, false},
		{map[int]int{1: 1}, map[int]int{2: 1}, false},
		{map[int]int{1: 1, 2: 100, 3: 1}, map[int]int{1: 1, 3: 1, 2: 100}, true},
		{map[int]int{1: 1, 2: 100, 3: 1}, map[int]int{1: 1, 3: 1, 2: 101}, false},
	}
	for _, test := range tests {
		actual := AreEqualMapsIntInt(test.map1, test.map2)
		if actual != test.expected {
			t.Errorf("AreEqualMapsIntInt(%v, %v) = %v; want %v", test.map1, test.map2, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsIntInt() {
	fmt.Println(AreEqualMapsIntInt(
		map[int]int{1: 1, 2: 100, 3: 1},
		map[int]int{1: 1, 3: 1, 2: 100},
	))
	fmt.Println(AreEqualMapsIntInt(
		map[int]int{1: 1, 2: 100, 3: 1},
		map[int]int{1: 1, 3: 1, 2: 101},
	))
	// Output:
	// true
	// false
}

func TestAreEqualMapsIntString(t *testing.T) {
	tests := []struct {
		map1     map[int]string
		map2     map[int]string
		expected bool
	}{
		{map[int]string{1: "1"}, map[int]string{}, false},
		{map[int]string{}, map[int]string{1: "1"}, false},
		{map[int]string{1: "1"}, map[int]string{1: "1"}, true},
		{map[int]string{1: "2"}, map[int]string{1: "1"}, false},
		{map[int]string{1: "1"}, map[int]string{2: "1"}, false},
		{map[int]string{1: "1", 2: "100", 3: "|"}, map[int]string{1: "1", 3: "1", 2: "101"}, false},
		{map[int]string{1: "1", 2: "101", 3: "1"}, map[int]string{1: "1", 3: "1", 2: "101"}, true},
	}
	for _, test := range tests {
		actual := AreEqualMapsIntString(test.map1, test.map2)
		if actual != test.expected {
			t.Errorf("AreEqualMapsIntString(%v, %v) = %v; want %v", test.map1, test.map2, actual, test.expected)
		}
	}
}

func ExampleAreEqualMapsIntString() {
	fmt.Println(AreEqualMapsIntString(
		map[int]string{1: "1", 2: "101", 3: "1"},
		map[int]string{1: "1", 3: "1", 2: "101"},
	))
	fmt.Println(AreEqualMapsIntString(
		map[int]string{1: "1", 2: "100", 3: "|"},
		map[int]string{1: "1", 3: "1", 2: "101"},
	))
	// Output:
	// true
	// false
}
