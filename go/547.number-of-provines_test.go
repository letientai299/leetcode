package main

import "testing"

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		want int
	}{

		{
			m: [][]int{
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
			want: 8,
		},

		{
			m: [][]int{
				{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
			},
			want: 3,
		},

		{
			m: [][]int{
				{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.m); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}