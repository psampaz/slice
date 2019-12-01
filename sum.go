package slice

import "errors"

// SumByte returns the sum of the values of a byte Slice or an error in case of a nil or empty slice
func SumByte(a []byte) (byte, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum byte
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumComplex128 returns the sum of the values of a complex128 Slice or an error in case of a nil or empty slice
func SumComplex128(a []complex128) (complex128, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum complex128
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumComplex64 returns the sum of the values of a complex64 Slice or an error in case of a nil or empty slice
func SumComplex64(a []complex64) (complex64, error) {
	if len(a) == 0 {
		return complex(0, 0), errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum complex64
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumFloat32 returns the sum of the values of a float32 Slice or an error in case of a nil or empty slice
func SumFloat32(a []float32) (float32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum float32
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumFloat64 returns the sum of the values of a float64 Slice or an error in case of a nil or empty slice
func SumFloat64(a []float64) (float64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum float64
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumInt returns the sum of the values of a int Slice or an error in case of a nil or empty slice
func SumInt(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum int
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumInt16 returns the sum of the values of a int16 Slice or an error in case of a nil or empty slice
func SumInt16(a []int16) (int16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum int16
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumInt32 returns the sum of the values of a int32 Slice or an error in case of a nil or empty slice
func SumInt32(a []int32) (int32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum int32
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumInt64 returns the sum of the values of a int64 Slice or an error in case of a nil or empty slice
func SumInt64(a []int64) (int64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum int64
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumInt8 returns the sum of the values of a int8 Slice or an error in case of a nil or empty slice
func SumInt8(a []int8) (int8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum int8
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumRune returns the sum of the values of a rune Slice or an error in case of a nil or empty slice
func SumRune(a []rune) (rune, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum rune
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUint returns the sum of the values of a uint Slice or an error in case of a nil or empty slice
func SumUint(a []uint) (uint, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uint
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUint16 returns the sum of the values of a uint16 Slice or an error in case of a nil or empty slice
func SumUint16(a []uint16) (uint16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uint16
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUint32 returns the sum of the values of a uint32 Slice or an error in case of a nil or empty slice
func SumUint32(a []uint32) (uint32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uint32
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUint64 returns the sum of the values of a uint64 Slice or an error in case of a nil or empty slice
func SumUint64(a []uint64) (uint64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uint64
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUint8 returns the sum of the values of a uint8 Slice or an error in case of a nil or empty slice
func SumUint8(a []uint8) (uint8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uint8
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}

// SumUintptr returns the sum of the values of a uintptr Slice or an error in case of a nil or empty slice
func SumUintptr(a []uintptr) (uintptr, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum uintptr
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}
