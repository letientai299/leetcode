package main

import (
	"fmt"
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"abbabbabb", true},
		{"", false},
		{"a", false},
		{"aa", true},
		{"aba", false},
		{"abc", false},
		{"abab", true},
		{"abbabb", true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := repeatedSubstringPattern(tt.s)
			if got != tt.want {
				t.Errorf("repeatedSubstringPattern(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
