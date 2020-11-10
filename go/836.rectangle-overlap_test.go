package main

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	tests := []struct {
		r1   []int
		r2   []int
		want bool
	}{
		{
			r1:   []int{7, 8, 13, 15},
			r2:   []int{10, 8, 12, 20},
			want: true,
		},

		{
			r1:   []int{5, 15, 8, 18},
			r2:   []int{0, 3, 7, 9},
			want: false,
		},

		{
			r1:   []int{11, 12, 13, 13},
			r2:   []int{17, 1, 17, 19},
			want: false,
		},

		{
			r1:   []int{7, 8, 13, 15},
			r2:   []int{10, 8, 12, 20},
			want: true,
		},

		{
			r1:   []int{0, 0, 2, 2},
			r2:   []int{1, 1, 3, 3},
			want: true,
		},

		{
			r1:   []int{0, 0, 1, 1},
			r2:   []int{1, 0, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isRectangleOverlap(tt.r1, tt.r2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
