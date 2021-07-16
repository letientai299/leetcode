package main

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "10010100", want: 3},
		{s: "00100", want: 2},
		{s: "0100", want: 1},
		{s: "0000", want: 2},
		{s: "01010", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minOperations(tt.s); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
