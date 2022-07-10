package lfs

import (
	"reflect"
	"testing"
)

func TestGetdivisors(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Getdivisors 100",
			args: args{
				i: 100,
			},
			want: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Getdivisors(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Getdivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
