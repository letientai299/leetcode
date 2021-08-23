package main

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder498(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want []int
	}{
		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},

		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: []int{1, 2, 4, 5, 3, 6},
		},

		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 10, 11, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder498(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
