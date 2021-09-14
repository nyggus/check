package check

import (
	"fmt"
	"testing"
)

func TestAllKeyValuePairsInMapStringString(t *testing.T) {
	tests := []struct {
		Map1     map[string]string
		Map2     map[string]string
		expected bool
	}{
		{map[string]string{}, map[string]string{"1": "a", "2": "b"}, false},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{}, false},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "2": "b"}, true},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"11": "a", "12": "b"}, false},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "B"}, false},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapStringString(test.Map1, test.Map2) != test.expected {
			t.Errorf("AllKeyValuePairsInMapStringString(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapStringString() {
	fmt.Println(AllKeyValuePairsInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "a", "3": "c", "2": "b"},
	))
	fmt.Println(AllKeyValuePairsInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "a", "3": "c", "2": "B"},
	))
	// Output:
	// true
	// false
}

func TestAnyKeyValuePairInMapStringString(t *testing.T) {
	tests := []struct {
		Map1     map[string]string
		Map2     map[string]string
		expected bool
	}{
		{map[string]string{}, map[string]string{"1": "a", "2": "b"}, false},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{}, false},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "2": "b"}, true},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true},
		{map[string]string{"1": "a", "2": "B"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true},
		{map[string]string{"1": "A", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "A", "3": "c", "2": "B"}, false},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapStringString(test.Map1, test.Map2) != test.expected {
			t.Errorf("AnyKeyValuePairInMapStringString(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapStringString() {
	fmt.Println(AnyKeyValuePairInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "a", "3": "c", "2": "X"},
	))
	fmt.Println(AnyKeyValuePairInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "B", "3": "c", "2": "B"},
	))
	// Output:
	// true
	// false
}

func TestWhichKeyValuePairsInMapStringString(t *testing.T) {
	tests := []struct {
		Map1         map[string]string
		Map2         map[string]string
		expectedBool bool
		expectedKeys map[string]string
	}{
		{map[string]string{}, map[string]string{"1": "a", "2": "b"}, false, map[string]string{}},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{}, false, map[string]string{}},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "2": "b"}, true, map[string]string{"1": "a", "2": "b"}},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true, map[string]string{"1": "a", "2": "b"}},
		{map[string]string{"1": "a", "2": "B"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true, map[string]string{"1": "a"}},
		{map[string]string{"1": "A", "2": "b"}, map[string]string{"1": "a", "3": "c", "2": "b"}, true, map[string]string{"2": "b"}},
		{map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "A", "3": "c", "2": "B"}, false, map[string]string{}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapStringString(test.Map1, test.Map2)
		if ok != test.expectedBool || !AreEqualMapsStringString(keys, test.expectedKeys) {
			t.Errorf("WhichKeyValuePairsInMapStringString(%v, %v) should be %v with keys %v", test.Map1, test.Map2, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichKeyValuePairsInMapStringString() {
	keys, ok := WhichKeyValuePairsInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "a", "3": "c", "2": "X"},
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapStringString(
		map[string]string{"1": "a", "2": "b"},
		map[string]string{"1": "B", "3": "c", "2": "B"},
	)
	fmt.Println(ok, keys)
	// Output:
	// true map[1:a]
	// false map[]
}

func TestAllKeyValuePairsInMapStringInt(t *testing.T) {
	tests := []struct {
		Map1     map[string]int
		Map2     map[string]int
		expected bool
	}{
		{map[string]int{}, map[string]int{"1": 1, "2": 2}, false},
		{map[string]int{"1": 1, "2": 1}, map[string]int{}, false},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "2": 2}, true},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"11": 1, "12": 2}, false},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"11": 1, "12": 2}, false},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "3": 3, "2": 2}, true},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "3": 3, "2": 40}, false},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapStringInt(test.Map1, test.Map2) != test.expected {
			t.Errorf("AllKeyValuePairsInMapStringInt(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapStringInt() {
	fmt.Println(AllKeyValuePairsInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 1, "2": 2},
	))
	fmt.Println(AllKeyValuePairsInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 1, "3": 3, "2": 40},
	))
	// Output:
	// true
	// false
}

func TestAnyKeyValuePairInMapStringInt(t *testing.T) {
	tests := []struct {
		Map1     map[string]int
		Map2     map[string]int
		expected bool
	}{
		{map[string]int{}, map[string]int{"1": 1, "2": 2}, false},
		{map[string]int{"1": 1, "2": 1}, map[string]int{}, false},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "2": 23}, true},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "3": 3, "2": 23}, true},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 21, "3": 23, "2": 40}, false},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapStringInt(test.Map1, test.Map2) != test.expected {
			t.Errorf("AnyKeyValuePairInMapStringInt(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapStringInt() {
	fmt.Println(AnyKeyValuePairInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 1, "2": 23},
	))
	fmt.Println(AnyKeyValuePairInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 21, "3": 3, "2": 40},
	))
	// Output:
	// true
	// false
}

func TestWhichKeyValuePairsInMapStringInt(t *testing.T) {
	tests := []struct {
		Map1         map[string]int
		Map2         map[string]int
		expectedBool bool
		expectedKeys map[string]int
	}{
		{map[string]int{}, map[string]int{"1": 1, "2": 2}, false, map[string]int{}},
		{map[string]int{"1": 1, "2": 1}, map[string]int{}, false, map[string]int{}},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "2": 23}, true, map[string]int{"1": 1}},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "3": 3, "2": 2}, true, map[string]int{"1": 1, "2": 2}},
		{map[string]int{"1": 1, "2": 2}, map[string]int{"1": 21, "3": 23, "2": 40}, false, map[string]int{}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapStringInt(test.Map1, test.Map2)
		if ok != test.expectedBool || !AreEqualMapsStringInt(keys, test.expectedKeys) {
			t.Errorf("WhichKeyValuePairsInMapStringInt(%v, %v) should be %v with keys %v", test.Map1, test.Map2, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichKeyValuePairsInMapStringInt() {

	fmt.Println(WhichKeyValuePairsInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 1, "2": 23},
	))
	fmt.Println(WhichKeyValuePairsInMapStringInt(
		map[string]int{"1": 1, "2": 2},
		map[string]int{"1": 21, "3": 3, "2": 40},
	))
	// Output:
	// map[1:1] true
	// map[] false
}

func TestAllKeyValuePairsInMapStringFloat64(t *testing.T) {
	tests := []struct {
		Map1     map[string]float64
		Map2     map[string]float64
		epsilon  float64
		expected bool
	}{
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, 0, false},
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, .1, false},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, 0, false},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, .1, false},
		{map[string]float64{"1": .1, "2": .2}, map[string]float64{"1": .1, "2": .2}, 0, true},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": .1, "2": .2}, 0, false},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"11": .1, "12": .2}, 0, false},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": .1, "2": .2}, .01, true},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapStringFloat64(test.Map1, test.Map2, test.epsilon) != test.expected {
			t.Errorf("AllKeyValuePairsInMapStringFloat64(%v, %v, %v) should be %v", test.Map1, test.Map2, test.epsilon, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapStringFloat64() {
	fmt.Println(AllKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .1, "2": .2},
		map[string]float64{"1": .1, "2": .2},
		0,
	))
	fmt.Println(AllKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .11, "2": .2},
		map[string]float64{"1": .1, "2": .2},
		0,
	))
	fmt.Println(AllKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .11, "2": .2},
		map[string]float64{"1": .1, "2": .2},
		.01,
	))
	// Output:
	// true
	// false
	// true
}

func TestAnyKeyValuePairInMapStringFloat64(t *testing.T) {
	tests := []struct {
		Map1     map[string]float64
		Map2     map[string]float64
		epsilon  float64
		expected bool
	}{
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, 0, false},
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, .1, false},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, 0, false},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, .1, false},
		{map[string]float64{"1": .1, "2": .2}, map[string]float64{"1": .1, "2": 2.2}, 0, true},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": 2.1, "2": 2.2}, 0, false},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": .1, "2": 2.2}, .01, true},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapStringFloat64(test.Map1, test.Map2, test.epsilon) != test.expected {
			t.Errorf("AnyKeyValuePairInMapStringFloat64(%v, %v, %v) should be %v", test.Map1, test.Map2, test.epsilon, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapStringFloat64() {
	fmt.Println(AnyKeyValuePairInMapStringFloat64(
		map[string]float64{"1": .1, "2": .2},
		map[string]float64{"1": .1, "2": 2.2},
		0,
	))
	fmt.Println(AnyKeyValuePairInMapStringFloat64(
		map[string]float64{"1": .11, "2": 2.2},
		map[string]float64{"1": .1, "2": .2},
		0,
	))
	fmt.Println(AnyKeyValuePairInMapStringFloat64(
		map[string]float64{"1": .11, "2": .2},
		map[string]float64{"1": .1, "2": 2.2},
		.01,
	))
	// Output:
	// true
	// false
	// true
}

func TestWhichKeyValuePairsInMapStringFloat64(t *testing.T) {
	tests := []struct {
		Map1         map[string]float64
		Map2         map[string]float64
		epsilon      float64
		expectedBool bool
		expectedKeys map[string]float64
	}{
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, 0, false, map[string]float64{}},
		{map[string]float64{}, map[string]float64{"1": .1, "2": 2.}, .1, false, map[string]float64{}},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, 0, false, map[string]float64{}},
		{map[string]float64{"1": .1, "2": .1}, map[string]float64{}, .1, false, map[string]float64{}},
		{map[string]float64{"1": .1, "2": .2}, map[string]float64{"1": .1, "2": 2.2}, 0, true, map[string]float64{"1": .1}},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": 2.1, "2": 2.2}, 0, false, map[string]float64{}},
		{map[string]float64{"1": .11, "2": .2}, map[string]float64{"1": .1, "2": 2.2}, .01, true, map[string]float64{"1": .11}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapStringFloat64(test.Map1, test.Map2, test.epsilon)
		if ok != test.expectedBool || !AreEqualMapsStringFloat64(keys, test.expectedKeys, 0) {
			t.Errorf(
				"WhichKeyValuePairsInMapStringFloat64(%v, %v, %v) should be %v with keys %v, not %v and %v",
				test.Map1, test.Map2, test.epsilon, test.expectedBool, test.expectedKeys, ok, keys,
			)
		}
	}
}

func ExampleWhichKeyValuePairsInMapStringFloat64() {
	keys, ok := WhichKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .1, "2": .2},
		map[string]float64{"1": .1, "2": 2.2},
		0,
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .11, "2": 2.2},
		map[string]float64{"1": .1, "2": .2},
		0,
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapStringFloat64(
		map[string]float64{"1": .11, "2": .2},
		map[string]float64{"1": .1, "2": 2.2},
		.01,
	)
	fmt.Println(ok, keys)
	// Output:
	// true map[1:0.1]
	// false map[]
	// true map[1:0.11]
}

func TestAllKeyValuePairsInMapIntString(t *testing.T) {
	tests := []struct {
		Map1     map[int]string
		Map2     map[int]string
		expected bool
	}{
		{map[int]string{}, map[int]string{1: "a", 2: "b"}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{11: "a", 12: "b"}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "b"}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "B"}, false},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapIntString(test.Map1, test.Map2) != test.expected {
			t.Errorf("AllKeyValuePairsInMapIntString(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapIntString() {
	fmt.Println(AllKeyValuePairsInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "a", 3: "c", 2: "b"},
	))
	fmt.Println(AllKeyValuePairsInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "a", 3: "c", 2: "B"},
	))
	// Output:
	// true
	// false
}

func TestAnyKeyValuePairInMapIntString(t *testing.T) {
	tests := []struct {
		Map1     map[int]string
		Map2     map[int]string
		expected bool
	}{
		{map[int]string{}, map[int]string{1: "a", 2: "b"}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "b"}, true},
		{map[int]string{1: "a", 2: "B"}, map[int]string{1: "a", 3: "c", 2: "b"}, true},
		{map[int]string{1: "A", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "b"}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "A", 3: "c", 2: "B"}, false},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapIntString(test.Map1, test.Map2) != test.expected {
			t.Errorf("AnyKeyValuePairInMapIntString(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapIntString() {
	fmt.Println(AnyKeyValuePairInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "a", 3: "c", 2: "X"},
	))
	fmt.Println(AnyKeyValuePairInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "B", 3: "c", 2: "B"},
	))
	// Output:
	// true
	// false
}

func TestWhichKeyValuePairsInMapIntString(t *testing.T) {
	tests := []struct {
		Map1         map[int]string
		Map2         map[int]string
		expectedBool bool
		expectedKeys map[int]string
	}{
		{map[int]string{}, map[int]string{1: "a", 2: "b"}, false, map[int]string{}},
		{map[int]string{1: "a", 2: "b"}, map[int]string{}, false, map[int]string{}},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}, true, map[int]string{1: "a", 2: "b"}},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "b"}, true, map[int]string{1: "a", 2: "b"}},
		{map[int]string{1: "a", 2: "B"}, map[int]string{1: "a", 3: "c", 2: "b"}, true, map[int]string{1: "a"}},
		{map[int]string{1: "A", 2: "b"}, map[int]string{1: "a", 3: "c", 2: "b"}, true, map[int]string{2: "b"}},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "A", 3: "c", 2: "B"}, false, map[int]string{}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapIntString(test.Map1, test.Map2)
		if ok != test.expectedBool || !AreEqualMapsIntString(keys, test.expectedKeys) {
			t.Errorf("WhichKeyValuePairsInMapIntString(%v, %v) should be %v with keys %v", test.Map1, test.Map2, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichKeyValuePairsInMapIntString() {
	keys, ok := WhichKeyValuePairsInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "a", 3: "c", 2: "b"},
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapIntString(
		map[int]string{1: "a", 2: "b"},
		map[int]string{1: "a", 3: "c", 2: "B"},
	)
	fmt.Println(ok, keys)
	// Output:
	// true map[1:a 2:b]
	// true map[1:a]
}

func TestAllKeyValuePairsInMapIntInt(t *testing.T) {
	tests := []struct {
		Map1     map[int]int
		Map2     map[int]int
		expected bool
	}{
		{map[int]int{}, map[int]int{1: 1, 2: 2}, false},
		{map[int]int{1: 1, 2: 1}, map[int]int{}, false},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}, true},
		{map[int]int{1: 1, 2: 2}, map[int]int{11: 1, 12: 2}, false},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 3: 3, 2: 2}, true},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 3: 3, 2: 40}, false},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapIntInt(test.Map1, test.Map2) != test.expected {
			t.Errorf("AllKeyValuePairsInMapIntInt(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapIntInt() {
	fmt.Println(AllKeyValuePairsInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 1, 2: 2},
	))
	fmt.Println(AllKeyValuePairsInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 1, 3: 3, 2: 40},
	))
	// Output:
	// true
	// false
}

func TestAnyKeyValuePairInMapIntInt(t *testing.T) {
	tests := []struct {
		Map1     map[int]int
		Map2     map[int]int
		expected bool
	}{
		{map[int]int{}, map[int]int{1: 1, 2: 2}, false},
		{map[int]int{1: 1, 2: 1}, map[int]int{}, false},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 23}, true},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 3: 3, 2: 23}, true},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 21, 3: 23, 2: 40}, false},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapIntInt(test.Map1, test.Map2) != test.expected {
			t.Errorf("AnyKeyValuePairInMapIntInt(%v, %v) should be %v", test.Map1, test.Map2, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapIntInt() {
	fmt.Println(AnyKeyValuePairInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 1, 2: 23},
	))
	fmt.Println(AnyKeyValuePairInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 21, 3: 3, 2: 40},
	))
	// Output:
	// true
	// false
}

func TestWhichKeyValuePairsInMapIntInt(t *testing.T) {
	tests := []struct {
		Map1         map[int]int
		Map2         map[int]int
		expectedBool bool
		expectedKeys map[int]int
	}{
		{map[int]int{}, map[int]int{1: 1, 2: 2}, false, map[int]int{}},
		{map[int]int{1: 1, 2: 1}, map[int]int{}, false, map[int]int{}},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 23}, true, map[int]int{1: 1}},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 3: 3, 2: 2}, true, map[int]int{1: 1, 2: 2}},
		{map[int]int{1: 1, 2: 2}, map[int]int{1: 21, 3: 23, 2: 40}, false, map[int]int{}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapIntInt(test.Map1, test.Map2)
		if ok != test.expectedBool || !AreEqualMapsIntInt(keys, test.expectedKeys) {
			t.Errorf("WhichKeyValuePairsInMapIntInt(%v, %v) should be %v with keys %v", test.Map1, test.Map2, test.expectedBool, test.expectedKeys)
		}
	}
}

func ExampleWhichKeyValuePairsInMapIntInt() {
	keys, ok := WhichKeyValuePairsInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 1, 2: 23},
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapIntInt(
		map[int]int{1: 1, 2: 2},
		map[int]int{1: 21, 3: 3, 2: 40},
	)
	fmt.Println(ok, keys)
	// Output:
	// true map[1:1]
	// false map[]
}

func TestAllKeyValuePairsInMapIntFloat64(t *testing.T) {
	tests := []struct {
		Map1     map[int]float64
		Map2     map[int]float64
		epsilon  float64
		expected bool
	}{
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, 0, false},
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, .1, false},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, 0, false},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, .1, false},
		{map[int]float64{1: .1, 2: .2}, map[int]float64{1: .1, 2: .2}, 0, true},
		{map[int]float64{1: .1, 2: .2}, map[int]float64{11: .1, 12: .2}, 0, false},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: .1, 2: .2}, 0, false},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: .1, 2: .2}, .01, true},
	}
	for _, test := range tests {
		if AllKeyValuePairsInMapIntFloat64(test.Map1, test.Map2, test.epsilon) != test.expected {
			t.Errorf("AllKeyValuePairsInMapIntFloat64(%v, %v, %v) should be %v", test.Map1, test.Map2, test.epsilon, test.expected)
		}
	}
}

func ExampleAllKeyValuePairsInMapIntFloat64() {
	fmt.Println(AllKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .1, 2: .2},
		map[int]float64{1: .1, 2: .2},
		0,
	))
	fmt.Println(AllKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .11, 2: .2},
		map[int]float64{1: .1, 2: .2},
		0,
	))
	fmt.Println(AllKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .11, 2: .2},
		map[int]float64{1: .1, 2: .2},
		.01,
	))
	// Output:
	// true
	// false
	// true
}

func TestAnyKeyValuePairInMapIntFloat64(t *testing.T) {
	tests := []struct {
		Map1     map[int]float64
		Map2     map[int]float64
		epsilon  float64
		expected bool
	}{
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, 0, false},
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, .1, false},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, 0, false},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, .1, false},
		{map[int]float64{1: .1, 2: .2}, map[int]float64{1: .1, 2: 2.2}, 0, true},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: 2.1, 2: 2.2}, 0, false},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: .1, 2: 2.2}, .01, true},
	}
	for _, test := range tests {
		if AnyKeyValuePairInMapIntFloat64(test.Map1, test.Map2, test.epsilon) != test.expected {
			t.Errorf("AnyKeyValuePairInMapIntFloat64(%v, %v, %v) should be %v", test.Map1, test.Map2, test.epsilon, test.expected)
		}
	}
}

func ExampleAnyKeyValuePairInMapIntFloat64() {
	fmt.Println(AnyKeyValuePairInMapIntFloat64(
		map[int]float64{1: .1, 2: .2},
		map[int]float64{1: .1, 2: 2.2},
		0,
	))
	fmt.Println(AnyKeyValuePairInMapIntFloat64(
		map[int]float64{1: .11, 2: 2.2},
		map[int]float64{1: .1, 2: .2},
		0,
	))
	fmt.Println(AnyKeyValuePairInMapIntFloat64(
		map[int]float64{1: .11, 2: .2},
		map[int]float64{1: .1, 2: 2.2},
		.01,
	))
	// Output:
	// true
	// false
	// true
}

func TestWhichKeyValuePairsInMapIntFloat64(t *testing.T) {
	tests := []struct {
		Map1         map[int]float64
		Map2         map[int]float64
		epsilon      float64
		expectedBool bool
		expectedKeys map[int]float64
	}{
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, 0, false, map[int]float64{}},
		{map[int]float64{}, map[int]float64{1: .1, 2: 2.}, .1, false, map[int]float64{}},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, 0, false, map[int]float64{}},
		{map[int]float64{1: .1, 2: .1}, map[int]float64{}, .1, false, map[int]float64{}},
		{map[int]float64{1: .1, 2: .2}, map[int]float64{1: .1, 2: 2.2}, 0, true, map[int]float64{1: .1}},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: 2.1, 2: 2.2}, 0, false, map[int]float64{}},
		{map[int]float64{1: .11, 2: .2}, map[int]float64{1: .1, 2: 2.2}, .01, true, map[int]float64{1: .11}},
	}
	for _, test := range tests {
		keys, ok := WhichKeyValuePairsInMapIntFloat64(test.Map1, test.Map2, test.epsilon)
		if ok != test.expectedBool || !AreEqualMapsIntFloat64(keys, test.expectedKeys, 0) {
			t.Errorf(
				"WhichKeyValuePairsInMapIntFloat64(%v, %v, %v) should be %v with keys %v, not %v and %v",
				test.Map1, test.Map2, test.epsilon, test.expectedBool, test.expectedKeys, ok, keys,
			)
		}
	}
}

func ExampleWhichKeyValuePairsInMapIntFloat64() {
	keys, ok := WhichKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .1, 2: .2},
		map[int]float64{1: .1, 2: 2.2},
		0,
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .11, 2: 2.2},
		map[int]float64{1: .1, 2: .2},
		0,
	)
	fmt.Println(ok, keys)
	keys, ok = WhichKeyValuePairsInMapIntFloat64(
		map[int]float64{1: .11, 2: .2},
		map[int]float64{1: .1, 2: 2.2},
		.01,
	)
	fmt.Println(ok, keys)
	// Output:
	// true map[1:0.1]
	// false map[]
	// true map[1:0.11]
}
