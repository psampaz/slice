package slice

import (
	"math/rand"
	"time"
)

// ShuffleBool shuffles (in place) a bool slice
func ShuffleBool(a []bool) []bool {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleByte shuffles (in place) a byte slice
func ShuffleByte(a []byte) []byte {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleComplex128 shuffles (in place) a complex128 slice
func ShuffleComplex128(a []complex128) []complex128 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleComplex64 shuffles (in place) a complex64 slice
func ShuffleComplex64(a []complex64) []complex64 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleFloat32 shuffles (in place) a float32 slice
func ShuffleFloat32(a []float32) []float32 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleFloat64 shuffles (in place) a float64 slice
func ShuffleFloat64(a []float64) []float64 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt shuffles (in place) a int slice
func ShuffleInt(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt16 shuffles (in place) a int16 slice
func ShuffleInt16(a []int16) []int16 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt32 shuffles (in place) a int32 slice
func ShuffleInt32(a []int32) []int32 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt64 shuffles (in place) a int64 slice
func ShuffleInt64(a []int64) []int64 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt8 shuffles (in place) a int8 slice
func ShuffleInt8(a []int8) []int8 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleRune shuffles (in place) a rune slice
func ShuffleRune(a []rune) []rune {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleString shuffles (in place) a string slice
func ShuffleString(a []string) []string {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint shuffles (in place) a uint slice
func ShuffleUint(a []uint) []uint {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint16 shuffles (in place) a uint16 slice
func ShuffleUint16(a []uint16) []uint16 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint32 shuffles (in place) a uint32 slice
func ShuffleUint32(a []uint32) []uint32 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint64 shuffles (in place) a uint64 slice
func ShuffleUint64(a []uint64) []uint64 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint8 shuffles (in place) a uint8 slice
func ShuffleUint8(a []uint8) []uint8 {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUintptr shuffles (in place) a uintptr slice
func ShuffleUintptr(a []uintptr) []uintptr {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
