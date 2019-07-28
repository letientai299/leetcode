package main

import (
	"fmt"
	"testing"
)

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 6, 3, 0}, want: 1},
		{nums: []int{1, 2, 3, 4}, want: -1},
		{nums: []int{1, 2}, want: 1},
		{nums: []int{2, 1, 0}, want: 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := dominantIndex(tt.nums)
			if got != tt.want {
				t.Errorf("dominantIndex(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
