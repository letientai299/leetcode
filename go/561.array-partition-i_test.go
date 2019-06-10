package main

import (
	"fmt"
	"testing"
)

func Test_arrayPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3, 2}, 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := arrayPairSum(tt.nums)
			if got != tt.want {
				t.Errorf("arrayPairSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
