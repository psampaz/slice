package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPopBool(t *testing.T) {
	type args struct {
		a []bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   []bool
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    false,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    false,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []bool{true},
			},
			want:    true,
			want1:   []bool{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []bool{true, false, true},
			},
			want:    true,
			want1:   []bool{true, false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopBool(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopBool() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopByte(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		want1   []byte
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []byte{1},
			},
			want:    1,
			want1:   []byte{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []byte{1, 2, 3},
			},
			want:    3,
			want1:   []byte{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopByte(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopByte() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopByte() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopComplex128(t *testing.T) {
	type args struct {
		a []complex128
	}
	tests := []struct {
		name    string
		args    args
		want    complex128
		want1   []complex128
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []complex128{1},
			},
			want:    1,
			want1:   []complex128{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex128{1, 2, 3},
			},
			want:    3,
			want1:   []complex128{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopComplex128(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopComplex128() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopComplex128() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopComplex64(t *testing.T) {
	type args struct {
		a []complex64
	}
	tests := []struct {
		name    string
		args    args
		want    complex64
		want1   []complex64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []complex64{1},
			},
			want:    1,
			want1:   []complex64{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex64{1, 2, 3},
			},
			want:    3,
			want1:   []complex64{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopComplex64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopComplex64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopComplex64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopFloat32(t *testing.T) {
	type args struct {
		a []float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		want1   []float32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []float32{1},
			},
			want:    1,
			want1:   []float32{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float32{1, 2, 3},
			},
			want:    3,
			want1:   []float32{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopFloat32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFloat32() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopFloat32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopFloat64(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		want1   []float64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []float64{1},
			},
			want:    1,
			want1:   []float64{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float64{1, 2, 3},
			},
			want:    3,
			want1:   []float64{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopFloat64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFloat64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopFloat64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   []int
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []int{1},
			},
			want:    1,
			want1:   []int{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int{1, 2, 3},
			},
			want:    3,
			want1:   []int{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopInt(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopInt() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopInt16(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		want1   []int16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []int16{1},
			},
			want:    1,
			want1:   []int16{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int16{1, 2, 3},
			},
			want:    3,
			want1:   []int16{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopInt16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopInt16() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopInt32(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		want1   []int32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []int32{1},
			},
			want:    1,
			want1:   []int32{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int32{1, 2, 3},
			},
			want:    3,
			want1:   []int32{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopInt32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopInt32() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopInt64(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   []int64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []int64{1},
			},
			want:    1,
			want1:   []int64{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int64{1, 2, 3},
			},
			want:    3,
			want1:   []int64{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopInt64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopInt64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopInt8(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		want1   []int8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []int8{1},
			},
			want:    1,
			want1:   []int8{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int8{1, 2, 3},
			},
			want:    3,
			want1:   []int8{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopInt8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopInt8() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopInt8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopRune(t *testing.T) {
	type args struct {
		a []rune
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		want1   []rune
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []rune{1},
			},
			want:    1,
			want1:   []rune{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []rune{1, 2, 3},
			},
			want:    3,
			want1:   []rune{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopRune(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopRune() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopRune() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopString(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []string{"a"},
			},
			want:    "a",
			want1:   []string{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []string{"a", "b", "c"},
			},
			want:    "c",
			want1:   []string{"a", "b"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopString(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopString() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUint(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		want1   []uint
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uint{1},
			},
			want:    1,
			want1:   []uint{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint{1, 2, 3},
			},
			want:    3,
			want1:   []uint{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUint(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUint() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUint() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUint16(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		want1   []uint16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uint16{1},
			},
			want:    1,
			want1:   []uint16{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint16{1, 2, 3},
			},
			want:    3,
			want1:   []uint16{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUint16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUint16() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUint16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUint32(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		want1   []uint32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uint32{1},
			},
			want:    1,
			want1:   []uint32{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint32{1, 2, 3},
			},
			want:    3,
			want1:   []uint32{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUint32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUint32() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUint32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUint64(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		want1   []uint64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uint64{1},
			},
			want:    1,
			want1:   []uint64{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint64{1, 2, 3},
			},
			want:    3,
			want1:   []uint64{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUint64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUint64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUint64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUint8(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		want1   []uint8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uint8{1},
			},
			want:    1,
			want1:   []uint8{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint8{1, 2, 3},
			},
			want:    3,
			want1:   []uint8{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUint8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUint8() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUint8() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPopUintptr(t *testing.T) {
	type args struct {
		a []uintptr
	}
	tests := []struct {
		name    string
		args    args
		want    uintptr
		want1   []uintptr
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "non empty slice - one element",
			args: args{
				a: []uintptr{1},
			},
			want:    1,
			want1:   []uintptr{},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uintptr{1, 2, 3},
			},
			want:    3,
			want1:   []uintptr{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopUintptr(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopUintptr() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopUintptr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func ExamplePopInt() {
	a := []int{1, 2, 3, 4, 5}
	v, a, err := PopInt(a)
	fmt.Printf("%d, %v, %v", v, a, err)
	// Output: 5, [1 2 3 4], <nil>
}
