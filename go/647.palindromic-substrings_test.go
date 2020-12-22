package main

import "testing"

func Test_countSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abbabbbaca", want: 19},
		{s: "aba", want: 4},
		{s: "aaaa", want: 10},
		{s: "abc", want: 3},
		{s: "aabc", want: 5},
		{
			s:    "cbcdbcaadccaadcdddadcccadbaabdacccdcddccdcdaaadcbcabadcbdcdbacadbacddcccdbb",
			want: 123,
		},
		{s: "aabcbaabcbc", want: 21},
		{s: "abbaabcbc", want: 15},
		{s: "aaaaa", want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := countSubstrings(tt.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
