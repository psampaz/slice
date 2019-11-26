![Build Status](https://github.com/psampaz/slice/workflows/build/badge.svg)
[![codecov](https://codecov.io/gh/psampaz/slice/branch/master/graph/badge.svg)](https://codecov.io/gh/psampaz/slice)
[![GoDoc](https://godoc.org/github.com/psampaz/slice?status.svg)](https://godoc.org/github.com/psampaz/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/psampaz/slice)](https://goreportcard.com/report/github.com/psampaz/slice)
[![golangci](https://golangci.com/badges/github.com/psampaz/slice.svg)](https://golangci.com/r/github.com/psampaz/slice)

Type-safe functions for common Go slice operations.


# Operations 

|            | Min |
| ---------- | --- |
| bool       | ✕   |
| byte       | ✔   |
| complex64  | ✕   |
| complex128 | ✕   |
| float32    | ✔   |
| float64    | ✔   |
| int        | ✔   |
| int8       | ✔   |
| int16      | ✔   |
| int32      | ✔   |
| int64      | ✔   |
| rune       | ✔   |
| string     | ✕   |
| uint       | ✔   |
| uint8      | ✔   |
| uint16     | ✔   |
| uint32     | ✔   |
| uint64     | ✔   |
| uintptr    | ✔   |

# Examples

## Min

Min returns the minimum value of a slice or an error in case of a nil or empty slice.
```go
    a := []int{1, 2, 3, 0, 4, 5}
    min, err := slice.MinInt(a)
```

# Contributing

You are very welcome to contribute new opearations in this library.

The minimum requirements to accept a pull are the following:

- 100% code coverage
- Pass golangci checks
- Precise documentation of each function
- If the PR changes an existing operation
    - apply changes to all functions of this operation
- If the PR introduces a new operation, then:
    - implement functions for all applicable types
    - include one testable example
    - update the table on the Operations section
    - add one example int he Examples section

