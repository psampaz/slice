package slice

import "errors"

// DeleteBool removes the element at i index of a bool slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteBool(a []bool, i int) ([]bool, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteByte removes the element at i index of a byte slice. 
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteByte(a []byte, i int) ([]byte, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteComplex128 removes the element at i index of a complex128 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteComplex128(a []complex128, i int) ([]complex128, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteComplex64 removes the element at i index of a complex64 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteComplex64(a []complex64, i int) ([]complex64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteFloat32 removes the element at i index of a float32 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteFloat32(a []float32, i int) ([]float32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteFloat64 removes the element at i index of a float64 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteFloat64(a []float64, i int) ([]float64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteInt removes the element at i index of an int slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteInt(a []int, i int) ([]int, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteInt16 removes the element at i index of an int16 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteInt16(a []int16, i int) ([]int16, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteInt32 removes the element at i index of an int32 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteInt32(a []int32, i int) ([]int32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteInt64 removes the element at i index of an int64 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteInt64(a []int64, i int) ([]int64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteInt8 removes the element at i index of an int8 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteInt8(a []int8, i int) ([]int8, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteRune removes the element at i index of a rune slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteRune(a []rune, i int) ([]rune, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteString removes the element at i index of a string slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteString(a []string, i int) ([]string, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUint removes the element at i index of a uint slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUint(a []uint, i int) ([]uint, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUint16 removes the element at i index of a uint16 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUint16(a []uint16, i int) ([]uint16, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUint32 removes the element at i index of a uint32 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUint32(a []uint32, i int) ([]uint32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUint64 removes the element at i index of a uint64 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUint64(a []uint64, i int) ([]uint64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUint8 removes the element at i index of a uint8 slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUint8(a []uint8, i int) ([]uint8, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}

// DeleteUintptr removes the element at i index of a uintptr slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteUintptr(a []uintptr, i int) ([]uintptr, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(a)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(a[:i], a[i+1:]...), nil
}
