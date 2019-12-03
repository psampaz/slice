package slice

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestCopyBool(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []bool(nil)
		got := CopyBool(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyBool() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []bool{}
		got := CopyBool(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyBool() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []bool{true, false, true}
		got := CopyBool(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyBool() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyBool() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyByte(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []byte(nil)
		got := CopyByte(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyByte() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []byte{}
		got := CopyByte(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyByte() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []byte{1, 2, 3}
		got := CopyByte(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyByte() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyByte() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyComplex128(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []complex128(nil)
		got := CopyComplex128(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyComplex128() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []complex128{}
		got := CopyComplex128(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyComplex128() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
		got := CopyComplex128(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyComplex128() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyComplex128() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyComplex64(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []complex64(nil)
		got := CopyComplex64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyComplex64() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []complex64{}
		got := CopyComplex64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyComplex64() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []complex64{1, 2, 3}
		got := CopyComplex64(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyComplex64() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyComplex64() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyFloat32(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []float32(nil)
		got := CopyFloat32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyFloat32() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []float32{}
		got := CopyFloat32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyFloat32() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []float32{1.1, 2.2, 3.3}
		got := CopyFloat32(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyFloat32() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyFloat32() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyFloat64(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []float64(nil)
		got := CopyFloat64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyFloat64() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []float64{}
		got := CopyFloat64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyFloat64() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []float64{1.1, 2.2, 3.3}
		got := CopyFloat64(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyFloat64() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyFloat64() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyInt(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []int(nil)
		got := CopyInt(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int{}
		got := CopyInt(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []int{1, 2, 3}
		got := CopyInt(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyInt() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyInt() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyInt16(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []int16(nil)
		got := CopyInt16(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt16() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int16{}
		got := CopyInt16(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt16() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []int16{1, 2, 3}
		got := CopyInt16(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyInt16() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyInt16() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyInt32(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []int32(nil)
		got := CopyInt32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt32() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int32{}
		got := CopyInt32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt32() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []int32{1, 2, 3}
		got := CopyInt32(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyInt32() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyInt32() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyInt64(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []int64(nil)
		got := CopyInt64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt64() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int64{}
		got := CopyInt64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt64() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []int64{1, 2, 3}
		got := CopyInt64(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyInt64() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyInt64() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyInt8(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []int8(nil)
		got := CopyInt8(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt8() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []int8{}
		got := CopyInt8(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyInt8() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []int8{1, 2, 3}
		got := CopyInt8(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyInt8() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyInt8() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyRune(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []rune(nil)
		got := CopyRune(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyRune() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []rune{}
		got := CopyRune(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyRune() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []rune{1, 2, 3}
		got := CopyRune(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyRune() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyRune() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyString(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []string(nil)
		got := CopyString(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyString() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []string{}
		got := CopyString(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyString() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []string{"a", "b", "c"}
		got := CopyString(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyString() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyString() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUint(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uint(nil)
		got := CopyUint(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uint{}
		got := CopyUint(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uint{1, 2, 3}
		got := CopyUint(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUint() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUint() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUint16(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uint16(nil)
		got := CopyUint16(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint16() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uint16{}
		got := CopyUint16(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint16() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uint16{1, 2, 3}
		got := CopyUint16(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUint16() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUint16() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUint32(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uint32(nil)
		got := CopyUint32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint32() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uint32{}
		got := CopyUint32(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint32() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uint32{1, 2, 3}
		got := CopyUint32(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUint32() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUint32() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUint64(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uint64(nil)
		got := CopyUint64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint64() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uint64{}
		got := CopyUint64(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint64() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uint64{1, 2, 3}
		got := CopyUint64(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUint64() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUint64() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUint8(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uint8(nil)
		got := CopyUint8(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint8() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uint8{}
		got := CopyUint8(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUint8() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uint8{1, 2, 3}
		got := CopyUint8(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUint8() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUint8() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}

func TestCopyUintptr(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		src := []uintptr(nil)
		got := CopyUintptr(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUintptr() = %v, want %v", got, src)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		src := []uintptr{}
		got := CopyUintptr(src)
		if !reflect.DeepEqual(got, src) {
			t.Errorf("CopyUintptr() = %v, want %v", got, src)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		src := []uintptr{1, 2, 3}
		got := CopyUintptr(src)
		// Check that the original and the resulting slice do not share the same underlying array
		h1 := (*reflect.SliceHeader)(unsafe.Pointer(&src))
		h2 := (*reflect.SliceHeader)(unsafe.Pointer(&got))
		if h1.Data == h2.Data {
			t.Errorf("CopyUintptr() = %v, Slice copy should not share the same underlying array with the original slice", got)
		}
		// Check that that contents of the slice copy are the same with the original slice.
		for k := range src {
			if src[k] != got[k] {
				t.Errorf("CopyUintptr() = %v, want %v. Contents of the slice copy are the same with the original slice", got, src)
			}

		}
	})
}
