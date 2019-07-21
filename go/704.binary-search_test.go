package main

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{

		{
			nums:   []int{},
			target: 0,
			want:   -1,
		},

		{
			nums:   nil,
			target: 0,
			want:   -1,
		},

		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},

		{
			nums:   []int{1, 2, 3},
			target: 0,
			want:   -1,
		},


		{
			nums:   []int{0, 1, 2},
			target: 0,
			want:   0,
		},

		{
			nums:   []int{0, 1, 2, 4},
			target: 3,
			want:   -1,
		},

		{
			nums:   []int{0, 1, 2, 4},
			target: 2,
			want:   2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.target,
		)
		t.Run(testName, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
