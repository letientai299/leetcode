package main

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			nums:   []int{5, 0, 2, 5, -5, 4, -5, 1, -1},
			target: -5,
			want: [][]int{
				{-5, -5, 0, 5},
				{-5, -5, 1, 4},
				{-5, -1, 0, 1},
			},
		},

		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},

		{
			nums:   []int{-3, -1, 0, 2, 4, 5},
			target: 0,
			want: [][]int{
				{-3, -1, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fourSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
