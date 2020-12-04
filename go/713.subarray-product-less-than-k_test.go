package main

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{nums: []int{10, 5, 2, 6}, k: 0, want: 0},
		{nums: []int{1, 2, 3, 4}, k: 6, want: 5},
		{nums: []int{2, 2, 2, 2, 2}, k: 9, want: 12},
		{nums: []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, k: 19, want: 18},
		{nums: []int{10, 5, 2, 6}, k: 100, want: 8},
		{nums: []int{1, 2, 3, 4}, k: 25, want: 10},
		{nums: []int{1, 2, 3, 4}, k: 12, want: 7},
		{nums: []int{1, 2, 3, 4}, k: 7, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.nums, tt.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
