package main

import (
	"testing"
)

func Test_longestNiceSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "", want: ""},
		{s: "c", want: ""},
		{s: "zabcABCi", want: "abcABC"},
		{s: "YazaAay", want: "aAa"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestNiceSubstring(tt.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
