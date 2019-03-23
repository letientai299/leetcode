package main

import (
	"fmt"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{2, 2, 1, 1, 1, 1, 1, 2, 2}, 1},
		{[]int{2, 3, 1, 1, 1, 2, 1, 1, 4}, 1},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
