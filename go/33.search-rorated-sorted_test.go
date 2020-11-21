package main

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 2,
			want:   6,
		},

		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: 1,
			want:   5,
		},

		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: -1,
			want:   -1,
		},

		{
			nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 8,
			want:   4,
		},

		{
			nums:   []int{1},
			target: 1,
			want:   0,
		},

		{
			nums:   []int{2, 2, 3, 4, 0, 1, 1, 2},
			target: 2,
			want:   0,
		},

		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: 3,
			want:   7,
		},

		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: 4,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
