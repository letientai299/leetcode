package main

import "testing"

func Test_numWays(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "10101", want: 4},
		{s: "1101", want: 2},
		{s: "111", want: 1},
		{s: "11", want: 0},
		{s: "0000", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numWays(tt.s); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
