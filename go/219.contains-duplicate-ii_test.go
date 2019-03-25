package main

import (
	"fmt"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 1}, 2, true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := containsNearbyDuplicate(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("containsNearbyDuplicate(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
