package main

import "testing"

func Test_minMoves2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 2}, want: 2},
		{nums: []int{0, 0, 3, 4, 6, 8}, want: 15},
		{nums: []int{0, 0, 3, 6, 8}, want: 14},
		{nums: []int{1, 2, 3, 6, 8}, want: 11},
		{nums: []int{1, 2, 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minMoves2(tt.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
