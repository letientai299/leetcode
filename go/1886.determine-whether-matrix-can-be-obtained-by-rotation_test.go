package main

import "testing"

func Test_findRotation(t *testing.T) {
	tests := []struct {
		name   string
		mat    [][]int
		target [][]int
		want   bool
	}{
		{
			mat: [][]int{
				{0, 0},
				{0, 1},
			},
			target: [][]int{
				{0, 0},
				{1, 0},
			},
			want: true,
		},

		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			target: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: true,
		},

		{
			mat:    [][]int{{0, 1}, {1, 0}},
			target: [][]int{{1, 0}, {0, 1}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotation(tt.mat, tt.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
