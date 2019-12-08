package slice

import (
	"reflect"
	"testing"
)

func TestCutType(t *testing.T) {
	type args struct {
		a    []Type
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []Type
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
				a:    []Type{},
				from: 0,
				to:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from index out of bounds",
			args: args{
				a:    []Type{1, 2, 3, 4, 5},
				from: -1,
				to:   3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - to index out of bounds",
			args: args{
				a:    []Type{1, 2, 3, 4, 5},
				from: 1,
				to:   6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice - from > to",
			args: args{
				a:    []Type{1, 2, 3, 4, 5},
				from: 3,
				to:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice single elements",
			args: args{
				a:    []Type{1},
				from: 0,
				to:   0,
			},
			want:    []Type{},
			wantErr: false,
		},
		{
			name: "non empty slice - cut one element",
			args: args{
				a:    []Type{1, 2, 3},
				from: 1,
				to:   1,
			},
			want:    []Type{1, 3},
			wantErr: false,
		},
		{
			name: "non empty slice",
			args: args{
				a:    []Type{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want:    []Type{1, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteManyType(tt.args.a, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteManyType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteManyType() = %v, want %v", got, tt.want)
			}
		})
	}
}
