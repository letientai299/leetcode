package main

import (
	"fmt"
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 2, 2, 2}, want: 4},
		{nums: []int{1, 2, 3}, want: 0},
		{nums: []int{1}, want: 0},
		{nums: []int{2, 1}, want: 2},
		{nums: []int{3, 2, 1}, want: 3},
		{nums: []int{3, 2, 1, 10}, want: 3},
		{nums: []int{2, 6, 4, 8, 10, 9, 15}, want: 5},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findUnsortedSubarray(tt.nums)
			if got != tt.want {
				t.Errorf("findUnsortedSubarray(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
