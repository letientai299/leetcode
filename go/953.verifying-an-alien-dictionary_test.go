package main

import "testing"

func Test_isAlienSorted(t *testing.T) {
	tests := []struct {
		words []string
		order string
		want  bool
	}{
		{
			words: []string{"world", "or"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  true,
		},

		{
			words: []string{"hello", "hello"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  true,
		},

		{
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  false,
		},

		{
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isAlienSorted(tt.words, tt.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
