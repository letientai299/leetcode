package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{0},
			want: 1,
		},

		{
			nums: []int{1},
			want: 1,
		},

		{
			nums: []int{1, 0, 1, 1, 0},
			want: 4,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 4,
		},

		{
			nums: []int{0, 1, 0, 1, 0, 1, 0, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxConsecutiveOnes(tt.nums), "findMaxConsecutiveOnes(%v)", tt.nums)
		})
	}
}
