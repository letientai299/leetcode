package main

import "testing"

func Test_numEnclaves(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		{
			mat: [][]int{
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
				{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
				{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.mat); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
