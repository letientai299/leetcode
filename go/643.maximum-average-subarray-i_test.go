package main

import (
	"fmt"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k: 4,
			want: 12.75,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findMaxAverage(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
