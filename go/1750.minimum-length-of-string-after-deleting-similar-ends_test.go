package main

import "testing"

func Test_minimumLength(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abcxcbaa", want: 1},
		{s: "ca", want: 2},
		{s: "abccbaa", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minimumLength(tt.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
