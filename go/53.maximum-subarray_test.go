package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{[]int{6}, 6},
		{[]int{6, -1}, 6},
		{[]int{6, 1}, 7},
		{[]int{}, 0},
		{[]int{3, 2, 1, -7, 4, 5, -2}, 9},
		{[]int{3, 2, 1, -2, 4, 5, -2}, 13},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
