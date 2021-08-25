package main

import "testing"

func Test_possibleBipartition(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		dislikes [][]int
		want     bool
	}{
		{
			n:        4,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.n, tt.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
