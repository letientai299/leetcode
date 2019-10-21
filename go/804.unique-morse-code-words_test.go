package main

import (
	"fmt"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{
			words: []string{"gin", "zen", "gig", "msg"},
			want:  2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.words,
		)
		t.Run(testName, func(t *testing.T) {
			got := uniqueMorseRepresentations(tt.words)
			if got != tt.want {
				t.Errorf("uniqueMorseRepresentations(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
