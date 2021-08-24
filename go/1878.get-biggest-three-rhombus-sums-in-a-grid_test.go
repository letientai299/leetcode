package main

import (
	"reflect"
	"testing"
)

func Test_getBiggestThree(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want []int
	}{
		{
			grid: [][]int{
				{228, 4, 5, 1, 3},
				{3, 3, 4, 2, 3},
				{20, 30, 200, 40, 10},
				{1, 5, 5, 4, 1},
				{4, 3, 2, 2, 5},
			},
			want: []int{228, 216, 211},
		},

		{
			grid: [][]int{{7, 7, 7}},
			want: []int{7},
		},

		{
			grid: [][]int{
				{3, 4, 5, 1, 3},
				{3, 3, 4, 2, 3},
				{20, 30, 200, 40, 10},
				{1, 5, 5, 4, 1},
				{4, 3, 2, 2, 5},
			},
			want: []int{228, 216, 211},
		},

		{
			grid: [][]int{
				{3, 4, 5, 1, 3},
				{3, 3, 4, 2, 3},
				{20, 30, 200, 40, 10},
				{1, 5, 5, 4, 1},
				{4, 3, 2, 2, 5000},
			},
			want: []int{5000, 228, 216},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBiggestThree(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBiggestThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
