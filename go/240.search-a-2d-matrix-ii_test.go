package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	big1 := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	big2 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	tests := []struct {
		name string
		mat  [][]int
		t    int
		want bool
	}{
		{
			mat: [][]int{
				{3, 3, 8, 13, 13, 18},
				{4, 5, 11, 13, 18, 20},
				{9, 9, 14, 15, 23, 23},
				{13, 18, 22, 22, 25, 27},
				{18, 22, 23, 28, 30, 33},
				{21, 25, 28, 30, 35, 35},
				{24, 25, 33, 36, 37, 40},
			},
			t:    21,
			want: true,
		},
		{mat: big2, t: 5, want: true},

		{mat: big1, t: 25, want: false},
		{mat: big1, t: 23, want: true},
		{mat: big1, t: 22, want: true},
		{mat: big1, t: 18, want: true},
		{mat: big1, t: 17, want: true},
		{mat: big1, t: 1000, want: false},
		{mat: big1, t: 0, want: false},
		{mat: big1, t: 2, want: true},

		{mat: [][]int{{1, 2}}, t: 2, want: true},
		{mat: [][]int{{1}}, t: 2, want: false},
		{mat: [][]int{{1}}, t: 1, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.mat, tt.t); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
