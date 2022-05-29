package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countHillValley(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{2, 4, 1, 1, 6, 5}, want: 3},
		{nums: []int{1, 2, 3}, want: 0},
		{nums: []int{1, 2, 1}, want: 1},
		{nums: []int{1}, want: 0},
		{nums: []int{11, 1, 1, 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countHillValley(tt.nums), "countHillValley(%v)", tt.nums)
		})
	}
}
