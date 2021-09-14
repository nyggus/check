# Examples

## Boolean values

### Slices

`Any()` and `All()` work with boolean slices (`[]bool`):

```go
Any([]bool{true, true}) // true
Any([]bool{true, false}) // true
Any([]bool{false, false}) // false

All([]bool{true, true}) // true
All([]bool{true, false}) // false
All([]bool{false, false}) // false
```

### Maps

The `check` package offers also functions to check if all values of a map are true, or any is true, or which is true. You can do so for `map[int]bool` and `map[string]bool`:

#### `map[string]bool`

```go
AnyInMapString(map[string]bool{"key1": true, "key2": true}) // true
AnyInMapString(map[string]bool{"key1": true, "key2": false}) // true
AnyInMapString(map[string]bool{"key1": false, "key2": false}) // false

AllInMapString(map[string]bool{"key1": true, "key2": true}) // true
AllInMapString(map[string]bool{"key1": true, "key2": false}) // false
AllInMapString(map[string]bool{"key1": false, "key2": false}) // false

// This is how you should use Which... functions:
keys, ok := WhichInMapString(map[string]bool{"key1": true, "key2": true}) // keys = [key1 key2], ok = true
WhichInMapString(map[string]bool{"key1": true, "key2": false}) // [key1] true
WhichInMapString(map[string]bool{"key1": false, "key2": false}) // [] false
```

#### `map[int]bool`

```go
AnyInMapInt(map[int]bool{1: true, 2: true}) // true
AnyInMapInt(map[int]bool{1: true, 2: false}) // true
AnyInMapInt(map[int]bool{1: false, 2: false}) // false

AllInMapInt(map[int]bool{1: true, 2: true}) // true
AllInMapInt(map[int]bool{1: true, 2: false}) // false
AllInMapInt(map[int]bool{1: false, 2: false}) // false

WhichInMapInt(map[int]bool{1: true, 2: true}) // [1 2] true
WhichInMapInt(map[int]bool{1: true, 2: false}) // [1] true
WhichInMapInt(map[int]bool{1: false, 2: false}) // [] false
```

## Check if a slice is unique

These functions check if a slice is unique, but the package also offers functions to create a slice with unique elements of a slice. See [here](#return-a-unique-slice).

#### `[]int`

```go
IsUniqueIntSlice([]int{1, 2}) // false
IsUniqueIntSlice([]int{1, 1}) // true
```

#### `[]string`

```go
IsUniqueStringSlice([]string{"a", "b"}) // false
IsUniqueStringSlice([]string{"a", "A"}) // false
IsUniqueStringSlice([]string{"a", "a"}) // true
```

#### `[]float64`

Here, you need to provide the value of the `Epsilon` parameter, which affects the comparison of two floats. Alternatively, you could first round the slice's values and use `Epsilon = 0`. To use full accuracy, set `Epsilon` to 0.

```go
IsUniqueFloat64Slice([]float64{.01, .01}, 0) // true
IsUniqueFloat64Slice([]float64{.01, .01001}, 0) // false
IsUniqueFloat64Slice([]float64{.01, .01001}, .000000001) // true
```

#### Generic version

The only generic function that `check` offers is `IsUnique`; it works for the three types of slices above. Think twice before using it, however, since it decreases performance a little bit, as you can read in the benchmarks shown in [README](https://github.com/nyggus/check/blob/master/README.md).

```go
IsUnique([]int{1, 2}) // false
IsUnique([]string{"a", "A"}) // true
IsUnique([]float64{.01, .01001}) // true
```

## Check if a map has unique values

These functions check if a map has unique values. Do note that this does not mean comparing key-value pairs, since they are always unique.

#### `map[int]int`

```go
IsUniqueMapIntInt(map[int]int{1: 1, 2: 1, 3: 1, 100: 1}) // false
IsUniqueMapIntInt(map[int]int{1: 1, 2: 2, 3: 3, 100: 4}) // true
```

#### `map[int]string`

```go
IsUniqueMapIntString(map[int]string{1: "a", 2: "a", 3: "a", 100: "a"}) // false
IsUniqueMapIntString(map[int]string{1: "a", 2: "b", 3: "B", 100: "A"}) // true
```

#### `map[int]float64`

```go
IsUniqueMapIntFloat64(map[int]float64{1: .1, 2: .11, 3: .101, 100: .1105}, 0)) // true
IsUniqueMapIntFloat64(map[int]float64{1: .1, 2: .11, 3: .101, 100: .1105}, .1) // false	
```

#### `map[string]int`

```go
IsUniqueMapStringInt(
    map[string]int{"a": 1, "b": 21, "c": 11, "Blues is my soul, blues is my heart": 2},
) // true
IsUniqueMapStringInt(
    map[string]int{"a": 1, "b": 1, "c": 1, "Blues is my soul, blues is my heart": 1},
) // false
```

#### `map[string]string`

```go
IsUniqueMapStringString(
    map[string]string{"key1": "a", "key2": "a", "key3": "Q"},
) // false
IsUniqueMapStringString(
    map[string]string{"key1": "a", "key2": "Q", "key3": "A"},
) // true
```

#### `map[string]float64`

```go
IsUniqueMapStringFloat64(map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, 0) // true
IsUniqueMapStringFloat64(map[string]float64{"a": .1, "b": .11, "c": .101, "Zigi says hi!": .1105}, .1) // false
```

## Compare two slices

You have two options:

* compare two slices in their original order, meaning that they differ even when they have the same values but are sorted differently; thus, `[]int{1, 2}` **is not equal** to `[]int{2, 1}`
* compare two slices ignoring the ordering; thus, `[]int{1, 2}` **is equal** to `[]int{2, 1}`

### Compare two slices in their original order

#### `[]string`

```go
AreEqualSlicesString([]int{"a", "b", "a", "b"}, []int{"a", "b", "a", "b"}) // true
AreEqualSlicesString([]int{"a", "b", "b", "a"}, []int{"a", "b", "a", "b"}) // false
```

Of course, two slices with different lengths are different:

```go
AreEqualSlicesString([]int{"a", "b", "a", "b"}, []int{"a", "b", "a"}) // false
```

#### `[]int`

```go
AreEqualSlicesInt([]string{1, 1, 1, 2}, []string{1, 1, 1, 2}) // true
AreEqualSlicesInt([]string{1, 1, 1, 2}, []string{1, 1, 2, 1}) // false
```

#### `[]float64`

Again, we need to provide the value of the `Epsilon` parameter:

```go
AreEqualSlicesFloat64([]float64{1, 1, 1, 2}, []float64{1, 1, 1, 2}, 0) // true
AreEqualSlicesFloat64([]float64{1, 1, 1, 2}, []float64{1, 1, 2, 1}, .0000001) // false
AreEqualSlicesFloat64([]float64{1.001, 1, 1, 2}, []float64{1, 1, 1, 2}, .0000001) // false
AreEqualSlicesFloat64([]float64{1.001, 1, 1, 2}, []float64{1, 1, 1, 2}, .1) // true
```

### Compare two slices ignoring their ordering

#### `[]string`

```go
AreEqualSortedSlicesString([]string{"a", "b", "a", "b"}, []string{"a", "b", "a", "b"}) // true
AreEqualSortedSlicesString([]string{"a", "b", "b", "a"}, []string{"a", "b", "a", "b"}) // true
```

#### `[]int`

```go
AreEqualSortedSlicesInt([]int{1, 1, 1, 2}, []int{1, 1, 1, 2}) // true
AreEqualSortedSlicesInt([]int{1, 1, 1, 2}, []int{1, 1, 2, 1}) // true
```

#### `[]float64`

```go
AreEqualSortedSlicesFloat64([]float64{1, 1, 1, 2}, []float64{1, 1, 1, 2}, .0000001) // true
AreEqualSortedSlicesFloat64([]float64{1, 1, 1, 2}, []float64{1, 1, 2, 1}, .0000001) // true
AreEqualSortedSlicesFloat64([]float64{1.001, 1, 1, 2}, []float64{1, 1, 1, 2}, .0000001) // false
AreEqualSortedSlicesFloat64([]float64{1.001, 1, 1, 2}, []float64{1, 1, 1, 2}, .1) // true
```

## Check if a value is in a slice

These functions enable you to check whether a slice contains a particular value; [later](#check-if-any-or-all-values-from-a-slice-are-in-another-slice) you will also find a way to check whether a slice contains any or all of several values, or which values (from another slice) it contains.

#### `[]string`

```go
IsValueInStringSlice("Whatever", []string{"Nothing", "Whatever"}) // true
IsValueInStringSlice("Whatever", []string{"Nothing", "Whatever else"}) // false
```

#### `[]int`

```go
IsValueInIntSlice(1, []int{5, 10, 100, 1}) // true
IsValueInIntSlice(1, []int{5, 10, 100, 10}) // false
```

#### `[]float64`

```go
IsValueInFloat64Slice(1., []int{5., 10., 100., 1.}, 0) // true
IsValueInFloat64Slice(1., []int{5., 10., 100., 1.01}, 0) // false
IsValueInFloat64Slice(1., []int{5., 10., 100., 1.01}, .1) // true
```

(Remember about the `Epsilon` parameter.)

## Check if any or all values from a slice are in another slice, or which values are

With these functions, you can check whether a slice contains any or all of several values, provided as another slice (of course, of the same type). You can also check which value from one slice is in another slice. The any- and all-functions return a boolean value while the which-functions return a tuple with two values: a map that provides the indices from `Slice1` as keys and the corresponding values as the map's values; and the same boolean value as the corresponding any-function.

#### `[]string`

```go
AnyValueInStringSlice([]string{"a", "b"}, []string{"a", "b", "c"}) // true
AnyValueInStringSlice([]string{"a", "b"}, []string{"a", "c", "d"}) // true
AnyValueInStringSlice([]string{"a", "b"}, []string{"f", "c", "d"}) // false

AllValuesInStringSlice([]string{"a", "b"}, []string{"a", "b", "c"}) // true
AllValuesInStringSlice([]string{"a", "b"}, []string{"a", "c", "d"}) // false

values, ok := WhichValuesInStringSlice(
    []string{"10", "Shout Bamalama!"},
    []string{"10", "Shout Bamalama!", "50"},
) // values = map[10:[0] Shout Bamalama![1]] key = true
values, ok = WhichValuesInStringSlice(
    []string{"Sing", "a", "sing", "about", "the", "sun!"},
    []string{"sing", "sun", "the"},
) // values = map[sing:[2] the:[4]] key = true
```

#### `[]int`

```go
AnyValueInIntSlice([]int{1, 2, 3}, []int{1, 2, 3, 4}) // true
AnyValueInIntSlice([]int{1, 2, 3}, []int{1, 3, 4, 5}) // true
AnyValueInIntSlice([]int{1, 2, 3}, []int{4, 4, 4, 5}) // false

AllValuesInIntSlice([]int{1, 2, 3}, []int{1, 2, 3, 4}) // true
AllValuesInIntSlice([]int{1, 2, 3}, []int{1, 3, 4, 5}) // false

values, ok := WhichValuesInIntSlice([]int{10, 20}, []int{10, 30, 50}) // values = map[10:[0]] ok = true
values, ok = WhichValuesInIntSlice([]int{10, 10, 10}, []int{20, 20, 10}) // values = map[10:[2]] ok = true
```

#### `[]float64`

```go
AnyValueInFloat64Slice([]float64{.001, .002, .0022}, []float64{.001, .002}, 0) // true
AnyValueInFloat64Slice([]float64{.002, .0022, .0022}, []float64{.001, .002}, .001)  // false
AnyValueInFloat64Slice([]float64{1, 2, 3}, []float64{1, 3, 4, 5}, 0) // true

AllValuesInFloat64Slice([]float64{.001, .002, .0022}, []float64{.001, .002}, 0) // false
AllValuesInFloat64Slice([]float64{.001, .002, .0022}, []float64{.001, .002}, .001)  // true

values, ok := WhichValuesInFloat64Slice([]float64{10, 20, 70}, []float64{10, 20}, .00000001) // map[10:[0] 20:[1]] true
values, ok = WhichValuesInFloat64Slice([]float64{.1, .2, .2}, []float64{.3, .11, .3}, .1) // map[0:0.1 1:0.2 2:0.2] true
values, ok = WhichValuesInFloat64Slice([]float64{.1, .2, .2}, []float64{.3, .11, .3}, 0) // map[] false
```

(Remember about the meaning of the `Epsilon` parameter for floats.)

## Check if a value is among a map's values

You can use these functions in order to check whether a map contains a particular value. Note that keys are ignored here, only values are looked for; if you want to check if the map contains a key-value pair, see [those functions](#check-if-any-or-all-key-value-pairs-are-in-a-map).

Unlike the other `Is`-functions, these do not return only a boolean value but also a slice with the keys. Look here:

```go
keys, ok := IsValueInMapStringInt(1, map[string]int{"key1": 2, "key2": 1, "key3": 1}) // [key2 key3] true
```
Here, `ok` is `true` because `1` was found among the map's values, and `keys` is a two-element slice, with values being `"key2"` and `"key3"`, for which value `1` was found.


#### `map[string]string`

```go
keys, ok := IsValueInMapStringString(
    "value",
    map[string]string{"key1": "not this one", "key2": "value"},
) // [key2] true
keys, ok = IsValueInMapStringString(
    "value",
    map[string]string{"key1": "not this one", "key2": "no value"},
) // [], false
keys, ok = IsValueInMapStringString(
    "value",
    map[string]string{"key1": "not this one", "key2": "value", "key3": "value"},
) // [key2 key3] true
```

#### `map[string]int`

```go
keys, ok := IsValueInMapStringInt(1, map[string]int{"key1": 2, "key2": 1}) // [key2] true
keys, ok = IsValueInMapStringInt(1, map[string]int{"key1": 2, "key2": 2}) // [] false
keys, ok = IsValueInMapStringInt(1, map[string]int{"key1": 2, "key2": 1, "key3": 1}) // [key2 key3] true
```

#### `map[string]float64`

Again, do remember about the meaning of the `Epsilon` parameter:

```go
keys, ok := IsValueInMapStringFloat64(1, map[string]int{"key1": 2.01, "key2": 1.01}, 0) // [] false
keys, ok = IsValueInMapStringFloat64(1, map[string]int{"key1": 2.01, "key2": 1.01}, .1) // [key2] true
keys, ok = IsValueInMapStringFloat64(
    1,
    map[string]int{"key1": 2.01, "key2": 1.01, "key3": 1.07},
    .1,
) // [key2 key3] true
```

#### `map[int]string`

```go
keys, ok := IsValueInMapIntString("value", map[int]string{1: "not this one", 2: "value"}) // [2] true
keys, ok = IsValueInMapIntString("value", map[int]string{1: "not this one", 2: "no value"}) // [] false
keys, ok = IsValueInMapIntString(
    "value",
    map[int]string{1: "not this one", 2: "value", 3: "value"},
) // [2 3] true
```

#### `map[int]int`

```go
keys, ok := IsValueInMapIntInt(1, map[int]int{1: 2, 2: 1}) // [2] true
keys, ok = IsValueInMapIntInt(1, map[int]int{1: 2, 2: 2}) // [] false
keys, ok = IsValueInMapIntInt(1, map[int]int{1: 2, 2: 1, 3: 1}) // [2 3] true
```

#### `map[int]float64`

```go
keys, ok := IsValueInMapIntFloat64(1, map[int]int{1: 2.01, 2: 1.01}, 0) // [] false
keys, ok = IsValueInMapIntFloat64(1, map[int]int{1: 2.01, 2: 1.01}, .1) // [2] true
keys, ok = IsValueInMapIntFloat64(1, map[int]int{1: 2.01, 2: 1.01, 3: 1.07}, .1) // [2 3] true
```

## Check if any or all values from a slice are among a map's values, or which are

Any- and all-functions enable you to check whether a map contains any or all values, provided as a slice. You can also check which values from the slice are in the map, using which-functions. The any- and all-functions return a boolean value while the which-functions return a tuple: a map that provides the values from `Slice1` as keys and the corresponding values as the map's keys, and the same boolean value as the corresponding any-function.

#### `map[string]string`

```go
AnyValueInMapStringString(
    []string{"a", "b", "c"},
    map[string]string{"1": "a", "2": "b"},
) // true
AnyValueInMapStringString(
    []string{"a", "b", "c"},
    map[string]string{"1": "nothing", "2": "else", "3": "matters"},
) // false
AnyValueInMapStringString(
    []string{"a", "b", "z"},
    map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"},
) // true

AllValuesInMapStringString(
    []string{"a", "b", "c"},
    map[string]string{"1": "a", "2": "b"},
) // false
AllValuesInMapStringString(
    []string{"a", "b", "c"},
    map[string]string{"1": "nothing", "2": "else", "3": "matters"},
) // false
AllValuesInMapStringString(
    []string{"a", "b", "z"},
    map[string]string{"1": "a", "2": "b", "3": "haha", "5": "z", "100": "Zigi says hi!"},
) // true

values, ok := WhichValuesInMapStringString(
    []string{"a", "b", "c"},
    map[string]string{"1": "a", "2": "b"},
) // map[a:[1] b:[2]] true
```

#### `map[string]int`

```go
AnyValueInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}) // true
AnyValueInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 30, "2": 40}) // false
AnyValueInMapStringInt(
    []int{1, 2, 7000},
    map[string]int{"1": 10, "2": 10, "3": 1, "5": 1, "100": 7000},
) // true

AllValuesInMapStringInt(
    []int{10, 20, 70},
    map[string]int{"1": 10, "2": 20, "a": 70},
) // true
AllValuesInMapStringInt([]int{10, 20, 70}, map[string]int{"1": 10, "2": 20}) // false
AllValuesInMapStringInt(
    []int{1, 2, 7000},
    map[string]int{"1": 10, "2": 10, "3": 1, "5": 2, "100": 7000},
) // true

values, ok := WhichValuesInMapStringInt(
    []int{10, 20, 7},
    map[string]int{"a": 10, "G": 7, "H": 190},
) // map[7:[G] 10:[a]] true
```

#### `map[string]float64`

```go
AnyValueInMapStringFloat64(
    []float64{.01, .2},
    map[string]float64{"1": .01, "2": .2},
    0,
) // true
AnyValueInMapStringFloat64(
    []float64{.001, .02, 70.8},
    map[string]float64{"1": .0011, "2": .021},
    .01,
) // true
AnyValueInMapStringFloat64(
    []float64{.01, .02, 70.8},
    map[string]float64{"1": .0111, "2": .0211},
    .001,
) // false

AllValuesInMapStringFloat64(
    []float64{.01, .2},
    map[string]float64{"1": .01, "2": .2},
    0,
) // true
AllValuesInMapStringFloat64(
    []float64{.001, .02, 70.8},
    map[string]float64{"1": .0011, "2": .021},
    .01,
) // false

values, ok := WhichValuesInMapStringFloat64(
    []float64{.01, .2},
    map[string]float64{"a": .01, "b": .2},
    0,
) // map[0.01:[a] 0.2:[b]] true
values, ok = WhichValuesInMapStringFloat64(
    []float64{.01002, .01, .2},
    map[string]float64{"a": .01, "b": .011},
    .0001
) // map[0.01:[a] 0.01002:[a]] true
```

#### `map[int]string`

```go
AnyValueInMapIntString([]string{"a", "b"}, map[int]string{1: "v", 2: "g"}) // false
AnyValueInMapIntString([]string{"a", "b"}, map[int]string{1: "a", 2: "g"}) // true

AllValuesInMapIntString([]string{"a", "b", "C"}, map[int]string{1: "a", 2: "b", 3: "C"}) // true
AllValuesInMapIntString([]string{"a", "C"}, map[int]string{1: "a", 2: "b"}) // false

keys, ok := WhichValuesInMapIntString(
    []string{"a", "b", "c"},
    map[int]string{1: "a", 2: "b"},
) // map[a:[1] b:[2]] true
keys, ok = WhichValuesInMapIntString(
    []string{"a", "b", "c"},
    map[int]string{1: "F", 2: "H"},
) // map[] false
```

#### `map[int]int`

```go
AnyValueInMapIntInt([]int{10, 20, 70}, map[int]int{1: 10, 2: 20}) // true
AnyValueInMapIntInt([]int{10, 20, 70}, map[int]int{1: 30, 2: 40}) // false
AnyValueInMapIntInt([]int{1, 2, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}) // true

AllValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 10, 2: 20}) // false
AllValuesInMapIntInt([]int{10, 20, 70}, map[int]int{1: 20, 2: 10, 5: 70}) // true
AllValuesInMapIntInt([]int{10, 1, 7000}, map[int]int{1: 10, 2: 10, 3: 1, 5: 1, 100: 7000}) // true

values, ok := WhichValuesInMapIntInt(
    []int{10, 20, 70},
    map[int]int{1: 10, 2: 20},
) // true map[10:[1] 20:[2]]
```

#### `map[int]float64`

```go
AnyValueInMapIntFloat64([]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, .01)) // true
AnyValueInMapIntFloat64([]float64{.001, .02}, map[int]float64{1: .0011, 2: .021}, 0) // false

AllValuesInMapIntFloat64([]float64{.01, .2}, map[int]float64{1: .01, 2: .2}, 0) // true
AllValuesInMapIntFloat64([]float64{.001, .02, 70.8}, map[int]float64{1: .0011, 2: .021}, .01) // false

values, ok := WhichValuesInMapIntFloat64(
    []float64{.01, .2},
    map[int]float64{1: .01, 2: .2},
    0,
) // map[0.01:[1] 0.2:[2]] true
values, ok = WhichValuesInMapIntFloat64(
    []float64{.01002, .01, .2},
    map[int]float64{1: .01, 2: .011},
    .0001,
) // map[0.01:[1] 0.01002:[1]] true
```

## Check if any or all key-value pairs are in a map, or which is

#### `map[string]string`

```go
AnyKeyValuePairInMapStringString(
    map[string]string{"1": "a", "2": "b"},
    map[string]string{"1": "a", "3": "c", "2": "X"},
) // true
AnyKeyValuePairInMapStringString(
    map[string]string{"1": "a", "2": "b"},
    map[string]string{"1": "B", "3": "c", "2": "B"},
) // false

AllKeyValuePairsInMapStringString(
    map[string]string{"1": "a", "2": "b"},
    map[string]string{"1": "a", "3": "c", "2": "b"},
) // true
AllKeyValuePairsInMapStringString(
    map[string]string{"1": "a", "2": "b"},
    map[string]string{"1": "a", "3": "c", "2": "B"},
) // false

keys, ok := WhichKeyValuePairsInMapStringString(
	map[string]string{"1": "a", "2": "b"},
	map[string]string{"1": "a", "3": "c", "2": "X"},
) // map[1:a] true
```

#### `map[string]int`

```go
AnyKeyValuePairInMapStringInt(map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "2": 23}) // true
AnyKeyValuePairInMapStringInt(map[string]int{"1": 1, "2": 2}, map[string]int{"1": 21, "3": 3, "2": 40})) // false

AllKeyValuePairsInMapStringInt(map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "2": 2}) // true
AllKeyValuePairsInMapStringInt(map[string]int{"1": 1, "2": 2}, map[string]int{"1": 1, "3": 3, "2": 40}) // false

keys, ok := WhichKeyValuePairsInMapStringInt(
	map[string]int{"1": 1, "2": 2},
	map[string]int{"1": 1, "2": 23},
) // map[1:1] true
```

#### `map[string]float64`

```go
AnyKeyValuePairInMapStringFloat64(
    map[string]float64{"1": .1, "2": .2},
    map[string]float64{"1": .1, "2": 2.2},
    0,
) // true
AnyKeyValuePairInMapStringFloat64(
    map[string]float64{"1": .11, "2": 2.2},
    map[string]float64{"1": .1, "2": .2},
    0,
) // false
AnyKeyValuePairInMapStringFloat64(
    map[string]float64{"1": .11, "2": .2},
    map[string]float64{"1": .1, "2": 2.2},
    .01,
) // true

AllKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .1, "2": .2},
    map[string]float64{"1": .1, "2": .2},
    0,
) // true
AllKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .11, "2": .2},
    map[string]float64{"1": .1, "2": .2},
    0,
) // false
AllKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .11, "2": .2},
    map[string]float64{"1": .1, "2": .2},
    .01,
) // true

keys, ok := WhichKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .1, "2": .2},
    map[string]float64{"1": .1, "2": 2.2},
    0,
) // map[1:0.1] true
keys, ok = WhichKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .11, "2": 2.2},
    map[string]float64{"1": .1, "2": .2},
    0,
) // map[] false
keys, ok = WhichKeyValuePairsInMapStringFloat64(
    map[string]float64{"1": .11, "2": .2},
    map[string]float64{"1": .1, "2": 2.2},
    .01,
) // map[1:0.11] true
```

#### `map[int]string`

```go
AnyKeyValuePairInMapIntString(
    map[int]string{1: "a", 2: "b"},
    map[int]string{1: "a", 3: "c", 2: "X"},
) // true
AnyKeyValuePairInMapIntString(
    map[int]string{1: "a", 2: "b"},
    map[int]string{1: "B", 3: "c", 2: "B"},
) // false

AllKeyValuePairsInMapIntString(
    map[int]string{1: "a", 2: "b"},
    map[int]string{1: "a", 3: "c", 2: "b"},
) // true
AllKeyValuePairsInMapIntString(
    map[int]string{1: "a", 2: "b"},
    map[int]string{1: "a", 3: "c", 2: "B"},
) // false

keys, ok := WhichKeyValuePairsInMapIntInt(
    map[int]int{1: 1, 2: 2},
    map[int]int{1: 1, 2: 23},
) // map[1:1] true
keys, ok = WhichKeyValuePairsInMapIntInt(
    map[int]int{1: 1, 2: 2},
    map[int]int{1: 21, 3: 3, 2: 40},
) // map[] false
```

#### `map[int]int`

```go
AnyKeyValuePairInMapIntInt(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 23}) // true
AnyKeyValuePairInMapIntInt(map[int]int{1: 1, 2: 2}, map[int]int{1: 21, 3: 3, 2: 40}) // false

AllKeyValuePairsInMapIntInt(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}) // true
AllKeyValuePairsInMapIntInt(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 3: 3, 2: 40}) // false

keys, ok := WhichKeyValuePairsInMapIntInt(
    map[int]int{1: 1, 2: 2},
    map[int]int{1: 1, 2: 23},
) // map[1:1] true
```

#### `map[int]float64`

Since we're comparing float values, remember to use the `Epsilon` argument:

```go
AnyKeyValuePairInMapIntFloat64(
    map[int]float64{1: .1, 2: .2},
    map[int]float64{1: .1, 2: 2.2},
    0,
) // true
AnyKeyValuePairInMapIntFloat64(
    map[int]float64{1: .11, 2: 2.2},
    map[int]float64{1: .1, 2: .2},
    0,
) // false
AnyKeyValuePairInMapIntFloat64(
    map[int]float64{1: .11, 2: .2},
    map[int]float64{1: .1, 2: 2.2},
    .01,
) // true

AllKeyValuePairsInMapIntFloat64(
    map[int]float64{1: .1, 2: .2},
    map[int]float64{1: .1, 2: .2},
    0,
) // true
AllKeyValuePairsInMapIntFloat64(
    map[int]float64{1: .11, 2: .2},
    map[int]float64{1: .1, 2: .2},
    0,
) // false
AllKeyValuePairsInMapIntFloat64(
    map[int]float64{1: .11, 2: .2},
    map[int]float64{1: .1, 2: .2},
    .01,
) // true

keys, ok := WhichKeyValuePairsInMapIntFloat64(
    map[int]float64{1: .1, 2: .2},
    map[int]float64{1: .1, 2: 2.2},
    0,
) // map[1:0.1] true
keys, ok = WhichKeyValuePairsInMapIntFloat64(
    map[int]float64{1: .11, 2: .2},
    map[int]float64{1: .1, 2: 2.2},
    .01,
)// map[1:0.11] true
```

## Comparison of two maps

When comparing two maps, remember that maps are unordered, so the below functions compare whether the maps have the same key-value pairs, ignoring the maps' ordering (which is random).

#### `map[string]string`

```go
AreEqualMapsStringString(
    map[string]string{"1": "the first", "2": "the second"},
    map[string]string{"1": "the first", "2": "the second"},
) // true
AreEqualMapsStringString(
    map[string]string{"1": "the first", "2": "the second"},
    map[string]string{"2": "the second", "1": "the first"},
) // true
AreEqualMapsStringString(
    map[string]string{"1": "the first", "2": "the second"},
    map[string]string{"2": "the third", "1": "the first"},
) // false
```

#### `map[string]int`

```go
AreEqualMapsStringInt(
    map[string]int{"1": 1, "2": 2},
    map[string]int{"1": 1, "2": 2},
) // true
AreEqualMapsStringInt(
    map[string]int{"1": 1, "2": 2},
    map[string]int{"2": 2, "1": 1},
) // true
AreEqualMapsStringInt(
    map[string]int{"1": 1, "2": 2},
    map[string]int{"2": 3, "1": 1},
) // false
```

#### `map[string]float64`

```go
AreEqualMapsStringFloat64(
    map[string]float64{"1": 1., "2": 2.},
    map[string]float64{"1": 1., "2": 2.},
    0,
) // true
AreEqualMapsStringFloat64(
    map[string]float64{"1": 1., "2": 2.},
    map[string]float64{"1": 1.00001, "2": 2.},
    0,
) // false
AreEqualMapsStringFloat64(
    map[string]float64{"1": 1., "2": 2.},
    map[string]float64{"1": 1.00001, "2": 2.},
    .001,
) // true
```


#### `map[int]string`

```go
AreEqualMapsIntString(
    map[int]string{1: "the first", 2: "the second"},
    map[int]string{1: "the first", 2: "the second"},
) // true
AreEqualMapsIntString(
    map[int]string{1: "the first", 2: "the second"},
    map[int]string{2: "the second", 1: "the first"},
) // true
AreEqualMapsIntString(
    map[int]string{1: "the first", 2: "the second"},
    map[int]string{2: "the third", 1: "the first"},
) // false
```

#### `map[int]int`

```go
AreEqualMapsIntInt(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}) // true
AreEqualMapsIntInt(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 1}) // true
AreEqualMapsIntInt(map[int]int{1: 1, 2: 2}, map[int]int{2: 3, 1: 1}) // false
```

#### `map[int]float64`

```go
AreEqualMapsIntFloat64(
    map[int]float64{"1": 1., "2": 2.},
    map[int]float64{"1": 1., "2": 2.},
    0,
) // true
AreEqualMapsIntFloat64(
    map[int]float64{"1": 1., "2": 2.},
    map[int]float64{"1": 1.00001, "2": 2.},
    0,
) // false
AreEqualMapsIntFloat64(
    map[int]float64{"1": 1., "2": 2.},
    map[int]float64{"1": 1.00001, "2": 2.},
    .001,
) // true
```

## Return a unique slice

These functions are slightly atypical for the package, since, unlike the previous functions, they do not check anything. They are here because many users may wish to check if a slice is unique (which they can do using [these functions](#check-if-a-slice-is-unique)), and if not, to create a slice with its unique elements. To facilitate this task, the `check` package offers you handy functions  for `[]string`, `[]int` and `[]float64` slices.

#### `[]string`

```go
UniqueStringSlice([]string{"a", "b"}) // [a b]
UniqueStringSlice([]string{"a", "b", "b", "a", "b", "c"}) // [a b c]
UniqueIntSlice([]string{"a", "a"}) // [a]
```

#### `[]int`

```go
UniqueIntSlice([]int{1, 2}) // [1 2]
UniqueIntSlice([]int{1, 1}) // [1]
```

#### `[]float64`

Also here you need to remember about providing the `Epsilon` argument!

```go
UniqueFloat64Slice([]float64{.01, .01}, 0) // [.01 .01]
UniqueFloat64Slice([]float64{.01, .01001}, 0) // [.01 .01001]
UniqueFloat64Slice([]float64{.01, .01001}, .000000001) // [.01]
```
