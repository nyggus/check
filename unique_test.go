package check

import (
	"fmt"
	"testing"
)

func TestIsUniqueFloat64Slice(t *testing.T) {
	tests := []struct {
		input    []float64
		epsilon  float64
		expected bool
	}{
		{[]float64{}, .0000001, true},
		{[]float64{.55, .55, .55}, .0000001, false},
		{[]float64{.55, .55, .551}, .0000001, false},
		{[]float64{.55, .551, .552}, .0000001, true},
		{[]float64{.55, .551, .552}, .01, false},
		{[]float64{1, 1}, .0000001, false},
		{[]float64{1, 2}, .0000001, true},
		{[]float64{.551, .552, .553, .554, .555, .556, .557}, 0, true},
		{[]float64{.551, .552, .553, .554, .555, .556, .557}, .1, false},
		{[]float64{.551, .552, .553, .554, .555, .556, .557}, .01, false},
		{[]float64{.551, .552, .553, .554, .555, .556, .557}, .0011, false},
	}
	for _, test := range tests {
		actual := IsUniqueFloat64Slice(test.input, test.epsilon)
		if actual != test.expected {
			t.Errorf("IsUniqueFloat64Slice(%v, %v) = %v; want %v", test.input, test.epsilon, actual, test.expected)
		}
	}
}

func ExampleIsUniqueFloat64Slice() {
	fmt.Println(IsUniqueFloat64Slice([]float64{.551, .552, .553}, .0000001))
	// Watch for floats! The below example shows this:
	fmt.Println(IsUniqueFloat64Slice([]float64{.551, .552, .553}, .001))
	// But
	fmt.Println(IsUniqueFloat64Slice([]float64{.551, .552, .553}, .0011))
	// Output:
	// true
	// true
	// false
}

func TestIsUniqueIntSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{55, 55, 55}, false},
		{[]int{55, 5, 551}, true},
		{[]int{551, 551, 551}, false},
		{[]int{1, 1}, false},
		{[]int{1, 2}, true},
		{[]int{3155, 3255, 3355, 3455, 3555, 3565, 3754, 155, 255, 355, 455, 555, 565, 754, 3155}, false},
		{[]int{3155, 3255, 3355, 3455, 3555, 3565, 3754, 155, 255, 355, 455, 555, 565, 754}, true},
	}
	for _, test := range tests {
		actual := IsUniqueIntSlice(test.input)
		if actual != test.expected {
			t.Errorf("IsUniqueIntSlice(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func ExampleIsUniqueIntSlice() {
	fmt.Println(IsUniqueIntSlice([]int{55, 55, 55}))
	fmt.Println(IsUniqueIntSlice([]int{55, 56, 57}))
	// Output:
	// false
	// true
}
func TestIsUniqueStringSlice(t *testing.T) {
	tests := []struct {
		slice  []string
		unique bool
	}{
		{[]string{}, true},
		{[]string{"1"}, true},
		{[]string{"1", "1", "1"}, false},
		{[]string{"a", "a", "a"}, false},
		{[]string{"1", "1", "2"}, false},
		{[]string{"a", " a", "a "}, true},
		{[]string{"a", "", "a"}, false},
		{[]string{" ", "a", "A"}, true},
		{[]string{" ", "a", "A", "A"}, false},
	}
	for _, test := range tests {
		if IsUniqueStringSlice(test.slice) != test.unique {
			t.Errorf("IsUniqueStringSlice(%v) should be %v", test.slice, test.unique)
		}
	}
}

func ExampleIsUniqueStringSlice() {
	fmt.Println(IsUniqueStringSlice([]string{"a", "a", "a"}))
	fmt.Println(IsUniqueStringSlice([]string{"a", "b", "c"}))
	// Output:
	// false
	// true
}

func TestUniqueStringSlice(t *testing.T) {
	tests := []struct {
		slice  []string
		unique []string
	}{
		{[]string{"1", "1", "1"}, []string{"1"}},
		{[]string{"a", "a", "a"}, []string{"a"}},
		{[]string{"1"}, []string{"1"}},
		{[]string{}, []string{}},
		{[]string{"1", "1", "2"}, []string{"1", "2"}},
		{[]string{"a", "a", "a "}, []string{"a", "a "}},
		{[]string{"a", "", "a"}, []string{"a", ""}},
		{[]string{"a", "a", "A"}, []string{"a", "A"}},
	}
	for _, test := range tests {
		got := UniqueStringSlice(test.slice)
		expected := test.unique
		if !AreEqualSlicesString(got, expected) {
			t.Errorf("UniqueStringSlice(%v) should be %v, not %v", test.slice, expected, got)
		}
	}
}

func ExampleUniqueStringSlice() {
	fmt.Println(UniqueStringSlice([]string{"a", "a", "a"}))
	fmt.Println(UniqueStringSlice([]string{"a", "b", "a"}))
	// Output:
	// [a]
	// [a b]
}

func TestUniqueIntSlice(t *testing.T) {
	tests := []struct {
		slice  []int
		unique []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 12, 1, 1, 12, 1}, []int{1, 12}},
	}
	for _, test := range tests {
		got := UniqueIntSlice(test.slice)
		expected := test.unique
		if !AreEqualSlicesInt(got, expected) {
			t.Errorf("UniqueIntSlice(%v) should be %v, not %v", test.slice, expected, got)
		}
	}
}

func ExampleUniqueIntSlice() {
	fmt.Println(UniqueIntSlice([]int{1, 1, 2}))
	fmt.Println(UniqueIntSlice([]int{1, 1, 12, 1, 1, 12, 1}))
	// Output:
	// [1 2]
	// [1 12]
}

func TestUniqueFloat64Slice(t *testing.T) {
	tests := []struct {
		slice   []float64
		epsilon float64
		unique  []float64
	}{
		{[]float64{}, .0000001, []float64{}},
		{[]float64{1, 1, 1}, .0000001, []float64{1}},
		{[]float64{1, 2, 3}, .0000001, []float64{1, 2, 3}},
		{[]float64{1}, .0000001, []float64{1}},
		{[]float64{1, 1, 2}, .0000001, []float64{1, 2}},
		{[]float64{1, 1, 12, 1, 1, 12, 1}, .0000001, []float64{1, 12}},
		{[]float64{1., 1.01, 2}, .0000001, []float64{1, 1.01, 2}},
		{[]float64{1., 1.01, 2}, .1, []float64{1, 2}},
		{[]float64{.1, .2, .2}, .01, []float64{.1, .2}},
	}
	for _, test := range tests {
		got := UniqueFloat64Slice(test.slice, test.epsilon)
		expected := test.unique
		if !AreEqualSlicesFloat64(got, expected, test.epsilon) {
			t.Errorf("UniqueFloat64Slice(%v, %v) should be %v, not %v", test.slice, test.epsilon, expected, got)
		}
	}
}

func ExampleUniqueFloat64Slice() {
	fmt.Println(UniqueFloat64Slice([]float64{1., 1.01, 2}, 0))
	fmt.Println(UniqueFloat64Slice([]float64{1., 1.01, 2}, .1))
	// Output:
	// [1 1.01 2]
	// [1 2]
}

func benchmarkIsUniqueFloat64SlicePositive(value float64, n int, b *testing.B) {
	xx := make([]float64, n)
	for i := 0; i < n; i++ {
		xx[i] = value
	}
	for n := 0; n < b.N; n++ {
		IsUniqueFloat64Slice(xx, 0)
	}
}

func BenchmarkIsUniqueFloat64SlicePositive10(b *testing.B) {
	benchmarkIsUniqueFloat64SlicePositive(15436793.555555, 10, b)
}
func BenchmarkIsUniqueFloat64SlicePositive100(b *testing.B) {
	benchmarkIsUniqueFloat64SlicePositive(15436793.555555, 100, b)
}
func BenchmarkIsUniqueFloat64SlicePositive1000(b *testing.B) {
	benchmarkIsUniqueFloat64SlicePositive(15436793.555555, 1000, b)
}
func BenchmarkIsUniqueFloat64SlicePositive10000(b *testing.B) {
	benchmarkIsUniqueFloat64SlicePositive(15436793.555555, 10000, b)
}

func benchmarkIsUniqueFloat64SliceNegative(value float64, n int, b *testing.B) {
	xx := make([]float64, n)
	for i := 0; i < n; i++ {
		xx[i] = value
	}
	xx[len(xx)-1] *= 2
	for n := 0; n < b.N; n++ {
		IsUniqueFloat64Slice(xx, 0)
	}
}

func BenchmarkIsUniqueFloat64SliceNegative10(b *testing.B) {
	benchmarkIsUniqueFloat64SliceNegative(15436793.555555, 10, b)
}
func BenchmarkIsUniqueFloat64SliceNegative100(b *testing.B) {
	benchmarkIsUniqueFloat64SliceNegative(15436793.555555, 100, b)
}
func BenchmarkIsUniqueFloat64SliceNegative1000(b *testing.B) {
	benchmarkIsUniqueFloat64SliceNegative(15436793.555555, 1000, b)
}
func BenchmarkIsUniqueFloat64SliceNegative10000(b *testing.B) {
	benchmarkIsUniqueFloat64SliceNegative(15436793.555555, 10000, b)
}

func TestIsUniqueMapIntFloat64(t *testing.T) {
	tests := []struct {
		input    map[int]float64
		epsilon  float64
		expected bool
	}{
		{map[int]float64{}, 0, true},
		{map[int]float64{1: .1}, 0, true},
		{map[int]float64{1: .1}, .1, true},
		{map[int]float64{1: .1, 2: .1}, 0, false},
		{map[int]float64{1: .1, 2: .2}, 0, true},
		{map[int]float64{1: .1, 2: .11}, .001, true},
		{map[int]float64{1: .1, 2: .11}, .005, true},
		{map[int]float64{1: .1, 2: .11}, .01, false},
		{map[int]float64{1: .1, 2: .11, 3: .101, 100: .1105}, 0, true},
		{map[int]float64{1: .1, 2: .11, 3: .101, 100: .11}, .1, false},
	}
	for _, test := range tests {
		actual := IsUniqueMapIntFloat64(test.input, test.epsilon)
		if actual != test.expected {
			t.Errorf("IsUniqueMapIntFloat64(%v, %v) = %v; want %v", test.input, test.epsilon, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapIntFloat64() {
	fmt.Println(IsUniqueMapIntFloat64(map[int]float64{1: .1, 2: .11, 3: .101, 100: .1105}, 0))
	fmt.Println(IsUniqueMapIntFloat64(map[int]float64{1: .1, 2: .11, 3: .101, 100: .1105}, .1))
	// Output:
	// true
	// false
}

func TestIsUniqueMapStringFloat64(t *testing.T) {
	tests := []struct {
		input    map[string]float64
		epsilon  float64
		expected bool
	}{
		{map[string]float64{}, 0, true},
		{map[string]float64{"a": .1}, 0, true},
		{map[string]float64{"a": .1}, .1, true},
		{map[string]float64{"a": .1, "b": .1}, 0, false},
		{map[string]float64{"a": .1, "b": .2}, 0, true},
		{map[string]float64{"a": .1, "b": .11}, .001, true},
		{map[string]float64{"a": .1, "b": .11}, .005, true},
		{map[string]float64{"a": .1, "b": .11}, .01, false},
		{map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, 0, true},
		{map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, .1, false},
	}
	for _, test := range tests {
		actual := IsUniqueMapStringFloat64(test.input, test.epsilon)
		if actual != test.expected {
			t.Errorf("IsUniqueMapIntFloat64(%v, %v) = %v; want %v", test.input, test.epsilon, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapStringFloat64() {
	fmt.Println(IsUniqueMapStringFloat64(map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, 0))
	fmt.Println(IsUniqueMapStringFloat64(map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, .1))
	// Output:
	// true
	// false
}

func TestIsUniqueMapIntInt(t *testing.T) {
	tests := []struct {
		input    map[int]int
		expected bool
	}{
		{map[int]int{}, true},
		{map[int]int{1: 1}, true},
		{map[int]int{1: 1, 2: 1}, false},
		{map[int]int{1: 1, 2: 2}, true},
		{map[int]int{1: 100, 2: 101}, true},
		{map[int]int{1: 1, 2: 1, 3: 1, 100: 1}, false},
		{map[int]int{1: 1, 2: 2, 3: 3, 100: 4}, true},
	}
	for _, test := range tests {
		actual := IsUniqueMapIntInt(test.input)
		if actual != test.expected {
			t.Errorf("IsUniqueMapIntInt(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapIntInt() {
	fmt.Println(IsUniqueMapIntInt(map[int]int{1: 1, 2: 2, 3: 3, 100: 4}))
	fmt.Println(IsUniqueMapIntInt(map[int]int{1: 1, 2: 1, 3: 1, 100: 1}))
	// Output:
	// true
	// false
}

func TestIsUniqueMapIntString(t *testing.T) {
	tests := []struct {
		input    map[int]string
		expected bool
	}{
		{map[int]string{}, true},
		{map[int]string{1: "a"}, true},
		{map[int]string{1: "a", 2: "a"}, false},
		{map[int]string{1: "a", 2: "b"}, true},
		{map[int]string{1: "Zigi says hi", 2: "Zigi says hi!"}, true},
		{map[int]string{1: "a", 2: "a", 3: "a", 100: "a"}, false},
		{map[int]string{1: "a", 2: "b", 3: "B", 100: "A"}, true},
	}
	for _, test := range tests {
		actual := IsUniqueMapIntString(test.input)
		if actual != test.expected {
			t.Errorf("IsUniqueMapIntString(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapIntString() {
	fmt.Println(IsUniqueMapIntString(map[int]string{1: "a", 2: "a", 3: "a", 100: "a"}))
	fmt.Println(IsUniqueMapIntString(map[int]string{1: "a", 2: "b", 3: "B", 100: "A"}))
	// Output:
	// false
	// true
}

func TestIsUniqueMapStringInt(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected bool
	}{
		{map[string]int{}, true},
		{map[string]int{"a": 1}, true},
		{map[string]int{"a": 1, "b": 1}, false},
		{map[string]int{"a": 1, "b": 2}, true},
		{map[string]int{"a": 100, "b": 101}, true},
		{map[string]int{"a": 1, "b": 1, "c": 1, "Blues is my soul, blues is my heart": 1}, false},
		{map[string]int{"a": 1, "b": 21, "c": 11, "Blues is my soul, blues is my heart": 2}, true},
	}
	for _, test := range tests {
		actual := IsUniqueMapStringInt(test.input)
		if actual != test.expected {
			t.Errorf("IsUniqueMapStringInt(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapStringInt() {
	fmt.Println(IsUniqueMapStringInt(
		map[string]int{"a": 1, "b": 21, "c": 11, "Blues is my soul, blues is my heart": 2},
	))
	fmt.Println(IsUniqueMapStringInt(
		map[string]int{"a": 1, "b": 1, "c": 1, "Blues is my soul, blues is my heart": 1},
	))
	// Output:
	// true
	// false
}
func TestIsUniqueMapStringString(t *testing.T) {
	tests := []struct {
		input    map[string]string
		expected bool
	}{
		{map[string]string{}, true},
		{map[string]string{"a": "Zigi"}, true},
		{map[string]string{"a": "Zigi", "b": "Zigi"}, false},
		{map[string]string{"a": "Zigi", "b": "zigi"}, true},
		{map[string]string{"a": "Zigi", "b": "oh", "c": "Zigi", "Blues is my soul, blues is my heart": "you are a liar!"}, false},
		{map[string]string{"a": "Zigi", "b": "why", "c": "oh why", "Blues is my soul, blues is my heart": "you are a liar!"}, true},
	}
	for _, test := range tests {
		actual := IsUniqueMapStringString(test.input)
		if actual != test.expected {
			t.Errorf("IsUniqueMapStringInt(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func ExampleIsUniqueMapStringString() {
	fmt.Println(IsUniqueMapStringString(
		map[string]string{"key1": "a", "key2": "a", "key3": "Q"},
	))
	fmt.Println(IsUniqueMapStringString(
		map[string]string{"key1": "a", "key2": "Q", "key3": "A"},
	))
	// Output:
	// false
	// true
}

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{[]float64{}, true},
		{[]float64{.55, .55, .55}, false},
		{[]float64{.55, .55, .551}, false},
		{[]float64{.55, .551, .552}, true},
		{[]float64{1, 1}, false},
		{[]float64{1, 2}, true},
		{[]float64{.551, .552, .553, .554, .555, .556, .557}, true},
		{[]string{}, true},
		{[]string{"1"}, true},
		{[]string{"1", "1", "1"}, false},
		{[]string{"a", "a", "a"}, false},
		{[]string{"1", "1", "2"}, false},
		{[]string{"a", " a", "a "}, true},
		{[]string{"a", "", "a"}, false},
		{[]string{" ", "a", "A"}, true},
		{[]string{" ", "a", "A", "A"}, false},
		{[]int{}, true},
		{[]int{55, 55, 55}, false},
		{[]int{55, 5, 551}, true},
		{[]int{551, 551, 551}, false},
		{[]int{1, 1}, false},
		{[]int{1, 2}, true},
		{[]int{3155, 3255, 3355, 3455, 3555, 3565, 3754, 155, 255, 355, 455, 555, 565, 754, 3155}, false},
		{[]int{3155, 3255, 3355, 3455, 3555, 3565, 3754, 155, 255, 355, 455, 555, 565, 754}, true},
	}
	for _, test := range tests {
		actual := IsUnique(test.input)
		if actual != test.expected {
			t.Errorf("IsUnique(%v) = %v; want %v", test.input, actual, test.expected)
		}
	}
}

func TestIsUniquePanics1(t *testing.T) {
	defer func() { recover() }()
	incorrectSlice := []int32{1, 2}
	IsUnique(incorrectSlice)
	t.Errorf("IsUnique did not panic for %v", incorrectSlice)
}

func TestIsUniquePanics2(t *testing.T) {
	defer func() { recover() }()
	incorrectSlice := []bool{true, false}
	IsUnique(incorrectSlice)
	t.Errorf("IsUnique did not panic for %v", incorrectSlice)
}

func TestIsUniquePanics3(t *testing.T) {
	defer func() { recover() }()
	incorrectSlice := 1000
	IsUnique(incorrectSlice)
	t.Errorf("IsUnique did not panic for %v", incorrectSlice)
}

func TestIsUniquePanics4(t *testing.T) {
	defer func() { recover() }()
	incorrectSlice := "MoD: Master of Distaster"
	IsUnique(incorrectSlice)
	t.Errorf("IsUnique did not panic for %v", incorrectSlice)
}
