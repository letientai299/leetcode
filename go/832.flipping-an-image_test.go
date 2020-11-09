package main

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		arr  [][]int
		want [][]int
	}{
		{
			arr: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := flipAndInvertImage(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
