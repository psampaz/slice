package slice

// ReverseBool performs in place reversal of a bool slice
func ReverseBool(a []bool) []bool {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseByte performs in place reversal of a byte slice
func ReverseByte(a []byte) []byte {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseComplex128 performs in place reversal of a complex128 slice
func ReverseComplex128(a []complex128) []complex128 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseComplex64 performs in place reversal of a complex64 slice
func ReverseComplex64(a []complex64) []complex64 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseFloat32 performs in place reversal of a float32 slice
func ReverseFloat32(a []float32) []float32 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseFloat64 performs in place reversal of a float64 slice
func ReverseFloat64(a []float64) []float64 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseInt performs in place reversal of an int slice
func ReverseInt(a []int) []int {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseInt16 performs in place reversal of an int16 slice
func ReverseInt16(a []int16) []int16 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseInt32 performs in place reversal of an int32 slice
func ReverseInt32(a []int32) []int32 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseInt64 performs in place reversal of an int64 slice
func ReverseInt64(a []int64) []int64 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseInt8 performs in place reversal of an int8 slice
func ReverseInt8(a []int8) []int8 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseRune performs in place reversal of a rune slice
func ReverseRune(a []rune) []rune {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseString performs in place reversal of a string slice
func ReverseString(a []string) []string {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUint performs in place reversal of a uint slice
func ReverseUint(a []uint) []uint {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUint16 performs in place reversal of a uint16 slice
func ReverseUint16(a []uint16) []uint16 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUint32 performs in place reversal of a uint32 slice
func ReverseUint32(a []uint32) []uint32 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUint64 performs in place reversal of a uint64 slice
func ReverseUint64(a []uint64) []uint64 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUint8 performs in place reversal of a uint8 slice
func ReverseUint8(a []uint8) []uint8 {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}

// ReverseUintptr performs in place reversal of a uintptr slice
func ReverseUintptr(a []uintptr) []uintptr {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}
