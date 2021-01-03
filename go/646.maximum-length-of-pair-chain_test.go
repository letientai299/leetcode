package main

import "testing"

func Test_findLongestChain(t *testing.T) {
	tests := []struct {
		name  string
		pairs [][]int
		want  int
	}{
		{
			pairs: [][]int{
				{7, 9}, {4, 5}, {7, 9}, {-7, -1}, {0, 10}, {3, 10}, {3, 6}, {2, 3},
			},
			want: 4,
		},

		{
			pairs: [][]int{
				{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
