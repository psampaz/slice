package slice

import (
	"reflect"
	"testing"
)

func TestReverseType(t *testing.T) {
	type args struct {
		src []Type
	}
	tests := []struct {
		name string
		args args
		want []Type
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []Type{},
			},
			want: []Type{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []Type{1, 2, 3, 4, 5},
			},
			want: []Type{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []Type{1, 2, 3, 4},
			},
			want: []Type{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseType(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseType() = %v, want %v", got, tt.want)
			}
		})
	}
}
