package main

import (
	"fmt"
	"testing"
)

func Test_findDuplicate287(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{

		{
			nums: []int{2, 1, 2},
			want: 2,
		},

		{
			nums: []int{2, 1, 1},
			want: 1,
		},

		{
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},

		{
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findDuplicate287(tt.nums)
			if got != tt.want {
				t.Errorf("findDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
