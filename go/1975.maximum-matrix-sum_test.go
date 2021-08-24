package main

import "testing"

func Test_maxMatrixSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int64
	}{
		{
			matrix: [][]int{
				{2, 9, 3},
				{5, 4, -4},
				{1, 7, 1},
			},
			want: 34,
		},

		{
			matrix: [][]int{
				{-1, 0, -1},
				{-2, 1, 3},
				{3, 2, 2},
			},
			want: 15,
		},

		{
			matrix: [][]int{
				{1, 2, 3},
				{-1, -2, -3},
				{1, 2, 3},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMatrixSum(tt.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
