package slice

// CopyBool creates a copy of a bool slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyBool(a []bool) []bool {
	return append(a[:0:0], a...)
}

// CopyByte creates a copy of a byte slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyByte(a []byte) []byte {
	return append(a[:0:0], a...)
}

// CopyComplex128 creates a copy of a complex128 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyComplex128(a []complex128) []complex128 {
	return append(a[:0:0], a...)
}

// CopyComplex64 creates a copy of a complex64 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyComplex64(a []complex64) []complex64 {
	return append(a[:0:0], a...)
}

// CopyFloat32 creates a copy of a float32 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyFloat32(a []float32) []float32 {
	return append(a[:0:0], a...)
}

// CopyFloat64 creates a copy of a float64 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyFloat64(a []float64) []float64 {
	return append(a[:0:0], a...)
}

// CopyInt creates a copy of an int slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyInt(a []int) []int {
	return append(a[:0:0], a...)
}

// CopyInt16 creates a copy of an int16 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyInt16(a []int16) []int16 {
	return append(a[:0:0], a...)
}

// CopyInt32 creates a copy of an int32 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyInt32(a []int32) []int32 {
	return append(a[:0:0], a...)
}

// CopyInt64 creates a copy of an int64 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyInt64(a []int64) []int64 {
	return append(a[:0:0], a...)
}

// CopyInt8 creates a copy of an int8 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyInt8(a []int8) []int8 {
	return append(a[:0:0], a...)
}

// CopyRune creates a copy of a rune slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyRune(a []rune) []rune {
	return append(a[:0:0], a...)
}

// CopyString creates a copy of a string slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyString(a []string) []string {
	return append(a[:0:0], a...)
}

// CopyUint creates a copy of a uint slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUint(a []uint) []uint {
	return append(a[:0:0], a...)
}

// CopyUint16 creates a copy of a uint16 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUint16(a []uint16) []uint16 {
	return append(a[:0:0], a...)
}

// CopyUint32 creates a copy of a uint32 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUint32(a []uint32) []uint32 {
	return append(a[:0:0], a...)
}

// CopyUint64 creates a copy of a uint64 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUint64(a []uint64) []uint64 {
	return append(a[:0:0], a...)
}

// CopyUint8 creates a copy of a uint8 slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUint8(a []uint8) []uint8 {
	return append(a[:0:0], a...)
}

// CopyUintptr creates a copy of a uintptr slice.
// The resulting slice has the same elements as the original but the underlying array is different.
// See https://github.com/go101/go101/wiki
func CopyUintptr(a []uintptr) []uintptr {
	return append(a[:0:0], a...)
}
