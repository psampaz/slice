package slice

// DeduplicateBool performs order preserving, in place deduplication of a bool slice
func DeduplicateBool(a []bool) []bool {
	if len(a) < 2 {
		return a
	}

	seen := make(map[bool]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateByte performs order preserving, in place deduplication of a byte slice
func DeduplicateByte(a []byte) []byte {
	if len(a) < 2 {
		return a
	}

	seen := make(map[byte]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateComplex128 performs order preserving, in place deduplication of a complex128 slice
func DeduplicateComplex128(a []complex128) []complex128 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[complex128]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateComplex64 performs order preserving, in place deduplication of a complex64 slice
func DeduplicateComplex64(a []complex64) []complex64 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[complex64]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateFloat32 performs order preserving, in place deduplication of a float32 slice
func DeduplicateFloat32(a []float32) []float32 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[float32]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateFloat64 performs order preserving, in place deduplication of a float64 slice
func DeduplicateFloat64(a []float64) []float64 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[float64]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateInt performs order preserving, in place deduplication of a int slice
func DeduplicateInt(a []int) []int {
	if len(a) < 2 {
		return a
	}

	seen := make(map[int]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateInt16 performs order preserving, in place deduplication of a int16 slice
func DeduplicateInt16(a []int16) []int16 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[int16]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateInt32 performs order preserving, in place deduplication of a int32 slice
func DeduplicateInt32(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[int32]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateInt64 performs order preserving, in place deduplication of a int64 slice
func DeduplicateInt64(a []int64) []int64 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[int64]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateInt8 performs order preserving, in place deduplication of a int8 slice
func DeduplicateInt8(a []int8) []int8 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[int8]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateRune performs order preserving, in place deduplication of a rune slice
func DeduplicateRune(a []rune) []rune {
	if len(a) < 2 {
		return a
	}

	seen := make(map[rune]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateString performs order preserving, in place deduplication of a string slice
func DeduplicateString(a []string) []string {
	if len(a) < 2 {
		return a
	}

	seen := make(map[string]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUint performs order preserving, in place deduplication of a uint slice
func DeduplicateUint(a []uint) []uint {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uint]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUint16 performs order preserving, in place deduplication of a uint16 slice
func DeduplicateUint16(a []uint16) []uint16 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uint16]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUint32 performs order preserving, in place deduplication of a uint32 slice
func DeduplicateUint32(a []uint32) []uint32 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uint32]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUint64 performs order preserving, in place deduplication of a uint64 slice
func DeduplicateUint64(a []uint64) []uint64 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uint64]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUint8 performs order preserving, in place deduplication of a uint8 slice
func DeduplicateUint8(a []uint8) []uint8 {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uint8]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}

// DeduplicateUintptr performs order preserving, in place deduplication of a uintptr slice
func DeduplicateUintptr(a []uintptr) []uintptr {
	if len(a) < 2 {
		return a
	}

	seen := make(map[uintptr]struct{}, 0)

	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}
