package main

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 1, 9, 1, 1, 10, 8}, 21},
		{[]int{1, 2, 1, 9, 2, 1, 10, 6}, 21},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("rob(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
