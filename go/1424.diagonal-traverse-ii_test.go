package main

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want []int
	}{
		{
			nums: [][]int{
				{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16},
			},
			want: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
