package main

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {if got := uniquePathsWithObstacles(tt.grid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
