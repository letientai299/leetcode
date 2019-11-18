package main

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{

		{
			nums: []int{1},
			k:    0,
			want: 0,
		},

		{
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},

		{
			nums: []int{1, -1, 1, -1},
			k:    0,
			want: 4,
		},

		{
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},

		{
			nums: []int{1, 1, 1},
			k:    1,
			want: 3,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := subarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("subarraySum(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
