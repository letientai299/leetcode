package main

import "testing"

func Test_maxCoins(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		want  int
	}{
		{
			piles: []int{2, 4, 1, 2, 7, 8},
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.piles); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
