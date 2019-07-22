package main

import (
	"fmt"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 1}, -1},
		{[]int{1, 2, 1}, 1},
		{[]int{1, 7, 3, 6, 5, 6}, 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Errorf("pivotIndex(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
