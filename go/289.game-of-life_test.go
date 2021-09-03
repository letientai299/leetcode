package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		{
			board: [][]int{
				{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.board)
			assert.Equal(t, tt.want, tt.board)
		})
	}
}
