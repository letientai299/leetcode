package main

import (
	"fmt"
	"testing"
)

func Test_wordPattern(t *testing.T) {
	tests := []struct {
		pattern string
		str     string
		want    bool
	}{
		{pattern: "abba", str: "a b b a", want: true},
		{pattern: "abba", str: "cat dog", want: false},
		{pattern: "abba", str: "cat dog dog cat cat", want: false},
		{pattern: "abba", str: "cat dog dog cat", want: true},
		{pattern: "a", str: "something", want: true},
		{pattern: "aa", str: "something something", want: true},
		{pattern: "ab", str: "xsomething something", want: true},
		{pattern: "aa", str: "xsomething something", want: false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.pattern, tt.str,
		)
		t.Run(testName, func(t *testing.T) {
			got := wordPattern(tt.pattern, tt.str)
			if got != tt.want {
				t.Errorf("wordPattern(%v, %v) = %v, want %v", tt.pattern, tt.str, got, tt.want)
			}
		})
	}
}
