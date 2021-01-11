package main

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
		{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, want: 6},
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
