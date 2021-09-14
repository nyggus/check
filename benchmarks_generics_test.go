package check

import (
	"testing"
)

/*
 The file contains a generic function allValuesInSlice(Slice1 interface{}, Slice2 interface{}), which that works with []string and []int slices. See the benchmarks:

goos: linux
goarch: amd64
pkg: github.com/nyggus/check
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkAllValuesInSliceString-8       18360554                69.73 ns/op
BenchmarkAllValuesInStringSlice-8       19585856                63.80 ns/op
BenchmarkAllValuesInSliceInt-8          62096694                19.67 ns/op
BenchmarkAllValuesInIntSlice-8          69703411                17.27 ns/op

The benchmarks compare this generic function with the corresponding functions for []string and []int slices. As we can see, the decrease in performance is rather small: about 9% for []string but about 13% for []int slices. Of course, these measures will slightly change from run to run. For much longer slices, however, the decrease might be negligible whatsoever.

Here are benchmarks for IsUnique against its three subcomponents, IsUniqueStringSlice, IsUniqueIntSlice and IsUniqueFloat64Slice:

goos: linux
goarch: amd64
pkg: github.com/nyggus/check
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkIsUniqueForStringsGeneric-8              137904              8294 ns/op
BenchmarkIsUniqueForStringsNotGeneric-8           142462              8072 ns/op
BenchmarkIsUniqueForIntsGeneric-8                 239313              4390 ns/op
BenchmarkIsUniqueForIntsNotGeneric-8              266353              4687 ns/op
BenchmarkIsUniqueForFloat64Generic-8                 146           7684877 ns/op
BenchmarkIsUniqueForFloat64NotGeneric-8              156           7587926 ns/op

Of course, the generic version is never slower than its original counterpart, despite what we see. The results do show, however, that for big slices (as those used here for integers and floats) the decrease in performance can even go unnoticed.

*/

func BenchmarkIsUniqueForStringsGeneric(b *testing.B) {
	slice := []string{
		"a", "g", "e", "h", "b", "r", "c",
		"A", "G", "E", "F", "H", "I", "J",
		"1", "2", "3", "4", "5", "6", "7",
		"1 ", "2 ", "3 ", "4 ", "5 ", "6 ", "7 ",
		" 1", " 2", " 3", " 4", " 5", " 6", " 7",
		"11", "12", "13", "14", "15", "16", "17",
		"21", "22", "23", "24", "25", "26", "27",
		"31", "32", "33", "34", "35", "36", "37",
		"41", "42", "43", "44", "45", "46", "47",
		"51", "52", "53", "54", "55", "56", "57",
		"61", "62", "63", "64", "65", "66", "66",
	}
	for n := 0; n < b.N; n++ {
		IsUnique(slice)
	}
}

func BenchmarkIsUniqueForStringsNotGeneric(b *testing.B) {
	slice := []string{
		"a", "g", "e", "h", "b", "r", "c",
		"A", "G", "E", "F", "H", "I", "J",
		"1", "2", "3", "4", "5", "6", "7",
		"1 ", "2 ", "3 ", "4 ", "5 ", "6 ", "7 ",
		" 1", " 2", " 3", " 4", " 5", " 6", " 7",
		"11", "12", "13", "14", "15", "16", "17",
		"21", "22", "23", "24", "25", "26", "27",
		"31", "32", "33", "34", "35", "36", "37",
		"41", "42", "43", "44", "45", "46", "47",
		"51", "52", "53", "54", "55", "56", "57",
		"61", "62", "63", "64", "65", "66", "66",
	}
	for n := 0; n < b.N; n++ {
		IsUniqueStringSlice(slice)
	}
}

func BenchmarkIsUniqueForIntsGeneric(b *testing.B) {
	slice := []int{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}
	slice = append(slice, 0)
	for n := 0; n < b.N; n++ {
		IsUnique(slice)
	}
}

func BenchmarkIsUniqueForIntsNotGeneric(b *testing.B) {
	slice := []int{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}
	slice = append(slice, 0)
	for n := 0; n < b.N; n++ {
		IsUniqueIntSlice(slice)
	}
}

func BenchmarkIsUniqueForFloat64Generic(b *testing.B) {
	slice := []float64{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, float64(i)/500)
	}
	slice = append(slice, 2)
	for n := 0; n < b.N; n++ {
		IsUnique(slice)
	}
}

func BenchmarkIsUniqueForFloat64NotGeneric(b *testing.B) {
	slice := []float64{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, float64(i)/500)
	}
	slice = append(slice, 2)
	for n := 0; n < b.N; n++ {
		IsUniqueFloat64Slice(slice, 0)
	}
}

func TestBenchmark(t *testing.T) {
	slice := []float64{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, float64(i)/500)
	}
	if !IsValueInFloat64Slice(2, slice, 0) {
		t.Error("slice should include 2. but does not")
	}
}

// allValuesInSlice is a generic function to work with either string or int slices.
// Its only aim is benchmarking.
func allValuesInSlice(Slice1 interface{}, Slice2 interface{}) bool {
	if slice1, ok := Slice1.([]int); ok {
		return AllValuesInIntSlice(slice1, Slice2.([]int))
	}
	if slice1, ok := Slice1.([]string); ok {
		return AllValuesInStringSlice(slice1, Slice2.([]string))
	}
	return false
}

func BenchmarkAllValuesInSliceString(b *testing.B) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "a", "g", "e", "e", "b", "r", "c"}
	for n := 0; n < b.N; n++ {
		allValuesInSlice(s1, s2)
	}
}
func BenchmarkAllValuesInStringSlice(b *testing.B) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "a", "g", "e", "e", "b", "r", "c"}
	for n := 0; n < b.N; n++ {
		AllValuesInStringSlice(s1, s2)
	}
}

func BenchmarkAllValuesInSliceInt(b *testing.B) {
	s1 := []int{1, 100, 1000}
	s2 := []int{1, 0, 10, 0, 0, 10, 100, 99, 1000}
	for n := 0; n < b.N; n++ {
		allValuesInSlice(s1, s2)
	}
}
func BenchmarkAllValuesInIntSlice(b *testing.B) {
	s1 := []int{1, 100, 1000}
	s2 := []int{1, 0, 10, 0, 0, 10, 100, 99, 1000}
	for n := 0; n < b.N; n++ {
		AllValuesInIntSlice(s1, s2)
	}
}

func TestAllValuesInSlice(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "a", "g", "c", "b"}
	is := allValuesInSlice(s1, s2)
	if !is {
		t.Errorf("ERROR!")
	}
	si1 := []int{1, 2, 3}
	si2 := []int{0, 0, 3, 1, 2}
	is = allValuesInSlice(si1, si2)
	if !is {
		t.Errorf("ERROR!")
	}
	s1 = []string{"a", "b", "c"}
	s2 = []string{"h", "h", "g"}
	is = allValuesInSlice(s1, s2)
	if is {
		t.Errorf("ERROR!")
	}
	si1 = []int{1, 2, 3}
	si2 = []int{0, 0, 0}
	is = allValuesInSlice(si1, si2)
	if is {
		t.Errorf("ERROR!")
	}
}
