package main

import (
	"fmt"
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{

		{[][]int{
			{1, 1, 1},
			{0, 1, 0},
		}, 10},

		{[][]int{{1}}, 4},
		{[][]int{{1, 0}}, 4},
		{[][]int{{0, 1, 0}}, 4},

		{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}, 4},

		{[][]int{
			{0, 1, 0},
			{0, 1, 0},
			{0, 0, 0},
		}, 6},

		{[][]int{
			{0, 1, 0},
			{1, 1, 0},
			{0, 0, 0},
		}, 8},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.grid,
		)
		t.Run(testName, func(t *testing.T) {
			got := islandPerimeter(tt.grid)
			if got != tt.want {
				t.Errorf("islandPerimeter(%v) = %v, want %v", tt.grid, got, tt.want)
			}
		})
	}
}
