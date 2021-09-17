package main

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		nums []int
		x    int
		name string
		want int
	}{
		{nums: []int{3, 2, 20, 1, 1, 3}, x: 10, want: 5},
		{nums: []int{5, 2, 3, 1, 1}, x: 5, want: 1},
		{nums: []int{5, 6, 7, 8, 9}, x: 4, want: -1},
		{nums: []int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}, x: 134365, want: 16},
		{nums: []int{1, 1, 4, 2, 3}, x: 5, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.nums, tt.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
