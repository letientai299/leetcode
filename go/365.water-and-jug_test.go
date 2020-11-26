package main

import "testing"

func Test_gcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 4, b: 7, want: 1},
		{a: 4, b: 8, want: 4},
		{a: 3, b: 8, want: 1},
		{a: 9, b: 12, want: 3},
		{a: 12, b: 9, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := gcd(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
