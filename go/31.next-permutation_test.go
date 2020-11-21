package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 1},
			want: []int{2, 1, 1},
		},

		{
			nums: []int{2, 2, 7, 5, 4, 3, 2, 2, 1},
			want: []int{2, 3, 1, 2, 2, 2, 4, 5, 7},
		},

		{
			nums: []int{1, 2, 3, 5, 4},
			want: []int{1, 2, 4, 3, 5},
		},

		{
			nums: []int{1},
			want: []int{1},
		},

		{
			nums: []int{},
			want: []int{},
		},

		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},

		{
			nums: []int{3, 1, 2},
			want: []int{3, 2, 1},
		},

		{
			nums: []int{1, 3, 2},
			want: []int{2, 1, 3},
		},

		{
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			nextPermutation(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
