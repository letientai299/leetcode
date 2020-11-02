package main

import (
	"fmt"
	"testing"
)

func Test_longestWord(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{
			words: []string{"w", "wo", "wor", "wos", "word"},
			want: "word",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.words), func(t *testing.T) {
			if got := longestWord(tt.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
