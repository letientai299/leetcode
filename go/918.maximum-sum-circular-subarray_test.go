package main

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{-2, -3, -1}, want: -1},
		{nums: []int{9, -4, -7, 9}, want: 18},
		{nums: []int{1, -2, 3, -2}, want: 3},
		{nums: []int{5, -3, 5}, want: 10},
		{nums: []int{3, -1, 2, -1}, want: 4},
		{nums: []int{3, -2, 2, -3}, want: 3},
		{nums: []int{1, 2, 3}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
