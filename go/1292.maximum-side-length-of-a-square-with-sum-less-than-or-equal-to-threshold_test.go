package main

import "testing"

func Test_maxSideLength(t *testing.T) {
	tests := []struct {
		name      string
		mat       [][]int
		threshold int
		want      int
	}{
		{
			mat: [][]int{
				{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0},
			},
			threshold: 6,
			want:      3,
		},

		{
			mat: [][]int{
				{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2},
			},
			threshold: 4,
			want:      2,
		},

		{
			mat: [][]int{
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
			},
			threshold: 17,
			want:      2,
		},

		{
			mat: [][]int{
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
			},
			threshold: 19,
			want:      3,
		},

		{
			mat: [][]int{
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
			},
			threshold: 2,
			want:      1,
		},

		{
			mat: [][]int{
				{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2},
			},
			threshold: 1,
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSideLength(tt.mat, tt.threshold); got != tt.want {
				t.Errorf("maxSideLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
