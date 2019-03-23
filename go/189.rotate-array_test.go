package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    4,
			want: []int{3, 4, 5, 6, 1, 2},
		},

		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    2,
			want: []int{5, 6, 1, 2, 3, 4},
		},

		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    1,
			want: []int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    2,
			want: []int{6, 7, 1, 2, 3, 4, 5},
		},

		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},

		{
			nums: []int{},
			k:    1,
			want: []int{},
		},

		{
			nums: []int{},
			k:    1,
			want: []int{},
		},

		{
			nums: []int{1, 2},
			k:    1,
			want: []int{2, 1},
		},

		{
			nums: []int{1, 2, 3, 4},
			k:    4,
			want: []int{1, 2, 3, 4},
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    2,
			want: []int{3, 4, 1, 2},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: []int{4, 5, 6, 1, 2, 3},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			fmt.Printf("Before: %v\n", tt.nums)
			rotate(tt.nums, tt.k)
			fmt.Printf("After:  %v\n", tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
