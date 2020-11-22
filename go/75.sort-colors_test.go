package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},

		{
			nums: []int{1, 2, 0},
			want: []int{0, 1, 2},
		},


		{
			nums: []int{1, 0, 0},
			want: []int{0, 0, 1},
		},

		{
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},

		{
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},

		{
			nums: []int{2},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sortColors(tt.nums)
			assert.EqualValues(t, tt.want, tt.nums)
		})
	}
}
