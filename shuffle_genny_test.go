package slice

import (
	"reflect"
	"testing"
)

func TestShuffleType(t *testing.T) {
	type args struct {
		a []Type
	}
	tests := []struct {
		name string
		args args
		want []Type
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				a: []Type{},
			},
			want: []Type{},
		},
		{
			name: "one element",
			args: args{
				a: []Type{1},
			},
			want: []Type{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleType(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
