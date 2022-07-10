package lfs

import "testing"

func TestGettrianglenumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Triangle 1",
			args: args{
				1,
			},
			want: 1,
		},
		{
			name: "Triangle 2",
			args: args{
				2,
			},
			want: 3,
		},
		{
			name: "Triangle 3",
			args: args{
				3,
			},
			want: 6,
		},
		{
			name: "Triangle 4",
			args: args{
				4,
			},
			want: 10,
		},
		{
			name: "Triangle 5",
			args: args{
				5,
			},
			want: 15,
		},
		{
			name: "Triangle 6",
			args: args{
				6,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gettrianglenumber(tt.args.n); got != tt.want {
				t.Errorf("Gettrianglenumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
