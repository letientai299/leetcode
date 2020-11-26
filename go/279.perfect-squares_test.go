package main

import "testing"

func Test_numSquares(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 16, want: 1},
		{n: 15, want: 4},
		{n: 14, want: 3},
		{n: 12, want: 3},
		{n: 13, want: 2},
		{n: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
