package main

import "testing"

func Test_countSubIslands(t *testing.T) {
	tests := []struct {
		name  string
		grid1 [][]int
		grid2 [][]int
		want  int
	}{
		{
			grid1: [][]int{
				{1, 1, 1, 1, 0, 0},
				{1, 1, 0, 1, 0, 0},
				{1, 0, 0, 1, 1, 1},
				{1, 1, 1, 0, 0, 1},
				{1, 1, 1, 1, 1, 0},
				{1, 0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 1, 1},
				{1, 0, 0, 0, 1, 0},
				{1, 1, 1, 1, 1, 0},
			},
			grid2: [][]int{
				{1, 1, 1, 1, 0, 1},
				{0, 0, 1, 0, 1, 0},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 0},
				{0, 1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1, 1},
				{1, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0, 0},
			},
			want: 0,
		},

		{
			grid1: [][]int{
				{1, 1, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 1, 1},
			},
			grid2: [][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 1, 1},
				{0, 1, 0, 0, 0},
				{1, 0, 1, 1, 0},
				{0, 1, 0, 1, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubIslands(tt.grid1, tt.grid2); got != tt.want {
				t.Errorf("countSubIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
