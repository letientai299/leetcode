package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeStones(t *testing.T) {
	tests := []struct {
		name   string
		stones [][]int
		want   int
	}{
		{
			stones: [][]int{
				{3, 2}, {0, 0}, {3, 3}, {2, 1}, {2, 3}, {2, 2}, {0, 2},
			},
			want: 6,
		},

		{
			stones: [][]int{{0, 1}, {1, 0}},
			want:   0,
		},

		{
			stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
			want:   5,
		},

		{
			stones: [][]int{{0, 0}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeStones(tt.stones), "removeStones(%v)", tt.stones)
		})
	}
}
