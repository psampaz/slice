package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteRangeBool(t *testing.T) {
	type args struct {
		a    []bool
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []bool{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []bool{true, true, false, true, false},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []bool{true, true, false, true, false},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []bool{true, true, false, true, false},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []bool{true},
				from: 0,
				to:   0,
			},
			want:    []bool{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []bool{true, false, true},
				from: 1,
				to:   1,
			},
			want:    []bool{true, true},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []bool{true, true, false, true, false},
				from: 1,
				to:   3,
			},
			want:    []bool{true, false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeBool(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeByte(t *testing.T) {
	type args struct {
		a    []byte
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []byte{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []byte{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []byte{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []byte{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []byte{1},
				from: 0,
				to:   0,
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []byte{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []byte{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []byte{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []byte{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeByte(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeComplex128(t *testing.T) {
	type args struct {
		a    []complex128
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []complex128{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []complex128{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []complex128{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []complex128{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []complex128{1},
				from: 0,
				to:   0,
			},
			want:    []complex128{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []complex128{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []complex128{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []complex128{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []complex128{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeComplex128(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeComplex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeComplex64(t *testing.T) {
	type args struct {
		a    []complex64
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []complex64{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []complex64{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []complex64{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []complex64{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []complex64{1},
				from: 0,
				to:   0,
			},
			want:    []complex64{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []complex64{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []complex64{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []complex64{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []complex64{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeComplex64(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeComplex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeFloat32(t *testing.T) {
	type args struct {
		a    []float32
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []float32{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []float32{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []float32{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []float32{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []float32{1},
				from: 0,
				to:   0,
			},
			want:    []float32{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []float32{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []float32{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []float32{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []float32{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeFloat32(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeFloat64(t *testing.T) {
	type args struct {
		a    []float64
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []float64{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []float64{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []float64{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []float64{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []float64{1},
				from: 0,
				to:   0,
			},
			want:    []float64{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []float64{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []float64{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []float64{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []float64{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeFloat64(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeInt(t *testing.T) {
	type args struct {
		a    []int
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []int{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []int{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []int{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []int{1},
				from: 0,
				to:   0,
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []int{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []int{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []int{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []int{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeInt(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeInt16(t *testing.T) {
	type args struct {
		a    []int16
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int16{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []int16{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []int16{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []int16{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []int16{1},
				from: 0,
				to:   0,
			},
			want:    []int16{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []int16{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []int16{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []int16{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []int16{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeInt16(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeInt32(t *testing.T) {
	type args struct {
		a    []int32
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int32{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []int32{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []int32{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []int32{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []int32{1},
				from: 0,
				to:   0,
			},
			want:    []int32{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []int32{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []int32{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []int32{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []int32{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeInt32(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeInt64(t *testing.T) {
	type args struct {
		a    []int64
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int64{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []int64{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []int64{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []int64{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []int64{1},
				from: 0,
				to:   0,
			},
			want:    []int64{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []int64{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []int64{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []int64{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []int64{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeInt64(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeInt8(t *testing.T) {
	type args struct {
		a    []int8
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []int8{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []int8{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []int8{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []int8{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []int8{1},
				from: 0,
				to:   0,
			},
			want:    []int8{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []int8{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []int8{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []int8{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []int8{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeInt8(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeRune(t *testing.T) {
	type args struct {
		a    []rune
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []rune{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []rune{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []rune{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []rune{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []rune{1},
				from: 0,
				to:   0,
			},
			want:    []rune{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []rune{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []rune{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []rune{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []rune{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeRune(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeString(t *testing.T) {
	type args struct {
		a    []string
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []string{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []string{"a", "b", "c", "d", "e"},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []string{"a", "b", "c", "d", "e"},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []string{"a", "b", "c", "d", "e"},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []string{"a"},
				from: 0,
				to:   0,
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []string{"a", "b", "c"},
				from: 1,
				to:   1,
			},
			want:    []string{"a", "c"},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []string{"a", "b", "c", "d", "e"},
				from: 1,
				to:   3,
			},
			want:    []string{"a", "e"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeString(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUint(t *testing.T) {
	type args struct {
		a    []uint
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uint{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uint{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uint{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uint{1},
				from: 0,
				to:   0,
			},
			want:    []uint{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uint{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uint{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uint{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uint{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUint(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUint16(t *testing.T) {
	type args struct {
		a    []uint16
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint16{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uint16{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uint16{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uint16{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uint16{1},
				from: 0,
				to:   0,
			},
			want:    []uint16{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uint16{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uint16{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uint16{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uint16{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUint16(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUint32(t *testing.T) {
	type args struct {
		a    []uint32
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint32{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uint32{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uint32{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uint32{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uint32{1},
				from: 0,
				to:   0,
			},
			want:    []uint32{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uint32{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uint32{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uint32{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uint32{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUint32(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUint64(t *testing.T) {
	type args struct {
		a    []uint64
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint64{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uint64{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uint64{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uint64{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uint64{1},
				from: 0,
				to:   0,
			},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uint64{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uint64{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uint64{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uint64{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUint64(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUint8(t *testing.T) {
	type args struct {
		a    []uint8
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uint8{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uint8{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uint8{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uint8{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uint8{1},
				from: 0,
				to:   0,
			},
			want:    []uint8{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uint8{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uint8{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uint8{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uint8{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUint8(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRangeUintptr(t *testing.T) {
	type args struct {
		a    []uintptr
		from int
		to   int
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
				a:    nil,
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a:    []uintptr{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []uintptr{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []uintptr{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []uintptr{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []uintptr{1},
				from: 0,
				to:   0,
			},
			want:    []uintptr{},
			wantErr: false,
		},
		{
			name: "non empty slice - delete one element",
			args: args{
				a:    []uintptr{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []uintptr{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []uintptr{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []uintptr{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRangeUintptr(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRangeUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRangeUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDeleteRangeInt() {
	a := []int{1, 2, 3, 4, 5}
	a, err := DeleteRangeInt(a, 2, 3)
	fmt.Printf("%v, %v", a, err)
	// Output: [1 2 5], <nil>
}
