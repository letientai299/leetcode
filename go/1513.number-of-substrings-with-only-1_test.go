package main

import "testing"

func Test_numSub(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "110011101100", want: 12},
		{s: "0011101100", want: 9},
		{s: "001110100", want: 7},
		{s: "11101", want: 7},
		{s: "111", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numSub(tt.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
