package main

import "testing"

func Test_numTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 4, want: 14},
		{n: 5, want: 42},
		{n: 3, want: 5},
		{n: 2, want: 2},
		{n: 1, want: 1},
		{n: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
