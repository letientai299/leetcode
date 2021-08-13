package main

import "testing"

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{1, 0, 1, 0, 0, 1, 1, 1, 0},
				{1, 1, 1, 0, 0, 0, 0, 0, 1},
				{0, 0, 1, 1, 0, 0, 0, 1, 1},
				{0, 1, 1, 0, 0, 1, 0, 0, 1},
				{1, 1, 0, 1, 1, 0, 0, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 0, 1, 0},
				{1, 1, 1, 0, 1, 0, 0, 0, 1},
				{0, 1, 1, 1, 1, 0, 0, 1, 0},
				{1, 0, 0, 1, 1, 1, 0, 0, 0},
			},
			want: 4,
		},

		{
			matrix: [][]byte{
				{0},
			},
			want: 0,
		},

		{
			matrix: [][]byte{
				{0, 1},
				{1, 0},
			},
			want: 1,
		},

		{
			matrix: [][]byte{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for y := 0; y < len(tt.matrix); y++ {
				for x := 0; x < len(tt.matrix[0]); x++ {
					tt.matrix[y][x] += '0'
				}
			}
			if got := maximalSquare(tt.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
