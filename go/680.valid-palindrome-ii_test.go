package main

import (
	"fmt"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"text", true},
		{"txt", true},
		{"t123xt", false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := validPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("validPalindrome(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
