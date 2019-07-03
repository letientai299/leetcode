package main

import (
	"fmt"
	"testing"
)

func Test_findLHS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 1, 1}, 0},
		{[]int{1, 3, 1}, 0},
		{[]int{1, 2, 1}, 3},
		{[]int{1, 2, 1, 4}, 3},
		{[]int{1, 2, 1, 2}, 4},
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findLHS(tt.nums)
			if got != tt.want {
				t.Errorf("findLHS(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
