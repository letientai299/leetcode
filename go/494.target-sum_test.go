package main

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 1}, k: 0, want: 2},
		{nums: []int{0, 0, 0, 0, 0, 1}, k: 1, want: 32},
		{nums: []int{0, 1}, k: 1, want: 2},
		{nums: []int{1, 0}, k: 1, want: 2},
		{nums: []int{1}, k: 1, want: 1},
		{nums: []int{1, 1, 1, 1, 1}, k: 3, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.nums, tt.k); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
