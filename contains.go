package slice

// ContainsBool checks if a value exists in a bool slice
func ContainsBool(a []bool, x bool) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsByte checks if a value exists in a byte slice
func ContainsByte(a []byte, x byte) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsComplex64 checks if a value exists in a complex64 slice
func ContainsComplex64(a []complex64, x complex64) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsComplex128 checks if a value exists in a complex128 slice
func ContainsComplex128(a []complex128, x complex128) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsFloat32 checks if a value exists in a float32 slice
func ContainsFloat32(a []float32, x float32) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsFloat64 checks if a value exists in a float64 slice
func ContainsFloat64(a []float64, x float64) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsInt checks if a value exists in an int slice
func ContainsInt(a []int, x int) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsInt16 checks if a value exists in an int16 slice
func ContainsInt16(a []int16, x int16) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsInt32 checks if a value exists in an int32 slice
func ContainsInt32(a []int32, x int32) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsInt64 checks if a value exists in an int64 slice
func ContainsInt64(a []int64, x int64) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsInt8 checks if a value exists in an int8 slice
func ContainsInt8(a []int8, x int8) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsRune checks if a value exists in a rune slice
func ContainsRune(a []rune, x rune) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsString checks if a value exists in a string slice
func ContainsString(a []string, x string) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUint checks if a value exists in a uint slice
func ContainsUint(a []uint, x uint) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUint8 checks if a value exists in a uint8 slice
func ContainsUint8(a []uint8, x uint8) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUint16 checks if a value exists in a uint16 slice
func ContainsUint16(a []uint16, x uint16) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUint32 checks if a value exists in a uint32 slice
func ContainsUint32(a []uint32, x uint32) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUint64 checks if a value exists in a uint64 slice
func ContainsUint64(a []uint64, x uint64) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}

// ContainsUintptr checks if a value exists in a uintptr slice
func ContainsUintptr(a []uintptr, x uintptr) bool {
	if len(a) == 0 {
		return false
	}
	for k := range a {
		if a[k] == x {
			return true
		}
	}
	return false
}
