package main

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{1, 2},
				{1, 1},
			},
			want: 3,
		},

		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
