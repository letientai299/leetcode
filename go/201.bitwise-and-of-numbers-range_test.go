package main

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{m: 5, n: 7, want: 4},
		{m: 5, n: 5, want: 5},
		{m: 1, n: 3, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.m, tt.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
