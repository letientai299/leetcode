package main

import "testing"

func Test_orangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{2, 1},
				{1, 1},
			},
			want: 2,
		},

		{
			grid: [][]int{
				{2, 1, 1}, {1, 1, 1}, {0, 1, 2},
			},
			want: 2,
		},

		{
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},

		{
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},

		{
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
