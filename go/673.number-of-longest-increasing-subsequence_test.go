package main

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 1, 2, 3, 1, 2, 3}, want: 10},
		{nums: []int{2, 3, 2, 3, 2}, want: 3},
		{nums: []int{2, 2, 2, 2, 2}, want: 5},
		{nums: []int{1, 2, 4, 3, 5, 4, 7, 2}, want: 3},
		{nums: []int{1, 3, 5, 4, 7}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
