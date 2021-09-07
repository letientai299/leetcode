package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cats and dog", "cat sand dog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
