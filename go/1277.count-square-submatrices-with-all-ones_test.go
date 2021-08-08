package main

import "testing"

func Test_countSquares(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 7,
		},

		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: 9,
		},

		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 0},
			},
			want: 7,
		},

		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
