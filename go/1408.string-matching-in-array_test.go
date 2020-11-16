package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringMatching(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{
			words: []string{"leetcode", "leetcoder", "od", "am", "hamlet"},
			want:  []string{"am", "od", "leetcode"},
		},

		{
			words: []string{"a", "bc", "aa", "bcd"},
			want:  []string{"a", "bc"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := stringMatching(tt.words); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
