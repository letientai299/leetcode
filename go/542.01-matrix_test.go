package main

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{
			mat:
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},

		{
			mat: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1},
			},
			want: [][]int{
				{2, 1, 2, 3, 4},
				{1, 0, 1, 2, 3},
				{2, 1, 2, 2, 3},
				{3, 2, 2, 1, 2},
				{3, 2, 1, 0, 1},
			},
		},

		{
			mat:
			[][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Error("updateMatrix() wrong", tt.want)

				t.Log("want")
				for _, row := range tt.want {
					t.Log(row)
				}

				t.Log("got")
				for _, row := range got {
					t.Log(row)
				}
			}
		})
	}
}
