package slice

import (
	"reflect"
	"testing"
)

func TestDeduplicateType(t *testing.T) {
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
			name: "non empty slice",
			args: args{
				a: []Type{1, 2, 3, 3, 2, 5, 3},
			},
			want: []Type{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeduplicateType(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeduplicateType() = %v, want %v", got, tt.want)
			}
		})
	}
}
