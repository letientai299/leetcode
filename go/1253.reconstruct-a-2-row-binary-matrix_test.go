package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reconstructMatrix(t *testing.T) {
	tests := []struct {
		name   string
		upper  int
		lower  int
		colsum []int
		want   [][]int
	}{
		{
			upper:  1,
			lower:  3,
			colsum: []int{1, 2, 1},
			want: [][]int{
				{0, 1, 0},
				{1, 1, 1},
			},
		},

		{
			upper:  0,
			lower:  3,
			colsum: []int{1, 1, 1},
			want: [][]int{
				{0, 0, 0},
				{1, 1, 1},
			},
		},

		{
			upper:  1,
			lower:  2,
			colsum: []int{1, 1, 1},
			want: [][]int{
				{0, 1, 0},
				{1, 0, 1},
			},
		},

		{
			upper:  4,
			lower:  1,
			colsum: []int{1, 1, 1},
			want:   [][]int{},
		},

		{
			upper:  2,
			lower:  1,
			colsum: []int{1, 1, 1},
			want: [][]int{
				{1, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructMatrix(tt.upper, tt.lower, tt.colsum); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
