package slice

import (
	"fmt"
	"testing"
)

func TestMaxByte(t *testing.T) {
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
				a: []byte{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxByte(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat32(t *testing.T) {
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
				a: []float32{1.1, 4.1, 5.1, -34.4, 2.2, 100.1},
			},
			want:    100.1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxFloat32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat64(t *testing.T) {
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
				a: []float64{1.1, 4.1, 5.1, -34.4, 2.2, 100.1},
			},
			want:    100.1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxFloat64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
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
				a: []int{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt8(t *testing.T) {
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
				a: []int8{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt16(t *testing.T) {
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
				a: []int16{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt32(t *testing.T) {
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
				a: []int32{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
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
				a: []int64{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxRune(t *testing.T) {
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
				a: []rune{1, 4, 5, -34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxRune(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint(t *testing.T) {
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
				a: []uint{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUint(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint16(t *testing.T) {
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
				a: []uint16{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUint16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint32(t *testing.T) {
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
				a: []uint32{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUint32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint64(t *testing.T) {
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
				a: []uint64{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUint64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint8(t *testing.T) {
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
				a: []uint8{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUint8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUintptr(t *testing.T) {
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
				a: []uintptr{1, 4, 5, 34, 2, 100},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxUintptr(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMaxInt() {
	a := []int{1, 2, 3, 0, 7, 5, 2}
	max, err := MaxInt(a)
	fmt.Printf("%d, %v", max, err)
	// Output: 7, <nil>
}
