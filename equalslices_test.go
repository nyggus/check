package check

import (
	"fmt"
	"testing"
)

func TestAreEqualSlicesFloat64(t *testing.T) {
	tests := []struct {
		inputSlice1 []float64
		inputSlice2 []float64
		epsilon     float64
		expected    bool
	}{
		{[]float64{}, []float64{}, .0001, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55}, .0001, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .55}, .0001, true},
		{[]float64{.55, .45, .55}, []float64{.55, .55, .45}, .0001, false},
		{[]float64{2, 1}, []float64{1, 2}, .0001, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .55}, .1, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .54}, .1, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .54}, .01, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55001, .55}, .000001, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55001, .55}, .001, true},
		{[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			0,
			true,
		},
		{[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			[]float64{1212121212.55, -111111111111111.55, .000000000000000054},
			0,
			false,
		},
	}
	for _, test := range tests {
		actual := AreEqualSlicesFloat64(test.inputSlice1, test.inputSlice2, test.epsilon)
		if actual != test.expected {
			t.Errorf("AreEqualSlicesFloat64: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSlicesFloat64() {
	fmt.Println(AreEqualSlicesFloat64([]float64{.55, .55, .55}, []float64{.55, .55, .54}, 0))
	fmt.Println(AreEqualSlicesFloat64([]float64{.55, .55, .55}, []float64{.55, .55, .54}, .1))
	fmt.Println(AreEqualSlicesFloat64([]float64{.55, .55, .55}, []float64{.55, .55001, .55}, .000001))
	// Output:
	// false
	// true
	// false
}

func TestAreEqualSortedSlicesFloat64(t *testing.T) {
	tests := []struct {
		inputSlice1 []float64
		inputSlice2 []float64
		epsilon     float64
		expected    bool
	}{
		{[]float64{}, []float64{}, .0001, true},
		{[]float64{.55, .45, .55}, []float64{.55, .55, .45}, .0001, true},
		{[]float64{.55, .45, .55}, []float64{.55, .55}, .0001, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .55}, .0001, true},
		{[]float64{2, 1}, []float64{1, 2}, .0001, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .55}, .1, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .54}, .1, true},
		{[]float64{.55, .55, .55}, []float64{.55, .55, .54}, .01, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55001, .55}, .000001, false},
		{[]float64{.55, .55, .55}, []float64{.55, .55001, .55}, .001, true},
		{[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			0,
			true,
		},
		{[]float64{1212121212.55, -111111111111111.55, .000000000000000055},
			[]float64{1212121212.55, -111111111111111.55, .000000000000000054},
			0,
			false,
		},
	}
	for _, test := range tests {
		actual := AreEqualSortedSlicesFloat64(test.inputSlice1, test.inputSlice2, test.epsilon)
		if actual != test.expected {
			t.Errorf("AreEqualSortedSlicesFloat64: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSortedSlicesFloat64() {
	fmt.Println(AreEqualSortedSlicesFloat64([]float64{.55, .51, .55}, []float64{.55, .55, .51}, 0))
	fmt.Println(AreEqualSortedSlicesFloat64([]float64{.55, .51, .55}, []float64{.55, .55, .52}, .0001))
	// Output:
	// true
	// false
}

func benchmarkAreEqualSlicesFloat64Positive(value float64, n int, b *testing.B) {
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = value
		y[i] = value
	}
	for n := 0; n < b.N; n++ {
		AreEqualSlicesFloat64(x, y, 0)
	}
}

// EqualSlices variants

func BenchmarkAreEqualSlicesFloat64Positive10(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Positive(15436793.555555, 10, b)
}
func BenchmarkAreEqualSlicesFloat64Positive100(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Positive(15436793.555555, 100, b)
}
func BenchmarkAreEqualSlicesFloat64Positive1000(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Positive(15436793.555555, 1000, b)
}
func BenchmarkAreEqualSlicesFloat64Positive10000(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Positive(15436793.555555, 10000, b)
}

func benchmarkAreEqualSlicesFloat64Negative(value float64, n int, b *testing.B) {
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = value
		y[i] = value
	}
	y[len(y)-1] += .00001
	for n := 0; n < b.N; n++ {
		AreEqualSlicesFloat64(x, y, 0)
	}
}

func BenchmarkAreEqualSlicesFloat64Negative10(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Negative(15436793.555555, 10, b)
}
func BenchmarkAreEqualSlicesFloat64Negative100(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Negative(15436793.555555, 100, b)
}
func BenchmarkAreEqualSlicesFloat64Negative1000(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Negative(15436793.555555, 1000, b)
}
func BenchmarkAreEqualSlicesFloat64Negative10000(b *testing.B) {
	benchmarkAreEqualSlicesFloat64Negative(15436793.555555, 10000, b)
}

func TestAreEqualSlicesInt(t *testing.T) {
	tests := []struct {
		inputSlice1 []int
		inputSlice2 []int
		expected    bool
	}{
		{[]int{}, []int{}, true},
		{[]int{55, 55}, []int{55, 55, 55}, false},
		{[]int{2, 1}, []int{1, 2}, false},
		{[]int{55, 55, 55}, []int{55, 55, 55}, true},
		{[]int{55, 55, 55, 5000000000000}, []int{55, 55, 55, 5000000000000}, true},
		{[]int{55, 55, 45}, []int{55, 45, 55}, false},
		{[]int{55, 55, 45}, []int{55, 55, 45}, true},
		{[]int{55, 55, 45}, []int{55, 55, 55}, false},
	}
	for _, test := range tests {
		actual := AreEqualSlicesInt(test.inputSlice1, test.inputSlice2)
		if actual != test.expected {
			t.Errorf("AreEqualSlicesInt: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSlicesInt() {
	fmt.Println(AreEqualSlicesInt([]int{55, 55, 45}, []int{55, 55, 45}))
	fmt.Println(AreEqualSlicesInt([]int{55, 55, 45}, []int{55, 45, 55}))
	// Output:
	// true
	// false
}

func TestAreEqualSortedSlicesInt(t *testing.T) {
	tests := []struct {
		inputSlice1 []int
		inputSlice2 []int
		expected    bool
	}{
		{[]int{}, []int{}, true},
		{[]int{55, 55}, []int{55, 55, 55}, false},
		{[]int{2, 1}, []int{1, 2}, true},
		{[]int{55, 55, 55}, []int{55, 55, 55}, true},
		{[]int{55, 55, 55, 5000000000000}, []int{55, 55, 55, 5000000000000}, true},
		{[]int{55, 55, 45}, []int{55, 45, 55}, true},
		{[]int{55, 55, 45}, []int{55, 55, 45}, true},
		{[]int{55, 55, 45}, []int{55, 55, 55}, false},
	}
	for _, test := range tests {
		actual := AreEqualSortedSlicesInt(test.inputSlice1, test.inputSlice2)
		if actual != test.expected {
			t.Errorf("AreEqualSortedSlicesInt: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSortedSlicesInt() {
	fmt.Println(AreEqualSortedSlicesInt([]int{55, 55, 45}, []int{55, 45, 55}))
	fmt.Println(AreEqualSortedSlicesInt([]int{1, 2, 3}, []int{3, 2, 1}))
	// Output:
	// true
	// true
}

func TestAreEqualSlicesString(t *testing.T) {
	tests := []struct {
		inputSlice1 []string
		inputSlice2 []string
		expected    bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"55", "55"}, []string{"55", "55", "55"}, false},
		{[]string{"a", "b"}, []string{"b", "a"}, false},
		{[]string{"55", "55", "55"}, []string{"55", "55", "55"}, true},
		{[]string{"55", "55", "55", "5000000000000"}, []string{"55", "55", "55", "5000000000000"}, true},
		{[]string{"55", "55", "45"}, []string{"55", "45", "55"}, false},
		{[]string{"55", "55", "45"}, []string{"55", "55", "45"}, true},
	}
	for _, test := range tests {
		actual := AreEqualSlicesString(test.inputSlice1, test.inputSlice2)
		if actual != test.expected {
			t.Errorf("AreEqualSlicesString: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSlicesString() {
	fmt.Println(AreEqualSlicesString([]string{"a", "a", "b"}, []string{"a", "b"}))
	fmt.Println(AreEqualSlicesString([]string{"a", "a", "b"}, []string{"a", "a", "b"}))
	fmt.Println(AreEqualSlicesString([]string{"a", "a", "b"}, []string{"a", "b", "a"}))
	// Output:
	// false
	// true
	// false
}

func TestAreEqualSortedSlicesString(t *testing.T) {
	tests := []struct {
		inputSlice1 []string
		inputSlice2 []string
		expected    bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"55", "55"}, []string{"55", "55", "55"}, false},
		{[]string{"a", "b"}, []string{"b", "a"}, true},
		{[]string{"55", "55", "55"}, []string{"55", "55", "55"}, true},
		{[]string{"55", "55", "55", "5000000000000"}, []string{"55", "55", "55", "5000000000000"}, true},
		{[]string{"55", "55", "45"}, []string{"55", "45", "55"}, true},
		{[]string{"55", "55", "45"}, []string{"55", "55", "45"}, true},
		{[]string{"55", "55", "45"}, []string{"55", "55", "65"}, false},
	}
	for _, test := range tests {
		actual := AreEqualSortedSlicesString(test.inputSlice1, test.inputSlice2)
		if actual != test.expected {
			t.Errorf("AreEqualSortedSlicesString: %v == %v should be %v", test.inputSlice1, test.inputSlice2, test.expected)
		}
	}
}

func ExampleAreEqualSortedSlicesString() {
	fmt.Println(AreEqualSortedSlicesString([]string{"a", "a", "b"}, []string{"a", "a", "b"}))
	fmt.Println(AreEqualSortedSlicesString([]string{"a", "a", "b"}, []string{"a", "b", "a"}))
	// Output:
	// true
	// true
}
