package slice

import "errors"

// PopBool removes and returns the last value a bool slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopBool(a []bool) (bool, []bool, error) {
	if len(a) == 0 {
		return false, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopByte removes and returns the last value a byte slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopByte(a []byte) (byte, []byte, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopComplex128 removes and returns the last value a complex128 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopComplex128(a []complex128) (complex128, []complex128, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopComplex64 removes and returns the last value a complex64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopComplex64(a []complex64) (complex64, []complex64, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopFloat32 removes and returns the last value a float32 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopFloat32(a []float32) (float32, []float32, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopFloat64 removes and returns the last value a float64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopFloat64(a []float64) (float64, []float64, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopInt removes and returns the last value a int slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopInt(a []int) (int, []int, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopInt16 removes and returns the last value a int16 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopInt16(a []int16) (int16, []int16, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopInt32 removes and returns the last value a int32 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopInt32(a []int32) (int32, []int32, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopInt64 removes and returns the last value a int64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopInt64(a []int64) (int64, []int64, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopInt8 removes and returns the last value a int8 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopInt8(a []int8) (int8, []int8, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopRune removes and returns the last value a rune slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopRune(a []rune) (rune, []rune, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopString removes and returns the last value a string slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopString(a []string) (string, []string, error) {
	if len(a) == 0 {
		return "", nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUint removes and returns the last value a uint slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUint(a []uint) (uint, []uint, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUint16 removes and returns the last value a uint16 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUint16(a []uint16) (uint16, []uint16, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUint32 removes and returns the last value a uint32 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUint32(a []uint32) (uint32, []uint32, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUint64 removes and returns the last value a uint64 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUint64(a []uint64) (uint64, []uint64, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUint8 removes and returns the last value a uint8 slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUint8(a []uint8) (uint8, []uint8, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}

// PopUintptr removes and returns the last value a uintptr slice and the remaining slice.
// An error is returned in case of a nil or empty slice.
func PopUintptr(a []uintptr) (uintptr, []uintptr, error) {
	if len(a) == 0 {
		return 0, nil, errors.New("Cannot pop from a nil or empty slice")
	}

	return a[len(a)-1], a[:len(a)-1], nil
}
