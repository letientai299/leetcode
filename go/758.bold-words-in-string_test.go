package main

import "testing"

func Test_boldWords(t *testing.T) {
	tests := []struct {
		words []string
		s     string
		want  string
	}{
		{
			words: []string{"ab", "bc"}, s: "aabcd", want: "a<b>abc</b>d",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := boldWords(tt.words, tt.s); got != tt.want {
				t.Errorf("boldWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
