package check

import (
	"fmt"
	"testing"
)

func TestIsValueInStringSlice(t *testing.T) {
	tests := []struct {
		x        string
		xx       []string
		expected bool
	}{
		{"a", []string{"a", "b", "c"}, true},
		{"A", []string{"a", "b", "c"}, false},
		{"a ", []string{"a", "b", "c"}, false},
		{"a", []string{"a"}, true},
		{"a", []string{}, false},
	}
	for _, test := range tests {
		if IsValueInStringSlice(test.x, test.xx) != test.expected {
			t.Errorf("IsValueInStringSlice(%v, %v) should be %v", test.x, test.xx, test.expected)
		}
	}
}

func ExampleIsValueInStringSlice() {
	fmt.Println(IsValueInStringSlice("a", []string{"a", "b", "c"}))
	fmt.Println(IsValueInStringSlice("A", []string{"a", "b", "c"}))
	// Output:
	// true
	// false
}

func TestIsValueInIntSlice(t *testing.T) {
	tests := []struct {
		x        int
		xx       []int
		expected bool
	}{
		{1, []int{1, 2, 3}, true},
		{1, []int{2, 2, 3}, false},
		{1, []int{1, 1, 1}, true},
		{1, []int{}, false},
	}
	for _, test := range tests {
		if IsValueInIntSlice(test.x, test.xx) != test.expected {
			t.Errorf("TestIsValueInIntSlice(%v, %v) should be %v", test.x, test.xx, test.expected)
		}
	}
}

func ExampleIsValueInIntSlice() {
	fmt.Println(IsValueInIntSlice(1, []int{1, 2, 3}))
	fmt.Println(IsValueInIntSlice(1, []int{2, 2, 3}))
	// Output:
	// true
	// false
}

func TestIsValueInFloat64Slice(t *testing.T) {
	tests := []struct {
		x        float64
		xx       []float64
		epsilon  float64
		expected bool
	}{
		{1, []float64{1, 2, 3}, .000000001, true},
		{1, []float64{2, 2, 3}, .000000001, false},
		{1, []float64{1, 1, 1}, .000000001, true},
		{1, []float64{.001, .001, .001}, .000000001, false},
		{.001, []float64{.001, .001, .001}, .000000001, true},
		{.001, []float64{.001, .001, .001}, .01, true},
		{.001, []float64{}, .01, false},
		{.11, []float64{.1, .2, .2}, 0, false},
		{.11, []float64{.1, .2, .2}, .1, true},
		{.11, []float64{.1, .2, .2}, .01, true},
		{.11, []float64{.1, .2, .2}, .001, false},
	}
	for _, test := range tests {
		if IsValueInFloat64Slice(test.x, test.xx, test.epsilon) != test.expected {
			t.Errorf("IsValueInFloat64Slice(%v, %v) should be %v", test.x, test.xx, test.expected)
		}
	}
}

func ExampleIsValueInFloat64Slice() {
	fmt.Println(IsValueInFloat64Slice(.001, []float64{.001, .001, .001}, 0))
	fmt.Println(IsValueInFloat64Slice(.001, []float64{.002, .011, .002}, .0001))
	fmt.Println(IsValueInFloat64Slice(.001, []float64{.002, .011, .002}, .1))
	// Output:
	// true
	// false
	// true
}
