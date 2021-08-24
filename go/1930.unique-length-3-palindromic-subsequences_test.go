package main

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "bbcbaba", want: 4},
		{s: "aabca", want: 3},
		{s: "adc", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
