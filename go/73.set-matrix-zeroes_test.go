package main

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
	}{
		{
			mat: [][]int{
				{1, 2, 3, 4},
				{5, 0, 7, 8},
				{0, 10, 11, 12},
				{13, 14, 15, 0},
			},
		},

		{
			mat: [][]int{
				{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
			},
		},
		{
			mat: [][]int{
				{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("------")
			for _, v := range tt.mat {
				fmt.Println(v)
			}
			fmt.Println("------")
			setZeroes(tt.mat)
			for _, v := range tt.mat {
				fmt.Println(v)
			}
		})
	}
}
