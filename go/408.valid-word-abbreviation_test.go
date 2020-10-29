package main

import (
	"fmt"
	"testing"
)

func Test_validWordAbbreviation(t *testing.T) {
	tests := []struct {
		word string
		abbr string
		want bool
	}{
		{word: "a", abbr: "01", want: false},
		{word: "word", abbr: "1ord", want: true},
		{word: "internationalization", abbr: "i12iz4n", want: true},
		{word: "word", abbr: "1o2", want: true},
		{word: "word", abbr: "2rd", want: true},
		{word: "word", abbr: "3d", want: true},
		{word: "word", abbr: "4", want: true},
		{word: "word", abbr: "4d", want: false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.word, tt.abbr,
		)
		t.Run(testName, func(t *testing.T) {
			got := validWordAbbreviation(tt.word, tt.abbr)
			if got != tt.want {
				t.Errorf("validWordAbbreviation(%v, %v) = %v, want %v", tt.word, tt.abbr, got, tt.want)
			}
		})
	}
}
