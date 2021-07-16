package main

import (
	"testing"
)

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{word1: "ab", word2: "pqrs", want: "apbqrs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.word1, tt.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
