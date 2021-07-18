package main

import "testing"

func Test_canBeTypedWords(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		brokenLetters string
		want          int
	}{
		{text: "hello world", brokenLetters: "ad", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeTypedWords(tt.text, tt.brokenLetters); got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
