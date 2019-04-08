package main

import (
	"fmt"
	"testing"
)

func Test_minMoves(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 3}, 2},
		{[]int{1, 1, 5}, 4},
		{[]int{1, 2, 7}, 7},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := minMoves(tt.nums)
			if got != tt.want {
				t.Errorf("minMoves(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
