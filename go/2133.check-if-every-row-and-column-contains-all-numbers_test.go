package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkValid(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 2, 2, 4, 5, 6, 6},
				{2, 2, 4, 5, 6, 6, 1},
				{2, 4, 5, 6, 6, 1, 2},
				{4, 5, 6, 6, 1, 2, 2},
				{5, 6, 6, 1, 2, 2, 4},
				{6, 6, 1, 2, 2, 4, 5},
				{6, 1, 2, 2, 4, 5, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkValid(tt.matrix), "checkValid(%v)", tt.matrix)
		})
	}
}
