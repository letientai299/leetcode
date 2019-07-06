package main

import (
	"fmt"
	"testing"
)

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 3, 4, 2, 5}, true},
		{[]int{1, 2, 4, 5, 3}, true},
		{[]int{3, 3, 2, 2}, false},
		{[]int{2, 2, 3, 2, 4}, true},
		{[]int{2, 3, 3, 2, 4}, true},
		{[]int{3}, true},
		{[]int{}, true},
		{[]int{1, 2, 4}, true},
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{1, 4, 2, 3}, true},
		{[]int{2, 4, 2, 3}, true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := checkPossibility(tt.nums)
			if got != tt.want {
				t.Errorf("checkPossibility(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
