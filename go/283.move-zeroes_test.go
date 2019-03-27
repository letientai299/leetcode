package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			nums: nil,
			want: nil,
		},

		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{0, 0},
			want: []int{0, 0},
		},
		{
			nums: []int{1, 0},
			want: []int{1, 0},
		},
		{
			nums: []int{1, 0, 1},
			want: []int{1, 1, 0},
		},
		{
			nums: []int{0, 0, 1},
			want: []int{1, 0, 0},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			moveZeroes(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
