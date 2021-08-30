package main

import "testing"

func Test_numSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		goal int
		want int
	}{
		{goal: 0, nums: []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, want: 27},
		{goal: 2, nums: []int{1, 0, 1, 0, 1}, want: 4},
		{goal: 0, nums: []int{0, 0, 0, 0, 0}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(tt.nums, tt.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
