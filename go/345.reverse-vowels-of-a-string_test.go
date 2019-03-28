package main

import (
	"fmt"
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "hello", want: "holle"},
		{s: "code", want: "cedo"},
		{s: "a.b,.", want: "a.b,."},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := reverseVowels(tt.s)
			if got != tt.want {
				t.Errorf("reverseVowels(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
