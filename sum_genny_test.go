package slice

import (
	"reflect"
	"testing"
)

func TestSumType(t *testing.T) {
	type args struct {
		a []Type
	}
	tests := []struct {
		name    string
		args    args
		want    Type
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: []Type{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non empty slice",
			args: args{
				a: []Type{1, 2, 3},
			},
			want:    6,
			wantErr: false,
		},					
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumType(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumType() = %v, want %v", got, tt.want)
			}
		})
	}
}
