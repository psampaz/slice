package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeduplicateBool(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []bool{true, false, true, false},
			},
			want: []bool{true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateBool(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateByte(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []byte{1, 2, 3, 3, 2, 5, 3},
			},
			want: []byte{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateByte(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateComplex128(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []complex128{complex(1, 1), complex(2, 2), complex(3, 3), complex(3, 3), complex(4, 4), complex(5, 5), complex(3, 3)},
			},
			want: []complex128{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4), complex(5, 5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateComplex128(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateComplex64(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []complex64{complex(1, 1), complex(2, 2), complex(3, 3), complex(3, 3), complex(4, 4), complex(5, 5), complex(3, 3)},
			},
			want: []complex64{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4), complex(5, 5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateComplex64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateFloat32(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []float32{1.1, 2.2, 3.3, 3.3, 2.2, 5.5, 3.3},
			},
			want: []float32{1.1, 2.2, 3.3, 5.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateFloat32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateFloat64(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []float64{1.1, 2.2, 3.3, 3.3, 2.2, 5.5, 3.3},
			},
			want: []float64{1.1, 2.2, 3.3, 5.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateFloat64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateInt(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []int{1, 2, 3, 3, 2, 5, 3},
			},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateInt(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateInt16(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []int16{1, 2, 3, 3, 2, 5, 3},
			},
			want: []int16{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateInt16(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateInt32(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []int32{1, 2, 3, 3, 2, 5, 3},
			},
			want: []int32{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateInt32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateInt64(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []int64{1, 2, 3, 3, 2, 5, 3},
			},
			want: []int64{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateInt64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateInt8(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []int8{1, 2, 3, 3, 2, 5, 3},
			},
			want: []int8{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateInt8(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateRune(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []rune{1, 2, 3, 3, 2, 5, 3},
			},
			want: []rune{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateRune(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateString(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []string{"a", "b", "c", "c", "d", "e"},
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateString(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUint(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uint{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uint{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUint(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUint16(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uint16{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uint16{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUint16(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUint32(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uint32{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uint32{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUint32(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUint64(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uint64{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uint64{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUint64(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUint8(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uint8{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uint8{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUint8(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateUintptr(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []uintptr{1, 2, 3, 3, 2, 5, 3},
			},
			want: []uintptr{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateUintptr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDeduplicateInt() {
	a := []int{1, 2, 3, 2, 5, 3}
	a = DeduplicateInt(a)
	fmt.Printf("%v", a)
	// Output: [1 2 3 5]
}
