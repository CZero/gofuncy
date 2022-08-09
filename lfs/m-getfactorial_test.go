package lfs

import "testing"

func TestGetFactorial(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Factorial 5",
			args: args{5},
			want: 120,
		},
		{
			name: "Factorial 6",
			args: args{6},
			want: 720,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFactorial(tt.args.input); got != tt.want {
				t.Errorf("GetFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
