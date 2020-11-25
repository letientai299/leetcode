package main

import "testing"

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abcabcabc", want: 2},
		{s: "acababa", want: 5},
		{s: "a", want: 1},
		{s: "", want: 0},
		{s: "aabaa", want: 5},
		{s: "aaaa", want: 4},
		{s: "eceba", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := lengthOfLongestSubstringTwoDistinct(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
