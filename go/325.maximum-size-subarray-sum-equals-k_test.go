package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArrayLen(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, -1, 5, -2, 3},
			k:    3,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubArrayLen(tt.nums, tt.k), "maxSubArrayLen(%v, %v)", tt.nums, tt.k)
		})
	}
}
