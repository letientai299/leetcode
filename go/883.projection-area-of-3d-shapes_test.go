package main

import "testing"

func Test_projectionArea(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := projectionArea(tt.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
