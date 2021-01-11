package main

import (
	"fmt"
	"testing"
)

func Test_shortestDistance(t *testing.T) {
	tests := []struct {
		words []string
		word1 string
		word2 string
		want  int
	}{
		{
			words: []string{"a", "b", "c", "a"},
			word1: "a",
			word2: "b",
			want:  1,
		},

		{
			words: []string{"b", "b", "c", "a"},
			word1: "a",
			word2: "b",
			want:  2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v",
			tt.words, tt.word1, tt.word2,
		)
		t.Run(testName, func(t *testing.T) {
			got := shortestDistance(tt.words, tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("shortestDistance(%v, %v, %v) = %v, want %v", tt.words, tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}
