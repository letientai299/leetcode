package main

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{m: 13, n: 23, want: 548354040},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
