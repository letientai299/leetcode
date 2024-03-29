package main

import (
	"testing"
)

func Test_checkSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 3, 0, 6},
			k:    6,
			want: true,
		},

		{
			nums: []int{1, 2, 12},
			k:    6,
			want: false,
		},

		{
			nums: []int{23, 0, 0},
			k:    6,
			want: true,
		},

		{
			nums: []int{0, 1, 0, 3, 0, 4, 0, 4, 0},
			k:    5,
			want: false,
		},

		{
			nums: []int{23, 2, 4, 6, 7},
			k:    6,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := checkSubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
