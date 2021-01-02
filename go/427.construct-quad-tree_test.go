package main

import (
	"testing"
)

func Test_construct(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
		},

		{
			grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := construct(tt.grid)
			t.Log(got)
		})
	}
}
