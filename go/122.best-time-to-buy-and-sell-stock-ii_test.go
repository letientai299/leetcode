package main

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			nil, 0,
		},
		{
			[]int{1}, 0,
		},
		{
			[]int{1, 1, 1, 1}, 0,
		},
		{
			[]int{1, 2, 2, 1}, 1,
		},
		{
			[]int{3, 2, 2, 1}, 0,
		},
		{
			[]int{3, 2, 1, 2, 3, 4, 5, 6}, 5,
		},
		{
			[]int{2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 1, 5}, 6,
		},
		{
			[]int{2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 1, 2, 2, 5, 5, 6, 6, 6}, 7,
		},
		{
			[]int{3, 3, 3, 3, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 1, 2, 2, 5, 5, 6, 6, 6}, 7,
		},
		{
			[]int{7, 6, 4, 3, 1}, 0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4}, 7,
		},
		{
			[]int{7, 1, 5, 3, 6, 4, 9, 8, 7}, 12,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.prices), func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
