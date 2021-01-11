package main

import (
	"testing"
)

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 2, want: 2},
		{n: 1, want: 1},
		{n: 3, want: 5},
		{n: 4, want: 14},
		{n: 5, want: 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.n)
			if tt.want != len(got) {
				t.Errorf("generateTrees() = len %d, want %v, trees=%v", len(got), tt.want, got)
			}
		})
	}
}
