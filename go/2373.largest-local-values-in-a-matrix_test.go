package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestLocal(t *testing.T) {
	tests := []struct {
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, tt.want, largestLocal(tt.grid), "largestLocal(%v)", tt.grid)
		})
	}
}
