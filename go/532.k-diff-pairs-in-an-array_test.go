package main

import (
	"fmt"
	"testing"
)

func Test_findPairs(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 1, 1, 2, 1}, k: 1, want: 1},
		{nums: []int{1, 2, 3, 4, 5}, k: 1, want: 4},
		{nums: []int{1, 2, 3, 4, 5}, k: 0, want: 0},
		{nums: []int{1, 2, 3, 4, 1}, k: 0, want: 1},
		{nums: []int{3, 1, 4, 1, 5}, k: 2, want: 2},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := findPairs(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findPairs(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
