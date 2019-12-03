![Build Status](https://github.com/psampaz/slice/workflows/build/badge.svg)
[![codecov](https://codecov.io/gh/psampaz/slice/branch/master/graph/badge.svg)](https://codecov.io/gh/psampaz/slice)
[![GoDoc](https://godoc.org/github.com/psampaz/slice?status.svg)](https://godoc.org/github.com/psampaz/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/psampaz/slice)](https://goreportcard.com/report/github.com/psampaz/slice)
[![golangci](https://golangci.com/badges/github.com/psampaz/slice.svg)](https://golangci.com/r/github.com/psampaz/slice)

Type-safe functions for common Go slice operations.


# Operations 

✔ = Supported 

✕ = Non supported 

\- = Not yet implemented

|            | bool | byte | complex(all) | float(all) | int(all) | string | uint(all) | uintptr |
| ---------- | ---- | ---- | ------------ | ---------- | -------- | ------ | --------- | ------- |
| Batch      | -    | -    | -            | -          | -        | -      | -         | -       | 
| Contains   | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Copy       | -    | -    | -            | -          | -        | -      | -         | -       | 
| Cut        | -    | -    | -            | -          | -        | -      | -         | -       | 
| Deduplicate| ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Delete     | -    | -    | -            | -          | -        | -      | -         | -       | 
| Filter     | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       |
| Insert     | -    | -    | -            | -          | -        | -      | -         | -       | 
| Max        | ✕    | ✔    | ✕            | ✔          | ✔        | ✕      | ✔         | ✔       |
| Min        | ✕    | ✔    | ✕            | ✔          | ✔        | ✕      | ✔         | ✔       |
| Pop        | -    | -    | -            | -          | -        | -      | -         | -       |
| Push       | -    | -    | -            | -          | -        | -      | -         | -       |
| Reverse    | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Shift      | -    | -    | -            | -          | -        | -      | -         | -       | 
| Shuffle    | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Sum        | -    | -    | -            | -          | -        | -      | -         | -       | 
| Unshift    | -    | -    | -            | -          | -        | -      | -         | -       | 

# Examples

## slice.Deduplicate

Deduplicate performs order preserving, in place deduplication of a slice
```go
    a := []int{1, 2, 3, 2, 5, 3}
    a = slice.DeduplicateInt(a) // [1, 2, 3, 5]
```

## slice.Contains

Contains checks if a specific value exists in a slice.
```go
    a := []int{"a","b","c","d"}
    exists := slice.Contains(a, "c") // true
```

## slice.Filter

Filter performs in place filtering of a slice based on a predicate
```go
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    keep := func(x int) bool {
        return x%2 == 0 
    }
    a = slice.Filter(a, keep) // [2, 4, 6, 8, 10]
```

## slice.Max

Max returns the maximum value of a slice or an error in case of a nil or empty slice.
```go
    a := []int{1, 2, 3, 0, 4, 5}
    max, err := slice.MaxInt(a) // 5, nil
```

## slice.Min

Min returns the minimum value of a slice or an error in case of a nil or empty slice.
```go
    a := []int{1, 2, 3, 0, 4, 5}
    min, err := slice.MinInt(a) // 0, nil
```

## slice.Reverse

Reverse performs in place reversal of a slice
```go
    a := []int{1, 2, 3, 4, 5}
    a = slice.ReverseInt(a) // [5, 4, 3, 2, 1]
```


## slice.Shuffle

Shuffle shuffles (in place) a slice
```go
    a := []int{1, 2, 3, 4, 5}
    a = slice.Shuffle(a)
```

# Contributing

You are very welcome to contribute new opearations in this library.

The minimum requirements to accept a pull request are the following:

- Use only functions. This is a function based library so struct based operations will not be accepted, in order to preserve simplicity and consistency 
- 100% code coverage
- Pass golangci checks
- Precise documentation of each function
- If the PR changes an existing operation
    - apply changes to all functions of this operation
- If the PR introduces a new operation, then:
    - implement functions for all applicable types
    - include one testable example at the end of the test file
    - update the table in the Operations section
    - add one example in the Examples section
    - use a separate file for each operation
    - use the type after the name of the function: use MinInt32  instead of Int32Min

# Test

if you want to run the test suite for this library:

```go
$ go test -v -cover
```

# Credits

- SliceTricks (https://github.com/golang/go/wiki/SliceTricks). This was the inspiration behind this library.
- Genny (https://github.com/cheekybits/genny). In order to speedup the development and avoid massive copy paste, the excellent Genny library was used.
