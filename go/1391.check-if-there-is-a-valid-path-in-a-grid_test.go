package main

import "testing"

func Test_hasValidPath(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want bool
	}{
		{
			grid: [][]int{
				{4, 3, 5},
				{6, 5, 2},
			},
			want: false,
		},

		{
			grid: [][]int{
				{6, 1, 3},
				{4, 1, 5},
			},
			want: true,
		},

		{
			grid: [][]int{
				{4, 1},
				{6, 1},
			},
			want: true,
		},

		{
			grid: [][]int{
				{2, 6},
			},
			want: false,
		},

		{
			grid: [][]int{
				{1, 2, 1},
				{1, 2, 1},
			},
			want: false,
		},

		{
			grid: [][]int{
				{1, 1, 2},
			},
			want: false,
		},

		{
			grid: [][]int{
				{2},
				{2},
				{2},
				{2},
				{2},
				{6},
			},
			want: true,
		},

		{
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 3},
			},
			want: true,
		},

		{
			grid: [][]int{
				{2, 4, 3},
				{6, 5, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasValidPath(tt.grid); got != tt.want {
				t.Errorf("hasValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
