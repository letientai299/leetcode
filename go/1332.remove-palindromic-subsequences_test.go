package main

import "testing"

func Test_removePalindromeSub(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "ababb", want: 2},
		{s: "abb", want: 2},
		{s: "abcbab", want: 2},
		{s: "caba", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := removePalindromeSub(tt.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
