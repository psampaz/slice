package slice

import "errors"

// MinByte returns the minimum element of a byte slice
func MinByte(a []byte) (byte, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinFloat32 returns the minimum element of a float32 slice
func MinFloat32(a []float32) (float32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinFloat64 returns the minimum element of a float64 slice
func MinFloat64(a []float64) (float64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt returns the minimum element of an int slice
func MinInt(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt8 returns the minimum element of an int8 slice
func MinInt8(a []int8) (int8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt16 returns the minimum element of an int16 slice
func MinInt16(a []int16) (int16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt32 returns the minimum element of an int32 slice
func MinInt32(a []int32) (int32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt64 returns the minimum element of an int64 slice
func MinInt64(a []int64) (int64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinRune returns the minimum element of a rune slice
func MinRune(a []rune) (rune, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint returns the minimum element of a uint slice
func MinUint(a []uint) (uint, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint8 returns the minimum element of a uint8 slice
func MinUint8(a []uint8) (uint8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint16 returns the minimum element of a uint16 slice
func MinUint16(a []uint16) (uint16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint32 returns the minimum element of a uint32 slice
func MinUint32(a []uint32) (uint32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint64 returns the minimum element of a uint64 slice
func MinUint64(a []uint64) (uint64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUintptr returns the minimum element of a uintptr slice
func MinUintptr(a []uintptr) (uintptr, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the min of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}
