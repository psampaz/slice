package slice

import "errors"

// MaxByte returns the maximum value of a byte slice or an error in case of a nil or empty slice
func MaxByte(a []byte) (byte, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxFloat32 returns the maximum value of a float32 slice or an error in case of a nil or empty slice
func MaxFloat32(a []float32) (float32, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxFloat64 returns the maximum value of a float64 slice or an error in case of a nil or empty slice
func MaxFloat64(a []float64) (float64, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxInt returns the maximum value of a int slice or an error in case of a nil or empty slice
func MaxInt(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxInt8 returns the maximum value of a int8 slice or an error in case of a nil or empty slice
func MaxInt8(a []int8) (int8, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxInt16 returns the maximum value of a int16 slice or an error in case of a nil or empty slice
func MaxInt16(a []int16) (int16, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxInt32 returns the maximum value of a int32 slice or an error in case of a nil or empty slice
func MaxInt32(a []int32) (int32, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxInt64 returns the maximum value of a int64 slice or an error in case of a nil or empty slice
func MaxInt64(a []int64) (int64, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxRune returns the maximum value of a rune slice or an error in case of a nil or empty slice
func MaxRune(a []rune) (rune, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUint returns the maximum value of a uint slice or an error in case of a nil or empty slice
func MaxUint(a []uint) (uint, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUint8 returns the maximum value of a uint8 slice or an error in case of a nil or empty slice
func MaxUint8(a []uint8) (uint8, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUint16 returns the maximum value of a uint16 slice or an error in case of a nil or empty slice
func MaxUint16(a []uint16) (uint16, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUint32 returns the maximum value of a uint32 slice or an error in case of a nil or empty slice
func MaxUint32(a []uint32) (uint32, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUint64 returns the maximum value of a uint64 slice or an error in case of a nil or empty slice
func MaxUint64(a []uint64) (uint64, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}

// MaxUintptr returns the maximum value of a uintptr slice or an error in case of a nil or empty slice
func MaxUintptr(a []uintptr) (uintptr, error) {
	if len(a) == 0 {
		return 0, errors.New("returns the maximum of a nil or empty slice")
	}

	max := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] > max {
			max = a[k]
		}
	}

	return max, nil
}
