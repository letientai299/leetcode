package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	tests := []struct {
		nums [][]int
		r    int
		c    int
		want [][]int
	}{

		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1,
			c: 4,
			want: [][]int{
				{1, 2, 3, 4},
			},
		},

		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 4,
			c: 1,
			want: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},

		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 4,
			c: 2,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v",
			tt.nums, tt.r, tt.c,
		)
		t.Run(testName, func(t *testing.T) {
			if got := matrixReshape(tt.nums, tt.r, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape(%v, %v, %v) = %v, want %v", tt.nums, tt.r, tt.c, got, tt.want)
			}
		})
	}
}
