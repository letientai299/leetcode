package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeAnagrams(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{
			words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			want:  []string{"abba", "cd"},
		},
		{
			words: []string{"a", "b", "c", "d", "e"},
			want:  []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(tt.words), func(t *testing.T) {
			assert.Equalf(t, tt.want, removeAnagrams(tt.words), "removeAnagrams(%v)", tt.words)
		})
	}
}
