package main

import (
	"fmt"
	"testing"
)

func Test_findShortestSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1,2}, 1},
		{[]int{1,1}, 2},
		{[]int{1}, 1},
		{[]int{1,2,2,1}, 2},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findShortestSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("findShortestSubArray(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
