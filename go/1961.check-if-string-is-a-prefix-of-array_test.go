package main

import "testing"

func Test_isPrefixString(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  bool
	}{
		{
			s:     "aaa",
			words: []string{"a", "a", "aLeetcode", "!"},
			want:  false,
		},

		{
			s:     "a",
			words: []string{"a", "aa", "Leetcode", "!"},
			want:  true,
		},

		{
			s:     "a",
			words: []string{"aa", "aa", "Leetcode", "!"},
			want:  false,
		},

		{
			s:     "IloveLeetcode32",
			words: []string{"aI", "love", "Leetcode", "!"},
			want:  false,
		},

		{
			s:     "IloveLeetcode32",
			words: []string{"I", "love", "Leetcode", "!"},
			want:  false,
		},

		{
			s:     "IloveLeetcode",
			words: []string{"I", "love", "Leetcode", "!"},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrefixString(tt.s, tt.words); got != tt.want {
				t.Errorf("isPrefixString() = %v, want %v", got, tt.want)
			}
		})
	}
}
