package main

import (
	"fmt"
	"testing"
)

func Test_missingNumber_268(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{1}, 0},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := missingNumber_268(tt.nums)
			if got != tt.want {
				t.Errorf("missingNumber(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
