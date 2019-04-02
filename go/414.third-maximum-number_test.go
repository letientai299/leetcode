package main

import (
	"fmt"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{3, 2, 1}, 1},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := thirdMax(tt.nums)
			if got != tt.want {
				t.Errorf("thirdMax(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
