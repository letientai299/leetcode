package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{3, 1, 2, 2, 2, 1, 3},
			k:    2,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countPairs(tt.nums, tt.k), "countPairs(%v, %v)", tt.nums, tt.k)
		})
	}
}
