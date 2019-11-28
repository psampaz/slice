package slice

import (
	"reflect"
	"testing"
)

func TestFilterType1(t *testing.T) {
	type args struct {
		a    []Type
		keep func(x Type) bool
	}
	tests := []struct {
		name string
		args args
		want []Type
	}{
		{
			name: "non empty slice",
			args: args{
				a: []Type{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				keep: func(x Type) bool {
					if x < 5 {
						return true
					} else {
						return false
					}
				},
			},
			want: []Type{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterType(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
