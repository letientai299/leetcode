package main

import (
	"fmt"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{0, 1, 1, 1, 1, 1}, 5},
		{[]int{0, 1, 1, 1, 1, 1, 0}, 5},
		{[]int{0, 0}, 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.nums)
			if got != tt.want {
				t.Errorf("findMaxConsecutiveOnes(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
