package slice

import "testing"

func TestContainsBool(t *testing.T) {
	type args struct {
		a []bool
		x bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: true,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []bool{},
				x: true,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []bool{true, false, true, false},
				x: true,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []bool{true, true},
				x: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsBool(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsByte(t *testing.T) {
	type args struct {
		a []byte
		x byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []byte{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []byte{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []byte{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsByte(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsComplex128(t *testing.T) {
	type args struct {
		a []complex128
		x complex128
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: complex(1, 1),
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex128{},
				x: complex(1, 1),
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []complex128{complex(1, 1), complex(2, 2), complex(3, 3)},
				x: complex(2, 2),
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []complex128{complex(1, 1), complex(2, 2), complex(3, 3)},
				x: complex(4, 4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsComplex128(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsComplex64(t *testing.T) {
	type args struct {
		a []complex64
		x complex64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: complex(1, 1),
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex64{},
				x: complex(1, 1),
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []complex64{complex(1, 1), complex(2, 2), complex(3, 3)},
				x: complex(2, 2),
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []complex64{complex(1, 1), complex(2, 2), complex(3, 3)},
				x: complex(4, 4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsComplex64(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFloat32(t *testing.T) {
	type args struct {
		a []float32
		x float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []float32{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []float32{1.1, 2.2, 3.3, 4.4, 5.5},
				x: 3.3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []float32{1.1, 2.2, 3.3, 4.4, 5.5},
				x: 6.6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFloat32(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFloat64(t *testing.T) {
	type args struct {
		a []float64
		x float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []float64{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				x: 3.3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				x: 6.6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFloat64(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []int{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt16(t *testing.T) {
	type args struct {
		a []int16
		x int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []int16{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []int16{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []int16{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt16(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt32(t *testing.T) {
	type args struct {
		a []int32
		x int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []int32{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt32(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt64(t *testing.T) {
	type args struct {
		a []int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []int64{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []int64{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []int64{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt64(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt8(t *testing.T) {
	type args struct {
		a []int8
		x int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []int8{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []int8{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []int8{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt8(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsRune(t *testing.T) {
	type args struct {
		a []rune
		x rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []rune{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []rune{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []rune{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsRune(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: "a",
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []string{},
				x: "a",
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []string{"a", "b", "c", "d", "e"},
				x: "a",
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []string{"a", "b", "c", "d", "e"},
				x: "f",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint(t *testing.T) {
	type args struct {
		a []uint
		x uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uint{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uint{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint16(t *testing.T) {
	type args struct {
		a []uint16
		x uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint16{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uint16{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uint16{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint16(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint32(t *testing.T) {
	type args struct {
		a []uint32
		x uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint32{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uint32{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uint32{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint32(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint64(t *testing.T) {
	type args struct {
		a []uint64
		x uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint64{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uint64{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uint64{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint64(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint8(t *testing.T) {
	type args struct {
		a []uint8
		x uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint8{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uint8{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uint8{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint8(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUintptr(t *testing.T) {
	type args struct {
		a []uintptr
		x uintptr
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				x: 0,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: []uintptr{},
				x: 0,
			},
			want: false,
		},
		{
			name: "non empty slice containing value",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5},
				x: 3,
			},
			want: true,
		},
		{
			name: "non empty slice not containing value",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5},
				x: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUintptr(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}
