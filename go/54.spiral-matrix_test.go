package main

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want []int
	}{
		{
			mat: [][]int{
				{1},
				{2},
				{3},
			},
			want: []int{1, 2, 3},
		},
		{
			mat: [][]int{
				{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 4, 3},
		},

		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
