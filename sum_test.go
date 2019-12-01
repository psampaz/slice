package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumByte(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []byte{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []byte{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumByte(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumComplex128(t *testing.T) {
	type args struct {
		a []complex128
	}
	tests := []struct {
		name    string
		args    args
		want    complex128
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    complex(0, 0),
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex128{},
			},
			want:    complex(0, 0),
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex128{complex(1, 1), complex(2, 2), complex(3, 3)},
			},
			want:    complex(6, 6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumComplex128(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumComplex64(t *testing.T) {
	type args struct {
		a []complex64
	}
	tests := []struct {
		name    string
		args    args
		want    complex64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    complex(0, 0),
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex64{},
			},
			want:    complex(0, 0),
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex64{complex(1, 1), complex(2, 2), complex(3, 3)},
			},
			want:    complex(6, 6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumComplex64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat32(t *testing.T) {
	type args struct {
		a []float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []float32{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float32{1.0, 2.0, 3.0},
			},
			want:    6.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFloat32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat64(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []float64{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float64{1.0, 2.0, 3.0},
			},
			want:    6.0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFloat64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt16(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int16{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int16{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt32(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int32{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int32{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt64(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int64{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int64{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt8(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int8{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int8{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumRune(t *testing.T) {
	type args struct {
		a []rune
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []rune{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []rune{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumRune(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUint(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUint(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUint16(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint16{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint16{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUint16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUint32(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint32{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint32{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUint32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUint64(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint64{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint64{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUint64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUint8(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint8{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint8{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUint8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUintptr(t *testing.T) {
	type args struct {
		a []uintptr
	}
	tests := []struct {
		name    string
		args    args
		want    uintptr
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uintptr{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uintptr{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUintptr(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSumInt() {
	a := []int{1, 2, 3}
	sum, err := SumInt(a)
	fmt.Printf("%d, %v", sum, err)
	// Output: 6, <nil>
}
