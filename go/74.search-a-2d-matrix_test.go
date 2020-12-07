package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		t    int
		want bool
	}{
		{
			mat: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			t:    10,
			want: true,
		},

		{
			mat: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			t:    100,
			want: false,
		},

		{
			mat: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			t:    0,
			want: false,
		},

		{
			mat: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			t:    3,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.mat, tt.t); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
