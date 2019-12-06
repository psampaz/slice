package slice

import (
	"reflect"
	"sort"
	"testing"
)

func TestShuffleBool_LessThan2Elements(t *testing.T) {
	type args struct {
		a []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []bool{},
			},
			want: []bool{},
		},
		{
			name: "one element",
			args: args{
				a: []bool{true},
			},
			want: []bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleBool(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleBool_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice. Let's assume that false < true for testing purposes.
	a := []bool{false, false, false, true, true, true}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]bool, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleBool(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUintptr(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return !shuffled[i] // this will put "false" first and "true" after
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleBool() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleByte_LessThan2Elements(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []byte{},
			},
			want: []byte{},
		},
		{
			name: "one element",
			args: args{
				a: []byte{1},
			},
			want: []byte{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleByte(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleByte_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []byte{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]byte, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleByte(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleByte(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleByte() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleComplex128_LessThan2Elements(t *testing.T) {
	type args struct {
		a []complex128
	}
	tests := []struct {
		name string
		args args
		want []complex128
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex128{},
			},
			want: []complex128{},
		},
		{
			name: "one element",
			args: args{
				a: []complex128{1},
			},
			want: []complex128{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleComplex128(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleComplex128_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice. Keep the imaginary part 0 to help with sorting operation later on.
	a := []complex128{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]complex128, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleComplex128(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleComplex128(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return real(shuffled[i]) < real(shuffled[j])
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleComplex128() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleComplex64_LessThan2Elements(t *testing.T) {
	type args struct {
		a []complex64
	}
	tests := []struct {
		name string
		args args
		want []complex64
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex64{},
			},
			want: []complex64{},
		},
		{
			name: "one element",
			args: args{
				a: []complex64{1},
			},
			want: []complex64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleComplex64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleComplex64_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice. Keep the imaginary part 0 to help with sorting operation later on.
	a := []complex64{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]complex64, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleComplex64(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleComplex64(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return real(shuffled[i]) < real(shuffled[j])
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleComplex64() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleFloat32_LessThan2Elements(t *testing.T) {
	type args struct {
		a []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []float32{},
			},
			want: []float32{},
		},
		{
			name: "one element",
			args: args{
				a: []float32{1},
			},
			want: []float32{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleFloat32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleFloat32_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []float32{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]float32, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleFloat32(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleFloat32(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleFloat32() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleFloat64_LessThan2Elements(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []float64{},
			},
			want: []float64{},
		},
		{
			name: "one element",
			args: args{
				a: []float64{1},
			},
			want: []float64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleFloat64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleFloat64_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []float64{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	shuffled := ShuffleFloat64(a)
	if reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleFloat64() = %v, Shuffled slice should not match the original", shuffled)
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleFloat64() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleInt_LessThan2Elements(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []int{},
			},
			want: []int{},
		},
		{
			name: "one element",
			args: args{
				a: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleInt(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleInt_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []int{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]int, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleInt(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleInt(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleInt() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleInt16_LessThan2Elements(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []int16{},
			},
			want: []int16{},
		},
		{
			name: "one element",
			args: args{
				a: []int16{1},
			},
			want: []int16{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleInt16(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleInt16_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []int16{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]int16, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleInt16(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleInt16(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleInt16() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleInt32_LessThan2Elements(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []int32{},
			},
			want: []int32{},
		},
		{
			name: "one element",
			args: args{
				a: []int32{1},
			},
			want: []int32{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleInt32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleInt32_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []int32{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]int32, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleInt32(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleInt32(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleInt32() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleInt64_LessThan2Elements(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []int64{},
			},
			want: []int64{},
		},
		{
			name: "one element",
			args: args{
				a: []int64{1},
			},
			want: []int64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleInt64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleInt64_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []int64{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]int64, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleInt64(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleInt64(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleInt64() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleInt8_LessThan2Elements(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []int8{},
			},
			want: []int8{},
		},
		{
			name: "one element",
			args: args{
				a: []int8{1},
			},
			want: []int8{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleInt8(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleInt8_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []int8{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]int8, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleInt8(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleInt8(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleInt8() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleRune_LessThan2Elements(t *testing.T) {
	type args struct {
		a []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []rune{},
			},
			want: []rune{},
		},
		{
			name: "one element",
			args: args{
				a: []rune{1},
			},
			want: []rune{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleRune(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleRune_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []rune{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]rune, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleRune(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleRune(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleRune() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleString_LessThan2Elements(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []string{},
			},
			want: []string{},
		},
		{
			name: "one element",
			args: args{
				a: []string{"a"},
			},
			want: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleString(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleString_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []string{"a", "b", "c", "d", "e", "f"}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]string, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleString(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleString(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleString() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUint_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint{},
			},
			want: []uint{},
		},
		{
			name: "one element",
			args: args{
				a: []uint{1},
			},
			want: []uint{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUint(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUint_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uint{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uint, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUint(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUint(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUint() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUint16_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint16{},
			},
			want: []uint16{},
		},
		{
			name: "one element",
			args: args{
				a: []uint16{1},
			},
			want: []uint16{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUint16(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUint16_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uint16{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uint16, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUint16(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUint16(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUint16() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUint32_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint32{},
			},
			want: []uint32{},
		},
		{
			name: "one element",
			args: args{
				a: []uint32{1},
			},
			want: []uint32{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUint32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUint32_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uint32{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uint32, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUint32(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUint32(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUint32() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUint64_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint64{},
			},
			want: []uint64{},
		},
		{
			name: "one element",
			args: args{
				a: []uint64{1},
			},
			want: []uint64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUint64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUint64_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uint64{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uint64, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUint64(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUint64(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUint64() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUint8_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint8{},
			},
			want: []uint8{},
		},
		{
			name: "one element",
			args: args{
				a: []uint8{1},
			},
			want: []uint8{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUint8(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUint8_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uint8{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uint8, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUint8(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUint8(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUint8() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}

func TestShuffleUintptr_LessThan2Elements(t *testing.T) {
	type args struct {
		a []uintptr
	}
	tests := []struct {
		name string
		args args
		want []uintptr
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []uintptr{},
			},
			want: []uintptr{},
		},
		{
			name: "one element",
			args: args{
				a: []uintptr{1},
			},
			want: []uintptr{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleUintptr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleUintptr_MoreThan2Elements(t *testing.T) {
	// start with a sorted slice
	a := []uintptr{1, 2, 3, 4, 5}

	// Make sure the shuffled slice elements are in different order that the original slice
	original := append(a[:0:0], a...)
	different := false
	shuffled := make([]uintptr, len(a))
	// This is a fragile test because the result of the shuffling might be exactly the same as the original
	// We run the shuffle 100 until until we get the first correct shuffle.
	// Todo: find a smarter way to test this
	for x := 0; x <= 100; x++ {
		shuffled = ShuffleUintptr(a)
		if !reflect.DeepEqual(shuffled, original) {
			different = true
			break
		}
	}
	if different == false {
		t.Errorf("ShuffleUintptr(): Shuffled slice should not match the original")
	}
	// Sort the shuffled slice in order to ensure that its elements are the same
	sort.Slice(shuffled, func(i, j int) bool {
		return shuffled[i] < shuffled[j]
	})
	if !reflect.DeepEqual(shuffled, original) {
		t.Errorf("ShuffleUintptr() = %v, Shuffled slice should have the same elements with the original %v", shuffled, original)
	}
}
