package main

import "testing"

func Test_maxRepOpt1(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{text: "aaabaaabbaaaaaaa", want: 8},
		{text: "abcdfaaaa", want: 5},
		{text: "adaaabaaa", want: 7},
		{text: "aaabaaa", want: 6},
		{text: "aaabaaaca", want: 7},
		{text: "aaabbaaa", want: 4},
		{text: "abcdef", want: 1},
		{text: "aaaa", want: 4},
		{text: "ababa", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if got := maxRepOpt1(tt.text); got != tt.want {
				t.Errorf("maxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
