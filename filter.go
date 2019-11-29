package slice

// FilterBool performs in place filtering of a bool slice based on a predicate
func FilterBool(a []bool, keep func(x bool) bool) []bool {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterByte performs in place filtering of a byte slice based on a predicate
func FilterByte(a []byte, keep func(x byte) bool) []byte {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterComplex128 performs in place filtering of a complex128 slice based on a predicate
func FilterComplex128(a []complex128, keep func(x complex128) bool) []complex128 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterComplex64 performs in place filtering of a complex64 slice based on a predicate
func FilterComplex64(a []complex64, keep func(x complex64) bool) []complex64 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterFloat32 performs in place filtering of a float32 slice based on a predicate
func FilterFloat32(a []float32, keep func(x float32) bool) []float32 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterFloat64 performs in place filtering of a float64 slice based on a predicate
func FilterFloat64(a []float64, keep func(x float64) bool) []float64 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterInt performs in place filtering of an int slice based on a predicate
func FilterInt(a []int, keep func(x int) bool) []int {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterInt16 performs in place filtering of an int16 slice based on a predicate
func FilterInt16(a []int16, keep func(x int16) bool) []int16 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterInt32 performs in place filtering of an int32 slice based on a predicate
func FilterInt32(a []int32, keep func(x int32) bool) []int32 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterInt64 performs in place filtering of an int64 slice based on a predicate
func FilterInt64(a []int64, keep func(x int64) bool) []int64 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterInt8 performs in place filtering of an int8 slice based on a predicate
func FilterInt8(a []int8, keep func(x int8) bool) []int8 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterRune performs in place filtering of a rune slice based on a predicate
func FilterRune(a []rune, keep func(x rune) bool) []rune {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterString performs in place filtering of a string slice based on a predicate
func FilterString(a []string, keep func(x string) bool) []string {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUint performs in place filtering of a uint slice based on a predicate
func FilterUint(a []uint, keep func(x uint) bool) []uint {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUint16 performs in place filtering of a uint16 slice based on a predicate
func FilterUint16(a []uint16, keep func(x uint16) bool) []uint16 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUint32 performs in place filtering of a uint32 slice based on a predicate
func FilterUint32(a []uint32, keep func(x uint32) bool) []uint32 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUint64 performs in place filtering of a uint64 slice based on a predicate
func FilterUint64(a []uint64, keep func(x uint64) bool) []uint64 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUint8 performs in place filtering of a uint8 slice based on a predicate
func FilterUint8(a []uint8, keep func(x uint8) bool) []uint8 {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}

// FilterUintptr performs in place filtering of a uintptr slice based on a predicate
func FilterUintptr(a []uintptr, keep func(x uintptr) bool) []uintptr {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}
