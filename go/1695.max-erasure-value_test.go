package main

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{10000, 1, 10000, 1, 1, 1, 1, 1, 1}, want: 10001},
		{nums: []int{4, 2, 4, 5, 6}, want: 17},
		{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
