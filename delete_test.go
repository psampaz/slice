package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteBool(t *testing.T) {
	type args struct {
		a []bool
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []bool{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []bool{true, false, true, false, true},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []bool{true, false, true, false, true},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []bool{true, false, true, false, true},
				i: 2,
			},
			want:    []bool{true, false, false, true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteBool(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteByte(t *testing.T) {
	type args struct {
		a []byte
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []byte{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []byte{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []byte{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []byte{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []byte{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteByte(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteComplex128(t *testing.T) {
	type args struct {
		a []complex128
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []complex128
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex128{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []complex128{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []complex128{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex128{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []complex128{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteComplex128(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteComplex64(t *testing.T) {
	type args struct {
		a []complex64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []complex64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []complex64{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []complex64{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []complex64{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []complex64{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []complex64{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteComplex64(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFloat32(t *testing.T) {
	type args struct {
		a []float32
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []float32{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []float32{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []float32{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float32{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []float32{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteFloat32(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFloat64(t *testing.T) {
	type args struct {
		a []float64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []float64{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []float64{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []float64{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []float64{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []float64{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteFloat64(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt(t *testing.T) {
	type args struct {
		a []int
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []int{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteInt(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt16(t *testing.T) {
	type args struct {
		a []int16
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int16{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []int16{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []int16{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int16{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []int16{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteInt16(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt32(t *testing.T) {
	type args struct {
		a []int32
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int32{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []int32{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteInt32(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt64(t *testing.T) {
	type args struct {
		a []int64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int64{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []int64{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []int64{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int64{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []int64{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteInt64(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteInt8(t *testing.T) {
	type args struct {
		a []int8
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []int8{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []int8{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []int8{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []int8{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []int8{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteInt8(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRune(t *testing.T) {
	type args struct {
		a []rune
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []rune
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []rune{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []rune{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []rune{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []rune{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []rune{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRune(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteString(t *testing.T) {
	type args struct {
		a []string
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []string{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []string{"a", "b", "c", "d", "e"},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []string{"a", "b", "c", "d", "e"},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []string{"a", "b", "c", "d", "e"},
				i: 2,
			},
			want:    []string{"a", "b", "d", "e"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteString(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint(t *testing.T) {
	type args struct {
		a []uint
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uint{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uint{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uint{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUint(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint16(t *testing.T) {
	type args struct {
		a []uint16
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint16{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uint16{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uint16{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint16{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uint16{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUint16(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint32(t *testing.T) {
	type args struct {
		a []uint32
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint32
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint32{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uint32{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uint32{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint32{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uint32{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUint32(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint64(t *testing.T) {
	type args struct {
		a []uint64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint64{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uint64{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uint64{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint64{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uint64{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUint64(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUint8(t *testing.T) {
	type args struct {
		a []uint8
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint8
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uint8{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uint8{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uint8{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uint8{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uint8{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUint8(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteUintptr(t *testing.T) {
	type args struct {
		a []uintptr
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []uintptr
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []uintptr{},
				i: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index larger than len(a)-1",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5},
				i: 5,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice. Index < 0",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5},
				i: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []uintptr{1, 2, 3, 4, 5},
				i: 2,
			},
			want:    []uintptr{1, 2, 4, 5},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteUintptr(tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDeleteInt() {
	a := []int{1, 2, 3, 4, 5}
	a, err := DeleteInt(a, 2)
	fmt.Printf("%v, %v", a, err)
	// Output: [1 2 4 5], <nil>
}
