package main

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{arr: []int{1, 2, 3, 4, 5, 0}, want: 1},
		{arr: []int{1, 2, 0, 3}, want: 2},
		{arr: []int{1, 0, 3, 2, 4}, want: 3},
		{arr: []int{1, 0, 2, 3, 4}, want: 4},
		{arr: []int{4, 3, 2, 1, 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
