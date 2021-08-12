package main

import "testing"

func Test_canPartitionKSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 1, 1, 1, 2, 2, 2, 2},
			k:    4,
			want: true,
		},

		{
			nums: []int{2, 2, 2, 2, 3, 4, 5},
			k:    4,
			want: false,
		},

		{
			nums: []int{4, 3, 2, 3, 5, 2, 1},
			k:    4,
			want: true,
		},

		{
			nums: []int{1, 2, 3, 4},
			k:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.nums, tt.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
