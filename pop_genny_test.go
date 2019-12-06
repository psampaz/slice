package slice

import (
	"reflect"
	"testing"
)

func TestPopType(t *testing.T) {
	type args struct {
		a []Type
	}
	tests := []struct {
		name    string
		args    args
		want    Type
		want1   []Type
		wantErr bool
	}{
		{
			name: "nil slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "empty slice",
			args: args{
				a: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},		
		{
			name: "non empty slice",
			args: args{
				a: []Type{1, 2, 3},
			},
			want:    3,
			want1:   []Type{1, 2},
			wantErr: false,
		},				
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopType(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopType() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
