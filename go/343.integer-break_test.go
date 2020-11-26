package main

import "testing"

func Test_integerBreak(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 15, want: 243},
		{n: 11, want: 54},
		{n: 10, want: 36},
		{n: 2, want: 1},
		{n: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := integerBreak(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
