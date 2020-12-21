package main

import "testing"

func Test_minSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{nums: []int{3, 6, 8, 1}, p: 8, want: -1},
		{nums: []int{3, 1, 4, 2}, p: 5, want: 0},
		{nums: []int{3, 1, 4, 2}, p: 2, want: 0},
		{nums: []int{3, 1, 4, 2}, p: 10, want: 0},
		{nums: []int{3, 1, 4, 2}, p: 7, want: 1},
		{nums: []int{3, 1, 4, 2}, p: 3, want: 1},
		{nums: []int{3, 1, 4, 2}, p: 8, want: 1},
		{nums: []int{3, 1, 4, 2}, p: 9, want: 1},
		{nums: []int{3, 1, 4, 2}, p: 6, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.nums, tt.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
