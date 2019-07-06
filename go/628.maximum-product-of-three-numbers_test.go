package main

import (
	"fmt"
	"testing"
)

func Test_maximumProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1000, 1000, 1000}, 1000000000},
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-5, 1, 2, 3}, 6},
		{[]int{-5, 0, 1, 2, 3}, 6},
		{[]int{-5, -2, 0, 1, 2, 3}, 30},
		{[]int{-5, -2, -1, 1, 2, 3}, 30},
		{[]int{-5, -2, -1}, -10},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := maximumProduct(tt.nums)
			if got != tt.want {
				t.Errorf("maximumProduct(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
