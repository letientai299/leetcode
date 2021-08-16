package main

import "testing"

func Test_numMatchingSubseq(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  int
	}{
		{
			s:     "abcde",
			words: []string{"a", "bb", "acd", "ace"},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMatchingSubseq(tt.s, tt.words); got != tt.want {
				t.Errorf("numMatchingSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
