package slice

import (
	"reflect"
	"testing"
)

func TestFilterBool(t *testing.T) {
	type args struct {
		a    []bool
		keep func(x bool) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(bool) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []bool{},
				keep: func(bool) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []bool{true, false, true, false},
				keep: func(x bool) bool {
					return x
				},
			},
			want: []bool{true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterBool(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterByte(t *testing.T) {
	type args struct {
		a    []byte
		keep func(x byte) bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(byte) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []byte{},
				keep: func(byte) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x byte) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []byte{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterByte(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterComplex128(t *testing.T) {
	type args struct {
		a    []complex128
		keep func(x complex128) bool
	}
	tests := []struct {
		name string
		args args
		want []complex128
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(complex128) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []complex128{},
				keep: func(complex128) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex128{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x complex128) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []complex128{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComplex128(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterComplex64(t *testing.T) {
	type args struct {
		a    []complex64
		keep func(x complex64) bool
	}
	tests := []struct {
		name string
		args args
		want []complex64
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(complex64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []complex64{},
				keep: func(complex64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x complex64) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []complex64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComplex64(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFloat32(t *testing.T) {
	type args struct {
		a    []float32
		keep func(x float32) bool
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(float32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []float32{},
				keep: func(float32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9},
				keep: func(x float32) bool {
					if x < 5.5 {
						return true
					}
					return false
				},
			},
			want: []float32{1.1, 2.2, 3.3, 4.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterFloat32(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFloat64(t *testing.T) {
	type args struct {
		a    []float64
		keep func(x float64) bool
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(float64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []float64{},
				keep: func(float64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9},
				keep: func(x float32) bool {
					if x < 5.5 {
						return true
					}
					return false
				},
			},
			want: []float64{1.1, 2.2, 3.3, 4.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterFloat64(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInt(t *testing.T) {
	type args struct {
		a    []int
		keep func(x int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(int) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int{},
				keep: func(int) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x int) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInt16(t *testing.T) {
	type args struct {
		a    []int16
		keep func(x int16) bool
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(int16) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int16{},
				keep: func(int16) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x int16) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []int16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt16(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInt32(t *testing.T) {
	type args struct {
		a    []int32
		keep func(x int32) bool
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(int32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int32{},
				keep: func(int32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x int32) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []int32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt32(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInt64(t *testing.T) {
	type args struct {
		a    []int64
		keep func(x int64) bool
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(int64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int64{},
				keep: func(int64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x int64) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt64(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInt8(t *testing.T) {
	type args struct {
		a    []int8
		keep func(x int8) bool
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(int8) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int8{},
				keep: func(int8) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x int8) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []int8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt8(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterRune(t *testing.T) {
	type args struct {
		a    []rune
		keep func(x rune) bool
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(rune) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []rune{},
				keep: func(rune) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []rune{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x rune) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []rune{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterRune(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	type args struct {
		a    []string
		keep func(x string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(string) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []string{},
				keep: func(string) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []string{"a","b","c","d"},
				keep: func(x string) bool {
					if x == "a" || x == "d" {
						return true
					}
					return false
				},
			},
			want: []string{"a","d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterString(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUint(t *testing.T) {
	type args struct {
		a    []uint
		keep func(x uint) bool
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uint) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint{},
				keep: func(uint) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uint) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uint{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUint(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUint16(t *testing.T) {
	type args struct {
		a    []uint16
		keep func(x uint16) bool
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uint16) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint16{},
				keep: func(uint16) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uint16) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uint16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUint16(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUint32(t *testing.T) {
	type args struct {
		a    []uint32
		keep func(x uint32) bool
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uint32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint32{},
				keep: func(uint32) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uint32) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uint32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUint32(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUint64(t *testing.T) {
	type args struct {
		a    []uint64
		keep func(x uint64) bool
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uint64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint64{},
				keep: func(uint64) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uint64) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uint64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUint64(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUint8(t *testing.T) {
	type args struct {
		a    []uint8
		keep func(x uint8) bool
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uint8) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint8{},
				keep: func(uint8) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uint8) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uint8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUint8(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUintptr(t *testing.T) {
	type args struct {
		a    []uintptr
		keep func(x uintptr) bool
	}
	tests := []struct {
		name string
		args args
		want []uintptr
	}{
		{
			name: "nil slice",
			args: args{
				a:    nil,
				keep: func(uintptr) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uintptr{},
				keep: func(uintptr) bool { panic("not implemented") },
			},
			want: nil,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x uintptr) bool {
					if x < 5 {
						return true
					}
					return false
				},
			},
			want: []uintptr{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterUintptr(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}
