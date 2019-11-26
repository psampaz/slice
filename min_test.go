package slice

import (
	"fmt"
	"testing"
)

func TestMinByte(t *testing.T) {
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
				a: []byte{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinByte(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat32(t *testing.T) {
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
				a: []float32{1.1, 3.3, 2.2, -5.5, 4.4, 6.6},
			},
			want:    -5.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinFloat32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat64(t *testing.T) {
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
				a: []float64{1.1, 3.3, 2.2, -5.5, 4.4, 6.6},
			},
			want:    -5.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinFloat64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
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
				a: []int{1, 3, 2, -4, 6, 5},
			},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt16(t *testing.T) {
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
				a: []int16{1, 3, 2, -4, 6, 5},
			},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt32(t *testing.T) {
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
				a: []int32{1, 3, 2, -4, 6, 5},
			},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64(t *testing.T) {
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
				a: []int64{1, 3, 2, -4, 6, 5},
			},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt8(t *testing.T) {
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
				a: []int8{1, 3, 2, -4, 6, 5},
			},
			want:    -4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinRune(t *testing.T) {
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
				a: []rune{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinRune(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint(t *testing.T) {
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
				a: []uint{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUint(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint16(t *testing.T) {
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
				a: []uint16{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUint16(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint32(t *testing.T) {
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
				a: []uint32{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUint32(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint64(t *testing.T) {
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
				a: []uint64{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUint64(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint8(t *testing.T) {
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
				a: []uint8{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUint8(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUintptr(t *testing.T) {
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
				a: []uintptr{1, 3, 2, 0, 5, 4},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinUintptr(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMinInt() {
	a := []int{1, 2, 3, 0, 7, 5, 2}
	min, err := MinInt(a)
	fmt.Printf("%d, %v", min, err)
	// Output: 0, <nil>
}
