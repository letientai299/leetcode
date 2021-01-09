package main

import "testing"

func Test_longestMountain(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{arr: []int{2, 3, 3, 2, 0, 2}, want: 0},
		{arr: []int{2, 3, 1, 2, 3, 4, 5, 6}, want: 3},
		{arr: []int{6, 5, 4, 3, 7, 8, 9}, want: 0},
		{arr: []int{6, 5, 4, 4, 2, 2, 1}, want: 0},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, want: 0},
		{arr: []int{1, 2, 1, 4, 4, 7, 3, 2, 5}, want: 4},
		{arr: []int{2, 2, 1, 4, 4, 7, 3, 2, 5}, want: 4},
		{arr: []int{2, 1, 4, 4, 7, 3, 2, 5}, want: 4},
		{arr: []int{2, 1, 4, 7, 3, 2, 5}, want: 5},
		{arr: []int{1, 1, 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.arr); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
