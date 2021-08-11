package main

import "testing"

func Test_breakPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "zzzz", want: "azzz"},
		{s: "aabaa", want: "aabab"},
		{s: "aaa", want: "aab"},
		{s: "bbb", want: "abb"},
		{s: "abba", want: "aaba"},
		{s: "aa", want: "ab"},
		{s: "a", want: ""},
		{s: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := breakPalindrome(tt.s); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
