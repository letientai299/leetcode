package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		b         []int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			b:         []int{4, 8},
			want:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},

		{
			intervals: [][]int{{1, 3}, {6, 9}},
			b:         []int{10, 11},
			want:      [][]int{{1, 3}, {6, 9}, {10, 11}},
		},

		{
			intervals: [][]int{{1, 3}, {6, 9}},
			b:         []int{-1, 0},
			want:      [][]int{{-1, 0}, {1, 3}, {6, 9}},
		},

		{
			intervals: [][]int{{1, 3}, {6, 9}},
			b:         []int{2, 6},
			want:      [][]int{{1, 9}},
		},

		{
			intervals: [][]int{{1, 3}, {6, 9}},
			b:         []int{2, 5},
			want:      [][]int{{1, 5}, {6, 9}},
		},

		{
			intervals: [][]int{{1, 3}, {6, 9}},
			b:         []int{4, 5},
			want:      [][]int{{1, 3}, {4, 5}, {6, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.intervals, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
