package main

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "cbacdcbc", want: "acdb"},
		{s: "mitnlruhznjfyzmtmfnstsxwktxlboxutbic", want: "ilrhjfyzmnstwkboxuc"},
		{s: "abacb", want: "abc"},
		{s: "cdadabcc", want: "adbc"},
		{s: "bcacb", want: "acb"},
		{s: "bcabc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
