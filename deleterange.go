package slice

import "errors"

// DeleteRangeBool deletes the elements between from and to index (inclusive) from a bool slice
func DeleteRangeBool(a []bool, from, to int) ([]bool, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeByte deletes the elements between from and to index (inclusive) from a byte slice
func DeleteRangeByte(a []byte, from, to int) ([]byte, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeComplex128 deletes the elements between from and to index (inclusive) from a complex128 slice
func DeleteRangeComplex128(a []complex128, from, to int) ([]complex128, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeComplex64 deletes the elements between from and to index (inclusive) from a complex64 slice
func DeleteRangeComplex64(a []complex64, from, to int) ([]complex64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeFloat32 deletes the elements between from and to index (inclusive) from a float32 slice
func DeleteRangeFloat32(a []float32, from, to int) ([]float32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeFloat64 deletes the elements between from and to index (inclusive) from a float64 slice
func DeleteRangeFloat64(a []float64, from, to int) ([]float64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeInt deletes the elements between from and to index (inclusive) from a int slice
func DeleteRangeInt(a []int, from, to int) ([]int, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeInt16 deletes the elements between from and to index (inclusive) from a int16 slice
func DeleteRangeInt16(a []int16, from, to int) ([]int16, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeInt32 deletes the elements between from and to index (inclusive) from a int32 slice
func DeleteRangeInt32(a []int32, from, to int) ([]int32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeInt64 deletes the elements between from and to index (inclusive) from a int64 slice
func DeleteRangeInt64(a []int64, from, to int) ([]int64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeInt8 deletes the elements between from and to index (inclusive) from a int8 slice
func DeleteRangeInt8(a []int8, from, to int) ([]int8, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeRune deletes the elements between from and to index (inclusive) from a rune slice
func DeleteRangeRune(a []rune, from, to int) ([]rune, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeString deletes the elements between from and to index (inclusive) from a string slice
func DeleteRangeString(a []string, from, to int) ([]string, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUint deletes the elements between from and to index (inclusive) from a uint slice
func DeleteRangeUint(a []uint, from, to int) ([]uint, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUint16 deletes the elements between from and to index (inclusive) from a uint16 slice
func DeleteRangeUint16(a []uint16, from, to int) ([]uint16, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUint32 deletes the elements between from and to index (inclusive) from a uint32 slice
func DeleteRangeUint32(a []uint32, from, to int) ([]uint32, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUint64 deletes the elements between from and to index (inclusive) from a uint64 slice
func DeleteRangeUint64(a []uint64, from, to int) ([]uint64, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUint8 deletes the elements between from and to index (inclusive) from a uint8 slice
func DeleteRangeUint8(a []uint8, from, to int) ([]uint8, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}

// DeleteRangeUintptr deletes the elements between from and to index (inclusive) from a uintptr slice
func DeleteRangeUintptr(a []uintptr, from, to int) ([]uintptr, error) {
	if len(a) == 0 {
		return nil, errors.New("Cannot delete from a nil or empty slice")
	}

	if from < 0 || from > len(a)-1 {
		return nil, errors.New("from index out of bounds")
	}

	if to < 0 || to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}
