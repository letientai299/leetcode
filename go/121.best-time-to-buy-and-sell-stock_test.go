package main

import (
	"fmt"
	"testing"
)

func Test_maxProfit_1(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			nil, 0,
		},
		{
			[]int{7, 6, 4, 3, 1}, 0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4}, 5,
		},
		{
			[]int{7, 1, 5, 3, 6, 4, 9, 8, 7}, 8,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.prices), func(t *testing.T) {
			if got := maxProfit_1(tt.prices); got != tt.want {
				t.Errorf("maxProfit_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
