# `check` Go package

The `check` package of Go helps one to check various conditions for

* slices:
    * `[]int`
    * `[]float64`
    * `[]string`
    * `[]bool`
* maps:
    * `map[string]int`
    * `map[string]string`
    * `map[string]float64`
    * `map[string]bool`
    * `map[int]int`
    * `map[int]string`
    * `map[int]float64`
    * `map[int]bool`

Such checks can be useful in regular code but also in testing. For instance, you can use the `check` package to compare obtained and expected slices (or maps) obtained. You will find many testing examples in the package's code in its [repository](https://github.com/nyggus/check)).

You can check the following types of conditions:

* check if all boolean values are `true` (in a slice and in a map)
* check if any boolean value is true (in a slice and in a map)
* check which of boolean values is true (in a slice and in a map)
* check if all elements of a slice have the same value
* check if all keys of a map have the same value
* check if two slices are the same (either taking into account the ordering of the slices or ignoring it)
* check if two maps are the same
* check if a value is in a slice
* check if several values are in a slice
* check if any of several values are in a slice
* check which of several values are in a slice
* check if a key-value pair is among a map's key-values
* check if a value is among a map's values
* check if several values from a slice are in a map
* check if any value from a slice are in a map
* check which values from a slice are in a map
* check if all key-value pairs from a map are in a map
* check if any of key-value pairs from a map are in a map
* check which key-value pairs from a map are in a map

In addition to these checks, the package enables one to create a unique slice (that is, with unique values) out of a slice. 

>> **Working with `float64` values**: All comparisons of `float64` values enable the user to use an epsilon. This means that two floats differ when their absolute difference is higher than the epsilon. If you do not want to use the epsilon, simply make it 0. **Always, always pay special attention to working with floats, and make sure what you're doing is what you need.**

Many functions resemble one another, but yet there are no generic functions (with the exception of `IsUnique`). See [here](#why-no-generics?) for explanation.

# Naming convention

Most function names are long, on purpose: this is to help the user fit the function to his or her needs, without the need to remember too many details. Below are some general rules. Note that `...` represent details about a slice or a map; e.g., `AnyInMap...` can be `WhichInMapString` or `WhichInMapInt` (which are short versions of `WhichInMapStringBool` or `WhichInMapIntBool`).

* `Any()` and `All()`, the only short names, work with boolean slices (`[]bool`)
* `AnyInMap...` and `AllInMap...` and `WhichInMap...` functions check if any or all values of a map is/are true (work with `map[int]bool` and `map[string]bool`), or which are
* `IsUnique...Slice` checks whether a slice of particular type has unique values (`IsUniqueIntSlice`, `IsUniqueStringSlice` and `IsUniqueFloat64Slice`); you can use `IsUnique` instead, a generic function working with `[]int`, `[]string` and `[]float64` slices, but not without decrease in performance
* `Unique...Slice` returns a new slice of the same type, with unique values of the original slice (so duplicates are removed)
* `IsUniqueMap...` checks whether a map has unique values; note that the map's keys are ignored (since each map has unique key-value pairs)
* `AreEqualSlices...` compares two slices of type `...` (again, `int`, `string` or `float64`); the functions compares both values and the slices' ordering
* `AreEqualSortedSlices...` compares two slices of type `...`; unlike the above functions, these ignore ordering
* `IsValueIn...Slice` checks if a slice has a particular value
* `AnyValueIn...Slice` checks if any of values provided as a slice is in another slice
* `AllValuesIn...Slice` checks if all values provided as a slice are in another slice
* `WhichValuesIn...Slice` checks which values provided as a slice are in another slice
* `AreEqualMaps...` compares whether two maps contain the same key-value pairs
* `IsValueInMap...` checks whether a map contains a particular value; here, `...` can be `StringString`, `IntFloat64` and the like (see above the types of maps that the `check` package works with)
* `AnyValueInMap...` checks whether any of values provided as a slice are among a map's values
* `AllValuesInMap...` checks whether all values provided as a slice are among a map's values
* `WhichValuesInMap...` checks which values provided as a slice are among a map's values
* `AnyKeyValuePairInMap...` checks whether a map contains any of the key-value pairs provided as a map
* `AllKeyValuePairsInMap...` checks whether a map contains all key-value pairs provided as a map
* `WhichKeyValuePairsInMap...` checks which key-value pairs provided as a map are in a map

As you see, these functions offer many various types of checks, hence their names were designed in such a way so that a function's name facilitates remembering what the function does. The consequence is long names, but this cost comes with readability.

Most functions starting with `Is`, `Any` and `All` return a boolean value. The only exceptions are functions starting with `IsValueInMap`, since they also return the keys for which this value occurs in the map.

# Examples

The above section shows the types of functions you will find in `check` and their general purposes. You will find many simple examples of using the package's functions in the package's documentation and in the docs folder of this repository, in [this file](https://github.com/nyggus/check/blob/master/docs/examples.md). Below, you will find just several examples that aim to shed some light on what the package offers.

### I want to check if two integer slices are the same

What do you mean? I mean, you can compare their values and ignore whether or not the slices are ordered in the same way, or you can take their ordering into account. Let's see:

```go
slice1 := []int{1, 1, 1, 2, 1, 1, 2}
slice2 := []int{1, 1, 2, 1, 1, 1, 2}
```

Thanks to the `check` package, we can check if the slices have the same values, using `AreEqualSlicesInt()` function:

```
AreEqualSortedSlicesInt(slice1, slice2) // true
```

As the name suggests, the slices are first sorted and only then compared; hence they're the same. If we want take into account that their orderings differ, then they should not differ, and they don't:

```
AreEqualSlicesInt(slice1, slice2) // false
```

### I want to check if a slice contains a particular value

Imagine you have the following slice:

```go
slice := []string{"This", "is", "a", "string", "slice", "."}
```

Does `slice` contain a word "is"? Let's check:

```go
IsValueInStringSlice("is", slice) // true
```

Yes, it does! And does it contain "Alice" and "string" and "thing"?

```go
AllValuesInStringSlice([]string{"Alice", "string", "thing"}, slice) // false
```

No, it does not contain all of them; but perhaps it contains any of them?

```go
AnyValueInStringSlice([]string{"Alice", "string", "thing"}, slice) // true
```

Yes, this time we do have `true`: this slice does contain at least one from among those three strings. Simple, but we do not know which of them - and even how many of them - are in the slice. `check` offers you a function to find this:

```go
values, ok := WhichValuesInStringSlice([]string{"Alice", "string", "thing"}, slice)
// ok is true
// values is map[1:string]
```

As you see, the `Which...` functions return two values: a boolean (which is actually the second returned value) that says whether any value from `slice1` is in `slice2`, and a map providing those values. In `map[1:string]`, the key (`1`) represents the index of the value (`"string"`) in the first slice. Consider another example of a `Which...` function:

```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 1, 2, 1, 1, 1, 2, 7, 33, 12, 33, 67, 90}
values, ok := WhichValuesInIntSlice(slice1, slice2)
// ok is true
// map[1:[0 1 3 4 5] 2:[2 6]]
```

### I want to check if a `float64` number is among a map's values

As mentioned above, be very careful when working with float numbers. To help you do this, all the functions that compare floats have the `Epsilon` parameter, which aims to achieve the desired level of accuracy. Of course, another approach could be to round all the values before any comparisons, though not always you will want to round them. Here's an example of how to check if a `float64` number is among a map's values.

```go
searchedFloat := .023
Map := map[int]float64{1:.0214, 2:.0212, 3:.0222, 4:.0231}
IsValueInMapIntFloat64(searchedFloat, Map, 0) // false
```

Here, we used `Epsilon = 0`, so we did not decrease the accuracy: what we have is what we got, and in that case, `.023` is _not_ among `Map`'s values. But when we change this level of accuracy, increasing it to, say, .001 (which means that if two values differ by .001 or less than they are considered equal), then 

```go
IsValueInMapIntFloat64(searchedFloat, Map, .001) // true
```

This is becuase `.0231 - .023` is equal to `Epsilon` (`.001`), so the searched float is in fact among the map's values.

### I want to check if two maps are the same

Let's compare if two maps are the same:

```go
Map1 := map[int]string{1: "What", 2: "a", 3: "mysterious", 4: "thing"}
Map2 := map[int]string{4: "thing", 1: "What", 2: "a", 3: "mysterious"}
Map1 := map[int]string{1: "What", 2: "a", 3: "mysterious", 4: "thing!"}

AreEqualMapsIntString(Map1, Map2) // true
AreEqualMapsIntString(Map1, Map3) // false
```

Of course, `Map1` and `Map2` are the same, since the ordering of a map is random and does not matter. `Map1` and `Map2`, however, differ from `Map3` because of the exclamation mark in `"thing!"` of the key `4` of the latter.

### I want to check is a slice is unique

You can do so for `[]int`, `[]string` and `[]float64` slices, e.g.,

```go
IsUniqueIntSlice([]int{1, 1, 2, 12, 3, 21, 71}) // false
IsUniqueIntSlice([]int{1, 15, 2, 12, 3, 21, 71}) // true
IsUniqueFloat64Slice([]float64{.0214, .0212, .0222, .0221}, .001) // false
```

Note that in the latter case, `Epsilon = .001`, so `.0222` and `.0221` are considered equal (their difference is equal to Epsilon). If you increase the accuracy (that is, decrease `Epsilon`), the result changes:

```go
IsUniqueFloat64Slice([]float64{.0214, .0212, .0222, .0221}, .0001) // true
```

### I want to create a unique slice

As already mentioned, you can create a unique slice (also `[]float64`, for which you will again have to use the `Epsilon` parameter), even if this does not seem to fit the `check` package's main purpose. You can do it in the following way:

```go
UniqueIntSlice([]int{1, 1, 2, 2, 3, 21, 1}) // [1 2 3 21]
UniqueStringSlice([]string{"a", "a", "C", "b"}) // [a C b]
UniqueFloat64Slice([]float64{.0021, .0024, .0022, .0031, .00311}, 0) // [.0021, .0024, .0022, .0031, .00311]
UniqueFloat64Slice([]float64{.0021, .0024, .0022, .0031, .00311}, .0001) // [.0021, .0024, .0022, .0031]
```

# Why bother? I can make all those checks directly in my code!

Sure you can! And frankly, this is not that difficult to do - but you need to be careful. The `check` package, however, offers you an alternative: a one-liner that will say what you're doing instead of five or ten additional lines, and you need not worry about the details. It's like with any package: It can help you save time and energy and work, and it's tested, and so you can use it without worrying that you have made a mistake somewhere there in the code or omitted something important.

# Why no generics?

Many of the `check` functions are similar, so it might be tempting to propose several generic functions. The `check` package, however, offers only one such generic function, `IsUnique` (working with  `[]int`, `[]string` and `[]float64` slices), and it does so only for the representation purpose.

One reason behind no generics is that the functions working with `float64` values use the `Epsilon` parameter, which is not used by functions working with `string` and `int` values. Hence to make a generic function to work with different signatures would mean either complicating or simplifying things, something better to be avoided. For instance, `IsUnique` sets `Epsilon` to 0, a simplification not always preferred.

Another issue could be performance. Generics would require type checking, which could slightly decrease the performance. Nonetheless, this effect is so small that can be considered negligible, unless for very small slices. You can find some benchmarks in [this file](https://github.com/nyggus/check/blob/master/benchmarks_generics_test.go).

Let's wait for the generics to appear in Go soon, and then we'll see.

# Contribution

You will find the package's code to the package in its [github repository](https://github.com/nyggus/check). Feel free to contribute if you see any bugs or have an idea for improving the package or adding a new functionality. You can either create an issue or submit a merge request.
