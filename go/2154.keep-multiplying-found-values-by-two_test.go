package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFinalValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		ori  int
		want int
	}{
		{
			nums: []int{5, 3, 6, 1, 12},
			ori:  3,
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findFinalValue(tt.nums, tt.ori), "findFinalValue(%v, %v)", tt.nums, tt.ori)
		})
	}
}
