package main

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{
			board: [][]int{
				{-1, -1, -1, -1, 48, 5, -1},
				{12, 29, 13, 9, -1, 2, 32},
				{-1, -1, 21, 7, -1, 12, 49},
				{42, 37, 21, 40, -1, 22, 12},
				{42, -1, 2, -1, -1, -1, 6},
				{39, -1, 35, -1, -1, 39, -1},
				{-1, 36, -1, -1, -1, -1, 5},
			},
			want: 3,
		},

		{
			board: [][]int{
				{-1, -1, 2, 21, -1},
				{16, -1, 24, -1, 4},
				{2, 3, -1, -1, -1},
				{-1, 11, 23, 18, -1},
				{-1, -1, -1, 23, -1},
			},
			want: 2,
		},

		{
			board: [][]int{
				{1, 1, -1},
				{1, 1, 1},
				{-1, 1, 1},
			},
			want: -1,
		},

		{
			board: [][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
