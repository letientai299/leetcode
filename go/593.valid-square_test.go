package main

import "testing"

func Test_validSquare(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
		p3 []int
		p4 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				p1: []int{0, 0},
				p2: []int{3, -4},
				p3: []int{5, 0},
				p4: []int{5, -5},
			},
			want: false,
		},

		{
			args: args{
				p1: []int{0, 0},
				p2: []int{1, 1},
				p3: []int{0, 1},
				p4: []int{0, -1},
			},
			want: false,
		},

		{
			args: args{
				p1: []int{0, 1},
				p2: []int{1, 2},
				p3: []int{0, 2},
				p4: []int{0, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSquare(tt.args.p1, tt.args.p2, tt.args.p3, tt.args.p4); got != tt.want {
				t.Errorf("validSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
