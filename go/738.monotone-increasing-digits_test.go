package main

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 21345267, want: 19999999},
		{n: 332, want: 299},
		{n: 1034, want: 999},
		{n: 1234, want: 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.n); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
