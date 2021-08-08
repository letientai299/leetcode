package main

import "testing"

func Test_maxNonOverlapping(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{0, 0, 0},
			target: 0,
			want:   3,
		},

		{
			nums:   []int{0, 0, 0},
			target: 3,
			want:   0,
		},

		{
			nums:   []int{-1, 3, 5, 1, 4, 2, -9},
			target: 6,
			want:   2,
		},

		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 2,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNonOverlapping(tt.nums, tt.target); got != tt.want {
				t.Errorf("maxNonOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
