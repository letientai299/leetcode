package main

import (
	"fmt"
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1, 0, 2}, 2},
		{[]int{3, 2, 1, 0}, 1},
		{[]int{1, 1, 1, 1}, 1},
		{[]int{1, 2, 1, 1, 1}, 2},
		{[]int{1, 3, 5, 4, 7}, 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findLengthOfLCIS(tt.nums)
			if got != tt.want {
				t.Errorf("findLengthOfLCIS(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
