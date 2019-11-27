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

|            | Min | Max | Contains | Reverse | Shuffle | Clone | Deduplicate | Insert | Delete |
| ---------- | --- | --- | -------- | ------- | ------- | ----- | ----------- | ------ | ------ |
| bool       | ✕   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| byte       | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| complex64  | ✕   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| complex128 | ✕   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| float32    | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| float64    | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| int        | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| int8       | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| int16      | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| int32      | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| int64      | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| rune       | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| string     | ✕   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uint       | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uint8      | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uint16     | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uint32     | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uint64     | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |
| uintptr    | ✔   | -   | ✔        | -       | -       | -     | -           | -      | -      |

# Examples

## Min

Min returns the minimum value of a slice or an error in case of a nil or empty slice.
```go
    a := []int{1, 2, 3, 0, 4, 5}
    min, err := slice.MinInt(a)
```

Contains checks if a specific value exists in a slice.
```go
    a := []int{"a","b","c","d"}
    b := "c"
    exists := slice.Contains(a, "c")
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
