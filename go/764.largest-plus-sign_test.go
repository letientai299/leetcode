package main

import "testing"

func Test_orderOfLargestPlusSign(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		mines [][]int
		want  int
	}{
		{n: 2, mines: [][]int{{0, 0}, {0, 1}, {1, 0}}, want: 1},
		{n: 5, mines: [][]int{{4, 3}}, want: 3},
		{n: 5, mines: [][]int{{4, 2}}, want: 2},
		{n: 1, mines: [][]int{{0, 0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.n, tt.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
