package main

import "testing"

func Test_numFactoredBinaryTrees(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{arr: []int{18, 3, 6, 2}, want: 12},
		{arr: []int{2, 4}, want: 3},
		{arr: []int{2, 4, 5, 10}, want: 7},
		{arr: []int{10, 2, 5, 4}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFactoredBinaryTrees(tt.arr); got != tt.want {
				t.Errorf("numFactoredBinaryTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
