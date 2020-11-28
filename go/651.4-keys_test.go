package main

import "testing"

func Test_maxA(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 11, want: 24},
		{n: 10, want: 18},
		{n: 4, want: 4},
		{n: 5, want: 5},
		{n: 8, want: 12},
		{n: 7, want: 9},
		{n: 6, want: 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxA(tt.n); got != tt.want {
				t.Errorf("maxA() = %v, want %v", got, tt.want)
			}
		})
	}
}
