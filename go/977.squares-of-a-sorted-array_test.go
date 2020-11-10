package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		a    []int
		want []int
	}{
		{
			a:    []int{0, 0, 1, 2},
			want: []int{0, 0, 1, 4},
		},

		{
			a:    []int{1, 2, 3},
			want: []int{1, 4, 9},
		},

		{
			a:    []int{-4, -1, -1},
			want: []int{1, 1, 16},
		},

		{
			a:    []int{-4, -1, -1, 3, 10},
			want: []int{1, 1, 9, 16, 100},
		},

		{
			a:    []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sortedSquares(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
