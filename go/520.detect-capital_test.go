package main

import (
	"fmt"
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"USA", true},
		{"Usa", true},
		{"usa", true},
		{"UsA", false},
		{"", true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.word,
		)
		t.Run(testName, func(t *testing.T) {
			got := detectCapitalUse(tt.word)
			if got != tt.want {
				t.Errorf("detectCapitalUse(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
