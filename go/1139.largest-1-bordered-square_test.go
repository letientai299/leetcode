package main

import "testing"

func Test_largest1BorderedSquare(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 1, 1, 1, 0},
				{1, 1, 0, 1, 1, 0},
				{1, 1, 0, 1, 0, 1},
				{1, 1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			want: 16,
		},

		{
			grid: [][]int{
				{0, 0},
				{1, 0},
			},
			want: 1,
		},

		{
			grid: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: 16,
		},

		{
			grid: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 0},
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: 9,
		},

		{
			grid: [][]int{
				{1, 1, 1, 0},
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: 9,
		},

		{
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: 9,
		},

		{
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 9,
		},

		{grid: [][]int{{1, 1}, {0, 0}}, want: 1},
		{grid: [][]int{{1}}, want: 1},
		{grid: [][]int{{0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largest1BorderedSquare(tt.grid); got != tt.want {
				t.Errorf("largest1BorderedSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
