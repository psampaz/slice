![Build Status](https://github.com/psampaz/slice/workflows/build/badge.svg)
[![codecov](https://codecov.io/gh/psampaz/slice/branch/master/graph/badge.svg)](https://codecov.io/gh/psampaz/slice)
[![GoDoc](https://godoc.org/github.com/psampaz/slice?status.svg)](https://godoc.org/github.com/psampaz/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/psampaz/slice)](https://goreportcard.com/report/github.com/psampaz/slice)
[![golangci](https://golangci.com/badges/github.com/psampaz/slice.svg)](https://golangci.com/r/github.com/psampaz/slice)

Type-safe functions for common Go slice operations.

# Installation

```
go get github.com/psampaz/slice
```

# Operations 

✔ = Supported 

✕ = Non supported 

\- = Not yet implemented

|            | bool | byte | complex(all) | float(all) | int(all) | string | uint(all) | uintptr |
| ---------- | ---- | ---- | ------------ | ---------- | -------- | ------ | --------- | ------- |
| Batch      | -    | -    | -            | -          | -        | -      | -         | -       | 
| Contains   | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Copy       | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Cut        | -    | -    | -            | -          | -        | -      | -         | -       | 
| Deduplicate| ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Delete     | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Filter     | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       |
| Insert     | -    | -    | -            | -          | -        | -      | -         | -       | 
| Max        | ✕    | ✔    | ✕            | ✔          | ✔        | ✕      | ✔         | ✔       |
| Min        | ✕    | ✔    | ✕            | ✔          | ✔        | ✕      | ✔         | ✔       |
| Pop        | -    | -    | -            | -          | -        | -      | -         | -       |
| Push       | -    | -    | -            | -          | -        | -      | -         | -       |
| Reverse    | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Shift      | -    | -    | -            | -          | -        | -      | -         | -       | 
| Shuffle    | ✔    | ✔    | ✔            | ✔          | ✔        | ✔      | ✔         | ✔       | 
| Sum        | ✕    | ✔    | ✔            | ✔          | ✔        | ✕      | ✔         | ✔       | 
| Unshift    | -    | -    | -            | -          | -        | -      | -         | -       | 

# Examples

## slice.Deduplicate

Deduplicate performs order preserving, in place deduplication of a slice
```go
    a := []int{1, 2, 3, 2, 5, 3}
    a = slice.DeduplicateInt(a) // [1, 2, 3, 5]
```

## slice.Delete

Delete removes an element at a specific index of a slice. An error is return in case the index is out of bounds or the slice is nil or empty.
```go
    a := []int{1, 2, 3, 4, 5}
    a, err = slice.DeleteInt(a, 2) // [1, 2, 4, 5], nil
```

## slice.Contains

Contains checks if a specific value exists in a slice.
```go
    a := []int{1, 2, 3, 4, 5}
    exists := slice.ContainsInt(a, 3) // true
```

## slice.Copy

Copy creates a copy of a slice.
The resulting slice has the same elements as the original but the underlying array is different.
See https://github.com/go101/go101/wiki
```go
    a := []int{1, 2, 3, 4}
    b := slice.CopyInt(a) // [1, 2, 3, 4]
```

## slice.Filter

Filter performs in place filtering of a slice based on a predicate
```go
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    keep := func(x int) bool {
        return x%2 == 0 
    }
    a = slice.FilterInt(a, keep) // [2, 4, 6, 8, 10]
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
    a = slice.ShuffleInt(a) // [3, 5, 1, 4, 2] (random output)
 ```

## slice.Sum

Sum returns the sum of the values of a slice or an error in case of a nil or empty slice

```go
    a := []int{1, 2, 3}
    sum, err := slice.SumInt(a) // 6, nil
```

# Tests

if you want to run the test suite for this library:

```go
$ go test -v -cover
```

# Credits

- SliceTricks (https://github.com/golang/go/wiki/SliceTricks). This was the inspiration behind this library.
- Genny (https://github.com/cheekybits/genny). In order to speedup the development and avoid massive copy paste, the excellent Genny library was used.

# Contributing

You are very welcome to contribute new operations or bug fixes in this library.

## Contribution guidelines (code)

1. Use only functions. This is a function based library so struct based operations will not be accepted, in order to preserve simplicity and consistency.

2. If the operation is not working on a nil or empty slice, then the function should return an error.

3. If the operation accepts slice indexes as parameters, then the function should guard against out of bound index values and return an error in that case.

4. All operations should be in place operations, meaning that they should alter the original slice. 

5. Each function should have precise documentation. 

6. Each operation should live in each own file. Example:

    ```
    min.go
    min_test.go
    ```
7. The naming convention for functions is OperationType. Example:

    ``` 
    MinInt32()
    ```

    instead of 

    ``` 
    Int32Min()
    ```
8. Implement ALL applicable types in the same PR.

9. Include one testable example for Int type at the end of the test file.

10. Include one example in the Examples section of README

11. Update the table in the Operation section of README 

12. Update the UNRELEASED section of CHANGELOG

## Contribution guidelines (tests)

1. All code should be 100% covered with tests
2. All operations should be tested for 3 scenarios at least:
    1. nil slice
    2. empty slice
    3. non empty slice

## Static code analysis    

golangci.com runs on all PRs. Code is checked with golint, go vet, gofmt, plus 20+ linters, and review comments will be automatically added in your PR in case of a failure. You can see the whole list of linters here: https://golangci.com/product#linters

## Steps for contributing new operations

1. Open an issue describing the new operation, the proposed name and the applicable types.
2. If the operation is approved to be included in the library, create a small PR the implementation and test for only only type. 
3. After code review you can proceed the implementation for the rest types. This is necessary because if you submit a PR with the implementation and test for all types, a small correction during review could eventually lead to a big refactor due to code duplication.

## Using Genny for fast implementation of all types of an operation (Optional)

The following steps are an example of how to use [https://github.com/cheekybits/genny](Genny) to implement the min operation:

1. Install Genny

    ```
    go get github.com/cheekybits/genny
    ```

2. Create a file named min_genny.go

    ```go
    package slice

    import (
        "errors"
        "github.com/cheekybits/genny/generic"
    )

    type Type generic.Type

    // MinType returns the minimum value of an Type slice or an error in case of a nil or empty slice
    func MinType(a []Type) (Type, error) {
        if len(a) == 0 {
            return 0, errors.New("Cannot get the minimum of a nil or empty slice")
        }

        min := a[0]
        for k := 1; k < len(a); k++ {
            if a[k] < min {
                min = a[k]
            }
        }

        return min, nil
    }
    ```

3. Use genny to generate code for all Go's built in types:

    ```
    cat min_genny.go | genny gen Type=BUILTINS > min.go
    ```

    This step will generate a file min.go with the following content:

    ```go
    package slice

    import "errors"

    // MinByte returns the minimum value of a byte slice or an error in case of a nil or empty slice
    func MinByte(a []byte) (byte, error) {
        if len(a) == 0 {
            return 0, errors.New("Cannot get the minimum of a nil or empty slice")
        }

        min := a[0]
        for k := 1; k < len(a); k++ {
            if a[k] < min {
                min = a[k]
            }
        }

        return min, nil
    }

    // MinFloat32 returns the minimum value of a float32 slice or an error in case of a nil or empty slice
    func MinFloat32(a []float32) (float32, error) {
        if len(a) == 0 {
            return 0, errors.New("Cannot get the minimum of a nil or empty slice")
        }

        min := a[0]
        for k := 1; k < len(a); k++ {
            if a[k] < min {
                min = a[k]
            }
        }

        return min, nil
    }
    .
    .
    .
    .
    ```

4. Delete the implementation for all types not applicable for the operation

5. Create a file named min_genny_test.go

    ```go
    package slice

    import (
        "fmt"
        "testing"
    )

    func TestMinType(t *testing.T) {
        type args struct {
            a []Type
        }
        tests := []struct {
            name    string
            args    args
            want    Type
            wantErr bool
        }{
            {
                name: "nil slice",
                args: args{
                    a: nil,
                },
                want:    0,
                wantErr: true,
            },
            {
                name: "empty slice",
                args: args{
                    a: []Type{},
                },
                want:    0,
                wantErr: true,
            },
            {
                name: "non empty slice",
                args: args{
                    a: []Type{1, 3, 2, 0, 5, 4},
                },
                want:    0,
                wantErr: false,
            },
        }
        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
                got, err := MinType(tt.args.a)
                if (err != nil) != tt.wantErr {
                    t.Errorf("MinType() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if got != tt.want {
                    t.Errorf("MinType() = %v, want %v", got, tt.want)
                }
            })
        }
    }

    ```

6. Use genny to generate tests for all Go's built in types:

    ```
    cat min_genny_test.go | genny gen Type=BUILTINS > min_test.go
    ```

    This step will generate a file min_test.go with tests for each one of Go's built in types.

7. Remove tests for non applicable types.

8. Adjust the tests for each one of the types.

9. Delete min_genny.go and min_genny_test.go

