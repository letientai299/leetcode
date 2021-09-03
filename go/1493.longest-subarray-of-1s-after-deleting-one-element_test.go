package main

import "testing"

func Test_longestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{0, 0, 1, 1, 1, 0, 1, 1, 0}, want: 5},
		{nums: []int{0, 0, 1, 1, 1, 0, 0, 1, 1, 0}, want: 3},
		{nums: []int{1, 1, 1}, want: 2},
		{nums: []int{0, 0, 0}, want: 0},
		{nums: []int{0, 1, 0, 0}, want: 1},
		{nums: []int{0, 1, 1, 0}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
