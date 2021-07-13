package main

import "testing"

func Test_maxRepeating(t *testing.T) {
	tests := []struct {
		name     string
		sequence string
		word     string
		want     int
	}{
		{sequence: "a", word: "a", want: 1},
		{sequence: "ababc", word: "ab", want: 2},
		{sequence: "ababc", word: "ba", want: 1},
		{sequence: "baaaab", word: "aa", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.sequence, tt.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
